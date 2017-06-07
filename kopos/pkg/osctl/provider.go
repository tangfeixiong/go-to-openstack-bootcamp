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
	"github.com/gophercloud/gophercloud/openstack/identity/v2/users"
	"github.com/gophercloud/gophercloud/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
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

func (provider *InfraProvider) createNetwork(name, tenantid string, adminstateup, shared bool) (*networks.Network, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewNetworkV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap neutron service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	adminStateUp := true
	nameFake := "virtual network"
	sharedFake := true
	tenantId := "00000000-0000-0000-0000-000000000000"

	opts := networks.CreateOpts{
		AdminStateUp: &adminStateUp,
		Name:         nameFake,
		Shared:       &sharedFake,
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

func (provider *InfraProvider) createSubnet(networkid, cidr, name, tenantid string, allocationpools []subnets.AllocationPool, gatewayip string, ipversion gophercloud.IPVersion, enabledhcp bool, dnsnameservers []string, hostroutes []subnets.HostRoute) (*subnets.Subnet, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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
		glog.Errorf("Could not create neutron subnet: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a new neutron subnet: %v", subnet)

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
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	nameFake := "virtual router"
	adminStateUp := false
	distributedFake := false
	tenantId := "00000000-0000-0000-0000-000000000000"
	gatewayInfo := routers.GatewayInfo{NetworkID: "00000000-0000-0000-0000-000000000000"}

	opts := routers.CreateOpts{
		Name:         nameFake,
		AdminStateUp: &adminStateUp,
		Distributed:  &distributedFake,
		TenantID:     tenantId,
		GatewayInfo:  &gatewayInfo,
	}
	opts = routers.CreateOpts{
		Name:         name,
		AdminStateUp: &adminstateup,
		Distributed:  &distributed,
		TenantID:     tenantid,
		GatewayInfo:  &gatewayinfo,
	}

	router, err := routers.Create(client, opts).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron router: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a new neutron router: %v", router)

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
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	nameFake := "required"
	adminStateUp := false
	distributedFake := false
	gatewayInfo := routers.GatewayInfo{NetworkID: "required"}
	routesFake := []routers.Route{{NextHop: "ip addr", DestinationCIDR: "ip range"}}

	opts := routers.UpdateOpts{
		Name:         nameFake,
		AdminStateUp: &adminStateUp,
		Distributed:  &distributedFake,
		GatewayInfo:  &gatewayInfo,
		Routes:       routesFake,
	}
	opts = routers.UpdateOpts{
		Name:         name,
		AdminStateUp: &adminstateup,
		Distributed:  &distributed,
		GatewayInfo:  &gatewayinfo,
		Routes:       routes,
	}

	resp, err := routers.Update(client, networkid, opts).Extract()
	if nil != err {
		glog.Errorf("Could not modify neutron router: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to modify neutron router: %v", resp)

	return resp, nil
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
	}

	port, err := ports.Create(client, opts).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron port: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a new neutron port: %v", port)

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
		return nil, err
	}
	glog.V(5).Infof("Reap neutron serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	ifinfo, err := routers.AddInterface(client, routerid, routers.AddInterfaceOpts{
		PortID: portid,
	}).Extract()
	if nil != err {
		glog.Errorf("Could not create neutron interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to a new neutron interface: %v", ifinfo)

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

func (provider *InfraProvider) createSecGroupRule(direction rules.RuleDirection, ethertype rules.RuleEtherType, secgroupid string, portrangemax, portrangemin int, protocol rules.RuleProtocol, remotegroupid, remoteipprefix, tenantid string) (*rules.SecGroupRule, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewImageServiceV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap glance service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap glance serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	limit := 0
	marker := "00000000-0000-0000-0000-000000000000"
	nameFake := "search condition"
	visibility := images.ImageVisibilityPublic
	memberStatus := images.ImageMemberStatusAll
	ownerFake := "uploader"
	status := images.ImageStatusActive
	sizeMin := int64(100)
	sizeMax := int64(10000)
	sortKey := "id"
	sortDir := "ASC"
	tagFake := ""

	opts := images.ListOpts{
		Limit:        limit,
		Marker:       marker,
		Name:         nameFake,
		Visibility:   visibility,
		MemberStatus: memberStatus,
		Owner:        ownerFake,
		Status:       status,
		SizeMin:      sizeMin,
		SizeMax:      sizeMax,
		SortKey:      sortKey,
		SortDir:      sortDir,
		Tag:          tagFake,
	}
	opts = images.ListOpts{
		Visibility:   visibility,
		MemberStatus: memberStatus,
		Status:       status,
	}

	result, err := images.List(client, opts).AllPages()
	if nil != err {
		glog.Errorf("Could not reap glance search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap glance search interface: %v", result)

	images, err := images.ExtractImages(result)
	if nil != err {
		glog.Errorf("Could not reap glance search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap glance search interface: %v", images)

	return images, nil
}

func (provider *InfraProvider) searchImages(name string) ([]images.Image, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewImageServiceV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap glance service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap glance serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	resp, err := images.List(client, images.ListOpts{
		// Limit:      limit,
		// Marker:      marker,
		Name: name,
		// Visibility:   images.ImageVisibilityPublic,
		// MemberStatus: images.ImageMemberStatusAll,
		// Owner:       owner,
		// Status: images.ImageStatusActive,
		// SizeMin: sizemin,
		// SizeMax: sizemax,
		// SortKey: sortkey,
		// SortDir: sortdir,
		// Tag:       tag,
	}).AllPages()
	if nil != err {
		glog.Errorf("Could not reap glance search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap glance search interface: %v", resp)

	result, err := images.ExtractImages(resp)
	if nil != err {
		glog.Errorf("Could not reap glance search interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap glance search interface: %v", result)

	return result, nil
}

func (provider *InfraProvider) queryImage(id string) (*images.Image, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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
		glog.Errorf("Could not reap glance reap interface: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Succeeded to reap glance reap interface: %v", result)

	return result, nil
}

func (provider *InfraProvider) gainComputeImages() ([]computeimages.Image, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap glance service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap glance serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	kind := "SERVER" // "BASE" "SERVER" "ALL"
	result, err := computeimages.ListDetail(client, computeimages.ListOpts{
		// ChangesSince: changesince,
		// Limit:      limit,
		// Marker:      marker,
		// Name: name,
		// Server:       server,
		Status: string(images.ImageStatusActive),
		Type:   kind,
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
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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

func (provider *InfraProvider) reapFlavors() ([]flavors.Flavor, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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

func (provider *InfraProvider) reapMachines() ([]servers.Server, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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

func (provider *InfraProvider) createMachine(name, imagename, flavorname string, networks []servers.Network) (*servers.Server, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return nil, err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	nameFake := "new vm name"
	imageref := "00000000-0000-0000-0000-000000000000"
	imageName := "cirros"
	flavorref := "00000000-0000-0000-0000-000000000000"
	flavorName := "small"
	securityGroups := []string{"default", "customization"}
	userData := []byte("/bin/cat /etc/OS-RELEASE")
	availabilityzone := "no zone"
	networksFake := []servers.Network{{
		UUID:    "", // Required unless Port is provided
		Port:    "00000000-0000-0000-0000-000000000000",
		FixedIP: "optional, maybe using DHCP",
	}}
	metadata := map[string]string{"key": "value"}
	personality := []*servers.File{&servers.File{Path: "path-to-injected-file", Contents: []byte("contents-of-injected-file")}}
	configDrive := false
	adminpass := "secret"
	accessipv4 := "0.0.0.0"
	accessipv6 := "::0"

	opts := servers.CreateOpts{
		Name:             nameFake,
		ImageRef:         imageref,
		ImageName:        imageName,
		FlavorRef:        flavorref,
		FlavorName:       flavorName,
		SecurityGroups:   securityGroups,
		UserData:         userData,
		AvailabilityZone: availabilityzone,
		Networks:         networksFake,
		Metadata:         metadata,
		Personality:      personality,
		ConfigDrive:      &configDrive,
		AdminPass:        adminpass,
		AccessIPv4:       accessipv4,
		AccessIPv6:       accessipv6,
		ServiceClient:    client,
	}
	opts = servers.CreateOpts{
		Name:          name,
		ImageName:     imagename,
		FlavorName:    flavorname,
		Networks:      networks,
		ServiceClient: client,
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
	if nil != provider.lasterr {
		return provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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

func (provider *InfraProvider) searchMachine(name string) ([]servers.Server, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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

func (provider *InfraProvider) identifyMachine(name string) (string, error) {
	if nil != provider.lasterr {
		return "", provider.lasterr
	}

	client, err := openstack.NewComputeV2(provider.providerclient, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	if nil != err {
		glog.Errorf("Could not reap compute service: %v", err)
		return "", err
	}
	glog.V(5).Infof("Reap compute serivce endpoint: %v%v", client.Endpoint, client.ResourceBase)

	resp, err := servers.IDFromName(client, name)
	if nil != err {
		glog.Errorf("Could not reap compute/servers search capability: %v", err)
		return "", err
	}
	glog.V(5).Infof("Succeeded to reap compute/servers search capability: %v", resp)

	return resp, nil
}

func (provider *InfraProvider) reapUsers() ([]users.User, error) {
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, nil, provider.lasterr
	}

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
	if nil != provider.lasterr {
		return nil, nil, provider.lasterr
	}

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
