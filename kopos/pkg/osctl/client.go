package osctl

import (
	"fmt"
	// "sync"
	// "os"
	"strings"
	"time"

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
	// "github.com/pborman/uuid"

	pbos "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack"
	pbkeystone "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack/identity"
	pbneutron "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack/neutron"
	pbnova "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack/nova"
)

type InfraResClient struct {
	InfraProvider
	config *CounterConfig
}

func InfraRes() *InfraResClient {
	return new(InfraResClient)
}

func (player *InfraResClient) Credential(config *CounterConfig) *InfraResClient {
	player.config = config
	return player
}

func (player *InfraResClient) DiscoverNetworks(discovery *pbos.NetworkDiscoveryReqRespData) (*pbos.NetworkDiscoveryReqRespData, error) {
	glog.V(2).Infoln("Go to list public networks")

	resp := new(pbos.NetworkDiscoveryReqRespData)
	resp.Status = discovery.Status
	resp.Networks = make([]*pbneutron.Network, 0)
	result, err := player.tryto().gainNetworks()
	if nil != err {
		return resp, err
	}
	for i := 0; i < len(result); i++ {
		resp.Networks = append(resp.Networks, &pbneutron.Network{
			Id:           result[i].ID,
			Name:         result[i].Name,
			AdminStateUp: result[i].AdminStateUp,
			Status:       result[i].Status,
			Subnets:      make([]*pbneutron.Subnet, 0),
			TenantId:     result[i].TenantID,
			Shared:       result[i].Shared,
		})
		for _, item := range result[i].Subnets {
			resp.Networks[i].Subnets = append(resp.Networks[i].Subnets, &pbneutron.Subnet{
				Id: item,
			})
		}
		//		value, err := player.tryto().searchSubnets(result[i].ID)
		//		if nil != err {
		//			return resp, err
		//		}
		//		for j := 0; j < len(value); j++ {
		//			resp.Networks[i].Subnetes = append(resp.Networks[i].Subnetes, &pbneutron.Subnet{
		//				Id:              value[j].ID,
		//				NetworkId:       value[j].NetworkID,
		//				Name:            value[j].Name,
		//				IpVersion:       value[j].IPVersion,
		//				Cidr:            value[j].CIDR,
		//				GatewayIp:       value[j].GatewayIP,
		//				DnsNameServers:  value[j].DNSNameservers,
		//				AllocationPools: make([]*pbneutron.AllocationPool, 0),
		//				HostRoutes:      make([]*pbneutron.HostRoute, 0),
		//				EnableDhcp:      value[j].EnableDHCP,
		//				TenantId:        value[j].TenantID,
		//			})
		//			for _, item := range value[j].AllocationPools {
		//				resp.Networks[i].Subnets[j].AllocationPools = append(resp.Networks[i].Subnets[j].AllocationPools, &pbneutron.AllocationPool{
		//					Start: item.Start,
		//					End:   item.End,
		//				})
		//			}
		//			for _, item := range value[j].HostRoutes {
		//				resp.Networks[i].Subnets[j].HostRoutes = append(resp.Networks[i].Subnets[j].HostRoutes, &pbneutron.HostRoute{
		//					DestinationCidr: item.DestinationCIDR,
		//					NextHop:         item.NextHop,
		//				})
		//			}
		//		}
	}

	return resp, nil
}

func (player *InfraResClient) DiscoverNetworkDetailed(discovery *pbneutron.Network) (*pbneutron.Network, error) {
	glog.V(2).Infoln("Go to discover network detailed")

	resp := new(pbneutron.Network)
	result, err := player.tryto().queryNetwork(discovery.Id)
	if nil != err {
		return resp, err
	}
	if nil == result {
		return resp, fmt.Errorf("Could not decide network id")
	}

	resp = &pbneutron.Network{
		Id:           result.ID,
		Name:         result.Name,
		AdminStateUp: result.AdminStateUp,
		Status:       result.Status,
		Subnets:      make([]*pbneutron.Subnet, 0),
		TenantId:     result.TenantID,
		Shared:       result.Shared,
	}
	for _, item := range result.Subnets {
		resp.Subnets = append(resp.Subnets, &pbneutron.Subnet{
			Id: item,
		})
	}

	return resp, nil
}

func (player *InfraResClient) DiscoverNetworkDetails(discovery *pbneutron.Network) (*pbneutron.Network, error) {
	glog.V(2).Infoln("Go to search network for details")

	resp := new(pbneutron.Network)
	result, err := player.tryto().searchNetworks(discovery.Name)
	if nil != err {
		return resp, err
	}

	if i := 0; i < len(result) {
		resp = &pbneutron.Network{
			Id:           result[i].ID,
			Name:         result[i].Name,
			AdminStateUp: result[i].AdminStateUp,
			Status:       result[i].Status,
			Subnets:      make([]*pbneutron.Subnet, 0),
			TenantId:     result[i].TenantID,
			Shared:       result[i].Shared,
		}
		for _, item := range result[i].Subnets {
			resp.Subnets = append(resp.Subnets, &pbneutron.Subnet{
				Id: item,
			})
		}
	}
	return resp, nil
}

func (player *InfraResClient) DiscoverSubnets(discovery *pbos.SubnetDiscoveryReqRespData) (*pbos.SubnetDiscoveryReqRespData, error) {
	glog.V(2).Infoln("Go to list subnets")

	resp := new(pbos.SubnetDiscoveryReqRespData)
	resp.NetworkId = discovery.NetworkId
	resp.Subnets = make([]*pbneutron.Subnet, 0)
	result, err := player.tryto().gainSubnets()
	if nil != err {
		return resp, err
	}
	for j := 0; j < len(result); j++ {
		resp.Subnets = append(resp.Subnets, &pbneutron.Subnet{
			Id:              result[j].ID,
			NetworkId:       result[j].NetworkID,
			Name:            result[j].Name,
			IpVersion:       int32(result[j].IPVersion),
			Cidr:            result[j].CIDR,
			GatewayIp:       result[j].GatewayIP,
			DnsNameServers:  result[j].DNSNameservers,
			AllocationPools: make([]*pbneutron.AllocationPool, 0),
			HostRoutes:      make([]*pbneutron.HostRoute, 0),
			EnableDhcp:      result[j].EnableDHCP,
			TenantId:        result[j].TenantID,
		})
		for _, item := range result[j].AllocationPools {
			resp.Subnets[j].AllocationPools = append(resp.Subnets[j].AllocationPools, &pbneutron.AllocationPool{
				Start: item.Start,
				End:   item.End,
			})
		}
		for _, item := range result[j].HostRoutes {
			resp.Subnets[j].HostRoutes = append(resp.Subnets[j].HostRoutes, &pbneutron.HostRoute{
				DestinationCidr: item.DestinationCIDR,
				NextHop:         item.NextHop,
			})
		}
	}

	return resp, nil
}

func (player *InfraResClient) CreateNetworkingLandscape(req *pbos.OpenstackNeutronLandscapeReqRespData) (*pbos.OpenstackNeutronLandscapeReqRespData, error) {
	glog.Infof("Go to establish network environments, input=%+v", req)
	resp := new(pbos.OpenstackNeutronLandscapeReqRespData)
	// To create openstack network and subnet
	// At first, if the network is listed in ifaces_info, treat as internal network.
	// Working as external network, verify its existing, or create it
	resp.Vnets = make([]*pbneutron.Network, 0)
	for i := 0; i < len(req.Vnets); i++ {
		if 0 != len(req.Vnets[i].Id) {
			k := -1
			for j := 0; j < len(req.IfacesInfo); j++ {
				if req.Vnets[i].Id == req.IfacesInfo[j].NetworkId {
					k = j
					break
				}
			}
			if k < 0 {
				resultvnet, err := player.tryto().queryNetwork(req.Vnets[i].Id)
				if nil != err {
					req.Vnets[i].Id = ""
				} else {
					req.Vnets[i].Name = resultvnet.Name
					// Treat as exiternal network
					resp.Vnets = append(resp.Vnets, &pbneutron.Network{
						Id:           resultvnet.ID,
						Name:         resultvnet.Name,
						AdminStateUp: resultvnet.AdminStateUp,
						Status:       resultvnet.Status,
						Subnets:      make([]*pbneutron.Subnet, 0),
						TenantId:     resultvnet.TenantID,
						Shared:       resultvnet.Shared,
					})
				}
			}
		}
		if 0 == len(req.Vnets[i].Id) && 0 != len(req.Vnets[i].Name) {
			k := -1
			for j := 0; j < len(req.IfacesInfo); j++ {
				if req.Vnets[i].Name == req.IfacesInfo[j].NetworkName {
					k = j
					break
				}
			}
			if k < 0 {
				resultvnets, err := player.tryto().searchNetworks(req.Vnets[i].Name)
				if nil != err || 0 == len(resultvnets) {
					req.Vnets[i].Name = ""
				} else if 0 != len(resultvnets) {
					req.Vnets[i].Id = resultvnets[0].ID
					// Treat as exiternal network
					resp.Vnets = append(resp.Vnets, &pbneutron.Network{
						Id:           resultvnets[0].ID,
						Name:         resultvnets[0].Name,
						AdminStateUp: resultvnets[0].AdminStateUp,
						Status:       resultvnets[0].Status,
						Subnets:      make([]*pbneutron.Subnet, 0),
						TenantId:     resultvnets[0].TenantID,
						Shared:       resultvnets[0].Shared,
					})
				}
			}
		}
		if 0 == len(req.Vnets[i].Id) && 0 == len(req.Vnets[i].Name) {
			resp.StateCode = int32(1000 + i + 1)
			resp.StateMessage = "Could not decide network"
			return resp, fmt.Errorf("Could not decide network")
		}
		if 0 != len(req.Vnets[i].Id) {
			// For external network, get subnets
			resultsubnets, err := player.tryto().searchSubnets(req.Vnets[i].Id)
			if nil != err {
				resp.StateCode = int32(1000+i+1) * 1000
				resp.StateMessage = "Could not decide subnets"
				return resp, fmt.Errorf("Could not decide subnets")
			}
			for j := 0; j < len(resultsubnets); j++ {
				req.Vnets[i].Subnets = append(req.Vnets[i].Subnets, &pbneutron.Subnet{
					Id:              resultsubnets[j].ID,
					NetworkId:       resultsubnets[j].ID,
					Name:            resultsubnets[j].Name,
					IpVersion:       int32(resultsubnets[j].IPVersion),
					Cidr:            resultsubnets[j].CIDR,
					GatewayIp:       resultsubnets[j].GatewayIP,
					DnsNameServers:  resultsubnets[j].DNSNameservers,
					AllocationPools: make([]*pbneutron.AllocationPool, 0),
					HostRoutes:      make([]*pbneutron.HostRoute, 0),
					EnableDhcp:      resultsubnets[j].EnableDHCP,
					TenantId:        resultsubnets[j].TenantID,
				})
				for x := 0; x < len(resultsubnets[j].AllocationPools); x++ {
					req.Vnets[i].Subnets[len(req.Vnets[i].Subnets)-1].AllocationPools = append(req.Vnets[i].Subnets[len(req.Vnets[i].Subnets)-1].AllocationPools, &pbneutron.AllocationPool{
						Start: resultsubnets[j].AllocationPools[x].Start,
						End:   resultsubnets[j].AllocationPools[x].End,
					})
				}
				for y := 0; y < len(resultsubnets[j].HostRoutes); y++ {
					req.Vnets[i].Subnets[len(req.Vnets[i].Subnets)-1].HostRoutes = append(req.Vnets[i].Subnets[len(req.Vnets[i].Subnets)-1].HostRoutes, &pbneutron.HostRoute{
						DestinationCidr: resultsubnets[j].HostRoutes[y].DestinationCIDR,
						NextHop:         resultsubnets[j].HostRoutes[y].NextHop,
					})
				}
			}
			continue
		}
		// For internal network, create it plus subnets
		l2net, err := player.createNetworkingLandscapeNetwork(req.Vnets[i])
		if err != nil {
			resp.StateCode = int32(1000 + i + 1)
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Could not establish landscape network: %v", req.Vnets[i].Name)
		}
		req.Vnets[i].Id = l2net.Id
		resp.Vnets = append(resp.Vnets, l2net)
		for j := 0; j < len(req.Vnets[i].Subnets); j++ {
			req.Vnets[i].Subnets[j].NetworkId = l2net.Id
			l3subnet, err := player.createNetworkingLandscapeSubnet(req.Vnets[i].Subnets[j])
			if err != nil {
				resp.StateCode = int32((1000+i+1)*1000 + 1 + j)
				resp.StateMessage = err.Error()
				return resp, fmt.Errorf("Could not establish landscape subnetwork: %v", req.Vnets[i].Subnets[j].Name)
			}
			req.Vnets[i].Subnets[j].Id = l3subnet.Id
			resp.Vnets[i].Subnets = append(resp.Vnets[i].Subnets, l3subnet)
		}
	}

	// Create router
	resultvrouter, err := player.createNetworkingLandscapeRouter(req.Vrouter)
	if err != nil {
		resp.StateCode = 2001
		resp.StateMessage = err.Error()
		return resp, fmt.Errorf("Could not establish landscape router: %v", req.Vrouter.Name)
	}
	req.Vrouter.Id = resultvrouter.Id
	resp.Vrouter = resultvrouter

	// Create secgroup
	resultsecgroup, err := player.createNetworkingLandscapeSecGroup(req.Secgroup)
	if err != nil {
		resp.StateCode = 3001
		resp.StateMessage = err.Error()
		return resp, fmt.Errorf("Could not establish landscape security group: %v", req.Secgroup.Name)
	}
	req.Secgroup.Id = resultsecgroup.Id
	resp.Secgroup = resultsecgroup

	// Create security group rules
	if 0 == len(req.Secgroup.SecurityGroupRules) {
		// all IPv4 tcp/udp/icmp ingress

		direction := string(rules.DirIngress)
		ethertype := string(rules.EtherType4)
		secgroupid := resultsecgroup.Id
		portrangemin := int32(1)
		portrangemax := int32(65535)
		protocol := string(rules.ProtocolTCP)
		// remotegroupid := ""
		remoteipprefix := "0.0.0.0/0"
		tenantid := resultsecgroup.TenantId
		req.Secgroup.SecurityGroupRules = append(req.Secgroup.SecurityGroupRules, &pbneutron.SecGroupRule{
			Direction:       direction,
			Ethertype:       ethertype,
			SecurityGroupId: secgroupid,
			PortRangeMin:    portrangemin,
			PortRangeMax:    portrangemax,
			Protocol:        protocol,
			RemoteIpPrefix:  remoteipprefix,
			TenantId:        tenantid,
		})

		protocol = string(rules.ProtocolUDP)
		req.Secgroup.SecurityGroupRules = append(req.Secgroup.SecurityGroupRules, &pbneutron.SecGroupRule{
			Direction:       direction,
			Ethertype:       ethertype,
			SecurityGroupId: secgroupid,
			PortRangeMin:    portrangemin,
			PortRangeMax:    portrangemax,
			Protocol:        protocol,
			RemoteIpPrefix:  remoteipprefix,
			TenantId:        tenantid,
		})

		portrangemax = -1
		portrangemin = -1
		protocol = string(rules.ProtocolICMP)
		req.Secgroup.SecurityGroupRules = append(req.Secgroup.SecurityGroupRules, &pbneutron.SecGroupRule{
			Direction:       direction,
			Ethertype:       ethertype,
			SecurityGroupId: secgroupid,
			PortRangeMin:    portrangemin,
			PortRangeMax:    portrangemax,
			Protocol:        protocol,
			RemoteIpPrefix:  remoteipprefix,
			TenantId:        tenantid,
		})

		// remoteipprefix = "::/0"
	}
	for i := 0; i < len(req.Secgroup.SecurityGroupRules); i++ {
		req.Secgroup.SecurityGroupRules[i].SecurityGroupId = resultsecgroup.Id
		resultrule, err := player.createNetworkingLandscapeSecGroupRule(req.Secgroup.SecurityGroupRules[i])
		if err != nil {
			resp.StateCode = 3001*1000 + int32(i) + 1
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Could not establish landscape security group rule: %v", req.Secgroup.Name)
		}
		req.Secgroup.SecurityGroupRules[i].Id = resultrule.Id
		resp.Secgroup.SecurityGroupRules = append(resp.Secgroup.SecurityGroupRules, resultrule)
	}

	// Create port into internal subnet, and create interface into router
	resp.IfacesInfo = make([]*pbos.IfaceInfo, 0)
	resp.Ports = make([]*pbneutron.Port, 0)
	resp.InterfacesInfo = make([]*pbneutron.InterfaceInfo, 0)
	for i := 0; i < len(req.IfacesInfo); i++ {
	LOOP_IFACES:
		for j := 0; j < len(resp.Vnets); j++ {
			if req.IfacesInfo[i].NetworkName == resp.Vnets[j].Name {
				req.IfacesInfo[i].NetworkId = resp.Vnets[j].Id
				for k := 0; k < len(resp.Vnets[j].Subnets); k++ {
					if req.IfacesInfo[i].SubnetName == resp.Vnets[j].Subnets[k].Name {
						req.IfacesInfo[i].SubnetId = resp.Vnets[j].Subnets[k].Id
						break LOOP_IFACES
					}
				}
				break
			}
		}
		if req.IfacesInfo[i].RouterName == resp.Vrouter.Name {
			req.IfacesInfo[i].RouterId = resp.Vrouter.Id
		}
		req.IfacesInfo[i].SecgroupsInfo = make([]*pbos.SecGroupInfo, 0)
		req.IfacesInfo[i].SecgroupsInfo = append(req.IfacesInfo[i].SecgroupsInfo, &pbos.SecGroupInfo{
			Id:   resp.Secgroup.Id,
			Name: resp.Secgroup.Name,
		})
		if 0 != len(req.IfacesInfo[i].NetworkId) && 0 != len(req.IfacesInfo[i].SubnetId) && 0 != len(req.IfacesInfo[i].RouterId) {
			// Port
			port := &pbneutron.Port{
				NetworkId:      req.IfacesInfo[i].NetworkId,
				Name:           req.IfacesInfo[i].SubnetName,
				AdminStateUp:   false,
				FixedIps:       make([]*pbneutron.IP, 0),
				TenantId:       player.config.ProjectId,
				SecurityGroups: make([]string, 0),
			}
			port.FixedIps = append(port.FixedIps, &pbneutron.IP{
				SubnetId: req.IfacesInfo[i].SubnetId,
			})
			for j := 0; j < len(req.IfacesInfo[i].SecgroupsInfo); j++ {
				port.SecurityGroups = append(port.SecurityGroups, req.IfacesInfo[i].SecgroupsInfo[j].Id)
			}
			resultport, err := player.createNetworkingLandscapePort(port)
			if err != nil {
				resp.StateCode = 4000 + int32(i) + 1
				resp.StateMessage = err.Error()
				return resp, fmt.Errorf("Could not establish landscape port: %v", port.Name)
			}
			port.Id = resultport.Id
			req.IfacesInfo[i].PortId = resultport.Id
			resp.Ports = append(resp.Ports, resultport)

			// Interface
			iface := &pbneutron.InterfaceInfo{
				SubnetId: req.IfacesInfo[i].SubnetId,
				PortId:   resultport.Id,
				TenantId: player.config.ProjectId,
			}
			resultiface, err := player.createNetworkingLandscapeInterface(req.IfacesInfo[i].RouterId, iface)
			if err != nil {
				resp.StateCode = (4000+int32(i)+1)*1000 + 1
				resp.StateMessage = err.Error()
				return resp, fmt.Errorf("Could not establish landscape interface: %v", port.Name)
			}
			req.IfacesInfo[i].InterfaceInfoId = resultiface.Id
			resp.InterfacesInfo = append(resp.InterfacesInfo, resultiface)

			// Detailed iface info
			resp.IfacesInfo = append(resp.IfacesInfo, &pbos.IfaceInfo{
				RouterId:        req.IfacesInfo[i].RouterId,
				RouterName:      req.IfacesInfo[i].RouterName,
				NetworkId:       req.IfacesInfo[i].NetworkId,
				NetworkName:     req.IfacesInfo[i].NetworkName,
				SubnetId:        req.IfacesInfo[i].SubnetId,
				SubnetName:      req.IfacesInfo[i].SubnetName,
				SecgroupsInfo:   make([]*pbos.SecGroupInfo, 0),
				PortId:          req.IfacesInfo[i].PortId,
				PortName:        req.IfacesInfo[i].PortName,
				InterfaceInfoId: req.IfacesInfo[i].InterfaceInfoId,
			})
			for x := 0; x < len(req.IfacesInfo[i].SecgroupsInfo); x++ {
				resp.IfacesInfo[len(resp.IfacesInfo)-1].SecgroupsInfo = append(resp.IfacesInfo[len(resp.IfacesInfo)-1].SecgroupsInfo, &pbos.SecGroupInfo{
					Id:   req.IfacesInfo[i].SecgroupsInfo[x].Id,
					Name: req.IfacesInfo[i].SecgroupsInfo[x].Name,
				})
			}
		}
	}

	// Set up gateway
	resp.GatewaysInfo = make([]*pbos.GatewayInfo, 0)
	for i := 0; i < len(req.GatewaysInfo); i++ {
		for j := 0; j < len(resp.Vnets); j++ {
			if req.GatewaysInfo[i].NetworkName == resp.Vnets[j].Name {
				req.GatewaysInfo[i].NetworkId = resp.Vnets[j].Id
				break
			}
		}
		if req.GatewaysInfo[i].RouterName == resp.Vrouter.Name {
			req.GatewaysInfo[i].RouterId = resp.Vrouter.Id
		}
		if 0 != len(req.GatewaysInfo[i].NetworkId) && 0 != len(req.GatewaysInfo[i].RouterId) {
			// req.Vrouter.TenantId = player.config.ProjectId
			// req.Vrouter.Status = "ACTIVE"
			req.Vrouter.GatewayInfo = &pbneutron.GatewayInfo{
				NetworkId: req.GatewaysInfo[i].NetworkId,
			}
			// req.Vrouter.Routes = []*pbneutron.Route{}

			resultrouter, err := player.updateNetworkingLandscapeRouterExternalGatewayInfo(req.Vrouter)
			if err != nil {
				resp.StateCode = 2001 * 1000
				resp.StateMessage = err.Error()
				return resp, fmt.Errorf("Could not establish landscape gateway: %v", req.Vrouter.Name)
			}
			resp.Vrouter.GatewayInfo = &pbneutron.GatewayInfo{
				NetworkId: resultrouter.GatewayInfo.NetworkId,
			}

			// Record gateway info
			resp.GatewaysInfo = append(resp.GatewaysInfo, &pbos.GatewayInfo{
				NetworkId:   req.GatewaysInfo[i].NetworkId,
				NetworkName: req.GatewaysInfo[i].NetworkName,
				RouterId:    resultrouter.Id,
				RouterName:  resultrouter.Name,
			})
		}
	}
	return resp, nil
}

func (player *InfraResClient) createNetworkingLandscapeNetwork(network *pbneutron.Network) (*pbneutron.Network, error) {
	glog.Infoln("Go to create network")
	name := network.Name
	tenantid := player.config.ProjectId
	adminstateup := false
	shared := false
	if 0 != len(network.TenantId) {
		tenantid = network.TenantId
	}
	adminstateup = network.AdminStateUp
	shared = network.Shared
	resultnet, err := player.tryto().createNetwork(name, tenantid, adminstateup, shared)
	if nil != err {
		return nil, err
	}

	l2net := &pbneutron.Network{
		Id:           resultnet.ID,
		Name:         resultnet.Name,
		AdminStateUp: resultnet.AdminStateUp,
		Status:       resultnet.Status,
		Subnets:      make([]*pbneutron.Subnet, 0),
		TenantId:     resultnet.TenantID,
		Shared:       resultnet.Shared,
	}
	return l2net, nil
}

func (player *InfraResClient) createNetworkingLandscapeSubnet(subnet *pbneutron.Subnet) (*pbneutron.Subnet, error) {
	glog.Infoln("Go to create subnet")
	l2netID := subnet.NetworkId
	name := subnet.Name
	tenantid := player.config.ProjectId
	subnetcidr := subnet.Cidr
	offset := strings.LastIndex(subnetcidr, ".")
	ipprefix := subnetcidr[:offset]
	allocstart := ipprefix + ".100"
	allocend := ipprefix + ".199"
	allocpools := []subnets.AllocationPool{{allocstart, allocend}}
	gatewayip := ipprefix + ".1"
	ipversion := gophercloud.IPv4
	enabledhcp := true
	dnsnameservers := []string{}
	hostroutes := []subnets.HostRoute{}
	if 0 != len(subnet.TenantId) {
		tenantid = subnet.TenantId
	}
	for _, item := range subnet.AllocationPools {
		allocpools = append(allocpools, subnets.AllocationPool{
			Start: item.Start,
			End:   item.End,
		})
	}
	if 1 < len(subnet.AllocationPools) {
		allocpools = allocpools[1:]
	}
	if 0 != len(subnet.GatewayIp) {
		gatewayip = subnet.GatewayIp
	}
	if 0 != subnet.IpVersion {
		ipversion = gophercloud.IPVersion(subnet.IpVersion)
	}
	enabledhcp = subnet.EnableDhcp
	dnsnameservers = append(dnsnameservers, subnet.DnsNameServers...)
	for _, item := range subnet.HostRoutes {
		hostroutes = append(hostroutes, subnets.HostRoute{
			NextHop:         item.NextHop,
			DestinationCIDR: item.DestinationCidr,
		})
	}

	resultsubnet, err := player.tryto().createSubnet(l2netID, subnetcidr, name, tenantid, allocpools, gatewayip, ipversion, enabledhcp, dnsnameservers, hostroutes)
	if nil != err {
		return nil, err
	}
	l3subnet := &pbneutron.Subnet{
		Id:              resultsubnet.ID,
		NetworkId:       resultsubnet.NetworkID,
		Name:            resultsubnet.Name,
		IpVersion:       int32(resultsubnet.IPVersion),
		Cidr:            resultsubnet.CIDR,
		GatewayIp:       resultsubnet.GatewayIP,
		DnsNameServers:  resultsubnet.DNSNameservers,
		AllocationPools: make([]*pbneutron.AllocationPool, 0),
		HostRoutes:      make([]*pbneutron.HostRoute, 0),
		EnableDhcp:      resultsubnet.EnableDHCP,
		TenantId:        resultsubnet.TenantID,
	}
	for x := 0; x < len(resultsubnet.AllocationPools); x++ {
		l3subnet.AllocationPools = append(l3subnet.AllocationPools, &pbneutron.AllocationPool{
			Start: resultsubnet.AllocationPools[x].Start,
			End:   resultsubnet.AllocationPools[x].End,
		})
	}
	for y := 0; y < len(resultsubnet.HostRoutes); y++ {
		l3subnet.HostRoutes = append(l3subnet.HostRoutes, &pbneutron.HostRoute{
			DestinationCidr: resultsubnet.HostRoutes[y].DestinationCIDR,
			NextHop:         resultsubnet.HostRoutes[y].NextHop,
		})
	}
	return l3subnet, nil
}

func (player *InfraResClient) createNetworkingLandscapeRouter(router *pbneutron.Router) (*pbneutron.Router, error) {
	glog.Infoln("Go to create router")
	name := router.Name
	adminstateup := true
	distributed := false
	tenantid := player.config.ProjectId
	gatewayinfo := routers.GatewayInfo{}

	adminstateup = router.AdminStateUp
	distributed = router.Distributed
	if 0 != len(router.TenantId) {
		tenantid = router.TenantId
	}
	if nil != router.GatewayInfo && 0 != len(router.GatewayInfo.NetworkId) {
		gatewayinfo.NetworkID = router.GatewayInfo.NetworkId
	}
	resultrouter, err := player.tryto().createRouter(name, adminstateup, distributed, tenantid, gatewayinfo)
	if nil != err {
		return nil, err
	}
	vrouter := &pbneutron.Router{
		Status:       resultrouter.Status,
		GatewayInfo:  new(pbneutron.GatewayInfo),
		AdminStateUp: resultrouter.AdminStateUp,
		Distributed:  resultrouter.Distributed,
		Name:         resultrouter.Name,
		Id:           resultrouter.ID,
		TenantId:     resultrouter.TenantID,
		Routes:       make([]*pbneutron.Route, 0),
	}
	if nil != vrouter.GatewayInfo {
		vrouter.GatewayInfo.NetworkId = resultrouter.GatewayInfo.NetworkID
	}
	for i := 0; i < len(resultrouter.Routes); i++ {
		vrouter.Routes = append(vrouter.Routes, &pbneutron.Route{
			NextHop:         resultrouter.Routes[i].NextHop,
			DestinationCidr: resultrouter.Routes[i].DestinationCIDR,
		})
	}
	return vrouter, nil
}

func (player *InfraResClient) createNetworkingLandscapeSecGroup(secgroup *pbneutron.SecGroup) (*pbneutron.SecGroup, error) {
	glog.V(2).Infoln("Go to create security group")
	name := secgroup.Name
	tenantid := player.config.ProjectId
	description := "The fight networking security group"
	if 0 != len(secgroup.TenantId) {
		tenantid = secgroup.TenantId
	}
	if 0 != len(secgroup.Description) {
		description = secgroup.Description
	}
	result, err := player.tryto().createSecurityGroup(name, tenantid, description)
	if nil != err {
		return nil, err
	}

	return &pbneutron.SecGroup{
		Id:                 result.ID,
		Name:               result.Name,
		Description:        result.Description,
		SecurityGroupRules: make([]*pbneutron.SecGroupRule, 0),
		TenantId:           result.TenantID,
	}, nil
}

func (player *InfraResClient) createNetworkingLandscapeSecGroupRule(rule *pbneutron.SecGroupRule) (*pbneutron.SecGroupRule, error) {
	glog.V(2).Infoln("Go to create security group rule")
	direction := rules.RuleDirection(rule.Direction)
	ethertype := rules.EtherType4
	switch rule.Ethertype {
	case "4", "v4", "IPv4", "ipv4", "IPV4":
		ethertype = rules.EtherType4
	case "6", "v6", "IPv6", "ipv6", "IPV6":
		ethertype = rules.EtherType6
	default:
		ethertype = rules.EtherType4
	}
	secgroupid := rule.SecurityGroupId
	portrangemax := int(rule.PortRangeMax)
	portrangemin := int(rule.PortRangeMin)
	protocol := rules.RuleProtocol(rule.Protocol)
	remotegroupid := ""
	remoteipprefix := rule.RemoteIpPrefix
	tenantid := rule.TenantId
	resultrule, err := player.tryto().createSecGroupRule(direction, ethertype, secgroupid, portrangemax, portrangemin, protocol, remotegroupid, remoteipprefix, tenantid)
	if nil != err {
		return nil, err
	}
	return &pbneutron.SecGroupRule{
		Id:              resultrule.ID,
		Direction:       resultrule.Direction,
		Ethertype:       resultrule.EtherType,
		SecurityGroupId: resultrule.SecGroupID,
		PortRangeMin:    int32(resultrule.PortRangeMin),
		PortRangeMax:    int32(resultrule.PortRangeMax),
		Protocol:        resultrule.Protocol,
		RemoteIpPrefix:  resultrule.RemoteIPPrefix,
		TenantId:        resultrule.TenantID,
	}, nil
}

func (player *InfraResClient) createNetworkingLandscapePort(port *pbneutron.Port) (*pbneutron.Port, error) {
	glog.V(2).Infoln("Go to create port")
	l2netID := port.NetworkId
	name := port.Name
	adminstateup := port.AdminStateUp
	macaddress := ""
	fixedips := []ports.IP{}
	deviceid := ""
	tenantid := player.config.ProjectId
	deviceowner := ""
	securitygroups := []string{}
	allowedaddresspairs := []ports.AddressPair{}
	if 0 != len(port.MacAddress) {
		macaddress = port.MacAddress
	}
	for _, item := range port.FixedIps {
		fixedips = append(fixedips, ports.IP{
			SubnetID:  item.SubnetId,
			IPAddress: item.IpAddress,
		})
	}
	if 0 != len(port.DeviceId) {
		deviceid = port.DeviceId
	}
	if 0 != len(port.TenantId) {
		tenantid = port.TenantId
	}
	if 0 != len(port.DeviceOwner) {
		deviceowner = port.DeviceOwner
	}
	securitygroups = append(securitygroups, port.SecurityGroups...)
	for _, item := range port.AllowedAddressPairs {
		allowedaddresspairs = append(allowedaddresspairs, ports.AddressPair{
			IPAddress:  item.IpAddress,
			MACAddress: item.MacAddress,
		})
	}
	resultport, err := player.tryto().createPort(l2netID, name, adminstateup, macaddress, fixedips, deviceid, deviceowner, tenantid, securitygroups, allowedaddresspairs)
	if nil != err {
		return nil, err
	}

	value := &pbneutron.Port{
		Id:                  resultport.ID,
		NetworkId:           resultport.NetworkID,
		Name:                resultport.Name,
		AdminStateUp:        resultport.AdminStateUp,
		Status:              resultport.Status,
		MacAddress:          resultport.MACAddress,
		FixedIps:            make([]*pbneutron.IP, 0),
		TenantId:            resultport.TenantID,
		DeviceOwner:         resultport.DeviceOwner,
		SecurityGroups:      resultport.SecurityGroups,
		DeviceId:            resultport.DeviceID,
		AllowedAddressPairs: make([]*pbneutron.AddressPair, 0),
	}
	for x := 0; x < len(resultport.FixedIPs); x++ {
		value.FixedIps = append(value.FixedIps, &pbneutron.IP{
			SubnetId:  resultport.FixedIPs[x].SubnetID,
			IpAddress: resultport.FixedIPs[x].IPAddress,
		})
	}
	for x := 0; x < len(resultport.AllowedAddressPairs); x++ {
		value.AllowedAddressPairs = append(value.AllowedAddressPairs, &pbneutron.AddressPair{
			IpAddress:  resultport.AllowedAddressPairs[x].IPAddress,
			MacAddress: resultport.AllowedAddressPairs[x].MACAddress,
		})
	}
	return value, nil
}

func (player *InfraResClient) createNetworkingLandscapeInterface(routerid string, iface *pbneutron.InterfaceInfo) (*pbneutron.InterfaceInfo, error) {
	glog.V(2).Infoln("Go to create interface")
	routerId := routerid
	portid := iface.PortId
	resultinfo, err := player.tryto().plugRouterIntoSubnet(routerId, portid)
	if nil != err {
		return nil, err
	}
	return &pbneutron.InterfaceInfo{
		Id:       resultinfo.ID,
		SubnetId: resultinfo.SubnetID,
		PortId:   resultinfo.PortID,
		TenantId: resultinfo.TenantID,
	}, nil
}

func (player *InfraResClient) updateNetworkingLandscapeRouterExternalGatewayInfo(router *pbneutron.Router) (*pbneutron.Router, error) {
	glog.V(2).Infoln("Go to create/update external gateway info")
	id := router.Id
	name := router.Name
	adminstateup := router.AdminStateUp
	distributed := router.Distributed
	// tenantid := router.TenantId
	gatewayinfo := routers.GatewayInfo{
		NetworkID: router.GatewayInfo.NetworkId,
	}
	routes := []routers.Route{}
	for i := 0; i < len(router.Routes); i++ {
		routes = append(routes, routers.Route{
			NextHop:         router.Routes[i].NextHop,
			DestinationCIDR: router.Routes[i].DestinationCidr,
		})
	}
	resultrouter, err := player.tryto().upstreamRouter(id, name, adminstateup, distributed, gatewayinfo, routes)
	if nil != err {
		return nil, err
	}
	vrouter := &pbneutron.Router{
		Status: resultrouter.Status,
		GatewayInfo: &pbneutron.GatewayInfo{
			NetworkId: resultrouter.GatewayInfo.NetworkID,
		},
		AdminStateUp: resultrouter.AdminStateUp,
		Distributed:  resultrouter.Distributed,
		Name:         resultrouter.Name,
		Id:           resultrouter.ID,
		TenantId:     resultrouter.TenantID,
		Routes:       []*pbneutron.Route{},
	}
	for _, item := range resultrouter.Routes {
		vrouter.Routes = append(vrouter.Routes, &pbneutron.Route{
			NextHop:         item.NextHop,
			DestinationCidr: item.DestinationCIDR,
		})
	}
	return vrouter, nil
}

func (player *InfraResClient) DiscoverImages(discovery *pbos.ImageDiscoveryReqRespData) (*pbos.ImageDiscoveryReqRespData, error) {
	glog.V(2).Infoln("Go to list active images")

	resp := new(pbos.ImageDiscoveryReqRespData)
	resp.Status = discovery.Status
	resp.Visibility = discovery.Visibility
	resp.MemberStatus = discovery.MemberStatus
	resp.Images = make([]*pbos.Image, 0)
	//	resultimages, err := player.tryto().gainImages()
	resultimages, err := player.tryto().gainComputeImages()
	if nil != err {
		return resp, err
	}
	for i := 0; i < len(resultimages); i++ {
		//		resp.Images = append(resp.Images, &pbos.Image{
		//			Id:              resultimages[i].ID,
		//			Name:            resultimages[i].Name,
		//			Status:          string(resultimages[i].Status),
		//			Tags:            resultimages[i].Tags,
		//			ContainerFormat: resultimages[i].ContainerFormat,
		//			DiskFormat:      resultimages[i].DiskFormat,
		//			MinDisk:         int32(resultimages[i].MinDiskGigabytes),
		//			MinRam:          int32(resultimages[i].MinRAMMegabytes),
		//			Owner:           resultimages[i].Owner,
		//			Protected:       resultimages[i].Protected,
		//			Visibility:      string(resultimages[i].Visibility),
		//			Checksum:        resultimages[i].Checksum,
		//			Size_:           resultimages[i].SizeBytes,
		//			Metadata:        resultimages[i].Metadata,
		//			Properties:      resultimages[i].Properties,
		//			CreateAt:        resultimages[i].CreatedAt.Format(time.RFC3339),
		//			UpdatedAt:       resultimages[i].UpdatedAt.Format(time.RFC3339),
		//			File:            resultimages[i].File,
		//			Schema:          resultimages[i].Schema,
		//		})
		resp.Images = append(resp.Images, &pbos.Image{
			Id:      resultimages[i].ID,
			Name:    resultimages[i].Name,
			Status:  resultimages[i].Status,
			MinDisk: int32(resultimages[i].MinDisk),
			MinRam:  int32(resultimages[i].MinRAM),
			// Metadata:        resultimages[i].Metadata,
			CreateAt:  resultimages[i].Created,
			UpdatedAt: resultimages[i].Updated,
		})
	}

	return resp, nil
}

func (player *InfraResClient) DiscoverImageDetailed(discovery *pbos.Image) (*pbos.Image, error) {
	glog.V(2).Infoln("Go to discover image detailed")

	resp := new(pbos.Image)
	resultimage, err := player.tryto().queryImage(discovery.Id)
	if nil != err {
		return resp, err
	}
	if nil == resultimage {
		return resp, fmt.Errorf("Could not decide image id")
	}

	resp = &pbos.Image{
		Id:              resultimage.ID,
		Name:            resultimage.Name,
		Status:          string(resultimage.Status),
		Tags:            resultimage.Tags,
		ContainerFormat: resultimage.ContainerFormat,
		DiskFormat:      resultimage.DiskFormat,
		MinDisk:         int32(resultimage.MinDiskGigabytes),
		MinRam:          int32(resultimage.MinRAMMegabytes),
		Owner:           resultimage.Owner,
		Protected:       resultimage.Protected,
		Visibility:      string(resultimage.Visibility),
		Checksum:        resultimage.Checksum,
		Size_:           resultimage.SizeBytes,
		Metadata:        resultimage.Metadata,
		Properties:      resultimage.Properties,
		CreateAt:        resultimage.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       resultimage.UpdatedAt.Format(time.RFC3339),
		File:            resultimage.File,
		Schema:          resultimage.Schema,
	}

	return resp, nil
}

func (player *InfraResClient) DiscoverImageDetails(discovery *pbos.Image) (*pbos.Image, error) {
	glog.V(2).Infoln("Go to search image for details")

	resp := new(pbos.Image)
	resultimages, err := player.tryto().searchImages(discovery.Name)
	if nil != err {
		return resp, err
	}

	if 0 != len(resultimages) {
		resp = &pbos.Image{
			Id:              resultimages[0].ID,
			Name:            resultimages[0].Name,
			Status:          string(resultimages[0].Status),
			Tags:            resultimages[0].Tags,
			ContainerFormat: resultimages[0].ContainerFormat,
			DiskFormat:      resultimages[0].DiskFormat,
			MinDisk:         int32(resultimages[0].MinDiskGigabytes),
			MinRam:          int32(resultimages[0].MinRAMMegabytes),
			Owner:           resultimages[0].Owner,
			Protected:       resultimages[0].Protected,
			Visibility:      string(resultimages[0].Visibility),
			Checksum:        resultimages[0].Checksum,
			Size_:           resultimages[0].SizeBytes,
			Metadata:        resultimages[0].Metadata,
			Properties:      resultimages[0].Properties,
			CreateAt:        resultimages[0].CreatedAt.Format(time.RFC3339),
			UpdatedAt:       resultimages[0].UpdatedAt.Format(time.RFC3339),
			File:            resultimages[0].File,
			Schema:          resultimages[0].Schema,
		}
	}
	return resp, nil
}

func (player *InfraResClient) DiscoverFlavors(discovery *pbos.FlavorDiscoveryReqRespData) (*pbos.FlavorDiscoveryReqRespData, error) {
	glog.V(2).Infoln("Go to list public flavors")

	resp := new(pbos.FlavorDiscoveryReqRespData)
	resp.AccessType = discovery.AccessType
	resp.Flavors = make([]*pbnova.Flavor, 0)
	result, err := player.tryto().gainFlavors()
	if nil != err {
		return resp, err
	}
	for i := 0; i < len(result); i++ {
		resp.Flavors = append(resp.Flavors, &pbnova.Flavor{
			Id:         result[i].ID,
			Disk:       int32(result[i].Disk),
			Ram:        int32(result[i].RAM),
			Name:       result[i].Name,
			RxtxFactor: result[i].RxTxFactor,
			Swap:       int32(result[i].Swap),
			Vcpus:      int32(result[i].VCPUs),
			IsPublic:   result[i].IsPublic,
		})
	}

	return resp, nil
}

func (player *InfraResClient) DiscoverFlavorDetailed(discovery *pbos.Flavor) (*pbos.Flavor, error) {
	glog.V(2).Infoln("Go to discover flavor detailed")

	resp := new(pbos.Flavor)
	result, err := player.tryto().queryFlavor(discovery.Id)
	if nil != err {
		return resp, err
	}
	if nil == result {
		return resp, fmt.Errorf("Could not decide flavor id")
	}

	resp = &pbos.Flavor{
		Id:         result.ID,
		Disk:       int32(result.Disk),
		Ram:        int32(result.RAM),
		Name:       result.Name,
		RxtxFactor: result.RxTxFactor,
		Swap:       int32(result.Swap),
		Vcpus:      int32(result.VCPUs),
		IsPublic:   result.IsPublic,
	}

	return resp, nil
}

func (player *InfraResClient) DiscoverFlavorDetails(discovery *pbos.Flavor) (*pbos.Flavor, error) {
	glog.V(2).Infoln("Go to search flavor for details")

	resp := new(pbos.Flavor)
	result, err := player.tryto().gainFlavors()
	if nil != err {
		return resp, err
	}

	for i := 0; i < len(result); i++ {
		if discovery.Name == result[i].Name {
			resp = &pbos.Flavor{
				Id:         result[i].ID,
				Disk:       int32(result[i].Disk),
				Ram:        int32(result[i].RAM),
				Name:       result[i].Name,
				RxtxFactor: result[i].RxTxFactor,
				Swap:       int32(result[i].Swap),
				Vcpus:      int32(result[i].VCPUs),
				IsPublic:   result[i].IsPublic,
			}
			break
		}
	}
	return resp, nil
}

func (player *InfraResClient) DiscoverMachines(discovery *pbos.MachineDiscoveryReqRespData) (*pbos.MachineDiscoveryReqRespData, error) {
	glog.V(2).Infoln("Go to list vms")

	resp := new(pbos.MachineDiscoveryReqRespData)
	resp.Status = discovery.Status
	resp.Vms = make([]*pbnova.Server, 0)
	result, err := player.tryto().gainMachines()
	if nil != err {
		return resp, err
	}
	for i := 0; i < len(result); i++ {
		resp.Vms = append(resp.Vms, &pbnova.Server{
			Id:             result[i].ID,
			TenantId:       result[i].TenantID,
			UserId:         result[i].UserID,
			Name:           result[i].Name,
			Updated:        result[i].Updated.Format(time.RFC3339),
			Created:        result[i].Created.Format(time.RFC3339),
			HostId:         result[i].HostID,
			Status:         result[i].Status,
			Progress:       int32(result[i].Progress),
			AccessIPv4:     result[i].AccessIPv4,
			AccessIPv6:     result[i].AccessIPv6,
			Images:         make(map[string]*pbnova.Image),
			Flavors:        make(map[string]*pbnova.Flavor),
			Addresses:      make(map[string]*pbnova.Addresses),
			MetadataInfo:   make(map[string]string),
			Links:          make([]string, 0),
			KeyName:        result[i].KeyName,
			AdminPass:      result[i].AdminPass,
			SecurityGroups: make([]*pbnova.SecurityGroups, 0),
		})
	}

	return resp, nil
}

func (player *InfraResClient) DestroyMachines(discovery *pbos.MachineDestroyReqRespData) (*pbos.MachineDestroyReqRespData, error) {
	glog.V(2).Infoln("Go to delete vms")

	resp := new(pbos.MachineDestroyReqRespData)
	resp.Vms = make([]*pbos.IdNamePair, 0)
	for i := 0; i < len(discovery.Vms); i++ {
		id := discovery.Vms[i].Id
		name := discovery.Vms[i].Name
		if len(id) == 0 && len(name) > 0 {
			value, err := player.tryto().identifyMachine(discovery.Vms[i].Name)
			if nil != err {
				resp.StateCode += 1
				resp.StateMessage += err.Error() + ";"
			}
			id = *value
		}
		if len(id) > 0 {
			err := player.tryto().deleteMachine(id)
			if err != nil {
				resp.StateCode += 1
				resp.StateMessage += err.Error() + ";"
			}
		}
	}

	return resp, nil
}

func (player *InfraResClient) RebootMachines(discovery *pbos.MachineRebootReqRespData) (*pbos.MachineRebootReqRespData, error) {
	glog.V(2).Infoln("Go to delete vms")

	resp := new(pbos.MachineRebootReqRespData)
	resp.Vms = make([]*pbos.IdNamePair, 0)
	for i := 0; i < len(discovery.Vms); i++ {
		id := discovery.Vms[i].Id
		name := discovery.Vms[i].Name
		if len(id) == 0 && len(name) > 0 {
			value, err := player.tryto().identifyMachine(discovery.Vms[i].Name)
			if nil != err {
				resp.StateCode += 1
				resp.StateMessage += err.Error() + ";"
			}
			id = *value
		}
		if len(id) > 0 {
			err := player.tryto().restartMachine(id, servers.SoftReboot)
			if err != nil {
				resp.StateCode += 1
				resp.StateMessage += err.Error() + ";"
			}
		}
	}

	return resp, nil
}

func (player *InfraResClient) BootVirtualMachines(req *pbos.OpenstackNovaBootReqRespData) (*pbos.OpenstackNovaBootReqRespData, error) {
	glog.Infof("Go to boot virtual machines, input=%+v", req)
	resp := new(pbos.OpenstackNovaBootReqRespData)

	if 0 != len(req.FlavorId) {
		resultflavor, err := player.tryto().queryFlavor(req.FlavorId)
		if nil != err {
			req.FlavorId = ""
		} else {
			resp.FlavorId = resultflavor.ID
			resp.FlavorName = resultflavor.Name
		}
	}
	if 0 == len(req.FlavorId) && 0 != len(req.FlavorName) {
		id, err := player.tryto().identifyFlavor(req.FlavorName)
		if nil != err {
			resp.StateCode = 100
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Could not decide machine flavor")
		}
		req.FlavorId = *id
		resp.FlavorId = *id
		resp.FlavorName = req.FlavorName
	}
	if 0 == len(req.FlavorId) && 0 == len(req.FlavorName) {
		resp.StateCode = 101
		resp.StateMessage = "Could not decide machine flavor"
		return resp, fmt.Errorf("Could not decide machine flavor")
	}

	if 0 != len(req.ImageId) {
		resultimage, err := player.tryto().queryImage(req.ImageId)
		if nil != err {
			req.ImageId = ""
		} else {
			resp.ImageId = resultimage.ID
			resp.ImageName = resultimage.Name
		}
	}
	if 0 == len(req.ImageId) && 0 != len(req.ImageName) {
		id, err := player.tryto().identifyImage(req.ImageName)
		if nil != err {
			resp.StateCode = 200
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Could not decide machine image")
		}
		req.ImageId = *id
		resp.ImageId = *id
		resp.ImageName = req.ImageName
	}
	if 0 == len(req.ImageId) && 0 == len(req.ImageName) {
		resp.StateCode = 201
		resp.StateMessage = "Could not decide machine image"
		return resp, fmt.Errorf("Could not decide machine image")
	}

	if 0 == req.MinCount && 0 == req.MaxCount {
		req.MinCount = 1
		req.MaxCount = 1
	} else if 0 != req.MinCount && 0 == req.MaxCount {
		req.MaxCount = req.MinCount
	} else if 0 == req.MinCount && 0 != req.MaxCount {
		req.MinCount = req.MaxCount
	}

	resp.SecgroupsInfo = make([]*pbos.SecGroupInfo, 0)
	for i := 0; i < len(req.SecgroupsInfo); i++ {
		if 0 != len(req.SecgroupsInfo[i].Id) {
			resultsecgroup, err := player.tryto().querySecurityGroup(req.SecgroupsInfo[i].Id)
			if nil != err {
				req.SecgroupsInfo[i].Id = ""
			} else {
				req.SecgroupsInfo[i].Name = resultsecgroup.Name
				resp.SecgroupsInfo = append(resp.SecgroupsInfo, &pbos.SecGroupInfo{
					Id:   resultsecgroup.ID,
					Name: resultsecgroup.Name,
				})
			}
		}
		if 0 == len(req.SecgroupsInfo[i].Id) && 0 != len(req.SecgroupsInfo[i].Name) {
			name, err := player.tryto().identifySecurityGroup(req.SecgroupsInfo[i].Name)
			if nil != err {
				req.SecgroupsInfo[i].Name = ""
			} else {
				req.SecgroupsInfo[i].Name = *name
				resp.SecgroupsInfo = append(resp.SecgroupsInfo, &pbos.SecGroupInfo{
					Id:   req.SecgroupsInfo[i].Id,
					Name: *name,
				})
			}
		}
	}

	if 0 != len(req.NetworkId) {
		resultnetwork, err := player.tryto().queryNetwork(req.NetworkId)
		if nil != err {
			req.NetworkId = ""
		} else {
			resp.NetworkId = resultnetwork.ID
			resp.NetworkName = resultnetwork.Name
		}
	}
	if 0 == len(req.NetworkId) && 0 != len(req.NetworkName) {
		id, err := player.tryto().identifyNetwork(req.NetworkName)
		if nil != err {
			resp.StateCode = 300
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Could not decide machine network")
		}
		req.NetworkId = *id
		resp.NetworkId = *id
		resp.NetworkName = req.NetworkName
	}
	if 0 == len(req.NetworkId) && 0 == len(req.NetworkName) {
		resp.StateCode = 301
		resp.StateMessage = "Could not decide machine network"
		return resp, fmt.Errorf("Could not decide machine network")
	}

	if 0 != len(req.FloatingNetworkId) {
		resultnetwork, err := player.tryto().queryNetwork(req.FloatingNetworkId)
		if nil != err {
			req.FloatingNetworkId = ""
		} else {
			resp.FloatingNetworkId = resultnetwork.ID
			resp.FloatingNetworkName = resultnetwork.Name
		}
	}
	if 0 == len(req.FloatingNetworkId) && 0 != len(req.FloatingNetworkName) {
		id, err := player.tryto().identifyNetwork(req.FloatingNetworkName)
		if nil != err {
			resp.StateCode = 300
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Could not decide floating network")
		}
		req.FloatingNetworkId = *id
		resp.FloatingNetworkId = *id
		resp.FloatingNetworkName = req.NetworkName
	}
	if 0 == len(req.FloatingNetworkId) && 0 == len(req.FloatingNetworkName) {
		resp.StateCode = 301
		resp.StateMessage = "Could not decide floating network"
		return resp, fmt.Errorf("Could not decide floating network")
	}

	resp.Ports = make([]*pbneutron.Port, 0)
	port := &pbneutron.Port{
		AdminStateUp:        true,
		MacAddress:          "",
		FixedIps:            []*pbneutron.IP{},
		DeviceId:            "",
		DeviceOwner:         "",
		TenantId:            player.config.ProjectId,
		SecurityGroups:      []string{},
		AllowedAddressPairs: []*pbneutron.AddressPair{},
		NetworkId:           req.NetworkId,
	}
	for i := 0; i < len(req.SecgroupsInfo); i++ {
		if len(req.SecgroupsInfo[i].Id) != 0 {
			port.SecurityGroups = append(port.SecurityGroups, req.SecgroupsInfo[i].Id)
		}
	}
	//	wg := &sync.WaitGroup{}
	for i := 0; i < int(req.MaxCount); i++ {
		port.Name = fmt.Sprintf("%s-%d", req.NamePrefix, i)
		//		go func() {
		//			wg.Add(1)
		//			defer wg.Done()
		resultport, err := player.createNetworkPortIntoActive(port)
		if nil != err {
			if int(req.MinCount) < i {
				break
			}
			resp.StateCode = 400
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Could not create enough network port, current=%d", i)
		}
		resp.Ports = append(resp.Ports, resultport)
		//		}()
	}
	//	wg.Wait()
	req.Ports = resp.Ports
	if len(req.Ports) < int(req.MaxCount) {
		req.MaxCount = int32(len(req.Ports))
		req.MaxCount = int32(len(req.Ports))
	}

	resp.Servers = make([]*pbnova.Server, 0)
	resp.PortServerPairs = make(map[string]string)
	for i := 0; i < int(req.MaxCount); i++ {
		servername := fmt.Sprintf("%s-%d", req.NamePrefix, i)
		imageref := req.ImageId
		imagename := req.ImageName
		flavorref := req.FlavorId
		flavorname := req.FlavorName
		securitygroups := []string{}
		for j := 0; j < len(req.SecgroupsInfo); j++ {
			securitygroups = append(securitygroups, req.SecgroupsInfo[j].Name)
		}
		userData := req.UserData
		networks := []servers.Network{{Port: req.Ports[i].Id}}
		personality := servers.Personality{}
		for j := 0; j < len(req.Personality); j++ {
			personality = append(personality, &servers.File{
				Path:     req.Personality[j].Path,
				Contents: req.Personality[j].Contents,
			})
		}
		adminPass := ""
		accessIpv4 := ""
		vm, err := player.tryto().createMachine(servername, imageref, imagename, flavorref, flavorname, securitygroups, userData, networks, personality, adminPass, accessIpv4)
		if nil != err {
			if int(req.MinCount) < i {
				break
			}
			resp.StateCode = 500
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Could not create enough vms, currnt=%d", i)
			return resp, err
		}
		resp.MaxCount = int32(i) + 1
		resp.Servers = append(resp.Servers, &pbnova.Server{
			Id:             vm.ID,
			TenantId:       vm.TenantID,
			UserId:         vm.UserID,
			Updated:        vm.Updated.Format(time.RFC3339),
			Created:        vm.Created.Format(time.RFC3339),
			HostId:         vm.HostID,
			Status:         vm.Status,
			Progress:       int32(vm.Progress),
			AccessIPv4:     vm.AccessIPv4,
			AccessIPv6:     vm.AccessIPv6,
			Images:         make(map[string]*pbnova.Image),
			Flavors:        make(map[string]*pbnova.Flavor),
			Addresses:      make(map[string]*pbnova.Addresses),
			MetadataInfo:   make(map[string]string),
			Links:          make([]string, 0),
			KeyName:        vm.KeyName,
			AdminPass:      vm.AdminPass,
			SecurityGroups: make([]*pbnova.SecurityGroups, 0),
		})
		resp.PortServerPairs[req.Ports[i].Id] = vm.ID
	}
	req.Servers = resp.Servers
	if len(req.Servers) < int(req.MaxCount) {
		req.MaxCount = int32(len(req.Servers))
		req.MaxCount = int32(len(req.Servers))
	}

	resp.FloatingIps = make([]*pbneutron.FloatingIP, 0)
	for i := 0; i < int(req.MaxCount); i++ {
		floatingnetworkid := req.FloatingNetworkId
		portid := req.Ports[i].Id
		resultfip, err := player.tryto().createFloatingIp(floatingnetworkid, "", portid, "", "")
		if nil != err {
			if int(req.MinCount) < i {
				break
			}
			resp.StateCode = 600
			resp.StateMessage = err.Error()
			return resp, fmt.Errorf("Could not create enough floating ip, currnt=%d", i)
			return resp, err
		}
		resp.MinCount = int32(i) + 1
		resp.FloatingIps = append(resp.FloatingIps, &pbneutron.FloatingIP{
			FloatingNetworkId: resultfip.FloatingNetworkID,
			FloatingIpAddress: resultfip.FloatingIP,
			PortId:            resultfip.PortID,
			FixedIpAddress:    resultfip.FixedIP,
			TenantId:          resultfip.TenantID,
		})
	}

	// watch spawn required

	return resp, nil
}

func (player *InfraResClient) createNetworkPortIntoActive(port *pbneutron.Port) (*pbneutron.Port, error) {
	glog.V(2).Infof("To create network port %s into active %s", port.Name, port.NetworkId)
	adminstateup := port.AdminStateUp
	macaddress := port.MacAddress
	fixedips := []ports.IP{}
	deviceid := port.DeviceId
	deviceowner := port.DeviceOwner
	tenantid := port.TenantId
	securitygroups := port.SecurityGroups
	allowedaddresspairs := []ports.AddressPair{}
	portname := port.Name
	networkid := port.NetworkId

	for _, item := range port.FixedIps {
		fixedips = append(fixedips, ports.IP{
			SubnetID:  item.SubnetId,
			IPAddress: item.IpAddress,
		})
	}
	for _, item := range port.AllowedAddressPairs {
		allowedaddresspairs = append(allowedaddresspairs, ports.AddressPair{
			IPAddress:  item.IpAddress,
			MACAddress: item.MacAddress,
		})
	}

	resultport, err := player.tryto().createPort(networkid, portname, adminstateup, macaddress, fixedips, deviceid, deviceowner, tenantid, securitygroups, allowedaddresspairs)
	if nil != err {
		return nil, err
	}
	id := resultport.ID
	c1 := make(chan error, 1)
	go func() {
		for i := 0; i < 15; i++ {
			time.Sleep(time.Second * 1)
			resultport, err = player.tryto().queryPort(id)
			if nil != err {
				c1 <- err
				return
			}
			if "ACTIVE" == port.Status {
				c1 <- nil
				return
			}
		}
		glog.Warningf("Port is unavailabe after check status %d times", 15)
		// c1 <- fmt.Errorf("Port is unavailabe after check status %d times", 15)
		c1 <- nil
	}()

	select {
	case res := <-c1:
		if nil != res {
			return nil, res
		}
	case <-time.After(time.Second * 15):
		glog.Warningf("Port is unavailabe after waiting for %d seconds", 15)
		// return nil, fmt.Errorf("Port is unavailabe after waiting for %d seconds", 15)
	}

	value := &pbneutron.Port{
		Id:                  resultport.ID,
		NetworkId:           resultport.NetworkID,
		Name:                resultport.Name,
		AdminStateUp:        resultport.AdminStateUp,
		Status:              resultport.Status,
		MacAddress:          resultport.MACAddress,
		FixedIps:            make([]*pbneutron.IP, 0),
		TenantId:            resultport.TenantID,
		DeviceOwner:         resultport.DeviceOwner,
		SecurityGroups:      resultport.SecurityGroups,
		DeviceId:            resultport.DeviceID,
		AllowedAddressPairs: make([]*pbneutron.AddressPair, 0),
	}
	for x := 0; x < len(resultport.FixedIPs); x++ {
		value.FixedIps = append(value.FixedIps, &pbneutron.IP{
			SubnetId:  resultport.FixedIPs[x].SubnetID,
			IpAddress: resultport.FixedIPs[x].IPAddress,
		})
	}
	for x := 0; x < len(resultport.AllowedAddressPairs); x++ {
		value.AllowedAddressPairs = append(value.AllowedAddressPairs, &pbneutron.AddressPair{
			IpAddress:  resultport.AllowedAddressPairs[x].IPAddress,
			MacAddress: resultport.AllowedAddressPairs[x].MACAddress,
		})
	}

	return value, nil
}

func (player *InfraResClient) CreateTargetDrone(name, tenantid, subnetcidr, gatewayip, description, imagename, flavorname string) {
	userData := []byte{}
	consoleNetworks := []servers.Network{}
	personality := servers.Personality{}
	adminPass := ""
	accessIpv4 := ""
	vm, err := player.tryto().createMachine(name, "", imagename, "", flavorname, []string{}, userData, consoleNetworks, personality, adminPass, accessIpv4)
	if nil != err {
		return
	}
	println(vm)

}

func (player *InfraResClient) tryto() *InfraResClient {
	var err error
	var opts gophercloud.AuthOptions

	if nil == player.providerclient {
		if "admin" == player.config.Username {
			if opts, err = openstack.AuthOptionsFromEnv(); nil != err {
				glog.Errorf("Could not load player openrc: %v", err)
				player.lasterr = err
				return player
			} else {
				glog.Infof("Load admin openrc: %v %v", opts.IdentityEndpoint, opts.Username)
				player.lasterr = nil
			}
		} else {
			opts = gophercloud.AuthOptions{
				IdentityEndpoint: player.config.IdentityHost,
				Username:         player.config.Username,
				Password:         player.config.Password,
				DomainName:       player.config.DomainName,
				TenantName:       player.config.ProjectName,
			}
		}

		if player.providerclient, err = openstack.AuthenticatedClient(opts); nil != err {
			glog.Errorf("Could not authenticate player: %v", err)
			player.lasterr = err
		} else {
			glog.Infof("Authenticated as token: %v", player.providerclient.TokenID)
			player.lasterr = nil
			player.config.Token = player.providerclient.TokenID
		}
		return player
	}

	if player.config.Token != player.providerclient.TokenID {
		if "admin" == player.config.Username {
			if opts, err = openstack.AuthOptionsFromEnv(); nil != err {
				glog.Errorf("Could not load player openrc: %v", err)
				// player.lasterr = err
				return player
			} else {
				glog.Infof("Load admin openrc: %v %v", opts.IdentityEndpoint, opts.Username)
			}
		} else {
			opts = gophercloud.AuthOptions{
				IdentityEndpoint: player.config.IdentityHost,
				Username:         player.config.Username,
				Password:         player.config.Password,
				DomainName:       player.config.DomainName,
				TenantName:       player.config.ProjectName,
			}
		}

		if player.providerclient, err = openstack.AuthenticatedClient(opts); nil != err {
			glog.Errorf("Could not authenticate player: %v", err)
			player.lasterr = err
		} else {
			glog.Infof("Authenticated as token: %v", player.providerclient.TokenID)
			player.lasterr = nil
			player.config.Token = player.providerclient.TokenID
		}
		return player
	}

	//	if nil != player.providerclient.lasterr {
	//		return nil, player.providerclient.lasterr
	//	}

	return player
}

func (player *InfraResClient) ValidateToken(in *pbos.TokenReqRespData) (*pbos.TokenReqRespData, error) {
	glog.Infoln("Going to validate token")
	out := in
	resulttoken, err := player.tryto().validateToken(player.tryto().providerclient.TokenID)
	if nil != err {
		return nil, err
	}
	out.Tgt = &pbkeystone.Token{
		Id:        resulttoken.ID,
		ExpiresAt: resulttoken.ExpiresAt.Format(time.RFC3339),
		Tenant: &pbkeystone.Tenant{
			Id:          resulttoken.Tenant.ID,
			Name:        resulttoken.Tenant.Name,
			Description: resulttoken.Tenant.Description,
			Enabled:     resulttoken.Tenant.Enabled,
		},
	}
	return out, nil
}
