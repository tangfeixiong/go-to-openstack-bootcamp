package osctl

import (
	"os"

	"github.com/golang/glog"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	_ "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	_ "github.com/gophercloud/gophercloud/openstack/utils"
)

type InfraProvider struct {
	identityendpoint string
	username         string
	password         string
	providerclient   *gophercloud.ProviderClient
	lasterr          error
}

func (provider *InfraProvider) createNetwork(name, tenantid string, adminstateup, shared bool) (*networks.Network, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
	}
	glog.Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	network, err := networks.Create(client, networks.CreateOpts{
		AdminStateUp: &adminstateup,
		Name:         name,
		Shared:       &shared,
		TenantID:     tenantid,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron network: %v", err)
		return nil, err
	}
	glog.Infof("Succeeded to a new neutron network: %v", network)

	return network, nil
}

func (provider *InfraProvider) createSubnet(networkid, cidr, name, tenantid string, allocationpools []subnets.AllocationPool, gatewayip string, ipversion gophercloud.IPVersion, enabledhcp bool, dnsnameservers []string, hostroutes []subnets.HostRoute) (*subnets.Subnet, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
	}
	glog.Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	subnet, err := subnets.Create(client, subnets.CreateOpts{
		NetworkID:       networkid,
		CIDR:            cidr,
		Name:            name,
		TenantID:        tenantid,
		AllocationPools: allocationpools,
		GatewayIP:       &gatewayip,
		IPVersion:       ipversion,
		EnableDHCP:      &enabledhcp,
		DNSNameservers:  dnsnameservers,
		HostRoutes:      hostroutes,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron subnet: %v", err)
		return nil, err
	}
	glog.Infof("Succeeded to a new neutron subnet: %v", subnet)

	return subnet, nil
}

func (provider *InfraProvider) createRouter(name string, adminstateup, distributed bool, tenantid string, gatewayinfo routers.GatewayInfo) (*routers.Router, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
	}
	glog.Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	router, err := routers.Create(client, routers.CreateOpts{
		Name:         name,
		AdminStateUp: &adminstateup,
		Distributed:  &distributed,
		TenantID:     tenantid,
		GatewayInfo:  &gatewayinfo,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron router: %v", err)
		return nil, err
	}
	glog.Infof("Succeeded to a new neutron router: %v", router)

	return router, nil
}

func (provider *InfraProvider) upstreamRouter(networkid, name string, adminstateup, distributed bool, gatewayinfo routers.GatewayInfo, routes []routers.Route) (*routers.Router, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
	}
	glog.Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	router, err := routers.Update(client, networkid, routers.UpdateOpts{
		Name:         name,
		AdminStateUp: &adminstateup,
		Distributed:  &distributed,
		GatewayInfo:  &gatewayinfo,
		Routes:       routes,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not modify neutron router: %v", err)
		return nil, err
	}
	glog.Infof("Succeeded to modify neutron router: %v", router)

	return router, nil
}

func (provider *InfraProvider) createPort(networkid, name string, adminstateup bool, macaddress string, fixedips []ports.IP, deviceid, deviceowner, tenantid string, securitygroups []string, allowedaddresspairs []ports.AddressPair) (*ports.Port, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
	}
	glog.Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	port, err := ports.Create(client, ports.CreateOpts{
		NetworkID:           networkid,
		Name:                name,
		AdminStateUp:        &adminstateup,
		MACAddress:          macaddress,
		FixedIPs:            fixedips,
		DeviceID:            deviceid,
		DeviceOwner:         deviceowner,
		TenantID:            tenantid,
		SecurityGroups:      securitygroups,
		AllowedAddressPairs: allowedaddresspairs,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron port: %v", err)
		return nil, err
	}
	glog.Infof("Succeeded to a new neutron port: %v", port)

	return port, nil
}

func (provider *InfraProvider) plugRouterIntoSubnet(routerid, portid string) (*routers.InterfaceInfo, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
	}
	glog.Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	ifinfo, err := routers.AddInterface(client, routerid, routers.AddInterfaceOpts{
		PortID: portid,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron interface: %v", err)
		return nil, err
	}
	glog.Infof("Succeeded to a new neutron interface: %v", ifinfo)

	return ifinfo, nil
}

func (provider *InfraProvider) createSecurityGroup(name, tenantid, description string) (*groups.SecGroup, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
	}
	glog.Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	secgroup, err := groups.Create(client, groups.CreateOpts{
		Name:        name,
		TenantID:    tenantid,
		Description: description,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron interface: %v", err)
		return nil, err
	}
	glog.Infof("Succeeded to a new neutron interface: %v", secgroup)

	return secgroup, nil
}

func (provider *InfraProvider) createSecGroupRule(direction rules.RuleDirection, ethertype rules.RuleEtherType, secgroupid string, portrangemax, portrangemin int, protocol rules.RuleProtocol, remotegroupid, remoteipprefix, tenantid string) (*rules.SecGroupRule, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
	}
	glog.Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	secgrouprule, err := rules.Create(client, rules.CreateOpts{
		Direction:      direction,
		EtherType:      ethertype,
		SecGroupID:     secgroupid,
		PortRangeMax:   portrangemax,
		PortRangeMin:   portrangemin,
		Protocol:       protocol,
		RemoteGroupID:  remotegroupid,
		RemoteIPPrefix: remoteipprefix,
		TenantID:       tenantid,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron interface: %v", err)
		return nil, err
	}
	glog.Infof("Succeeded to a new neutron interface: %v", secgrouprule)

	return secgrouprule, nil
}

/*
func tutorial() {
	// Option 1: Pass in the values yourself
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://openstack.example.com:5000/v2.0",
		Username:         "{username}",
		Password:         "{password}",
	}

	// Option 2: Use a utility function to retrieve all your environment variables
	opts, err := openstack.AuthOptionsFromEnv()

	provider, err := openstack.AuthenticatedClient(opts)

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})

	server, err := servers.Create(client, servers.CreateOpts{
		Name:      "My new server!",
		FlavorRef: "flavor_id",
		ImageRef:  "image_id",
	}).Extract()
}
*/
