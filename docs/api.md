# about API

Swagger 2.0

![屏幕快照 2017-06-24 下午1.49.46.png](./屏幕快照%202017-06-24%20下午1.49.46.png)

## Release 0.1.1

* create base networking
* get all images
* get all flavors

### Networking

test JSON
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

test run
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ kopos test net
2017-06-25 05:26:17.486172 I | {[name:"int-stage-0" subnets:<name:"int-192-168-128-0-slash-24" cidr:"192.168.128.0/24" enable_dhcp:true >  name:"int-stage-1" subnets:<name:"int-192-168-129-0-slash-24" cidr:"192.168.129.0/24" enable_dhcp:true >  name:"int-stage-2" subnets:<name:"int-192-168-130-0-slash-24" cidr:"192.168.130.0/24" enable_dhcp:true >  name:"int-stage-3" subnets:<name:"int-192-168-131-0-slash-24" cidr:"192.168.131.0/24" enable_dhcp:true >  name:"public" admin_state_up:true subnets:<name:"10.100.151.0/24" cidr:"10.100.151.0/24" gateway_ip:"10.100.151.1" allocation_pools:<start:"10.100.151.50" end:"10.100.151.240" > > shared:true ] name:"hack"  name:"hack" security_group_rules:<direction:"ingress" protocol:"tcp" > security_group_rules:<direction:"ingress" protocol:"udp" > security_group_rules:<direction:"ingress" protocol:"icmp" >  [router_name:"hack" network_name:"int-stage-0" subnet_name:"int-192-168-128-0-slash-24" secgroups_info:<name:"hack" >  router_name:"hack" network_name:"int-stage-1" subnet_name:"int-192-168-129-0-slash-24" secgroups_info:<name:"hack" >  router_name:"hack" network_name:"int-stage-2" subnet_name:"int-192-168-130-0-slash-24" secgroups_info:<name:"hack" >  router_name:"hack" network_name:"int-stage-3" subnet_name:"int-192-168-131-0-slash-24" secgroups_info:<name:"hack" > ] [network_name:"public" router_name:"hack" ] 0  [] []}
response Status: 200 OK
response Headers: map[Content-Type:[application/json] Date:[Sun, 25 Jun 2017 12:26:20 GMT]]
response Body: {"vnets":[{"id":"87e3d043-56b8-419c-84ea-53e0c70c1fbe","name":"int-stage-0","status":"ACTIVE","subnets":[{"id":"949cf048-93b5-4bf7-946c-e94c0cf9454b","network_id":"87e3d043-56b8-419c-84ea-53e0c70c1fbe","name":"int-192-168-128-0-slash-24","ip_version":4,"cidr":"192.168.128.0/24","gateway_ip":"192.168.128.1","allocation_pools":[{"start":"192.168.128.100","end":"192.168.128.199"}],"enable_dhcp":true,"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"7291d11d-a2c5-4492-8d3d-f0a49da58209","name":"int-stage-1","status":"ACTIVE","subnets":[{"id":"93b35c15-2953-44f4-84e1-93a674c613d3","network_id":"7291d11d-a2c5-4492-8d3d-f0a49da58209","name":"int-192-168-129-0-slash-24","ip_version":4,"cidr":"192.168.129.0/24","gateway_ip":"192.168.129.1","allocation_pools":[{"start":"192.168.129.100","end":"192.168.129.199"}],"enable_dhcp":true,"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"a9e90774-cc8e-4a4b-86c6-8aeb2a63cd6f","name":"int-stage-2","status":"ACTIVE","subnets":[{"id":"9951f699-4708-4f1d-bb67-6df619c3b448","network_id":"a9e90774-cc8e-4a4b-86c6-8aeb2a63cd6f","name":"int-192-168-130-0-slash-24","ip_version":4,"cidr":"192.168.130.0/24","gateway_ip":"192.168.130.1","allocation_pools":[{"start":"192.168.130.100","end":"192.168.130.199"}],"enable_dhcp":true,"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"68b92034-2d78-4fae-9148-d049c822905b","name":"int-stage-3","status":"ACTIVE","subnets":[{"id":"1d0d239c-10a4-4169-af39-34f9eec7f69f","network_id":"68b92034-2d78-4fae-9148-d049c822905b","name":"int-192-168-131-0-slash-24","ip_version":4,"cidr":"192.168.131.0/24","gateway_ip":"192.168.131.1","allocation_pools":[{"start":"192.168.131.100","end":"192.168.131.199"}],"enable_dhcp":true,"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","name":"public","admin_state_up":true,"status":"ACTIVE","tenant_id":"8907b30a998647d5991547e9bbffa69a"}],"vrouter":{"status":"ACTIVE","gateway_info":{"network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d"},"name":"hack","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},"secgroup":{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack","description":"The fight networking security group","security_group_rules":[{"id":"f47550db-3f26-4151-abd8-06ee34e473d8","direction":"ingress","ethertype":"IPv4","security_group_id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","protocol":"tcp","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"70558f8b-df11-45b6-a293-4fb8788e9d15","direction":"ingress","ethertype":"IPv4","security_group_id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","protocol":"udp","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"id":"f4f9493b-dc22-4663-9e57-14ccf701b114","direction":"ingress","ethertype":"IPv4","security_group_id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","protocol":"icmp","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},"ifaces_info":[{"router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack","network_id":"87e3d043-56b8-419c-84ea-53e0c70c1fbe","network_name":"int-stage-0","subnet_id":"949cf048-93b5-4bf7-946c-e94c0cf9454b","subnet_name":"int-192-168-128-0-slash-24","secgroups_info":[{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack"}],"port_id":"b583e434-8129-4949-9d41-510e2e0c038c","interface_info_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121"},{"router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack","network_id":"7291d11d-a2c5-4492-8d3d-f0a49da58209","network_name":"int-stage-1","subnet_id":"93b35c15-2953-44f4-84e1-93a674c613d3","subnet_name":"int-192-168-129-0-slash-24","secgroups_info":[{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack"}],"port_id":"6cb972d1-5130-491f-8ecb-194acb12b4dc","interface_info_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121"},{"router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack","network_id":"a9e90774-cc8e-4a4b-86c6-8aeb2a63cd6f","network_name":"int-stage-2","subnet_id":"9951f699-4708-4f1d-bb67-6df619c3b448","subnet_name":"int-192-168-130-0-slash-24","secgroups_info":[{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack"}],"port_id":"705081cc-1a92-4130-9417-424dfc25f6d3","interface_info_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121"},{"router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack","network_id":"68b92034-2d78-4fae-9148-d049c822905b","network_name":"int-stage-3","subnet_id":"1d0d239c-10a4-4169-af39-34f9eec7f69f","subnet_name":"int-192-168-131-0-slash-24","secgroups_info":[{"id":"75578fa7-4ca8-4a7b-9314-1aa45192d322","name":"hack"}],"port_id":"bbaa38e5-79ae-49ac-8665-6dd59e60b276","interface_info_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121"}],"gateways_info":[{"network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","network_name":"public","router_id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","router_name":"hack"}],"ports":[{"id":"b583e434-8129-4949-9d41-510e2e0c038c","network_id":"87e3d043-56b8-419c-84ea-53e0c70c1fbe","name":"int-192-168-128-0-slash-24","status":"DOWN","mac_address":"fa:16:3e:12:13:43","fixed_ips":[{"subnet_id":"949cf048-93b5-4bf7-946c-e94c0cf9454b","ip_address":"192.168.128.100"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["75578fa7-4ca8-4a7b-9314-1aa45192d322"]},{"id":"6cb972d1-5130-491f-8ecb-194acb12b4dc","network_id":"7291d11d-a2c5-4492-8d3d-f0a49da58209","name":"int-192-168-129-0-slash-24","status":"DOWN","mac_address":"fa:16:3e:5a:4d:d7","fixed_ips":[{"subnet_id":"93b35c15-2953-44f4-84e1-93a674c613d3","ip_address":"192.168.129.100"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["75578fa7-4ca8-4a7b-9314-1aa45192d322"]},{"id":"705081cc-1a92-4130-9417-424dfc25f6d3","network_id":"a9e90774-cc8e-4a4b-86c6-8aeb2a63cd6f","name":"int-192-168-130-0-slash-24","status":"DOWN","mac_address":"fa:16:3e:7c:85:dd","fixed_ips":[{"subnet_id":"9951f699-4708-4f1d-bb67-6df619c3b448","ip_address":"192.168.130.100"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["75578fa7-4ca8-4a7b-9314-1aa45192d322"]},{"id":"bbaa38e5-79ae-49ac-8665-6dd59e60b276","network_id":"68b92034-2d78-4fae-9148-d049c822905b","name":"int-192-168-131-0-slash-24","status":"DOWN","mac_address":"fa:16:3e:94:02:d6","fixed_ips":[{"subnet_id":"1d0d239c-10a4-4169-af39-34f9eec7f69f","ip_address":"192.168.131.100"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["75578fa7-4ca8-4a7b-9314-1aa45192d322"]}],"interfaces_info":[{"subnet_id":"949cf048-93b5-4bf7-946c-e94c0cf9454b","port_id":"b583e434-8129-4949-9d41-510e2e0c038c","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"subnet_id":"93b35c15-2953-44f4-84e1-93a674c613d3","port_id":"6cb972d1-5130-491f-8ecb-194acb12b4dc","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"subnet_id":"9951f699-4708-4f1d-bb67-6df619c3b448","port_id":"705081cc-1a92-4130-9417-424dfc25f6d3","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"subnet_id":"1d0d239c-10a4-4169-af39-34f9eec7f69f","port_id":"bbaa38e5-79ae-49ac-8665-6dd59e60b276","id":"abc8b477-fab7-40c6-a11d-f9c2d5980121","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}]}
```

Via openstack command client
```
[vagrant@localhost echopb]$ neutron net-list
+--------------------------------------+-------------+-------------------------------------------------------+
| id                                   | name        | subnets                                               |
+--------------------------------------+-------------+-------------------------------------------------------+
| 5a63e1f5-32b3-4dad-a080-839ae9253b77 | Name        |                                                       |
| 68b92034-2d78-4fae-9148-d049c822905b | int-stage-3 | 1d0d239c-10a4-4169-af39-34f9eec7f69f 192.168.131.0/24 |
| 7291d11d-a2c5-4492-8d3d-f0a49da58209 | int-stage-1 | 93b35c15-2953-44f4-84e1-93a674c613d3 192.168.129.0/24 |
| 7931ca33-5cc6-47bc-8c0c-28ed9b89f09c | private_net | 5bce8649-630b-401f-aec6-4ce644f9917f 10.20.30.0/24    |
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
| 5bce8649-630b-401f-aec6-4ce644f9917f | private_subnet             | 10.20.30.0/24    | {"start": "10.20.30.2", "end": "10.20.30.254"}         |
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
| 7a44e06c-310f-439c-b466-19ee7b8434d0 | default | egress, IPv4                                                         |
|                                      |         | egress, IPv6                                                         |
|                                      |         | ingress, IPv4, remote_group_id: 7a44e06c-310f-439c-b466-19ee7b8434d0 |
|                                      |         | ingress, IPv6, remote_group_id: 7a44e06c-310f-439c-b466-19ee7b8434d0 |
| 7e0ac9fa-4b84-4c7d-9c43-f19f455640f3 | default | egress, IPv4                                                         |
|                                      |         | egress, IPv6                                                         |
|                                      |         | ingress, IPv4, remote_group_id: 7e0ac9fa-4b84-4c7d-9c43-f19f455640f3 |
|                                      |         | ingress, IPv6, remote_group_id: 7e0ac9fa-4b84-4c7d-9c43-f19f455640f3 |
| ba595dd6-88c5-423e-aba7-ff3878c8f826 | default | egress, IPv4                                                         |
|                                      |         | egress, IPv6                                                         |
|                                      |         | ingress, IPv4, remote_group_id: ba595dd6-88c5-423e-aba7-ff3878c8f826 |
|                                      |         | ingress, IPv6, remote_group_id: ba595dd6-88c5-423e-aba7-ff3878c8f826 |
+--------------------------------------+---------+----------------------------------------------------------------------+
[vagrant@localhost echopb]$ neutron security-group-rules-list
Unknown command [u'security-group-rules-list']
[vagrant@localhost echopb]$ neutron security-group-rule-list
+--------------------------------------+----------------+-----------+-----------+---------------+------------------+
| id                                   | security_group | direction | ethertype | protocol/port | remote           |
+--------------------------------------+----------------+-----------+-----------+---------------+------------------+
| 058df718-8fca-4019-85cc-fa817cf879bf | default        | ingress   | IPv4      | any           | default (group)  |
| 0b15d85a-603f-4bfc-8967-aff5ab7e5816 | default        | ingress   | IPv4      | 1-65535/udp   | 0.0.0.0/0 (CIDR) |
| 0ebee491-fc1f-4fb2-918e-4f04d812f2a1 | hack           | egress    | IPv4      | any           | any              |
| 0f181667-2cb9-4ef5-86da-50c5f051fb3b | default        | egress    | IPv4      | any           | any              |
| 1c1ca39c-aa7c-4d62-a9ea-69bb96eaaa18 | default        | ingress   | IPv6      | any           | default (group)  |
| 1e3da529-47ec-408f-b067-f1e68b17d3b8 | default        | ingress   | IPv4      | 1-65535/tcp   | 0.0.0.0/0 (CIDR) |
| 1e4c18f9-1601-4812-97bd-8eae74c1b64e | default        | egress    | IPv6      | any           | any              |
| 1e6856ef-e1ad-4ac7-b241-be650286cb78 | default        | egress    | IPv4      | any           | any              |
| 1f8d8d28-4e7c-40d6-b2b8-5fa778995a18 | default        | egress    | IPv4      | any           | any              |
| 2295e9ea-c06d-4fe3-af4b-6769f6bd3018 | default        | egress    | IPv6      | any           | any              |
| 2b6f556c-857f-4da3-9572-0ef97bc7a921 | default        | ingress   | IPv6      | any           | default (group)  |
| 309eb89f-9678-4791-af39-ad2d9968c3f6 | default        | ingress   | IPv6      | any           | default (group)  |
| 56b0eb3c-5145-433b-8dd7-652ef7b245bf | default        | egress    | IPv6      | any           | any              |
| 57f2214a-d285-485f-a2f0-fefac0b5a6b3 | default        | egress    | IPv4      | any           | any              |
| 6c757a47-77c6-4b6f-ac67-88620eb98a68 | default        | egress    | IPv6      | any           | any              |
| 6ce6a067-7cbc-46cb-9084-5a88306bf2fb | default        | ingress   | IPv4      | any           | default (group)  |
| 70558f8b-df11-45b6-a293-4fb8788e9d15 | hack           | ingress   | IPv4      | udp           | any              |
| af0051fe-b640-421d-add8-1d5017841806 | default        | ingress   | IPv4      | any           | default (group)  |
| b9cf64af-1882-4235-b28e-45a5866e1ff0 | default        | ingress   | IPv4      | icmp          | 0.0.0.0/0 (CIDR) |
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
| 00d44e7a-6d31-4995-bb92-0088bda70ab3 | 111-2                      | fa:16:3e:c4:a4:23 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.22"}    |
| 07c1c2ae-7fa7-4847-8cca-06c3a6ddf907 | 哈哈哈哈-48                | fa:16:3e:e2:f0:d0 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.8"}     |
| 0ea2da16-7a9d-4354-8437-8b7b9babd3c8 | 333-2                      | fa:16:3e:52:1c:22 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.116"}   |
| 16608213-5023-4507-8d0b-15c1b1ab6ec1 |                            | fa:16:3e:22:18:78 | {"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.79"}   |
| 1f1cf7d0-5d79-4b4a-8721-f13de420c192 | 111-0                      | fa:16:3e:3c:85:68 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.47"}    |
| 234e498f-cace-49bd-b1bc-7251b3092271 | 111-1                      | fa:16:3e:f4:90:4a | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.57"}    |
| 26d955b6-4422-456d-8b9c-d66c1034d999 | 111-2                      | fa:16:3e:38:5e:9a | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.61"}    |
| 2ba78099-0c0a-43cc-9212-05ca0d81dbc2 | 哈哈哈哈-46                | fa:16:3e:5a:9f:55 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.78"}    |
| 2bc9af08-6b2a-4ea5-b699-9d13bdf89db5 | 哈哈哈哈-44                | fa:16:3e:05:1a:27 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.76"}    |
| 305faf7a-1a3d-4440-b57a-ef948fd17a0b | 111-2                      | fa:16:3e:e6:72:6d | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.34"}    |
| 3709fce5-446a-4e46-9092-540a1179d6f2 | 1111-0                     | fa:16:3e:1c:7d:8a | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.84"}    |
| 3819089e-5642-4677-94a4-b55eea757cda | 哈哈哈哈-41                | fa:16:3e:96:1d:c3 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.73"}    |
| 3a10474d-9357-488d-b298-6440288a4c05 | 111-1                      | fa:16:3e:80:48:78 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.21"}    |
| 446d5791-97a0-48a6-aadf-0d288019e4c9 | 111-0                      | fa:16:3e:82:29:70 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.50"}    |
| 4799c6df-cbec-4ead-8022-df42fe4600b4 | 哈哈哈哈-45                | fa:16:3e:6f:d3:e1 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.77"}    |
| 48606d5d-3fd4-4479-863a-3746c19dfafb | 333-0                      | fa:16:3e:86:ae:01 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.120"}   |
| 4868476a-c642-4d29-bbd7-94ef98893346 |                            | fa:16:3e:09:6d:97 | {"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.50"}   |
| 4b9bd22f-5c41-48a2-a507-7e25791373bf | 111-1                      | fa:16:3e:8c:73:30 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.18"}    |
| 4cdf23cc-047a-48f3-8465-5bd7a509faca | 哈哈哈哈-43                | fa:16:3e:7f:d4:2b | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.75"}    |
| 4d688f24-1462-4b91-babe-caa645aad224 | 1111-1                     | fa:16:3e:56:2c:b6 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.85"}    |
| 4f9f35f6-f1c8-4e06-b917-828fbe23b89f | 111-0                      | fa:16:3e:b0:29:95 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.56"}    |
| 5952c07d-0602-4bbe-97a9-723556a00bb7 | 111-2                      | fa:16:3e:28:5f:c7 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.49"}    |
| 64879b56-8a68-41e5-b863-05f77e39d7d3 | 111-1                      | fa:16:3e:6d:43:5e | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.60"}    |
| 657d644b-33a6-4d82-925f-1be5ea29ac0b | 333-0                      | fa:16:3e:75:3c:13 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.114"}   |
| 6cb972d1-5130-491f-8ecb-194acb12b4dc | int-192-168-129-0-slash-24 | fa:16:3e:5a:4d:d7 | {"subnet_id": "93b35c15-2953-44f4-84e1-93a674c613d3", "ip_address": "192.168.129.100"} |
| 6f0a48c6-0598-4cef-b21b-8e97a9b69e35 | 333-1                      | fa:16:3e:b7:48:41 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.121"}   |
| 705081cc-1a92-4130-9417-424dfc25f6d3 | int-192-168-130-0-slash-24 | fa:16:3e:7c:85:dd | {"subnet_id": "9951f699-4708-4f1d-bb67-6df619c3b448", "ip_address": "192.168.130.100"} |
| 76633581-b6ab-4e46-9b0c-181f800dac7b | 333-2                      | fa:16:3e:d4:e6:7c | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.122"}   |
| 784b022d-93f4-4890-94ac-27a595f6fa61 | 1111-2                     | fa:16:3e:5b:c8:a9 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.86"}    |
| 7936f4fd-0f8a-453d-9415-e552fa799ddd | 哈哈哈哈-49                | fa:16:3e:71:e2:c2 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.80"}    |
| 7c561ce7-ab33-498a-89d6-1cf0f54bafe3 | 111-0                      | fa:16:3e:f3:b8:0e | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.32"}    |
| 7dfc4c35-bce4-492f-a23f-cbe74ebdcd4b |                            | fa:16:3e:5c:bd:21 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.2"}     |
| 7ecfb26e-5741-47fb-a6e7-6de62a6c8870 | 111-0                      | fa:16:3e:50:13:a2 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.59"}    |
| 7f70337f-60c7-4b10-97f7-3fb6949be174 |                            | fa:16:3e:08:ed:c8 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.1"}     |
| 93583b29-a0dd-4017-90b5-8abcd78a0845 | 哈哈哈哈-47                | fa:16:3e:1b:ce:70 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.79"}    |
| 9483d682-5ebd-4efb-876c-629ac54c0870 | 111-2                      | fa:16:3e:fa:93:85 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.40"}    |
| a29ff001-73be-469c-a44b-090f91222125 | 111-2                      | fa:16:3e:dd:88:cc | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.52"}    |
| a9099e99-0c77-451d-b321-568cb4bfc9b7 | 333-1                      | fa:16:3e:b0:be:57 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.115"}   |
| b071211d-ed9c-48d2-ae92-86d6baa6b8d1 | 333-0                      | fa:16:3e:29:5c:bd | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.117"}   |
| b4bb6c17-f85b-4d6f-9302-da6767d0a872 |                            | fa:16:3e:2a:db:0f | {"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.77"}   |
| b583e434-8129-4949-9d41-510e2e0c038c | int-192-168-128-0-slash-24 | fa:16:3e:12:13:43 | {"subnet_id": "949cf048-93b5-4bf7-946c-e94c0cf9454b", "ip_address": "192.168.128.100"} |
| b9a2f027-007e-4e33-8522-e1f73e117503 |                            | fa:16:3e:16:31:b7 | {"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.76"}   |
| bbaa38e5-79ae-49ac-8665-6dd59e60b276 | int-192-168-131-0-slash-24 | fa:16:3e:94:02:d6 | {"subnet_id": "1d0d239c-10a4-4169-af39-34f9eec7f69f", "ip_address": "192.168.131.100"} |
| c4118d1b-1c67-4d6f-86eb-11ef690e86b1 | 111-1                      | fa:16:3e:6e:f7:1a | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.48"}    |
| c8fec633-ed59-49a8-9c24-9a3246927398 | 111-1                      | fa:16:3e:37:a7:b8 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.51"}    |
| cab1e4b4-9c88-4455-81de-5a42f052dcb0 | 111-0                      | fa:16:3e:e6:cf:05 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.20"}    |
| cb08c09a-0c96-4ab0-b45c-bb20f07fb039 | 333-1                      | fa:16:3e:b1:73:0c | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.118"}   |
| d2a89f75-266d-4489-b768-d20df30030be | 111-1                      | fa:16:3e:18:73:e3 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.39"}    |
| d6778ca5-c24e-45ba-89e7-72ccc1fb3047 | 111-2                      | fa:16:3e:9e:a2:61 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.58"}    |
| d86ff47f-ebdf-41b5-a629-bf8457a41b1c | 111-1                      | fa:16:3e:1a:67:aa | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.30"}    |
| db03c619-b6bb-4e2f-ae7d-d5b989a9f91c | 哈哈哈哈-42                | fa:16:3e:d3:5c:d3 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.74"}    |
| dcd07ccb-02bf-4885-8341-e3a2ee1287c1 |                            | fa:16:3e:4d:3b:1b | {"subnet_id": "5bce8649-630b-401f-aec6-4ce644f9917f", "ip_address": "10.20.30.3"}      |
| dd025f5c-ca40-4e9e-ab43-70718bfe605b | 111-0                      | fa:16:3e:10:9a:6f | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.29"}    |
| e2a1ebc3-504e-47c0-a94f-8dc736dd6196 |                            | fa:16:3e:c8:b4:b6 | {"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.78"}   |
| e4259eb7-a59e-4da4-ada8-02a307902860 | 111-0                      | fa:16:3e:0a:4c:6f | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.17"}    |
| e55ff203-9f0e-497a-8b4f-4ca27c10d8d0 | 111-2                      | fa:16:3e:4f:64:4b | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.31"}    |
| ec553da2-3fa4-4904-9aa0-656dc4075150 | 111-0                      | fa:16:3e:8b:1a:e8 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.38"}    |
| eca2cef3-6789-4c8d-8267-6232e4f09592 | 111-2                      | fa:16:3e:20:36:0b | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.64"}    |
| eedd3022-4553-4d84-8640-8f3260a02aa6 | 111-2                      | fa:16:3e:27:6f:d3 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.19"}    |
| ef43972c-24b8-47af-9436-8370a814f15e | 111-1                      | fa:16:3e:df:6e:6f | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.33"}    |
| f6e85b28-f8f5-4590-8816-b7e3074828aa | 333-2                      | fa:16:3e:19:4a:33 | {"subnet_id": "931c8e3a-47df-42a6-aecd-9d81789b5fb7", "ip_address": "192.168.0.119"}   |
+--------------------------------------+----------------------------+-------------------+----------------------------------------------------------------------------------------+
```

### ImageService

Test run
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ curl -iL http://127.0.0.1:10001/v1/flavors
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 25 Jun 2017 08:39:06 GMT
Content-Length: 503

{"flavors":[{"id":"1","disk":1,"ram":512,"name":"m1.tiny","rxtx_factor":1,"vcpus":1},{"id":"2","disk":20,"ram":2048,"name":"m1.small","rxtx_factor":1,"vcpus":1},{"id":"3","disk":40,"ram":4096,"name":"m1.medium","rxtx_factor":1,"vcpus":2},{"id":"4","disk":80,"ram":8192,"name":"m1.large","rxtx_factor":1,"vcpus":4},{"id":"5","disk":160,"ram":16384,"name":"m1.xlarge","rxtx_factor":1,"vcpus":8},{"id":"bf289de9-bd91-4d12-9d86-a4e8889f412f","disk":200,"ram":2048
```

Via openstack command client
```
[vagrant@localhost echopb]$ openstack image list
+--------------------------------------+------------------------------------------------------+
| ID                                   | Name                                                 |
+--------------------------------------+------------------------------------------------------+
| 4ac53f68-3ac2-49df-8d69-afcea235f8b9 | target_Struts2032-Struts2032-Struts2S2-032-disk1-cl1 |
| c611c095-17a0-4236-9310-438aa9a133e3 | target_Weblogic-Weblogic-Weblogic-disk1-cl1          |
| ab4f5bd8-22a7-4c3a-9e03-1736f612d1f5 | target_editor-editor-disk1-cl1.img-editor-disk1-cl1  |
| 80f1f10b-8128-4c05-a19c-00d05fb6f4e9 | controll_test-kali.img-kali                          |
| 665caf29-72fe-47e9-948b-c9e46cdb07a0 | target66a-Joomla_JFilterInputXSSBypass-Ubuntu-cl1    |
| d085372a-4c45-4c76-bea9-d08a312bcd79 | target_k_009-dedecms_5.5-Ubuntu-cl2                  |
| f99f0481-4011-4ab6-ab53-7e53952179d3 | target_k_003-centos+ecshop2.7.2-CentOS5.8-cl1        |
| 83a93848-2b94-4cde-89b8-271bb9b8bd83 | target86a-phpcmsV9_blindSQL-Ubuntu-cl2               |
| 58f0711d-9d29-4f49-8aab-5a3ab5f0088e | controll048-zdyx-048_SFTP-SFTP                       |
| 7eaf6779-8269-448b-bb7c-4e000ac18d3a | target_k_002-bash_remote-Ubuntu-cl2                  |
| ee8f84b2-c994-4475-9158-3a02e1db5f49 | target_js_022-6-6                                    |
| c425edd1-1a94-4371-bbba-b6e6dcef8ba1 | target_k_001-Adult-Ubuntu-cl3                        |
| 5b715762-8536-4720-896f-bcf61ed3f8f9 | target_js_008-windows2003+appcms-appcms              |
| 30df7065-3658-4415-9226-be10c464aee3 | target_js_027-11-win_11                              |
| 4ba3329f-5e70-439d-8709-2ef691108e59 | target_js_021-5-5                                    |
| bc8ff445-5930-4d98-9684-7c66d47fa444 | target012a-Oracle_11g_11.2.0.1.0-win_zd012           |
| fba67ad9-9791-44ac-87ce-25e4916cda20 | target_k_12-NovelleDirectory8.8.7-linux              |
| 0902ab92-f056-40ee-84e5-1a85317f04c2 | controll102-zdyx-102_tb-tb                           |
| 07a0cfe2-2e14-4f38-8794-364855effe43 | target093-Apachetomcat_remoteexploit-Ubuntu-cl1      |
| dedfe98b-1501-4c06-9c3e-95be40f9cf0b | target_k_010-DedeCMS_5.7-Ubuntu-cl2                  |
| a78cf43c-475e-45d6-a57e-0268ba661ba5 | target38-42a-38-42a-Ubuntu-cl2                       |
| 413d5a4b-49f3-4cf8-a329-66f6f8f70bba | target008a-Splunk5_RCE-win_zd008                     |
| 967bf61b-55b5-4fa2-930c-2a3cf3efa797 | target_k_17-WebTester5.0-Ubuntu-cl3                  |
| 26ef9873-ab8f-40d8-a7c7-e4ffb579daaf | target78a-OpenSSL_Heartbeat-Ubuntu-cl2               |
| bfbdea57-35ed-4bab-b3fb-9472501c623a | target_k_004-centos+Joomla1.5-CentOS5.8-cl1          |
+--------------------------------------+------------------------------------------------------+
```

### Compute flavor service

Test run
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ curl -iL http://127.0.0.1:10001/v1/images
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 25 Jun 2017 08:32:42 GMT
Transfer-Encoding: chunked

{"images":[{"id":"4ac53f68-3ac2-49df-8d69-afcea235f8b9","name":"target_Struts2032-Struts2032-Struts2S2-032-disk1-cl1","status":"ACTIVE","create_at":"2017-06-22T21:05:12Z","updated_at":"2017-06-22T21:05:23Z"},{"id":"c611c095-17a0-4236-9310-438aa9a133e3","name":"target_Weblogic-Weblogic-Weblogic-disk1-cl1","status":"ACTIVE","create_at":"2017-06-22T20:57:24Z","updated_at":"2017-06-22T20:57:56Z"},{"id":"ab4f5bd8-22a7-4c3a-9e03-1736f612d1f5","name":"target_editor-editor-disk1-cl1.img-editor-disk1-cl1","status":"ACTIVE","create_at":"2017-06-22T20:54:56Z","updated_at":"2017-06-22T20:55:06Z"},{"id":"80f1f10b-8128-4c05-a19c-00d05fb6f4e9","name":"controll_test-kali.img-kali","status":"ACTIVE","create_at":"2017-03-10T15:48:37Z","updated_at":"2017-03-10T15:49:55Z"},{"id":"665caf29-72fe-47e9-948b-c9e46cdb07a0","name":"target66a-Joomla_JFilterInputXSSBypass-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T11:47:54Z","updated_at":"2017-02-15T11:49:19Z"},{"id":"d085372a-4c45-4c76-bea9-d08a312bcd79","name":"target_k_009-dedecms_5.5-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T11:45:35Z","updated_at":"2017-02-15T11:47:52Z"},{"id":"f99f0481-4011-4ab6-ab53-7e53952179d3","name":"target_k_003-centos+ecshop2.7.2-CentOS5.8-cl1","status":"ACTIVE","create_at":"2017-02-15T11:44:46Z","updated_at":"2017-02-15T11:45:33Z"},{"id":"83a93848-2b94-4cde-89b8-271bb9b8bd83","name":"target86a-phpcmsV9_blindSQL-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T11:43:20Z","updated_at":"2017-02-15T11:44:44Z"},{"id":"58f0711d-9d29-4f49-8aab-5a3ab5f0088e","name":"controll048-zdyx-048_SFTP-SFTP","status":"ACTIVE","create_at":"2017-02-15T11:40:35Z","updated_at":"2017-02-15T11:43:18Z"},{"id":"7eaf6779-8269-448b-bb7c-4e000ac18d3a","name":"target_k_002-bash_remote-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T11:39:38Z","updated_at":"2017-02-15T11:40:33Z"},{"id":"ee8f84b2-c994-4475-9158-3a02e1db5f49","name":"target_js_022-6-6","status":"ACTIVE","create_at":"2017-02-15T11:39:19Z","updated_at":"2017-02-15T11:39:36Z"},{"id":"c425edd1-1a94-4371-bbba-b6e6dcef8ba1","name":"target_k_001-Adult-Ubuntu-cl3","status":"ACTIVE","create_at":"2017-02-15T11:36:26Z","updated_at":"2017-02-15T11:39:17Z"},{"id":"5b715762-8536-4720-896f-bcf61ed3f8f9","name":"target_js_008-windows2003+appcms-appcms","status":"ACTIVE","create_at":"2017-02-15T11:35:43Z","updated_at":"2017-02-15T11:36:23Z"},{"id":"30df7065-3658-4415-9226-be10c464aee3","name":"target_js_027-11-win_11","status":"ACTIVE","create_at":"2017-02-15T11:34:56Z","updated_at":"2017-02-15T11:35:41Z"},{"id":"4ba3329f-5e70-439d-8709-2ef691108e59","name":"target_js_021-5-5","status":"ACTIVE","create_at":"2017-02-15T11:34:12Z","updated_at":"2017-02-15T11:34:54Z"},{"id":"bc8ff445-5930-4d98-9684-7c66d47fa444","name":"target012a-Oracle_11g_11.2.0.1.0-win_zd012","status":"ACTIVE","create_at":"2017-02-15T11:30:18Z","updated_at":"2017-02-15T11:34:10Z"},{"id":"fba67ad9-9791-44ac-87ce-25e4916cda20","name":"target_k_12-NovelleDirectory8.8.7-linux","status":"ACTIVE","create_at":"2017-02-15T11:29:27Z","updated_at":"2017-02-15T11:30:16Z"},{"id":"0902ab92-f056-40ee-84e5-1a85317f04c2","name":"controll102-zdyx-102_tb-tb","status":"ACTIVE","create_at":"2017-02-15T11:28:11Z","updated_at":"2017-02-15T11:29:25Z"},{"id":"07a0cfe2-2e14-4f38-8794-364855effe43","name":"target093-Apachetomcat_remoteexploit-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T11:26:40Z","updated_at":"2017-02-15T11:28:09Z"},{"id":"dedfe98b-1501-4c06-9c3e-95be40f9cf0b","name":"target_k_010-DedeCMS_5.7-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T11:25:03Z","updated_at":"2017-02-15T11:26:38Z"},{"id":"a78cf43c-475e-45d6-a57e-0268ba661ba5","name":"target38-42a-38-42a-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T11:23:31Z","updated_at":"2017-02-15T11:25:01Z"},{"id":"413d5a4b-49f3-4cf8-a329-66f6f8f70bba","name":"target008a-Splunk5_RCE-win_zd008","status":"ACTIVE","create_at":"2017-02-15T11:22:50Z","updated_at":"2017-02-15T11:23:29Z"},{"id":"967bf61b-55b5-4fa2-930c-2a3cf3efa797","name":"target_k_17-WebTester5.0-Ubuntu-cl3","status":"ACTIVE","create_at":"2017-02-15T11:21:14Z","updated_at":"2017-02-15T11:22:48Z"},{"id":"26ef9873-ab8f-40d8-a7c7-e4ffb579daaf","name":"target78a-OpenSSL_Heartbeat-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T11:19:56Z","updated_at":"2017-02-15T11:21:12Z"},{"id":"bfbdea57-35ed-4bab-b3fb-9472501c623a","name":"target_k_004-centos+Joomla1.5-CentOS5.8-cl1","status":"ACTIVE","create_at":"2017-02-15T11:19:09Z","updated_at":"2017-02-15T11:19:54Z"},{"id":"fbd945a4-9d34-4bd2-b3b3-2e1b5f262e6a","name":"target080-IIS_mljx-IIS_mljx","status":"ACTIVE","create_at":"2017-02-15T11:13:53Z","updated_at":"2017-02-15T11:19:07Z"},{"id":"fdfcf658-ba81-4af5-82b9-e572bff2ecdd","name":"target052a-nginx_dir-Ubuntu-cl3","status":"ACTIVE","create_at":"2017-02-15T11:13:06Z","updated_at":"2017-02-15T11:13:51Z"},{"id":"b8e1efb7-e9b4-4283-828b-a40d348bd395","name":"target082-coldfusion_passlogin-coldfusion","status":"ACTIVE","create_at":"2017-02-15T11:08:26Z","updated_at":"2017-02-15T11:13:04Z"},{"id":"34f4c0e6-276a-4c14-b800-eb77754e5a77","name":"target049-zdyx-049_ftpfu-ftpfu","status":"ACTIVE","create_at":"2017-02-15T11:06:15Z","updated_at":"2017-02-15T11:08:25Z"},{"id":"5ed2fda8-e595-46d3-95dc-ca5ccd9d6d04","name":"target88a-Drupal_7.x_path-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T11:05:18Z","updated_at":"2017-02-15T11:06:13Z"},{"id":"fd7cf08f-7576-40ae-8ac6-18b251f15d5b","name":"target_k_21-Trojan_horse_back_tracking-Ubuntu","status":"ACTIVE","create_at":"2017-02-15T11:03:48Z","updated_at":"2017-02-15T11:05:16Z"},{"id":"230cc6fa-0dc4-4809-837e-8ee5faae755d","name":"target018a-haneWIN1.5.3_Dos-win_zd018","status":"ACTIVE","create_at":"2017-02-15T11:01:50Z","updated_at":"2017-02-15T11:03:46Z"},{"id":"0a8ee430-5a65-46a5-9ecc-e1b9a05081d3","name":"target_js_025-9-9","status":"ACTIVE","create_at":"2017-02-15T11:01:07Z","updated_at":"2017-02-15T11:01:48Z"},{"id":"4586bbd9-7f90-42aa-b1bf-0562446279c5","name":"target024a-apache_roller-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T10:59:48Z","updated_at":"2017-02-15T11:01:05Z"},{"id":"97f82e86-36b0-488e-9717-c3a75660abad","name":"target65a-JoomlaXClonerComponent_RCE-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T10:58:28Z","updated_at":"2017-02-15T10:59:46Z"},{"id":"09e60af0-9a5b-4699-81f3-e5c3a6d6167b","name":"controll_js_kali-kalijingsai.img-kalijingsai","status":"ACTIVE","create_at":"2017-02-15T10:55:43Z","updated_at":"2017-02-15T10:58:26Z"},{"id":"c29eaac5-a579-479b-9a54-4f4011885b57","name":"target109-zdyx-109_Oracle01-Oracle1","status":"ACTIVE","create_at":"2017-02-15T10:52:01Z","updated_at":"2017-02-15T10:55:41Z"},{"id":"d517e884-7ec2-4c6a-8674-0cfec839846a","name":"target_js_024-8-8","status":"ACTIVE","create_at":"2017-02-15T10:51:40Z","updated_at":"2017-02-15T10:51:57Z"},{"id":"a3071cf5-d318-45ca-a408-15292d1af1f7","name":"target_js_033-17-win_17","status":"ACTIVE","create_at":"2017-02-15T10:50:57Z","updated_at":"2017-02-15T10:51:38Z"},{"id":"5f0c0a0f-e581-40aa-b9cf-a22d119ca71c","name":"target004a-eDirectory8.8.7-linux","status":"ACTIVE","create_at":"2017-02-15T10:49:44Z","updated_at":"2017-02-15T10:50:55Z"},{"id":"412e9764-f68c-4b73-b3d0-9beaaf3d31ba","name":"target009a-DedeCMS5.7_SQLi-recommend.php-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T10:48:23Z","updated_at":"2017-02-15T10:49:43Z"},{"id":"9cad9adc-182a-400c-b84b-33f10c2542c5","name":"target091-Apache_server_status-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T10:46:55Z","updated_at":"2017-02-15T10:48:21Z"},{"id":"2bad1744-0e17-402b-aded-5f70f965f182","name":"controll_2003-2003-win_con_2003","status":"ACTIVE","create_at":"2017-02-15T10:45:59Z","updated_at":"2017-02-15T10:46:54Z"},{"id":"d020ac78-5a70-4387-b798-cd2d1301baab","name":"target81a-Apache_houzhui-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T10:44:57Z","updated_at":"2017-02-15T10:45:57Z"},{"id":"a5274557-af78-4f8c-9678-9f8f6896a893","name":"target_js_018-Discuz_X3.1-discuz","status":"ACTIVE","create_at":"2017-02-15T10:39:32Z","updated_at":"2017-02-15T10:44:53Z"},{"id":"8b5ce896-7ada-4041-a6e6-a2158701519a","name":"target021a-deepofix_bypass-DeepOfix_Mail_Server-cl1","status":"ACTIVE","create_at":"2017-02-15T10:38:56Z","updated_at":"2017-02-15T10:39:31Z"},{"id":"09c4c4e0-8da9-44bb-abf2-48daa3e8461e","name":"target_k_14-NginxBlankNullByte_RCE-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T10:37:15Z","updated_at":"2017-02-15T10:38:54Z"},{"id":"12cc0a8e-0d32-408a-b232-866feee85058","name":"target64a-Joomlacom_collectorComponent_ArbitraryFileUploadVulnerability-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T10:35:54Z","updated_at":"2017-02-15T10:37:13Z"},{"id":"2e721c48-46b3-443c-bb20-c27e7db499a8","name":"target_js_007-windows2003+ckfinder-ckfinder","status":"ACTIVE","create_at":"2017-02-15T10:35:16Z","updated_at":"2017-02-15T10:35:52Z"},{"id":"8d2f3378-e3e1-4351-818b-be4a5047e489","name":"target_js_020-3-3","status":"ACTIVE","create_at":"2017-02-15T10:34:31Z","updated_at":"2017-02-15T10:35:14Z"},{"id":"4ea8854d-0c33-442e-af83-86c8c2a990c2","name":"target_js_001-anquangou-sqli","status":"ACTIVE","create_at":"2017-02-15T10:33:03Z","updated_at":"2017-02-15T10:34:29Z"},{"id":"1ff9cca3-4b9f-46c3-a3b4-b1b62bb543a7","name":"target_js_012-MysqlMOF_tiquan-mysql","status":"ACTIVE","create_at":"2017-02-15T10:27:16Z","updated_at":"2017-02-15T10:33:01Z"},{"id":"1718b5f9-87de-472f-a582-bbe1c393d81b","name":"target050-TFS-TFS","status":"ACTIVE","create_at":"2017-02-15T10:25:04Z","updated_at":"2017-02-15T10:27:15Z"},{"id":"f123ece6-277f-4fa8-8a97-d45b56c25e18","name":"target_js_013-office_rar_pojie-office","status":"ACTIVE","create_at":"2017-02-15T10:23:29Z","updated_at":"2017-02-15T10:25:02Z"},{"id":"e27ce6af-7998-4623-8e6b-6f06a6761e08","name":"target63a-JoomlaModuleSimpleFileUpload_RCE-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T10:22:31Z","updated_at":"2017-02-15T10:23:26Z"},{"id":"920bb4e7-3dee-4998-a042-66f50ed0ea64","name":"target048-zdxy_048_SFTP-SFTP","status":"ACTIVE","create_at":"2017-02-15T10:20:19Z","updated_at":"2017-02-15T10:22:29Z"},{"id":"548840f7-5291-4431-b05b-ccae8ee71eff","name":"target_js_006-windows2003+ecshop-ecshop","status":"ACTIVE","create_at":"2017-02-15T10:19:34Z","updated_at":"2017-02-15T10:20:17Z"},{"id":"76ed0878-59b1-40a7-96a2-e6d0cf41ef2d","name":"target002b-Bind9_DNS-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T10:18:16Z","updated_at":"2017-02-15T10:19:32Z"},{"id":"b1145176-6cc5-4305-8639-78890a4b1de1","name":"target_js_015-HFS_2.3.X_remote-HFS","status":"ACTIVE","create_at":"2017-02-15T10:17:36Z","updated_at":"2017-02-15T10:18:14Z"},{"id":"73d6eb1c-3f15-44ec-9a62-adfaf887a75e","name":"target_js_028-12-win_12","status":"ACTIVE","create_at":"2017-02-15T10:16:32Z","updated_at":"2017-02-15T10:17:34Z"},{"id":"d58daaa5-eb28-4144-9a02-14b80a5ca06e","name":"target69a-TechfolioJoomlaComponent_SQLInjectionVulnerability-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T10:15:41Z","updated_at":"2017-02-15T10:16:31Z"},{"id":"59be2f70-1e17-4d17-a63e-6c942f6b697e","name":"target_js_014-IIS+FCKeditor_FileUpload-iisfck","status":"ACTIVE","create_at":"2017-02-15T10:14:57Z","updated_at":"2017-02-15T10:15:39Z"},{"id":"a51c69a8-12b9-4363-8cf0-389c9ae065fe","name":"controll_win7-win7-win7","status":"ACTIVE","create_at":"2017-02-15T10:06:54Z","updated_at":"2017-02-15T10:14:54Z"},{"id":"54c98761-3391-4977-878a-b996a1f357e3","name":"target_js_003-windows2003+PHP168V6.02-php","status":"ACTIVE","create_at":"2017-02-15T10:06:10Z","updated_at":"2017-02-15T10:06:53Z"},{"id":"9520f5f0-123b-42d6-b949-d1433f9217c3","name":"target046-zdyx-046_qfc-qfc","status":"ACTIVE","create_at":"2017-02-15T10:04:10Z","updated_at":"2017-02-15T10:06:08Z"},{"id":"2998a8ff-23b5-4f13-83af-6813b531a885","name":"target83a-phpmyadmin-Ubuntu","status":"ACTIVE","create_at":"2017-02-15T10:02:16Z","updated_at":"2017-02-15T10:04:08Z"},{"id":"21422e57-f26a-46eb-b090-b82a9dc5c512","name":"target_js_004-windows2003+lpk-winser03","status":"ACTIVE","create_at":"2017-02-15T10:02:05Z","updated_at":"2017-02-15T10:02:14Z"},{"id":"3b57fef0-e73b-4ca9-b72a-fb547fcabc4a","name":"target_js_005-windows2003+iis6-iis6","status":"ACTIVE","create_at":"2017-02-15T10:01:24Z","updated_at":"2017-02-15T10:02:03Z"},{"id":"bf91a960-53f3-4e78-9535-0e220fd8b85c","name":"target_k_18-WordPress2.8.5-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T09:59:11Z","updated_at":"2017-02-15T10:01:23Z"},{"id":"2d0d68d1-1cc9-4c2e-9cc0-a97534b56934","name":"target_k_19-phpcmsv9_LFI-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T09:56:59Z","updated_at":"2017-02-15T09:59:09Z"},{"id":"181f7295-cb1d-42fd-8f4c-033de37635a6","name":"controll222-leichi.img-leichi","status":"ACTIVE","create_at":"2017-02-15T09:56:05Z","updated_at":"2017-02-15T09:56:57Z"},{"id":"32ed7ff8-e56e-4188-a746-84f047f5ad86","name":"controll105-zdyx_105-chb","status":"ACTIVE","create_at":"2017-02-15T09:53:59Z","updated_at":"2017-02-15T09:56:03Z"},{"id":"80598dc7-a7a7-45a5-8d43-3afc38700e3f","name":"target_js_023-7-7","status":"ACTIVE","create_at":"2017-02-15T09:53:14Z","updated_at":"2017-02-15T09:53:58Z"},{"id":"13d09989-d009-468e-9c95-d445c2a42f51","name":"target_js_011-oracle_input-oracle","status":"ACTIVE","create_at":"2017-02-15T09:49:50Z","updated_at":"2017-02-15T09:53:12Z"},{"id":"8f8f1d38-f222-47d7-9b44-d5e6df86df78","name":"target_js_034-18-18","status":"ACTIVE","create_at":"2017-02-15T09:49:03Z","updated_at":"2017-02-15T09:49:48Z"},{"id":"8948e77b-b52a-4804-8c70-7f2c3223490b","name":"target_k_005-centos+PHPWind7.5-CentOS5.8-cl1","status":"ACTIVE","create_at":"2017-02-15T09:48:19Z","updated_at":"2017-02-15T09:49:02Z"},{"id":"eebf0a8e-1415-4920-85f5-d5328033ea45","name":"target118a-xss_raoguo-win_zd118","status":"ACTIVE","create_at":"2017-02-15T09:47:37Z","updated_at":"2017-02-15T09:48:17Z"},{"id":"a416f5ac-c2a5-4606-88c3-ed82ebf34c51","name":"target68a-TNREnhancedJoomlaSearch-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T09:46:14Z","updated_at":"2017-02-15T09:47:36Z"},{"id":"ea88d594-9eda-45ee-a4ee-00f5d306100e","name":"target_k_20-Hackgame-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T09:43:41Z","updated_at":"2017-02-15T09:46:12Z"},{"id":"cb89829e-d5c0-47f2-bbc7-d3e0d1400fb5","name":"target_js_010-securitymanager_SQLi-sqli","status":"ACTIVE","create_at":"2017-02-15T09:43:02Z","updated_at":"2017-02-15T09:43:39Z"},{"id":"37ea887e-820f-4686-bf4f-db83ab59f924","name":"target_k_008-centos5+mantis-CentOS-cl1","status":"ACTIVE","create_at":"2017-02-15T09:42:52Z","updated_at":"2017-02-15T09:43:00Z"},{"id":"12c07753-29a2-4148-ae0d-e6980f259f08","name":"target_js_035-Splunk5_RCE-splunk","status":"ACTIVE","create_at":"2017-02-15T09:42:08Z","updated_at":"2017-02-15T09:42:50Z"},{"id":"2157b3b8-6af5-4aba-b626-52194be0ab99","name":"target_k_007-centos5+dz7.2-CentOS-cl1","status":"ACTIVE","create_at":"2017-02-15T09:41:59Z","updated_at":"2017-02-15T09:42:06Z"},{"id":"eb686065-b141-453e-8481-f163c8a39450","name":"target_js_026-10-win_10","status":"ACTIVE","create_at":"2017-02-15T09:41:17Z","updated_at":"2017-02-15T09:41:57Z"},{"id":"0b2822bf-f71e-4b0f-939c-edf608202d2b","name":"target006a-win_zd006.img-win_zd006","status":"ACTIVE","create_at":"2017-02-15T09:40:36Z","updated_at":"2017-02-15T09:41:16Z"},{"id":"da893adb-11d2-4e3d-9b2d-cd2746dae8fc","name":"controll109-zdyx-109_Oracle01-Oracle1","status":"ACTIVE","create_at":"2017-02-15T09:36:57Z","updated_at":"2017-02-15T09:40:34Z"},{"id":"a423577d-27d3-42af-b930-b72b20f53152","name":"target108-zdyx-108_Oracle02-Oracle2","status":"ACTIVE","create_at":"2017-02-15T09:33:38Z","updated_at":"2017-02-15T09:36:55Z"},{"id":"8bfe2eff-b841-4c84-83ed-1eb27ed2c1c5","name":"controll_kali-kali.img-kali","status":"ACTIVE","create_at":"2017-02-15T09:31:00Z","updated_at":"2017-02-15T09:33:37Z"},{"id":"fb2b5f00-4c0e-4090-a0e7-0f05b2b44a63","name":"target_k_11-dz25remote-CentOS-cl1","status":"ACTIVE","create_at":"2017-02-15T09:30:49Z","updated_at":"2017-02-15T09:30:58Z"},{"id":"30267564-5775-4208-a752-582ad0fcf050","name":"controll106,107-MiddleAttack-MiddleAttack","status":"ACTIVE","create_at":"2017-02-15T09:30:08Z","updated_at":"2017-02-15T09:30:47Z"},{"id":"93b77de2-0ac3-4d14-b3c0-5818dfed5064","name":"target019a-haneWIN_SEH-win_zd019","status":"ACTIVE","create_at":"2017-02-15T09:28:15Z","updated_at":"2017-02-15T09:30:06Z"},{"id":"e0b5dcc7-484d-4ee0-8944-79603f9065ee","name":"target100-zdyx-100_officewjjm-ofjm","status":"ACTIVE","create_at":"2017-02-15T09:26:59Z","updated_at":"2017-02-15T09:28:14Z"},{"id":"a9f6a51d-3154-4657-a747-c5e3a960c75b","name":"target_js_009-shopex4.8.5-shopex","status":"ACTIVE","create_at":"2017-02-15T09:26:12Z","updated_at":"2017-02-15T09:26:57Z"},{"id":"a14faded-d813-4dbc-8c51-ff147444ec09","name":"controll108-zdyx-108_Oracle02-Oracle2","status":"ACTIVE","create_at":"2017-02-15T09:22:36Z","updated_at":"2017-02-15T09:26:11Z"},{"id":"c02e6575-3ed4-4de2-a9ac-5cbed2a6290f","name":"target200-sqljoin-Metasploitable","status":"ACTIVE","create_at":"2017-02-15T09:22:26Z","updated_at":"2017-02-15T09:22:34Z"},{"id":"ff4795e5-6aaa-4ce8-859f-c5bedd6cf2e0","name":"target_js_002-windows2003+PHPCMS9.5.6_getwebshell-phpcms","status":"ACTIVE","create_at":"2017-02-15T09:21:44Z","updated_at":"2017-02-15T09:22:25Z"},{"id":"ca01c769-3cff-4a5e-beea-3dc6e090d05f","name":"target002c-Bind9_DNS-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T09:19:49Z","updated_at":"2017-02-15T09:21:43Z"},{"id":"8987af61-f36b-4348-960f-ab8dbb06788e","name":"target_js_016-FCKeditor-fckeditor","status":"ACTIVE","create_at":"2017-02-15T09:18:15Z","updated_at":"2017-02-15T09:19:48Z"},{"id":"cdc8e055-7367-4bcd-9a76-695cf7408698","name":"target095-IIS_hzjx-hzjx","status":"ACTIVE","create_at":"2017-02-15T09:13:25Z","updated_at":"2017-02-15T09:18:14Z"},{"id":"8ac027f9-cdcb-4de5-92ba-2955e0145bfb","name":"controll_z_ubuntu-ubuntu-ubuntu","status":"ACTIVE","create_at":"2017-02-15T09:11:00Z","updated_at":"2017-02-15T09:13:23Z"},{"id":"80ca9c58-e3cd-4f73-b185-5ef0905bf7d8","name":"target_js_032-16-16","status":"ACTIVE","create_at":"2017-02-15T09:10:08Z","updated_at":"2017-02-15T09:10:58Z"},{"id":"71eb9d26-ab04-4744-817b-ce96575ef83c","name":"target_js_030-14-win_14","status":"ACTIVE","create_at":"2017-02-15T09:08:54Z","updated_at":"2017-02-15T09:10:06Z"},{"id":"86b5ef33-c5c5-4257-a10f-f1db6b270a2c","name":"controll_yinxie-yinxie-centos","status":"ACTIVE","create_at":"2017-02-15T09:07:44Z","updated_at":"2017-02-15T09:08:52Z"},{"id":"400e6473-af5b-4170-842d-5a2bba9de209","name":"target_k_16-tomcat_RCE-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T09:05:03Z","updated_at":"2017-02-15T09:07:43Z"},{"id":"0e71f3e9-9c7a-4547-b9d7-d5bed2e1bc32","name":"target77a-dedecms_plussql-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T09:03:24Z","updated_at":"2017-02-15T09:05:01Z"},{"id":"f2749e75-17cc-49d1-8236-198d2a1daa91","name":"target87a-phpcmsV9.3_include-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T09:01:01Z","updated_at":"2017-02-15T09:03:23Z"},{"id":"edb3b98c-9f00-48b9-8de8-4334f619c042","name":"target053a-JBoss_RemoteCodeExecution-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T08:59:25Z","updated_at":"2017-02-15T09:01:00Z"},{"id":"a978e83e-4007-4c22-a100-20539a8903d9","name":"target067-Joomla-Ubuntu","status":"ACTIVE","create_at":"2017-02-15T08:57:53Z","updated_at":"2017-02-15T08:59:24Z"},{"id":"9621364f-59fd-4c1a-83e9-4250c3ed433f","name":"target_js_017-Ewebeditor-editor","status":"ACTIVE","create_at":"2017-02-15T08:56:02Z","updated_at":"2017-02-15T08:57:52Z"},{"id":"a0bd9bae-a1e1-42ec-9492-e494def9183a","name":"controll_tomcat-tomcat-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T08:54:11Z","updated_at":"2017-02-15T08:56:01Z"},{"id":"a056d712-5d6f-4df4-b137-90763d727c6d","name":"target003a-U-mail_XSS-win_2003","status":"ACTIVE","create_at":"2017-02-15T08:53:14Z","updated_at":"2017-02-15T08:54:10Z"},{"id":"d26751ce-7780-4c3f-8ae5-43c271b37040","name":"target026a-MoinMoin_ArbitraryCommandExecution-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T08:51:39Z","updated_at":"2017-02-15T08:53:13Z"},{"id":"b85852d9-607b-47de-8cac-4fab64a31103","name":"controll101-zdyx-101_rarjm-jm","status":"ACTIVE","create_at":"2017-02-15T08:50:50Z","updated_at":"2017-02-15T08:51:38Z"},{"id":"8fb8e8e0-7771-4b16-9c9b-7963a9bc76ae","name":"target_k_006-centos5+dz7.1-CentOS-cl1","status":"ACTIVE","create_at":"2017-02-15T08:50:02Z","updated_at":"2017-02-15T08:50:49Z"},{"id":"47330c30-45b2-4f33-a677-43bd57cc4ed6","name":"target084-phpcmsV9_file_read-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T08:48:04Z","updated_at":"2017-02-15T08:50:01Z"},{"id":"cb7e6ea6-8713-4c5c-bf6c-6b98af0764cf","name":"target002a-Bind9_DNS-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T08:46:32Z","updated_at":"2017-02-15T08:48:03Z"},{"id":"c41ed23c-c04d-413f-98ae-b9a1357188a5","name":"target_k_13-lighttpdSQLinjection-Ubuntu-cl3","status":"ACTIVE","create_at":"2017-02-15T08:44:43Z","updated_at":"2017-02-15T08:46:31Z"},{"id":"fd99d697-87f3-40a0-935a-abd8dad58a54","name":"target_js_031-15-15","status":"ACTIVE","create_at":"2017-02-15T08:42:11Z","updated_at":"2017-02-15T08:44:42Z"},{"id":"5bccfa06-53f2-4ae9-8920-af4bf84e52af","name":"controll_js_03-win03.img-win03","status":"ACTIVE","create_at":"2017-02-15T08:41:21Z","updated_at":"2017-02-15T08:42:10Z"},{"id":"01df6e60-535b-48af-9d69-1d25ec8ebabd","name":"target023a-zdyx-023_NJStar-win_zd23","status":"ACTIVE","create_at":"2017-02-15T08:39:02Z","updated_at":"2017-02-15T08:41:19Z"},{"id":"94651a03-6bd5-402b-890f-97f0624fed4f","name":"target27-30,35-37a-27-30,35-37-Ubuntu-cl1","status":"ACTIVE","create_at":"2017-02-15T08:36:50Z","updated_at":"2017-02-15T08:39:01Z"},{"id":"c67f74d6-70fe-48ce-b3fe-0618516c4a78","name":"target_k_15-OpenWebAnalytics1.5.4-Ubuntu-cl2","status":"ACTIVE","create_at":"2017-02-15T08:35:00Z","updated_at":"2017-02-15T08:36:48Z"},{"id":"7eb3b9e1-03de-4431-9c36-878787fbc77d","name":"target_js_029-13-13","status":"ACTIVE","create_at":"2017-02-15T08:34:13Z","updated_at":"2017-02-15T08:34:59Z"},{"id":"12c871d8-8a13-4684-90ea-9e600975b275","name":"controll112-word.img-word","status":"ACTIVE","create_at":"2017-02-15T08:32:43Z","updated_at":"2017-02-15T08:34:12Z"},{"id":"31181ed2-f9e4-4c2c-a588-3286731a7468","name":"target015a-Struts2.0_RCE(S2-016)-Ubuntu","status":"ACTIVE","create_at":"2017-02-15T08:31:11Z","updated_at":"2017-02-15T08:32:42Z"},{"id":"cbaf5e9f-83b4-4716-ba49-1216cff125b5","name":"controll100-zdyx-100_officewjjm-wjjm","status":"ACTIVE","create_at":"2017-02-15T08:29:40Z","updated_at":"2017-02-15T08:31:09Z"},{"id":"4a218637-8a6c-47dd-a0c8-ea333efad9fc","name":"target_js_019-2-2","status":"ACTIVE","create_at":"2017-02-15T08:29:27Z","updated_at":"2017-02-15T08:29:39Z"}]}
```

Via openstack command client
```
[vagrant@localhost echopb]$ openstack flavor list
+--------------------------------------+-----------+-------+------+-----------+-------+-----------+
| ID                                   | Name      |   RAM | Disk | Ephemeral | VCPUs | Is Public |
+--------------------------------------+-----------+-------+------+-----------+-------+-----------+
| 1                                    | m1.tiny   |   512 |    1 |         0 |     1 | True      |
| 2                                    | m1.small  |  2048 |   20 |         0 |     1 | True      |
| 3                                    | m1.medium |  4096 |   40 |         0 |     2 | True      |
| 4                                    | m1.large  |  8192 |   80 |         0 |     4 | True      |
| 5                                    | m1.xlarge | 16384 |  160 |         0 |     8 | True      |
| bf289de9-bd91-4d12-9d86-a4e8889f412f | large     |  2048 |  200 |         0 |     1 | True      |
+--------------------------------------+-----------+-------+------+-----------+-------+-----------+
```

## Release 0.1

* bulk create VMs

### Computing

Test JSON
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

Test run
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ kopos test boot
response Status: 200 OK
response Headers: map[Content-Type:[application/json] Date:[Fri, 23 Jun 2017 10:07:24 GMT]]
response Body: {"flavor_id":"bf289de9-bd91-4d12-9d86-a4e8889f412f","flavor_name":"large","image_id":"83a93848-2b94-4cde-89b8-271bb9b8bd83","image_name":"target86a-phpcmsV9_blindSQL-Ubuntu-cl2","min_count":4,"max_count":4,"network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","network_name":"private","floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_network_name":"private","ports":[{"id":"806194d4-f561-43bc-8e25-ce1d1987d9f8","network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","name":"............-0","status":"DOWN","mac_address":"fa:16:3e:ba:86:f9","fixed_ips":[{"subnet_id":"931c8e3a-47df-42a6-aecd-9d81789b5fb7","ip_address":"192.168.0.5"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["4da19416-41cb-469b-b5fb-b02d24aaff47"]},{"id":"16c2bd2a-fdaa-42ba-b854-4508ba486d94","network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","name":"............-1","status":"DOWN","mac_address":"fa:16:3e:50:8c:82","fixed_ips":[{"subnet_id":"931c8e3a-47df-42a6-aecd-9d81789b5fb7","ip_address":"192.168.0.6"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["4da19416-41cb-469b-b5fb-b02d24aaff47"]},{"id":"5f022746-4998-4c0a-8291-4acae6c11e04","network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","name":"............-2","status":"DOWN","mac_address":"fa:16:3e:2a:d0:7e","fixed_ips":[{"subnet_id":"931c8e3a-47df-42a6-aecd-9d81789b5fb7","ip_address":"192.168.0.7"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["4da19416-41cb-469b-b5fb-b02d24aaff47"]},{"id":"036a26cb-8647-4522-8538-abfa7068bdf8","network_id":"830548b8-c7fc-435e-b144-b81f29b1e312","name":"............-3","status":"DOWN","mac_address":"fa:16:3e:ab:fe:26","fixed_ips":[{"subnet_id":"931c8e3a-47df-42a6-aecd-9d81789b5fb7","ip_address":"192.168.0.8"}],"tenant_id":"a2a01453f7ed456a8d0d270ed5207697","security_groups":["4da19416-41cb-469b-b5fb-b02d24aaff47"]}],"servers":[{"id":"82b0066a-3c96-4854-a466-f1dd1ea1569f","updated":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","adminPass":"m9JhekPZgcUR"},{"id":"ed2c2068-c717-4b09-9652-50ece6b153a4","updated":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","adminPass":"XeEJzyr9ZenJ"},{"id":"25436d62-49ee-4e32-aedc-fbab11d7cec2","updated":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","adminPass":"Bvu7dgFZD6yU"},{"id":"f93382f2-8447-424a-9f1e-8061c3b6eb71","updated":"0001-01-01T00:00:00Z","created":"0001-01-01T00:00:00Z","adminPass":"ML5aqubvQDRp"}],"floating_ips":[{"floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_ip_address":"10.100.151.106","port_id":"806194d4-f561-43bc-8e25-ce1d1987d9f8","fixed_ip_address":"192.168.0.5","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_ip_address":"10.100.151.107","port_id":"16c2bd2a-fdaa-42ba-b854-4508ba486d94","fixed_ip_address":"192.168.0.6","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_ip_address":"10.100.151.108","port_id":"5f022746-4998-4c0a-8291-4acae6c11e04","fixed_ip_address":"192.168.0.7","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"},{"floating_network_id":"92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d","floating_ip_address":"10.100.151.109","port_id":"036a26cb-8647-4522-8538-abfa7068bdf8","fixed_ip_address":"192.168.0.8","tenant_id":"a2a01453f7ed456a8d0d270ed5207697"}],"port_server_pairs":{"036a26cb-8647-4522-8538-abfa7068bdf8":"f93382f2-8447-424a-9f1e-8061c3b6eb71","16c2bd2a-fdaa-42ba-b854-4508ba486d94":"ed2c2068-c717-4b09-9652-50ece6b153a4","5f022746-4998-4c0a-8291-4acae6c11e04":"25436d62-49ee-4e32-aedc-fbab11d7cec2","806194d4-f561-43bc-8e25-ce1d1987d9f8":"82b0066a-3c96-4854-a466-f1dd1ea1569f"}}
```

Via openstack command client
```
[vagrant@localhost echopb]$ openstack server list
+--------------------------------------+-------------------------------------------------------+--------+------------------------+
| ID                                   | Name                                                  | Status | Networks               |
+--------------------------------------+-------------------------------------------------------+--------+------------------------+
| f93382f2-8447-424a-9f1e-8061c3b6eb71 | ............-3                                        | BUILD  |                        |
| 25436d62-49ee-4e32-aedc-fbab11d7cec2 | ............-2                                        | BUILD  |                        |
| ed2c2068-c717-4b09-9652-50ece6b153a4 | ............-1                                        | BUILD  |                        |
| 82b0066a-3c96-4854-a466-f1dd1ea1569f | ............-0                                        | BUILD  |                        |
| c0af9fef-6e50-41d8-8d90-cec121cb1537 | jkk                                                   | ACTIVE | private_net=10.20.30.9 |
| b50ac0ed-3e9a-4b75-8988-3bda581096f0 | gfj                                                   | ACTIVE | private_net=10.20.30.8 |
| fd61f0c7-0a45-4f33-9ae7-dc1b7ad4a81c | sefef                                                 | ACTIVE | private_net=10.20.30.7 |
| 48abf3ca-1f8a-499b-a65f-4a3570f0c5db | ct-_ctf-a8db1d82db78ed452ba0882fb9554fc9-vnn2vpygyszt | ACTIVE | private_net=10.20.30.2 |
+--------------------------------------+-------------------------------------------------------+--------+------------------------+
```
