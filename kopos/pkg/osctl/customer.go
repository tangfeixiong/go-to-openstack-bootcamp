package osctl

import (
	// "os"
	"strings"

	"github.com/golang/glog"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	_ "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	_ "github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	_ "github.com/gophercloud/gophercloud/openstack/utils"
)

type InfraResPlayer struct {
	InfraProvider
}

func InfraResClient() *InfraResPlayer {
	return new(InfraResPlayer)
}

func (player *InfraResPlayer) CreateConsoleIntoDnatWithNetworkAndMachine(name, tenandid, subnetcidr, gatewayip, description, imagename, flavorname string) {

	adminstateup := false
	shared := false
	network, err := player.tryto().createNetwork(name, tenandid, adminstateup, shared)
	if nil != err {
		return
	}

	offset := strings.LastIndex(subnetcidr, ".")
	ipprefix := subnetcidr[:offset]
	ipversion := gophercloud.IPv4
	enabledhcp := true
	dnsnameservers := []string(nil)        // []string{}
	hostroutes := []subnets.HostRoute(nil) //[]subnets.HostRoute{}
	subnet, err := player.tryto().createSubnet(network.ID, subnetcidr, name, tenandid, []subnets.AllocationPool{{ipprefix + ".50", ipprefix + ".200"}}, gatewayip, ipversion, enabledhcp, dnsnameservers, hostroutes)
	if nil != err {
		return
	}
	println(subnet)

	distributed := false
	gatewayinfo := routers.GatewayInfo{NetworkID: network.ID}
	router, err := player.tryto().createRouter(name, adminstateup, distributed, tenandid, gatewayinfo)
	if nil != err {
		return
	}

	secgroup, err := player.tryto().createSecurityGroup(name, tenandid, description)
	if nil != err {
		return
	}

	securitygroups := []string{}
	// remoteipprefix = "::/0"
	portrangemax := 65535
	portrangemin := 0
	protocol := rules.ProtocolTCP
	remotegroupid := ""
	remoteipprefix := "0.0.0.0/0"
	secgrouprule, err := player.tryto().createSecGroupRule(rules.DirIngress, rules.EtherType4, secgroup.ID, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenandid)
	if nil != err {
		return
	}
	securitygroups = append(securitygroups, secgrouprule.ID)

	protocol = rules.ProtocolUDP
	secgrouprule, err = player.tryto().createSecGroupRule(rules.DirIngress, rules.EtherType4, secgroup.ID, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenandid)
	if nil != err {
		return
	}
	securitygroups = append(securitygroups, secgrouprule.ID)

	protocol = rules.ProtocolICMP
	portrangemax = -1
	portrangemin = -1
	secgrouprule, err = player.tryto().createSecGroupRule(rules.DirIngress, rules.EtherType4, secgroup.ID, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenandid)
	if nil != err {
		return
	}
	securitygroups = append(securitygroups, secgrouprule.ID)

	macaddress := ""
	fixedips := []ports.IP{}
	deviceid := ""
	deviceowner := ""
	allowedaddresspairs := []ports.AddressPair{}
	port, err := player.tryto().createPort(network.ID, name, adminstateup, macaddress, fixedips, deviceid, deviceowner, tenandid, securitygroups, allowedaddresspairs)
	if nil != err {
		return
	}

	ifinfo, err := player.tryto().plugRouterIntoSubnet(router.ID, port.ID)
	if nil != err {
		return
	}
	println(ifinfo)

	consoleNetworks := []servers.Network{}
	vm, err := player.tryto().createMachine(name, imagename, flavorname, consoleNetworks)
	if nil != err {
		return
	}
	println(vm)
}

func (player *InfraResPlayer) CreateTargetDrone(name, tenandid, subnetcidr, gatewayip, description, imagename, flavorname string) {

	consoleNetworks := []servers.Network{}
	vm, err := player.tryto().createMachine(name, imagename, flavorname, consoleNetworks)
	if nil != err {
		return
	}
	println(vm)

}

func (player *InfraResPlayer) tryto() *InfraResPlayer {
	if nil == player.providerclient {
		if opts, err := openstack.AuthOptionsFromEnv(); nil != err {
			glog.Errorf("Could not load player openrc: %v", err)
			player.lasterr = err
		} else {
			glog.Infof("Load player openrc: %v %v", opts.IdentityEndpoint, opts.Username)
			if provider, err := openstack.AuthenticatedClient(opts); nil != err {
				glog.Errorf("Could not authenticate player: %v", err)
				player.lasterr = err
			} else {
				glog.Infof("Authenticated for token: %v", provider.TokenID)
				player.lasterr = nil
				player.providerclient = provider
			}
		}
	}
	return player
}

func (player *InfraResPlayer) IdentityEndpoint(url string) *InfraResPlayer {
	player.identityendpoint = url
	return player
}

func (player *InfraResPlayer) BasicAuthCredential(username, password string) *InfraResPlayer {
	player.username, player.password = username, password
	return player
}
