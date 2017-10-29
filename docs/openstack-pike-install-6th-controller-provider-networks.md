# OpenStack Pike Installation

## Table of content

控制节点
* [网络配置](#network)
* [SSH客户端排错](#ssh-trouble-shooting)
* [防火墙](#firewall)
* [NTP服务](#chrony)
* [Openstack Pike版本YUM仓库](#openstack-Repository)
* [Openstack 命令工具包](#openstack-client)
* [MariaDB数据库](#database)
* [RabbitMQ消息队列](#queue)
* [Memcached缓存](#cache)

## Controller Node

### Network

Config _eth2_ as promicuous mode
```
[vagrant@localhost ~]$ sudo nmcli con 
NAME                UUID                                  TYPE            DEVICE 
System eth0         5fb06bd0-0bb0-7ffb-45f1-d6edd65f3e03  802-3-ethernet  eth0   
Wired connection 1  3c55af49-6222-3b6b-b91d-eb1b82b6005e  802-3-ethernet  eth1   
Wired connection 2  40ae81b3-2e97-30d3-b588-aecb3e142583  802-3-ethernet  eth2   
[vagrant@localhost ~]$ sudo nmcli con delete 40ae81b3-2e97-30d3-b588-aecb3e142583
Connection 'Wired connection 2' (40ae81b3-2e97-30d3-b588-aecb3e142583) successfully deleted.
[vagrant@localhost ~]$ sudo nmcli con reload
[vagrant@localhost ~]$ sudo nmcli con 
NAME                UUID                                  TYPE            DEVICE 
System eth0         5fb06bd0-0bb0-7ffb-45f1-d6edd65f3e03  802-3-ethernet  eth0   
Wired connection 1  3c55af49-6222-3b6b-b91d-eb1b82b6005e  802-3-ethernet  eth1   
```

```
[vagrant@localhost ~]$ ls /etc/sysconfig/network-scripts/ifcfg-*
/etc/sysconfig/network-scripts/ifcfg-eth0  /etc/sysconfig/network-scripts/ifcfg-lo  /etc/sysconfig/network-scripts/ifcfg-Wired_connection_1
```

```
[vagrant@localhost ~]$ cat <<EOF > ifcfg-eth2
> DEVICE=eth2
> TYPE=Ethernet
> ONBOOT="yes"
> BOOTPROTO="none"
> EOF
```

```
[vagrant@localhost ~]$ sudo cp ifcfg-eth2 /etc/sysconfig/network-scripts/
```

Last 
```
[vagrant@localhost ~]$ ip a show eth2
4: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 08:00:27:f6:b8:b3 brd ff:ff:ff:ff:ff:ff
```

### Service Registration

[Neutron](https://docs.openstack.org/neutron/pike/install/controller-install-rdo.html)

Service user
```
[vagrant@localhost ~]$ openstack user create --domain default --password SERVICE_PASS neutron
+---------------------+----------------------------------+
| Field               | Value                            |
+---------------------+----------------------------------+
| domain_id           | default                          |
| enabled             | True                             |
| id                  | a2a3e896c86342198f7297d77edbcb4f |
| name                | neutron                          |
| options             | {}                               |
| password_expires_at | None                             |
+---------------------+----------------------------------+
```

Service role
```
[vagrant@localhost ~]$ openstack role add --project service --user neutron admin
```

Service name
```
[vagrant@localhost ~]$ openstack service create --name neutron --description "OpenStack Networking" network
+-------------+----------------------------------+
| Field       | Value                            |
+-------------+----------------------------------+
| description | OpenStack Networking             |
| enabled     | True                             |
| id          | 2c6d504fc9344ebf98d34506f60882c2 |
| name        | neutron                          |
| type        | network                          |
+-------------+----------------------------------+
```

Endpoint public
```
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne network public http://10.64.33.64:9696
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | 2911e9da6e5b4c798acd0c28d886c605 |
| interface    | public                           |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 2c6d504fc9344ebf98d34506f60882c2 |
| service_name | neutron                          |
| service_type | network                          |
| url          | http://10.64.33.64:9696          |
+--------------+----------------------------------+
```

Endpoint internal
```
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne network internal http://10.64.33.64:9696
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | 83ee3b84083642eca1fd98ecd6a04104 |
| interface    | internal                         |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 2c6d504fc9344ebf98d34506f60882c2 |
| service_name | neutron                          |
| service_type | network                          |
| url          | http://10.64.33.64:9696          |
+--------------+----------------------------------+
```

Endpoint admin
```
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne network admin http://10.64.33.64:9696
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | 5467907e39284043a8ab13403ce02f8f |
| interface    | admin                            |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 2c6d504fc9344ebf98d34506f60882c2 |
| service_name | neutron                          |
| service_type | network                          |
| url          | http://10.64.33.64:9696          |
+--------------+----------------------------------+
```

### Database

MariaDB
```
[vagrant@localhost ~]$ mysql -u root -e "CREATE DATABASE neutron;GRANT ALL PRIVILEGES ON neutron.* TO 'neutron'@'localhost' IDENTIFIED BY 'SERVICE_DBPASS';GRANT ALL PRIVILEGES ON neutron.* TO 'neutron'@'%' IDENTIFIED BY 'SERVICE_DBPASS';"
```

```
[vagrant@localhost ~]$ mysql -u root -e "show databases;"
+--------------------+
| Database           |
+--------------------+
| glance             |
| information_schema |
| keystone           |
| mysql              |
| neutron            |
| nova               |
| nova_api           |
| nova_cell0         |
| performance_schema |
| test               |
+--------------------+
```

### Provider networks

Install [Provider Networks](https://docs.openstack.org/neutron/pike/install/controller-install-option1-rdo.html)
```
[vagrant@localhost ~]$ sudo yum install -y openstack-neutron openstack-neutron-ml2 openstack-neutron-linuxbridge ebtables
Loaded plugins: fastestmirror
base                                                                                                                       | 3.6 kB  00:00:00     
centos-ceph-jewel                                                                                                          | 2.9 kB  00:00:00     
centos-openstack-pike                                                                                                      | 2.9 kB  00:00:00     
centos-qemu-ev                                                                                                             | 2.9 kB  00:00:00     
extras                                                                                                                     | 3.4 kB  00:00:00     
updates                                                                                                                    | 3.4 kB  00:00:00     
updates/7/x86_64/primary_db                                                                                                | 2.9 MB  00:00:03     
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.btte.net
 * updates: mirrors.btte.net
Package ebtables-2.0.10-15.el7.x86_64 already installed and latest version
Resolving Dependencies
--> Running transaction check
---> Package openstack-neutron.noarch 1:11.0.1-1.el7 will be installed
--> Processing Dependency: openstack-neutron-common = 1:11.0.1-1.el7 for package: 1:openstack-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: haproxy >= 1.5.0 for package: 1:openstack-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: radvd for package: 1:openstack-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: keepalived for package: 1:openstack-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: dnsmasq-utils for package: 1:openstack-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: dnsmasq for package: 1:openstack-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: dibbler-client for package: 1:openstack-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: conntrack-tools for package: 1:openstack-neutron-11.0.1-1.el7.noarch
---> Package openstack-neutron-linuxbridge.noarch 1:11.0.1-1.el7 will be installed
--> Processing Dependency: bridge-utils for package: 1:openstack-neutron-linuxbridge-11.0.1-1.el7.noarch
---> Package openstack-neutron-ml2.noarch 1:11.0.1-1.el7 will be installed
--> Processing Dependency: python-ncclient for package: 1:openstack-neutron-ml2-11.0.1-1.el7.noarch
--> Running transaction check
---> Package bridge-utils.x86_64 0:1.5-9.el7 will be installed
---> Package conntrack-tools.x86_64 0:1.4.4-3.el7_3 will be installed
--> Processing Dependency: libnetfilter_cttimeout.so.1(LIBNETFILTER_CTTIMEOUT_1.1)(64bit) for package: conntrack-tools-1.4.4-3.el7_3.x86_64
--> Processing Dependency: libnetfilter_cttimeout.so.1(LIBNETFILTER_CTTIMEOUT_1.0)(64bit) for package: conntrack-tools-1.4.4-3.el7_3.x86_64
--> Processing Dependency: libnetfilter_cthelper.so.0(LIBNETFILTER_CTHELPER_1.0)(64bit) for package: conntrack-tools-1.4.4-3.el7_3.x86_64
--> Processing Dependency: libnetfilter_queue.so.1()(64bit) for package: conntrack-tools-1.4.4-3.el7_3.x86_64
--> Processing Dependency: libnetfilter_cttimeout.so.1()(64bit) for package: conntrack-tools-1.4.4-3.el7_3.x86_64
--> Processing Dependency: libnetfilter_cthelper.so.0()(64bit) for package: conntrack-tools-1.4.4-3.el7_3.x86_64
---> Package dibbler-client.x86_64 0:1.0.1-0.RC1.2.el7 will be installed
---> Package dnsmasq.x86_64 0:2.76-2.el7_4.2 will be installed
---> Package dnsmasq-utils.x86_64 0:2.76-2.el7_4.2 will be installed
---> Package haproxy.x86_64 0:1.5.18-6.el7 will be installed
---> Package keepalived.x86_64 0:1.3.5-1.el7 will be installed
--> Processing Dependency: libnetsnmpmibs.so.31()(64bit) for package: keepalived-1.3.5-1.el7.x86_64
--> Processing Dependency: libnetsnmpagent.so.31()(64bit) for package: keepalived-1.3.5-1.el7.x86_64
--> Processing Dependency: libnetsnmp.so.31()(64bit) for package: keepalived-1.3.5-1.el7.x86_64
---> Package openstack-neutron-common.noarch 1:11.0.1-1.el7 will be installed
--> Processing Dependency: python-neutron = 1:11.0.1-1.el7 for package: 1:openstack-neutron-common-11.0.1-1.el7.noarch
---> Package python-ncclient.noarch 0:0.4.2-2.el7 will be installed
--> Processing Dependency: libxslt-python for package: python-ncclient-0.4.2-2.el7.noarch
---> Package radvd.x86_64 0:1.9.2-9.el7 will be installed
--> Running transaction check
---> Package libnetfilter_cthelper.x86_64 0:1.0.0-9.el7 will be installed
---> Package libnetfilter_cttimeout.x86_64 0:1.0.0-6.el7 will be installed
---> Package libnetfilter_queue.x86_64 0:1.0.2-2.el7_2 will be installed
---> Package libxslt-python.x86_64 0:1.1.28-5.el7 will be installed
---> Package net-snmp-agent-libs.x86_64 1:5.7.2-28.el7 will be installed
--> Processing Dependency: libsensors.so.4()(64bit) for package: 1:net-snmp-agent-libs-5.7.2-28.el7.x86_64
---> Package net-snmp-libs.x86_64 1:5.7.2-28.el7 will be installed
---> Package python-neutron.noarch 1:11.0.1-1.el7 will be installed
--> Processing Dependency: python-weakrefmethod >= 1.0.2 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-ryu >= 4.14 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-pyroute2 >= 0.4.19 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-pecan >= 1.0.0 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-os-xenapi >= 0.2.0 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-neutron-lib >= 1.9.0 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-designateclient >= 1.5.0 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-ovsdbapp for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Running transaction check
---> Package lm_sensors-libs.x86_64 0:3.4.0-4.20160601gitf9185e5.el7 will be installed
---> Package python-neutron-lib.noarch 0:1.9.1-1.el7 will be installed
---> Package python2-designateclient.noarch 0:2.7.0-1.el7 will be installed
---> Package python2-os-xenapi.noarch 0:0.2.0-1.el7 will be installed
---> Package python2-ovsdbapp.noarch 0:0.4.0-1.el7 will be installed
--> Processing Dependency: python-openvswitch for package: python2-ovsdbapp-0.4.0-1.el7.noarch
---> Package python2-pecan.noarch 0:1.1.2-1.el7 will be installed
--> Processing Dependency: python-webtest for package: python2-pecan-1.1.2-1.el7.noarch
--> Processing Dependency: python-singledispatch for package: python2-pecan-1.1.2-1.el7.noarch
--> Processing Dependency: python-logutils for package: python2-pecan-1.1.2-1.el7.noarch
---> Package python2-pyroute2.noarch 0:0.4.19-1.el7 will be installed
---> Package python2-ryu.noarch 0:4.15-1.el7 will be installed
--> Processing Dependency: python-ryu-common = 4.15-1.el7 for package: python2-ryu-4.15-1.el7.noarch
--> Processing Dependency: python-tinyrpc for package: python2-ryu-4.15-1.el7.noarch
---> Package python2-weakrefmethod.noarch 0:1.0.2-3.el7 will be installed
--> Running transaction check
---> Package python-logutils.noarch 0:0.3.3-3.el7 will be installed
---> Package python-ryu-common.noarch 0:4.15-1.el7 will be installed
---> Package python-webtest.noarch 0:2.0.23-1.el7 will be installed
--> Processing Dependency: python-waitress for package: python-webtest-2.0.23-1.el7.noarch
--> Processing Dependency: python-beautifulsoup4 for package: python-webtest-2.0.23-1.el7.noarch
---> Package python2-openvswitch.noarch 1:2.7.2-3.1fc27.el7 will be installed
---> Package python2-singledispatch.noarch 0:3.4.0.3-4.el7 will be installed
---> Package python2-tinyrpc.noarch 0:0.5-4.20170523git1f38ac.el7 will be installed
--> Processing Dependency: python-zmq for package: python2-tinyrpc-0.5-4.20170523git1f38ac.el7.noarch
--> Processing Dependency: python-werkzeug for package: python2-tinyrpc-0.5-4.20170523git1f38ac.el7.noarch
--> Processing Dependency: python-gevent for package: python2-tinyrpc-0.5-4.20170523git1f38ac.el7.noarch
--> Running transaction check
---> Package python-beautifulsoup4.noarch 0:4.6.0-1.el7 will be installed
---> Package python-waitress.noarch 0:0.8.9-5.el7 will be installed
---> Package python-werkzeug.noarch 0:0.9.1-2.el7 will be installed
---> Package python-zmq.x86_64 0:14.7.0-2.el7 will be installed
--> Processing Dependency: libzmq.so.4()(64bit) for package: python-zmq-14.7.0-2.el7.x86_64
---> Package python2-gevent.x86_64 0:1.1.2-2.el7 will be installed
--> Processing Dependency: libev.so.4()(64bit) for package: python2-gevent-1.1.2-2.el7.x86_64
--> Processing Dependency: libcares.so.2()(64bit) for package: python2-gevent-1.1.2-2.el7.x86_64
--> Running transaction check
---> Package c-ares.x86_64 0:1.10.0-3.el7 will be installed
---> Package libev.x86_64 0:4.15-7.el7 will be installed
---> Package zeromq.x86_64 0:4.0.5-4.el7 will be installed
--> Processing Dependency: libpgm-5.2.so.0()(64bit) for package: zeromq-4.0.5-4.el7.x86_64
--> Running transaction check
---> Package openpgm.x86_64 0:5.2.122-2.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                                   Arch               Version                                     Repository                         Size
==================================================================================================================================================
Installing:
 openstack-neutron                         noarch             1:11.0.1-1.el7                              centos-openstack-pike              27 k
 openstack-neutron-linuxbridge             noarch             1:11.0.1-1.el7                              centos-openstack-pike              14 k
 openstack-neutron-ml2                     noarch             1:11.0.1-1.el7                              centos-openstack-pike              13 k
Installing for dependencies:
 bridge-utils                              x86_64             1.5-9.el7                                   base                               32 k
 c-ares                                    x86_64             1.10.0-3.el7                                base                               78 k
 conntrack-tools                           x86_64             1.4.4-3.el7_3                               base                              186 k
 dibbler-client                            x86_64             1.0.1-0.RC1.2.el7                           centos-openstack-pike             409 k
 dnsmasq                                   x86_64             2.76-2.el7_4.2                              updates                           277 k
 dnsmasq-utils                             x86_64             2.76-2.el7_4.2                              updates                            29 k
 haproxy                                   x86_64             1.5.18-6.el7                                base                              834 k
 keepalived                                x86_64             1.3.5-1.el7                                 base                              327 k
 libev                                     x86_64             4.15-7.el7                                  extras                             44 k
 libnetfilter_cthelper                     x86_64             1.0.0-9.el7                                 base                               18 k
 libnetfilter_cttimeout                    x86_64             1.0.0-6.el7                                 base                               18 k
 libnetfilter_queue                        x86_64             1.0.2-2.el7_2                               base                               23 k
 libxslt-python                            x86_64             1.1.28-5.el7                                base                               59 k
 lm_sensors-libs                           x86_64             3.4.0-4.20160601gitf9185e5.el7              base                               41 k
 net-snmp-agent-libs                       x86_64             1:5.7.2-28.el7                              base                              704 k
 net-snmp-libs                             x86_64             1:5.7.2-28.el7                              base                              748 k
 openpgm                                   x86_64             5.2.122-2.el7                               centos-openstack-pike             172 k
 openstack-neutron-common                  noarch             1:11.0.1-1.el7                              centos-openstack-pike             246 k
 python-beautifulsoup4                     noarch             4.6.0-1.el7                                 centos-openstack-pike             171 k
 python-logutils                           noarch             0.3.3-3.el7                                 centos-openstack-pike              42 k
 python-ncclient                           noarch             0.4.2-2.el7                                 centos-openstack-pike             164 k
 python-neutron                            noarch             1:11.0.1-1.el7                              centos-openstack-pike             2.0 M
 python-neutron-lib                        noarch             1.9.1-1.el7                                 centos-openstack-pike             170 k
 python-ryu-common                         noarch             4.15-1.el7                                  centos-openstack-pike              51 k
 python-waitress                           noarch             0.8.9-5.el7                                 centos-openstack-pike             152 k
 python-webtest                            noarch             2.0.23-1.el7                                centos-openstack-pike              84 k
 python-werkzeug                           noarch             0.9.1-2.el7                                 extras                            562 k
 python-zmq                                x86_64             14.7.0-2.el7                                centos-openstack-pike             495 k
 python2-designateclient                   noarch             2.7.0-1.el7                                 centos-openstack-pike             112 k
 python2-gevent                            x86_64             1.1.2-2.el7                                 centos-openstack-pike             443 k
 python2-openvswitch                       noarch             1:2.7.2-3.1fc27.el7                         centos-openstack-pike             166 k
 python2-os-xenapi                         noarch             0.2.0-1.el7                                 centos-openstack-pike              35 k
 python2-ovsdbapp                          noarch             0.4.0-1.el7                                 centos-openstack-pike              46 k
 python2-pecan                             noarch             1.1.2-1.el7                                 centos-openstack-pike             268 k
 python2-pyroute2                          noarch             0.4.19-1.el7                                centos-openstack-pike             377 k
 python2-ryu                               noarch             4.15-1.el7                                  centos-openstack-pike             1.9 M
 python2-singledispatch                    noarch             3.4.0.3-4.el7                               centos-openstack-pike              18 k
 python2-tinyrpc                           noarch             0.5-4.20170523git1f38ac.el7                 centos-openstack-pike              32 k
 python2-weakrefmethod                     noarch             1.0.2-3.el7                                 centos-openstack-pike              13 k
 radvd                                     x86_64             1.9.2-9.el7                                 base                               85 k
 zeromq                                    x86_64             4.0.5-4.el7                                 centos-openstack-pike             434 k

Transaction Summary
==================================================================================================================================================
Install  3 Packages (+41 Dependent packages)

Total download size: 12 M
Installed size: 52 M
Downloading packages:
(1/44): bridge-utils-1.5-9.el7.x86_64.rpm                                                                                  |  32 kB  00:00:02     
(2/44): dnsmasq-2.76-2.el7_4.2.x86_64.rpm                                                                                  | 277 kB  00:00:02     
(3/44): conntrack-tools-1.4.4-3.el7_3.x86_64.rpm                                                                           | 186 kB  00:00:02     
(4/44): c-ares-1.10.0-3.el7.x86_64.rpm                                                                                     |  78 kB  00:00:03     
(5/44): haproxy-1.5.18-6.el7.x86_64.rpm                                                                                    | 834 kB  00:00:00     
(6/44): libev-4.15-7.el7.x86_64.rpm                                                                                        |  44 kB  00:00:00     
(7/44): libnetfilter_cttimeout-1.0.0-6.el7.x86_64.rpm                                                                      |  18 kB  00:00:00     
(8/44): libnetfilter_queue-1.0.2-2.el7_2.x86_64.rpm                                                                        |  23 kB  00:00:00     
(9/44): libxslt-python-1.1.28-5.el7.x86_64.rpm                                                                             |  59 kB  00:00:00     
(10/44): lm_sensors-libs-3.4.0-4.20160601gitf9185e5.el7.x86_64.rpm                                                         |  41 kB  00:00:00     
(11/44): net-snmp-agent-libs-5.7.2-28.el7.x86_64.rpm                                                                       | 704 kB  00:00:00     
(12/44): net-snmp-libs-5.7.2-28.el7.x86_64.rpm                                                                             | 748 kB  00:00:00     
(13/44): dibbler-client-1.0.1-0.RC1.2.el7.x86_64.rpm                                                                       | 409 kB  00:00:05     
(14/44): dnsmasq-utils-2.76-2.el7_4.2.x86_64.rpm                                                                           |  29 kB  00:00:02     
(15/44): keepalived-1.3.5-1.el7.x86_64.rpm                                                                                 | 327 kB  00:00:02     
(16/44): openstack-neutron-11.0.1-1.el7.noarch.rpm                                                                         |  27 kB  00:00:00     
(17/44): libnetfilter_cthelper-1.0.0-9.el7.x86_64.rpm                                                                      |  18 kB  00:00:02     
(18/44): openstack-neutron-common-11.0.1-1.el7.noarch.rpm                                                                  | 246 kB  00:00:00     
(19/44): openstack-neutron-linuxbridge-11.0.1-1.el7.noarch.rpm                                                             |  14 kB  00:00:00     
(20/44): openstack-neutron-ml2-11.0.1-1.el7.noarch.rpm                                                                     |  13 kB  00:00:00     
(21/44): python-beautifulsoup4-4.6.0-1.el7.noarch.rpm                                                                      | 171 kB  00:00:00     
(22/44): python-logutils-0.3.3-3.el7.noarch.rpm                                                                            |  42 kB  00:00:00     
(23/44): openpgm-5.2.122-2.el7.x86_64.rpm                                                                                  | 172 kB  00:00:04     
(24/44): python-ncclient-0.4.2-2.el7.noarch.rpm                                                                            | 164 kB  00:00:00     
(25/44): python-neutron-lib-1.9.1-1.el7.noarch.rpm                                                                         | 170 kB  00:00:00     
(26/44): python-ryu-common-4.15-1.el7.noarch.rpm                                                                           |  51 kB  00:00:00     
(27/44): python-waitress-0.8.9-5.el7.noarch.rpm                                                                            | 152 kB  00:00:00     
(28/44): python-webtest-2.0.23-1.el7.noarch.rpm                                                                            |  84 kB  00:00:00     
(29/44): python-neutron-11.0.1-1.el7.noarch.rpm                                                                            | 2.0 MB  00:00:02     
(30/44): python2-designateclient-2.7.0-1.el7.noarch.rpm                                                                    | 112 kB  00:00:00     
(31/44): python-zmq-14.7.0-2.el7.x86_64.rpm                                                                                | 495 kB  00:00:01     
(32/44): python-werkzeug-0.9.1-2.el7.noarch.rpm                                                                            | 562 kB  00:00:02     
(33/44): python2-openvswitch-2.7.2-3.1fc27.el7.noarch.rpm                                                                  | 166 kB  00:00:00     
(34/44): python2-gevent-1.1.2-2.el7.x86_64.rpm                                                                             | 443 kB  00:00:00     
(35/44): python2-os-xenapi-0.2.0-1.el7.noarch.rpm                                                                          |  35 kB  00:00:00     
(36/44): python2-ovsdbapp-0.4.0-1.el7.noarch.rpm                                                                           |  46 kB  00:00:00     
(37/44): python2-pecan-1.1.2-1.el7.noarch.rpm                                                                              | 268 kB  00:00:00     
(38/44): python2-pyroute2-0.4.19-1.el7.noarch.rpm                                                                          | 377 kB  00:00:00     
(39/44): python2-singledispatch-3.4.0.3-4.el7.noarch.rpm                                                                   |  18 kB  00:00:00     
(40/44): python2-tinyrpc-0.5-4.20170523git1f38ac.el7.noarch.rpm                                                            |  32 kB  00:00:00     
(41/44): radvd-1.9.2-9.el7.x86_64.rpm                                                                                      |  85 kB  00:00:00     
(42/44): python2-weakrefmethod-1.0.2-3.el7.noarch.rpm                                                                      |  13 kB  00:00:00     
(43/44): zeromq-4.0.5-4.el7.x86_64.rpm                                                                                     | 434 kB  00:00:01     
(44/44): python2-ryu-4.15-1.el7.noarch.rpm                                                                                 | 1.9 MB  00:00:03     
--------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                             699 kB/s |  12 MB  00:00:17     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 1:net-snmp-libs-5.7.2-28.el7.x86_64                                                                                           1/44 
  Installing : 1:python2-openvswitch-2.7.2-3.1fc27.el7.noarch                                                                                2/44 
  Installing : python2-ovsdbapp-0.4.0-1.el7.noarch                                                                                           3/44 
  Installing : python-beautifulsoup4-4.6.0-1.el7.noarch                                                                                      4/44 
  Installing : libev-4.15-7.el7.x86_64                                                                                                       5/44 
  Installing : python2-pyroute2-0.4.19-1.el7.noarch                                                                                          6/44 
  Installing : python-waitress-0.8.9-5.el7.noarch                                                                                            7/44 
  Installing : python-webtest-2.0.23-1.el7.noarch                                                                                            8/44 
  Installing : dnsmasq-2.76-2.el7_4.2.x86_64                                                                                                 9/44 
  Installing : lm_sensors-libs-3.4.0-4.20160601gitf9185e5.el7.x86_64                                                                        10/44 
  Installing : 1:net-snmp-agent-libs-5.7.2-28.el7.x86_64                                                                                    11/44 
  Installing : keepalived-1.3.5-1.el7.x86_64                                                                                                12/44 
  Installing : libxslt-python-1.1.28-5.el7.x86_64                                                                                           13/44 
  Installing : python-ncclient-0.4.2-2.el7.noarch                                                                                           14/44 
  Installing : bridge-utils-1.5-9.el7.x86_64                                                                                                15/44 
  Installing : dnsmasq-utils-2.76-2.el7_4.2.x86_64                                                                                          16/44 
  Installing : python2-designateclient-2.7.0-1.el7.noarch                                                                                   17/44 
  Installing : radvd-1.9.2-9.el7.x86_64                                                                                                     18/44 
  Installing : python-werkzeug-0.9.1-2.el7.noarch                                                                                           19/44 
  Installing : python-neutron-lib-1.9.1-1.el7.noarch                                                                                        20/44 
  Installing : c-ares-1.10.0-3.el7.x86_64                                                                                                   21/44 
  Installing : python2-gevent-1.1.2-2.el7.x86_64                                                                                            22/44 
  Installing : python2-weakrefmethod-1.0.2-3.el7.noarch                                                                                     23/44 
  Installing : openpgm-5.2.122-2.el7.x86_64                                                                                                 24/44 
  Installing : zeromq-4.0.5-4.el7.x86_64                                                                                                    25/44 
  Installing : python-zmq-14.7.0-2.el7.x86_64                                                                                               26/44 
  Installing : python2-tinyrpc-0.5-4.20170523git1f38ac.el7.noarch                                                                           27/44 
  Installing : haproxy-1.5.18-6.el7.x86_64                                                                                                  28/44 
  Installing : libnetfilter_queue-1.0.2-2.el7_2.x86_64                                                                                      29/44 
  Installing : python-logutils-0.3.3-3.el7.noarch                                                                                           30/44 
  Installing : libnetfilter_cthelper-1.0.0-9.el7.x86_64                                                                                     31/44 
  Installing : python-ryu-common-4.15-1.el7.noarch                                                                                          32/44 
  Installing : python2-ryu-4.15-1.el7.noarch                                                                                                33/44 
  Installing : dibbler-client-1.0.1-0.RC1.2.el7.x86_64                                                                                      34/44 
  Installing : python2-os-xenapi-0.2.0-1.el7.noarch                                                                                         35/44 
  Installing : libnetfilter_cttimeout-1.0.0-6.el7.x86_64                                                                                    36/44 
  Installing : conntrack-tools-1.4.4-3.el7_3.x86_64                                                                                         37/44 
  Installing : python2-singledispatch-3.4.0.3-4.el7.noarch                                                                                  38/44 
  Installing : python2-pecan-1.1.2-1.el7.noarch                                                                                             39/44 
  Installing : 1:python-neutron-11.0.1-1.el7.noarch                                                                                         40/44 
  Installing : 1:openstack-neutron-common-11.0.1-1.el7.noarch                                                                               41/44 
  Installing : 1:openstack-neutron-linuxbridge-11.0.1-1.el7.noarch                                                                          42/44 
  Installing : 1:openstack-neutron-11.0.1-1.el7.noarch                                                                                      43/44 
  Installing : 1:openstack-neutron-ml2-11.0.1-1.el7.noarch                                                                                  44/44 
  Verifying  : 1:python2-openvswitch-2.7.2-3.1fc27.el7.noarch                                                                                1/44 
  Verifying  : python2-singledispatch-3.4.0.3-4.el7.noarch                                                                                   2/44 
  Verifying  : libnetfilter_cttimeout-1.0.0-6.el7.x86_64                                                                                     3/44 
  Verifying  : python2-os-xenapi-0.2.0-1.el7.noarch                                                                                          4/44 
  Verifying  : dibbler-client-1.0.1-0.RC1.2.el7.x86_64                                                                                       5/44 
  Verifying  : python-ryu-common-4.15-1.el7.noarch                                                                                           6/44 
  Verifying  : libnetfilter_cthelper-1.0.0-9.el7.x86_64                                                                                      7/44 
  Verifying  : python-logutils-0.3.3-3.el7.noarch                                                                                            8/44 
  Verifying  : conntrack-tools-1.4.4-3.el7_3.x86_64                                                                                          9/44 
  Verifying  : python-zmq-14.7.0-2.el7.x86_64                                                                                               10/44 
  Verifying  : 1:openstack-neutron-linuxbridge-11.0.1-1.el7.noarch                                                                          11/44 
  Verifying  : python2-ryu-4.15-1.el7.noarch                                                                                                12/44 
  Verifying  : python2-ovsdbapp-0.4.0-1.el7.noarch                                                                                          13/44 
  Verifying  : libnetfilter_queue-1.0.2-2.el7_2.x86_64                                                                                      14/44 
  Verifying  : 1:python-neutron-11.0.1-1.el7.noarch                                                                                         15/44 
  Verifying  : haproxy-1.5.18-6.el7.x86_64                                                                                                  16/44 
  Verifying  : openpgm-5.2.122-2.el7.x86_64                                                                                                 17/44 
  Verifying  : python2-weakrefmethod-1.0.2-3.el7.noarch                                                                                     18/44 
  Verifying  : c-ares-1.10.0-3.el7.x86_64                                                                                                   19/44 
  Verifying  : python2-tinyrpc-0.5-4.20170523git1f38ac.el7.noarch                                                                           20/44 
  Verifying  : python-neutron-lib-1.9.1-1.el7.noarch                                                                                        21/44 
  Verifying  : python-werkzeug-0.9.1-2.el7.noarch                                                                                           22/44 
  Verifying  : radvd-1.9.2-9.el7.x86_64                                                                                                     23/44 
  Verifying  : python2-designateclient-2.7.0-1.el7.noarch                                                                                   24/44 
  Verifying  : zeromq-4.0.5-4.el7.x86_64                                                                                                    25/44 
  Verifying  : python2-pecan-1.1.2-1.el7.noarch                                                                                             26/44 
  Verifying  : dnsmasq-utils-2.76-2.el7_4.2.x86_64                                                                                          27/44 
  Verifying  : bridge-utils-1.5-9.el7.x86_64                                                                                                28/44 
  Verifying  : libxslt-python-1.1.28-5.el7.x86_64                                                                                           29/44 
  Verifying  : python-webtest-2.0.23-1.el7.noarch                                                                                           30/44 
  Verifying  : lm_sensors-libs-3.4.0-4.20160601gitf9185e5.el7.x86_64                                                                        31/44 
  Verifying  : keepalived-1.3.5-1.el7.x86_64                                                                                                32/44 
  Verifying  : dnsmasq-2.76-2.el7_4.2.x86_64                                                                                                33/44 
  Verifying  : 1:net-snmp-libs-5.7.2-28.el7.x86_64                                                                                          34/44 
  Verifying  : 1:openstack-neutron-11.0.1-1.el7.noarch                                                                                      35/44 
  Verifying  : python-waitress-0.8.9-5.el7.noarch                                                                                           36/44 
  Verifying  : python2-pyroute2-0.4.19-1.el7.noarch                                                                                         37/44 
  Verifying  : libev-4.15-7.el7.x86_64                                                                                                      38/44 
  Verifying  : 1:openstack-neutron-common-11.0.1-1.el7.noarch                                                                               39/44 
  Verifying  : 1:openstack-neutron-ml2-11.0.1-1.el7.noarch                                                                                  40/44 
  Verifying  : python-ncclient-0.4.2-2.el7.noarch                                                                                           41/44 
  Verifying  : 1:net-snmp-agent-libs-5.7.2-28.el7.x86_64                                                                                    42/44 
  Verifying  : python-beautifulsoup4-4.6.0-1.el7.noarch                                                                                     43/44 
  Verifying  : python2-gevent-1.1.2-2.el7.x86_64                                                                                            44/44 

Installed:
  openstack-neutron.noarch 1:11.0.1-1.el7    openstack-neutron-linuxbridge.noarch 1:11.0.1-1.el7    openstack-neutron-ml2.noarch 1:11.0.1-1.el7   

Dependency Installed:
  bridge-utils.x86_64 0:1.5-9.el7                                     c-ares.x86_64 0:1.10.0-3.el7                                                
  conntrack-tools.x86_64 0:1.4.4-3.el7_3                              dibbler-client.x86_64 0:1.0.1-0.RC1.2.el7                                   
  dnsmasq.x86_64 0:2.76-2.el7_4.2                                     dnsmasq-utils.x86_64 0:2.76-2.el7_4.2                                       
  haproxy.x86_64 0:1.5.18-6.el7                                       keepalived.x86_64 0:1.3.5-1.el7                                             
  libev.x86_64 0:4.15-7.el7                                           libnetfilter_cthelper.x86_64 0:1.0.0-9.el7                                  
  libnetfilter_cttimeout.x86_64 0:1.0.0-6.el7                         libnetfilter_queue.x86_64 0:1.0.2-2.el7_2                                   
  libxslt-python.x86_64 0:1.1.28-5.el7                                lm_sensors-libs.x86_64 0:3.4.0-4.20160601gitf9185e5.el7                     
  net-snmp-agent-libs.x86_64 1:5.7.2-28.el7                           net-snmp-libs.x86_64 1:5.7.2-28.el7                                         
  openpgm.x86_64 0:5.2.122-2.el7                                      openstack-neutron-common.noarch 1:11.0.1-1.el7                              
  python-beautifulsoup4.noarch 0:4.6.0-1.el7                          python-logutils.noarch 0:0.3.3-3.el7                                        
  python-ncclient.noarch 0:0.4.2-2.el7                                python-neutron.noarch 1:11.0.1-1.el7                                        
  python-neutron-lib.noarch 0:1.9.1-1.el7                             python-ryu-common.noarch 0:4.15-1.el7                                       
  python-waitress.noarch 0:0.8.9-5.el7                                python-webtest.noarch 0:2.0.23-1.el7                                        
  python-werkzeug.noarch 0:0.9.1-2.el7                                python-zmq.x86_64 0:14.7.0-2.el7                                            
  python2-designateclient.noarch 0:2.7.0-1.el7                        python2-gevent.x86_64 0:1.1.2-2.el7                                         
  python2-openvswitch.noarch 1:2.7.2-3.1fc27.el7                      python2-os-xenapi.noarch 0:0.2.0-1.el7                                      
  python2-ovsdbapp.noarch 0:0.4.0-1.el7                               python2-pecan.noarch 0:1.1.2-1.el7                                          
  python2-pyroute2.noarch 0:0.4.19-1.el7                              python2-ryu.noarch 0:4.15-1.el7                                             
  python2-singledispatch.noarch 0:3.4.0.3-4.el7                       python2-tinyrpc.noarch 0:0.5-4.20170523git1f38ac.el7                        
  python2-weakrefmethod.noarch 0:1.0.2-3.el7                          radvd.x86_64 0:1.9.2-9.el7                                                  
  zeromq.x86_64 0:4.0.5-4.el7                                        

Complete!
```

Backup "neutron.conf"
```
[vagrant@localhost ~]$ sudo cp /etc/neutron/neutron.conf etc0x2Fneutron0x2Fneutron.conf
```

Sed "neutron.conf"
```
[vagrant@localhost ~]$ sudo sed   's%^\[DEFAULT\]$%&\ndebug=true\nverbose=true\ncore_plugin=ml2\nservice_plugins=\ntransport_url=rabbit://openstack:RABBIT_PASS@${controller}\nmy_ip=${my_ip}\nauth_strategy=keystone\nnotify_nova_on_port_status_changes=true\nnotify_nova_on_port_data_changes=true\n%;s%^\[database\]$%&\nconnection=mysql+pymysql://neutron:SERVICE_DBPASS@${controller}/neutron\n%;s%^\[keystone_authtoken\]$%&\nauth_uri=http://${controller}:5000\nauth_url=http://${controller}:35357\nmemcached_servers=${controller}:11211\nauth_type=password\nproject_domain_name=default\nuser_domain_name=default\nproject_name=service\nusername=neutron\npassword=SERVICE_PASS\n%;s%^\[nova\]$%&\nauth_url=http://${controller}:35357\nauth_type=password\nproject_domain_name=default\nuser_domain_name=default\nregion_name=RegionOne\nproject_name=service\nusername=nova\npassword=SERVICE_PASS\n%;s%^\[oslo_concurrency\]$%&\nlock_path=/var/lib/neutron/tmp\n%' etc0x2Fneutron0x2Fneutron.conf | env controller=10.64.33.64 my_ip=10.64.33.64 envsubst > neutron.conf
```

Modify "neutron.conf"
```
[vagrant@localhost ~]$ sudo cp neutron.conf /etc/neutron/
```

Backup "ml2_conf.ini"
```
[vagrant@localhost ~]$ sudo cp /etc/neutron/plugins/ml2/ml2_conf.ini etc0x2Fneutron0x2Fplugins0x2Fml20x2Fml2_conf.ini
```

Sed "ml2_conf.ini"
```
[vagrant@localhost ~]$ sudo sed 's/^\[ml2\]$/&\ntype_drivers=flat,vlan\ntenant_network_types=\nmechanism_drivers=linuxbridge\nextension_drivers=port_security\n/;s/^\[ml2_type_flat\]$/&\nflat_networks=provider\n/;s/^\[securitygroup\]$/&\nenable_ipset=true\n/' etc0x2Fneutron0x2Fplugins0x2Fml20x2Fml2_conf.ini > ml2_conf.ini.provider-networks
```

Modify "ml2_conf.ini"
```
[vagrant@localhost ~]$ sudo cp ml2_conf.ini.provider-networks /etc/neutron/plugins/ml2/ml2_conf.ini 
```

Specify current plugin
```
[vagrant@localhost ~]$ sudo ln -s /etc/neutron/plugins/ml2/ml2_conf.ini /etc/neutron/plugin.ini
```

View "ml2_conf.ini"
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

Backup "linuxbridge_agent.ini"
```
[vagrant@localhost ~]$ sudo cp /etc/neutron/plugins/ml2/linuxbridge_agent.ini etc0x2Fneutron0x2Fplugins0x2Fml20x2Flinuxbridge_agent.ini
```

Sed "linuxbridge_agent.ini"
```
[vagrant@localhost ~]$ sudo sed 's/^\[linux_bridge\]$/&\nphysical_interface_mappings=provider:eth2\n/;s/^\[vxlan\]$/&\nenable_vxlan=false\n/;s/^\[securitygroup\]$/&\nenable_security_group=true\nfirewall_driver=neutron.agent.linux.iptables_firewall.IptablesFirewallDriver\n/' etc0x2Fneutron0x2Fplugins0x2Fml20x2Flinuxbridge_agent.ini > linuxbridge_agent.ini.provider-networks
```

Modify "linuxbridge_agen.ini"
```
[vagrant@localhost ~]$ sudo cp linuxbridge_agent.ini.provider-networks /etc/neutron/plugins/ml2/linuxbridge_agent.ini 
```

View "linuxbridge_agent.ini"
```
[vagrant@controller-10-64-33-64 ~]$ sudo cat /etc/neutron/plugins/ml2/linuxbridge_agent.ini | egrep '^[^#]'
[DEFAULT]
[agent]
[linux_bridge]
physical_interface_mappings=provider:eth2
[securitygroup]
enable_security_group=true
firewall_driver=neutron.agent.linux.iptables_firewall.IptablesFirewallDriver
[vxlan]
enable_vxlan=false
```

Backup "dhcp_agent.ini"
```
[vagrant@localhost ~]$ sudo cp /etc/neutron/dhcp_agent.ini etc0x2Fneutron0x2Fdhcp_agent.ini
```

Sed "dhcp_agent.ini"
```
[vagrant@localhost ~]$ sudo sed 's/^\[DEFAULT\]$/&\ninterface_driver=linuxbridge\ndhcp_driver=neutron.agent.linux.dhcp.Dnsmasq\nenable_isolated_metadata=true\n/' etc0x2Fneutron0x2Fdhcp_agent.ini > dhcp_agent.ini.provider-networks
```

Modify "dhcp_agent.ini"
```
[vagrant@localhost ~]$ sudo cp dhcp_agent.ini.provider-networks /etc/neutron/dhcp_agent.ini 
```

Backup "metadata_agent.ini"
```
[vagrant@localhost ~]$ sudo cp /etc/neutron/metadata_agent.ini etc0x2Fneutron0x2Fmetadata_agent.ini
```

Sed "metadata_agent.ini"
```
[vagrant@localhost ~]$ sudo sed 's/^\[DEFAULT\]$/&\nnova_metadata_host=${metadata_host}\nmetadata_proxy_shared_secret=METADATA_SECRET\n/' etc0x2Fneutron0x2Fmetadata_agent.ini | env metadata_host=10.64.33.64 envsubst > metadata_agent.ini.provider-networks
```

Modify "metadata_agent.ini"
```
[vagrant@localhost ~]$ sudo cp metadata_agent.ini.provider-networks /etc/neutron/metadata_agent.ini 
```

### Re-configure Nova

Sed "nova.conf"
```
[vagrant@localhost ~]$ sudo sed 's%^\[neutron]$%&\nurl=http://${controller}:9696\nauth_url=http://${controller}:35357\nauth_type=password\nproject_domain_name=default\nuser_domain_name=default\nregion_name=RegionOne\nproject_name=service\nusername=neutron\npassword=SERVICE_PASS\nservice_metadata_proxy=true\nmetadata_proxy_shared_secret=METADATA_SECRET\n%' nova.conf | env controller=10.64.33.64 envsubst > nova.conf.neutron
```

Modify "nova.conf"
```
[vagrant@localhost ~]$ sudo cp nova.conf.neutron /etc/nova/nova.conf
```

### Database

Neutron
```
[vagrant@localhost ~]$ sudo su -s /bin/sh -c "neutron-db-manage --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/plugins/ml2/ml2_conf.ini upgrade head" neutron
Traceback (most recent call last):
  File "/bin/neutron-db-manage", line 10, in <module>
    sys.exit(main())
  File "/usr/lib/python2.7/site-packages/neutron/db/migration/cli.py", line 687, in main
    return_val |= bool(CONF.command.func(config, CONF.command.name))
  File "/usr/lib/python2.7/site-packages/neutron/db/migration/cli.py", line 206, in do_upgrade
    run_sanity_checks(config, revision)
  File "/usr/lib/python2.7/site-packages/neutron/db/migration/cli.py", line 671, in run_sanity_checks
    script_dir.run_env()
  File "/usr/lib/python2.7/site-packages/alembic/script/base.py", line 416, in run_env
    util.load_python_file(self.dir, 'env.py')
  File "/usr/lib/python2.7/site-packages/alembic/util/pyfiles.py", line 93, in load_python_file
    module = load_module_py(module_id, path)
  File "/usr/lib/python2.7/site-packages/alembic/util/compat.py", line 79, in load_module_py
    mod = imp.load_source(module_id, path, fp)
  File "/usr/lib/python2.7/site-packages/neutron/db/migration/alembic_migrations/env.py", line 120, in <module>
    run_migrations_online()
  File "/usr/lib/python2.7/site-packages/neutron/db/migration/alembic_migrations/env.py", line 106, in run_migrations_online
    with DBConnection(neutron_config.database.connection, connection) as conn:
  File "/usr/lib/python2.7/site-packages/neutron/db/migration/connection.py", line 32, in __enter__
    self.engine = session.create_engine(self.connection_url)
  File "/usr/lib/python2.7/site-packages/oslo_db/sqlalchemy/engines.py", line 179, in create_engine
    test_conn = _test_connection(engine, max_retries, retry_interval)
  File "/usr/lib/python2.7/site-packages/oslo_db/sqlalchemy/engines.py", line 357, in _test_connection
    return engine.connect()
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/engine/base.py", line 2091, in connect
    return self._connection_cls(self, **kwargs)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/engine/base.py", line 90, in __init__
    if connection is not None else engine.raw_connection()
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/engine/base.py", line 2177, in raw_connection
    self.pool.unique_connection, _connection)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/engine/base.py", line 2151, in _wrap_pool_connect
    e, dialect, self)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/engine/base.py", line 1461, in _handle_dbapi_exception_noconnection
    util.raise_from_cause(newraise, exc_info)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/util/compat.py", line 203, in raise_from_cause
    reraise(type(exception), exception, tb=exc_tb, cause=cause)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/engine/base.py", line 2147, in _wrap_pool_connect
    return fn()
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/pool.py", line 328, in unique_connection
    return _ConnectionFairy._checkout(self)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/pool.py", line 766, in _checkout
    fairy = _ConnectionRecord.checkout(pool)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/pool.py", line 516, in checkout
    rec = pool._do_get()
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/pool.py", line 1138, in _do_get
    self._dec_overflow()
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/util/langhelpers.py", line 66, in __exit__
    compat.reraise(exc_type, exc_value, exc_tb)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/pool.py", line 1135, in _do_get
    return self._create_connection()
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/pool.py", line 333, in _create_connection
    return _ConnectionRecord(self)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/pool.py", line 461, in __init__
    self.__connect(first_connect_check=True)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/pool.py", line 651, in __connect
    connection = pool._invoke_creator(self)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/engine/strategies.py", line 105, in connect
    return dialect.connect(*cargs, **cparams)
  File "/usr/lib64/python2.7/site-packages/sqlalchemy/engine/default.py", line 393, in connect
    return self.dbapi.connect(*cargs, **cparams)
  File "/usr/lib/python2.7/site-packages/pymysql/__init__.py", line 90, in Connect
    return Connection(*args, **kwargs)
  File "/usr/lib/python2.7/site-packages/pymysql/connections.py", line 706, in __init__
    self.connect()
  File "/usr/lib/python2.7/site-packages/pymysql/connections.py", line 932, in connect
    self._request_authentication()
  File "/usr/lib/python2.7/site-packages/pymysql/connections.py", line 1152, in _request_authentication
    auth_packet = self._read_packet()
  File "/usr/lib/python2.7/site-packages/pymysql/connections.py", line 1014, in _read_packet
    packet.check_error()
  File "/usr/lib/python2.7/site-packages/pymysql/connections.py", line 393, in check_error
    err.raise_mysql_exception(self._data)
  File "/usr/lib/python2.7/site-packages/pymysql/err.py", line 107, in raise_mysql_exception
    raise errorclass(errno, errval)
sqlalchemy.exc.OperationalError: (pymysql.err.OperationalError) (1044, u"Access denied for user 'nova'@'%' to database 'neutron'")
```

Misstake
`
[database]
connection=mysql+pymysql://nova:SERVICE_DBPASS@10.64.33.64/neutron
```
Should "...//neutron:..."


Correct
```
[vagrant@localhost ~]$ sudo su -s /bin/sh -c "neutron-db-manage --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/plugins/ml2/ml2_conf.ini upgrade head" neutron
INFO  [alembic.runtime.migration] Context impl MySQLImpl.
INFO  [alembic.runtime.migration] Will assume non-transactional DDL.
  Running upgrade for neutron ...
INFO  [alembic.runtime.migration] Context impl MySQLImpl.
INFO  [alembic.runtime.migration] Will assume non-transactional DDL.
INFO  [alembic.runtime.migration] Running upgrade  -> kilo, kilo_initial
INFO  [alembic.runtime.migration] Running upgrade kilo -> 354db87e3225, nsxv_vdr_metadata.py
INFO  [alembic.runtime.migration] Running upgrade 354db87e3225 -> 599c6a226151, neutrodb_ipam
INFO  [alembic.runtime.migration] Running upgrade 599c6a226151 -> 52c5312f6baf, Initial operations in support of address scopes
INFO  [alembic.runtime.migration] Running upgrade 52c5312f6baf -> 313373c0ffee, Flavor framework
INFO  [alembic.runtime.migration] Running upgrade 313373c0ffee -> 8675309a5c4f, network_rbac
INFO  [alembic.runtime.migration] Running upgrade 8675309a5c4f -> 45f955889773, quota_usage
INFO  [alembic.runtime.migration] Running upgrade 45f955889773 -> 26c371498592, subnetpool hash
INFO  [alembic.runtime.migration] Running upgrade 26c371498592 -> 1c844d1677f7, add order to dnsnameservers
INFO  [alembic.runtime.migration] Running upgrade 1c844d1677f7 -> 1b4c6e320f79, address scope support in subnetpool
INFO  [alembic.runtime.migration] Running upgrade 1b4c6e320f79 -> 48153cb5f051, qos db changes
INFO  [alembic.runtime.migration] Running upgrade 48153cb5f051 -> 9859ac9c136, quota_reservations
INFO  [alembic.runtime.migration] Running upgrade 9859ac9c136 -> 34af2b5c5a59, Add dns_name to Port
INFO  [alembic.runtime.migration] Running upgrade 34af2b5c5a59 -> 59cb5b6cf4d, Add availability zone
INFO  [alembic.runtime.migration] Running upgrade 59cb5b6cf4d -> 13cfb89f881a, add is_default to subnetpool
INFO  [alembic.runtime.migration] Running upgrade 13cfb89f881a -> 32e5974ada25, Add standard attribute table
INFO  [alembic.runtime.migration] Running upgrade 32e5974ada25 -> ec7fcfbf72ee, Add network availability zone
INFO  [alembic.runtime.migration] Running upgrade ec7fcfbf72ee -> dce3ec7a25c9, Add router availability zone
INFO  [alembic.runtime.migration] Running upgrade dce3ec7a25c9 -> c3a73f615e4, Add ip_version to AddressScope
INFO  [alembic.runtime.migration] Running upgrade c3a73f615e4 -> 659bf3d90664, Add tables and attributes to support external DNS integration
INFO  [alembic.runtime.migration] Running upgrade 659bf3d90664 -> 1df244e556f5, add_unique_ha_router_agent_port_bindings
INFO  [alembic.runtime.migration] Running upgrade 1df244e556f5 -> 19f26505c74f, Auto Allocated Topology - aka Get-Me-A-Network
INFO  [alembic.runtime.migration] Running upgrade 19f26505c74f -> 15be73214821, add dynamic routing model data
INFO  [alembic.runtime.migration] Running upgrade 15be73214821 -> b4caf27aae4, add_bgp_dragent_model_data
INFO  [alembic.runtime.migration] Running upgrade b4caf27aae4 -> 15e43b934f81, rbac_qos_policy
INFO  [alembic.runtime.migration] Running upgrade 15e43b934f81 -> 31ed664953e6, Add resource_versions row to agent table
INFO  [alembic.runtime.migration] Running upgrade 31ed664953e6 -> 2f9e956e7532, tag support
INFO  [alembic.runtime.migration] Running upgrade 2f9e956e7532 -> 3894bccad37f, add_timestamp_to_base_resources
INFO  [alembic.runtime.migration] Running upgrade 3894bccad37f -> 0e66c5227a8a, Add desc to standard attr table
INFO  [alembic.runtime.migration] Running upgrade 0e66c5227a8a -> 45f8dd33480b, qos dscp db addition
INFO  [alembic.runtime.migration] Running upgrade 45f8dd33480b -> 5abc0278ca73, Add support for VLAN trunking
INFO  [alembic.runtime.migration] Running upgrade kilo -> 30018084ec99, Initial no-op Liberty contract rule.
INFO  [alembic.runtime.migration] Running upgrade 30018084ec99 -> 4ffceebfada, network_rbac
INFO  [alembic.runtime.migration] Running upgrade 4ffceebfada -> 5498d17be016, Drop legacy OVS and LB plugin tables
INFO  [alembic.runtime.migration] Running upgrade 5498d17be016 -> 2a16083502f3, Metaplugin removal
INFO  [alembic.runtime.migration] Running upgrade 2a16083502f3 -> 2e5352a0ad4d, Add missing foreign keys
INFO  [alembic.runtime.migration] Running upgrade 2e5352a0ad4d -> 11926bcfe72d, add geneve ml2 type driver
INFO  [alembic.runtime.migration] Running upgrade 11926bcfe72d -> 4af11ca47297, Drop cisco monolithic tables
INFO  [alembic.runtime.migration] Running upgrade 4af11ca47297 -> 1b294093239c, Drop embrane plugin table
INFO  [alembic.runtime.migration] Running upgrade 1b294093239c -> 8a6d8bdae39, standardattributes migration
INFO  [alembic.runtime.migration] Running upgrade 8a6d8bdae39 -> 2b4c2465d44b, DVR sheduling refactoring
INFO  [alembic.runtime.migration] Running upgrade 2b4c2465d44b -> e3278ee65050, Drop NEC plugin tables
INFO  [alembic.runtime.migration] Running upgrade e3278ee65050 -> c6c112992c9, rbac_qos_policy
INFO  [alembic.runtime.migration] Running upgrade c6c112992c9 -> 5ffceebfada, network_rbac_external
INFO  [alembic.runtime.migration] Running upgrade 5ffceebfada -> 4ffceebfcdc, standard_desc
INFO  [alembic.runtime.migration] Running upgrade 4ffceebfcdc -> 7bbb25278f53, device_owner_ha_replicate_int
INFO  [alembic.runtime.migration] Running upgrade 7bbb25278f53 -> 89ab9a816d70, Rename ml2_network_segments table
INFO  [alembic.runtime.migration] Running upgrade 5abc0278ca73 -> d3435b514502, Add device_id index to Port
INFO  [alembic.runtime.migration] Running upgrade d3435b514502 -> 30107ab6a3ee, provisioning_blocks.py
INFO  [alembic.runtime.migration] Running upgrade 30107ab6a3ee -> c415aab1c048, add revisions table
INFO  [alembic.runtime.migration] Running upgrade c415aab1c048 -> a963b38d82f4, add dns name to portdnses
INFO  [alembic.runtime.migration] Running upgrade 89ab9a816d70 -> c879c5e1ee90, Add segment_id to subnet
INFO  [alembic.runtime.migration] Running upgrade c879c5e1ee90 -> 8fd3918ef6f4, Add segment_host_mapping table.
INFO  [alembic.runtime.migration] Running upgrade 8fd3918ef6f4 -> 4bcd4df1f426, Rename ml2_dvr_port_bindings
INFO  [alembic.runtime.migration] Running upgrade 4bcd4df1f426 -> b67e765a3524, Remove mtu column from networks.
INFO  [alembic.runtime.migration] Running upgrade a963b38d82f4 -> 3d0e74aa7d37, Add flavor_id to Router
INFO  [alembic.runtime.migration] Running upgrade 3d0e74aa7d37 -> 030a959ceafa, uniq_routerports0port_id
INFO  [alembic.runtime.migration] Running upgrade 030a959ceafa -> a5648cfeeadf, Add support for Subnet Service Types
INFO  [alembic.runtime.migration] Running upgrade a5648cfeeadf -> 0f5bef0f87d4, add_qos_minimum_bandwidth_rules
INFO  [alembic.runtime.migration] Running upgrade 0f5bef0f87d4 -> 67daae611b6e, add standardattr to qos policies
INFO  [alembic.runtime.migration] Running upgrade 67daae611b6e -> 6b461a21bcfc, uniq_floatingips0floating_network_id0fixed_port_id0fixed_ip_addr
INFO  [alembic.runtime.migration] Running upgrade 6b461a21bcfc -> 5cd92597d11d, Add ip_allocation to port
INFO  [alembic.runtime.migration] Running upgrade 5cd92597d11d -> 929c968efe70, add_pk_version_table
INFO  [alembic.runtime.migration] Running upgrade 929c968efe70 -> a9c43481023c, extend_pk_with_host_and_add_status_to_ml2_port_binding
INFO  [alembic.runtime.migration] Running upgrade a9c43481023c -> 804a3c76314c, Add data_plane_status to Port
INFO  [alembic.runtime.migration] Running upgrade 804a3c76314c -> 2b42d90729da, qos add direction to bw_limit_rule table
INFO  [alembic.runtime.migration] Running upgrade 2b42d90729da -> 62c781cb6192, add is default to qos policies
INFO  [alembic.runtime.migration] Running upgrade 62c781cb6192 -> c8c222d42aa9, logging api
INFO  [alembic.runtime.migration] Running upgrade c8c222d42aa9 -> 349b6fd605a6, Add dns_domain to portdnses
INFO  [alembic.runtime.migration] Running upgrade 349b6fd605a6 -> 7d32f979895f, add mtu for networks
INFO  [alembic.runtime.migration] Running upgrade b67e765a3524 -> a84ccf28f06a, migrate dns name from port
INFO  [alembic.runtime.migration] Running upgrade a84ccf28f06a -> 7d9d8eeec6ad, rename tenant to project
INFO  [alembic.runtime.migration] Running upgrade 7d9d8eeec6ad -> a8b517cff8ab, Add routerport bindings for L3 HA
INFO  [alembic.runtime.migration] Running upgrade a8b517cff8ab -> 3b935b28e7a0, migrate to pluggable ipam
INFO  [alembic.runtime.migration] Running upgrade 3b935b28e7a0 -> b12a3ef66e62, add standardattr to qos policies
INFO  [alembic.runtime.migration] Running upgrade b12a3ef66e62 -> 97c25b0d2353, Add Name and Description to the networksegments table
INFO  [alembic.runtime.migration] Running upgrade 97c25b0d2353 -> 2e0d7a8a1586, Add binding index to RouterL3AgentBinding
INFO  [alembic.runtime.migration] Running upgrade 2e0d7a8a1586 -> 5c85685d616d, Remove availability ranges.
  OK
```

Have a look at
```
[vagrant@localhost ~]$ mysql -u root -e "show tables in neutron;"
+-----------------------------------------+
| Tables_in_neutron                       |
+-----------------------------------------+
| address_scopes                          |
| agents                                  |
| alembic_version                         |
| allowedaddresspairs                     |
| arista_provisioned_nets                 |
| arista_provisioned_tenants              |
| arista_provisioned_vms                  |
| auto_allocated_topologies               |
| bgp_peers                               |
| bgp_speaker_dragent_bindings            |
| bgp_speaker_network_bindings            |
| bgp_speaker_peer_bindings               |
| bgp_speakers                            |
| brocadenetworks                         |
| brocadeports                            |
| cisco_csr_identifier_map                |
| cisco_hosting_devices                   |
| cisco_ml2_apic_contracts                |
| cisco_ml2_apic_host_links               |
| cisco_ml2_apic_names                    |
| cisco_ml2_n1kv_network_bindings         |
| cisco_ml2_n1kv_network_profiles         |
| cisco_ml2_n1kv_policy_profiles          |
| cisco_ml2_n1kv_port_bindings            |
| cisco_ml2_n1kv_profile_bindings         |
| cisco_ml2_n1kv_vlan_allocations         |
| cisco_ml2_n1kv_vxlan_allocations        |
| cisco_ml2_nexus_nve                     |
| cisco_ml2_nexusport_bindings            |
| cisco_port_mappings                     |
| cisco_router_mappings                   |
| consistencyhashes                       |
| default_security_group                  |
| dnsnameservers                          |
| dvr_host_macs                           |
| externalnetworks                        |
| extradhcpopts                           |
| firewall_policies                       |
| firewall_rules                          |
| firewalls                               |
| flavors                                 |
| flavorserviceprofilebindings            |
| floatingipdnses                         |
| floatingips                             |
| ha_router_agent_port_bindings           |
| ha_router_networks                      |
| ha_router_vrid_allocations              |
| healthmonitors                          |
| ikepolicies                             |
| ipallocationpools                       |
| ipallocations                           |
| ipamallocationpools                     |
| ipamallocations                         |
| ipamsubnets                             |
| ipsec_site_connections                  |
| ipsecpeercidrs                          |
| ipsecpolicies                           |
| logs                                    |
| lsn                                     |
| lsn_port                                |
| maclearningstates                       |
| members                                 |
| meteringlabelrules                      |
| meteringlabels                          |
| ml2_brocadenetworks                     |
| ml2_brocadeports                        |
| ml2_distributed_port_bindings           |
| ml2_flat_allocations                    |
| ml2_geneve_allocations                  |
| ml2_geneve_endpoints                    |
| ml2_gre_allocations                     |
| ml2_gre_endpoints                       |
| ml2_nexus_vxlan_allocations             |
| ml2_nexus_vxlan_mcast_groups            |
| ml2_port_binding_levels                 |
| ml2_port_bindings                       |
| ml2_ucsm_port_profiles                  |
| ml2_vlan_allocations                    |
| ml2_vxlan_allocations                   |
| ml2_vxlan_endpoints                     |
| multi_provider_networks                 |
| networkconnections                      |
| networkdhcpagentbindings                |
| networkdnsdomains                       |
| networkgatewaydevicereferences          |
| networkgatewaydevices                   |
| networkgateways                         |
| networkqueuemappings                    |
| networkrbacs                            |
| networks                                |
| networksecuritybindings                 |
| networksegments                         |
| neutron_nsx_network_mappings            |
| neutron_nsx_port_mappings               |
| neutron_nsx_router_mappings             |
| neutron_nsx_security_group_mappings     |
| nexthops                                |
| nsxv_edge_dhcp_static_bindings          |
| nsxv_edge_vnic_bindings                 |
| nsxv_firewall_rule_bindings             |
| nsxv_internal_edges                     |
| nsxv_internal_networks                  |
| nsxv_port_index_mappings                |
| nsxv_port_vnic_mappings                 |
| nsxv_router_bindings                    |
| nsxv_router_ext_attributes              |
| nsxv_rule_mappings                      |
| nsxv_security_group_section_mappings    |
| nsxv_spoofguard_policy_network_mappings |
| nsxv_tz_network_bindings                |
| nsxv_vdr_dhcp_bindings                  |
| nuage_net_partition_router_mapping      |
| nuage_net_partitions                    |
| nuage_provider_net_bindings             |
| nuage_subnet_l2dom_mapping              |
| poolloadbalanceragentbindings           |
| poolmonitorassociations                 |
| pools                                   |
| poolstatisticss                         |
| portbindingports                        |
| portdataplanestatuses                   |
| portdnses                               |
| portqueuemappings                       |
| ports                                   |
| portsecuritybindings                    |
| providerresourceassociations            |
| provisioningblocks                      |
| qos_bandwidth_limit_rules               |
| qos_dscp_marking_rules                  |
| qos_minimum_bandwidth_rules             |
| qos_network_policy_bindings             |
| qos_policies                            |
| qos_policies_default                    |
| qos_port_policy_bindings                |
| qospolicyrbacs                          |
| qosqueues                               |
| quotas                                  |
| quotausages                             |
| reservations                            |
| resourcedeltas                          |
| router_extra_attributes                 |
| routerl3agentbindings                   |
| routerports                             |
| routerroutes                            |
| routerrules                             |
| routers                                 |
| securitygroupportbindings               |
| securitygrouprules                      |
| securitygroups                          |
| segmenthostmappings                     |
| serviceprofiles                         |
| sessionpersistences                     |
| standardattributes                      |
| subnet_service_types                    |
| subnetpoolprefixes                      |
| subnetpools                             |
| subnetroutes                            |
| subnets                                 |
| subports                                |
| tags                                    |
| trunks                                  |
| tz_network_bindings                     |
| vcns_router_bindings                    |
| vips                                    |
| vpnservices                             |
+-----------------------------------------+
```

### Restart Nova API

Restart
```
[vagrant@localhost ~]$ sudo systemctl restart openstack-nova-api.service
```

```
[vagrant@localhost ~]$ sudo tail /var/log/nova/nova-api.log
[vagrant@localhost ~]$ sudo tail -f /var/log/nova/nova-api.log 
2017-10-21 21:36:01.022 7381 DEBUG oslo_service.service [req-129ab858-9987-42ed-9600-609563c368f6 - - - - -] upgrade_levels.console         = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 21:36:01.022 7381 DEBUG oslo_service.service [req-129ab858-9987-42ed-9600-609563c368f6 - - - - -] upgrade_levels.consoleauth     = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 21:36:01.022 7381 DEBUG oslo_service.service [req-129ab858-9987-42ed-9600-609563c368f6 - - - - -] upgrade_levels.intercell       = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 21:36:01.022 7381 DEBUG oslo_service.service [req-129ab858-9987-42ed-9600-609563c368f6 - - - - -] upgrade_levels.network         = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 21:36:01.022 7381 DEBUG oslo_service.service [req-129ab858-9987-42ed-9600-609563c368f6 - - - - -] upgrade_levels.scheduler       = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 21:36:01.023 7381 DEBUG oslo_service.service [req-129ab858-9987-42ed-9600-609563c368f6 - - - - -] key_manager.api_class          = castellan.key_manager.barbican_key_manager.BarbicanKeyManager log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 21:36:01.023 7381 DEBUG oslo_service.service [req-129ab858-9987-42ed-9600-609563c368f6 - - - - -] key_manager.fixed_key          = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 21:36:01.023 7381 DEBUG oslo_service.service [req-129ab858-9987-42ed-9600-609563c368f6 - - - - -] osapi_v21.project_id_regex     = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 21:36:01.023 7381 DEBUG oslo_service.service [req-129ab858-9987-42ed-9600-609563c368f6 - - - - -] ******************************************************************************** log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2889
2017-10-21 21:36:01.027 7394 INFO nova.metadata.wsgi.server [req-a9ba62e1-5e71-4dc4-aec3-a6924657a014 - - - - -] (7394) wsgi starting up on http://0.0.0.0:8775
```

### Start Neutron Server

neutron server
```
[vagrant@localhost ~]$ sudo systemctl start neutron-server.service

[vagrant@localhost ~]$ sudo systemctl -l status neutron-server.service
● neutron-server.service - OpenStack Neutron Server
   Loaded: loaded (/usr/lib/systemd/system/neutron-server.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 21:43:50 UTC; 2min 23s ago
 Main PID: 7542 (neutron-server)
   CGroup: /system.slice/neutron-server.service
           ├─7542 /usr/bin/python2 /usr/bin/neutron-server --config-file /usr/share/neutron/neutron-dist.conf --config-dir /usr/share/neutron/server --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/plugin.ini --config-dir /etc/neutron/conf.d/common --config-dir /etc/neutron/conf.d/neutron-server --log-file /var/log/neutron/server.log
           ├─7553 /usr/bin/python2 /usr/bin/neutron-server --config-file /usr/share/neutron/neutron-dist.conf --config-dir /usr/share/neutron/server --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/plugin.ini --config-dir /etc/neutron/conf.d/common --config-dir /etc/neutron/conf.d/neutron-server --log-file /var/log/neutron/server.log
           ├─7554 /usr/bin/python2 /usr/bin/neutron-server --config-file /usr/share/neutron/neutron-dist.conf --config-dir /usr/share/neutron/server --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/plugin.ini --config-dir /etc/neutron/conf.d/common --config-dir /etc/neutron/conf.d/neutron-server --log-file /var/log/neutron/server.log
           ├─7555 /usr/bin/python2 /usr/bin/neutron-server --config-file /usr/share/neutron/neutron-dist.conf --config-dir /usr/share/neutron/server --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/plugin.ini --config-dir /etc/neutron/conf.d/common --config-dir /etc/neutron/conf.d/neutron-server --log-file /var/log/neutron/server.log
           └─7556 /usr/bin/python2 /usr/bin/neutron-server --config-file /usr/share/neutron/neutron-dist.conf --config-dir /usr/share/neutron/server --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/plugin.ini --config-dir /etc/neutron/conf.d/common --config-dir /etc/neutron/conf.d/neutron-server --log-file /var/log/neutron/server.log

Oct 21 21:43:48 localhost.localdomain systemd[1]: Starting OpenStack Neutron Server...
Oct 21 21:43:48 localhost.localdomain neutron-server[7542]: Guru meditation now registers SIGUSR1 and SIGUSR2 by default for backward compatibility. SIGUSR1 will no longer be registered in a future release, so please use SIGUSR2 to generate reports.
Oct 21 21:43:50 localhost.localdomain systemd[1]: Started OpenStack Neutron Server.

[vagrant@localhost ~]$ sudo cat /var/log/neutron/server.log | grep ERROR
2017-10-21 21:43:50.364 7542 DEBUG oslo_db.sqlalchemy.engines [req-f7cf9037-7f75-4fde-8836-d730279d4229 - - - - -] MySQL server mode set to STRICT_TRANS_TABLES,STRICT_ALL_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,TRADITIONAL,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION _check_effective_sql_mode /usr/lib/python2.7/site-packages/oslo_db/sqlalchemy/engines.py:285
2017-10-21 21:43:50.677 7542 ERROR neutron.api.extensions [req-f7cf9037-7f75-4fde-8836-d730279d4229 - - - - -] Unable to process extensions (auto-allocated-topology) because the configured plugins do not satisfy their requirements. Some features will not work as expected.
2017-10-21 21:43:50.772 7542 DEBUG oslo_service.service [-] logging_exception_prefix       = %(asctime)s.%(msecs)03d %(process)d ERROR %(name)s %(instance)s log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2879
2017-10-21 21:43:50.808 7542 DEBUG oslo_service.service [-] logging_exception_prefix       = %(asctime)s.%(msecs)03d %(process)d ERROR %(name)s %(instance)s log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2879
```

Linuxbridge angent
```
[vagrant@localhost ~]$ sudo systemctl start neutron-linuxbridge-agent.service
[vagrant@localhost ~]$ sudo systemctl -l status neutron-linuxbridge-agent.service
● neutron-linuxbridge-agent.service - OpenStack Neutron Linux Bridge Agent
   Loaded: loaded (/usr/lib/systemd/system/neutron-linuxbridge-agent.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 21:46:54 UTC; 11s ago
  Process: 7615 ExecStartPre=/usr/bin/neutron-enable-bridge-firewall.sh (code=exited, status=0/SUCCESS)
 Main PID: 7624 (neutron-linuxbr)
   CGroup: /system.slice/neutron-linuxbridge-agent.service
           ├─7624 /usr/bin/python2 /usr/bin/neutron-linuxbridge-agent --config-file /usr/share/neutron/neutron-dist.conf --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/plugins/ml2/linuxbridge_agent.ini --config-dir /etc/neutron/conf.d/common --config-dir /etc/neutron/conf.d/neutron-linuxbridge-agent --log-file /var/log/neutron/linuxbridge-agent.log
           ├─7637 sudo neutron-rootwrap-daemon /etc/neutron/rootwrap.conf
           └─7638 /usr/bin/python2 /usr/bin/neutron-rootwrap-daemon /etc/neutron/rootwrap.conf

Oct 21 21:46:54 localhost.localdomain systemd[1]: Starting OpenStack Neutron Linux Bridge Agent...
Oct 21 21:46:54 localhost.localdomain neutron-enable-bridge-firewall.sh[7615]: net.bridge.bridge-nf-call-iptables = 1
Oct 21 21:46:54 localhost.localdomain neutron-enable-bridge-firewall.sh[7615]: net.bridge.bridge-nf-call-ip6tables = 1
Oct 21 21:46:54 localhost.localdomain systemd[1]: Started OpenStack Neutron Linux Bridge Agent.
Oct 21 21:46:54 localhost.localdomain neutron-linuxbridge-agent[7624]: Guru meditation now registers SIGUSR1 and SIGUSR2 by default for backward compatibility. SIGUSR1 will no longer be registered in a future release, so please use SIGUSR2 to generate reports.
Oct 21 21:46:56 localhost.localdomain sudo[7637]:  neutron : TTY=unknown ; PWD=/ ; USER=root ; COMMAND=/bin/neutron-rootwrap-daemon /etc/neutron/rootwrap.conf

[vagrant@localhost ~]$ sudo cat /var/log/neutron/linuxbridge-agent.log | grep ERROR
2017-10-21 21:46:56.422 7624 DEBUG oslo_service.service [-] logging_exception_prefix       = %(asctime)s.%(msecs)03d %(process)d ERROR %(name)s %(instance)s log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2879
```

DHCP agent
```
[vagrant@localhost ~]$ sudo systemctl start neutron-dhcp-agent.service
[vagrant@localhost ~]$ sudo systemctl -l status neutron-dhcp-agent.service
● neutron-dhcp-agent.service - OpenStack Neutron DHCP Agent
   Loaded: loaded (/usr/lib/systemd/system/neutron-dhcp-agent.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 21:49:02 UTC; 11s ago
 Main PID: 7707 (neutron-dhcp-ag)
   CGroup: /system.slice/neutron-dhcp-agent.service
           └─7707 /usr/bin/python2 /usr/bin/neutron-dhcp-agent --config-file /usr/share/neutron/neutron-dist.conf --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/dhcp_agent.ini --config-dir /etc/neutron/conf.d/common --config-dir /etc/neutron/conf.d/neutron-dhcp-agent --log-file /var/log/neutron/dhcp-agent.log

Oct 21 21:49:02 localhost.localdomain systemd[1]: Started OpenStack Neutron DHCP Agent.
Oct 21 21:49:02 localhost.localdomain systemd[1]: Starting OpenStack Neutron DHCP Agent...
Oct 21 21:49:02 localhost.localdomain neutron-dhcp-agent[7707]: Guru meditation now registers SIGUSR1 and SIGUSR2 by default for backward compatibility. SIGUSR1 will no longer be registered in a future release, so please use SIGUSR2 to generate reports.

[vagrant@localhost ~]$ sudo cat /var/log/neutron/dhcp-agent.log | grep ERROR
2017-10-21 21:49:04.345 7707 DEBUG oslo_service.service [-] logging_exception_prefix       = %(asctime)s.%(msecs)03d %(process)d ERROR %(name)s %(instance)s log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2879
```

Metadata agent
```
[vagrant@localhost ~]$ sudo systemctl start neutron-metadata-agent.service
[vagrant@localhost ~]$ sudo systemctl -l status neutron-metadata-agent.service
● neutron-metadata-agent.service - OpenStack Neutron Metadata Agent
   Loaded: loaded (/usr/lib/systemd/system/neutron-metadata-agent.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 21:50:42 UTC; 10s ago
 Main PID: 7762 (neutron-metadat)
   CGroup: /system.slice/neutron-metadata-agent.service
           └─7762 /usr/bin/python2 /usr/bin/neutron-metadata-agent --config-file /usr/share/neutron/neutron-dist.conf --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/metadata_agent.ini --config-dir /etc/neutron/conf.d/common --config-dir /etc/neutron/conf.d/neutron-metadata-agent --log-file /var/log/neutron/metadata-agent.log

Oct 21 21:50:42 localhost.localdomain systemd[1]: Started OpenStack Neutron Metadata Agent.
Oct 21 21:50:42 localhost.localdomain systemd[1]: Starting OpenStack Neutron Metadata Agent...
Oct 21 21:50:43 localhost.localdomain neutron-metadata-agent[7762]: Guru meditation now registers SIGUSR1 and SIGUSR2 by default for backward compatibility. SIGUSR1 will no longer be registered in a future release, so please use SIGUSR2 to generate reports.

[vagrant@localhost ~]$ sudo cat /var/log/neutron/metadata-agent.log | grep ERROR
2017-10-21 21:50:44.828 7762 DEBUG neutron.agent.metadata_agent [-] logging_exception_prefix       = %(asctime)s.%(msecs)03d %(process)d ERROR %(name)s %(instance)s log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2879
2017-10-21 21:50:44.861 7762 DEBUG neutron.wsgi [req-68cfe22a-b782-42ff-8943-d388197e5bad - - - - -] logging_exception_prefix       = %(asctime)s.%(msecs)03d %(process)d ERROR %(name)s %(instance)s log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2879
```

Networking
```
[vagrant@localhost ~]$ sudo netstat -tpnl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 127.0.0.1:11211         0.0.0.0:*               LISTEN      893/memcached       
tcp        0      0 0.0.0.0:9292            0.0.0.0:*               LISTEN      891/python2         
tcp        0      0 0.0.0.0:111             0.0.0.0:*               LISTEN      1/systemd           
tcp        0      0 0.0.0.0:4369            0.0.0.0:*               LISTEN      1/systemd           
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      899/sshd            
tcp        0      0 127.0.0.1:25            0.0.0.0:*               LISTEN      1236/master         
tcp        0      0 0.0.0.0:9696            0.0.0.0:*               LISTEN      7542/python2        
tcp        0      0 0.0.0.0:6080            0.0.0.0:*               LISTEN      888/python2         
tcp        0      0 0.0.0.0:8774            0.0.0.0:*               LISTEN      7381/python2        
tcp        0      0 0.0.0.0:8775            0.0.0.0:*               LISTEN      7381/python2        
tcp        0      0 0.0.0.0:9191            0.0.0.0:*               LISTEN      898/python2         
tcp        0      0 0.0.0.0:25672           0.0.0.0:*               LISTEN      895/beam            
tcp        0      0 10.64.33.64:3306        0.0.0.0:*               LISTEN      1150/mysqld         
tcp6       0      0 ::1:11211               :::*                    LISTEN      893/memcached       
tcp6       0      0 :::111                  :::*                    LISTEN      1/systemd           
tcp6       0      0 :::80                   :::*                    LISTEN      897/httpd           
tcp6       0      0 :::22                   :::*                    LISTEN      899/sshd            
tcp6       0      0 ::1:25                  :::*                    LISTEN      1236/master         
tcp6       0      0 :::35357                :::*                    LISTEN      897/httpd           
tcp6       0      0 :::5672                 :::*                    LISTEN      895/beam            
tcp6       0      0 :::5000                 :::*                    LISTEN      897/httpd           
tcp6       0      0 :::8778                 :::*                    LISTEN      897/httpd           
```

Auto starging
```
[vagrant@localhost ~]$ sudo systemctl enable neutron-server.service neutron-linuxbridge-agent.service neutron-dhcp-agent.service neutron-metadata-agent.service 
Created symlink from /etc/systemd/system/multi-user.target.wants/neutron-server.service to /usr/lib/systemd/system/neutron-server.service.
Created symlink from /etc/systemd/system/multi-user.target.wants/neutron-linuxbridge-agent.service to /usr/lib/systemd/system/neutron-linuxbridge-agent.service.
Created symlink from /etc/systemd/system/multi-user.target.wants/neutron-dhcp-agent.service to /usr/lib/systemd/system/neutron-dhcp-agent.service.
Created symlink from /etc/systemd/system/multi-user.target.wants/neutron-metadata-agent.service to /usr/lib/systemd/system/neutron-metadata-agent.service.
```

### Verifying

cli
```
[vagrant@localhost ~]$ openstack extension list --network
+----------------------------------------------------------------------------------------------+---------------------------+----------------------------------------------------------------------------------------------------------------------------------------------------------+
| Name                                                                                         | Alias                     | Description                                                                                                                                              |
+----------------------------------------------------------------------------------------------+---------------------------+----------------------------------------------------------------------------------------------------------------------------------------------------------+
| Default Subnetpools                                                                          | default-subnetpools       | Provides ability to mark and use a subnetpool as the default                                                                                             |
| Network IP Availability                                                                      | network-ip-availability   | Provides IP availability data for each network and subnet.                                                                                               |
| Network Availability Zone                                                                    | network_availability_zone | Availability zone support for network.                                                                                                                   |
| Network MTU (writable)                                                                       | net-mtu-writable          | Provides a writable MTU attribute for a network resource.                                                                                                |
| Port Binding                                                                                 | binding                   | Expose port bindings of a virtual port to external application                                                                                           |
| agent                                                                                        | agent                     | The agent management extension.                                                                                                                          |
| Subnet Allocation                                                                            | subnet_allocation         | Enables allocation of subnets from a subnet pool                                                                                                         |
| DHCP Agent Scheduler                                                                         | dhcp_agent_scheduler      | Schedule networks among dhcp agents                                                                                                                      |
| Tag support                                                                                  | tag                       | Enables to set tag on resources.                                                                                                                         |
| Neutron external network                                                                     | external-net              | Adds external network attribute to network resource.                                                                                                     |
| Neutron Service Flavors                                                                      | flavors                   | Flavor specification for Neutron advanced services                                                                                                       |
| Network MTU                                                                                  | net-mtu                   | Provides MTU attribute for a network resource.                                                                                                           |
| Availability Zone                                                                            | availability_zone         | The availability zone extension.                                                                                                                         |
| Quota management support                                                                     | quotas                    | Expose functions for quotas management per tenant                                                                                                        |
| Tag support for resources with standard attribute: trunk, policy, security_group, floatingip | standard-attr-tag         | Enables to set tag on resources with standard attribute.                                                                                                 |
| If-Match constraints based on revision_number                                                | revision-if-match         | Extension indicating that If-Match based on revision_number is supported.                                                                                |
| Provider Network                                                                             | provider                  | Expose mapping of virtual networks to physical networks                                                                                                  |
| Multi Provider Network                                                                       | multi-provider            | Expose mapping of virtual networks to multiple physical networks                                                                                         |
| Quota details management support                                                             | quota_details             | Expose functions for quotas usage statistics per project                                                                                                 |
| Address scope                                                                                | address-scope             | Address scopes extension.                                                                                                                                |
| Subnet service types                                                                         | subnet-service-types      | Provides ability to set the subnet service_types field                                                                                                   |
| Resource timestamps                                                                          | standard-attr-timestamp   | Adds created_at and updated_at fields to all Neutron resources that have Neutron standard attributes.                                                    |
| Neutron Service Type Management                                                              | service-type              | API for retrieving service providers for Neutron advanced services                                                                                       |
| Tag support for resources: subnet, subnetpool, port, router                                  | tag-ext                   | Extends tag support to more L2 and L3 resources.                                                                                                         |
| Neutron Extra DHCP options                                                                   | extra_dhcp_opt            | Extra options configuration for DHCP. For example PXE boot options to DHCP clients can be specified (e.g. tftp-server, server-ip-address, bootfile-name) |
| Resource revision numbers                                                                    | standard-attr-revisions   | This extension will display the revision number of neutron resources.                                                                                    |
| Pagination support                                                                           | pagination                | Extension that indicates that pagination is enabled.                                                                                                     |
| Sorting support                                                                              | sorting                   | Extension that indicates that sorting is enabled.                                                                                                        |
| security-group                                                                               | security-group            | The security groups extension.                                                                                                                           |
| RBAC Policies                                                                                | rbac-policies             | Allows creation and modification of policies that control tenant access to resources.                                                                    |
| standard-attr-description                                                                    | standard-attr-description | Extension to add descriptions to standard attributes                                                                                                     |
| Port Security                                                                                | port-security             | Provides port security                                                                                                                                   |
| Allowed Address Pairs                                                                        | allowed-address-pairs     | Provides allowed address pairs                                                                                                                           |
| project_id field enabled                                                                     | project-id                | Extension that indicates that project_id field is enabled.                                                                                               |
+----------------------------------------------------------------------------------------------+---------------------------+----------------------------------------------------------------------------------------------------------------------------------------------------------+
```

[vagrant@localhost ~]$ openstack network agent list
+--------------------------------------+--------------------+-----------------------+-------------------+-------+-------+---------------------------+
| ID                                   | Agent Type         | Host                  | Availability Zone | Alive | State | Binary                    |
+--------------------------------------+--------------------+-----------------------+-------------------+-------+-------+---------------------------+
| 59d59396-ebd2-4f3e-a2da-387cbcac2027 | Metadata agent     | localhost.localdomain | None              | :-)   | UP    | neutron-metadata-agent    |
| 88e03f05-8d37-423b-b7be-1d74d166d408 | DHCP agent         | localhost.localdomain | nova              | :-)   | UP    | neutron-dhcp-agent        |
| ea6174e3-62d2-48f4-a268-236883c00f9d | Linux bridge agent | localhost.localdomain | None              | :-)   | UP    | neutron-linuxbridge-agent |
+--------------------------------------+--------------------+-----------------------+-------------------+-------+-------+---------------------------+
```
