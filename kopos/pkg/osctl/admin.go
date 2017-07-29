package osctl

import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/hypervisors"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	_ "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	_ "github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	// "github.com/gophercloud/gophercloud/openstack/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"

	"golang.org/x/crypto/ssh"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb"
	pbos "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack"
	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/pkg/util"
)

type AdminInfra struct {
	InfraProvider
	config       *CounterConfig
	workerconfig util.WorkerConfig
}

func Admin() *AdminInfra {
	return new(AdminInfra)
}

func (admin *AdminInfra) Credential(config *CounterConfig) *AdminInfra {
	admin.config = config
	return admin
}

func (admin *AdminInfra) Config(config *util.WorkerConfig) *AdminInfra {
	admin.workerconfig = *config
	return admin
}

func (admin *AdminInfra) BasicAuthCredential(username, password string) *AdminInfra {
	admin.username, admin.password = username, password
	return admin
}

func (admin *AdminInfra) tryto() *AdminInfra {
	var err error
	var opts gophercloud.AuthOptions

	if nil == admin.providerclient {
		if "admin" == admin.config.Username {
			if opts, err = openstack.AuthOptionsFromEnv(); nil != err {
				glog.Errorf("Could not load admin openrc: %v", err)
				admin.lasterr = err
				return admin
			} else {
				glog.Infof("Load admin openrc: %v %v", opts.IdentityEndpoint, opts.Username)
				admin.lasterr = nil
			}
		} else {
			opts = gophercloud.AuthOptions{
				IdentityEndpoint: admin.config.IdentityHost,
				Username:         admin.config.Username,
				Password:         admin.config.Password,
				DomainName:       admin.config.DomainName,
				TenantName:       admin.config.ProjectName,
			}
		}

		if admin.providerclient, err = openstack.AuthenticatedClient(opts); nil != err {
			glog.Errorf("Could not authenticate admin: %v", err)
			admin.lasterr = err
		} else {
			glog.Infof("Authenticated as token: %v", admin.providerclient.TokenID)
			admin.lasterr = nil
			admin.config.Token = admin.providerclient.TokenID
		}
		return admin
	}

	if admin.config.Token != admin.providerclient.TokenID {
		if "admin" == admin.config.Username {
			if opts, err = openstack.AuthOptionsFromEnv(); nil != err {
				glog.Errorf("Could not load admin openrc: %v", err)
				// admin.lasterr = err
				return admin
			} else {
				glog.Infof("Load admin openrc: %v %v", opts.IdentityEndpoint, opts.Username)
			}
		} else {
			opts = gophercloud.AuthOptions{
				IdentityEndpoint: admin.config.IdentityHost,
				Username:         admin.config.Username,
				Password:         admin.config.Password,
				DomainName:       admin.config.DomainName,
				TenantName:       admin.config.ProjectName,
			}
		}

		if admin.providerclient, err = openstack.AuthenticatedClient(opts); nil != err {
			glog.Errorf("Could not authenticate admin: %v", err)
			admin.lasterr = err
		} else {
			glog.Infof("Authenticated as token: %v", admin.providerclient.TokenID)
			admin.lasterr = nil
			admin.config.Token = admin.providerclient.TokenID
		}
		return admin
	}

	//	if nil != admin.providerclient.lasterr {
	//		return nil, admin.providerclient.lasterr
	//	}

	return admin
}

func (admin *AdminInfra) GetLibvirtDomainVNCDisplay(req *pbos.LibvirtDomainReqRespData) (*pbos.LibvirtDomainReqRespData, error) {
	glog.V(2).Infof("Go to get libvirt domain vnc display for server: %+v", req.ServerId)

	resp := new(pbos.LibvirtDomainReqRespData)
	resp.DomainInfo = new(pbos.LibvirtDomainInfo)
	req.DomainInfo = resp.DomainInfo
	if req == nil || len(req.ServerId) == 0 {
		resp.StateCode = 1
		resp.StateMessage = "Must specify id or name of VM"
		return resp, fmt.Errorf(resp.StateMessage)
	}

	vm, err := admin.getServer(req.ServerId)
	if nil != err {
		resp.StateCode = 2
		resp.StateMessage = err.Error()
		return resp, fmt.Errorf("Failed to call compute API of server get: %v", err)
	}
	if nil == vm {
		resp.StateCode = 3
		resp.StateMessage = "Could not find VM"
		return resp, fmt.Errorf(resp.StateMessage)
	}
	if 0 == len(vm.InstanceName) {
		resp.StateCode = 4
		resp.StateMessage = "Could not find libvirt domain name"
		return resp, fmt.Errorf(resp.StateMessage)
	}
	if 0 == len(vm.HypervisorHostname) {
		resp.StateCode = 5
		resp.StateMessage = "Could not find libvirtd name"
		return resp, fmt.Errorf(resp.StateMessage)
	}

	hvs, err := admin.getHypervisors()
	if nil != err {
		resp.StateCode = 6
		resp.StateMessage = err.Error()
		return resp, fmt.Errorf("Failed to call compute API of hypervisor list: %v", err)
	}
	var hv *hypervisors.Hypervisor = nil
	for i := 0; i < len(hvs); i++ {
		if vm.HypervisorHostname == hvs[i].HypervisorHostname {
			hv = &hvs[i]
			break
		}
	}
	if nil == hv {
		resp.StateCode = 7
		resp.StateMessage = "Could not find libvirtd host"
		return resp, fmt.Errorf(resp.StateMessage)
	}
	// if 0 == len(hv.HypervisorHostname) {
	if 0 == len(hv.HostIP) {
		resp.StateCode = 8
		resp.StateMessage = "Could not find libvirtd service"
		return resp, fmt.Errorf(resp.StateMessage)
	}
	req.DomainInfo.Name = vm.InstanceName

	if admin.workerconfig.GRPCUsed {
		// conn, err := grpc.Dial(hv.HypervisorHostname+":10010", grpc.WithInsecure())
		conn, err := grpc.Dial(hv.HostIP+":10010", grpc.WithInsecure())
		if err != nil {
			glog.V(2).Infof("Could not create gRPC connection: %v", err)
			resp.StateCode = 11
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Failed to call gRPC dial: %v", err)
		}
		defer conn.Close()
		c := pb.NewWorkerServiceClient(conn)

		r, err := c.GetLibvirtDomainInfo(context.Background(), req.DomainInfo)
		if err != nil {
			glog.V(2).Infof("Could not call gRPC server: %v", err)
			resp.StateCode = 12
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Failed to call gRPC server: %v", err)
		}
		resp.DomainInfo = r
	} else {
		sshConfig := &ssh.ClientConfig{
			User: admin.workerconfig.SSHUser.Name,
			Auth: []ssh.AuthMethod{
			// ssh.Password("your_password"),
			// PublicKeyFile("/path/to/your/pub/certificate/key")
			// SSHAgent(),
			},
		}
		if admin.workerconfig.SSHAgent {
			sshConfig.Auth = append(sshConfig.Auth, util.SSHAgent())
		}
		if len(admin.workerconfig.SSHUser.Password) != 0 {
			sshConfig.Auth = append(sshConfig.Auth, ssh.Password(admin.workerconfig.SSHUser.Password))
		} else {
			if len(admin.workerconfig.SSHUser.RSAKeyPath) != 0 {
				sshConfig.Auth = append(sshConfig.Auth, util.PublicKeyFile(admin.workerconfig.SSHUser.RSAKeyPath))
			}
		}

		client := &util.SSHClient{
			Config: sshConfig,
			Host:   hv.HostIP,
			Port:   22,
		}

		inbuf := bytes.Buffer{}
		outbuf := new(bytes.Buffer)
		errbuf := new(bytes.Buffer)
		cmd := &util.SSHCommand{
			Path:   "virsh vncdisplay " + req.DomainInfo.Name,
			Env:    []string{},
			Stdin:  &inbuf, // os.Stdin,
			Stdout: outbuf, // os.Stdout,
			Stderr: errbuf, // os.Stderr,
		}

		fmt.Printf("Running command: %s\n", cmd.Path)
		if err := client.RunCommand(cmd); err != nil {
			fmt.Fprintf(os.Stderr, "command run error: %s\n", err)
			resp.StateCode = 21
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Failed to call via SSH: %v", err)
		}
		time.Sleep(time.Second * 5)

		resultlines := []string{}
		if errbuf.Len() > 0 {
			scanner := bufio.NewScanner(errbuf)
			for scanner.Scan() {
				resultlines = append(resultlines, scanner.Text())
				resp.StateCode = 31
			}
			if err := scanner.Err(); err != nil && err != io.EOF {
				fmt.Fprintln(os.Stderr, "reading input:", err)
				resp.StateCode = 32
			}
			resp.StateMessage = "Execution failed! " + strings.Join(resultlines, "\n")
			return resp, fmt.Errorf(resp.StateMessage)
		}
		if outbuf.Len() > 0 {
			scanner := bufio.NewScanner(outbuf)
			for scanner.Scan() {
				resultlines = append(resultlines, scanner.Text())
			}
			if err := scanner.Err(); err != nil && err != io.EOF {
				fmt.Fprintln(os.Stderr, "reading input:", err)
				resp.StateCode = 0
				resp.StateMessage = "Execution incomplete! " + err.Error()
				return resp, nil
			}
		}
		if 0 != len(resultlines) {
			req.DomainInfo.Display = resultlines[0]
		}
	}

	return resp, nil
}

func (admin *AdminInfra) getServer(id string) (*Server, error) {
	glog.V(2).Infoln("Go to get server details")

	provider := admin.tryto()
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	//	resp, err := servers.Get(client, id).Extract()
	resp, err := deserializeServer(servers.Get(client, id))
	if nil != err {
		glog.Errorf("Could not reap compute interface of servers get: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute interface of servers get: %v", resp)

	return resp, nil
}

func (admin *AdminInfra) getNamedServers(name string) ([]Server, error) {
	glog.V(2).Info("Go to get server by name: %s", name)

	provider := admin.tryto()
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	changesSince := "1970-12-31T11:59:59PM"
	image_fake := "http://openstack.glance.api.endpoint/images/id"
	flavor_fake := "http://openstack.nova.api.endpoint/flavors/id"
	name_fake := "cirros"
	status_fake := "ACTIVE"
	host_fake := "all of hypervisors"
	marker_fake := "00000000-0000-0000-0000-000000000000"
	limit_fake := 0
	allTenants := false

	opts := servers.ListOpts{
		ChangesSince: changesSince,
		Image:        image_fake,
		Flavor:       flavor_fake,
		Name:         name_fake,
		Status:       status_fake,
		Host:         host_fake,
		Marker:       marker_fake,
		Limit:        limit_fake,
		AllTenants:   allTenants,
	}
	opts = servers.ListOpts{
		Name: name,
	}
	resp, err := servers.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap compute interface of servers list: %v", err)
		return nil, err
	}

	// result, err := servers.ExtractServers(resp)
	result, err := deserializeServers(resp)
	if nil != err {
		glog.Errorf("Could not reap compute interface of servers list: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute interface of servers list: %v", result)

	return result, nil
}

/*
{
    "hypervisor": {
        "cpu_info": {
            "arch": "x86_64",
            "model": "Nehalem",
            "vendor": "Intel",
            "features": [
                "pge",
                "clflush"
            ],
            "topology": {
                "cores": 1,
                "threads": 1,
                "sockets": 4
            }
        },
        "state": "up",
        "status": "enabled",
        "current_workload": 0,
        "disk_available_least": 0,
        "host_ip": "1.1.1.1",
        "free_disk_gb": 1028,
        "free_ram_mb": 7680,
        "hypervisor_hostname": "fake-mini",
        "hypervisor_type": "fake",
        "hypervisor_version": 1000,
        "id": 1,
        "local_gb": 1028,
        "local_gb_used": 0,
        "memory_mb": 8192,
        "memory_mb_used": 512,
        "running_vms": 0,
        "service": {
            "host": "043b3cacf6f34c90a7245151fc8ebcda",
            "id": 2,
            "disabled_reason": null
        },
        "vcpus": 1,
        "vcpus_used": 0
    }
}
*/
func (admin *AdminInfra) getHypervisors() ([]hypervisors.Hypervisor, error) {
	glog.V(2).Infoln("Go to get hypervisors detailed")

	provider := admin.tryto()
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	resp, err := hypervisors.List(client).AllPages()
	if nil != err {
		glog.Errorf("Could not reap compute interface of hypervisor list: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute interface of hypervisor list: %v", resp)

	result, err := hypervisors.ExtractHypervisors(resp)
	if nil != err {
		glog.Errorf("Could not reap compute interface of hypervisor list: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute interface of hypervisor list: %v", result)

	return result, nil
}

type User struct {
	Name string
	Age  int
}

func pq_tut() {
	/*
		 A backslash will escape the next character in values:
		"user=space\ man password='it\'s valid'
	*/
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	// db, err := sql.Open("postgres", "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	println(rows)

	rows, err = db.Query(`SELECT name FROM users WHERE favorite_fruit = $1
	OR age BETWEEN $2 AND $2 + 3`, "orange", 64)
	println(rows)

	/*
	   pq does not support the LastInsertId() method of the Result type in database/sql.
	   To return the identifier of an INSERT (or UPDATE or DELETE),
	   use the Postgres RETURNING clause with a standard Query or QueryRow call:
	*/
	var userid int
	err = db.QueryRow(`INSERT INTO users(name, favorite_fruit, age)
	VALUES('beatrice', 'starfruit', 93) RETURNING id`).Scan(&userid)

	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
	}

	// Bulk imports
	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := txn.Prepare(pq.CopyIn("users", "name", "age"))
	if err != nil {
		log.Fatal(err)
	}

	users := []User{{"foor", 19}, {"bar", 20}}

	for _, user := range users {
		_, err = stmt.Exec(user.Name, int64(user.Age))
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

type Account struct {
	ID      int `gorm:"primary_key"`
	Balance int
}

func gorm_tut() {
	// ORM
	const addr = "postgresql://maxroach@localhost:26257/bank?sslmode=disable"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Automatically create the "accounts" table based on the Account model.
	db.AutoMigrate(&Account{})

	// Insert two rows into the "accounts" table.
	db.Create(&Account{ID: 1, Balance: 1000})
	db.Create(&Account{ID: 2, Balance: 250})

	// Print out the balances.
	var accounts []Account
	db.Find(&accounts)
	fmt.Println("Initial balances:")
	for _, account := range accounts {
		fmt.Printf("%d %d\n", account.ID, account.Balance)
	}
}

func (admin *AdminInfra) MockSSH(req *pb.SSHReqRespData) (*pb.SSHReqRespData, error) {
	glog.V(2).Infof("Go to mock SSH serving: %+v, +%v", req.Cmd, req.Env)

	resp := new(pb.SSHReqRespData)
	hostIP := "172.17.4.50"
	if v, ok := os.LookupEnv("MOCK_HOST"); ok {
		hostIP = v
	}
	if 0 == len(req.Cmd) && 0 == len(req.Env) && 0 == len(req.Result) {
		resp.Result = []string{util.Pub_vagrant}
		return resp, nil
	}

	sshConfig := &ssh.ClientConfig{
		User: admin.workerconfig.SSHUser.Name,
		Auth: []ssh.AuthMethod{
		// ssh.Password("your_password"),
		// PublicKeyFile("/path/to/your/pub/certificate/key")
		// SSHAgent(),
		},
	}
	if admin.workerconfig.SSHAgent {
		fmt.Printf("Using sshagent, %s\n", admin.workerconfig.SSHUser.Name)
		sshConfig.Auth = append(sshConfig.Auth, util.SSHAgent())
	}
	if len(admin.workerconfig.SSHUser.Password) != 0 {
		fmt.Printf("Using basic auth: %s, %s\n", admin.workerconfig.SSHUser.Name, strings.Repeat("*", len(admin.workerconfig.SSHUser.Password)))
		sshConfig.Auth = append(sshConfig.Auth, ssh.Password(admin.workerconfig.SSHUser.Password))
	} else {
		if len(admin.workerconfig.SSHUser.RSAKeyPath) != 0 {
			fmt.Printf("Using rsa private key: %s, %s\n", admin.workerconfig.SSHUser.Name, admin.workerconfig.SSHUser.RSAKeyPath)
			sshConfig.Auth = append(sshConfig.Auth, util.PublicKeyFile(admin.workerconfig.SSHUser.RSAKeyPath))
		}
	}

	//		connection, err := ssh.Dial("tcp", hv.HostIP+":22", sshConfig)
	//		if err != nil {
	//			return nil, fmt.Errorf("Failed to dial: %s", err)
	//		}
	client := &util.SSHClient{
		Config: sshConfig,
		Host:   hostIP,
		Port:   22,
	}

	inbuf := bytes.Buffer{}
	outbuf := new(bytes.Buffer)
	errbuf := new(bytes.Buffer)
	cmd := &util.SSHCommand{
		Path:   req.Cmd,
		Env:    req.Env,
		Stdin:  &inbuf, // os.Stdin,
		Stdout: outbuf, // os.Stdout,
		Stderr: errbuf, // os.Stderr,
	}

	fmt.Printf("Running command: %s\n", cmd.Path)
	if err := client.RunCommand(cmd); err != nil {
		fmt.Fprintf(os.Stderr, "command run error: %s\n", err)
		return resp, fmt.Errorf("Failed to execute via SSH: %v", err)
	}
	time.Sleep(time.Second * 5)
	if errbuf.Len() > 0 {
		scanner := bufio.NewScanner(errbuf)
		for scanner.Scan() {
			resp.Result = append(resp.Result, scanner.Text())
		}
		if err := scanner.Err(); err != io.EOF {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		return resp, fmt.Errorf("Faild to execute into stderr")
	}
	if outbuf.Len() > 0 {
		scanner := bufio.NewScanner(outbuf)
		for scanner.Scan() {
			resp.Result = append(resp.Result, scanner.Text())
		}
		if err := scanner.Err(); err != io.EOF {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		return resp, nil
	}
	return resp, fmt.Errorf("Nothing output")
}

func (admin *AdminInfra) CreateSharedNet(name, tenandid, subnetcidr, gatewayip, description string) {
	adminstateup := true
	shared := true
	network, err := admin.tryto().createNetwork(name, tenandid, adminstateup, shared)
	if nil != err {
		return
	}

	offset := strings.LastIndex(subnetcidr, ".")
	ipprefix := subnetcidr[:offset]
	ipversion := gophercloud.IPv4
	enabledhcp := true
	dnsnameservers := []string{}
	hostroutes := []subnets.HostRoute{}
	subnet, err := admin.tryto().createSubnet(network.ID, subnetcidr, name, tenandid, []subnets.AllocationPool{{ipprefix + ".50", ipprefix + ".200"}}, gatewayip, ipversion, enabledhcp, dnsnameservers, hostroutes)
	if nil != err {
		return
	}
	println(subnet)

	distributed := false
	gatewayinfo := routers.GatewayInfo{}
	router, err := admin.tryto().createRouter(name, adminstateup, distributed, tenandid, gatewayinfo)
	if nil != err {
		return
	}

	secgroup, err := admin.tryto().createSecurityGroup(name, tenandid, description)
	if nil != err {
		return
	}

	// remoteipprefix = "::/0"

	securitygroups := []string{}
	portrangemax := 65535
	portrangemin := 0
	protocol := rules.ProtocolTCP
	remotegroupid := ""
	remoteipprefix := "0.0.0.0/0"
	secgrouprule, err := admin.tryto().createSecGroupRule(rules.DirIngress, rules.EtherType4, secgroup.ID, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenandid)
	if nil != err {
		return
	}
	securitygroups = append(securitygroups, secgrouprule.ID)

	protocol = rules.ProtocolUDP
	secgrouprule, err = admin.tryto().createSecGroupRule(rules.DirIngress, rules.EtherType4, secgroup.ID, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenandid)
	if nil != err {
		return
	}
	securitygroups = append(securitygroups, secgrouprule.ID)

	protocol = rules.ProtocolICMP
	portrangemax = -1
	portrangemin = -1
	secgrouprule, err = admin.tryto().createSecGroupRule(rules.DirIngress, rules.EtherType4, secgroup.ID, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenandid)
	if nil != err {
		return
	}
	securitygroups = append(securitygroups, secgrouprule.ID)

	macaddress := ""
	fixedips := []ports.IP{}
	deviceid := ""
	deviceowner := ""
	allowedaddresspairs := []ports.AddressPair{}
	port, err := admin.tryto().createPort(network.ID, name, adminstateup, macaddress, fixedips, deviceid, deviceowner, tenandid, securitygroups, allowedaddresspairs)
	if nil != err {
		return
	}

	ifinfo, err := admin.tryto().plugRouterIntoSubnet(router.ID, port.ID)
	if nil != err {
		return
	}
	println(ifinfo)
}
