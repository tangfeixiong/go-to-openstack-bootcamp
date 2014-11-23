OpenStack租户网络建设简要
=========================
用软件定义（SDN，Software-Defined Network）你自己的私有网络。
设计内网
--------
在OpenStack仪表盘，选择Network菜单，一般已经有一个命名类似public的网络，理解为联通/电信/移动已为你家里接进了Internet网
1. 给你家的网取个名
    .. image:: /os/image/demo-tenant-private-net-naming.png
2. 设计你家的局域网
    设置局域网IP地址，这里使用10.0.0.0/8的A类私有地址（http://baike.baidu.com/view/39496.htm）
    .. image:: /os/image/demo-tenant-private-subnet-summary.png
    设置DHCP和DNS，这样家里的计算机都能够傻瓜式上网
	.. image:: /os/image/demo-tenant-private-subnet-auto.png
    一般的私有网就是这样了
	.. image:: /os/image/demo-tenant-private-network.png
设置虚拟路由器
--------------
在Router菜单，创建一个路由器
1. 把入户的Internet网线插到WAN口
    .. image:: /os/image/demo-tenant-router-gateway.png
2. 把你家的私有网络连到虚拟路由器
    .. image:: /os/image/demo-tenant-router-lan.png
在Networking Topology菜单，查看网络效果
    .. image:: /os/image/demo-tenant-networking-topology.png
检查网络工作能力
----------------

