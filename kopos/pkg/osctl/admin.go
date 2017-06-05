package osctl

import (
	// "os"
	"strings"

	"github.com/golang/glog"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	_ "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	_ "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	_ "github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	_ "github.com/gophercloud/gophercloud/openstack/utils"
)

type AdminInfra struct {
	InfraProvider
}

func Admin() *AdminInfra {
	return new(AdminInfra)
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

func (admin *AdminInfra) tryto() *AdminInfra {
	if nil == admin.providerclient {
		if opts, err := openstack.AuthOptionsFromEnv(); nil != err {
			glog.Errorf("Could not load admin openrc: %v", err)
			admin.lasterr = err
		} else {
			glog.Infof("Load admin openrc: %v %v", opts.IdentityEndpoint, opts.Username)
			if provider, err := openstack.AuthenticatedClient(opts); nil != err {
				glog.Errorf("Could not authenticate admin: %v", err)
				admin.lasterr = err
			} else {
				glog.Infof("Authenticated for token: %v", provider.TokenID)
				admin.lasterr = nil
				admin.providerclient = provider
			}
		}
	}
	return admin
}

func (admin *AdminInfra) IdentityEndpoint(url string) *AdminInfra {
	admin.identityendpoint = url
	return admin
}

func (admin *AdminInfra) BasicAuthCredential(username, password string) *AdminInfra {
	admin.username, admin.password = username, password
	return admin
}
