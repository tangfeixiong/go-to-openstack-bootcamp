# OpenStack Pike Installation

## Table of content

控制节点
* [之前章节](./openstack-pike-install-1st-controller.md)
* [之前Keystone部署中重要操作](./openstack-pike-install-2nd-keystone.md#system-project)
* [Glance镜像服务](#image)
    * [在数据库中登记image服务](#service-discovery)
    * [安装及配置](#packages)
    * [启动](#start)
    * [检查日志](#logs)
    * [cirros镜像](#create)

## Controller

### Image

[Glance](https://docs.openstack.org/glance/pike/install/install-rdo.html)


Client
```
[vagrant@localhost ~]$ . openstack-admin.sh 
```

### Service Discovery

创建glance用户
```
[vagrant@localhost ~]$ openstack user create --domain default --password SERVICE_PASS glance 
+---------------------+----------------------------------+
| Field               | Value                            |
+---------------------+----------------------------------+
| domain_id           | default                          |
| enabled             | True                             |
| id                  | a404e7fe5d544719ad71a3e883653409 |
| name                | glance                           |
| options             | {}                               |
| password_expires_at | None                             |
+---------------------+----------------------------------+
```

分配角色
```
[vagrant@localhost ~]$ openstack role add --project service --user glance admin
```

登记服务
```
[vagrant@localhost ~]$ openstack service create --name glance --description "OpenStack Image" image
+-------------+----------------------------------+
| Field       | Value                            |
+-------------+----------------------------------+
| description | OpenStack Image                  |
| enabled     | True                             |
| id          | 842171ace1434340826a17eb48b94302 |
| name        | glance                           |
| type        | image                            |
+-------------+----------------------------------+
```

服务发现
```
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne image public http://10.64.33.64:9292
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | d205c75f37a14b3aa7122ea1de8d6d6d |
| interface    | public                           |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 842171ace1434340826a17eb48b94302 |
| service_name | glance                           |
| service_type | image                            |
| url          | http://10.64.33.64:9292          |
+--------------+----------------------------------+
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne image internal http://10.64.33.64:9292
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | 470edd05f14145b5848609380f260038 |
| interface    | internal                         |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 842171ace1434340826a17eb48b94302 |
| service_name | glance                           |
| service_type | image                            |
| url          | http://10.64.33.64:9292          |
+--------------+----------------------------------+
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne image admin http://10.64.33.64:9292
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | 8a10a8bbda4c49c5b3d6d7857e5af927 |
| interface    | admin                            |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 842171ace1434340826a17eb48b94302 |
| service_name | glance                           |
| service_type | image                            |
| url          | http://10.64.33.64:9292          |
+--------------+----------------------------------+
```

### Packages

Install
```
[vagrant@localhost ~]$ sudo yum install -y openstack-glance
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
Resolving Dependencies
--> Running transaction check
---> Package openstack-glance.noarch 1:15.0.0-1.el7 will be installed
--> Processing Dependency: python-glance = 1:15.0.0-1.el7 for package: 1:openstack-glance-15.0.0-1.el7.noarch
--> Running transaction check
---> Package python-glance.noarch 1:15.0.0-1.el7 will be installed
--> Processing Dependency: python-wsme >= 0.8 for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: python-taskflow >= 2.7.0 for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: python-swiftclient >= 2.2.0 for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-vmware >= 0.11.1 for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: python-os-brick >= 1.8.0 for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: python-glance-store >= 0.21.0 for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: python-retrying for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: python-httplib2 for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: python-cursive for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: python-boto for package: 1:python-glance-15.0.0-1.el7.noarch
--> Processing Dependency: pysendfile for package: 1:python-glance-15.0.0-1.el7.noarch
--> Running transaction check
---> Package pysendfile.x86_64 0:2.0.0-5.el7 will be installed
---> Package python-boto.noarch 0:2.34.0-4.el7 will be installed
--> Processing Dependency: python-rsa for package: python-boto-2.34.0-4.el7.noarch
---> Package python-httplib2.noarch 0:0.9.2-1.el7 will be installed
---> Package python-retrying.noarch 0:1.2.3-4.el7 will be installed
---> Package python2-cursive.noarch 0:0.1.2-1.el7 will be installed
--> Processing Dependency: python-lxml >= 2.3 for package: python2-cursive-0.1.2-1.el7.noarch
--> Processing Dependency: python-castellan >= 0.4.0 for package: python2-cursive-0.1.2-1.el7.noarch
---> Package python2-glance-store.noarch 0:0.22.0-1.el7 will be installed
--> Processing Dependency: python-oslo-privsep >= 1.9.0 for package: python2-glance-store-0.22.0-1.el7.noarch
--> Processing Dependency: python-oslo-rootwrap for package: python2-glance-store-0.22.0-1.el7.noarch
---> Package python2-os-brick.noarch 0:1.15.3-1.el7 will be installed
--> Processing Dependency: python-os-win >= 2.0.0 for package: python2-os-brick-1.15.3-1.el7.noarch
---> Package python2-oslo-vmware.noarch 0:2.23.0-1.el7 will be installed
--> Processing Dependency: python-oslo-vmware-lang = 2.23.0-1.el7 for package: python2-oslo-vmware-2.23.0-1.el7.noarch
--> Processing Dependency: python-suds >= 0.6 for package: python2-oslo-vmware-2.23.0-1.el7.noarch
---> Package python2-swiftclient.noarch 0:3.4.0-1.el7 will be installed
---> Package python2-taskflow.noarch 0:2.14.0-1.el7 will be installed
--> Processing Dependency: python-networkx >= 1.10 for package: python2-taskflow-2.14.0-1.el7.noarch
--> Processing Dependency: python-automaton >= 0.5.0 for package: python2-taskflow-2.14.0-1.el7.noarch
--> Processing Dependency: python-networkx-core for package: python2-taskflow-2.14.0-1.el7.noarch
---> Package python2-wsme.noarch 0:0.9.2-1.el7 will be installed
--> Processing Dependency: python-simplegeneric for package: python2-wsme-0.9.2-1.el7.noarch
--> Running transaction check
---> Package python-lxml.x86_64 0:3.2.1-4.el7 will be installed
--> Processing Dependency: libxslt.so.1(LIBXML2_1.1.9)(64bit) for package: python-lxml-3.2.1-4.el7.x86_64
--> Processing Dependency: libxslt.so.1(LIBXML2_1.1.26)(64bit) for package: python-lxml-3.2.1-4.el7.x86_64
--> Processing Dependency: libxslt.so.1(LIBXML2_1.1.2)(64bit) for package: python-lxml-3.2.1-4.el7.x86_64
--> Processing Dependency: libxslt.so.1(LIBXML2_1.0.24)(64bit) for package: python-lxml-3.2.1-4.el7.x86_64
--> Processing Dependency: libxslt.so.1(LIBXML2_1.0.22)(64bit) for package: python-lxml-3.2.1-4.el7.x86_64
--> Processing Dependency: libxslt.so.1(LIBXML2_1.0.18)(64bit) for package: python-lxml-3.2.1-4.el7.x86_64
--> Processing Dependency: libxslt.so.1(LIBXML2_1.0.11)(64bit) for package: python-lxml-3.2.1-4.el7.x86_64
--> Processing Dependency: libxslt.so.1()(64bit) for package: python-lxml-3.2.1-4.el7.x86_64
--> Processing Dependency: libexslt.so.0()(64bit) for package: python-lxml-3.2.1-4.el7.x86_64
---> Package python-networkx.noarch 0:1.10-1.el7 will be installed
---> Package python-networkx-core.noarch 0:1.10-1.el7 will be installed
--> Processing Dependency: scipy for package: python-networkx-core-1.10-1.el7.noarch
---> Package python-oslo-vmware-lang.noarch 0:2.23.0-1.el7 will be installed
---> Package python-simplegeneric.noarch 0:0.8-7.el7 will be installed
---> Package python2-automaton.noarch 0:1.12.1-1.el7 will be installed
---> Package python2-castellan.noarch 0:0.12.0-1.el7 will be installed
---> Package python2-os-win.noarch 0:2.2.0-1.el7 will be installed
---> Package python2-oslo-privsep.noarch 0:1.22.0-1.el7 will be installed
--> Processing Dependency: python-oslo-privsep-lang = 1.22.0-1.el7 for package: python2-oslo-privsep-1.22.0-1.el7.noarch
---> Package python2-oslo-rootwrap.noarch 0:5.9.0-1.el7 will be installed
---> Package python2-rsa.noarch 0:3.3-2.el7 will be installed
---> Package python2-suds.noarch 0:0.7-0.4.94664ddd46a6.el7 will be installed
--> Running transaction check
---> Package libxslt.x86_64 0:1.1.28-5.el7 will be installed
---> Package python-oslo-privsep-lang.noarch 0:1.22.0-1.el7 will be installed
---> Package python2-scipy.x86_64 0:0.18.0-3.el7 will be installed
--> Processing Dependency: numpy for package: python2-scipy-0.18.0-3.el7.x86_64
--> Processing Dependency: libgfortran.so.3(GFORTRAN_1.4)(64bit) for package: python2-scipy-0.18.0-3.el7.x86_64
--> Processing Dependency: libgfortran.so.3(GFORTRAN_1.0)(64bit) for package: python2-scipy-0.18.0-3.el7.x86_64
--> Processing Dependency: libtatlas.so.3()(64bit) for package: python2-scipy-0.18.0-3.el7.x86_64
--> Processing Dependency: libquadmath.so.0()(64bit) for package: python2-scipy-0.18.0-3.el7.x86_64
--> Processing Dependency: libgfortran.so.3()(64bit) for package: python2-scipy-0.18.0-3.el7.x86_64
--> Running transaction check
---> Package atlas.x86_64 0:3.10.1-12.el7 will be installed
---> Package libgfortran.x86_64 0:4.8.5-16.el7 will be installed
---> Package libquadmath.x86_64 0:4.8.5-16.el7 will be installed
---> Package python2-numpy.x86_64 1:1.11.2-2.el7 will be installed
--> Processing Dependency: python-nose for package: 1:python2-numpy-1.11.2-2.el7.x86_64
--> Running transaction check
---> Package python-nose.noarch 0:1.3.7-7.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                                 Arch                  Version                                 Repository                            Size
==================================================================================================================================================
Installing:
 openstack-glance                        noarch                1:15.0.0-1.el7                          centos-openstack-pike                 74 k
Installing for dependencies:
 atlas                                   x86_64                3.10.1-12.el7                           base                                 4.5 M
 libgfortran                             x86_64                4.8.5-16.el7                            base                                 296 k
 libquadmath                             x86_64                4.8.5-16.el7                            base                                 186 k
 libxslt                                 x86_64                1.1.28-5.el7                            base                                 242 k
 pysendfile                              x86_64                2.0.0-5.el7                             centos-openstack-pike                 10 k
 python-boto                             noarch                2.34.0-4.el7                            centos-openstack-pike                1.6 M
 python-glance                           noarch                1:15.0.0-1.el7                          centos-openstack-pike                778 k
 python-httplib2                         noarch                0.9.2-1.el7                             centos-openstack-pike                115 k
 python-lxml                             x86_64                3.2.1-4.el7                             base                                 758 k
 python-networkx                         noarch                1.10-1.el7                              centos-openstack-pike                7.8 k
 python-networkx-core                    noarch                1.10-1.el7                              centos-openstack-pike                1.6 M
 python-nose                             noarch                1.3.7-7.el7                             centos-openstack-pike                276 k
 python-oslo-privsep-lang                noarch                1.22.0-1.el7                            centos-openstack-pike                8.0 k
 python-oslo-vmware-lang                 noarch                2.23.0-1.el7                            centos-openstack-pike                9.2 k
 python-retrying                         noarch                1.2.3-4.el7                             centos-openstack-pike                 16 k
 python-simplegeneric                    noarch                0.8-7.el7                               centos-openstack-pike                 12 k
 python2-automaton                       noarch                1.12.1-1.el7                            centos-openstack-pike                 37 k
 python2-castellan                       noarch                0.12.0-1.el7                            centos-openstack-pike                 93 k
 python2-cursive                         noarch                0.1.2-1.el7                             centos-openstack-pike                 26 k
 python2-glance-store                    noarch                0.22.0-1.el7                            centos-openstack-pike                215 k
 python2-numpy                           x86_64                1:1.11.2-2.el7                          centos-openstack-pike                3.2 M
 python2-os-brick                        noarch                1.15.3-1.el7                            centos-openstack-pike                331 k
 python2-os-win                          noarch                2.2.0-1.el7                             centos-openstack-pike                396 k
 python2-oslo-privsep                    noarch                1.22.0-1.el7                            centos-openstack-pike                 30 k
 python2-oslo-rootwrap                   noarch                5.9.0-1.el7                             centos-openstack-pike                 38 k
 python2-oslo-vmware                     noarch                2.23.0-1.el7                            centos-openstack-pike                188 k
 python2-rsa                             noarch                3.3-2.el7                               centos-openstack-pike                 63 k
 python2-scipy                           x86_64                0.18.0-3.el7                            centos-openstack-pike                 12 M
 python2-suds                            noarch                0.7-0.4.94664ddd46a6.el7                centos-openstack-pike                234 k
 python2-swiftclient                     noarch                3.4.0-1.el7                             centos-openstack-pike                156 k
 python2-taskflow                        noarch                2.14.0-1.el7                            centos-openstack-pike                678 k
 python2-wsme                            noarch                0.9.2-1.el7                             centos-openstack-pike                193 k

Transaction Summary
==================================================================================================================================================
Install  1 Package (+32 Dependent packages)

Total download size: 28 M
Installed size: 121 M
Downloading packages:
(1/33): libgfortran-4.8.5-16.el7.x86_64.rpm                                                                                | 296 kB  00:00:00     
(2/33): libxslt-1.1.28-5.el7.x86_64.rpm                                                                                    | 242 kB  00:00:00     
(3/33): libquadmath-4.8.5-16.el7.x86_64.rpm                                                                                | 186 kB  00:00:00     
(4/33): pysendfile-2.0.0-5.el7.x86_64.rpm                                                                                  |  10 kB  00:00:00     
(5/33): atlas-3.10.1-12.el7.x86_64.rpm                                                                                     | 4.5 MB  00:00:01     
(6/33): openstack-glance-15.0.0-1.el7.noarch.rpm                                                                           |  74 kB  00:00:02     
(7/33): python-glance-15.0.0-1.el7.noarch.rpm                                                                              | 778 kB  00:00:02     
(8/33): python-lxml-3.2.1-4.el7.x86_64.rpm                                                                                 | 758 kB  00:00:00     
(9/33): python-httplib2-0.9.2-1.el7.noarch.rpm                                                                             | 115 kB  00:00:00     
(10/33): python-networkx-1.10-1.el7.noarch.rpm                                                                             | 7.8 kB  00:00:00     
(11/33): python-boto-2.34.0-4.el7.noarch.rpm                                                                               | 1.6 MB  00:00:05     
(12/33): python-nose-1.3.7-7.el7.noarch.rpm                                                                                | 276 kB  00:00:00     
(13/33): python-oslo-privsep-lang-1.22.0-1.el7.noarch.rpm                                                                  | 8.0 kB  00:00:00     
(14/33): python-oslo-vmware-lang-2.23.0-1.el7.noarch.rpm                                                                   | 9.2 kB  00:00:00     
(15/33): python-retrying-1.2.3-4.el7.noarch.rpm                                                                            |  16 kB  00:00:00     
(16/33): python-simplegeneric-0.8-7.el7.noarch.rpm                                                                         |  12 kB  00:00:00     
(17/33): python2-automaton-1.12.1-1.el7.noarch.rpm                                                                         |  37 kB  00:00:00     
(18/33): python-networkx-core-1.10-1.el7.noarch.rpm                                                                        | 1.6 MB  00:00:03     
(19/33): python2-castellan-0.12.0-1.el7.noarch.rpm                                                                         |  93 kB  00:00:00     
(20/33): python2-cursive-0.1.2-1.el7.noarch.rpm                                                                            |  26 kB  00:00:00     
(21/33): python2-glance-store-0.22.0-1.el7.noarch.rpm                                                                      | 215 kB  00:00:00     
(22/33): python2-os-brick-1.15.3-1.el7.noarch.rpm                                                                          | 331 kB  00:00:01     
(23/33): python2-os-win-2.2.0-1.el7.noarch.rpm                                                                             | 396 kB  00:00:01     
(24/33): python2-oslo-privsep-1.22.0-1.el7.noarch.rpm                                                                      |  30 kB  00:00:00     
(25/33): python2-oslo-rootwrap-5.9.0-1.el7.noarch.rpm                                                                      |  38 kB  00:00:00     
(26/33): python2-oslo-vmware-2.23.0-1.el7.noarch.rpm                                                                       | 188 kB  00:00:00     
(27/33): python2-rsa-3.3-2.el7.noarch.rpm                                                                                  |  63 kB  00:00:00     
(28/33): python2-numpy-1.11.2-2.el7.x86_64.rpm                                                                             | 3.2 MB  00:00:10     
(29/33): python2-suds-0.7-0.4.94664ddd46a6.el7.noarch.rpm                                                                  | 234 kB  00:00:01     
(30/33): python2-swiftclient-3.4.0-1.el7.noarch.rpm                                                                        | 156 kB  00:00:00     
(31/33): python2-taskflow-2.14.0-1.el7.noarch.rpm                                                                          | 678 kB  00:00:01     
(32/33): python2-wsme-0.9.2-1.el7.noarch.rpm                                                                               | 193 kB  00:00:01     
(33/33): python2-scipy-0.18.0-3.el7.x86_64.rpm                                                                             |  12 MB  00:00:22     
--------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                             771 kB/s |  28 MB  00:00:37     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : libquadmath-4.8.5-16.el7.x86_64                                                                                               1/33 
  Installing : libgfortran-4.8.5-16.el7.x86_64                                                                                               2/33 
  Installing : atlas-3.10.1-12.el7.x86_64                                                                                                    3/33 
  Installing : python-retrying-1.2.3-4.el7.noarch                                                                                            4/33 
  Installing : python-httplib2-0.9.2-1.el7.noarch                                                                                            5/33 
  Installing : python2-suds-0.7-0.4.94664ddd46a6.el7.noarch                                                                                  6/33 
  Installing : python-oslo-privsep-lang-1.22.0-1.el7.noarch                                                                                  7/33 
  Installing : python2-oslo-privsep-1.22.0-1.el7.noarch                                                                                      8/33 
  Installing : python-oslo-vmware-lang-2.23.0-1.el7.noarch                                                                                   9/33 
  Installing : python2-os-win-2.2.0-1.el7.noarch                                                                                            10/33 
  Installing : python2-os-brick-1.15.3-1.el7.noarch                                                                                         11/33 
  Installing : libxslt-1.1.28-5.el7.x86_64                                                                                                  12/33 
  Installing : python-lxml-3.2.1-4.el7.x86_64                                                                                               13/33 
  Installing : python2-oslo-vmware-2.23.0-1.el7.noarch                                                                                      14/33 
  Installing : pysendfile-2.0.0-5.el7.x86_64                                                                                                15/33 
  Installing : python2-oslo-rootwrap-5.9.0-1.el7.noarch                                                                                     16/33 
  Installing : python2-glance-store-0.22.0-1.el7.noarch                                                                                     17/33 
  Installing : python-nose-1.3.7-7.el7.noarch                                                                                               18/33 
  Installing : 1:python2-numpy-1.11.2-2.el7.x86_64                                                                                          19/33 
  Installing : python2-scipy-0.18.0-3.el7.x86_64                                                                                            20/33 
  Installing : python-networkx-core-1.10-1.el7.noarch                                                                                       21/33 
  Installing : python-networkx-1.10-1.el7.noarch                                                                                            22/33 
  Installing : python2-rsa-3.3-2.el7.noarch                                                                                                 23/33 
  Installing : python-boto-2.34.0-4.el7.noarch                                                                                              24/33 
  Installing : python2-automaton-1.12.1-1.el7.noarch                                                                                        25/33 
  Installing : python2-taskflow-2.14.0-1.el7.noarch                                                                                         26/33 
  Installing : python2-castellan-0.12.0-1.el7.noarch                                                                                        27/33 
  Installing : python2-cursive-0.1.2-1.el7.noarch                                                                                           28/33 
  Installing : python-simplegeneric-0.8-7.el7.noarch                                                                                        29/33 
  Installing : python2-wsme-0.9.2-1.el7.noarch                                                                                              30/33 
  Installing : python2-swiftclient-3.4.0-1.el7.noarch                                                                                       31/33 
  Installing : 1:python-glance-15.0.0-1.el7.noarch                                                                                          32/33 
  Installing : 1:openstack-glance-15.0.0-1.el7.noarch                                                                                       33/33 
  Verifying  : python-networkx-core-1.10-1.el7.noarch                                                                                        1/33 
  Verifying  : python2-swiftclient-3.4.0-1.el7.noarch                                                                                        2/33 
  Verifying  : libgfortran-4.8.5-16.el7.x86_64                                                                                               3/33 
  Verifying  : python-simplegeneric-0.8-7.el7.noarch                                                                                         4/33 
  Verifying  : 1:openstack-glance-15.0.0-1.el7.noarch                                                                                        5/33 
  Verifying  : python2-wsme-0.9.2-1.el7.noarch                                                                                               6/33 
  Verifying  : python-lxml-3.2.1-4.el7.x86_64                                                                                                7/33 
  Verifying  : python2-scipy-0.18.0-3.el7.x86_64                                                                                             8/33 
  Verifying  : atlas-3.10.1-12.el7.x86_64                                                                                                    9/33 
  Verifying  : python2-os-brick-1.15.3-1.el7.noarch                                                                                         10/33 
  Verifying  : python2-castellan-0.12.0-1.el7.noarch                                                                                        11/33 
  Verifying  : python2-automaton-1.12.1-1.el7.noarch                                                                                        12/33 
  Verifying  : python2-rsa-3.3-2.el7.noarch                                                                                                 13/33 
  Verifying  : python2-glance-store-0.22.0-1.el7.noarch                                                                                     14/33 
  Verifying  : python-retrying-1.2.3-4.el7.noarch                                                                                           15/33 
  Verifying  : 1:python2-numpy-1.11.2-2.el7.x86_64                                                                                          16/33 
  Verifying  : python-nose-1.3.7-7.el7.noarch                                                                                               17/33 
  Verifying  : python2-taskflow-2.14.0-1.el7.noarch                                                                                         18/33 
  Verifying  : 1:python-glance-15.0.0-1.el7.noarch                                                                                          19/33 
  Verifying  : python2-oslo-rootwrap-5.9.0-1.el7.noarch                                                                                     20/33 
  Verifying  : pysendfile-2.0.0-5.el7.x86_64                                                                                                21/33 
  Verifying  : libxslt-1.1.28-5.el7.x86_64                                                                                                  22/33 
  Verifying  : python-networkx-1.10-1.el7.noarch                                                                                            23/33 
  Verifying  : python2-os-win-2.2.0-1.el7.noarch                                                                                            24/33 
  Verifying  : python-oslo-vmware-lang-2.23.0-1.el7.noarch                                                                                  25/33 
  Verifying  : libquadmath-4.8.5-16.el7.x86_64                                                                                              26/33 
  Verifying  : python2-cursive-0.1.2-1.el7.noarch                                                                                           27/33 
  Verifying  : python-oslo-privsep-lang-1.22.0-1.el7.noarch                                                                                 28/33 
  Verifying  : python2-suds-0.7-0.4.94664ddd46a6.el7.noarch                                                                                 29/33 
  Verifying  : python2-oslo-vmware-2.23.0-1.el7.noarch                                                                                      30/33 
  Verifying  : python2-oslo-privsep-1.22.0-1.el7.noarch                                                                                     31/33 
  Verifying  : python-httplib2-0.9.2-1.el7.noarch                                                                                           32/33 
  Verifying  : python-boto-2.34.0-4.el7.noarch                                                                                              33/33 

Installed:
  openstack-glance.noarch 1:15.0.0-1.el7                                                                                                          

Dependency Installed:
  atlas.x86_64 0:3.10.1-12.el7                     libgfortran.x86_64 0:4.8.5-16.el7                libquadmath.x86_64 0:4.8.5-16.el7           
  libxslt.x86_64 0:1.1.28-5.el7                    pysendfile.x86_64 0:2.0.0-5.el7                  python-boto.noarch 0:2.34.0-4.el7           
  python-glance.noarch 1:15.0.0-1.el7              python-httplib2.noarch 0:0.9.2-1.el7             python-lxml.x86_64 0:3.2.1-4.el7            
  python-networkx.noarch 0:1.10-1.el7              python-networkx-core.noarch 0:1.10-1.el7         python-nose.noarch 0:1.3.7-7.el7            
  python-oslo-privsep-lang.noarch 0:1.22.0-1.el7   python-oslo-vmware-lang.noarch 0:2.23.0-1.el7    python-retrying.noarch 0:1.2.3-4.el7        
  python-simplegeneric.noarch 0:0.8-7.el7          python2-automaton.noarch 0:1.12.1-1.el7          python2-castellan.noarch 0:0.12.0-1.el7     
  python2-cursive.noarch 0:0.1.2-1.el7             python2-glance-store.noarch 0:0.22.0-1.el7       python2-numpy.x86_64 1:1.11.2-2.el7         
  python2-os-brick.noarch 0:1.15.3-1.el7           python2-os-win.noarch 0:2.2.0-1.el7              python2-oslo-privsep.noarch 0:1.22.0-1.el7  
  python2-oslo-rootwrap.noarch 0:5.9.0-1.el7       python2-oslo-vmware.noarch 0:2.23.0-1.el7        python2-rsa.noarch 0:3.3-2.el7              
  python2-scipy.x86_64 0:0.18.0-3.el7              python2-suds.noarch 0:0.7-0.4.94664ddd46a6.el7   python2-swiftclient.noarch 0:3.4.0-1.el7    
  python2-taskflow.noarch 0:2.14.0-1.el7           python2-wsme.noarch 0:0.9.2-1.el7               

Complete!
```

Configure
```
[vagrant@localhost ~]$ sudo sed -i 's/^\[DEFAULT\]$/&\ndebug=true\nverbose=true\n/;s%^\[database\]$%&\nconnection=mysql+pymysql://glance:SERVICE_DBPASS@10.64.33.64/glance\n%;s%^\[keystone_authtoken\]$%&\nauth_uri=http://10.64.33.64:5000\nauth_url=http://10.64.33.64:35357\nmemcached_servers=10.64.33.64:11211\nauth_type=password\nproject_domain_name=default\nuser_domain_name=default\nproject_name=service\nusername=glance\npassword=SERVICE_PASS\n%;s/^\[paste_deploy\]$/&\nflavor=keystone\n/;s%^\[glance_store\]$%&\nstores=file,http\ndefault_store=file\nfilesystem_store_datadir=/var/lib/glance/images/\n%' /etc/glance/glance-api.conf
```

```
[vagrant@localhost ~]$ sudo sed -i 's/^\[DEFAULT\]$/&\ndebug=true\nverbose=true\n/;s%^\[database\]$%&\nconnection=mysql+pymysql://glance:SERVICE_DBPASS@10.64.33.64/glance\n%;s%^\[keystone_authtoken\]$%&\nauth_uri=http://10.64.33.64:5000\nauth_url=http://10.64.33.64:35357\nmemcached_servers=10.64.33.64:11211\nauth_type=password\nproject_domain_name=default\nuser_domain_name=default\nproject_name=service\nusername=glance\npassword=SERVICE_PASS\n%;s/^\[paste_deploy\]$/&\nflavor=keystone\n/' /etc/glance/glance-registry.conf
```

View "glance-api.conf"
```
[vagrant@controller-10-64-33-64 ~]$ sudo cat /etc/glance/glance-api.conf | egrep '^[^#]'
[DEFAULT]
debug=true
verbose=true
[cors]
[database]
connection=mysql+pymysql://glance:SERVICE_DBPASS@10.64.33.64/glance
[glance_store]
stores=file,http
default_store=file
filesystem_store_datadir=/var/lib/glance/images/
[image_format]
[keystone_authtoken]
auth_uri=http://10.64.33.64:5000
auth_url=http://10.64.33.64:35357
memcached_servers=10.64.33.64:11211
auth_type=password
project_domain_name=default
user_domain_name=default
project_name=service
username=glance
password=SERVICE_PASS
[matchmaker_redis]
[oslo_concurrency]
[oslo_messaging_amqp]
[oslo_messaging_kafka]
[oslo_messaging_notifications]
[oslo_messaging_rabbit]
[oslo_messaging_zmq]
[oslo_middleware]
[oslo_policy]
[paste_deploy]
flavor=keystone
[profiler]
[store_type_location_strategy]
[task]
[taskflow_executor]
```

View "glance-registry.conf"
```
[vagrant@controller-10-64-33-64 ~]$ sudo cat /etc/glance/glance-registry.conf | egrep '^[^#]'
[DEFAULT]
debug=true
verbose=true
[database]
connection=mysql+pymysql://glance:SERVICE_DBPASS@10.64.33.64/glance
[keystone_authtoken]
auth_uri=http://10.64.33.64:5000
auth_url=http://10.64.33.64:35357
memcached_servers=10.64.33.64:11211
auth_type=password
project_domain_name=default
user_domain_name=default
project_name=service
username=glance
password=SERVICE_PASS
[matchmaker_redis]
[oslo_messaging_amqp]
[oslo_messaging_kafka]
[oslo_messaging_notifications]
[oslo_messaging_rabbit]
[oslo_messaging_zmq]
[oslo_policy]
[paste_deploy]
flavor=keystone
[profiler]
```

### Start

Database
```
[vagrant@localhost ~]$ mysql -u root -e "CREATE DATABASE glance;GRANT ALL PRIVILEGES ON glance.* TO 'glance'@'localhost' IDENTIFIED BY 'SERVICE_DBPASS';GRANT ALL PRIVILEGES ON glance.* TO 'glance'@'%' IDENTIFIED BY 'SERVICE_DBPASS';"
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
| performance_schema |
| test               |
+--------------------+
```

```
[vagrant@localhost ~]$ sudo su -s /bin/sh -c "glance-manage db_sync" glance
/usr/lib/python2.7/site-packages/oslo_db/sqlalchemy/enginefacade.py:1328: OsloDBDeprecationWarning: EngineFacade is deprecated; please use oslo_db.sqlalchemy.enginefacade
  expire_on_commit=expire_on_commit, _conf=conf)
INFO  [alembic.runtime.migration] Context impl MySQLImpl.
INFO  [alembic.runtime.migration] Will assume non-transactional DDL.
INFO  [alembic.runtime.migration] Running upgrade  -> liberty, liberty initial
INFO  [alembic.runtime.migration] Running upgrade liberty -> mitaka01, add index on created_at and updated_at columns of 'images' table
INFO  [alembic.runtime.migration] Running upgrade mitaka01 -> mitaka02, update metadef os_nova_server
INFO  [alembic.runtime.migration] Running upgrade mitaka02 -> ocata01, add visibility to and remove is_public from images
INFO  [alembic.runtime.migration] Running upgrade ocata01 -> pike01, drop glare artifacts tables
INFO  [alembic.runtime.migration] Context impl MySQLImpl.
INFO  [alembic.runtime.migration] Will assume non-transactional DDL.
Upgraded database to: pike01, current revision(s): pike01
```

```
[vagrant@localhost ~]$ mysql -u root -e "show tables in glance;"
+----------------------------------+
| Tables_in_glance                 |
+----------------------------------+
| alembic_version                  |
| image_locations                  |
| image_members                    |
| image_properties                 |
| image_tags                       |
| images                           |
| metadef_namespace_resource_types |
| metadef_namespaces               |
| metadef_objects                  |
| metadef_properties               |
| metadef_resource_types           |
| metadef_tags                     |
| migrate_version                  |
| task_info                        |
| tasks                            |
+----------------------------------+
[vagrant@localhost ~]$ mysql -u root -e "select * from glance.alembic_version;"
+-------------+
| version_num |
+-------------+
| pike01      |
+-------------+
[vagrant@localhost ~]$ mysql -u root -e "select * from glance.migrate_version;"
+-------------------+--------------------------------------------------------------------+---------+
| repository_id     | repository_path                                                    | version |
+-------------------+--------------------------------------------------------------------+---------+
| Glance Migrations | /usr/lib/python2.7/site-packages/glance/db/sqlalchemy/migrate_repo |       0 |
+-------------------+--------------------------------------------------------------------+---------+
```


```
[vagrant@localhost ~]$ sudo systemctl start openstack-glance-api.service
[vagrant@localhost ~]$ systemctl -l status openstack-glance-api.service
● openstack-glance-api.service - OpenStack Image Service (code-named Glance) API server
   Loaded: loaded (/usr/lib/systemd/system/openstack-glance-api.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 00:37:49 UTC; 17s ago
 Main PID: 29381 (glance-api)
   CGroup: /system.slice/openstack-glance-api.service
           ├─29381 /usr/bin/python2 /usr/bin/glance-api
           └─29390 /usr/bin/python2 /usr/bin/glance-api
[vagrant@localhost ~]$ sudo systemctl start openstack-glance-registry.service
[vagrant@localhost ~]$ systemctl -l status openstack-glance-registry.service
● openstack-glance-registry.service - OpenStack Image Service (code-named Glance) Registry server
   Loaded: loaded (/usr/lib/systemd/system/openstack-glance-registry.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 00:38:23 UTC; 8s ago
 Main PID: 29407 (glance-registry)
   CGroup: /system.slice/openstack-glance-registry.service
           ├─29407 /usr/bin/python2 /usr/bin/glance-registry
           └─29416 /usr/bin/python2 /usr/bin/glance-registry
```

Set Auto Starting when reboot
```
[vagrant@localhost ~]$ sudo systemctl enable openstack-glance-api.service openstack-glance-registry.service
Created symlink from /etc/systemd/system/multi-user.target.wants/openstack-glance-api.service to /usr/lib/systemd/system/openstack-glance-api.service.
Created symlink from /etc/systemd/system/multi-user.target.wants/openstack-glance-registry.service to /usr/lib/systemd/system/openstack-glance-registry.service.
```

Networking
```
[vagrant@localhost ~]$ sudo netstat -tpnl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       PID/Program name    
tcp        0      0 127.0.0.1:11211         0.0.0.0:*               LISTEN      25037/memcached     
tcp        0      0 0.0.0.0:9292            0.0.0.0:*               LISTEN      29381/python2       
tcp        0      0 0.0.0.0:111             0.0.0.0:*               LISTEN      1/systemd           
tcp        0      0 0.0.0.0:4369            0.0.0.0:*               LISTEN      1/systemd           
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN      877/sshd            
tcp        0      0 127.0.0.1:25            0.0.0.0:*               LISTEN      1102/master         
tcp        0      0 0.0.0.0:9191            0.0.0.0:*               LISTEN      29407/python2       
tcp        0      0 0.0.0.0:25672           0.0.0.0:*               LISTEN      23134/beam          
tcp        0      0 10.64.33.64:3306        0.0.0.0:*               LISTEN      22899/mysqld        
tcp6       0      0 ::1:11211               :::*                    LISTEN      25037/memcached     
tcp6       0      0 :::111                  :::*                    LISTEN      1/systemd           
tcp6       0      0 :::80                   :::*                    LISTEN      27161/httpd         
tcp6       0      0 :::22                   :::*                    LISTEN      877/sshd            
tcp6       0      0 ::1:25                  :::*                    LISTEN      1102/master         
tcp6       0      0 :::35357                :::*                    LISTEN      27161/httpd         
tcp6       0      0 :::5000                 :::*                    LISTEN      27161/httpd         
tcp6       0      0 :::5672                 :::*                    LISTEN      23134/beam          
```

### Logs

Verifying
```
[vagrant@localhost ~]$ sudo tail /var/log/glance/api.log
2017-10-21 00:37:50.698 29381 DEBUG glance_store.backend [-] Attempting to import store file _load_store /usr/lib/python2.7/site-packages/glance_store/backend.py:231
2017-10-21 00:37:50.699 29381 INFO glance_store._drivers.filesystem [-] Directory to write image files does not exist (/var/lib/glance/images/). Creating.
2017-10-21 00:37:50.699 29381 DEBUG glance_store.capabilities [-] Store glance_store._drivers.filesystem.Store doesn't support updating dynamic storage capabilities. Please overwrite 'update_capabilities' method of the store to implement updating logics if needed. update_capabilities /usr/lib/python2.7/site-packages/glance_store/capabilities.py:97
2017-10-21 00:37:50.699 29381 DEBUG glance_store.backend [-] Registering store file with schemes ('file', 'filesystem') create_stores /usr/lib/python2.7/site-packages/glance_store/backend.py:278
2017-10-21 00:37:50.699 29381 DEBUG glance_store.driver [-] Late loading location class glance_store._drivers.filesystem.StoreLocation get_store_location_class /usr/lib/python2.7/site-packages/glance_store/driver.py:89
2017-10-21 00:37:50.699 29381 DEBUG glance_store.location [-] Registering scheme file with {'location_class': <class 'glance_store._drivers.filesystem.StoreLocation'>, 'store': <glance_store._drivers.filesystem.Store object at 0x5849ed0>, 'store_entry': 'file'} register_scheme_map /usr/lib/python2.7/site-packages/glance_store/location.py:88
2017-10-21 00:37:50.699 29381 DEBUG glance_store.location [-] Registering scheme filesystem with {'location_class': <class 'glance_store._drivers.filesystem.StoreLocation'>, 'store': <glance_store._drivers.filesystem.Store object at 0x5849ed0>, 'store_entry': 'file'} register_scheme_map /usr/lib/python2.7/site-packages/glance_store/location.py:88
2017-10-21 00:37:50.700 29381 INFO glance.common.wsgi [-] Starting 1 workers
2017-10-21 00:37:50.702 29381 INFO glance.common.wsgi [-] Started child 29390
2017-10-21 00:37:50.704 29390 INFO eventlet.wsgi.server [-] (29390) wsgi starting up on http://0.0.0.0:9292
```

```
[vagrant@localhost ~]$ sudo tail /var/log/glance/registry.log
2017-10-21 00:38:24.397 29407 DEBUG glance.common.config [-] keystone_authtoken.username    = glance log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 00:38:24.397 29407 DEBUG glance.common.config [-] image_format.container_formats = ['ami', 'ari', 'aki', 'bare', 'ovf', 'ova', 'docker'] log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 00:38:24.397 29407 DEBUG glance.common.config [-] image_format.disk_formats      = ['ami', 'ari', 'aki', 'vhd', 'vhdx', 'vmdk', 'raw', 'qcow2', 'vdi', 'iso', 'ploop'] log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 00:38:24.397 29407 DEBUG glance.common.config [-] oslo_policy.policy_default_rule = default log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 00:38:24.397 29407 DEBUG glance.common.config [-] oslo_policy.policy_dirs        = ['policy.d'] log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 00:38:24.398 29407 DEBUG glance.common.config [-] oslo_policy.policy_file        = policy.json log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 00:38:24.398 29407 DEBUG glance.common.config [-] ******************************************************************************** log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2889
2017-10-21 00:38:24.398 29407 INFO glance.common.wsgi [-] Starting 1 workers
2017-10-21 00:38:24.399 29407 INFO glance.common.wsgi [-] Started child 29416
2017-10-21 00:38:24.401 29416 INFO eventlet.wsgi.server [-] (29416) wsgi starting up on http://0.0.0.0:9191
```

### Create

Cirros
```
[vagrant@localhost ~]$ curl -O http://download.cirros-cloud.net/0.3.5/cirros-0.3.5-x86_64-disk.img
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 12.6M  100 12.6M    0     0   211k      0  0:01:01  0:01:01 --:--:--  170k
```

Upload
```
[vagrant@localhost ~]$ openstack image create "cirros" --file cirros-0.3.5-x86_64-disk.img --disk-format qcow2 --container-format bare --public
+------------------+------------------------------------------------------+
| Field            | Value                                                |
+------------------+------------------------------------------------------+
| checksum         | f8ab98ff5e73ebab884d80c9dc9c7290                     |
| container_format | bare                                                 |
| created_at       | 2017-10-21T00:46:19Z                                 |
| disk_format      | qcow2                                                |
| file             | /v2/images/40d37d6a-2dba-4323-b648-e806f3acb857/file |
| id               | 40d37d6a-2dba-4323-b648-e806f3acb857                 |
| min_disk         | 0                                                    |
| min_ram          | 0                                                    |
| name             | cirros                                               |
| owner            | a0be38aef8c74d4abca3e4e100ee7910                     |
| protected        | False                                                |
| schema           | /v2/schemas/image                                    |
| size             | 13267968                                             |
| status           | active                                               |
| tags             |                                                      |
| updated_at       | 2017-10-21T00:46:19Z                                 |
| virtual_size     | None                                                 |
| visibility       | public                                               |
+------------------+------------------------------------------------------+
```

```
[vagrant@localhost ~]$ sudo ls -lh /var/lib/glance/images
total 13M
-rw-r-----. 1 glance glance 13M Oct 21 00:46 40d37d6a-2dba-4323-b648-e806f3acb857
```

```
[vagrant@localhost ~]$ openstack image list
+--------------------------------------+--------+--------+
| ID                                   | Name   | Status |
+--------------------------------------+--------+--------+
| 40d37d6a-2dba-4323-b648-e806f3acb857 | cirros | active |
+--------------------------------------+--------+--------+
```





