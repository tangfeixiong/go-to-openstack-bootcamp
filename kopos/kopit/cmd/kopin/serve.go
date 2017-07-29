package kopin

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb"
	pbos "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack"
	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/libvirtclient"
)

func serve(c *cobra.Command, args []string) {
	srv := new(gRPCService)
	s := grpc.NewServer()
	echopb.RegisterWorkerServiceServer(s, srv)

	l, err := net.Listen("tcp", ":10101")
	if err != nil {
		panic(err)
	}

	fmt.Println("starting gRPC on host", l.Addr())
	if err := s.Serve(l); nil != err {
		panic(err)
	}
}

type gRPCService struct {
	temp string
}

func (gs *gRPCService) StartPacketCapturing(ctx context.Context, in *pbos.OSILayersReqRespData) (*pbos.OSILayersReqRespData, error) {
	return new(pbos.OSILayersReqRespData), nil
}

func (gs *gRPCService) StopPacketCapturing(ctx context.Context, in *pbos.OSILayersReqRespData) (*pbos.OSILayersReqRespData, error) {
	return new(pbos.OSILayersReqRespData), nil
}

func (gs *gRPCService) GetLibvirtDomainInfo(ctx context.Context, in *pbos.LibvirtDomainInfo) (*pbos.LibvirtDomainInfo, error) {
	return libvirtclient.Execute()
}

func (gs *gRPCService) Echo(ctx context.Context, req *echopb.EchoMessage) (*echopb.EchoMessage, error) {
	resp := new(echopb.EchoMessage)
	if nil != req {
		if 0 != len(req.Value) {
			resp = req
		} else {
			resp.Value = "So what!"
		}
	}
	return resp, fmt.Errorf("Bad request")
}
