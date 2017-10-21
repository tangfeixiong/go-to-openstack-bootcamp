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
[vagrant@localhost ~]$ cat /etc/sysconfig/network-scripts/ifcfg-
ifcfg-eth0                ifcfg-lo                  ifcfg-Wired_connection_1  
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
