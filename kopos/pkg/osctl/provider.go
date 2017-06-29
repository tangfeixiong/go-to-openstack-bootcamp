package osctl

import (
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/flavors"
	computeimages "github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/identity/v2/tenants"
	"github.com/gophercloud/gophercloud/openstack/identity/v2/tokens"
	"github.com/gophercloud/gophercloud/openstack/identity/v2/users"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	"github.com/gophercloud/gophercloud/openstack/utils"
)

type InfraProvider struct {
	identityendpoint string
	username         string
	password         string
	providerclient   *gophercloud.ProviderClient
	lasterr          error
}

func (provider *InfraProvider) gainNetworks() ([]networks.Network, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap network service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap network serivce: %v%v", client.Endpoint, client.ResourceBase)

	status_fake := "ACTIVE"
	name_fake := "fake network name"
	adminStateUp := false
	tenantId := "fake project id"
	shared_fake := false
	id_fake := "fake network id"
	marker_fake := "00000000-0000-0000-0000-000000000000"
	limit_fake := 0
	sortKey := "ID"
	sortDir := "ASC"

	opts := networks.ListOpts{
		Status:       status_fake,
		Name:         name_fake,
		AdminStateUp: &adminStateUp,
		TenantID:     tenantId,
		Shared:       &shared_fake,
		ID:           id_fake,
		Marker:       marker_fake,
		Limit:        limit_fake,
		SortKey:      sortKey,
		SortDir:      sortDir,
	}
	glog.V(9).Infoln(opts)

	resp, err := networks.List(client, nil).AllPages()
	if nil != err {
		glog.Errorf("Could not reap network-list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap network-list capability: %v", resp)

	result, err := networks.ExtractNetworks(resp)
	if nil != err {
		glog.Errorf("Could not unmarshall networks result: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to unmarshall networks result: %v", result)

	return result, nil
}

func (provider *InfraProvider) searchNetworks(name string) ([]networks.Network, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap network service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap network serivce: %v%v", client.Endpoint, client.ResourceBase)

	status_fake := "ACTIVE"
	name_fake := "fake network name"
	adminStateUp := false
	tenantId := "fake project id"
	shared_fake := false
	id_fake := "fake network id"
	marker_fake := "00000000-0000-0000-0000-000000000000"
	limit_fake := 0
	sortKey := "ID"
	sortDir := "ASC"

	opts := networks.ListOpts{
		Status:       status_fake,
		Name:         name_fake,
		AdminStateUp: &adminStateUp,
		TenantID:     tenantId,
		Shared:       &shared_fake,
		ID:           id_fake,
		Marker:       marker_fake,
		Limit:        limit_fake,
		SortKey:      sortKey,
		SortDir:      sortDir,
	}
	opts = networks.ListOpts{
		Name: name,
	}

	resp, err := networks.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap network-list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap network-list capability: %v", resp)

	result, err := networks.ExtractNetworks(resp)
	if nil != err {
		glog.Errorf("Could not unmarshall networks result: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to unmarshall networks result: %v", result)

	return result, nil
}

func (provider *InfraProvider) createNetwork(name, tenantid string, adminstateup, shared bool) (*networks.Network, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	adminStateUp := true
	name_fake := "virtual network"
	shared_fake := true
	tenantId := "00000000-0000-0000-0000-000000000000"

	opts := networks.CreateOpts{
		AdminStateUp: &adminStateUp,
		Name:         name_fake,
		Shared:       &shared_fake,
		TenantID:     tenantId,
	}
	opts = networks.CreateOpts{
		AdminStateUp: &adminstateup,
		Name:         name,
		Shared:       &shared,
		TenantID:     tenantid,
	}

	resp, err := networks.Create(client, opts).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron network: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a new neutron network: %v", resp)

	return resp, nil
}

func (provider *InfraProvider) queryNetwork(id string) (*networks.Network, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap network service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap network serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	resp, err := networks.Get(client, id).Extract()
	if nil != err {
		glog.Errorf("Could not reap network-show interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap network-show interface: %v", resp)

	return resp, nil
}

func (provider *InfraProvider) identifyNetwork(name string) (*string, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap network service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap network serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	resp, err := networks.IDFromName(client, name)
	if nil != err {
		glog.Errorf("Could not reap network-list interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap network-list interface: %v", resp)

	return &resp, nil
}

func (provider *InfraProvider) gainSubnets() ([]subnets.Subnet, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap network service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap network serivce: %v%v", client.Endpoint, client.ResourceBase)

	name_fake := "fake subnet name"
	enableDHCP := true
	networkId := "fake network id"
	tenantId := "fake project id"
	ipVersion := int(gophercloud.IPv4)
	gatewayIp := "0.0.0.1"
	cidr_fake := "0.0.0.0/0"
	id_fake := "fake subnet id"
	limit_fake := 0
	marker_fake := "00000000-0000-0000-0000-000000000000"
	sortKey := "ID"
	sortDir := "ASC"

	opts := subnets.ListOpts{
		Name:       name_fake,
		EnableDHCP: &enableDHCP,
		NetworkID:  networkId,
		TenantID:   tenantId,
		IPVersion:  ipVersion,
		GatewayIP:  gatewayIp,
		CIDR:       cidr_fake,
		ID:         id_fake,
		Limit:      limit_fake,
		Marker:     marker_fake,
		SortKey:    sortKey,
		SortDir:    sortDir,
	}
	glog.V(9).Infoln(opts)

	resp, err := subnets.List(client, nil).AllPages()
	if nil != err {
		glog.Errorf("Could not reap subnet-list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap subnet-list capability: %v", resp)

	result, err := subnets.ExtractSubnets(resp)
	if nil != err {
		glog.Errorf("Could not unmarshall subnets result: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to unmarshall subnets result: %v", result)

	return result, nil
}

func (provider *InfraProvider) searchSubnet(name string) ([]subnets.Subnet, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap network service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap network serivce: %v%v", client.Endpoint, client.ResourceBase)

	name_fake := "fake subnet name"
	enableDHCP := true
	networkId := "fake network id"
	tenantId := "fake project id"
	ipVersion := int(gophercloud.IPv4)
	gatewayIp := "0.0.0.1"
	cidr_fake := "0.0.0.0/0"
	id_fake := "fake subnet id"
	limit_fake := 0
	marker_fake := "00000000-0000-0000-0000-000000000000"
	sortKey := "ID"
	sortDir := "ASC"

	opts := subnets.ListOpts{
		Name:       name_fake,
		EnableDHCP: &enableDHCP,
		NetworkID:  networkId,
		TenantID:   tenantId,
		IPVersion:  ipVersion,
		GatewayIP:  gatewayIp,
		CIDR:       cidr_fake,
		ID:         id_fake,
		Limit:      limit_fake,
		Marker:     marker_fake,
		SortKey:    sortKey,
		SortDir:    sortDir,
	}
	opts = subnets.ListOpts{
		Name: name,
	}

	resp, err := subnets.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap subnet-list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap subnet-list capability: %v", resp)

	result, err := subnets.ExtractSubnets(resp)
	if nil != err {
		glog.Errorf("Could not unmarshall subnets result: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to unmarshall subnets result: %v", result)

	return result, nil
}

func (provider *InfraProvider) searchSubnets(networkid string) ([]subnets.Subnet, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap networking service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap networking serivce: %v%v", client.Endpoint, client.ResourceBase)

	name_fake := "fake subnet name"
	enableDHCP := true
	networkId := "fake network id"
	tenantId := "fake project id"
	ipVersion := int(gophercloud.IPv4)
	gatewayIp := "0.0.0.1"
	cidr_fake := "0.0.0.0/0"
	id_fake := "fake subnet id"
	limit_fake := 0
	marker_fake := "00000000-0000-0000-0000-000000000000"
	sortKey := "ID"
	sortDir := "ASC"

	opts := subnets.ListOpts{
		Name:       name_fake,
		EnableDHCP: &enableDHCP,
		NetworkID:  networkId,
		TenantID:   tenantId,
		IPVersion:  ipVersion,
		GatewayIP:  gatewayIp,
		CIDR:       cidr_fake,
		ID:         id_fake,
		Limit:      limit_fake,
		Marker:     marker_fake,
		SortKey:    sortKey,
		SortDir:    sortDir,
	}
	opts = subnets.ListOpts{
		NetworkID: networkid,
	}

	resp, err := subnets.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap subnet list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap subnet list capability: %v", resp)

	result, err := subnets.ExtractSubnets(resp)
	if nil != err {
		glog.Errorf("Could not reap subnets unmarshall capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap subnets unmarshall capability: %v", result)

	return result, nil
}

func (provider *InfraProvider) createSubnet(networkid, cidr, name, tenantid string, allocationpools []subnets.AllocationPool, gatewayip string, ipversion gophercloud.IPVersion, enabledhcp bool, dnsnameservers []string, hostroutes []subnets.HostRoute) (*subnets.Subnet, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	networkId := "00000000-0000-0000-0000-000000000000"
	cidrFake := "0.0.0.0/0"
	nameFake := "ip net"
	tenantId := "00000000-0000-0000-0000-000000000000"
	allocationPools := []subnets.AllocationPool{{"0.0.0.2", "0.0.0.254"}}
	gatewayIp := "0.0.0.1"
	ipVersion := gophercloud.IPv4
	enableDhcp := true
	dnsNameServers := []string{"8.8.4.4", "8.8.8.8"}
	hostRoutes := []subnets.HostRoute{{DestinationCIDR: "10.1.100.0/24", NextHop: "172.17.14.10"}}

	opts := subnets.CreateOpts{
		NetworkID:       networkId,
		CIDR:            cidrFake,
		Name:            nameFake,
		TenantID:        tenantId,
		AllocationPools: allocationPools,
		GatewayIP:       &gatewayIp,
		IPVersion:       ipVersion,
		EnableDHCP:      &enableDhcp,
		DNSNameservers:  dnsNameServers,
		HostRoutes:      hostRoutes,
	}
	opts = subnets.CreateOpts{
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
	}

	subnet, err := subnets.Create(client, opts).Extract()
	if nil != err {
		glog.Errorf("Could not reap subnet-create interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to resp subnet-create interface: %v", subnet)

	return subnet, nil
}

func (provider *InfraProvider) identifySubnet(name string) (*string, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap networking service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap networking serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	resp, err := subnets.IDFromName(client, name)
	if nil != err {
		glog.Errorf("Could not reap subnet-list interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap subnet-list interface: %v", resp)

	return &resp, nil
}

func (provider *InfraProvider) searchRouters(name string) ([]routers.Router, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap network service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap network serivce: %v%v", client.Endpoint, client.ResourceBase)

	id_fake := "fake subnet id"
	name_fake := "fake subnet name"
	adminStateUp := false
	distributed_fake := false
	status_fake := "ACTIVE"
	tenantId := "fake project id"
	limit_fake := 0
	marker_fake := "00000000-0000-0000-0000-000000000000"
	sortKey := "ID"
	sortDir := "ASC"

	opts := routers.ListOpts{
		ID:           id_fake,
		Name:         name_fake,
		AdminStateUp: &adminStateUp,
		Distributed:  &distributed_fake,
		Status:       status_fake,
		TenantID:     tenantId,
		Limit:        limit_fake,
		Marker:       marker_fake,
		SortKey:      sortKey,
		SortDir:      sortDir,
	}
	opts = routers.ListOpts{
		Name: name,
	}

	resp, err := routers.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap router-list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap router-list capability: %v", resp)

	result, err := routers.ExtractRouters(resp)
	if nil != err {
		glog.Errorf("Could not unmarshall routers result: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to unmarshall routers result: %v", result)

	return result, nil
}

func (provider *InfraProvider) createRouter(name string, adminstateup, distributed bool, tenantid string, gatewayinfo routers.GatewayInfo) (*routers.Router, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	name_fake := "virtual router"
	adminStateUp := false
	distributed_fake := false
	tenantId := "00000000-0000-0000-0000-000000000000"
	gatewayInfo := routers.GatewayInfo{NetworkID: "00000000-0000-0000-0000-000000000000"}

	opts := routers.CreateOpts{
		Name:         name_fake,
		AdminStateUp: &adminStateUp,
		Distributed:  &distributed_fake,
		TenantID:     tenantId,
		GatewayInfo:  &gatewayInfo,
	}
	opts = routers.CreateOpts{
		Name:         name,
		AdminStateUp: &adminstateup,
		Distributed:  &distributed,
		TenantID:     tenantid,
	}
	if 0 != len(gatewayinfo.NetworkID) {
		opts.GatewayInfo = &gatewayinfo
	}

	router, err := routers.Create(client, opts).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron router: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a new neutron router: %v", router)

	return router, nil
}

func (provider *InfraProvider) upstreamRouter(id, name string, adminstateup, distributed bool, gatewayinfo routers.GatewayInfo, routes []routers.Route) (*routers.Router, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	name_fake := "required"
	adminStateUp := false
	distributed_fake := false
	gatewayInfo := routers.GatewayInfo{NetworkID: "required"}
	routes_fake := []routers.Route{{NextHop: "ip addr", DestinationCIDR: "ip range"}}

	opts := routers.UpdateOpts{
		Name:         name_fake,
		AdminStateUp: &adminStateUp,
		Distributed:  &distributed_fake,
		GatewayInfo:  &gatewayInfo,
		Routes:       routes_fake,
	}
	opts = routers.UpdateOpts{
		// Name: name,
		// AdminStateUp: &adminstateup,
		// Distributed:  &distributed,
		GatewayInfo: &gatewayinfo,
		// Routes:       routes,
	}

	resp, err := routers.Update(client, id, opts).Extract()
	if nil != err {
		glog.Errorf("Could not modify neutron router: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to modify neutron router: %v", resp)

	return resp, nil
}

func (provider *InfraProvider) gainFloatingIps() ([]floatingips.FloatingIP, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	id_fake := "The id"
	floatingNetworkId := "network id"
	floatingIp := "ip"
	portId := "port id"
	fixedIp := "ip"
	tenantId := "fake"
	limit_fake := 0
	marker_fake := "FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF"
	sortKey := "ID"
	sortDir := "ASC"
	routerId := "fake"

	opts := floatingips.ListOpts{
		ID:                id_fake,
		FloatingNetworkID: floatingNetworkId,
		FloatingIP:        floatingIp,
		PortID:            portId,
		FixedIP:           fixedIp,
		TenantID:          tenantId,
		Limit:             limit_fake,
		Marker:            marker_fake,
		SortKey:           sortKey,
		SortDir:           sortDir,
		RouterID:          routerId,
	}
	opts = floatingips.ListOpts{
		SortDir: sortDir,
	}

	result, err := floatingips.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap networking floatingips capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap networking floating capability: %v", result)

	resp, err := floatingips.ExtractFloatingIPs(result)
	if nil != err {
		glog.Errorf("Could not reap networking floatingips marshall: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap networking floating marshall: %v", resp)

	return resp, nil
}

func (provider *InfraProvider) createFloatingIp(floatingnetworkid, floatingip, portid, fixedip, tenantid string) (*floatingips.FloatingIP, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	floatingNetworkId := "required"
	floatingIp := "optional"
	portId := "required"
	fixedIp := "optional"
	tenantId := "optional"

	opts := floatingips.CreateOpts{
		FloatingNetworkID: floatingNetworkId,
		FloatingIP:        floatingIp,
		PortID:            portId,
		FixedIP:           fixedIp,
		TenantID:          tenantId,
	}
	opts = floatingips.CreateOpts{
		FloatingNetworkID: floatingnetworkid,
		PortID:            portid,
	}
	if 0 != len(floatingip) {
		opts.FloatingIP = floatingip
	}
	if 0 != len(fixedip) {
		opts.FixedIP = fixedip
	}
	if 0 != len(tenantid) {
		opts.TenantID = tenantid
	}

	result, err := floatingips.Create(client, opts).Extract()
	if nil != err {
		glog.Errorf("Could not reap networking floatingips capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap networking floating capability: %v", result)

	return result, nil
}

func (provider *InfraProvider) queryFloatingIp(id string) (*floatingips.FloatingIP, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := floatingips.Get(client, id).Extract()
	if nil != err {
		glog.Errorf("Could not reap networking floatingips capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap networking floating capability: %v", result)

	return result, nil
}

func (provider *InfraProvider) createPort(networkid, name string, adminstateup bool, macaddress string, fixedips []ports.IP, deviceid, deviceowner, tenantid string, securitygroups []string, allowedaddresspairs []ports.AddressPair) (*ports.Port, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	networkId := "required"
	nameFake := "required"
	adminStateUp := false
	macAddress := "00-00-00-00-00"
	fixedIps := []ports.IP{{SubnetID: "required", IPAddress: "0.0.0.0"}}
	deviceId := ""
	deviceOwner := ""
	tenantId := "required"
	securityGroups := []string{"default", "user-defined"}
	allowedAddressPairs := []ports.AddressPair{{IPAddress: "0.0.0.0", MACAddress: "00-00-00-00-00-00"}}

	opts := ports.CreateOpts{
		NetworkID:           networkId,
		Name:                nameFake,
		AdminStateUp:        &adminStateUp,
		MACAddress:          macAddress,
		FixedIPs:            fixedIps,
		DeviceID:            deviceId,
		DeviceOwner:         deviceOwner,
		TenantID:            tenantId,
		SecurityGroups:      securityGroups,
		AllowedAddressPairs: allowedAddressPairs,
	}
	opts = ports.CreateOpts{
		NetworkID: networkid,
	}
	if 0 != len(name) {
		opts.Name = name
	}
	if !adminstateup {
		opts.AdminStateUp = &adminstateup
	}
	if 0 != len(macaddress) {
		opts.MACAddress = macaddress
	}
	if 0 != len(fixedips) {
		opts.FixedIPs = fixedips
	}
	if 0 != len(deviceid) {
		opts.DeviceID = deviceid
	}
	if 0 != len(deviceowner) {
		opts.DeviceOwner = deviceowner
	}
	if 0 != len(tenantid) {
		opts.TenantID = tenantid
	}
	if 0 != len(securitygroups) {
		opts.SecurityGroups = securitygroups
	}
	if 0 != len(allowedaddresspairs) {
		opts.AllowedAddressPairs = allowedaddresspairs
	}

	port, err := ports.Create(client, opts).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron port: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a new neutron port: %v", port)

	return port, nil
}

func (provider *InfraProvider) queryPort(id string) (*ports.Port, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := ports.Get(client, id).Extract()
	if nil != err {
		glog.Errorf("Could not reap networking port creation capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap networking port creation capability: %v", result)

	return result, nil
}

func (provider *InfraProvider) plugRouterIntoSubnet(routerid, portid string) (*routers.InterfaceInfo, error) {
	glog.Infof("Invoke add router interface: router=%v, port=%v", routerid, portid)
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	ifinfo, err := routers.AddInterface(client, routerid, routers.AddInterfaceOpts{
		SubnetID: "",
		PortID:   portid,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not reap router interface creation: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap router interface creation: %v", ifinfo)

	return ifinfo, nil
}

func (provider *InfraProvider) createSecurityGroup(name, tenantid, description string) (*groups.SecGroup, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	secgroup, err := groups.Create(client, groups.CreateOpts{
		Name:        name,
		TenantID:    tenantid,
		Description: description,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a new neutron interface: %v", secgroup)

	return secgroup, nil
}

func (provider *InfraProvider) querySecurityGroup(id string) (*groups.SecGroup, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := groups.Get(client, id).Extract()
	if nil != err {
		glog.Errorf("Could not reap networking query capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a resp networking query capability: %v", result)

	return result, nil
}

func (provider *InfraProvider) identifySecurityGroup(name string) (*string, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := groups.IDFromName(client, name)
	if nil != err {
		glog.Errorf("Could not reap networking query capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a resp networking query capability: %v", result)

	return &result, nil
}

func (provider *InfraProvider) createSecGroupRule(direction rules.RuleDirection, ethertype rules.RuleEtherType, secgroupid string, portrangemax, portrangemin int, protocol rules.RuleProtocol, remotegroupid, remoteipprefix, tenantid string) (*rules.SecGroupRule, error) {
	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

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
	glog.V(5).Infof("Succeeded to a new neutron interface: %v", secgrouprule)

	return secgrouprule, nil
}

func (provider *InfraProvider) gainImages() ([]images.Image, error) {
	client, err := openstack.NewImageServiceV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap glance service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap glance serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	limit_fake := 0
	marker_fake := "00000000-0000-0000-0000-000000000000"
	name_fake := "search condition"
	visibility_fake := images.ImageVisibilityPublic
	memberStatus := images.ImageMemberStatusAll
	owner_fake := "uploader"
	status_fake := images.ImageStatusActive
	sizeMin := int64(100)
	sizeMax := int64(10000)
	sortKey := "id"
	sortDir := "ASC"
	tag_fake := ""

	opts := images.ListOpts{
		Limit:        limit_fake,
		Marker:       marker_fake,
		Name:         name_fake,
		Visibility:   visibility_fake,
		MemberStatus: memberStatus,
		Owner:        owner_fake,
		Status:       status_fake,
		SizeMin:      sizeMin,
		SizeMax:      sizeMax,
		SortKey:      sortKey,
		SortDir:      sortDir,
		Tag:          tag_fake,
	}
	glog.V(9).Info(opts)
	// opts = images.ListOpts{
	// 	Status: status_fake,
	// }

	result, err := images.List(client, nil).AllPages()
	if nil != err {
		glog.Errorf("Could not reap image list interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap image list interface: %v", result)

	images, err := images.ExtractImages(result)
	if nil != err {
		glog.Errorf("Could not reap image list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap image list capability: %v", images)

	return images, nil
}

func (provider *InfraProvider) searchImages(name string) ([]images.Image, error) {
	client, err := openstack.NewImageServiceV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap glance service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap glance serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	limit_fake := 0
	marker_fake := "00000000-0000-0000-0000-000000000000"
	name_fake := "search condition"
	visibility_fake := images.ImageVisibilityPublic
	memberStatus := images.ImageMemberStatusAll
	owner_fake := "uploader"
	status_fake := images.ImageStatusActive
	sizeMin := int64(100)
	sizeMax := int64(10000)
	sortKey := "id"
	sortDir := "ASC"
	tag_fake := ""

	opts := images.ListOpts{
		Limit:        limit_fake,
		Marker:       marker_fake,
		Name:         name_fake,
		Visibility:   visibility_fake,
		MemberStatus: memberStatus,
		Owner:        owner_fake,
		Status:       status_fake,
		SizeMin:      sizeMin,
		SizeMax:      sizeMax,
		SortKey:      sortKey,
		SortDir:      sortDir,
		Tag:          tag_fake,
	}
	opts = images.ListOpts{
		Name: name,
	}

	resp, err := images.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap glance search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap glance search interface: %v", resp)

	result, err := images.ExtractImages(resp)
	if nil != err {
		glog.Errorf("Could not complete glance search capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to glance search capability: %v", result)

	return result, nil
}

func (provider *InfraProvider) queryImage(id string) (*images.Image, error) {
	client, err := openstack.NewImageServiceV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap glance service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap glance serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := images.Get(client, id).Extract()
	if nil != err {
		glog.Errorf("Could not complete glance get capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to glance get capability: %v", result)

	return result, nil
}

func (provider *InfraProvider) identifyImage(name string) (*string, error) {
	client, err := openstack.NewImageServiceV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap image service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap image serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := images.List(client, images.ListOpts{
		Name: name,
	}).AllPages()
	if nil != err {
		glog.Errorf("Could not reap image search capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap image search capability: %v", result)

	resp, err := images.ExtractImages(result)
	if nil != err {
		glog.Errorf("Could not reap image search capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap image search capability: %v", resp)
	if 0 == len(resp) {
		return nil, fmt.Errorf("No existed")
	}
	id := resp[0].ID
	return &id, nil
}

func (provider *InfraProvider) gainComputeImages() ([]computeimages.Image, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap glance service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap glance serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	// kind := "SERVER" // "BASE" "SERVER" "ALL"
	result, err := computeimages.ListDetail(client, computeimages.ListOpts{
		// ChangesSince: changesince,
		// Limit:      limit,
		// Marker:      marker,
		// Name: name,
		// Server:       server,
		Status: string(images.ImageStatusActive),
		// Type:   kind,
	}).AllPages()
	if nil != err {
		glog.Errorf("Could not reap compute/images search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/images search interface: %v", result)

	images, err := computeimages.ExtractImages(result)
	if nil != err {
		glog.Errorf("Could not gain compute/images result: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to gain compute/images result: %v", result)

	return images, nil
}

func (provider *InfraProvider) searchComputeImages(name string) ([]computeimages.Image, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap glance service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap glance serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	kind := "SERVER" // "BASE" "SERVER" "ALL"
	resp, err := computeimages.ListDetail(client, computeimages.ListOpts{
		// ChangesSince: changesince,
		// Limit:      limit,
		// Marker:      marker,
		Name: name,
		// Server:       server,
		Status: string(images.ImageStatusActive),
		Type:   kind,
	}).AllPages()
	if nil != err {
		glog.Errorf("Could not reap glance search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap glance search interface: %v", resp)

	result, err := computeimages.ExtractImages(resp)
	if nil != err {
		glog.Errorf("Could not reap glance search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap glance reap interface: %v", result)

	return result, nil
}

func (provider *InfraProvider) queryComputeImage(id string) (*computeimages.Image, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap glance service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap glance serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := computeimages.Get(client, id).Extract()
	if nil != err {
		glog.Errorf("Could not reap glance reap interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap glance reap interface: %v", result)

	return result, nil
}

func (provider *InfraProvider) gainFlavors() ([]flavors.Flavor, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	changesince := "1990-12-30T01-59-59AM"
	mindisk := 10000
	minram := 2000
	marker := "00000000-0000-0000-0000-000000000000"
	limit := 0
	accesstype := flavors.PublicAccess
	opts := flavors.ListOpts{
		ChangesSince: changesince,
		MinDisk:      mindisk,
		MinRAM:       minram,
		Marker:       marker,
		Limit:        limit,
		AccessType:   accesstype,
	}
	glog.V(9).Infoln(opts)

	result, err := flavors.ListDetail(client, nil).AllPages()
	if nil != err {
		glog.Errorf("Could not reap compute/flavors list interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/flavors list interface: %v", result)

	flavors, err := flavors.ExtractFlavors(result)
	if nil != err {
		glog.Errorf("Could not reap compute/flavors list interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/flavors list interface: %v", flavors)

	return flavors, nil
}

func (provider *InfraProvider) queryFlavor(id string) (*flavors.Flavor, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := flavors.Get(client, id).Extract()
	if nil != err {
		glog.Errorf("Could not reap compute/flavors search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/flavors search interface: %v", result)

	return result, nil
}

func (provider *InfraProvider) identifyFlavor(name string) (*string, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := flavors.IDFromName(client, name)
	if nil != err {
		glog.Errorf("Could not reap compute/flavors search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/falvors search interface: %v", result)

	return &result, nil
}

func (provider *InfraProvider) gainMachines() ([]servers.Server, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	changesSince := "1970-12-31T11:59:59PM"
	image := "http://openstack.glance.api.endpoint/images/id"
	flavor := "http://openstack.nova.api.endpoint/flavors/id"
	nameFake := "cirros"
	status := "ACTIVE"
	host := "all of hypervisors"
	marker := "00000000-0000-0000-0000-000000000000"
	limit := 0
	alltenants := false

	opts := servers.ListOpts{
		ChangesSince: changesSince,
		Image:        image,
		Flavor:       flavor,
		Name:         nameFake,
		Status:       status,
		Host:         host,
		Marker:       marker,
		Limit:        limit,
		AllTenants:   alltenants,
	}
	glog.V(9).Infoln(opts)

	resp, err := servers.List(client, nil).AllPages()
	if nil != err {
		glog.Errorf("Could not reap compute/servers list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/servers list capability: %v", resp)

	result, err := servers.ExtractServers(resp)
	if nil != err {
		glog.Errorf("Could not decode compute/servers result: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to decode compute/servers result: %v", result)

	return result, nil
}

func (provider *InfraProvider) searchMachines(name string) ([]servers.Server, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	changesSince := "1970-12-31T11:59:59PM"
	image := "http://openstack.glance.api.endpoint/images/id"
	flavor := "http://openstack.nova.api.endpoint/flavors/id"
	nameFake := "cirros"
	status := "ACTIVE"
	host := "all of hypervisors"
	marker := "00000000-0000-0000-0000-000000000000"
	limit := 0
	alltenants := false

	opts := servers.ListOpts{
		ChangesSince: changesSince,
		Image:        image,
		Flavor:       flavor,
		Name:         nameFake,
		Status:       status,
		Host:         host,
		Marker:       marker,
		Limit:        limit,
		AllTenants:   alltenants,
	}
	opts = servers.ListOpts{
		Name:       name,
		AllTenants: alltenants,
	}

	resp, err := servers.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap compute/servers list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/servers list capability: %v", resp)

	result, err := servers.ExtractServers(resp)
	if nil != err {
		glog.Errorf("Could not decode compute/servers result: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to decode compute/serveres result: %v", result)

	return result, nil
}

func (provider *InfraProvider) createMachine(name, imageid, imagename, flavorid, flavorname string, securitygroups []string, userdata []byte, networks []servers.Network, personality servers.Personality, adminpass, accessipv4 string) (*servers.Server, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	name_fake := "new vm name"
	imageRef := "00000000-0000-0000-0000-000000000000"
	imageName := "cirros"
	flavorRef := "00000000-0000-0000-0000-000000000000"
	flavorName := "m1.small"
	securityGroups := []string{"default", "customization"}
	userData := []byte("/bin/cat /etc/OS-RELEASE")
	availabilityZone := "no zone"
	networks_fake := []servers.Network{{
		UUID:    "", // Required unless Port is provided
		Port:    "00000000-0000-0000-0000-000000000000",
		FixedIP: "optional, maybe using DHCP",
	}}
	metadata_fake := map[string]string{"key": "value"}
	personality_fake := []*servers.File{&servers.File{Path: "path-to-injected-file", Contents: []byte("contents-of-injected-file")}}
	configDrive := false
	adminPass := "secret"
	accessIpv4 := "0.0.0.0"
	accessIpv6 := "::"
	serviceClient := client

	opts := servers.CreateOpts{
		Name:             name_fake,
		ImageRef:         imageRef,
		ImageName:        imageName,
		FlavorRef:        flavorRef,
		FlavorName:       flavorName,
		SecurityGroups:   securityGroups,
		UserData:         userData,
		AvailabilityZone: availabilityZone,
		Networks:         networks_fake,
		Metadata:         metadata_fake,
		Personality:      personality_fake,
		ConfigDrive:      &configDrive,
		AdminPass:        adminPass,
		AccessIPv4:       accessIpv4,
		AccessIPv6:       accessIpv6,
		ServiceClient:    serviceClient,
	}
	opts = servers.CreateOpts{
		Name:          name,
		ImageRef:      imageid,
		ImageName:     imagename,
		FlavorRef:     flavorid,
		FlavorName:    flavorname,
		Networks:      networks,
		ServiceClient: client,
	}
	if 0 != len(securitygroups) {
		opts.SecurityGroups = securitygroups
	}
	if 0 != len(userdata) {
		opts.UserData = userdata
	}
	if 0 != len(personality) {
		opts.Personality = personality
	}
	if 0 != len(adminpass) {
		opts.AdminPass = adminpass
	}
	if 0 != len(accessipv4) {
		opts.AccessIPv4 = accessipv4
	}

	result, err := servers.Create(client, opts).Extract()
	if nil != err {
		glog.Errorf("Could not reap compute spawn capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute spawn capability: %v", result)

	return result, nil
}

func (provider *InfraProvider) deleteMachine(id string) error {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	if err := servers.Delete(client, id).ExtractErr(); nil != err {
		glog.Errorf("Could not reap compute removing capability: %v", err)
		return err
	}
	glog.V(5).Infoln("Succeeded to reap compute removing capability")
	return nil
}

func (provider *InfraProvider) queryMachine(id string) (*servers.Server, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	resp, err := servers.Get(client, id).Extract()
	if nil != err {
		glog.Errorf("Could not reap compute/servers query capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/servers query capability: %v", resp)

	return resp, nil
}

func (provider *InfraProvider) restartMachine(id string, rebootmethod servers.RebootMethod) error {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	rebootMethod := servers.SoftReboot

	opts := &servers.RebootOpts{
		Type: rebootMethod,
	}
	opts.Type = rebootmethod

	if err := servers.Reboot(client, id, opts).ExtractErr(); nil != err {
		glog.Errorf("Could not reap compute/servers reboot capability: %v", err)
		return err
	}
	glog.V(5).Infof("Succeeded to reap compute/servers reboot capability")

	return nil
}

func (provider *InfraProvider) searchMachine(name string) ([]servers.Server, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	changesSince := "1970-12-31T11:59:59PM"
	image := "http://openstack.glance.api.endpoint/images/id"
	flavor := "http://openstack.nova.api.endpoint/flavors/id"
	nameFake := "cirros"
	status := "ACTIVE"
	host := "all of hypervisors"
	marker := "00000000-0000-0000-0000-000000000000"
	limit := 0
	alltenants := false

	opts := servers.ListOpts{
		ChangesSince: changesSince,
		Image:        image,
		Flavor:       flavor,
		Name:         nameFake,
		Status:       status,
		Host:         host,
		Marker:       marker,
		Limit:        limit,
		AllTenants:   alltenants,
	}
	opts = servers.ListOpts{
		Name:       name,
		AllTenants: alltenants,
	}

	resp, err := servers.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap compute/servers list capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/servers list capability: %v", resp)

	result, err := servers.ExtractServers(resp)
	if nil != err {
		glog.Errorf("Could not decode compute/servers result: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to decode compute/servers result: %v", result)

	return result, nil
}

func (provider *InfraProvider) identifyMachine(name string) (*string, error) {
	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	result, err := servers.IDFromName(client, name)
	if nil != err {
		glog.Errorf("Could not reap compute/servers search capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap compute/servers search capability: %v", result)

	return &result, nil
}

func (provider *InfraProvider) reapUsers() ([]users.User, error) {
	chosen, _, err := utils.ChooseVersion(provider.providerclient, []*utils.Version{
		{ID: "v2.0", Priority: 20, Suffix: "/v2.0/"},
		{ID: "v3.0", Priority: 30, Suffix: "/v3/"},
	})
	if err != nil {
		return nil, err
	}

	switch chosen.ID {
	case "v2.0":
		client, err := openstack.NewIdentityV2(provider.providerclient, gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		})
		if nil != err {
			glog.Errorf("Could not reap identity service: %v", err)
			return nil, err
		}
		glog.V(5).Infof("Reap identity serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

		resp, err := users.List(client).AllPages()
		if nil != err {
			glog.Errorf("Could not reap identity list capability: %v", err)
			return nil, err
		}
		glog.V(5).Infof("Succeeded to reap identity list capability: %v", resp)

		result, err := users.ExtractUsers(resp)
		if nil != err {
			glog.Errorf("Could not decode identity result: %v", err)
			return nil, err
		}
		glog.V(5).Infof("Succeeded to decode identity result: %v", result)

		return result, nil
	case "v3.0":
		return nil, fmt.Errorf("Not implemented for this identity version: %s", chosen.ID)
	default:
		// The switch statement must be out of date from the versions list.
		return nil, fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

func (provider *InfraProvider) createUser(username, tenantid string, enabled bool, email string) (*users.User, error) {
	chosen, _, err := utils.ChooseVersion(provider.providerclient, []*utils.Version{
		{ID: "v2.0", Priority: 20, Suffix: "/v2.0/"},
		{ID: "v3.0", Priority: 30, Suffix: "/v3/"},
	})
	if err != nil {
		return nil, err
	}

	switch chosen.ID {
	case "v2.0":
		client, err := openstack.NewIdentityV2(provider.providerclient, gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		})
		if nil != err {
			glog.Errorf("Could not reap identity service: %v", err)
			return nil, err
		}
		glog.V(5).Infof("Reap identity serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

		resp, err := users.Create(client, users.CreateOpts{
			Username: username,
			TenantID: tenantid,
			Enabled:  &enabled,
			Email:    email,
		}).Extract()
		if nil != err {
			glog.Errorf("Could not reap identity/user creation capability: %v", err)
			return nil, err
		}
		glog.V(5).Infof("Succeeded to reap identity/user creation capability: %v", resp)

		return resp, nil
	case "v3.0":
		return nil, fmt.Errorf("Not implemented for this identity version: %s", chosen.ID)
	default:
		// The switch statement must be out of date from the versions list.
		return nil, fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

func (provider *InfraProvider) reapTenantsOrProjects() ([]tenants.Tenant, []projects.Project, error) {
	chosen, _, err := utils.ChooseVersion(provider.providerclient, []*utils.Version{
		{ID: "v2.0", Priority: 20, Suffix: "/v2.0/"},
		{ID: "v3.0", Priority: 30, Suffix: "/v3/"},
	})
	if err != nil {
		return nil, nil, err
	}

	switch chosen.ID {
	case "v2.0":
		client, err := openstack.NewIdentityV2(provider.providerclient, gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		})
		if nil != err {
			glog.Errorf("Could not reap identity service: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Reap identity serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

		marker := "00000000-0000-0000-0000-000000000000"
		limit := 0
		opts := tenants.ListOpts{
			Marker: marker,
			Limit:  limit,
		}
		glog.V(9).Infoln(opts)

		resp, err := tenants.List(client, nil).AllPages()
		if nil != err {
			glog.Errorf("Could not reap identity/tenants list capability: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to reap identity/tenants list capability: %v", resp)

		result, err := tenants.ExtractTenants(resp)
		if nil != err {
			glog.Errorf("Could not decode identity/tenants result: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to decode identity/tenants result: %v", result)

		return result, nil, nil
	case "v3.0":
		client, err := openstack.NewIdentityV3(provider.providerclient, gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		})
		if nil != err {
			glog.Errorf("Could not reap identity service: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Reap identity serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

		domainid := "00000000-0000-0000-0000-000000000000"
		enabled := true
		isdomain := false
		name := "demo"
		parentid := "00000000-0000-0000-0000-000000000000"
		opts := projects.ListOpts{
			DomainID: domainid,
			Enabled:  &enabled,
			IsDomain: &isdomain,
			Name:     name,
			ParentID: parentid,
		}
		glog.V(9).Infoln(opts)

		resp, err := projects.List(client, nil).AllPages()
		if nil != err {
			glog.Errorf("Could not reap identity/projects list capability: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to reap identity/projects list capability: %v", resp)

		result, err := projects.ExtractProjects(resp)
		if nil != err {
			glog.Errorf("Could not decode identity/projects result: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to decode identity/projects result: %v", result)

		return nil, result, nil
	default:
		// The switch statement must be out of date from the versions list.
		return nil, nil, fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

func (provider *InfraProvider) searchProject(name string) ([]tenants.Tenant, []projects.Project, error) {
	chosen, _, err := utils.ChooseVersion(provider.providerclient, []*utils.Version{
		{ID: "v2.0", Priority: 20, Suffix: "/v2.0/"},
		{ID: "v3.0", Priority: 30, Suffix: "/v3/"},
	})
	if err != nil {
		return nil, nil, err
	}

	switch chosen.ID {
	case "v2.0":
		client, err := openstack.NewIdentityV2(provider.providerclient, gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		})
		if nil != err {
			glog.Errorf("Could not reap identity service: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Reap identity serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

		marker := "00000000-0000-0000-0000-000000000000"
		limit := 0
		opts := tenants.ListOpts{
			Marker: marker,
			Limit:  limit,
		}
		glog.V(9).Infoln(opts)

		resp, err := tenants.List(client, nil).AllPages()
		if nil != err {
			glog.Errorf("Could not reap identity/tenants list capability: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to reap identity/tenants list capability: %v", resp)

		result, err := tenants.ExtractTenants(resp)
		if nil != err {
			glog.Errorf("Could not decode identity/tenants result: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to decode identity/tenants result: %v", result)

		value := []tenants.Tenant{}
		for i := 0; i < len(result); i++ {
			if name == result[i].Name {
				value = append(value, result[i])
				break
			}
		}

		return value, nil, nil
	case "v3.0":
		client, err := openstack.NewIdentityV3(provider.providerclient, gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		})
		if nil != err {
			glog.Errorf("Could not reap identity service: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Reap identity serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

		domainid := "00000000-0000-0000-0000-000000000000"
		enabled := true
		isdomain := false
		// name := "demo"
		parentid := "00000000-0000-0000-0000-000000000000"
		opts := projects.ListOpts{
			DomainID: domainid,
			Enabled:  &enabled,
			IsDomain: &isdomain,
			Name:     name,
			ParentID: parentid,
		}
		glog.V(9).Infoln(opts)

		resp, err := projects.List(client, nil).AllPages()
		if nil != err {
			glog.Errorf("Could not reap identity/projects list capability: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to reap identity/projects list capability: %v", resp)

		result, err := projects.ExtractProjects(resp)
		if nil != err {
			glog.Errorf("Could not decode identity/projects result: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to decode identity/projects result: %v", result)

		return nil, result, nil
	default:
		// The switch statement must be out of date from the versions list.
		return nil, nil, fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

func (provider *InfraProvider) queryProject(id string) (*tenants.Tenant, *projects.Project, error) {
	chosen, _, err := utils.ChooseVersion(provider.providerclient, []*utils.Version{
		{ID: "v2.0", Priority: 20, Suffix: "/v2.0/"},
		{ID: "v3.0", Priority: 30, Suffix: "/v3/"},
	})
	if err != nil {
		return nil, nil, err
	}

	switch chosen.ID {
	case "v2.0":
		return nil, nil, fmt.Errorf("Not implemented identity version: %s", chosen.ID)
	case "v3.0":
		client, err := openstack.NewIdentityV3(provider.providerclient, gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		})
		if nil != err {
			glog.Errorf("Could not reap identity service: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Reap identity serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

		resp, err := projects.Get(client, id, projects.GetOpts{}).Extract()
		if nil != err {
			glog.Errorf("Could not reap identity/projects get capability: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to reap identity/projects get capability: %v", resp)

		return nil, resp, nil
	default:
		// The switch statement must be out of date from the versions list.
		return nil, nil, fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

func (provider *InfraProvider) createProject(domainid string, enabled, isdomain bool, name, parentid, description string) (*tenants.Tenant, *projects.Project, error) {
	chosen, _, err := utils.ChooseVersion(provider.providerclient, []*utils.Version{
		{ID: "v2.0", Priority: 20, Suffix: "/v2.0/"},
		{ID: "v3.0", Priority: 30, Suffix: "/v3/"},
	})
	if err != nil {
		return nil, nil, err
	}

	switch chosen.ID {
	case "v2.0":
		return nil, nil, fmt.Errorf("Not implemented identity version: %s", chosen.ID)
	case "v3.0":
		client, err := openstack.NewIdentityV3(provider.providerclient, gophercloud.EndpointOpts{
			Region: os.Getenv("OS_REGION_NAME"),
		})
		if nil != err {
			glog.Errorf("Could not reap identity service: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Reap identity serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

		//		domainid := "00000000-0000-0000-0000-000000000000"
		//		enabled := true
		//		isdomain := false
		//		name := "demo"
		//		parentid := "00000000-0000-0000-0000-000000000000"
		//		description := "project"
		opts := projects.CreateOpts{
			DomainID:    domainid,
			Enabled:     &enabled,
			IsDomain:    &isdomain,
			Name:        name,
			ParentID:    parentid,
			Description: description,
		}
		glog.V(9).Infoln(opts)

		resp, err := projects.Create(client, nil).Extract()
		if nil != err {
			glog.Errorf("Could not reap identity/projects get capability: %v", err)
			return nil, nil, err
		}
		glog.V(5).Infof("Succeeded to reap identity/projects get capability: %v", resp)

		return nil, resp, nil
	default:
		// The switch statement must be out of date from the versions list.
		return nil, nil, fmt.Errorf("Unrecognized identity version: %s", chosen.ID)
	}
}

func (provider *InfraProvider) validateToken(id string) (*tokens.Token, error) {
	client, err := openstack.NewIdentityV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap identity service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap identity serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	resp, err := tokens.Get(client, id).ExtractToken()
	if nil != err {
		glog.Errorf("Could not reap identity query capability: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap identity query capability: %v", resp)

	return resp, nil
}

/*
func spawnVM() {
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
