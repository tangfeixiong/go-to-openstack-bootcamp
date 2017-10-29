# OpenStack Pike Installation － 创建虚拟机步骤

## Table of contents
控制节点
* 确认环境
* [套餐](#flavor)
* [公私钥](#keypair)
* [安全组](#security-grup)
* [准备虚拟网络](#dhcp-networking)
* [创建虚拟网络](#provider-network)
* [创建虚拟机](#launch)
计算节点
* [iproute2调试](#vm-networking)
* [arp和virsh调试][#debug]

## Controller

参考[Launch an instance](https://docs.openstack.org/install-guide/launch-instance.html)
```
[vagrant@controller-10-64-33-64 ~]$ . admin-openrc 
```

### 确认环境

Nova services
```
[vagrant@controller-10-64-33-64 ~]$ openstack compute service list
+----+------------------+------------------------+----------+---------+-------+----------------------------+
| ID | Binary           | Host                   | Zone     | Status  | State | Updated At                 |
+----+------------------+------------------------+----------+---------+-------+----------------------------+
|  7 | nova-compute     | compute-10-64-33-65    | nova     | enabled | up    | 2017-10-28T09:00:01.000000 |
|  8 | nova-consoleauth | controller-10-64-33-64 | internal | enabled | up    | 2017-10-28T09:00:00.000000 |
|  9 | nova-conductor   | controller-10-64-33-64 | internal | enabled | up    | 2017-10-28T09:00:08.000000 |
| 10 | nova-scheduler   | controller-10-64-33-64 | internal | enabled | up    | 2017-10-28T09:00:08.000000 |
+----+------------------+------------------------+----------+---------+-------+----------------------------+
```

Nova Controller
```
[vagrant@controller-10-64-33-64 ~]$ sudo nova-status upgrade check
+---------------------------+
| Upgrade Check Results     |
+---------------------------+
| Check: Cells v2           |
| Result: Success           |
| Details: None             |
+---------------------------+
| Check: Placement API      |
| Result: Success           |
| Details: None             |
+---------------------------+
| Check: Resource Providers |
| Result: Success           |
| Details: None             |
+---------------------------+
```

Neutron controller
```
[vagrant@controller-10-64-33-64 ~]$ openstack network agent list
+--------------------------------------+--------------------+------------------------+-------------------+-------+-------+---------------------------+
| ID                                   | Agent Type         | Host                   | Availability Zone | Alive | State | Binary                    |
+--------------------------------------+--------------------+------------------------+-------------------+-------+-------+---------------------------+
| 321066aa-bf65-4c43-9b3a-9994d97530b7 | Metadata agent     | controller-10-64-33-64 | None              | :-)   | UP    | neutron-metadata-agent    |
| 5ac11672-2f43-4504-a0c5-07d43f24b818 | DHCP agent         | controller-10-64-33-64 | nova              | :-)   | UP    | neutron-dhcp-agent        |
| a69362b5-9511-4e37-9c98-75babbf5ca2f | Linux bridge agent | compute-10-64-33-65    | None              | :-)   | UP    | neutron-linuxbridge-agent |
| b6e0c38b-fa23-4ad8-b42f-280aa600783d | Linux bridge agent | controller-10-64-33-64 | None              | :-)   | UP    | neutron-linuxbridge-agent |
+--------------------------------------+--------------------+------------------------+-------------------+-------+-------+---------------------------+
```

### Flavor

create
```
[vagrant@controller-10-64-33-64 ~]$ openstack flavor create --id 0 --vcpus 1 --ram 64 --disk 1 m1.nano
+----------------------------+---------+
| Field                      | Value   |
+----------------------------+---------+
| OS-FLV-DISABLED:disabled   | False   |
| OS-FLV-EXT-DATA:ephemeral  | 0       |
| disk                       | 1       |
| id                         | 0       |
| name                       | m1.nano |
| os-flavor-access:is_public | True    |
| properties                 |         |
| ram                        | 64      |
| rxtx_factor                | 1.0     |
| swap                       |         |
| vcpus                      | 1       |
+----------------------------+---------+
```

### Keypair

参考 https://docs.openstack.org/install-guide/launch-instance.html#generate-a-key-pair
```
[vagrant@controller-10-64-33-64 ~]$ openstack keypair list
```

### Security Group

Default
```
[vagrant@controller-10-64-33-64 ~]$ openstack security group list
+--------------------------------------+---------+------------------------+----------------------------------+
| ID                                   | Name    | Description            | Project                          |
+--------------------------------------+---------+------------------------+----------------------------------+
| a29a8bda-173d-4785-802f-fdb4e2a19664 | default | Default security group | a0be38aef8c74d4abca3e4e100ee7910 |
+--------------------------------------+---------+------------------------+----------------------------------+
```

```
[vagrant@controller-10-64-33-64 ~]$ openstack security group rule list --long default
+--------------------------------------+-------------+----------+------------+-----------+-----------+--------------------------------------+
| ID                                   | IP Protocol | IP Range | Port Range | Direction | Ethertype | Remote Security Group                |
+--------------------------------------+-------------+----------+------------+-----------+-----------+--------------------------------------+
| 40484746-2f04-4d5a-88f8-7f9a99370fab | None        | None     |            | ingress   | IPv4      | a29a8bda-173d-4785-802f-fdb4e2a19664 |
| 56745761-72d9-4df4-8951-9d0813cf3d89 | None        | None     |            | ingress   | IPv6      | a29a8bda-173d-4785-802f-fdb4e2a19664 |
| af5511c6-f631-48cc-8312-33fc931d84b1 | None        | None     |            | egress    | IPv4      | None                                 |
| b47af435-ca45-434f-aff7-7f47d7ac196a | None        | None     |            | egress    | IPv6      | None                                 |
+--------------------------------------+-------------+----------+------------+-----------+-----------+--------------------------------------+
```

ICMP
```
[vagrant@controller-10-64-33-64 ~]$ openstack security group rule create --proto icmp default
+-------------------+--------------------------------------+
| Field             | Value                                |
+-------------------+--------------------------------------+
| created_at        | 2017-10-28T10:04:41Z                 |
| description       |                                      |
| direction         | ingress                              |
| ether_type        | IPv4                                 |
| id                | b6475d58-130c-429f-ad44-55b56d65acfa |
| name              | None                                 |
| port_range_max    | None                                 |
| port_range_min    | None                                 |
| project_id        | a0be38aef8c74d4abca3e4e100ee7910     |
| protocol          | icmp                                 |
| remote_group_id   | None                                 |
| remote_ip_prefix  | 0.0.0.0/0                            |
| revision_number   | 0                                    |
| security_group_id | a29a8bda-173d-4785-802f-fdb4e2a19664 |
| updated_at        | 2017-10-28T10:04:41Z                 |
+-------------------+--------------------------------------+
```

SSH
```
[vagrant@controller-10-64-33-64 ~]$ openstack security group rule create --proto tcp --dst-port 22 default
+-------------------+--------------------------------------+
| Field             | Value                                |
+-------------------+--------------------------------------+
| created_at        | 2017-10-28T10:05:23Z                 |
| description       |                                      |
| direction         | ingress                              |
| ether_type        | IPv4                                 |
| id                | 5eeffc99-708a-401a-951b-4092d7a8506f |
| name              | None                                 |
| port_range_max    | 22                                   |
| port_range_min    | 22                                   |
| project_id        | a0be38aef8c74d4abca3e4e100ee7910     |
| protocol          | tcp                                  |
| remote_group_id   | None                                 |
| remote_ip_prefix  | 0.0.0.0/0                            |
| revision_number   | 0                                    |
| security_group_id | a29a8bda-173d-4785-802f-fdb4e2a19664 |
| updated_at        | 2017-10-28T10:05:23Z                 |
+-------------------+--------------------------------------+
```

### Image

Cirros
```
[vagrant@controller-10-64-33-64 ~]$ openstack image list
+--------------------------------------+--------+--------+
| ID                                   | Name   | Status |
+--------------------------------------+--------+--------+
| 40d37d6a-2dba-4323-b648-e806f3acb857 | cirros | active |
+--------------------------------------+--------+--------+
```

### DHCP Networking

混杂（promiscuous）模式?

controller node
```
[vagrant@controller-10-64-33-64 ~]$ sudo ls /etc/sysconfig/network-scripts/ifcfg-*
/etc/sysconfig/network-scripts/ifcfg-eth0  /etc/sysconfig/network-scripts/ifcfg-lo
/etc/sysconfig/network-scripts/ifcfg-eth2  /etc/sysconfig/network-scripts/ifcfg-Wired_connection_1
[vagrant@controller-10-64-33-64 ~]$ sudo cat /etc/sysconfig/network-scripts/ifcfg-eth2
DEVICE=eth2
TYPE=Ethernet
ONBOOT=yes
BOOTPROTO=none
[vagrant@controller-10-64-33-64 ~]$ sudo nmcli c load /etc/sysconfig/network-scripts/ifcfg-eth2
[vagrant@controller-10-64-33-64 ~]$ sudo nmcli c reload
[vagrant@controller-10-64-33-64 ~]$ sudo nmcli c show
NAME                UUID                                  TYPE            DEVICE 
System eth0         5fb06bd0-0bb0-7ffb-45f1-d6edd65f3e03  802-3-ethernet  eth0   
System eth2         3a73717e-65ab-93e8-b518-24f5af32dc0d  802-3-ethernet  eth2   
Wired connection 1  3c55af49-6222-3b6b-b91d-eb1b82b6005e  802-3-ethernet  eth1   
[vagrant@controller-10-64-33-64 ~]$ sudo nmcli d status
DEVICE  TYPE      STATE      CONNECTION         
eth0    ethernet  connected  System eth0        
eth1    ethernet  connected  Wired connection 1 
eth2    ethernet  connected  System eth2        
lo      loopback  unmanaged  --                 
```

```
[vagrant@controller-10-64-33-64 ~]$ ip a show eth2
4: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 08:00:27:f6:b8:b3 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::a00:27ff:fef6:b8b3/64 scope link 
       valid_lft forever preferred_lft forever
```

compute node
```
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli d set eth2 managed no
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli d status
DEVICE  TYPE      STATE      CONNECTION         
eth0    ethernet  connected  System eth0        
eth1    ethernet  connected  Wired connection 1 
eth2    ethernet  unmanaged  --                 
lo      loopback  unmanaged  --                 
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli c show
NAME                UUID                                  TYPE            DEVICE 
System eth0         5fb06bd0-0bb0-7ffb-45f1-d6edd65f3e03  802-3-ethernet  eth0   
Wired connection 1  0ab4e0ff-f77e-3886-9ba9-1838e4120850  802-3-ethernet  eth1   
System eth2         3a73717e-65ab-93e8-b518-24f5af32dc0d  802-3-ethernet  --     
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli d set eth2 autoconnect yes
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli d 
DEVICE  TYPE      STATE      CONNECTION         
eth0    ethernet  connected  System eth0        
eth1    ethernet  connected  Wired connection 1 
eth2    ethernet  unmanaged  --                 
lo      loopback  unmanaged  --                 
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli c
NAME                UUID                                  TYPE            DEVICE 
System eth0         5fb06bd0-0bb0-7ffb-45f1-d6edd65f3e03  802-3-ethernet  eth0   
Wired connection 1  0ab4e0ff-f77e-3886-9ba9-1838e4120850  802-3-ethernet  eth1   
System eth2         3a73717e-65ab-93e8-b518-24f5af32dc0d  802-3-ethernet  --     
```

### Provider Network

Create (参考 https://docs.openstack.org/install-guide/launch-instance-networks-provider.html)
```
[vagrant@controller-10-64-33-64 ~]$ sudo cat /etc/neutron/plugin.ini | egrep '^[^#]'
[DEFAULT]
[ml2]
type_drivers=flat,vlan
tenant_network_types=
mechanism_drivers=linuxbridge
extension_drivers=port_security
[ml2_type_flat]
flat_networks=provider
[ml2_type_geneve]
[ml2_type_gre]
[ml2_type_vlan]
[ml2_type_vxlan]
[securitygroup]
enable_ipset=true
```

```
[vagrant@controller-10-64-33-64 ~]$ sudo arping -I eth2 -c3 192.168.1.1
ARPING 192.168.1.1 from 10.0.2.15 eth2
Sent 3 probes (3 broadcast(s))
Received 0 response(s)
[vagrant@controller-10-64-33-64 ~]$ sudo arping -I eth2 -c3 192.168.1.100
ARPING 192.168.1.100 from 10.0.2.15 eth2
Sent 3 probes (3 broadcast(s))
Received 0 response(s)
```

```
[vagrant@controller-10-64-33-64 ~]$ openstack network create --share --external --provider-physical-network provider --provider-network-type flat provider1
+---------------------------+--------------------------------------+
| Field                     | Value                                |
+---------------------------+--------------------------------------+
| admin_state_up            | UP                                   |
| availability_zone_hints   |                                      |
| availability_zones        |                                      |
| created_at                | 2017-10-29T00:38:09Z                 |
| description               |                                      |
| dns_domain                | None                                 |
| id                        | 59fe11c2-2131-4a9b-b708-e855c3c6ac25 |
| ipv4_address_scope        | None                                 |
| ipv6_address_scope        | None                                 |
| is_default                | None                                 |
| is_vlan_transparent       | None                                 |
| mtu                       | 1500                                 |
| name                      | provider1                            |
| port_security_enabled     | True                                 |
| project_id                | a0be38aef8c74d4abca3e4e100ee7910     |
| provider:network_type     | flat                                 |
| provider:physical_network | provider                             |
| provider:segmentation_id  | None                                 |
| qos_policy_id             | None                                 |
| revision_number           | 3                                    |
| router:external           | External                             |
| segments                  | None                                 |
| shared                    | True                                 |
| status                    | ACTIVE                               |
| subnets                   |                                      |
| tags                      |                                      |
| updated_at                | 2017-10-29T00:38:09Z                 |
+---------------------------+--------------------------------------+
```

```
[vagrant@controller-10-64-33-64 ~]$ openstack subnet create --network provider1 --allocation-pool start=192.168.31.70,end=192.168.31.79 --dns-nameserver 8.8.4.4 --gateway 192.168.31.1 --subnet-range 192.168.31.64/27 provider1subnet
+-------------------------+--------------------------------------+
| Field                   | Value                                |
+-------------------------+--------------------------------------+
| allocation_pools        | 192.168.31.70-192.168.31.79          |
| cidr                    | 192.168.31.64/27                     |
| created_at              | 2017-10-29T00:52:57Z                 |
| description             |                                      |
| dns_nameservers         | 8.8.4.4                              |
| enable_dhcp             | True                                 |
| gateway_ip              | 192.168.31.1                         |
| host_routes             |                                      |
| id                      | 90315432-195b-4334-8094-1f0228093c16 |
| ip_version              | 4                                    |
| ipv6_address_mode       | None                                 |
| ipv6_ra_mode            | None                                 |
| name                    | provider1subnet                      |
| network_id              | 59fe11c2-2131-4a9b-b708-e855c3c6ac25 |
| project_id              | a0be38aef8c74d4abca3e4e100ee7910     |
| revision_number         | 0                                    |
| segment_id              | None                                 |
| service_types           |                                      |
| subnetpool_id           | None                                 |
| tags                    |                                      |
| updated_at              | 2017-10-29T00:52:57Z                 |
| use_default_subnet_pool | None                                 |
+-------------------------+--------------------------------------+

```

```
[vagrant@controller-10-64-33-64 ~]$ ps -ef | grep dns
nobody   26446     1  0 00:52 ?        00:00:00 dnsmasq --no-hosts --no-resolv --strict-order --except-interface=lo --pid-file=/var/lib/neutron/dhcp/59fe11c2-2131-4a9b-b708-e855c3c6ac25/pid --dhcp-hostsfile=/var/lib/neutron/dhcp/59fe11c2-2131-4a9b-b708-e855c3c6ac25/host --addn-hosts=/var/lib/neutron/dhcp/59fe11c2-2131-4a9b-b708-e855c3c6ac25/addn_hosts --dhcp-optsfile=/var/lib/neutron/dhcp/59fe11c2-2131-4a9b-b708-e855c3c6ac25/opts --dhcp-leasefile=/var/lib/neutron/dhcp/59fe11c2-2131-4a9b-b708-e855c3c6ac25/leases --dhcp-match=set:ipxe,175 --bind-interfaces --interface=ns-8e7ff5b5-3f --dhcp-range=set:tag0,192.168.31.64,static,255.255.255.224,86400s --dhcp-option-force=option:mtu,1500 --dhcp-lease-max=32 --conf-file= --domain=openstacklocal
vagrant  27395 24923  0 01:34 pts/0    00:00:00 grep --color=auto dns
```

```
[vagrant@controller-10-64-33-64 ~]$ brctl show
bridge name	bridge id		STP enabled	interfaces
brq59fe11c2-21		8000.080027f6b8b3	no		eth2
							tap8e7ff5b5-3f
```

```
[vagrant@controller-10-64-33-64 ~]$ ip a show brq59fe11c2-21
13: brq59fe11c2-21: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP qlen 1000
    link/ether 08:00:27:f6:b8:b3 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::a064:c9ff:fedf:25b4/64 scope link 
       valid_lft forever preferred_lft forever
```

```
[vagrant@controller-10-64-33-64 ~]$ ip a show tap8e7ff5b5-3f
12: tap8e7ff5b5-3f@if2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master brq59fe11c2-21 state UP qlen 1000
    link/ether ae:26:45:86:d8:78 brd ff:ff:ff:ff:ff:ff link-netnsid 0
```

```
[vagrant@controller-10-64-33-64 ~]$ ip netns
qdhcp-59fe11c2-2131-4a9b-b708-e855c3c6ac25 (id: 0)
```

```
[vagrant@controller-10-64-33-64 ~]$ sudo ip netns exec qdhcp-59fe11c2-2131-4a9b-b708-e855c3c6ac25 ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: ns-8e7ff5b5-3f@if12: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP qlen 1000
    link/ether fa:16:3e:19:f7:3b brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 169.254.169.254/16 brd 169.254.255.255 scope global ns-8e7ff5b5-3f
       valid_lft forever preferred_lft forever
    inet 192.168.31.70/27 brd 192.168.31.95 scope global ns-8e7ff5b5-3f
       valid_lft forever preferred_lft forever
    inet6 fe80::f816:3eff:fe19:f73b/64 scope link 
       valid_lft forever preferred_lft forever
```

```
[vagrant@controller-10-64-33-64 ~]$ sudo ip netns exec qdhcp-59fe11c2-2131-4a9b-b708-e855c3c6ac25 ip r
default via 192.168.31.1 dev ns-8e7ff5b5-3f 
169.254.0.0/16 dev ns-8e7ff5b5-3f proto kernel scope link src 169.254.169.254 
192.168.31.1 dev ns-8e7ff5b5-3f scope link 
192.168.31.64/27 dev ns-8e7ff5b5-3f proto kernel scope link src 192.168.31.70 
```

由于没有路由
```
[vagrant@controller-10-64-33-64 ~]$ ping -c3 192.168.31.70
PING 192.168.31.70 (192.168.31.70) 56(84) bytes of data.

--- 192.168.31.70 ping statistics ---
3 packets transmitted, 0 received, 100% packet loss, time 2000ms
```

为bridge配置ip地址作为provider网关
```
[vagrant@controller-10-64-33-64 ~]$ sudo ip a add 192.168.31.1/24 dev brq59fe11c2-21
```

```
[vagrant@controller-10-64-33-64 ~]$ ip r
default via 10.0.2.2 dev eth0 proto static metric 100 
10.0.2.0/24 dev eth0 proto kernel scope link src 10.0.2.15 metric 100 
10.64.33.0/24 dev eth1 proto kernel scope link src 10.64.33.64 metric 100 
192.168.31.0/24 dev brq59fe11c2-21 proto kernel scope link src 192.168.31.1 
```

```
[vagrant@controller-10-64-33-64 ~]$ arping -I brq59fe11c2-21 -c3 192.168.31.70
ARPING 192.168.31.70 from 192.168.31.1 brq59fe11c2-21
Unicast reply from 192.168.31.70 [FA:16:3E:19:F7:3B]  0.719ms
Unicast reply from 192.168.31.70 [FA:16:3E:19:F7:3B]  0.546ms
Unicast reply from 192.168.31.70 [FA:16:3E:19:F7:3B]  0.570ms
Sent 3 probes (1 broadcast(s))
Received 3 response(s)
```

```
[vagrant@controller-10-64-33-64 ~]$ arp -n
Address                  HWtype  HWaddress           Flags Mask            Iface
10.64.33.1               ether   0a:00:27:00:00:17   C                     eth1
169.254.169.254          ether   fa:16:3e:19:f7:3b   C                     brq59fe11c2-21
10.0.2.2                 ether   52:54:00:12:35:02   C                     eth0
10.0.2.3                 ether   52:54:00:12:35:03   C                     eth0
192.168.31.70            ether   fa:16:3e:19:f7:3b   C                     brq59fe11c2-21
10.64.33.65              ether   08:00:27:b4:e4:88   C                     eth1
```

现在可到达netns了
```
[vagrant@controller-10-64-33-64 ~]$ ping -c3 192.168.31.70
PING 192.168.31.70 (192.168.31.70) 56(84) bytes of data.
64 bytes from 192.168.31.70: icmp_seq=1 ttl=64 time=0.058 ms
64 bytes from 192.168.31.70: icmp_seq=2 ttl=64 time=0.084 ms
64 bytes from 192.168.31.70: icmp_seq=3 ttl=64 time=0.088 ms

--- 192.168.31.70 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 1999ms
rtt min/avg/max/mdev = 0.058/0.076/0.088/0.016 ms
```

```
[vagrant@controller-10-64-33-64 ~]$ sudo ip netns exec qdhcp-59fe11c2-2131-4a9b-b708-e855c3c6ac25 arping -c3 192.168.31.1
ARPING 192.168.31.1 from 169.254.169.254 ns-8e7ff5b5-3f
Unicast reply from 192.168.31.1 [08:00:27:F6:B8:B3]  0.840ms
Unicast reply from 192.168.31.1 [08:00:27:F6:B8:B3]  0.569ms
Unicast reply from 192.168.31.1 [08:00:27:F6:B8:B3]  0.570ms
Sent 3 probes (1 broadcast(s))
Received 3 response(s)
```

```
[vagrant@controller-10-64-33-64 ~]$ sudo ip netns exec qdhcp-59fe11c2-2131-4a9b-b708-e855c3c6ac25 arp -n
Address                  HWtype  HWaddress           Flags Mask            Iface
192.168.31.1             ether   08:00:27:f6:b8:b3   C                     ns-8e7ff5b5-3f
```

等下在计算节点上也需要手动检查网络


### Launch

Networking
```
[vagrant@controller-10-64-33-64 ~]$ openstack network list
+--------------------------------------+-----------+--------------------------------------+
| ID                                   | Name      | Subnets                              |
+--------------------------------------+-----------+--------------------------------------+
| 59fe11c2-2131-4a9b-b708-e855c3c6ac25 | provider1 | 90315432-195b-4334-8094-1f0228093c16 |
+--------------------------------------+-----------+--------------------------------------+
```

```
[vagrant@controller-10-64-33-64 ~]$ openstack subnet list
+--------------------------------------+-----------------+--------------------------------------+------------------+
| ID                                   | Name            | Network                              | Subnet           |
+--------------------------------------+-----------------+--------------------------------------+------------------+
| 90315432-195b-4334-8094-1f0228093c16 | provider1subnet | 59fe11c2-2131-4a9b-b708-e855c3c6ac25 | 192.168.31.64/27 |
+--------------------------------------+-----------------+--------------------------------------+------------------+
```

参考 https://docs.openstack.org/install-guide/launch-instance-provider.html
```
[vagrant@controller-10-64-33-64 ~]$ openstack server create --flavor m1.nano --image cirros --nic net-id=provider1 --security-group default cirros1
+-------------------------------------+-----------------------------------------------+
| Field                               | Value                                         |
+-------------------------------------+-----------------------------------------------+
| OS-DCF:diskConfig                   | MANUAL                                        |
| OS-EXT-AZ:availability_zone         |                                               |
| OS-EXT-SRV-ATTR:host                | None                                          |
| OS-EXT-SRV-ATTR:hypervisor_hostname | None                                          |
| OS-EXT-SRV-ATTR:instance_name       |                                               |
| OS-EXT-STS:power_state              | NOSTATE                                       |
| OS-EXT-STS:task_state               | scheduling                                    |
| OS-EXT-STS:vm_state                 | building                                      |
| OS-SRV-USG:launched_at              | None                                          |
| OS-SRV-USG:terminated_at            | None                                          |
| accessIPv4                          |                                               |
| accessIPv6                          |                                               |
| addresses                           |                                               |
| adminPass                           | ynCS89i6DzkE                                  |
| config_drive                        |                                               |
| created                             | 2017-10-29T01:42:40Z                          |
| flavor                              | m1.nano (0)                                   |
| hostId                              |                                               |
| id                                  | 90e44b5b-31d1-49f6-84ae-db0d85fdae7a          |
| image                               | cirros (40d37d6a-2dba-4323-b648-e806f3acb857) |
| key_name                            | None                                          |
| name                                | cirros1                                       |
| progress                            | 0                                             |
| project_id                          | a0be38aef8c74d4abca3e4e100ee7910              |
| properties                          |                                               |
| security_groups                     | name='a29a8bda-173d-4785-802f-fdb4e2a19664'   |
| status                              | BUILD                                         |
| updated                             | 2017-10-29T01:42:40Z                          |
| user_id                             | 44e6ee1df8ae436986d2d50f7b358aa0              |
| volumes_attached                    |                                               |
+-------------------------------------+-----------------------------------------------+
```

```
[vagrant@controller-10-64-33-64 ~]$ openstack server list
+--------------------------------------+---------+--------+-------------------------+--------+---------+
| ID                                   | Name    | Status | Networks                | Image  | Flavor  |
+--------------------------------------+---------+--------+-------------------------+--------+---------+
| 90e44b5b-31d1-49f6-84ae-db0d85fdae7a | cirros1 | ACTIVE | provider1=192.168.31.74 | cirros | m1.nano |
+--------------------------------------+---------+--------+-------------------------+--------+---------+
```

vnc
```
[vagrant@controller-10-64-33-64 ~]$ openstack console url show cirros1
+-------+----------------------------------------------------------------------------------+
| Field | Value                                                                            |
+-------+----------------------------------------------------------------------------------+
| type  | novnc                                                                            |
| url   | http://10.64.33.64:6080/vnc_auto.html?token=4dc627f5-8ef4-4b06-9214-0e56bc0ec93d |
+-------+----------------------------------------------------------------------------------+
```

## Compute

### VM networking

Environment
```
[vagrant@compute-10-64-33-65 ~]$ ls /var/lib/nova/instances/
90e44b5b-31d1-49f6-84ae-db0d85fdae7a  _base  compute_nodes  locks
```

```
[vagrant@compute-10-64-33-65 ~]$ brctl show
bridge name	bridge id		STP enabled	interfaces
brq59fe11c2-21		8000.0800275b61bc	no		eth2
							tap14d72434-b8
```

```
[vagrant@compute-10-64-33-65 ~]$ ip a show brq59fe11c2-21
7: brq59fe11c2-21: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP qlen 1000
    link/ether 08:00:27:5b:61:bc brd ff:ff:ff:ff:ff:ff
```

```
[vagrant@compute-10-64-33-65 ~]$ ip a show tap14d72434-b8
8: tap14d72434-b8: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast master brq59fe11c2-21 state UNKNOWN qlen 1000
    link/ether fe:16:3e:cd:96:47 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::fc16:3eff:fecd:9647/64 scope link 
       valid_lft forever preferred_lft forever

```

```
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli d
DEVICE          TYPE      STATE      CONNECTION         
brq59fe11c2-21  bridge    connected  brq59fe11c2-21     
eth0            ethernet  connected  System eth0        
eth1            ethernet  connected  Wired connection 1 
eth2            ethernet  connected  System eth2        
tap14d72434-b8  tun       connected  tap14d72434-b8     
lo              loopback  unmanaged  --                 
```

```
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli d set eth2 managed no
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli d set brq59fe11c2-21 managed no
[vagrant@compute-10-64-33-65 ~]$ sudo nmcli d set tap14d72434-b8 managed no
```

```
[vagrant@compute-10-64-33-65 ~]$ sudo ip a add 192.168.31.5/24 dev brq59fe11c2-21
```

```
[vagrant@compute-10-64-33-65 ~]$ arping -I brq59fe11c2-21 -c3 192.168.31.1
ARPING 192.168.31.1 from 192.168.31.5 brq59fe11c2-21
Unicast reply from 192.168.31.1 [08:00:27:F6:B8:B3]  2.200ms
Unicast reply from 192.168.31.1 [08:00:27:F6:B8:B3]  0.722ms
Unicast reply from 192.168.31.1 [08:00:27:F6:B8:B3]  1.104ms
Sent 3 probes (1 broadcast(s))
Received 3 response(s)
[vagrant@compute-10-64-33-65 ~]$ arping -I brq59fe11c2-21 -c3 192.168.31.70
ARPING 192.168.31.70 from 192.168.31.5 brq59fe11c2-21
Unicast reply from 192.168.31.70 [FA:16:3E:19:F7:3B]  2.275ms
Sent 3 probes (1 broadcast(s))
Received 1 response(s)
[vagrant@compute-10-64-33-65 ~]$ arp -n
Address                  HWtype  HWaddress           Flags Mask            Iface
10.0.2.3                 ether   52:54:00:12:35:03   C                     eth0
192.168.31.70            ether   fa:16:3e:19:f7:3b   C                     brq59fe11c2-21
10.0.2.2                 ether   52:54:00:12:35:02   C                     eth0
192.168.31.1             ether   08:00:27:f6:b8:b3   C                     brq59fe11c2-21
192.168.31.74                    (incomplete)                              brq59fe11c2-21
10.64.33.1               ether   0a:00:27:00:00:17   C                     eth1
10.64.33.64              ether   08:00:27:68:ae:d9   C                     eth1
```

### debug

virsh
```
[vagrant@compute-10-64-33-65 ~]$ sudo virsh list
 Id    Name                           State
----------------------------------------------------
 2     instance-00000002              running
[vagrant@compute-10-64-33-65 ~]$ sudo virsh domiflist 2
Interface  Type       Source     Model       MAC
-------------------------------------------------------
tap14d72434-b8 bridge     brq59fe11c2-21 virtio      fa:16:3e:cd:96:47

[vagrant@compute-10-64-33-65 ~]$ sudo virsh domifaddr 2
 Name       MAC address          Protocol     Address
-------------------------------------------------------------------------------

[vagrant@compute-10-64-33-65 ~]$ sudo virsh domstate 2
running

[vagrant@compute-10-64-33-65 ~]$ sudo virsh domstats 2
Domain: 'instance-00000002'
  state.state=1
  state.reason=5
  cpu.time=95448760595
  cpu.user=30790000000
  cpu.system=1600000000
  balloon.current=65536
  balloon.maximum=65536
  balloon.last-update=0
  balloon.rss=233920
  vcpu.current=1
  vcpu.maximum=1
  net.count=1
  net.0.name=tap14d72434-b8
  net.0.rx.bytes=2058
  net.0.rx.pkts=39
  net.0.rx.errs=0
  net.0.rx.drop=0
  net.0.tx.bytes=2928
  net.0.tx.pkts=18
  net.0.tx.errs=0
  net.0.tx.drop=0
  block.count=1
  block.0.name=vda
  block.0.path=/var/lib/nova/instances/90e44b5b-31d1-49f6-84ae-db0d85fdae7a/disk
  block.0.rd.reqs=1997
  block.0.rd.bytes=40820736
  block.0.rd.times=12570091638
  block.0.wr.reqs=200
  block.0.wr.bytes=555008
  block.0.wr.times=1859711706
  block.0.fl.reqs=33
  block.0.fl.times=23235572
  block.0.allocation=2097152
  block.0.capacity=1073741824
  block.0.physical=2101248

```

按Eenter登录
```
[vagrant@compute-10-64-33-65 ~]$ sudo virsh console 2
Connected to domain instance-00000002
Escape character is ^]


login as 'cirros' user. default password: 'cubswin:)'. use 'sudo' for root.
cirros login: cirros
Password: 
$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 16436 qdisc noqueue 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast qlen 1000
    link/ether fa:16:3e:cd:96:47 brd ff:ff:ff:ff:ff:ff
    inet6 fe80::f816:3eff:fecd:9647/64 scope link 
       valid_lft forever preferred_lft forever
```

```
$ sudo ip a add 192.168.31.74/27 dev eth0
```

```
$ sudo arping -c3 -I eth0 192.168.31.5
ARPING to 192.168.31.5 from 192.168.31.74 via eth0
Unicast reply from 192.168.31.5 [8:0:27:5b:61:bc] 1.065ms
Unicast reply from 192.168.31.5 [8:0:27:5b:61:bc] 0.378ms
Unicast reply from 192.168.31.5 [8:0:27:5b:61:bc] 1.476ms
Sent 3 probe(s) (1 broadcast(s))
Received 3 reply (0 request(s), 0 broadcast(s))
```

```
$ sudo ip r add default dev eth0
```

提示按ctrl-]退出
```
[vagrant@compute-10-64-33-65 ~]$ sudo arping -I brq59fe11c2-21 -c3 192.168.31.74
ARPING 192.168.31.74 from 192.168.31.5 brq59fe11c2-21
Unicast reply from 192.168.31.74 [FA:16:3E:CD:96:47]  0.939ms
Unicast reply from 192.168.31.74 [FA:16:3E:CD:96:47]  1.450ms
Unicast reply from 192.168.31.74 [FA:16:3E:CD:96:47]  0.763ms
Sent 3 probes (1 broadcast(s))
Received 3 response(s)
```

```
[vagrant@compute-10-64-33-65 ~]$ ssh cirros@192.168.31.74
The authenticity of host '192.168.31.74 (192.168.31.74)' can't be established.
RSA key fingerprint is SHA256:qd2kmEg0QvozmNOvQTSKKw1RYp70pNhKkGBun5V0vo0.
RSA key fingerprint is MD5:d5:7c:3d:15:7a:fc:22:e5:b8:04:21:a2:d1:1c:be:7d.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '192.168.31.74' (RSA) to the list of known hosts.
cirros@192.168.31.74's password: 
```