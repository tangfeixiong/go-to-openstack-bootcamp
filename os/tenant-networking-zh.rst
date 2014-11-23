简要设置OpenStack租户网络
=========================
用软件定义你的私有网络（SDN，Software-Defined Network）。

设计内网
--------
在OpenStack仪表盘，选择Network菜单，已经有一个命名类似public的网络，理解为联通/电信/移动给你家接入了宽带

    1. 为你家的网取个名
            .. image:: /os/image/demo-tenant-private-net-naming.png
    2. 设计你家的局域网
        设置局域网IP地址，这里使用10.0.0.0/8的A类私有地址（http://baike.baidu.com/view/39496.htm）
            .. image:: /os/image/demo-tenant-private-subnet-sum.png
    
	    设置DHCP和DNS，这样计算机都能够傻瓜式上网
	        .. image:: /os/image/demo-tenant-private-subnet-auto.png
    
	    一般的私有网就这样了
	        .. image:: /os/image/demo-tenant-private-network.png
	
设置虚拟路由器
--------------
在Router菜单，创建一个路由器

    1. 把入户的宽带插到WAN口
        .. image:: /os/image/demo-tenant-router-gateway.png
    2. 把你家的私有网络连到虚拟路由器
        .. image:: /os/image/demo-tenant-router-lan.png

在Networking Topology菜单，查看网络效果
    .. image:: /os/image/demo-tenant-networking-topology.png
	
检查网络工作能力
----------------
作为租户，自然是创建一台虚拟机，然后试试能不能上网。

而作为OpenStack私有云网管，要负责检查租户的虚拟网络设施：在OpenStack网络节点上进行CLI操作，通过路由器的ID，在该路由器上操作网络命令	

.. code:: bash
[root@cfhost01 ~]# . keystonerc_admin

[root@cfhost01 ~(keystone_admin)]# neutron router-list
+--------------------------------------+-------------+-----------------------------------------------------------------------------+
| id                                   | name        | external_gateway_info                                                       |
+--------------------------------------+-------------+-----------------------------------------------------------------------------+
| 961b4753-ba8a-492f-9176-9970c72a9bec | ...         | {"network_id": "b1aa5d97-c0f5-46f6-a18e-533044e214d2", "enable_snat": true} |
| c9611b53-4ac9-4fdc-9ac4-41aa67e40dff | ...         | {"network_id": "b1aa5d97-c0f5-46f6-a18e-533044e214d2", "enable_snat": true} |
| e6f8e625-6ab0-43dc-bc21-59dea0e30992 | router1     | {"network_id": "b1aa5d97-c0f5-46f6-a18e-533044e214d2", "enable_snat": true} |
| fd676abc-c561-40ca-a708-4eed5a57b1dc | routerAdmin | {"network_id": "b1aa5d97-c0f5-46f6-a18e-533044e214d2", "enable_snat": true} |
+--------------------------------------+-------------+-----------------------------------------------------------------------------+
	
[root@cfhost01 ~(keystone_admin)]# ip netns exec qrouter-e6f8e625-6ab0-43dc-bc21-59dea0e30992 ping -c 3 www.baidu.com

PING www.a.shifen.com (61.135.169.121) 56(84) bytes of data.

64 bytes from 61.135.169.121: icmp_seq=1 ttl=54 time=2.82 ms

64 bytes from 61.135.169.121: icmp_seq=2 ttl=54 time=2.74 ms

64 bytes from 61.135.169.121: icmp_seq=3 ttl=54 time=2.78 ms


--- www.a.shifen.com ping statistics ---

3 packets transmitted, 3 received, 0% packet loss, time 2005ms

rtt min/avg/max/mdev = 2.747/2.787/2.826/0.053 ms

[root@cfhost01 ~(keystone_admin)]# ip netns exec qrouter-e6f8e625-6ab0-43dc-bc21-59dea0e30992 ping -c 3 192.168.70.162

PING 192.168.70.162 (192.168.70.162) 56(84) bytes of data.

64 bytes from 192.168.70.162: icmp_seq=1 ttl=64 time=1.35 ms

64 bytes from 192.168.70.162: icmp_seq=2 ttl=64 time=0.501 ms

64 bytes from 192.168.70.162: icmp_seq=3 ttl=64 time=0.068 ms


--- 192.168.70.162 ping statistics ---

3 packets transmitted, 3 received, 0% packet loss, time 2001ms

rtt min/avg/max/mdev = 0.068/0.640/1.351/0.532 ms

[root@cfhost01 ~]# ip netns exec qrouter-e6f8e625-6ab0-43dc-bc21-59dea0e30992 ip addr

203: lo: <LOOPBACK,UP,LOWER_UP> mtu 16436 qdisc noqueue state UNKNOWN
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
3542: qr-bbba4328-e1: <BROADCAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN
    link/ether fa:16:3e:34:42:61 brd ff:ff:ff:ff:ff:ff
    inet 10.0.0.1/24 brd 10.0.0.255 scope global qr-bbba4328-e1
    inet6 fe80::f816:3eff:fe34:4261/64 scope link
       valid_lft forever preferred_lft forever
3561: qg-ba49d758-77: <BROADCAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UNKNOWN
    link/ether fa:16:3e:cb:1c:8c brd ff:ff:ff:ff:ff:ff
    inet 192.168.74.143/24 brd 192.168.74.255 scope global qg-ba49d758-77
    inet6 fe80::f816:3eff:fecb:1c8c/64 scope link
       valid_lft forever preferred_lft forever

[root@cfhost01 ~]# ip netns exec qrouter-e6f8e625-6ab0-43dc-bc21-59dea0e30992 ip route

192.168.70.0/24 via 192.168.74.2 dev qg-ba49d758-77

10.0.0.0/24 dev qr-bbba4328-e1  proto kernel  scope link  src 10.0.0.1

192.168.74.0/24 dev qg-ba49d758-77  proto kernel  scope link  src 192.168.74.143

default via 192.168.74.1 dev qg-ba49d758-77

[root@cfhost01 ~]# ip netns exec qrouter-e6f8e625-6ab0-43dc-bc21-59dea0e30992 route -n

Kernel IP routing table

Destination     Gateway         Genmask         Flags Metric Ref    Use Iface

192.168.70.0    192.168.74.2    255.255.255.0   UG    0      0        0 qg-ba49d758-77

10.0.0.0        0.0.0.0         255.255.255.0   U     0      0        0 qr-bbba4328-e1

192.168.74.0    0.0.0.0         255.255.255.0   U     0      0        0 qg-ba49d758-77

0.0.0.0         192.168.74.1    0.0.0.0         UG    0      0        0 qg-ba49d758-77	   
	   
[root@cfhost01 ~]# ip netns exec qrouter-e6f8e625-6ab0-43dc-bc21-59dea0e30992 traceroute 8.8.4.4

traceroute to 8.8.4.4 (8.8.4.4), 30 hops max, 60 byte packets

 1  192.168.74.1 (192.168.74.1)  0.805 ms  0.713 ms  0.647 ms
 
 2  ???.???.??.?? (???.???.??.??)  1.439 ms  2.482 ms  2.456 ms
 
 3  * * *
 
 4  61.49.163.129 (61.49.163.129)  3.958 ms  3.953 ms  3.946 ms
 
 5  61.49.163.129 (61.49.163.129)  3.910 ms  3.940 ms  3.870 ms
 
 6  61.148.156.77 (61.148.156.77)  9.666 ms  9.427 ms  5.722 ms
 
 7  61.148.158.245 (61.148.158.245)  6.270 ms  14.660 ms  14.591 ms
 
 8  123.126.0.141 (123.126.0.141)  14.555 ms  14.562 ms  14.646 ms
 
 9  219.158.105.246 (219.158.105.246)  36.527 ms  36.600 ms  36.486 ms
 
10  219.158.96.226 (219.158.96.226)  49.524 ms 219.158.23.6 (219.158.23.6)  37.177 ms  37.184 ms

11  219.158.97.30 (219.158.97.30)  88.343 ms  88.337 ms 219.158.96.246 (219.158.96.246)  81.815 ms

12  219.158.29.54 (219.158.29.54)  106.702 ms  106.706 ms 219.158.3.238 (219.158.3.238)  82.381 ms

13  72.14.215.130 (72.14.215.130)  221.355 ms  225.325 ms  225.218 ms

14  209.85.255.241 (209.85.255.241)  287.971 ms 209.85.255.243 (209.85.255.243)  287.395 ms 209.85.255.237 (209.85.255.237)  241.909 ms

15  google-public-dns-b.google.com (8.8.4.4)  275.116 ms  281.573 ms  277.861 ms	   
	   
[root@cfhost01 ~]# ip netns exec qrouter-e6f8e625-6ab0-43dc-bc21-59dea0e30992 nslookup 8.8.8.8

Server:         114.114.114.114

Address:        114.114.114.114#53


Non-authoritative answer:

8.8.8.8.in-addr.arpa    name = google-public-dns-a.google.com.

Authoritative answers can be found from:

	   
更多请参考
----------
    #. http://docs.openstack.org/user-guide/content/dashboard_create_networks.html
    #. https://openstack.redhat.com/Networking_in_too_much_detail
