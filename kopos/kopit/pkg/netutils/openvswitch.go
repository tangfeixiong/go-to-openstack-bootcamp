/*
   Alternation
     https://github.com/socketplane/libovsdb
*/

package netutils

import (
	"fmt"

	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util"
	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util/ovs"
)

type OVSctl struct{}

func (ctl OVSctl) List() error {
	util.Logger.Println("Go to show OpenVSwitch DB")

	cli := ovs.Client_vsctl
	out, err := cli.Show()
	if nil != err {
		util.Logger.Println(err)
		return fmt.Errorf("Failed to list ip stacks: %s", err.Error())
	}
	util.Logger.Println(string(out))
	return nil
}
