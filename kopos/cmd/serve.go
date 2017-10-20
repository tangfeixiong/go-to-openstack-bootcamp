// Inspired from https://github.com/philips/grpc-gateway-example

package cmd

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/philips/go-bindata-assetfs"
	"github.com/spf13/cobra"
	// "github.com/spf13/pflag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb"
	pbos "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack"
	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/pkg/osctl"
	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/pkg/ui/data/swagger"
	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/pkg/util"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches the example webserver on https://localhost:10000",
	Run: func(cmd *cobra.Command, args []string) {
		//		seed := pflag.Bool("seed", false, "Only pretend to seed.")
		//		pflag.Lookup("seed").NoOptDefVal = "false"
		//		if *seed {
		//			go func() {
		//				seedinit()
		//				seedmain()
		//			}()
		//		}

		go func() {
			wg.Add(1)
			defer wg.Done()
			vulnerable(0)
		}()
		go func() {
			wg.Add(1)
			defer wg.Done()
			vulnerable(1)
		}()

		serve()
		wg.Wait()
	},
}

var (
	wg         sync.WaitGroup
	workconfig util.WorkerConfig
)

func init() {
	// bridge glog with pflag
	GLog(RootCmd.PersistentFlags())
	RootCmd.AddCommand(serveCmd)
	// "/Users/fanhongling/Downloads/tmp"
	path := filepath.Join("/Users", "fanhongling", "Downloads", "tmp")
	serveCmd.Flags().BoolVar(&workconfig.GRPCUsed, "grpcworker", false, "using gRPC worker service")
	serveCmd.Flags().BoolVar(&workconfig.SSHAgent, "sshagent", false, "using ssh agent")
	serveCmd.Flags().StringVarP(&workconfig.SSHUser.Name, "sshusername", "u", "root", "ssh user")
	serveCmd.Flags().StringVarP(&workconfig.SSHUser.Password, "sshpassword", "p", "", "ssh password")
	serveCmd.Flags().StringVarP(&workconfig.SSHUser.RSAKeyPath, "sshrsapubpath", "k", path, "ssh private key, like $HOME/.ssh/id_rsa")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0600); nil != err {
			fmt.Printf("Not writen init: %v\n", err)
		}
	}
	pub := []byte(util.Key_vagrant)
	path = filepath.Join(path, "vagrant")
	if err := ioutil.WriteFile(path, pub, 0600); nil != err {
		fmt.Printf("Not writen init: %v\n", err)
	}
	// f, err := os.Create(path)
	// if nil != err {
	//	 fmt.Printf("Not writen init: %v\n", err)
	// }
	// defer f.Close()
	// if _, err := f.Write(pub); nil != err {
	//	 fmt.Printf("Not writen init: %v\n", err)
	// }
}

type myService struct {
	config       *osctl.CounterConfig
	workerconfig *util.WorkerConfig
}

func (m *myService) DiscoverNetworks(ctx context.Context, in *pbos.NetworkDiscoveryReqRespData) (*pbos.NetworkDiscoveryReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverNetworks(in)
}

func (m *myService) DiscoverSubnets(ctx context.Context, in *pbos.SubnetDiscoveryReqRespData) (*pbos.SubnetDiscoveryReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverSubnets(in)
}

func (m *myService) DiscoverNetworkingTopology(ctx context.Context, req *pbos.NetworkTopologyReqRespData) (*pbos.NetworkTopologyReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverNetworkingTopology(req)
}

func (m *myService) EstablishNetworkLandscape(ctx context.Context, req *pbos.OpenstackNeutronLandscapeReqRespData) (*pbos.OpenstackNeutronLandscapeReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).CreateNetworkingLandscape(req)
}

func (m *myService) DiscoverImages(ctx context.Context, in *pbos.ImageDiscoveryReqRespData) (*pbos.ImageDiscoveryReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverImages(in)
}

func (m *myService) DiscoverImageDetailed(ctx context.Context, in *pbos.Image) (*pbos.Image, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverImageDetailed(in)
}

func (m *myService) SearchImageDetails(ctx context.Context, in *pbos.Image) (*pbos.Image, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverImageDetails(in)
}

func (m *myService) DiscoverFlavors(ctx context.Context, in *pbos.FlavorDiscoveryReqRespData) (*pbos.FlavorDiscoveryReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverFlavors(in)
}

func (m *myService) DiscoverFlavorDetailed(ctx context.Context, in *pbos.Flavor) (*pbos.Flavor, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverFlavorDetailed(in)
}

func (m *myService) SearchFlavorDetails(ctx context.Context, in *pbos.Flavor) (*pbos.Flavor, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverFlavorDetails(in)
}

func (m *myService) SpawnMachines(ctx context.Context, in *pbos.MachineSpawnsReqRespData) (*pbos.MachineSpawnsReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).SpawnMachines(in)
}

func (m *myService) DiscoverMachines(ctx context.Context, in *pbos.MachineDiscoveryReqRespData) (*pbos.MachineDiscoveryReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).DiscoverMachines(in)
}

func (m *myService) DestroyMachines(ctx context.Context, in *pbos.MachineDestroyReqRespData) (*pbos.MachineDestroyReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).DestroyMachines(in)
}

func (m *myService) RebootMachines(ctx context.Context, in *pbos.MachineRebootReqRespData) (*pbos.MachineRebootReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).RebootMachines(in)
}

func (m *myService) GetLibvirtDomainVNCDisplay(ctx context.Context, in *pbos.LibvirtDomainReqRespData) (*pbos.LibvirtDomainReqRespData, error) {
	return osctl.Admin().Credential(m.config).Config(m.workerconfig).GetLibvirtDomainVNCDisplay(in)
}

func (m *myService) BootVirtualMachines(ctx context.Context, in *pbos.OpenstackNovaBootReqRespData) (*pbos.OpenstackNovaBootReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).BootVirtualMachines(in)
}

func (m *myService) Echo(c context.Context, s *pb.EchoMessage) (*pb.EchoMessage, error) {
	fmt.Printf("rpc request Echo(%q)\n", s.Value)
	return s, nil
}

func (m *myService) ValidateToken(ctx context.Context, req *pbos.TokenReqRespData) (*pbos.TokenReqRespData, error) {
	return osctl.InfraRes().Credential(m.config).ValidateToken(req)
}

func (m *myService) MockSSH(ctx context.Context, req *pb.SSHReqRespData) (*pb.SSHReqRespData, error) {
	return osctl.Admin().Credential(m.config).Config(m.workerconfig).MockSSH(req)
}

func (m *myService) AdminSharedNetworkCreation(ctx context.Context, req *pbos.OpenstackNeutronNetRequestData) (*pbos.OpenstackNeutronNetResponseData, error) {
	osctl.Admin().CreateSharedNet("Name", "TenantID", "subnetCIDR", "GatewayIP", "Description")
	return new(pbos.OpenstackNeutronNetResponseData), nil
}

func (m *myService) ApplyConsoleIntoDnatWithNetworkAndMachine(ctx context.Context, req *pbos.ConsoleResourceRequestData) (*pbos.ConsoleResourceResponseData, error) {
	fmt.Printf("rpc ApplyConsoleIntoDnatWithNetworkAndMachine(%v)\n", req)
	return new(pbos.ConsoleResourceResponseData), nil
}

func (m *myService) OrderTargetDroneIntoTrainee(ctx context.Context, req *pbos.TraineeDroneRequestData) (*pbos.TraineeDroneResponseData, error) {
	fmt.Printf("rpc OrderTargetDroneIntoTrainee(%v)\n", req)
	return new(pbos.TraineeDroneResponseData), nil
}

func (m *myService) OrderTargetDroneIntoDefenseFortification(ctx context.Context, req *pbos.DefensiveDroneRequestData) (*pbos.DefensiveDroneResponseData, error) {
	fmt.Printf("rpc OrderTargetDroneIntoDefenseFortification(%v)\n", req)
	return new(pbos.DefensiveDroneResponseData), nil
}

func newServer() *myService {
	hvusername := os.Getenv("HV_USERNAME")
	hvpassword := os.Getenv("HV_PASSWORD")
	hvkeypath := os.Getenv("HV_KEYPATH")
	// "/Users/fanhongling/Downloads/tmp/vagrant"
	path := filepath.Join("/Users", "fanhongling", "Downloads", "tmp", "vagrant")
	if 0 != len(hvusername) && hvusername != "fake user" && hvusername != workconfig.SSHUser.Name {
		workconfig.SSHUser.Name = hvusername
	}
	if 0 != len(hvpassword) && hvpassword != "fake secret" && 0 == len(workconfig.SSHUser.Password) {
		workconfig.SSHUser.Password = hvpassword
	}
	if 0 != len(hvkeypath) && hvkeypath != path && 0 == len(workconfig.SSHUser.RSAKeyPath) {
		workconfig.SSHUser.RSAKeyPath = hvkeypath
	}

	// return new(myService)
	return &myService{
		config:       osctl.NewCounterConfig(),
		workerconfig: &workconfig,
	}
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO(tamird): point to merged gRPC code rather than a PR.
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func serveSwagger(mux *http.ServeMux) {
	mime.AddExtensionType(".svg", "image/svg+xml")

	// Expose files in third_party/swagger-ui/ on <host>/swagger-ui
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

func serve() {
	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000"))}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEchoServiceServer(grpcServer, newServer())
	ctx := context.Background()

	dcreds := credentials.NewTLS(&tls.Config{
		ServerName: demoAddr,
		RootCAs:    demoCertPool,
	})
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(dcreds)}

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		io.Copy(w, strings.NewReader(pb.Swagger))
	})

	gwmux := runtime.NewServeMux()
	err := pb.RegisterEchoServiceHandlerFromEndpoint(ctx, gwmux, demoAddr, dopts)
	if err != nil {
		fmt.Printf("serve: %v\n", err)
		return
	}

	mux.Handle("/", gwmux)
	serveSwagger(mux)

	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:    demoAddr,
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*demoKeyPair},
			NextProtos:   []string{"h2"},
		},
	}

	fmt.Printf("grpc on port: %d\n", port)
	err = srv.Serve(tls.NewListener(conn, srv.TLSConfig))

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	return
}

func vulnerable(index int) {
	hosts := []string{"localhost:10002", ":10001"}
	if v, ok := os.LookupEnv("GRPC_PORT_OSBOOTCAMP"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			hosts[0] = v
		} else {
			hosts[0] = "localhost:" + v
		}
	}
	if v, ok := os.LookupEnv("HTTP_PORT_OSBOOTCAMP"); ok && 0 != len(v) {
		if strings.Contains(v, ":") {
			hosts[1] = v
		} else {
			hosts[1] = ":" + v
		}
	}

	switch index {
	case 0:
		s := grpc.NewServer()
		pb.RegisterEchoServiceServer(s, newServer())

		l, err := net.Listen("tcp", hosts[0])
		if err != nil {
			panic(err)
		}

		fmt.Printf("grpc on host: %s\n", hosts[0])
		if err := s.Serve(l); nil != err {
			panic(err)
		}
	default:
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := http.NewServeMux()

		dopts := []grpc.DialOption{grpc.WithInsecure()}

		gwmux := runtime.NewServeMux()
		if err := pb.RegisterEchoServiceHandlerFromEndpoint(ctx, gwmux, hosts[0], dopts); err != nil {
			fmt.Printf("serve: %v\n", err)
			return
		}

		mux.Handle("/", gwmux)

		fmt.Printf("http on host: %s\n", hosts[1])
		if err := http.ListenAndServe(hosts[1], allowCORS(mux)); nil != err {
			fmt.Fprintf(os.Stderr, "Server died: %s\n", err)
		}

		//	lstn, err := net.Listen("tcp", hosts[1])
		//	if nil != err {
		//		panic(err)
		//	}

		//	srv := &http.Server{
		//		Addr: hosts[1],
		//		// Handler: grpcHandlerFunc(grpcServer, mux),
		//		Handler: mux,
		//	}

		//	if err := srv.Serve(lstn); nil != err {
		//		fmt.Fprintf(os.Stderr, "Server died: %s\n", err)
		//	}
	}
}

// allowCORS allows Cross Origin Resoruce Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

var swaggerDir string = "examples/examplepb"

func serveSwagger2(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, ".swagger.json") {
		// glog.Errorf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	// glog.Infof("Serving %s", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join(swaggerDir, p)
	http.ServeFile(w, r, p)
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	// glog.Infof("preflight request for %s", r.URL.Path)
	return
}
