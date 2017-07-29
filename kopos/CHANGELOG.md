# API instruction

__OpenAPI__

Swagger 2.0

swagger-ui为https://localhost:10000/swagger-ui

swagger.json为https://localhost:10000/swagger.json

__Table of Contents__

1. [Batch create Virtual Machines](#Batch create Virtual Machines)
2. [Batch create simple Networks](#Batch create simple Networks)
3. [List flavors](#Flavor list)
3. [List images](#Image list)
3. [Get flavor](#Flavor get)
3. [Get image](#Image get)
4. [List networks](#Network list)
4. [List subnetes](#Subnet list)
4. [List VMs](#Virtual machine list)
4. [Reboot VMs](#Reboot virtual machines)
4. [Destroy VMs](#Destroy virtual machines)
5. [Spawn VMs](#Spawn Machines)
5. [Networking Architecture](#Networking Topology)

## V1

API path prefix  _/v1/..._


### 0.1.2-alpha3

2012-07-24

Fix bug  /v1/spawn   with correct json
```
  {
    vms: [
      {
        "flavor_name": "m1.small",
        "image_name": "cirros",
        "min_count": 2,
        "max_count": 4,
        "secgroups_info": [],
        "user_data": [],
        "network_name": "private",
        "floating_network_name": "public",
        "personality": [],
        "name_prefix": "awesome-VM"
      },
      {
        ...
      }
    ]
  }
```

### 0.1.2-alpha2

2012-07-12

### Openstack compute config

1）因获取libvirt的vncdisplay需要，对现用openstack的compute node的配置需显式配置（CHANGELOG）
如在10.121.198.4环境上
```
[root@compute01 ~]# head /etc/nova/nova.conf
[default]
my_ip=10.121.198.4
```

请devops务必在今后，或其它openstack环境中添加my_ip配置项，如10.100.151.149，配置my_ip＝10.100.151.149

检查配置
```
[vagrant@localhost kopos]$ nova hypervisor-show compute01
+---------------------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Property                  | Value                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
+---------------------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cpu_info_arch             | x86_64                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| cpu_info_features         | ["ssse3", "pge", "avx", "clflush", "sep", "syscall", "vme", "dtes64", "invpcid", "msr", "sse", "xsave", "vmx", "erms", "xtpr", "cmov", "tsc", "smep", "pbe", "est", "pat", "monitor", "smx", "lm", "abm", "nx", "fxsr", "tm", "sse4.1", "pae", "sse4.2", "pclmuldq", "acpi", "fma", "tsc-deadline", "popcnt", "mmx", "osxsave", "cx8", "mce", "mtrr", "rdtscp", "ht", "pse", "lahf_lm", "pdcm", "mca", "pdpe1gb", "apic", "fsgsbase", "f16c", "ds", "invtsc", "pni", "tm2", "avx2", "aes", "sse2", "ss", "bmi1", "bmi2", "pcid", "de", "fpu", "cx16", "pse36", "ds_cpl", "movbe", "rdrand", "x2apic"] |
| cpu_info_model            | SandyBridge                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| cpu_info_topology_cores   | 4                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| cpu_info_topology_sockets | 1                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| cpu_info_topology_threads | 2                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| cpu_info_vendor           | Intel                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| current_workload          | 0                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| disk_available_least      | -607                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| free_disk_gb              | -532                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| free_ram_mb               | 15055                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| host_ip                   | 10.121.198.4                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| hypervisor_hostname       | compute01                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| hypervisor_type           | QEMU                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| hypervisor_version        | 1005003                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| id                        | 1                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| local_gb                  | 872                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| local_gb_used             | 1404                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| memory_mb                 | 31951                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| memory_mb_used            | 16896                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| running_vms               | 11                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| service_disabled_reason   | -                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| service_host              | compute01                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| service_id                | 5                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| state                     | up                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| status                    | enabled                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| vcpus                     | 8                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| vcpus_used                | 11                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
+---------------------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
```

默认是127.0.0.1
```
[vagrant@localhost kopos]$ nova hypervisor-show compute01
+---------------------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Property                  | Value                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
+---------------------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| cpu_info_arch             | x86_64                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| cpu_info_features         | ["ssse3", "pge", "avx", "clflush", "sep", "syscall", "vme", "dtes64", "invpcid", "msr", "sse", "xsave", "vmx", "erms", "xtpr", "cmov", "tsc", "smep", "pbe", "est", "pat", "monitor", "smx", "lm", "abm", "nx", "fxsr", "tm", "sse4.1", "pae", "sse4.2", "pclmuldq", "acpi", "fma", "tsc-deadline", "popcnt", "mmx", "osxsave", "cx8", "mce", "mtrr", "rdtscp", "ht", "pse", "lahf_lm", "pdcm", "mca", "pdpe1gb", "apic", "fsgsbase", "f16c", "ds", "invtsc", "pni", "tm2", "avx2", "aes", "sse2", "ss", "bmi1", "bmi2", "pcid", "de", "fpu", "cx16", "pse36", "ds_cpl", "movbe", "rdrand", "x2apic"] |
| cpu_info_model            | SandyBridge                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| cpu_info_topology_cores   | 4                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| cpu_info_topology_sockets | 1                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| cpu_info_topology_threads | 2                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| cpu_info_vendor           | Intel                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| current_workload          | 0                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| disk_available_least      | -36                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| free_disk_gb              | 31                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| free_ram_mb               | 207                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| host_ip                   | 127.0.0.1                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| hypervisor_hostname       | compute01                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| hypervisor_type           | QEMU                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| hypervisor_version        | 1005003                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| id                        | 1                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| local_gb                  | 872                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| local_gb_used             | 841                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| memory_mb                 | 31951                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| memory_mb_used            | 31744                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| running_vms               | 6                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| service_disabled_reason   | -                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| service_host              | compute01                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |
| service_id                | 5                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| state                     | up                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| status                    | enabled                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| vcpus                     | 8                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| vcpus_used                | 16                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
+---------------------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
```

2017-07-08

#### Fetch VNC displaoy port

_GET /v1/libvirt-domains0x3f/{server_id}_ （新增），获取vm的VNCDisplay（virsh vncdisplay $instance_name）

Response
```
{
    "server_id": "012345678-01234-01234-01234-0123456789ab",
    "state_code": 返回的正数值,
    "string state_message": "返回的字符串值",
    "domain_info": {
        "display": "10.100.151.234:1"
    }
}
```

### 0.1.2-alpha1

2017-07-06

#### Networking Topology

_POST /v1/topology_ （新增），显示路由连接下的内外网关系，Request/Response是基于/v1/landscape的JSON数组

Response
```
{
  "vnets": [
    {
      "name": "int-stage-0",
      "subnets": [
        {
          "name": "int-192-168-128-0-slash-24",
          "cidr": "192.168.128.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "int-stage-1",
      "subnets": [
        {
          "name": "int-192-168-129-0-slash-24",
          "cidr": "192.168.129.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "int-stage-2",
      "subnets": [
        {
          "name": "int-192-168-130-0-slash-24",
          "cidr": "192.168.130.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "int-stage-3",
      "subnets": [
        {
          "name": "int-192-168-131-0-slash-24",
          "cidr": "192.168.131.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    }
  ],
  "vrouter": {
    "name": "hack",
    "routes": [],
    "admint_state_up": true
  },
  "secgroup": {},
  "ifaces_info": [],
  "gateways_info": [
    {
      "network_name": "public",
      "router_name": "hack"
    }
  ],
  "ports": [
    {
      "id": "port_id",
      "name": "port_name"
    }
  ],
  "interfaces_info": []
}
```

### Spawn Machines

_POST /v1/spawn_ （新增），创建不同规格的虚拟机，Request/Response是基于/v1/boot的JOSN数组
```
{ [
  {
    "flavor_name": "large",
    "image_name": "cirros",
    "min_count": 1,
    "max_count": 1,
    "secgroups_info": [],
    "user_data": [],
    "network_name": "private",
    "floating_network_name": "public-admin1",
    "personality": [],
    "name_prefix": "............"
  },
  {
    "flavor_name": "large",
    "image_name": "target86a-phpcmsV9_blindSQL-Ubuntu-cl2",
    "min_count": 1,
    "max_count": 1,
    "secgroups_info": [],
    "user_data": [],
    "network_name": "private",
    "floating_network_name": "",
    "personality": [],
    "name_prefix": "............"
  }
] }
```

#### Enhanced

_POST /v1/boot_ （增强）如果不提供floating_network_name，则不创建floatingip，如果有空余的floatingip，则不新建

#### Fixed bug

_GET /v1/servers_ （修复）显示虚拟机的ip地址
```
{
  ...
  "Addresses": 
    {
      "租户网络名": [
        {
          "addr": "192.168.1.100",
          "version": 4,
          "mac_addr": "23:45:67:89:ab:cd",
          "assigned_type": "dhcp"
        }
      ],
      "Floating网络名": [
        ...
      ]
    }
  ...
}
```

### 0.1.2-alpha0

2017-06-29

#### Virtual machine list

_GET /v1/servers_

#### Destroy virtual machines

_POST /v1/recycle_

#### Reboot virtual machines

_POST /v1/reboot_

#### Network list

_GET /v1/networks_

#### Subnet list

_GET /v1/subnets_

#### Fixed bug in [boot VMs](#Batch create Virtual Machines), waiting newly created port to state UP
```
[vagrant@localhost go-to-openstack-bootcamp]$ nova list
+--------------------------------------+----------------+--------+------------+-------------+-----------------------------------------------+
| ID                                   | Name           | Status | Task State | Power State | Networks                                      |
+--------------------------------------+----------------+--------+------------+-------------+-----------------------------------------------+
| 1cfca805-d9db-4bfa-9952-c0b5144b81d8 | z1t138beoclp-0 | ACTIVE | -          | Running     | private-admin1=192.168.128.103, 10.100.151.36 |
| 44093845-3d44-4d63-99ce-89948cd12961 | z1t138beoclp-1 | ACTIVE | -          | Running     | private-admin1=192.168.128.104, 10.100.151.37 |
| 58939fb2-8385-4ae1-8ebe-9c0c15b2e283 | z1t138beoclp-2 | ACTIVE | -          | Running     | private-admin1=192.168.128.105, 10.100.151.38 |
| 76778ba0-58f9-4371-a10d-21d72626eb19 | z1t138beoclp-3 | ACTIVE | -          | Running     | private-admin1=192.168.128.106, 10.100.151.39 |
+--------------------------------------+----------------+--------+------------+-------------+-----------------------------------------------+
fanhonglingdeMacBook-Pro:~ fanhongling$ ssh cirros@10.100.151.37
The authenticity of host '10.100.151.37 (10.100.151.37)' can't be established.
RSA key fingerprint is f2:34:db:f5:59:5a:d7:26:b0:46:a8:b7:3d:18:50:43.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '10.100.151.37' (RSA) to the list of known hosts.
cirros@10.100.151.37's password: 
$ ip a show eth0
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast qlen 1000
    link/ether fa:16:3e:95:d2:c1 brd ff:ff:ff:ff:ff:ff
    inet 192.168.128.104/24 brd 192.168.128.255 scope global eth0
    inet6 fe80::f816:3eff:fe95:d2c1/64 scope link 
       valid_lft forever preferred_lft forever
$ exit
Connection to 10.100.151.37 closed.
[vagrant@localhost go-to-openstack-bootcamp]$ neutron router-list
+--------------------------------------+---------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------+-------+
| id                                   | name          | external_gateway_info                                                                                                                                                                     | distributed | ha    |
+--------------------------------------+---------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------+-------+
| 9601f30a-fd2f-4d1c-a620-e4e0dba834ee | router-admin1 | {"network_id": "a7276117-d06e-4d9a-b3dd-4f31e0d4d861", "enable_snat": true, "external_fixed_ips": [{"subnet_id": "c31a05b1-d146-4a6e-b456-bc5113f4567e", "ip_address": "10.100.151.33"}]} | False       | False |
+--------------------------------------+---------------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------+-------+
[root@network ~]# ip netns
qdhcp-05fc56e3-0aac-4812-8c9e-56a54c20fdc6
qrouter-9601f30a-fd2f-4d1c-a620-e4e0dba834ee
[root@network ~]# ip netns exec qrouter-9601f30a-fd2f-4d1c-a620-e4e0dba834ee ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN 
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
30: qg-4a74f6d9-40: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN 
    link/ether fa:16:3e:70:a4:09 brd ff:ff:ff:ff:ff:ff
    inet 10.100.151.33/24 brd 10.100.151.255 scope global qg-4a74f6d9-40
       valid_lft forever preferred_lft forever
    inet 10.100.151.36/32 brd 10.100.151.36 scope global qg-4a74f6d9-40
       valid_lft forever preferred_lft forever
    inet 10.100.151.38/32 brd 10.100.151.38 scope global qg-4a74f6d9-40
       valid_lft forever preferred_lft forever
    inet 10.100.151.39/32 brd 10.100.151.39 scope global qg-4a74f6d9-40
       valid_lft forever preferred_lft forever
    inet 10.100.151.37/32 brd 10.100.151.37 scope global qg-4a74f6d9-40
       valid_lft forever preferred_lft forever
    inet6 fe80::f816:3eff:fe70:a409/64 scope link 
       valid_lft forever preferred_lft forever
31: qr-5901c3bc-80: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN 
    link/ether fa:16:3e:0d:cb:57 brd ff:ff:ff:ff:ff:ff
    inet 192.168.128.1/24 brd 192.168.128.255 scope global qr-5901c3bc-80
       valid_lft forever preferred_lft forever
    inet6 fe80::f816:3eff:fe0d:cb57/64 scope link 
       valid_lft forever preferred_lft forever
[root@network ~]# ip netns exec qrouter-9601f30a-fd2f-4d1c-a620-e4e0dba834ee ping -c3 192.168.128.1
PING 192.168.128.1 (192.168.128.1) 56(84) bytes of data.
64 bytes from 192.168.128.1: icmp_seq=1 ttl=64 time=0.067 ms
64 bytes from 192.168.128.1: icmp_seq=2 ttl=64 time=0.027 ms
64 bytes from 192.168.128.1: icmp_seq=3 ttl=64 time=0.044 ms

--- 192.168.128.1 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 1999ms
rtt min/avg/max/mdev = 0.027/0.046/0.067/0.016 ms
[root@network ~]# ip netns exec qrouter-9601f30a-fd2f-4d1c-a620-e4e0dba834ee ping -c3 192.168.128.103
PING 192.168.128.103 (192.168.128.103) 56(84) bytes of data.
64 bytes from 192.168.128.103: icmp_seq=1 ttl=64 time=1.85 ms
64 bytes from 192.168.128.103: icmp_seq=2 ttl=64 time=0.555 ms
64 bytes from 192.168.128.103: icmp_seq=3 ttl=64 time=0.443 ms

--- 192.168.128.103 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2002ms
rtt min/avg/max/mdev = 0.443/0.951/1.857/0.642 ms
[root@network ~]# ip netns exec qrouter-9601f30a-fd2f-4d1c-a620-e4e0dba834ee ping -c3 192.168.128.104
PING 192.168.128.104 (192.168.128.104) 56(84) bytes of data.
64 bytes from 192.168.128.104: icmp_seq=1 ttl=64 time=1.68 ms
64 bytes from 192.168.128.104: icmp_seq=2 ttl=64 time=0.682 ms
64 bytes from 192.168.128.104: icmp_seq=3 ttl=64 time=0.526 ms

--- 192.168.128.104 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2002ms
rtt min/avg/max/mdev = 0.526/0.963/1.682/0.512 ms
[root@network ~]# ip netns exec qrouter-9601f30a-fd2f-4d1c-a620-e4e0dba834ee ping -c3 192.168.128.105
PING 192.168.128.105 (192.168.128.105) 56(84) bytes of data.
64 bytes from 192.168.128.105: icmp_seq=1 ttl=64 time=1.89 ms
64 bytes from 192.168.128.105: icmp_seq=2 ttl=64 time=0.528 ms
64 bytes from 192.168.128.105: icmp_seq=3 ttl=64 time=0.472 ms

--- 192.168.128.105 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2002ms
rtt min/avg/max/mdev = 0.472/0.963/1.891/0.657 ms
[root@network ~]# ip netns exec qrouter-9601f30a-fd2f-4d1c-a620-e4e0dba834ee ping -c3 10.100.151.36
PING 10.100.151.36 (10.100.151.36) 56(84) bytes of data.
64 bytes from 10.100.151.36: icmp_seq=1 ttl=64 time=0.863 ms
64 bytes from 10.100.151.36: icmp_seq=2 ttl=64 time=0.539 ms
64 bytes from 10.100.151.36: icmp_seq=3 ttl=64 time=0.451 ms

--- 10.100.151.36 ping statistics ---
3 packets transmitted, 3 received, 0% packet loss, time 2000ms
rtt min/avg/max/mdev = 0.451/0.617/0.863/0.179 ms
[root@network ~]# ip netns exec qrouter-9601f30a-fd2f-4d1c-a620-e4e0dba834ee ssh cirros@192.168.128.103
The authenticity of host '192.168.128.103 (192.168.128.103)' can't be established.
RSA key fingerprint is 7b:9c:79:d1:00:5c:fc:ce:6a:f1:36:a2:cd:df:ca:66.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '192.168.128.103' (RSA) to the list of known hosts.
cirros@192.168.128.103's password: 
$ ip a show eth0
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast qlen 1000
    link/ether fa:16:3e:d9:b1:5e brd ff:ff:ff:ff:ff:ff
    inet 192.168.128.103/24 brd 192.168.128.255 scope global eth0
    inet6 fe80::f816:3eff:fed9:b15e/64 scope link 
       valid_lft forever preferred_lft forever
$ exit
Connection to 192.168.128.103 closed.
[root@network ~]# ip netns exec qrouter-9601f30a-fd2f-4d1c-a620-e4e0dba834ee ssh cirros@10.100.151.36
The authenticity of host '10.100.151.36 (10.100.151.36)' can't be established.
RSA key fingerprint is 7b:9c:79:d1:00:5c:fc:ce:6a:f1:36:a2:cd:df:ca:66.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '10.100.151.36' (RSA) to the list of known hosts.
cirros@10.100.151.36's password: 
$ ip a show eth0
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast qlen 1000
    link/ether fa:16:3e:d9:b1:5e brd ff:ff:ff:ff:ff:ff
    inet 192.168.128.103/24 brd 192.168.128.255 scope global eth0
    inet6 fe80::f816:3eff:fed9:b15e/64 scope link 
       valid_lft forever preferred_lft forever
$ exit
Connection to 10.100.151.36 closed.
```

#### Image get

_GET /v1/images/0x3f/{id}_

_GET /v1/images/0x3fname/{name}_

Command
```
fanhonglingdeMacBook-Pro:~ fanhongling$ curl http://127.0.0.1:10001/v1/flavors/0x3fname/large
{"id":"3b36b050-4b98-4902-810c-79df2f0c80fe","disk":200,"ram":2048,"name":"large","rxtx_factor":1,"vcpus":1}
fanhonglingdeMacBook-Pro:~ fanhongling$ curl http://127.0.0.1:10001/v1/flavors/0x3f/3b36b050-4b98-4902-810c-79df2f0c80fe
{"id":"3b36b050-4b98-4902-810c-79df2f0c80fe","disk":200,"ram":2048,"name":"large","rxtx_factor":1,"vcpus":1}
```

#### Flavor get

_GET /v1/flavors/0x3f/{id}_

_GET /v1/flavors/0x3fname/{name}_

Command
```
fanhonglingdeMacBook-Pro:~ fanhongling$ curl http://127.0.0.1:10001/v1/images/0x3fname/cirros
{"id":"30c1093a-d98a-4a41-8d39-cd5ce8e0d45a","name":"cirros","status":"active","container_format":"bare","disk_format":"qcow2","owner":"1ba84201548b469aae6ec1303abbb482","visibility":"private","checksum":"d972013792949d0d3ba628fbe8685bce","size":"13147648","create_at":"2017-06-28T16:22:03Z","updated_at":"2017-06-28T16:22:03Z","file":"/v2/images/30c1093a-d98a-4a41-8d39-cd5ce8e0d45a/file","schema":"/v2/schemas/image"}
fanhonglingdeMacBook-Pro:~ fanhongling$ curl http://127.0.0.1:10001/v1/images/0x3f/30c1093a-d98a-4a41-8d39-cd5ce8e0d45a
{"id":"30c1093a-d98a-4a41-8d39-cd5ce8e0d45a","name":"cirros","status":"active","container_format":"bare","disk_format":"qcow2","owner":"1ba84201548b469aae6ec1303abbb482","visibility":"private","checksum":"d972013792949d0d3ba628fbe8685bce","size":"13147648","create_at":"2017-06-28T16:22:03Z","updated_at":"2017-06-28T16:22:03Z","file":"/v2/images/30c1093a-d98a-4a41-8d39-cd5ce8e0d45a/file","schema":"/v2/schemas/image"}
```

### 0.1.1-beta0

2017-06-25

#### Image list

_GET /v1/images_

Command
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ curl -iL http://127.0.0.1:10001/v1/images
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 25 Jun 2017 08:32:42 GMT
Transfer-Encoding: chunked

{"images":[{"id":"4ac53f68-3ac2-49df-8d69-afcea235f8b9","name":"target_Struts2032-Struts2032-Struts2S2-032-disk1-cl1","status":"ACTIVE","create_at":"2017-06-22T21:05:12Z","updated_at":"2017-06-22T21:05:23Z"},{"id":"c611c095-17a0-4236-9310-438aa9a133e3","name":"target_Weblogic-Weblogic-Weblogic-disk1-cl1","status":"ACTIVE","create_at":"2017-06-22T20:57:24Z","updated_at":"2017-06-22T20:57:56Z"}]}
```
#### Flavor list

_GET /v1/flavors_

Command
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ curl -iL http://127.0.0.1:10001/v1/flavors
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 25 Jun 2017 08:33:07 GMT
Content-Length: 503

{"flavors":[{"id":"1","disk":1,"ram":512,"name":"m1.tiny","rxtx_factor":1,"vcpus":1},{"id":"2","disk":20,"ram":2048,"name":"m1.small","rxtx_factor":1,"vcpus":1},{"id":"3","disk":40,"ram":4096,"name":"m1.medium","rxtx_factor":1,"vcpus":2},{"id":"4","disk":80,"ram":8192,"name":"m1.large","rxtx_factor":1,"vcpus":4},{"id":"5","disk":160,"ram":16384,"name":"m1.xlarge","rxtx_factor":1,"vcpus":8},{"id":"bf289de9-bd91-4d12-9d86-a4e8889f412f","disk":200,"ram":2048,"name":"large","rxtx_factor":1,"vcpus":1}]}
```

### 0.1.1-alpha0

2017-06-25

#### Batch create simple Networks

_POST /v1/landscape_

Request
```
{
  "vnets": [
    {
      "name": "int-stage-0",
      "subnets": [
        {
          "name": "int-192-168-128-0-slash-24",
          "cidr": "192.168.128.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "int-stage-1",
      "subnets": [
        {
          "name": "int-192-168-129-0-slash-24",
          "cidr": "192.168.129.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "int-stage-2",
      "subnets": [
        {
          "name": "int-192-168-130-0-slash-24",
          "cidr": "192.168.130.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "int-stage-3",
      "subnets": [
        {
          "name": "int-192-168-131-0-slash-24",
          "cidr": "192.168.131.0/24",
          "dns_name_servers": [],
          "allocation_pools": [],
          "host_routes": [],
          "enable_dhcp": true
        }
      ]
    },
    {
      "name": "public",
      "admin_state_up": true,
      "subnets": [
        {
          "name": "10.100.151.0/24",
          "cidr": "10.100.151.0/24",
          "gateway_ip": "10.100.151.1",
          "dns_name_servers": [],
          "allocation_pools": [
            {
              "start": "10.100.151.50",
              "end": "10.100.151.240"
            }
          ],
          "host_routes": [],
          "enable_dhcp": false
        }
      ],
      "shared": true
    }
  ],
  "vrouter": {
    "name": "hack",
    "routes": [],
    "admint_state_up": true
  },
  "secgroup": {
    "name": "hack",
    "security_group_rules": [
      {
        "direction": "ingress",
        "protocol": "tcp"
      },
      {
        "direction": "ingress",
        "protocol": "udp"
      },
      {
        "direction": "ingress",
        "protocol": "icmp"
      }
    ]
  },
  "ifaces_info": [
    {
      "router_name": "hack",
      "network_name": "int-stage-0",
      "subnet_name": "int-192-168-128-0-slash-24",
      "secgroups_info": [
        {
          "name": "hack"
        }
      ]
    },
    {
      "router_name": "hack",
      "network_name": "int-stage-1",
      "subnet_name": "int-192-168-129-0-slash-24",
      "secgroups_info": [
        {
          "name": "hack"
        }
      ]
    },
    {
      "router_name": "hack",
      "network_name": "int-stage-2",
      "subnet_name": "int-192-168-130-0-slash-24",
      "secgroups_info": [
        {
          "name": "hack"
        }
      ]
    },
    {
      "router_name": "hack",
      "network_name": "int-stage-3",
      "subnet_name": "int-192-168-131-0-slash-24",
      "secgroups_info": [
        {
          "name": "hack"
        }
      ]
    }
  ],
  "gateways_info": [
    {
      "network_name": "public",
      "router_name": "hack"
    }
  ]
}
```

Response
```
{"vnets":[{"id":"87e3d043-56b8-419c-84ea-53e0c70c1fbe","name":"int-stage-0","status":"ACTIVE","subnets":[{"id":"949cf048-93b5-4bf7-946c-e94c0cf9454b","network_id":"87e3d043-56b8-419c-84ea-53e0c70c1fbe","name":"int-192-168-128-0-slash-24","ip_version":4,"cidr":"192.168.128.0/24","gateway_ip":"192.168.128.1","allocation_pools":[{"start":"192.168.128.100","end":"192.168.128.199"}],"enable_dhcp":true,"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"7291d11d-a2c5-4492-8d3d-f0a49da58209","name":"int-stage-1","status":"ACTIVE","subnets":[{"id":"93b35c15-2953-44f4-84e1-93a674c613d3","network_id":"7291d11d-a2c5-4492-8d3d-f0a49da58209","name":"int-192-168-129-0-slash-24","ip_version":4,"cidr":"192.168.129.0/24","gateway_ip":"192.168.129.1","allocation_pools":[{"start":"192.168.129.100","end":"192.168.129.199"}],"enable_dhcp":true,"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"a9e90774-cc8e-4a4b-86c6-8aeb2a63cd6f","name":"int-stage-2","status":"ACTIVE","subnets":[{"id":"9951f699-4708-4f1d-bb67-6df619c3b448","network_id":"a9e90774-cc8e-4a4b-86c6-8aeb2a63cd6f","name":"int-192-168-130-0-slash-24","ip_version":4,"cidr":"192.168.130.0/24","gateway_ip":"192.168.130.1","allocation_pools":[{"start":"192.168.130.100","end":"192.168.130.199"}],"enable_dhcp":true,"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"68b92034-2d78-4fae-9148-d049c822905b","name":"int-stage-3","status":"ACTIVE","subnets":[{"id":"1d0d239c-10a4-4169-af39-34f9eec7f69f","network_id":"68b92034-2d78-4fae-9148-d049c822905b","name":"int-192-168-131-0-slash-24","ip_version":4,"cidr":"192.168.131.0/24","gateway_ip":"192.168.131.1","allocation_pools":[{"start":"192.168.131.100","end":"192.168.131.199"}],"enable_dhcp":true,"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","name":"public","admin_state_up":true,"status":"ACTIVE","tenant_id":"8907b30a998647d5991547e9bbffa69a"}],"vrouter":{"status":"ACTIVE","gateway_info":{"network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d"},"name":"hack","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},"secgroup":{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack","description":"The fight networking security group","security_group_rules":[{"id":"f47550db-3f26-4151-abd8-06ee34e473d8","direction":"ingress","ethertype":"IPv4","security_group_id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","protocol":"tcp","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"70558f8b-df11-45b6-a293-4fb8788e9d15","direction":"ingress","ethertype":"IPv4","security_group_id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","protocol":"udp","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"f4f9493b-dc22-4663-9e57-14ccf701b114","direction":"ingress","ethertype":"IPv4","security_group_id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","protocol":"icmp","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},"ifaces_info":[{"router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack","network_id":"87e3d043-56b8-419c-84ea-53e0c70c1fbe","network_name":"int-stage-0","subnet_id":"949cf048-93b5-4bf7-946c-e94c0cf9454b","subnet_name":"int-192-168-128-0-slash-24","secgroups_info":[{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack"}],"port_id":"b583e434-8129-4949-9d41-510e2e0c038c","interface_info_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121"},{"router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack","network_id":"7291d11d-a2c5-4492-8d3d-f0a49da58209","network_name":"int-stage-1","subnet_id":"93b35c15-2953-44f4-84e1-93a674c613d3","subnet_name":"int-192-168-129-0-slash-24","secgroups_info":[{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack"}],"port_id":"6cb972d1-5130-491f-8ecb-194acb12b4dc","interface_info_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121"},{"router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack","network_id":"a9e90774-cc8e-4a4b-86c6-8aeb2a63cd6f","network_name":"int-stage-2","subnet_id":"9951f699-4708-4f1d-bb67-6df619c3b448","subnet_name":"int-192-168-130-0-slash-24","secgroups_info":[{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack"}],"port_id":"705081cc-1a92-4130-9417-424dfc25f6d3","interface_info_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121"},{"router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack","network_id":"68b92034-2d78-4fae-9148-d049c822905b","network_name":"int-stage-3","subnet_id":"1d0d239c-10a4-4169-af39-34f9eec7f69f","subnet_name":"int-192-168-131-0-slash-24","secgroups_info":[{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack"}],"port_id":"bbaa38e5-79ae-49ac-8665-6dd59e60b276","interface_info_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121"}],"gateways_info":[{"network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","network_name":"public","router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack"}],"ports":[{"id":"b583e434-8129-4949-9d41-510e2e0c038c","network_id":"87e3d043-56b8-419c-84ea-53e0c70c1fbe","name":"int-192-168-128-0-slash-24","status":"DOWN","mac_address":"fa:16:3e:12:13:43","fixed_ips":[{"subnet_id":"949cf048-93b5-4bf7-946c-e94c0cf9454b","ip_address":"192.168.128.100"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["75578fa7-4ca8-4a7b-9314-1aa45192d322"]},{"id":"6cb972d1-5130-491f-8ecb-194acb12b4dc","network_id":"7291d11d-a2c5-4492-8d3d-f0a49da58209","name":"int-192-168-129-0-slash-24","status":"DOWN","mac_address":"fa:16:3e:5a:4d:d7","fixed_ips":[{"subnet_id":"93b35c15-2953-44f4-84e1-93a674c613d3","ip_address":"192.168.129.100"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["75578fa7-4ca8-4a7b-9314-1aa45192d322"]},{"id":"705081cc-1a92-4130-9417-424dfc25f6d3","network_id":"a9e90774-cc8e-4a4b-86c6-8aeb2a63cd6f","name":"int-192-168-130-0-slash-24","status":"DOWN","mac_address":"fa:16:3e:7c:85:dd","fixed_ips":[{"subnet_id":"9951f699-4708-4f1d-bb67-6df619c3b448","ip_address":"192.168.130.100"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["75578fa7-4ca8-4a7b-9314-1aa45192d322"]},{"id":"bbaa38e5-79ae-49ac-8665-6dd59e60b276","network_id":"68b92034-2d78-4fae-9148-d049c822905b","name":"int-192-168-131-0-slash-24","status":"DOWN","mac_address":"fa:16:3e:94:02:d6","fixed_ips":[{"subnet_id":"1d0d239c-10a4-4169-af39-34f9eec7f69f","ip_address":"192.168.131.100"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["75578fa7-4ca8-4a7b-9314-1aa45192d322"]}],"interfaces_info":[{"subnet_id":"949cf048-93b5-4bf7-946c-e94c0cf9454b","port_id":"b583e434-8129-4949-9d41-510e2e0c038c","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"subnet_id":"93b35c15-2953-44f4-84e1-93a674c613d3","port_id":"6cb972d1-5130-491f-8ecb-194acb12b4dc","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"subnet_id":"9951f699-4708-4f1d-bb67-6df619c3b448","port_id":"705081cc-1a92-4130-9417-424dfc25f6d3","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"subnet_id":"1d0d239c-10a4-4169-af39-34f9eec7f69f","port_id":"bbaa38e5-79ae-49ac-8665-6dd59e60b276","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}]}
```

Command
```
[vagrant@localhost echopb]$ neutron net-list
+--------------------------------------+-------------+-------------------------------------------------------+
| id                                   | name        | subnets                                               |
+--------------------------------------+-------------+-------------------------------------------------------+
| 5a63e1f5-32b3-4dad-a080-839ae9253b77 | Name        |                                                       |
| 68b92034-2d78-4fae-9148-d049c822905b | int-stage-3 | 1d0d239c-10a4-4169-af39-34f9eec7f69f 192.168.131.0/24 |
| 7291d11d-a2c5-4492-8d3d-f0a49da58209 | int-stage-1 | 93b35c15-2953-44f4-84e1-93a674c613d3 192.168.129.0/24 |
| 830548b8-c7fc-435e-b144-b81f29b1e312 | private     | 931c8e3a-47df-42a6-aecd-9d81789b5fb7 192.168.0.0/24   |
| 87e3d043-56b8-419c-84ea-53e0c70c1fbe | int-stage-0 | 949cf048-93b5-4bf7-946c-e94c0cf9454b 192.168.128.0/24 |
| 92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d | public      | 27147d15-ce56-4913-9097-25d24a6d590e 10.100.151.0/24  |
| a9e90774-cc8e-4a4b-86c6-8aeb2a63cd6f | int-stage-2 | 9951f699-4708-4f1d-bb67-6df619c3b448 192.168.130.0/24 |
+--------------------------------------+-------------+-------------------------------------------------------+
[vagrant@localhost echopb]$ neutron subnet-list
+--------------------------------------+----------------------------+------------------+--------------------------------------------------------+
| id                                   | name                       | cidr             | allocation_pools                                       |
+--------------------------------------+----------------------------+------------------+--------------------------------------------------------+
| 1d0d239c-10a4-4169-af39-34f9eec7f69f | int-192-168-131-0-slash-24 | 192.168.131.0/24 | {"start": "192.168.131.100", "end": "192.168.131.199"} |
| 27147d15-ce56-4913-9097-25d24a6d590e | 10.100.151.0/24            | 10.100.151.0/24  | {"start": "10.100.151.50", "end": "10.100.151.240"}    |
| 931c8e3a-47df-42a6-aecd-9d81789b5fb7 | 192.168.0.0/24             | 192.168.0.0/24   | {"start": "192.168.0.2", "end": "192.168.0.254"}       |
| 93b35c15-2953-44f4-84e1-93a674c613d3 | int-192-168-129-0-slash-24 | 192.168.129.0/24 | {"start": "192.168.129.100", "end": "192.168.129.199"} |
| 949cf048-93b5-4bf7-946c-e94c0cf9454b | int-192-168-128-0-slash-24 | 192.168.128.0/24 | {"start": "192.168.128.100", "end": "192.168.128.199"} |
| 9951f699-4708-4f1d-bb67-6df619c3b448 | int-192-168-130-0-slash-24 | 192.168.130.0/24 | {"start": "192.168.130.100", "end": "192.168.130.199"} |
+--------------------------------------+----------------------------+------------------+--------------------------------------------------------+
[vagrant@localhost echopb]$ neutron router-list
+--------------------------------------+------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------+-------+
| id                                   | name | external_gateway_info                                                                                                                                                                     | distributed | ha    |
+--------------------------------------+------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------+-------+
| 951d2d2f-5417-413e-a18d-b0627d09b8dd | test | {"network_id": "92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d", "enable_snat": true, "external_fixed_ips": [{"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.50"}]} | False       | False |
| abc8b477-fab7-40c6-a11d-f9c2d5980121 | hack | {"network_id": "92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d", "enable_snat": true, "external_fixed_ips": [{"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.79"}]} | False       | False |
+--------------------------------------+------+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------+-------+
[vagrant@localhost echopb]$ neutron security-group-list
+--------------------------------------+---------+----------------------------------------------------------------------+
| id                                   | name    | security_group_rules                                                 |
+--------------------------------------+---------+----------------------------------------------------------------------+
| 4da19416-41cb-469b-b5fb-b02d24aaff47 | default | egress, IPv4                                                         |
|                                      |         | egress, IPv6                                                         |
|                                      |         | ingress, IPv4, 1-65535/tcp, remote_ip_prefix: 0.0.0.0/0              |
|                                      |         | ingress, IPv4, 1-65535/udp, remote_ip_prefix: 0.0.0.0/0              |
|                                      |         | ingress, IPv4, icmp, remote_ip_prefix: 0.0.0.0/0                     |
|                                      |         | ingress, IPv4, remote_group_id: 4da19416-41cb-469b-b5fb-b02d24aaff47 |
|                                      |         | ingress, IPv6, remote_group_id: 4da19416-41cb-469b-b5fb-b02d24aaff47 |
| 75578fa7-4ca8-4a7b-9314-1aa45192d322 | hack    | egress, IPv4                                                         |
|                                      |         | egress, IPv6                                                         |
|                                      |         | ingress, IPv4, icmp                                                  |
|                                      |         | ingress, IPv4, tcp                                                   |
|                                      |         | ingress, IPv4, udp                                                   |
+--------------------------------------+---------+----------------------------------------------------------------------+
[vagrant@localhost echopb]$ neutron security-group-rule-list
+--------------------------------------+----------------+-----------+-----------+---------------+------------------+
| id                                   | security_group | direction | ethertype | protocol/port | remote           |
+--------------------------------------+----------------+-----------+-----------+---------------+------------------+
| 0ebee491-fc1f-4fb2-918e-4f04d812f2a1 | hack           | egress    | IPv4      | any           | any              |
| 70558f8b-df11-45b6-a293-4fb8788e9d15 | hack           | ingress   | IPv4      | udp           | any              |
| cc2e908d-646c-4e4d-90c7-80528624d495 | hack           | egress    | IPv6      | any           | any              |
| cd9e67d8-5335-4e20-b76e-1fc66d89198f | default        | ingress   | IPv4      | any           | default (group)  |
| ed8d662a-2c33-4363-bed6-f349050360eb | default        | ingress   | IPv6      | any           | default (group)  |
| f47550db-3f26-4151-abd8-06ee34e473d8 | hack           | ingress   | IPv4      | tcp           | any              |
| f4f9493b-dc22-4663-9e57-14ccf701b114 | hack           | ingress   | IPv4      | icmp          | any              |
+--------------------------------------+----------------+-----------+-----------+---------------+------------------+
[vagrant@localhost echopb]$ neutron port-list
+--------------------------------------+----------------------------+-------------------+----------------------------------------------------------------------------------------+
| id                                   | name                       | mac_address       | fixed_ips                                                                              |
+--------------------------------------+----------------------------+-------------------+----------------------------------------------------------------------------------------+
| 6cb972d1-5130-491f-8ecb-194acb12b4dc | int-192-168-129-0-slash-24 | fa:16:3e:5a:4d:d7 | {"subnet_id": "93b35c15-2953-44f4-84e1-93a674c613d3", "ip_address": "192.168.129.100"} |
| 705081cc-1a92-4130-9417-424dfc25f6d3 | int-192-168-130-0-slash-24 | fa:16:3e:7c:85:dd | {"subnet_id": "9951f699-4708-4f1d-bb67-6df619c3b448", "ip_address": "192.168.130.100"} |
| b583e434-8129-4949-9d41-510e2e0c038c | int-192-168-128-0-slash-24 | fa:16:3e:12:13:43 | {"subnet_id": "949cf048-93b5-4bf7-946c-e94c0cf9454b", "ip_address": "192.168.128.100"} |
| bbaa38e5-79ae-49ac-8665-6dd59e60b276 | int-192-168-131-0-slash-24 | fa:16:3e:94:02:d6 | {"subnet_id": "1d0d239c-10a4-4169-af39-34f9eec7f69f", "ip_address": "192.168.131.100"} |
+--------------------------------------+----------------------------+-------------------+----------------------------------------------------------------------------------------+
```

### 0.1.0

2017-06-23

#### Batch create Virtual Machines

_POST /v1/boot_

Request
```
{
  "flavor_name": "large",
  "image_name": "target86a-phpcmsV9_blindSQL-Ubuntu-cl2",
  "min_count": 2,
  "max_count": 4,
  "secgroups_info": [],
  "user_data": [],
  "network_name": "private",
  "floating_network_name": "public",
  "personality": [],
  "name_prefix": "............"
}
```

Response
```
{"flavor_id":"bf289de9-bd91-4d12-9d86-a4e8889f412f","flavor_name":"large","image_id":"83a93848-2b94-4cde-89b8-271bb9b8bd83","image_name":"target86a-phpcmsV9_blindSQL-Ubuntu-cl2","min_count":4,"max_count":4,"network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","network_name":"private","floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_network_name":"private","ports":[{"id":"806194d4-f561-43bc-8e25-ce1d1987d9f8","network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","name":"............-0","status":"DOWN","mac_address":"fa:16:3e:ba:86:f9","fixed_ips":[{"subnet_id":"931c8e3a-47df-42a6-aecd-9d81789b5fb7","ip_address":"192.168.0.5"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["4da19416-41cb-469b-b5fb-b02d24aaff47"]},{"id":"16c2bd2a-fdaa-42ba-b854-4508ba486d94","network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","name":"............-1","status":"DOWN","mac_address":"fa:16:3e:50:8c:82","fixed_ips":[{"subnet_id":"931c8e3a-47df-42a6-aecd-9d81789b5fb7","ip_address":"192.168.0.6"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["4da19416-41cb-469b-b5fb-b02d24aaff47"]},{"id":"5f022746-4998-4c0a-8291-4acae6c11e04","network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","name":"............-2","status":"DOWN","mac_address":"fa:16:3e:2a:d0:7e","fixed_ips":[{"subnet_id":"931c8e3a-47df-42a6-aecd-9d81789b5fb7","ip_address":"192.168.0.7"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["4da19416-41cb-469b-b5fb-b02d24aaff47"]},{"id":"036a26cb-8647-4522-8538-abfa7068bdf8","network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","name":"............-3","status":"DOWN","mac_address":"fa:16:3e:ab:fe:26","fixed_ips":[{"subnet_id":"931c8e3a-47df-42a6-aecd-9d81789b5fb7","ip_address":"192.168.0.8"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["4da19416-41cb-469b-b5fb-b02d24aaff47"]}],"servers":[{"id":"82b0066a-3c96-4854-a466-f1dd1ea1569f","updated":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","adminPass":"m9JhekPZgcUR"},{"id":"ed2c2068-c717-4b09-9652-50ece6b153a4","updated":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","adminPass":"XeEJzyr9ZenJ"},{"id":"25436d62-49ee-4e32-aedc-fbab11d7cec2","updated":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","adminPass":"Bvu7dgFZD6yU"},{"id":"f93382f2-8447-424a-9f1e-8061c3b6eb71","updated":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","adminPass":"ML5aqubvQDRp"}],"floating_ips":[{"floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_ip_address":"10.100.151.106","port_id":"806194d4-f561-43bc-8e25-ce1d1987d9f8","fixed_ip_address":"192.168.0.5","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_ip_address":"10.100.151.107","port_id":"16c2bd2a-fdaa-42ba-b854-4508ba486d94","fixed_ip_address":"192.168.0.6","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_ip_address":"10.100.151.108","port_id":"5f022746-4998-4c0a-8291-4acae6c11e04","fixed_ip_address":"192.168.0.7","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_ip_address":"10.100.151.109","port_id":"036a26cb-8647-4522-8538-abfa7068bdf8","fixed_ip_address":"192.168.0.8","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"port_server_pairs":{"036a26cb-8647-4522-8538-abfa7068bdf8":"f93382f2-8447-424a-9f1e-8061c3b6eb71","16c2bd2a-fdaa-42ba-b854-4508ba486d94":"ed2c2068-c717-4b09-9652-50ece6b153a4","5f022746-4998-4c0a-8291-4acae6c11e04":"25436d62-49ee-4e32-aedc-fbab11d7cec2","806194d4-f561-43bc-8e25-ce1d1987d9f8":"82b0066a-3c96-4854-a466-f1dd1ea1569f"}}
```

Command
```
[vagrant@localhost echopb]$ openstack server list
+--------------------------------------+-------------------------------------------------------+--------+------------------------+
| ID                                   | Name                                                  | Status | Networks               |
+--------------------------------------+-------------------------------------------------------+--------+------------------------+
| f93382f2-8447-424a-9f1e-8061c3b6eb71 | ............-3                                        | BUILD  |                        |
| 25436d62-49ee-4e32-aedc-fbab11d7cec2 | ............-2                                        | BUILD  |                        |
| ed2c2068-c717-4b09-9652-50ece6b153a4 | ............-1                                        | BUILD  |                        |
| 82b0066a-3c96-4854-a466-f1dd1ea1569f | ............-0                                        | BUILD  |                        |
+--------------------------------------+-------------------------------------------------------+--------+------------------------+
[vagrant@localhost echopb]$ openstack server list
+--------------------------------------+-------------------------------------------------------+--------+-------------------------------------+
| ID                                   | Name                                                  | Status | Networks                            |
+--------------------------------------+-------------------------------------------------------+--------+-------------------------------------+
| f93382f2-8447-424a-9f1e-8061c3b6eb71 | ............-3                                        | ACTIVE | private=192.168.0.8, 10.100.151.109 |
| 25436d62-49ee-4e32-aedc-fbab11d7cec2 | ............-2                                        | ACTIVE | private=192.168.0.7, 10.100.151.108 |
| ed2c2068-c717-4b09-9652-50ece6b153a4 | ............-1                                        | ACTIVE | private=192.168.0.6, 10.100.151.107 |
| 82b0066a-3c96-4854-a466-f1dd1ea1569f | ............-0                                        | ACTIVE | private=192.168.0.5, 10.100.151.106 |
+--------------------------------------+-------------------------------------------------------+--------+-------------------------------------+
```