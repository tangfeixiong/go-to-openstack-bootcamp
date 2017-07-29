/*
   Inspired from
     https://github.com/openshift/origin/blob/master/pkg/cmd/openshift/openshift.go
     https://stackoverflow.com/questions/34053881/golang-how-can-i-use-pflag-with-other-packages-that-use-flag
*/
package kopin

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	// "github.com/google/gopacket/dumpcommand"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	dumpcommand "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/gopaketctl"
	logutil "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util"
)

var (
	/*tmplLong = templates.LongDesc*/ tmplLong string = (`
		%[2]s

		The %[3]s helps you build, deploy, and manage your applications on top of
		Docker containers. To start an all-in-one server with the default configuration, run:

		    $ %[1]s start &`)
	logger = logutil.Logger
)

type Option struct {
	autocap_ifaces []string
	iface          *string
	snaplen        *int
	tstype         *string
	promisc        *bool
	handles        []*pcap.Handle
}

func buildrun(opt *Option, out io.Writer) func(*cobra.Command, []string) {
	var wg sync.WaitGroup

	var iface = flag.String("i", "eth1", "Interface to read packets from")
	var fname = flag.String("r", "", "Filename to read from, overrides -i")
	var snaplen = flag.Int("s", 65536, "Snap length (number of bytes max to read per packet")
	var tstype = flag.String("timestamp_type", "", "Type of timestamps to use")
	var promisc = flag.Bool("promisc", true, "Set promiscuous mode")

	return func(c *cobra.Command, args []string) {
		c.SetOutput(out)

		if 0 != len(*iface) {
			go func() {
				wg.Add(1)
				defer wg.Done()
				/*
				   Inspired from
				     https://github.com/google/gopacket/bolb/master/examples/pcapdump/main.go
				*/
				defer util.Run()()
				var handle *pcap.Handle
				var err error
				if *fname != "" {
					if handle, err = pcap.OpenOffline(*fname); err != nil {
						log.Fatal("PCAP OpenOffline error:", err)
					}
					defer handle.Close()
					opt.handles = append(opt.handles, handle)
				} else {
					// This is a little complicated because we want to allow all possible options
					// for creating the packet capture handle... instead of all this you can
					// just call pcap.OpenLive if you want a simple handle.
					inactive, err := pcap.NewInactiveHandle(*iface)
					if err != nil {
						// log.Fatalf("could not create: %v", err)
						logger.Fatalf("could not create: %v", err)
					}
					defer inactive.CleanUp()
					if err = inactive.SetSnapLen(*snaplen); err != nil {
						// log.Fatalf("could not set snap length: %v", err)
						logger.Fatalf("could not set snap length: %v", err)
					} else if err = inactive.SetPromisc(*promisc); err != nil {
						// log.Fatalf("could not set promisc mode: %v", err)
						logger.Fatalf("could not set promisc mode: %v", err)
					} else if err = inactive.SetTimeout(time.Second * 3); err != nil {
						// log.Fatalf("could not set timeout: %v", err)
						logger.Fatalf("could not set timeout: %v", err)
					}
					if *tstype != "" {
						if t, err := pcap.TimestampSourceFromString(*tstype); err != nil {
							// log.Fatalf("Supported timestamp types: %v", inactive.SupportedTimestamps())
							logger.Fatalf("Supported timestamp types: %v", inactive.SupportedTimestamps())
						} else if err := inactive.SetTimestampSource(t); err != nil {
							// log.Fatalf("Supported timestamp types: %v", inactive.SupportedTimestamps())
							logger.Fatalf("Supported timestamp types: %v", inactive.SupportedTimestamps())
						}
					}
					if handle, err = inactive.Activate(); err != nil {
						// log.Fatal("PCAP Activate error:", err)
						logger.Fatal("PCAP Activate error:", err)
					}
					defer handle.Close()

					opt.handles = append(opt.handles, handle)
					time.Sleep(time.Second)
				}

				if len(flag.Args()) > 0 {
					bpffilter := strings.Join(flag.Args(), " ")
					fmt.Fprintf(os.Stderr, "Using BPF filter %q\n", bpffilter)
					if err = handle.SetBPFFilter(bpffilter); err != nil {
						// log.Fatal("BPF filter error:", err)
						logger.Fatal("BPF filter error:", err)
					}
				}
				err = dumpcommand.Run(handle)
				if err != nil {
					logger.Fatal(err)
				}
			}()
		}

		// c.Help()
		serve(c, args)
		wg.Wait()
	}
}

func NewRootCommand(name string) *cobra.Command {
	var opt Option
	// in, out, errout := os.Stdin, os.Stdout, os.Stderr

	root := &cobra.Command{
		Use:   name,
		Short: "Build, deploy, and manage your cloud applications",
		Long:  fmt.Sprintf(tmplLong, name, name, name), // fmt.Sprintf(openshiftLong, name, cmdutil.GetPlatformName(name), cmdutil.GetDistributionName(name)),
		Run:   buildrun(&opt, os.Stdout),               // kcmdutil.DefaultSubCommandRun(out),
	}
	//	root.AddCommand(pcapCommand(&opt))
	//	root.AddCommand(libvirtCommand(&opt))
	//	root.AddCommand(dockerCommand(&opt))
	//	root.AddCommand(linuxserverCommand(&opt))

	/* Two approach for this. One with pflags AddGoFlags()*/
	/*
		f := pflag.NewFlagSet("goFlags", pflag.ExitOnError)
		f.AddGoFlagSet(flag.CommandLine)
		f.Parse([]string{
			"--v=2", // Your go flags that will be pass to other programs.
		})
		flag.Parse()
	*/
	/* Another approach you can use. This is not by the book but i have seen some used this. Get the flag valu by another pflag.*/
	/*
	   pflag.StringVar(&f, "f", "**", "***")
	   pflag.Parse()
	*/
	/* Set the value as flag value.*/
	/*
	   flag.Set("v", f)
	   flag.Parse()
	*/
	/* And you are good to go */
	//  root.Flags().StringSliceVar(&opt.autocap_ifaces, "autocap-ifaces", []string{}, "interfaces to auto read packets from")
	//	if 0 != len(opt.autocap_ifaces) {
	//		flag.Set("iface", opt.autocap_ifaces[0])
	//	}
	//	loglevel := root.LocalNonPersistentFlags().Int32("loglevel", 2, "Set the level of log output (0-5)")
	//	flag.Set("v", strconv.Itoa(int(*loglevel)))
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()

	//	f := clientcmd.New(pflag.NewFlagSet("", pflag.ContinueOnError))

	//	startAllInOne, _ := start.NewCommandStartAllInOne(name, out, errout)
	//	root.AddCommand(startAllInOne)
	//	root.AddCommand(admin.NewCommandAdmin("admin", name+" admin", in, out, errout))
	//	root.AddCommand(cli.NewCommandCLI("cli", name+" cli", in, out, errout))
	//	root.AddCommand(cli.NewCmdKubectl("kube", out))
	//	root.AddCommand(newExperimentalCommand("ex", name+" ex"))
	//	root.AddCommand(newCompletionCommand("completion", name+" completion"))
	//	root.AddCommand(cmd.NewCmdVersion(name, f, out, cmd.VersionOptions{PrintEtcdVersion: true, IsServer: true}))

	//	// infra commands are those that are bundled with the binary but not displayed to end users
	//	// directly
	//	infra := &cobra.Command{
	//		Use: "infra", // Because this command exposes no description, it will not be shown in help
	//	}

	//	infra.AddCommand(
	//		irouter.NewCommandTemplateRouter("router"),
	//		irouter.NewCommandF5Router("f5-router"),
	//		deployer.NewCommandDeployer("deploy"),
	//		recycle.NewCommandRecycle("recycle", out),
	//		builder.NewCommandS2IBuilder("sti-build"),
	//		builder.NewCommandDockerBuilder("docker-build"),
	//		diagnostics.NewCommandPodDiagnostics("diagnostic-pod", out),
	//		diagnostics.NewCommandNetworkPodDiagnostics("network-diagnostic-pod", out),
	//	)
	//	root.AddCommand(infra)

	//	root.AddCommand(cmd.NewCmdOptions(out))

	//	// TODO: add groups
	//	templates.ActsAsRootCommand(root, []string{"options"})

	return root
}

func pcapCommand(opt *Option) *cobra.Command {

	thecmd := &cobra.Command{
		Use:   "pcap",
		Short: "Using libpcap to capture packets from Linux network stack",
		Long: ` The pcap client implementing
        
        To be continued
        `,
		Run: buildrun(opt, os.Stdout),
	}

	return thecmd
}

func libvirtCommand(opt *Option) *cobra.Command {

	thecmd := &cobra.Command{
		Use:   "libvirt",
		Short: "Using libvirt to manage vms from Linux daemon",
		Long: ` The libvirtd client implementing
        
        To be continued
        `,
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("N/A, capability is planned in future.")
		},
	}

	return thecmd
}

func dockerCommand(opt *Option) *cobra.Command {

	thecmd := &cobra.Command{
		Use:   "docker",
		Short: "Using docker client (docker compose client) to manage running containers from Linux daemon",
		Long: ` The docker client implementing
        
        To be continued
        `,
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("N/A, capability is planned in future.")
		},
	}

	return thecmd
}

func linuxserverCommand(opt *Option) *cobra.Command {

	thecmd := &cobra.Command{
		Use:   "linuxserver",
		Short: "To manage Linux server",
		Long: ` The mgmt client implementing
        
        To be continued
        `,
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("N/A, capability is planned in future.")
		},
	}

	return thecmd
}
