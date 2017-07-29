package netutils

import (
	"fmt"
	"testing"

	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util"
)

func TestIPAddrList(t *testing.T) {
	util.Logger.Println("Go to show all IP address")

	cli := util.Client
	out, err := cli.AddrShow("")
	if nil != err {
		util.Logger.Println(err)
		t.Fail()
	}

	fmt.Println(string(out))
}

func Test_IPAddrList(t *testing.T) {
	util.Logger.Println("Go to show all IP address")

	result, err := List()

	if err != nil {
		t.Fail()
	}

	for _, item := range result {
		fmt.Println(item)
	}

	t.Log(result)
}
