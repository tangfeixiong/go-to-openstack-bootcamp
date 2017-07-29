package osctl

import (
	"encoding/json"
	"reflect"

	"github.com/golang/glog"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
	"github.com/mitchellh/mapstructure"

	pbnova "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack/nova"
)

func convertServerAddresses(v interface{}) *pbnova.Addresses {
	addresses := make([]*pbnova.Address, 0)
	switch reflect.TypeOf(v).Kind() {
	case reflect.Slice:
		result := reflect.ValueOf(v)
		for i := 0; i < result.Len(); i++ {
			item := result.Index(i).Interface()
			if reflect.TypeOf(item).Kind() == reflect.Map {
				addr := new(pbnova.Address)
				elem := reflect.ValueOf(item)
				for _, key := range elem.MapKeys() {
					mv := elem.MapIndex(key)
					mk := key.Interface().(string)
					switch mk {
					case "addr":
						addr.Addr = mv.Interface().(string)
					case "version":
						kind := reflect.TypeOf(mv.Interface()).Kind()
						if kind == reflect.Float64 {
							addr.Version = int32(mv.Interface().(float64))
						} else if kind == reflect.Int {
							addr.Version = int32(mv.Interface().(int))
						} else {
							addr.Version = 4
						}
					case "OS-EXT-IPS-MAC:mac_addr":
						addr.MacAddr = mv.Interface().(string)
					case "OS-EXT-IPS:type":
						addr.AssignedType = mv.Interface().(string)
					default:
						if val, err := json.Marshal(mv.Interface()); nil != err {
							glog.V(2).Infof("Could not do with OpenStack extension %v: %v", mk, mv)
						} else {
							addr.Ext[mk] = val
						}
					}
				}
				addresses = append(addresses, addr)
			} else {
				glog.Infof("Unknown server address: %v", v)
			}
		}
	default:
		// [map[OS-EXT-IPS-MAC:mac_addr:fa:16:3e:12:66:b0 OS-EXT-IPS:type:fixed addr:192.168.128.160 version:4] map[OS-EXT-IPS-MAC:mac_addr:fa:16:3e:12:66:b0 OS-EXT-IPS:type:floating addr:10.100.151.32 version:4]]
		glog.Infof("Unknown server address: %v", v)
	}

	return &pbnova.Addresses{
		Addresses: addresses,
	}
}

/*
{
    "server": {
        "OS-DCF:diskConfig": "AUTO",
        "OS-EXT-AZ:availability_zone": "nova",
        "OS-EXT-SRV-ATTR:host": "compute",
        "OS-EXT-SRV-ATTR:hostname": "new-server-test",
        "OS-EXT-SRV-ATTR:hypervisor_hostname": "fake-mini",
        "OS-EXT-SRV-ATTR:instance_name": "instance-00000001",
        "OS-EXT-SRV-ATTR:kernel_id": "",
        "OS-EXT-SRV-ATTR:launch_index": 0,
        "OS-EXT-SRV-ATTR:ramdisk_id": "",
        "OS-EXT-SRV-ATTR:reservation_id": "r-ov3q80zj",
        "OS-EXT-SRV-ATTR:root_device_name": "/dev/sda",
        "OS-EXT-SRV-ATTR:user_data": "IyEvYmluL2Jhc2gKL2Jpbi9zdQplY2hvICJJIGFtIGluIHlvdSEiCg==",
        "OS-EXT-STS:power_state": 1,
        "OS-EXT-STS:task_state": null,
        "OS-EXT-STS:vm_state": "active",
        "OS-SRV-USG:launched_at": "2017-02-14T19:23:59.895661",
        "OS-SRV-USG:terminated_at": null,
        "accessIPv4": "1.2.3.4",
        "accessIPv6": "80fe::",
        "addresses": {
            "private": [
                {
                    "OS-EXT-IPS-MAC:mac_addr": "aa:bb:cc:dd:ee:ff",
                    "OS-EXT-IPS:type": "fixed",
                    "addr": "192.168.0.3",
                    "version": 4
                }
            ]
        },
        "config_drive": "",
        "created": "2017-02-14T19:23:58Z",
        "description": null,
        "flavor": {
            "disk": 1,
            "ephemeral": 0,
            "extra_specs": {
                "hw:cpu_model": "SandyBridge",
                "hw:mem_page_size": "2048",
                "hw:cpu_policy": "dedicated"
            },
            "original_name": "m1.tiny.specs",
            "ram": 512,
            "swap": 0,
            "vcpus": 1
        },
        "hostId": "2091634baaccdc4c5a1d57069c833e402921df696b7f970791b12ec6",
        "host_status": "UP",
        "id": "9168b536-cd40-4630-b43f-b259807c6e87",
        "image": {
            "id": "70a599e0-31e7-49b7-b260-868f441e862b",
            "links": [
                {
                    "href": "http://openstack.example.com/6f70656e737461636b20342065766572/images/70a599e0-31e7-49b7-b260-868f441e862b",
                    "rel": "bookmark"
                }
            ]
        },
        "key_name": null,
        "links": [
            {
                "href": "http://openstack.example.com/v2.1/6f70656e737461636b20342065766572/servers/9168b536-cd40-4630-b43f-b259807c6e87",
                "rel": "self"
            },
            {
                "href": "http://openstack.example.com/6f70656e737461636b20342065766572/servers/9168b536-cd40-4630-b43f-b259807c6e87",
                "rel": "bookmark"
            }
        ],
        "locked": false,
        "metadata": {
            "My Server Name": "Apache1"
        },
        "name": "new-server-test",
        "os-extended-volumes:volumes_attached": [
            {
                "delete_on_termination": false,
                "id": "volume_id1"
            },
            {
                "delete_on_termination": false,
                "id": "volume_id2"
            }
        ],
        "progress": 0,
        "security_groups": [
            {
                "name": "default"
            }
        ],
        "status": "ACTIVE",
        "tags": [],
        "tenant_id": "6f70656e737461636b20342065766572",
        "updated": "2017-02-14T19:24:00Z",
        "user_id": "fake"
    }
}
*/
type Server struct {
	Host               string `mapstructure:"OS-EXT-SRV-ATTR:host"`
	Hostname           string `mapstructure:"OS-EXT-SRV-ATTR:hostname"`
	HypervisorHostname string `mapstructure:"OS-EXT-SRV-ATTR:hypervisor_hostname"`
	InstanceName       string `mapstructure:"OS-EXT-SRV-ATTR:instance_name"`
	KernelID           string `mapstructure:"OS-EXT-SRV-ATTR:kernel_id"`
	LaunchIndex        int    `mapstructure:"OS-EXT-SRV-ATTR:launch_index"`
	RAMDiskID          string `mapstructure:"OS-EXT-SRV-ATTR:ramdisk_id"`
	ReservationID      string `mapstructure:"OS-EXT-SRV-ATTR:reservation_id"`
	RootDeviceName     string `mapstructure:"OS-EXT-SRV-ATTR:root_device_name"`
	UserData           string `mapstructure:"OS-EXT-SRC-ATTR:user_data"`
	servers.Server
}

func deserializeServer(r servers.GetResult) (*Server, error) {
	if r.Err != nil {
		return nil, r.Err
	}

	var response struct {
		Server Server `mapstructure:"server"`
	}

	config := &mapstructure.DecoderConfig{
		DecodeHook: func(from reflect.Kind, to reflect.Kind, data interface{}) (interface{}, error) {
			if (from == reflect.String) && (to == reflect.Map) {
				return map[string]interface{}{}, nil
			}
			return data, nil
		},
		Result: &response,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(r.Body)
	if err != nil {
		return nil, err
	}

	return &response.Server, nil
}

func deserializeServers(page pagination.Page) ([]Server, error) {
	casted := page.(servers.ServerPage).Body

	var response struct {
		Servers []Server `mapstructure:"servers"`
	}

	config := &mapstructure.DecoderConfig{
		DecodeHook: func(from reflect.Kind, to reflect.Kind, data interface{}) (interface{}, error) {
			if (from == reflect.String) && (to == reflect.Map) {
				return map[string]interface{}{}, nil
			}
			return data, nil
		},
		Result: &response,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return nil, err
	}

	err = decoder.Decode(casted)

	return response.Servers, err
}

/*
{
    "router": {
        "admin_state_up": true,
        "availability_zone_hints": [],
        "availability_zones": [
            "nova"
        ],
        "description": "",
        "distributed": false,
        "external_gateway_info": {
            "enable_snat": true,
            "external_fixed_ips": [
                {
                    "ip_address": "172.24.4.6",
                    "subnet_id": "b930d7f6-ceb7-40a0-8b81-a425dd994ccf"
                },
                {
                    "ip_address": "2001:db8::9",
                    "subnet_id": "0c56df5d-ace5-46c8-8f4c-45fa4e334d18"
                }
            ],
            "network_id": "ae34051f-aa6c-4c75-abf5-50dc9ac99ef3"
        },
        "ha": false,
        "id": "f8a44de0-fc8e-45df-93c7-f79bf3b01c95",
        "name": "router1",
        "routes": [],
        "status": "ACTIVE",
        "project_id": "0bd18306d801447bb457a46252d82d13",
        "tenant_id": "0bd18306d801447bb457a46252d82d13"
    }
}
*/
