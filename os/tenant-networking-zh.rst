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

而作为OpenStack私有云网管，要负责检查租户的虚拟网络设施：

    1. 通过路由器的ID，在OpenStack网络节点上进行CLI操作	
	2. 在该路由器上操作网络命令
	
更多请参考
----------
    #. http://docs.openstack.org/user-guide/content/dashboard_create_networks.html
    #. https://openstack.redhat.com/Networking_in_too_much_detail
