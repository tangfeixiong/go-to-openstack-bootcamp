// tangfeixiong <tangfx128@gmail.com>

package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"

	pb "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Example echo gRPC service CLI client",
	Run: func(cmd *cobra.Command, args []string) {
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

		copts := []grpc.CallOption{grpc.EmptyCallOption{}}
		in := &pb.OpenstackNeutronNetRequestData{Name: "test"}
		resp, err := client.AdminSharedNetworkCreation(context.Background(), in, copts...)
		println(resp)
	},
}

func init() {
	RootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
