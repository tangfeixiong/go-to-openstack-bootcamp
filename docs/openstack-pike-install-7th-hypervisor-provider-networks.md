# OpenStack Pike Installation

## Table of content

计算节点
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
[vagrant@localhost ~]$ sudo nmcli d
DEVICE  TYPE      STATE      CONNECTION         
eth0    ethernet  connected  System eth0        
eth1    ethernet  connected  Wired connection 1 
eth2    ethernet  connected  Wired connection 2 
lo      loopback  unmanaged  --                 
[vagrant@localhost ~]$ sudo nmcli c
NAME                UUID                                  TYPE            DEVICE 
System eth0         5fb06bd0-0bb0-7ffb-45f1-d6edd65f3e03  802-3-ethernet  eth0   
Wired connection 1  0ab4e0ff-f77e-3886-9ba9-1838e4120850  802-3-ethernet  eth1   
Wired connection 2  eb0a2ba8-d0c6-36d2-8959-8ed82367330c  802-3-ethernet  eth2   
```

```
[vagrant@localhost ~]$ sudo nmcli con delete "Wired connection 2"
Connection 'Wired connection 2' (eb0a2ba8-d0c6-36d2-8959-8ed82367330c) successfully deleted.
[vagrant@localhost ~]$ sudo nmcli con reload
[vagrant@localhost ~]$ sudo nmcli con
NAME                UUID                                  TYPE            DEVICE 
System eth0         5fb06bd0-0bb0-7ffb-45f1-d6edd65f3e03  802-3-ethernet  eth0   
Wired connection 1  0ab4e0ff-f77e-3886-9ba9-1838e4120850  802-3-ethernet  eth1   
[vagrant@localhost ~]$ sudo nmcli dev
DEVICE  TYPE      STATE         CONNECTION         
eth0    ethernet  connected     System eth0        
eth1    ethernet  connected     Wired connection 1 
eth2    ethernet  disconnected  --                 
lo      loopback  unmanaged     --                  
```
Last 
```
[vagrant@localhost ~]$ ip a show eth2
4: eth2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 08:00:27:f6:b8:b3 brd ff:ff:ff:ff:ff:ff
```

```
[vagrant@localhost ~]$ ls /etc/sysconfig/network-scripts/ifcfg-*
/etc/sysconfig/network-scripts/ifcfg-eth0  /etc/sysconfig/network-scripts/ifcfg-lo  /etc/sysconfig/network-scripts/ifcfg-Wired_connection_1
```

Generate
```
[vagrant@localhost ~]$ cat <<EOF > ifcfg-eth2
> DEVICE=eth2
> TYPE=Ethernet
> ONBOOT="yes"
> BOOTPROTO="none"
> EOF
```

Modify
```
[vagrant@localhost ~]$ sudo cp ifcfg-eth2 /etc/sysconfig/network-scripts/
```



### Hostname

Backup
```
[vagrant@localhost ~]$ sudo cp /etc/hostname etc0x2Fhostname
[vagrant@localhost ~]$ sudo echo "compute-10-64-33-65" > hostname
```

Modify
```
[vagrant@localhost ~]$ sudo hostname --file hostname
```

Check
```
[vagrant@localhost ~]$ sudo cp hostname /etc/hostname
[vagrant@localhost ~]$ hostname
compute-10-64-33-65
```


### Delivery Packages

[Neutron](https://docs.openstack.org/neutron/pike/install/compute-install-rdo.html)
```
[vagrant@localhost ~]$ sudo yum install -y openstack-neutron-linuxbridge ebtables ipset
Loaded plugins: fastestmirror
base                                                                                                                       | 3.6 kB  00:00:00     
centos-ceph-jewel                                                                                                          | 2.9 kB  00:00:00     
centos-openstack-pike                                                                                                      | 2.9 kB  00:00:00     
centos-qemu-ev                                                                                                             | 2.9 kB  00:00:00     
extras                                                                                                                     | 3.4 kB  00:00:00     
updates                                                                                                                    | 3.4 kB  00:00:00     
Loading mirror speeds from cached hostfile
 * base: mirrors.btte.net
 * extras: mirrors.sohu.com
 * updates: mirrors.sohu.com
Package ebtables-2.0.10-15.el7.x86_64 already installed and latest version
Package ipset-6.29-1.el7.x86_64 already installed and latest version
Resolving Dependencies
--> Running transaction check
---> Package openstack-neutron-linuxbridge.noarch 1:11.0.1-1.el7 will be installed
--> Processing Dependency: openstack-neutron-common = 1:11.0.1-1.el7 for package: 1:openstack-neutron-linuxbridge-11.0.1-1.el7.noarch
--> Running transaction check
---> Package openstack-neutron-common.noarch 1:11.0.1-1.el7 will be installed
--> Processing Dependency: python-neutron = 1:11.0.1-1.el7 for package: 1:openstack-neutron-common-11.0.1-1.el7.noarch
--> Running transaction check
---> Package python-neutron.noarch 1:11.0.1-1.el7 will be installed
--> Processing Dependency: python-weakrefmethod >= 1.0.2 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-ryu >= 4.14 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-pyroute2 >= 0.4.19 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-pecan >= 1.0.0 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-osprofiler >= 1.4.0 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-os-xenapi >= 0.2.0 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-neutron-lib >= 1.9.0 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-httplib2 >= 0.7.5 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-designateclient >= 1.5.0 for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Processing Dependency: python-ovsdbapp for package: 1:python-neutron-11.0.1-1.el7.noarch
--> Running transaction check
---> Package python-httplib2.noarch 0:0.9.2-1.el7 will be installed
---> Package python-neutron-lib.noarch 0:1.9.1-1.el7 will be installed
---> Package python2-designateclient.noarch 0:2.7.0-1.el7 will be installed
---> Package python2-os-xenapi.noarch 0:0.2.0-1.el7 will be installed
---> Package python2-osprofiler.noarch 0:1.11.0-1.el7 will be installed
---> Package python2-ovsdbapp.noarch 0:0.4.0-1.el7 will be installed
--> Processing Dependency: python-openvswitch for package: python2-ovsdbapp-0.4.0-1.el7.noarch
---> Package python2-pecan.noarch 0:1.1.2-1.el7 will be installed
--> Processing Dependency: python-webtest for package: python2-pecan-1.1.2-1.el7.noarch
--> Processing Dependency: python-singledispatch for package: python2-pecan-1.1.2-1.el7.noarch
--> Processing Dependency: python-simplegeneric for package: python2-pecan-1.1.2-1.el7.noarch
--> Processing Dependency: python-logutils for package: python2-pecan-1.1.2-1.el7.noarch
---> Package python2-pyroute2.noarch 0:0.4.19-1.el7 will be installed
---> Package python2-ryu.noarch 0:4.15-1.el7 will be installed
--> Processing Dependency: python-ryu-common = 4.15-1.el7 for package: python2-ryu-4.15-1.el7.noarch
--> Processing Dependency: python-tinyrpc for package: python2-ryu-4.15-1.el7.noarch
---> Package python2-weakrefmethod.noarch 0:1.0.2-3.el7 will be installed
--> Running transaction check
---> Package python-logutils.noarch 0:0.3.3-3.el7 will be installed
---> Package python-ryu-common.noarch 0:4.15-1.el7 will be installed
---> Package python-simplegeneric.noarch 0:0.8-7.el7 will be installed
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
 Package                                    Arch                Version                                  Repository                          Size
==================================================================================================================================================
Installing:
 openstack-neutron-linuxbridge              noarch              1:11.0.1-1.el7                           centos-openstack-pike               14 k
Installing for dependencies:
 c-ares                                     x86_64              1.10.0-3.el7                             base                                78 k
 libev                                      x86_64              4.15-7.el7                               extras                              44 k
 openpgm                                    x86_64              5.2.122-2.el7                            centos-openstack-pike              172 k
 openstack-neutron-common                   noarch              1:11.0.1-1.el7                           centos-openstack-pike              246 k
 python-beautifulsoup4                      noarch              4.6.0-1.el7                              centos-openstack-pike              171 k
 python-httplib2                            noarch              0.9.2-1.el7                              centos-openstack-pike              115 k
 python-logutils                            noarch              0.3.3-3.el7                              centos-openstack-pike               42 k
 python-neutron                             noarch              1:11.0.1-1.el7                           centos-openstack-pike              2.0 M
 python-neutron-lib                         noarch              1.9.1-1.el7                              centos-openstack-pike              170 k
 python-ryu-common                          noarch              4.15-1.el7                               centos-openstack-pike               51 k
 python-simplegeneric                       noarch              0.8-7.el7                                centos-openstack-pike               12 k
 python-waitress                            noarch              0.8.9-5.el7                              centos-openstack-pike              152 k
 python-webtest                             noarch              2.0.23-1.el7                             centos-openstack-pike               84 k
 python-werkzeug                            noarch              0.9.1-2.el7                              extras                             562 k
 python-zmq                                 x86_64              14.7.0-2.el7                             centos-openstack-pike              495 k
 python2-designateclient                    noarch              2.7.0-1.el7                              centos-openstack-pike              112 k
 python2-gevent                             x86_64              1.1.2-2.el7                              centos-openstack-pike              443 k
 python2-openvswitch                        noarch              1:2.7.2-3.1fc27.el7                      centos-openstack-pike              166 k
 python2-os-xenapi                          noarch              0.2.0-1.el7                              centos-openstack-pike               35 k
 python2-osprofiler                         noarch              1.11.0-1.el7                             centos-openstack-pike              114 k
 python2-ovsdbapp                           noarch              0.4.0-1.el7                              centos-openstack-pike               46 k
 python2-pecan                              noarch              1.1.2-1.el7                              centos-openstack-pike              268 k
 python2-pyroute2                           noarch              0.4.19-1.el7                             centos-openstack-pike              377 k
 python2-ryu                                noarch              4.15-1.el7                               centos-openstack-pike              1.9 M
 python2-singledispatch                     noarch              3.4.0.3-4.el7                            centos-openstack-pike               18 k
 python2-tinyrpc                            noarch              0.5-4.20170523git1f38ac.el7              centos-openstack-pike               32 k
 python2-weakrefmethod                      noarch              1.0.2-3.el7                              centos-openstack-pike               13 k
 zeromq                                     x86_64              4.0.5-4.el7                              centos-openstack-pike              434 k

Transaction Summary
==================================================================================================================================================
Install  1 Package (+28 Dependent packages)

Total download size: 8.3 M
Installed size: 41 M
Downloading packages:
(1/29): c-ares-1.10.0-3.el7.x86_64.rpm                                                                                     |  78 kB  00:00:02     
(2/29): libev-4.15-7.el7.x86_64.rpm                                                                                        |  44 kB  00:00:02     
(3/29): openpgm-5.2.122-2.el7.x86_64.rpm                                                                                   | 172 kB  00:00:04     
(4/29): openstack-neutron-common-11.0.1-1.el7.noarch.rpm                                                                   | 246 kB  00:00:04     
(5/29): openstack-neutron-linuxbridge-11.0.1-1.el7.noarch.rpm                                                              |  14 kB  00:00:00     
(6/29): python-httplib2-0.9.2-1.el7.noarch.rpm                                                                             | 115 kB  00:00:00     
(7/29): python-logutils-0.3.3-3.el7.noarch.rpm                                                                             |  42 kB  00:00:00     
(8/29): python-beautifulsoup4-4.6.0-1.el7.noarch.rpm                                                                       | 171 kB  00:00:02     
(9/29): python-neutron-lib-1.9.1-1.el7.noarch.rpm                                                                          | 170 kB  00:00:00     
(10/29): python-ryu-common-4.15-1.el7.noarch.rpm                                                                           |  51 kB  00:00:00     
(11/29): python-simplegeneric-0.8-7.el7.noarch.rpm                                                                         |  12 kB  00:00:00     
(12/29): python-waitress-0.8.9-5.el7.noarch.rpm                                                                            | 152 kB  00:00:00     
(13/29): python-webtest-2.0.23-1.el7.noarch.rpm                                                                            |  84 kB  00:00:00     
(14/29): python-neutron-11.0.1-1.el7.noarch.rpm                                                                            | 2.0 MB  00:00:03     
(15/29): python-werkzeug-0.9.1-2.el7.noarch.rpm                                                                            | 562 kB  00:00:00     
(16/29): python2-designateclient-2.7.0-1.el7.noarch.rpm                                                                    | 112 kB  00:00:00     
(17/29): python-zmq-14.7.0-2.el7.x86_64.rpm                                                                                | 495 kB  00:00:01     
(18/29): python2-gevent-1.1.2-2.el7.x86_64.rpm                                                                             | 443 kB  00:00:00     
(19/29): python2-os-xenapi-0.2.0-1.el7.noarch.rpm                                                                          |  35 kB  00:00:00     
(20/29): python2-openvswitch-2.7.2-3.1fc27.el7.noarch.rpm                                                                  | 166 kB  00:00:00     
(21/29): python2-osprofiler-1.11.0-1.el7.noarch.rpm                                                                        | 114 kB  00:00:00     
(22/29): python2-ovsdbapp-0.4.0-1.el7.noarch.rpm                                                                           |  46 kB  00:00:00     
(23/29): python2-pecan-1.1.2-1.el7.noarch.rpm                                                                              | 268 kB  00:00:00     
(24/29): python2-pyroute2-0.4.19-1.el7.noarch.rpm                                                                          | 377 kB  00:00:03     
(25/29): python2-singledispatch-3.4.0.3-4.el7.noarch.rpm                                                                   |  18 kB  00:00:00     
(26/29): python2-tinyrpc-0.5-4.20170523git1f38ac.el7.noarch.rpm                                                            |  32 kB  00:00:00     
(27/29): python2-weakrefmethod-1.0.2-3.el7.noarch.rpm                                                                      |  13 kB  00:00:00     
(28/29): python2-ryu-4.15-1.el7.noarch.rpm                                                                                 | 1.9 MB  00:00:04     
(29/29): zeromq-4.0.5-4.el7.x86_64.rpm                                                                                     | 434 kB  00:00:01     
--------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                             505 kB/s | 8.3 MB  00:00:16     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 1:python2-openvswitch-2.7.2-3.1fc27.el7.noarch                                                                                1/29 
  Installing : python2-ovsdbapp-0.4.0-1.el7.noarch                                                                                           2/29 
  Installing : python-beautifulsoup4-4.6.0-1.el7.noarch                                                                                      3/29 
  Installing : python-httplib2-0.9.2-1.el7.noarch                                                                                            4/29 
  Installing : python2-pyroute2-0.4.19-1.el7.noarch                                                                                          5/29 
  Installing : python-waitress-0.8.9-5.el7.noarch                                                                                            6/29 
  Installing : python-webtest-2.0.23-1.el7.noarch                                                                                            7/29 
  Installing : python2-designateclient-2.7.0-1.el7.noarch                                                                                    8/29 
  Installing : python-werkzeug-0.9.1-2.el7.noarch                                                                                            9/29 
  Installing : libev-4.15-7.el7.x86_64                                                                                                      10/29 
  Installing : python2-osprofiler-1.11.0-1.el7.noarch                                                                                       11/29 
  Installing : python-neutron-lib-1.9.1-1.el7.noarch                                                                                        12/29 
  Installing : c-ares-1.10.0-3.el7.x86_64                                                                                                   13/29 
  Installing : python2-gevent-1.1.2-2.el7.x86_64                                                                                            14/29 
  Installing : python2-weakrefmethod-1.0.2-3.el7.noarch                                                                                     15/29 
  Installing : openpgm-5.2.122-2.el7.x86_64                                                                                                 16/29 
  Installing : zeromq-4.0.5-4.el7.x86_64                                                                                                    17/29 
  Installing : python-zmq-14.7.0-2.el7.x86_64                                                                                               18/29 
  Installing : python2-tinyrpc-0.5-4.20170523git1f38ac.el7.noarch                                                                           19/29 
  Installing : python-logutils-0.3.3-3.el7.noarch                                                                                           20/29 
  Installing : python-ryu-common-4.15-1.el7.noarch                                                                                          21/29 
  Installing : python2-ryu-4.15-1.el7.noarch                                                                                                22/29 
  Installing : python2-os-xenapi-0.2.0-1.el7.noarch                                                                                         23/29 
  Installing : python2-singledispatch-3.4.0.3-4.el7.noarch                                                                                  24/29 
  Installing : python-simplegeneric-0.8-7.el7.noarch                                                                                        25/29 
  Installing : python2-pecan-1.1.2-1.el7.noarch                                                                                             26/29 
  Installing : 1:python-neutron-11.0.1-1.el7.noarch                                                                                         27/29 
  Installing : 1:openstack-neutron-common-11.0.1-1.el7.noarch                                                                               28/29 
  Installing : 1:openstack-neutron-linuxbridge-11.0.1-1.el7.noarch                                                                          29/29 
  Verifying  : 1:python2-openvswitch-2.7.2-3.1fc27.el7.noarch                                                                                1/29 
  Verifying  : python-simplegeneric-0.8-7.el7.noarch                                                                                         2/29 
  Verifying  : python2-singledispatch-3.4.0.3-4.el7.noarch                                                                                   3/29 
  Verifying  : python2-tinyrpc-0.5-4.20170523git1f38ac.el7.noarch                                                                            4/29 
  Verifying  : python2-os-xenapi-0.2.0-1.el7.noarch                                                                                          5/29 
  Verifying  : python-ryu-common-4.15-1.el7.noarch                                                                                           6/29 
  Verifying  : python-logutils-0.3.3-3.el7.noarch                                                                                            7/29 
  Verifying  : 1:openstack-neutron-linuxbridge-11.0.1-1.el7.noarch                                                                           8/29 
  Verifying  : python2-ryu-4.15-1.el7.noarch                                                                                                 9/29 
  Verifying  : python2-ovsdbapp-0.4.0-1.el7.noarch                                                                                          10/29 
  Verifying  : 1:python-neutron-11.0.1-1.el7.noarch                                                                                         11/29 
  Verifying  : openpgm-5.2.122-2.el7.x86_64                                                                                                 12/29 
  Verifying  : python2-weakrefmethod-1.0.2-3.el7.noarch                                                                                     13/29 
  Verifying  : c-ares-1.10.0-3.el7.x86_64                                                                                                   14/29 
  Verifying  : python-neutron-lib-1.9.1-1.el7.noarch                                                                                        15/29 
  Verifying  : zeromq-4.0.5-4.el7.x86_64                                                                                                    16/29 
  Verifying  : python2-pecan-1.1.2-1.el7.noarch                                                                                             17/29 
  Verifying  : python2-osprofiler-1.11.0-1.el7.noarch                                                                                       18/29 
  Verifying  : libev-4.15-7.el7.x86_64                                                                                                      19/29 
  Verifying  : python-werkzeug-0.9.1-2.el7.noarch                                                                                           20/29 
  Verifying  : python-zmq-14.7.0-2.el7.x86_64                                                                                               21/29 
  Verifying  : python2-designateclient-2.7.0-1.el7.noarch                                                                                   22/29 
  Verifying  : python-webtest-2.0.23-1.el7.noarch                                                                                           23/29 
  Verifying  : python-waitress-0.8.9-5.el7.noarch                                                                                           24/29 
  Verifying  : python2-pyroute2-0.4.19-1.el7.noarch                                                                                         25/29 
  Verifying  : 1:openstack-neutron-common-11.0.1-1.el7.noarch                                                                               26/29 
  Verifying  : python-httplib2-0.9.2-1.el7.noarch                                                                                           27/29 
  Verifying  : python-beautifulsoup4-4.6.0-1.el7.noarch                                                                                     28/29 
  Verifying  : python2-gevent-1.1.2-2.el7.x86_64                                                                                            29/29 

Installed:
  openstack-neutron-linuxbridge.noarch 1:11.0.1-1.el7                                                                                             

Dependency Installed:
  c-ares.x86_64 0:1.10.0-3.el7                                         libev.x86_64 0:4.15-7.el7                                                  
  openpgm.x86_64 0:5.2.122-2.el7                                       openstack-neutron-common.noarch 1:11.0.1-1.el7                             
  python-beautifulsoup4.noarch 0:4.6.0-1.el7                           python-httplib2.noarch 0:0.9.2-1.el7                                       
  python-logutils.noarch 0:0.3.3-3.el7                                 python-neutron.noarch 1:11.0.1-1.el7                                       
  python-neutron-lib.noarch 0:1.9.1-1.el7                              python-ryu-common.noarch 0:4.15-1.el7                                      
  python-simplegeneric.noarch 0:0.8-7.el7                              python-waitress.noarch 0:0.8.9-5.el7                                       
  python-webtest.noarch 0:2.0.23-1.el7                                 python-werkzeug.noarch 0:0.9.1-2.el7                                       
  python-zmq.x86_64 0:14.7.0-2.el7                                     python2-designateclient.noarch 0:2.7.0-1.el7                               
  python2-gevent.x86_64 0:1.1.2-2.el7                                  python2-openvswitch.noarch 1:2.7.2-3.1fc27.el7                             
  python2-os-xenapi.noarch 0:0.2.0-1.el7                               python2-osprofiler.noarch 0:1.11.0-1.el7                                   
  python2-ovsdbapp.noarch 0:0.4.0-1.el7                                python2-pecan.noarch 0:1.1.2-1.el7                                         
  python2-pyroute2.noarch 0:0.4.19-1.el7                               python2-ryu.noarch 0:4.15-1.el7                                            
  python2-singledispatch.noarch 0:3.4.0.3-4.el7                        python2-tinyrpc.noarch 0:0.5-4.20170523git1f38ac.el7                       
  python2-weakrefmethod.noarch 0:1.0.2-3.el7                           zeromq.x86_64 0:4.0.5-4.el7                                                

Complete!
```

### Configure

Backup "neutron.conf"
```
[vagrant@localhost ~]$ sudo cp /etc/neutron/neutron.conf etc0x2Fneutron0x2Fneutron.conf
```

Sed "neutron.conf"
```
[vagrant@localhost ~]$ sudo sed   's%^\[DEFAULT\]$%&\ndebug=true\nverbose=true\ntransport_url=rabbit://openstack:RABBIT_PASS@${controller}\n#my_ip=${my_ip}\nauth_strategy=keystone\n%;s%^\[keystone_authtoken\]$%&\nauth_uri=http://${controller}:5000\nauth_url=http://${controller}:35357\nmemcached_servers=${controller}:11211\nauth_type=password\nproject_domain_name=default\nuser_domain_name=default\nproject_name=service\nusername=neutron\npassword=SERVICE_PASS\n%;s%^\[oslo_concurrency\]$%&\nlock_path=/var/lib/neutron/tmp\n%' etc0x2Fneutron0x2Fneutron.conf | env controller=10.64.33.64 my_ip=10.64.33.65 envsubst > neutron.conf
```

Check "neutron.conf"
```
[vagrant@localhost ~]$ sudo cat neutron.conf | egrep '^[^#]'
[DEFAULT]
debug=true
verbose=true
transport_url=rabbit://openstack:RABBIT_PASS@10.64.33.64
auth_strategy=keystone
[agent]
[cors]
[database]
[keystone_authtoken]
auth_uri=http://10.64.33.64:5000
auth_url=http://10.64.33.64:35357
memcached_servers=10.64.33.64:11211
auth_type=password
project_domain_name=default
user_domain_name=default
project_name=service
username=neutron
password=SERVICE_PASS
[matchmaker_redis]
[nova]
[oslo_concurrency]
lock_path=/var/lib/neutron/tmp
[oslo_messaging_amqp]
[oslo_messaging_kafka]
[oslo_messaging_notifications]
[oslo_messaging_rabbit]
[oslo_messaging_zmq]
[oslo_middleware]
[oslo_policy]
[quotas]
[ssl]
```

Modify "neutron.conf"
```
[vagrant@localhost ~]$ sudo cp neutron.conf /etc/neutron/neutron.conf 
```

Backup "linuxbridge_agent.ini"
```
[vagrant@localhost ~]$ sudo cp /etc/neutron/plugins/ml2/linuxbridge_agent.ini etc0x2Fneutron0x2Fplugins0x2Fml20x2Flinuxbridge_agent.ini
```

Sed "linuxbridge_agent.ini"
```
[vagrant@localhost ~]$ sudo sed 's/^\[linux_bridge\]$/&\nphysical_interface_mappings=provider:eth2\n/;s/^\[vxlan\]$/&\nenable_vxlan=false\n/;s/^\[securitygroup\]$/&\nenable_security_group=true\nfirewall_driver=neutron.agent.linux.iptables_firewall.IptablesFirewallDriver\n/' etc0x2Fneutron0x2Fplugins0x2Fml20x2Flinuxbridge_agent.ini > linuxbridge_agent.ini.provider-networks
```

Check "linuxbridge_agent.ini"
```
[vagrant@localhost ~]$ sudo cat linuxbridge_agent.ini.provider-networks | egrep '^[^#]'
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

Modify "linuxbridge_agent.ini"
```
[vagrant@localhost ~]$ sudo cp linuxbridge_agent.ini.provider-networks /etc/neutron/plugins/ml2/linuxbridge_agent.ini 
```


### Re-configure Nova

Sed "nova.conf"
```
[vagrant@localhost ~]$ sudo sed 's%^\[neutron]$%&\nurl=http://${controller}:9696\nauth_url=http://${controller}:35357\nauth_type=password\nproject_domain_name=default\nuser_domain_name=default\nregion_name=RegionOne\nproject_name=service\nusername=neutron\npassword=SERVICE_PASS\n%' nova.conf | env controller=10.64.33.64 envsubst > nova.conf.neutron
```

Check "nova.conf"
```
[vagrant@localhost ~]$ sudo cat nova.conf.neutron | egrep '^[^#]'
[DEFAULT]
debug=true
verbose=true
enabled_apis=osapi_compute,metadata
transport_url=rabbit://openstack:RABBIT_PASS@10.64.33.64
my_ip=10.64.33.65
user_neutron=True
firewall_driver=nova.virt.firewall.NoopFirewallDriver
[api]
auth_strategy=keystone
[api_database]
[barbican]
[cache]
[cells]
[cinder]
[compute]
[conductor]
[console]
[consoleauth]
[cors]
[crypto]
[database]
[ephemeral_storage_encryption]
[filter_scheduler]
[glance]
api_servers=http://10.64.33.64:9292
[guestfs]
[healthcheck]
[hyperv]
[ironic]
[key_manager]
[keystone_authtoken]
auth_uri=http://10.64.33.64:5000
auth_url=http://10.64.33.64:35357
memcached_servers=10.64.33.64:11211
auth_type=password
project_domain_name=default
user_domain_name=default
project_name=service
username=nova
password=SERVICE_PASS
[libvirt]
virt_type=qemu
[matchmaker_redis]
[metrics]
[mks]
[neutron]
url=http://10.64.33.64:9696
auth_url=http://10.64.33.64:35357
auth_type=password
project_domain_name=default
user_domain_name=default
region_name=RegionOne
project_name=service
username=neutron
password=SERVICE_PASS
[notifications]
[osapi_v21]
[oslo_concurrency]
lock_path=/var/lib/nova/tmp
[oslo_messaging_amqp]
[oslo_messaging_kafka]
[oslo_messaging_notifications]
[oslo_messaging_rabbit]
[oslo_messaging_zmq]
[oslo_middleware]
[oslo_policy]
[pci]
[placement]
os_region_name=RegionOne
project_domain_name=Default
project_name=service
auth_type=password
user_domain_name=Default
auth_url=http://10.64.33.64:35357/v3
username=placement
password=SERVICE_PASS
[quota]
[rdp]
[remote_debug]
[scheduler]
[serial_console]
[service_user]
[spice]
[trusted_computing]
[upgrade_levels]
[vendordata_dynamic_auth]
[vmware]
[vnc]
enabled=True
vncserver_listen=0.0.0.0
vncserver_proxyclient_address=10.64.33.65
novncproxy_base_url=http://10.64.33.64:6080/vnc_auto.html
[workarounds]
[wsgi]
[xenserver]
[xvp]
```

Modify "nova.conf"
```
[vagrant@localhost ~]$ sudo cp nova.conf.neutron /etc/nova/nova.conf
```

### Start

Nova Compute
```
[vagrant@localhost ~]$ sudo systemctl restart openstack-nova-compute.service
[vagrant@localhost ~]$ sudo systemctl -l status openstack-nova-compute.service
● openstack-nova-compute.service - OpenStack Nova Compute Server
   Loaded: loaded (/usr/lib/systemd/system/openstack-nova-compute.service; enabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 23:17:56 UTC; 19s ago
 Main PID: 4485 (nova-compute)
   CGroup: /system.slice/openstack-nova-compute.service
           └─4485 /usr/bin/python2 /usr/bin/nova-compute

Oct 21 23:17:54 localhost.localdomain systemd[1]: Starting OpenStack Nova Compute Server...
Oct 21 23:17:56 localhost.localdomain systemd[1]: Started OpenStack Nova Compute Server.
[vagrant@localhost ~]$ sudo tail /var/log/nova/nova-compute.log
2017-10-21 23:17:57.257 4485 DEBUG oslo_concurrency.lockutils [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] Lock "compute_resources" acquired by "nova.compute.resource_tracker._update_available_resource" :: waited 0.000s inner /usr/lib/python2.7/site-packages/oslo_concurrency/lockutils.py:270
2017-10-21 23:17:57.802 4485 DEBUG nova.scheduler.client.report [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] Grabbing aggregate associations for resource provider 61a37417-a33f-4512-ad04-1c218db039e6 _ensure_resource_provider /usr/lib/python2.7/site-packages/nova/scheduler/client/report.py:510
2017-10-21 23:17:58.225 4485 DEBUG nova.compute.resource_tracker [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] Total usable vcpus: 1, total allocated vcpus: 0 _report_final_resource_view /usr/lib/python2.7/site-packages/nova/compute/resource_tracker.py:792
2017-10-21 23:17:58.225 4485 INFO nova.compute.resource_tracker [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] Final resource view: name=localhost.localdomain phys_ram=2047MB used_ram=512MB phys_disk=37GB used_disk=0GB total_vcpus=1 used_vcpus=0 pci_stats=[]
2017-10-21 23:17:58.243 4485 DEBUG nova.scheduler.client.report [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] Refreshing aggregate associations for resource provider 61a37417-a33f-4512-ad04-1c218db039e6 _ensure_resource_provider /usr/lib/python2.7/site-packages/nova/scheduler/client/report.py:498
2017-10-21 23:17:58.555 4485 DEBUG nova.compute.resource_tracker [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] Compute_service record updated for localhost.localdomain:localhost.localdomain _update_available_resource /usr/lib/python2.7/site-packages/nova/compute/resource_tracker.py:732
2017-10-21 23:17:58.556 4485 DEBUG oslo_concurrency.lockutils [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] Lock "compute_resources" released by "nova.compute.resource_tracker._update_available_resource" :: held 1.298s inner /usr/lib/python2.7/site-packages/oslo_concurrency/lockutils.py:282
2017-10-21 23:17:58.556 4485 DEBUG nova.service [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] Creating RPC server for service compute start /usr/lib/python2.7/site-packages/nova/service.py:166
2017-10-21 23:17:58.573 4485 DEBUG nova.service [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] Join ServiceGroup membership for this service compute start /usr/lib/python2.7/site-packages/nova/service.py:184
2017-10-21 23:17:58.574 4485 DEBUG nova.servicegroup.drivers.db [req-7efc7c63-2fc1-4141-bb09-74707491a12b - - - - -] DB_Driver: join new ServiceGroup member localhost.localdomain to the compute group, service = <Service: host=localhost.localdomain, binary=nova-compute, manager_class_name=nova.compute.manager.ComputeManager> join /usr/lib/python2.7/site-packages/nova/servicegroup/drivers/db.py:47
```

Neutron linux bridge agent
```
[vagrant@localhost ~]$ sudo systemctl start neutron-linuxbridge-agent.service
[vagrant@localhost ~]$ sudo systemctl -l status neutron-linuxbridge-agent.service
● neutron-linuxbridge-agent.service - OpenStack Neutron Linux Bridge Agent
   Loaded: loaded (/usr/lib/systemd/system/neutron-linuxbridge-agent.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 23:19:31 UTC; 10s ago
  Process: 4553 ExecStartPre=/usr/bin/neutron-enable-bridge-firewall.sh (code=exited, status=0/SUCCESS)
 Main PID: 4560 (neutron-linuxbr)
   CGroup: /system.slice/neutron-linuxbridge-agent.service
           ├─4560 /usr/bin/python2 /usr/bin/neutron-linuxbridge-agent --config-file /usr/share/neutron/neutron-dist.conf --config-file /etc/neutron/neutron.conf --config-file /etc/neutron/plugins/ml2/linuxbridge_agent.ini --config-dir /etc/neutron/conf.d/common --config-dir /etc/neutron/conf.d/neutron-linuxbridge-agent --log-file /var/log/neutron/linuxbridge-agent.log
           ├─4573 sudo neutron-rootwrap-daemon /etc/neutron/rootwrap.conf
           └─4574 /usr/bin/python2 /usr/bin/neutron-rootwrap-daemon /etc/neutron/rootwrap.conf

Oct 21 23:19:31 localhost.localdomain systemd[1]: Starting OpenStack Neutron Linux Bridge Agent...
Oct 21 23:19:31 localhost.localdomain neutron-enable-bridge-firewall.sh[4553]: net.bridge.bridge-nf-call-iptables = 1
Oct 21 23:19:31 localhost.localdomain neutron-enable-bridge-firewall.sh[4553]: net.bridge.bridge-nf-call-ip6tables = 1
Oct 21 23:19:31 localhost.localdomain systemd[1]: Started OpenStack Neutron Linux Bridge Agent.
Oct 21 23:19:32 localhost.localdomain neutron-linuxbridge-agent[4560]: Guru meditation now registers SIGUSR1 and SIGUSR2 by default for backward compatibility. SIGUSR1 will no longer be registered in a future release, so please use SIGUSR2 to generate reports.
Oct 21 23:19:33 localhost.localdomain sudo[4573]:  neutron : TTY=unknown ; PWD=/ ; USER=root ; COMMAND=/bin/neutron-rootwrap-daemon /etc/neutron/rootwrap.conf
[vagrant@localhost ~]$ sudo tail /var/log/neutron/linuxbridge-agent.log
2017-10-21 23:19:33.439 4560 DEBUG neutron_lib.callbacks.manager [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Subscribe: <bound method LinuxBridgeTrunkDriver.agent_port_change of <neutron.services.trunk.drivers.linuxbridge.agent.driver.LinuxBridgeTrunkDriver object at 0x47d5bd0>> port_device after_update subscribe /usr/lib/python2.7/site-packages/neutron_lib/callbacks/manager.py:41
2017-10-21 23:19:33.440 4560 DEBUG neutron_lib.callbacks.manager [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Subscribe: <bound method LinuxBridgeTrunkDriver.agent_port_delete of <neutron.services.trunk.drivers.linuxbridge.agent.driver.LinuxBridgeTrunkDriver object at 0x47d5bd0>> port_device after_delete subscribe /usr/lib/python2.7/site-packages/neutron_lib/callbacks/manager.py:41
2017-10-21 23:19:33.440 4560 DEBUG neutron.api.rpc.callbacks.resource_manager [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Registering callback for Trunk register /usr/lib/python2.7/site-packages/neutron/api/rpc/callbacks/resource_manager.py:64
2017-10-21 23:19:33.441 4560 DEBUG neutron.api.rpc.callbacks.resource_manager [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Registering callback for SubPort register /usr/lib/python2.7/site-packages/neutron/api/rpc/callbacks/resource_manager.py:64
2017-10-21 23:19:33.577 4560 INFO neutron.plugins.ml2.drivers.agent._common_agent [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Linux bridge agent Agent RPC Daemon Started!
2017-10-21 23:19:33.577 4560 INFO neutron.plugins.ml2.drivers.agent._common_agent [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Linux bridge agent Agent out of sync with plugin!
2017-10-21 23:19:33.578 4560 DEBUG oslo_concurrency.lockutils [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Lock "ebtables" acquired by "neutron.plugins.ml2.drivers.linuxbridge.agent.arp_protect.delete_unreferenced_arp_protection" :: waited 0.000s inner /usr/lib/python2.7/site-packages/oslo_concurrency/lockutils.py:270
2017-10-21 23:19:33.578 4560 DEBUG neutron.agent.linux.utils [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Running command (rootwrap daemon): ['ebtables', '--concurrent', '-L'] execute_rootwrap_daemon /usr/lib/python2.7/site-packages/neutron/agent/linux/utils.py:108
2017-10-21 23:19:33.592 4560 INFO neutron.plugins.ml2.drivers.linuxbridge.agent.arp_protect [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Clearing orphaned ARP spoofing entries for devices []
2017-10-21 23:19:33.592 4560 DEBUG oslo_concurrency.lockutils [req-f0cecd31-4951-4cb2-92c0-0aa112a890d8 - - - - -] Lock "ebtables" released by "neutron.plugins.ml2.drivers.linuxbridge.agent.arp_protect.delete_unreferenced_arp_protection" :: held 0.014s inner /usr/lib/python2.7/site-packages/oslo_concurrency/lockutils.py:282
```

Auto Starting
```
[vagrant@localhost ~]$ sudo systemctl enable neutron-linuxbridge-agent.service
Created symlink from /etc/systemd/system/multi-user.target.wants/neutron-linuxbridge-agent.service to /usr/lib/systemd/system/neutron-linuxbridge-agent.service.
````



### Verifying

From controller
```
[vagrant@localhost ~]$ openstack network agent list
+--------------------------------------+--------------------+-----------------------+-------------------+-------+-------+---------------------------+
| ID                                   | Agent Type         | Host                  | Availability Zone | Alive | State | Binary                    |
+--------------------------------------+--------------------+-----------------------+-------------------+-------+-------+---------------------------+
| 59d59396-ebd2-4f3e-a2da-387cbcac2027 | Metadata agent     | localhost.localdomain | None              | :-)   | UP    | neutron-metadata-agent    |
| 88e03f05-8d37-423b-b7be-1d74d166d408 | DHCP agent         | localhost.localdomain | nova              | :-)   | UP    | neutron-dhcp-agent        |
| a69362b5-9511-4e37-9c98-75babbf5ca2f | Linux bridge agent | compute-10-64-33-65   | None              | :-)   | UP    | neutron-linuxbridge-agent |
| ea6174e3-62d2-48f4-a268-236883c00f9d | Linux bridge agent | localhost.localdomain | None              | :-)   | UP    | neutron-linuxbridge-agent |
+--------------------------------------+--------------------+-----------------------+-------------------+-------+-------+---------------------------+
```
