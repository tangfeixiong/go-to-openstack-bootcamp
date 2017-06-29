// tangfeixiong <tangfx128@gmail.com>

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"k8s.io/kubernetes/pkg/util/rand"

	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/cmd/sapcc-koposcmd"
	pb "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb"
	pbos "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Example test gRPC service CLI client",
	Run: func(cmd *cobra.Command, args []string) {
		var insecure bool
		pflag.BoolVar(&insecure, "insecure", true, "using http.")
		pflag.Parse()
		if insecure {
			fmt.Println("http...")
			url := "http://localhost:10001/v1/battlefields"

			var netTransport = &http.Transport{
				Dial: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).Dial,
				// TLSHandshakeTimeout: 5 * time.Second,
			}
			var netClient = &http.Client{
				Timeout:   time.Second * 10,
				Transport: netTransport,
			}

			in, err := json.Marshal(&pbos.OpenstackNeutronNetRequestData{Name: "test"})
			if err != nil {
				panic(err)
			}

			response, err := netClient.Post(url, "application/json", bytes.NewBuffer(in))
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			fmt.Println("response Status:", response.Status)
			fmt.Println("response Headers:", response.Header)
			respbody, _ := ioutil.ReadAll(response.Body)
			fmt.Println("response Body:", string(respbody))

			var jsonStr = []byte(`{"name": "again"}`)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			req.Header.Set("X-Custom-Header", "test again")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			fmt.Println("response Status:", resp.Status)
			fmt.Println("response Headers:", resp.Header)
			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println("response Body:", string(body))

			return
		}

		println("grpc with tls")
		var opts []grpc.DialOption
		creds := credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000")
		opts = append(opts, grpc.WithTransportCredentials(creds))
		conn, err := grpc.Dial(demoAddr, opts...)
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := pb.NewEchoServiceClient(conn)

		// msg, err := client.Echo(context.Background(), &pb.EchoMessage{strings.Join(os.Args[2:], " ")})
		// println(msg.Value)
		println(strings.Join(os.Args[1:], " "))

		in := &pbos.OpenstackNeutronNetRequestData{Name: "test"}
		// copts := []grpc.CallOption{grpc.EmptyCallOption{}}
		copts := []grpc.CallOption{}
		resp, err := client.AdminSharedNetworkCreation(context.Background(), in, copts...)
		println(resp)
	},
}

func init() {
	testCmd.AddCommand(koposcmd.SapccKoposCmd)
	testCmd.AddCommand(tokenCommand, bootCommand(), netCommand())
	RootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoCmd.Flags().String("foo", "", "A help for foo")

	testCmd.PersistentFlags().BoolP("insecure", "n", true, "using insecure rpc/http")
	tokenCommand.Flags().BoolP("http", "t", true, "using insecure http")
}

var tokenCommand = &cobra.Command{
	Use:   "token",
	Short: "Example token gRPC service CLI client",
	Run: func(cmd *cobra.Command, args []string) {
		var insecure, ishttp bool
		pflag.BoolVar(&insecure, "insecure", true, "using insecure rpc/http.")
		pflag.BoolVar(&ishttp, "http", true, "using http.")
		pflag.Parse()

		if insecure {
			if ishttp {
				var url = "http://localhost:10001/v1/token"
				var jsonStr = []byte(`{"src": null, "tgt": null}`)
				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
				req.Header.Set("X-Custom-Header", "test again")
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					panic(err)
				}
				defer resp.Body.Close()

				fmt.Println("response Status:", resp.Status)
				fmt.Println("response Headers:", resp.Header)
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("response Body:", string(body))
				return
			}

			conn, err := grpc.Dial("localhost:10002", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := pb.NewEchoServiceClient(conn)

			// Contact the server and print out its response.
			if len(os.Args) > 1 {
				log.Printf("args: %+v", os.Args[1:])
			}
			r, err := c.ValidateToken(context.Background(), &pbos.TokenReqRespData{})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %v", r)
			return
		}

		var opts []grpc.DialOption
		creds := credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000")
		opts = append(opts, grpc.WithTransportCredentials(creds))
		conn, err := grpc.Dial(demoAddr, opts...)
		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}
		defer conn.Close()
		client := pb.NewEchoServiceClient(conn)

		println(strings.Join(os.Args[1:], " "))

		in := &pbos.TokenReqRespData{}
		// copts := []grpc.CallOption{grpc.EmptyCallOption{}}
		copts := []grpc.CallOption{}
		resp, err := client.ValidateToken(context.Background(), in, copts...)
		if err != nil {
			grpclog.Printf("fail to greet: %v", err)
		}
		println(resp)

	},
}

func bootCommand() *cobra.Command {
	bootcmd := &cobra.Command{
		Use:   "boot",
		Short: "Example boot gRPC service CLI client",
		Run: func(cmd *cobra.Command, args []string) {
			var insecure, ishttp bool
			pflag.BoolVar(&insecure, "insecure", true, "using insecure rpc/http.")
			pflag.BoolVar(&ishttp, "http", true, "using http.")
			pflag.Parse()

			if insecure {
				if ishttp {
					var url = "http://localhost:10001/v1/boot"
					var jsonStr = []byte(strings.Replace(`
{
  "flavor_name": "m1.tiny",
  "image_name": "cirros",
  "min_count": 2,
  "max_count": 4,
  "secgroups_info": [
    {
      "id": "15895b59-9aa7-4e6e-9d23-75a9324030f5",
      "name": "default"
    }
  ],
  "user_data": [],
  "network_name": "private-admin1",
  "floating_network_name": "public",
  "personality": [],
  "name_prefix": "............"
}
`, "............", rand.String(12), 1))
					req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
					req.Header.Set("X-Custom-Header", "test again")
					req.Header.Set("Content-Type", "application/json")

					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()

					fmt.Println("response Status:", resp.Status)
					fmt.Println("response Headers:", resp.Header)
					body, _ := ioutil.ReadAll(resp.Body)
					fmt.Println("response Body:", string(body))
					return
				}

				conn, err := grpc.Dial("localhost:10002", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				defer conn.Close()
				c := pb.NewEchoServiceClient(conn)

				// Contact the server and print out its response.
				if len(os.Args) > 1 {
					log.Printf("args: %+v", os.Args[1:])
				}
				r, err := c.ValidateToken(context.Background(), &pbos.TokenReqRespData{})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				log.Printf("Greeting: %v", r)
				return
			}

			var opts []grpc.DialOption
			creds := credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000")
			opts = append(opts, grpc.WithTransportCredentials(creds))
			conn, err := grpc.Dial(demoAddr, opts...)
			if err != nil {
				grpclog.Fatalf("fail to dial: %v", err)
			}
			defer conn.Close()
			client := pb.NewEchoServiceClient(conn)

			println(strings.Join(os.Args[1:], " "))

			in := &pbos.TokenReqRespData{}
			// copts := []grpc.CallOption{grpc.EmptyCallOption{}}
			copts := []grpc.CallOption{}
			resp, err := client.ValidateToken(context.Background(), in, copts...)
			if err != nil {
				grpclog.Printf("fail to greet: %v", err)
			}
			println(resp)

		},
	}
	bootcmd.Flags().StringP("name", "", "", "vm name and/or prefix")
	bootcmd.Flags().StringP("flavor_name", "", "m1.small", "vm spec")
	bootcmd.Flags().StringP("image_name", "", "cirros", "vm image")
	bootcmd.Flags().IntP("count", "", 4, "vm count")
	bootcmd.Flags().StringP("secgroup_names", "", "default", "security group name array seperated with comma")
	bootcmd.Flags().StringP("network_name", "", "", "vm network name")
	bootcmd.Flags().StringP("floating_network_name", "", "", "vm floating ip network name")
	bootcmd.Flags().StringP("userdata_file", "", "", "user data file")
	bootcmd.Flags().StringP("personality_files", "", "", "personality file array seperated with comma")

	return bootcmd
}

func netCommand() *cobra.Command {
	netcmd := &cobra.Command{
		Use:   "net",
		Short: "Example networking gRPC service CLI client",
		Run: func(cmd *cobra.Command, args []string) {
			var insecure, ishttp bool
			pflag.BoolVar(&insecure, "insecure", true, "using insecure rpc/http.")
			pflag.BoolVar(&ishttp, "http", true, "using http.")
			pflag.Parse()

			if insecure {
				if ishttp {
					var url = "http://localhost:10001/v1/landscape"
					var jsonStr = []byte(strings.Replace(`
{
  "vnets": [
    {
      "name": "int-stage-0",
      "subnets": [
        {
          "name": "int-192-168-128-0-slash-24",
          "cidr": "192.168.128.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "int-stage-1",
      "subnets": [
        {
          "name": "int-192-168-129-0-slash-24",
          "cidr": "192.168.129.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "int-stage-2",
      "subnets": [
        {
          "name": "int-192-168-130-0-slash-24",
          "cidr": "192.168.130.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "int-stage-3",
      "subnets": [
        {
          "name": "int-192-168-131-0-slash-24",
          "cidr": "192.168.131.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "public",
      "admin_state_up": true,
      "subnets": [
        {
          "name": "10.100.151.0/24",
          "cidr": "10.100.151.0/24",
          "gateway_ip": "10.100.151.1",
          "dns_name_servers": [],
          "allocation_pools": [
            {
              "start": "10.100.151.50",
              "end": "10.100.151.240"
            }
          ],
          "host_routes": [],
          "enable_dhcp": false
        }
      ],
      "shared": true
    }
  ],
  "vrouter": {
    "name": "hack",
    "routes": [],
    "admint_state_up": true
  },
  "secgroup": {
    "name": "hack",
    "security_group_rules": [
      {
        "direction": "ingress",
        "protocol": "tcp"
      },
      {
        "direction": "ingress",
        "protocol": "udp"
      },
      {
        "direction": "ingress",
        "protocol": "icmp"
      }
    ]
  },
  "ifaces_info": [
    {
      "router_name": "hack",
      "network_name": "int-stage-0",
      "subnet_name": "int-192-168-128-0-slash-24",
      "secgroups_info": [
        {
          "name": "hack"
        }
      ]
    },
    {
      "router_name": "hack",
      "network_name": "int-stage-1",
      "subnet_name": "int-192-168-129-0-slash-24",
      "secgroups_info": [
        {
          "name": "hack"
        }
      ]
    },
    {
      "router_name": "hack",
      "network_name": "int-stage-2",
      "subnet_name": "int-192-168-130-0-slash-24",
      "secgroups_info": [
        {
          "name": "hack"
        }
      ]
    },
    {
      "router_name": "hack",
      "network_name": "int-stage-3",
      "subnet_name": "int-192-168-131-0-slash-24",
      "secgroups_info": [
        {
          "name": "hack"
        }
      ]
    }
  ],
  "gateways_info": [
    {
      "network_name": "public",
      "router_name": "hack"
    }
  ]
}
`, "............", rand.String(12), 1))
					obj := new(pbos.OpenstackNeutronLandscapeReqRespData)
					if err := json.Unmarshal(jsonStr, obj); err != nil {
						log.Fatal(err)
					} else {
						log.Println(*obj)
					}
					req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
					req.Header.Set("X-Custom-Header", "test again")
					req.Header.Set("Content-Type", "application/json")

					client := &http.Client{}
					resp, err := client.Do(req)
					if err != nil {
						panic(err)
					}
					defer resp.Body.Close()

					fmt.Println("response Status:", resp.Status)
					fmt.Println("response Headers:", resp.Header)
					body, _ := ioutil.ReadAll(resp.Body)
					fmt.Println("response Body:", string(body))
					return
				}

				conn, err := grpc.Dial("localhost:10002", grpc.WithInsecure())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				defer conn.Close()
				c := pb.NewEchoServiceClient(conn)

				// Contact the server and print out its response.
				if len(os.Args) > 1 {
					log.Printf("args: %+v", os.Args[1:])
				}
				r, err := c.EstablishNetworkLandscape(context.Background(), &pbos.OpenstackNeutronLandscapeReqRespData{})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				log.Printf("Greeting: %v", r)
				return
			}

			var opts []grpc.DialOption
			creds := credentials.NewClientTLSFromCert(demoCertPool, "localhost:10000")
			opts = append(opts, grpc.WithTransportCredentials(creds))
			conn, err := grpc.Dial(demoAddr, opts...)
			if err != nil {
				grpclog.Fatalf("fail to dial: %v", err)
			}
			defer conn.Close()
			client := pb.NewEchoServiceClient(conn)

			println(strings.Join(os.Args[1:], " "))

			in := &pbos.OpenstackNeutronLandscapeReqRespData{}
			// copts := []grpc.CallOption{grpc.EmptyCallOption{}}
			copts := []grpc.CallOption{}
			resp, err := client.EstablishNetworkLandscape(context.Background(), in, copts...)
			if err != nil {
				grpclog.Printf("fail to greet: %v", err)
			}
			println(resp)

		},
	}
	netcmd.Flags().StringP("vnets", "", "", "json expr")
	netcmd.Flags().StringP("vrouter", "", "", "json expr")
	netcmd.Flags().StringP("secgroup", "", "", "json expr")
	netcmd.Flags().StringP("ifaces_info", "", "", "json expr")
	netcmd.Flags().StringP("gateways_info", "", "", "json expr")

	return netcmd
}
