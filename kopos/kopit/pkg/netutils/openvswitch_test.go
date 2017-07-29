package netutils

import (
	"fmt"
	"testing"

	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util/ovs"
)

func TestOVSvsctl_show(t *testing.T) {
	cli := ovs.Client_vsctl
	out, err := cli.Show()
	if nil != err {
		t.Fail()
	}

	fmt.Println(string(out))
}
