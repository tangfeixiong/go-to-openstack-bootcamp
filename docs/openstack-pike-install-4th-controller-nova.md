# OpenStack Pike Installation

## Table of content

控制节点
* [之前章节](./openstack-pike-install-3rd-image-glance.md)
* [Nova服务](#compute)

## Controller

### Compute

[Nova](https://docs.openstack.org/nova/pike/install/install-rdo.html)


Database
```
[vagrant@localhost ~]$ mysql -u root -e "CREATE DATABASE nova_api;GRANT ALL PRIVILEGES ON nova_api.* TO 'nova'@'localhost' IDENTIFIED BY 'SERVICE_DBPASS';GRANT ALL PRIVILEGES ON nova_api.* TO 'nova'@'%' IDENTIFIED BY 'SERVICE_DBPASS';"
[vagrant@localhost ~]$ mysql -u root -e "CREATE DATABASE nova;GRANT ALL PRIVILEGES ON nova.* TO 'nova'@'localhost' IDENTIFIED BY 'SERVICE_DBPASS';GRANT ALL PRIVILEGES ON nova.* TO 'nova'@'%' IDENTIFIED BY 'SERVICE_DBPASS';"
[vagrant@localhost ~]$ mysql -u root -e "CREATE DATABASE nova_cell0;GRANT ALL PRIVILEGES ON nova_cell0.* TO 'nova'@'localhost' IDENTIFIED BY 'SERVICE_DBPASS';GRANT ALL PRIVILEGES ON nova_cell0.* TO 'nova'@'%' IDENTIFIED BY 'SERVICE_DBPASS';"
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
| nova               |
| nova_api           |
| nova_cell0         |
| performance_schema |
| test               |
+--------------------+
```


Client
```
[vagrant@localhost ~]$ . openstack-admin.sh 
```

Registration
```
[vagrant@localhost ~]$ openstack user create --domain default --password SERVICE_PASS nova 
+---------------------+----------------------------------+
| Field               | Value                            |
+---------------------+----------------------------------+
| domain_id           | default                          |
| enabled             | True                             |
| id                  | 302b5318ebf048a8ae97d0ba8b2d7b14 |
| name                | nova                             |
| options             | {}                               |
| password_expires_at | None                             |
+---------------------+----------------------------------+
[vagrant@localhost ~]$ openstack role add --project service --user nova admin
[vagrant@localhost ~]$ openstack service create --name nova --description "OpenStack Compute" compute
+-------------+----------------------------------+
| Field       | Value                            |
+-------------+----------------------------------+
| description | OpenStack Compute                |
| enabled     | True                             |
| id          | 5f88f8d5c9494308b4e57b695422015d |
| name        | nova                             |
| type        | compute                          |
+-------------+----------------------------------+
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne compute public http://10.64.33.64:8774/v2.1
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | fa1658ccec0447f8b907d32142bf3588 |
| interface    | public                           |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 5f88f8d5c9494308b4e57b695422015d |
| service_name | nova                             |
| service_type | compute                          |
| url          | http://10.64.33.64:8774/v2.1     |
+--------------+----------------------------------+
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne compute internal http://10.64.33.64:8774/v2.1
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | b5e44769ad7f4ebdafaa4b4911ed667f |
| interface    | internal                         |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 5f88f8d5c9494308b4e57b695422015d |
| service_name | nova                             |
| service_type | compute                          |
| url          | http://10.64.33.64:8774/v2.1     |
+--------------+----------------------------------+
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne compute admin http://10.64.33.64:8774/v2.1
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | d4d9a2860cf446249a648b20f22b643e |
| interface    | admin                            |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 5f88f8d5c9494308b4e57b695422015d |
| service_name | nova                             |
| service_type | compute                          |
| url          | http://10.64.33.64:8774/v2.1     |
+--------------+----------------------------------+
```

```
[vagrant@localhost ~]$ openstack user create --domain default --password SERVICE_PASS placement
+---------------------+----------------------------------+
| Field               | Value                            |
+---------------------+----------------------------------+
| domain_id           | default                          |
| enabled             | True                             |
| id                  | c75b8cc66dd149dcbc18e3146181bed6 |
| name                | placement                        |
| options             | {}                               |
| password_expires_at | None                             |
+---------------------+----------------------------------+
[vagrant@localhost ~]$ openstack role add --project service --user placement admin
[vagrant@localhost ~]$ openstack service create --name placement --description "Placement API" placement
+-------------+----------------------------------+
| Field       | Value                            |
+-------------+----------------------------------+
| description | Placement API                    |
| enabled     | True                             |
| id          | 669cf4a419904eb9a2e4763a5fcc7744 |
| name        | placement                        |
| type        | placement                        |
+-------------+----------------------------------+
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne placement public http://10.64.33.64:8778
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | f4ca03f933734b74ba78966f6040c220 |
| interface    | public                           |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 669cf4a419904eb9a2e4763a5fcc7744 |
| service_name | placement                        |
| service_type | placement                        |
| url          | http://10.64.33.64:8778          |
+--------------+----------------------------------+
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne placement internal http://10.64.33.64:8778
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | 234787bf252a4521bdd8188e69d2971c |
| interface    | internal                         |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 669cf4a419904eb9a2e4763a5fcc7744 |
| service_name | placement                        |
| service_type | placement                        |
| url          | http://10.64.33.64:8778          |
+--------------+----------------------------------+
[vagrant@localhost ~]$ openstack endpoint create --region RegionOne placement admin http://10.64.33.64:8778
+--------------+----------------------------------+
| Field        | Value                            |
+--------------+----------------------------------+
| enabled      | True                             |
| id           | 0e2086ec94cf4ec3b7f64d90e8844f25 |
| interface    | admin                            |
| region       | RegionOne                        |
| region_id    | RegionOne                        |
| service_id   | 669cf4a419904eb9a2e4763a5fcc7744 |
| service_name | placement                        |
| service_type | placement                        |
| url          | http://10.64.33.64:8778          |
+--------------+----------------------------------+
```

Install
```
[vagrant@localhost ~]$ sudo yum install -y openstack-nova-api openstack-nova-conductor openstack-nova-console openstack-nova-novncproxy openstack-nova-scheduler openstack-nova-placement-api
Loaded plugins: fastestmirror
base                                                                                                                       | 3.6 kB  00:00:00     
centos-ceph-jewel                                                                                                          | 2.9 kB  00:00:00     
centos-openstack-pike                                                                                                      | 2.9 kB  00:00:00     
centos-qemu-ev                                                                                                             | 2.9 kB  00:00:00     
extras                                                                                                                     | 3.4 kB  00:00:00     
updates                                                                                                                    | 3.4 kB  00:00:00     
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
Resolving Dependencies
--> Running transaction check
---> Package openstack-nova-api.noarch 1:16.0.1-1.el7 will be installed
--> Processing Dependency: openstack-nova-common = 1:16.0.1-1.el7 for package: 1:openstack-nova-api-16.0.1-1.el7.noarch
---> Package openstack-nova-conductor.noarch 1:16.0.1-1.el7 will be installed
---> Package openstack-nova-console.noarch 1:16.0.1-1.el7 will be installed
--> Processing Dependency: python-websockify >= 0.8.0 for package: 1:openstack-nova-console-16.0.1-1.el7.noarch
---> Package openstack-nova-novncproxy.noarch 1:16.0.1-1.el7 will be installed
--> Processing Dependency: novnc for package: 1:openstack-nova-novncproxy-16.0.1-1.el7.noarch
---> Package openstack-nova-placement-api.noarch 1:16.0.1-1.el7 will be installed
---> Package openstack-nova-scheduler.noarch 1:16.0.1-1.el7 will be installed
--> Running transaction check
---> Package novnc.noarch 0:0.5.1-2.el7 will be installed
---> Package openstack-nova-common.noarch 1:16.0.1-1.el7 will be installed
--> Processing Dependency: python-nova = 1:16.0.1-1.el7 for package: 1:openstack-nova-common-16.0.1-1.el7.noarch
---> Package python-websockify.noarch 0:0.8.0-1.el7 will be installed
--> Running transaction check
---> Package python-nova.noarch 1:16.0.1-1.el7 will be installed
--> Processing Dependency: python-paramiko >= 2.0 for package: 1:python-nova-16.0.1-1.el7.noarch
--> Processing Dependency: python-oslo-versionedobjects >= 1.17.0 for package: 1:python-nova-16.0.1-1.el7.noarch
--> Processing Dependency: python-oslo-reports >= 0.6.0 for package: 1:python-nova-16.0.1-1.el7.noarch
--> Processing Dependency: python-os-vif >= 1.7.0 for package: 1:python-nova-16.0.1-1.el7.noarch
--> Processing Dependency: python-microversion-parse >= 0.1.2 for package: 1:python-nova-16.0.1-1.el7.noarch
--> Processing Dependency: python-tooz for package: 1:python-nova-16.0.1-1.el7.noarch
--> Processing Dependency: python-psutil for package: 1:python-nova-16.0.1-1.el7.noarch
--> Processing Dependency: python-os-traits for package: 1:python-nova-16.0.1-1.el7.noarch
--> Running transaction check
---> Package python-paramiko.noarch 0:2.1.1-2.el7 will be installed
---> Package python-tooz.noarch 0:1.58.0-1.el7 will be installed
--> Processing Dependency: python-voluptuous >= 0.8.9 for package: python-tooz-1.58.0-1.el7.noarch
--> Processing Dependency: python-zake for package: python-tooz-1.58.0-1.el7.noarch
--> Processing Dependency: python-redis for package: python-tooz-1.58.0-1.el7.noarch
---> Package python2-microversion-parse.noarch 0:0.1.4-1.el7 will be installed
---> Package python2-os-traits.noarch 0:0.3.3-1.el7 will be installed
---> Package python2-os-vif.noarch 0:1.7.0-1.el7 will be installed
---> Package python2-oslo-reports.noarch 0:1.22.0-1.el7 will be installed
---> Package python2-oslo-versionedobjects.noarch 0:1.26.0-1.el7 will be installed
--> Processing Dependency: python-oslo-versionedobjects-lang = 1.26.0-1.el7 for package: python2-oslo-versionedobjects-1.26.0-1.el7.noarch
--> Processing Dependency: python-mock for package: python2-oslo-versionedobjects-1.26.0-1.el7.noarch
---> Package python2-psutil.x86_64 0:5.2.2-2.el7 will be installed
--> Running transaction check
---> Package python-oslo-versionedobjects-lang.noarch 0:1.26.0-1.el7 will be installed
---> Package python-redis.noarch 0:2.10.3-1.el7 will be installed
---> Package python-voluptuous.noarch 0:0.8.9-1.el7 will be installed
---> Package python2-mock.noarch 0:2.0.0-1.el7 will be installed
---> Package python2-zake.noarch 0:0.2.2-2.el7 will be installed
--> Processing Dependency: python-kazoo for package: python2-zake-0.2.2-2.el7.noarch
--> Running transaction check
---> Package python-kazoo.noarch 0:2.2.1-1.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                                          Arch                  Version                        Repository                            Size
==================================================================================================================================================
Installing:
 openstack-nova-api                               noarch                1:16.0.1-1.el7                 centos-openstack-pike                7.4 k
 openstack-nova-conductor                         noarch                1:16.0.1-1.el7                 centos-openstack-pike                5.0 k
 openstack-nova-console                           noarch                1:16.0.1-1.el7                 centos-openstack-pike                6.0 k
 openstack-nova-novncproxy                        noarch                1:16.0.1-1.el7                 centos-openstack-pike                5.4 k
 openstack-nova-placement-api                     noarch                1:16.0.1-1.el7                 centos-openstack-pike                5.2 k
 openstack-nova-scheduler                         noarch                1:16.0.1-1.el7                 centos-openstack-pike                5.0 k
Installing for dependencies:
 novnc                                            noarch                0.5.1-2.el7                    centos-openstack-pike                176 k
 openstack-nova-common                            noarch                1:16.0.1-1.el7                 centos-openstack-pike                367 k
 python-kazoo                                     noarch                2.2.1-1.el7                    centos-openstack-pike                130 k
 python-nova                                      noarch                1:16.0.1-1.el7                 centos-openstack-pike                3.3 M
 python-oslo-versionedobjects-lang                noarch                1.26.0-1.el7                   centos-openstack-pike                7.8 k
 python-paramiko                                  noarch                2.1.1-2.el7                    extras                               267 k
 python-redis                                     noarch                2.10.3-1.el7                   centos-openstack-pike                 94 k
 python-tooz                                      noarch                1.58.0-1.el7                   centos-openstack-pike                 94 k
 python-voluptuous                                noarch                0.8.9-1.el7                    centos-openstack-pike                 36 k
 python-websockify                                noarch                0.8.0-1.el7                    centos-openstack-pike                 69 k
 python2-microversion-parse                       noarch                0.1.4-1.el7                    centos-openstack-pike                 16 k
 python2-mock                                     noarch                2.0.0-1.el7                    centos-openstack-pike                120 k
 python2-os-traits                                noarch                0.3.3-1.el7                    centos-openstack-pike                 22 k
 python2-os-vif                                   noarch                1.7.0-1.el7                    centos-openstack-pike                 59 k
 python2-oslo-reports                             noarch                1.22.0-1.el7                   centos-openstack-pike                 52 k
 python2-oslo-versionedobjects                    noarch                1.26.0-1.el7                   centos-openstack-pike                 72 k
 python2-psutil                                   x86_64                5.2.2-2.el7                    centos-openstack-pike                310 k
 python2-zake                                     noarch                0.2.2-2.el7                    centos-openstack-pike                 39 k

Transaction Summary
==================================================================================================================================================
Install  6 Packages (+18 Dependent packages)

Total download size: 5.2 M
Installed size: 22 M
Downloading packages:
(1/24): openstack-nova-api-16.0.1-1.el7.noarch.rpm                                                                         | 7.4 kB  00:00:00     
(2/24): novnc-0.5.1-2.el7.noarch.rpm                                                                                       | 176 kB  00:00:02     
(3/24): openstack-nova-conductor-16.0.1-1.el7.noarch.rpm                                                                   | 5.0 kB  00:00:00     
(4/24): openstack-nova-console-16.0.1-1.el7.noarch.rpm                                                                     | 6.0 kB  00:00:00     
(5/24): openstack-nova-common-16.0.1-1.el7.noarch.rpm                                                                      | 367 kB  00:00:02     
(6/24): openstack-nova-novncproxy-16.0.1-1.el7.noarch.rpm                                                                  | 5.4 kB  00:00:00     
(7/24): openstack-nova-placement-api-16.0.1-1.el7.noarch.rpm                                                               | 5.2 kB  00:00:00     
(8/24): openstack-nova-scheduler-16.0.1-1.el7.noarch.rpm                                                                   | 5.0 kB  00:00:00     
(9/24): python-kazoo-2.2.1-1.el7.noarch.rpm                                                                                | 130 kB  00:00:00     
(10/24): python-paramiko-2.1.1-2.el7.noarch.rpm                                                                            | 267 kB  00:00:00     
(11/24): python-oslo-versionedobjects-lang-1.26.0-1.el7.noarch.rpm                                                         | 7.8 kB  00:00:00     
(12/24): python-redis-2.10.3-1.el7.noarch.rpm                                                                              |  94 kB  00:00:00     
(13/24): python-tooz-1.58.0-1.el7.noarch.rpm                                                                               |  94 kB  00:00:00     
(14/24): python-voluptuous-0.8.9-1.el7.noarch.rpm                                                                          |  36 kB  00:00:00     
(15/24): python-websockify-0.8.0-1.el7.noarch.rpm                                                                          |  69 kB  00:00:00     
(16/24): python2-microversion-parse-0.1.4-1.el7.noarch.rpm                                                                 |  16 kB  00:00:00     
(17/24): python2-mock-2.0.0-1.el7.noarch.rpm                                                                               | 120 kB  00:00:00     
(18/24): python2-os-traits-0.3.3-1.el7.noarch.rpm                                                                          |  22 kB  00:00:00     
(19/24): python2-os-vif-1.7.0-1.el7.noarch.rpm                                                                             |  59 kB  00:00:00     
(20/24): python2-oslo-reports-1.22.0-1.el7.noarch.rpm                                                                      |  52 kB  00:00:00     
(21/24): python2-oslo-versionedobjects-1.26.0-1.el7.noarch.rpm                                                             |  72 kB  00:00:00     
(22/24): python2-psutil-5.2.2-2.el7.x86_64.rpm                                                                             | 310 kB  00:00:00     
(23/24): python2-zake-0.2.2-2.el7.noarch.rpm                                                                               |  39 kB  00:00:00     
(24/24): python-nova-16.0.1-1.el7.noarch.rpm                                                                               | 3.3 MB  00:00:05     
--------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                             562 kB/s | 5.2 MB  00:00:09     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : python-websockify-0.8.0-1.el7.noarch                                                                                          1/24 
  Installing : python2-psutil-5.2.2-2.el7.x86_64                                                                                             2/24 
  Installing : python2-oslo-reports-1.22.0-1.el7.noarch                                                                                      3/24 
  Installing : novnc-0.5.1-2.el7.noarch                                                                                                      4/24 
  Installing : python2-os-traits-0.3.3-1.el7.noarch                                                                                          5/24 
  Installing : python-voluptuous-0.8.9-1.el7.noarch                                                                                          6/24 
  Installing : python-paramiko-2.1.1-2.el7.noarch                                                                                            7/24 
  Installing : python2-mock-2.0.0-1.el7.noarch                                                                                               8/24 
  Installing : python-kazoo-2.2.1-1.el7.noarch                                                                                               9/24 
  Installing : python2-zake-0.2.2-2.el7.noarch                                                                                              10/24 
  Installing : python-redis-2.10.3-1.el7.noarch                                                                                             11/24 
  Installing : python-tooz-1.58.0-1.el7.noarch                                                                                              12/24 
  Installing : python2-microversion-parse-0.1.4-1.el7.noarch                                                                                13/24 
  Installing : python-oslo-versionedobjects-lang-1.26.0-1.el7.noarch                                                                        14/24 
  Installing : python2-oslo-versionedobjects-1.26.0-1.el7.noarch                                                                            15/24 
  Installing : python2-os-vif-1.7.0-1.el7.noarch                                                                                            16/24 
  Installing : 1:python-nova-16.0.1-1.el7.noarch                                                                                            17/24 
  Installing : 1:openstack-nova-common-16.0.1-1.el7.noarch                                                                                  18/24 
  Installing : 1:openstack-nova-scheduler-16.0.1-1.el7.noarch                                                                               19/24 
  Installing : 1:openstack-nova-novncproxy-16.0.1-1.el7.noarch                                                                              20/24 
  Installing : 1:openstack-nova-console-16.0.1-1.el7.noarch                                                                                 21/24 
  Installing : 1:openstack-nova-api-16.0.1-1.el7.noarch                                                                                     22/24 
  Installing : 1:openstack-nova-placement-api-16.0.1-1.el7.noarch                                                                           23/24 
  Installing : 1:openstack-nova-conductor-16.0.1-1.el7.noarch                                                                               24/24 
  Verifying  : python2-oslo-reports-1.22.0-1.el7.noarch                                                                                      1/24 
  Verifying  : 1:openstack-nova-scheduler-16.0.1-1.el7.noarch                                                                                2/24 
  Verifying  : python-oslo-versionedobjects-lang-1.26.0-1.el7.noarch                                                                         3/24 
  Verifying  : python2-zake-0.2.2-2.el7.noarch                                                                                               4/24 
  Verifying  : 1:openstack-nova-novncproxy-16.0.1-1.el7.noarch                                                                               5/24 
  Verifying  : python2-microversion-parse-0.1.4-1.el7.noarch                                                                                 6/24 
  Verifying  : 1:openstack-nova-console-16.0.1-1.el7.noarch                                                                                  7/24 
  Verifying  : 1:python-nova-16.0.1-1.el7.noarch                                                                                             8/24 
  Verifying  : python-redis-2.10.3-1.el7.noarch                                                                                              9/24 
  Verifying  : python-kazoo-2.2.1-1.el7.noarch                                                                                              10/24 
  Verifying  : python2-mock-2.0.0-1.el7.noarch                                                                                              11/24 
  Verifying  : python-paramiko-2.1.1-2.el7.noarch                                                                                           12/24 
  Verifying  : 1:openstack-nova-api-16.0.1-1.el7.noarch                                                                                     13/24 
  Verifying  : python-voluptuous-0.8.9-1.el7.noarch                                                                                         14/24 
  Verifying  : novnc-0.5.1-2.el7.noarch                                                                                                     15/24 
  Verifying  : python2-oslo-versionedobjects-1.26.0-1.el7.noarch                                                                            16/24 
  Verifying  : 1:openstack-nova-common-16.0.1-1.el7.noarch                                                                                  17/24 
  Verifying  : 1:openstack-nova-placement-api-16.0.1-1.el7.noarch                                                                           18/24 
  Verifying  : python-tooz-1.58.0-1.el7.noarch                                                                                              19/24 
  Verifying  : python2-psutil-5.2.2-2.el7.x86_64                                                                                            20/24 
  Verifying  : python2-os-vif-1.7.0-1.el7.noarch                                                                                            21/24 
  Verifying  : 1:openstack-nova-conductor-16.0.1-1.el7.noarch                                                                               22/24 
  Verifying  : python2-os-traits-0.3.3-1.el7.noarch                                                                                         23/24 
  Verifying  : python-websockify-0.8.0-1.el7.noarch                                                                                         24/24 

Installed:
  openstack-nova-api.noarch 1:16.0.1-1.el7                                  openstack-nova-conductor.noarch 1:16.0.1-1.el7                        
  openstack-nova-console.noarch 1:16.0.1-1.el7                              openstack-nova-novncproxy.noarch 1:16.0.1-1.el7                       
  openstack-nova-placement-api.noarch 1:16.0.1-1.el7                        openstack-nova-scheduler.noarch 1:16.0.1-1.el7                        

Dependency Installed:
  novnc.noarch 0:0.5.1-2.el7                                                openstack-nova-common.noarch 1:16.0.1-1.el7                          
  python-kazoo.noarch 0:2.2.1-1.el7                                         python-nova.noarch 1:16.0.1-1.el7                                    
  python-oslo-versionedobjects-lang.noarch 0:1.26.0-1.el7                   python-paramiko.noarch 0:2.1.1-2.el7                                 
  python-redis.noarch 0:2.10.3-1.el7                                        python-tooz.noarch 0:1.58.0-1.el7                                    
  python-voluptuous.noarch 0:0.8.9-1.el7                                    python-websockify.noarch 0:0.8.0-1.el7                               
  python2-microversion-parse.noarch 0:0.1.4-1.el7                           python2-mock.noarch 0:2.0.0-1.el7                                    
  python2-os-traits.noarch 0:0.3.3-1.el7                                    python2-os-vif.noarch 0:1.7.0-1.el7                                  
  python2-oslo-reports.noarch 0:1.22.0-1.el7                                python2-oslo-versionedobjects.noarch 0:1.26.0-1.el7                  
  python2-psutil.x86_64 0:5.2.2-2.el7                                       python2-zake.noarch 0:0.2.2-2.el7                                    

Complete!
[vagrant@localhost ~]$ 
```

Configure
```
[vagrant@localhost ~]$ sudo sed -i  's%^\[DEFAULT\]$%&\ndebug=true\nverbose=true\ntransport_url=rabbit://openstack:RABBIT_PASS@10.64.33.64\nmy_ip=10.64.33.64\nuser_neutron=True\nfirewall_driver=nova.virt.firewall_NoopFirewallDriver\n%;s%^\[api_database\]$%&\nconnection=mysql+pymysql://nova:SERVICE_DBPASS@10.64.33.64/nova_api\n%;s%^\[database\]$%&\nconnection=mysql+pymysql://nova:SERVICE_DBPASS@10.64.33.64/nova\n%;s/^\[api\]$/&\nauth_strategy=keystone\n/;s%^\[keystone_authtoken\]$%&\nauth_uri=http://10.64.33.64:5000\nauth_url=http://10.64.33.64:35357\nmemcached_servers=10.64.33.64:11211\nauth_type=password\nproject_domain_name=default\nuser_domain_name=default\nproject_name=service\nusername=nova\npassword=SERVICE_PASS\n%;s/^\[vnc\]$/&\nenabled=true\nvncserver_listen=$my_ip\nvncserver_proxyclient_address=$my_ip\n/;s%^\[glance\]$%&\napi_servers=http://10.64.33.64:9292\n%;s%^\[oslo_concurrency\]$%&\nlock_path=/var/lib/nova/tmp\n%;s%^\[placement\]$%&\nos_region_name=RegionOne\nproject_domain_name=Default\nproject_name=service\nauth_type=password\nuser_domain_name=Default\nauth_url=http://10.64.33.64:35357/v3\nusername=placement\npassword=SERVICE_PASS\n%' /etc/nova/nova.conf
```

Patching
```
[vagrant@localhost ~]$ [ ! -f 00-nova-placement-api.conf.orig ] && sudo cp /etc/httpd/conf.d/00-nova-placement-api.conf 00-nova-placement-api.conf.orig
```

```
[vagrant@localhost ~]$ sudo cat 00-nova-placement-api.conf.orig > 00-nova-placement-api.conf && echo -e "\n\n<Directory /usr/bin>\n  <IfVersion >= 2.4>\n    Require all granted\n  </IfVersion>\n  <IfVersion < 2.4 >\n    Order allow,deny\n    Allow from all\n  </IfVersion>\n</Directory>" | tee -a 00-nova-placement-api.conf
Listen 8778

<VirtualHost *:8778>
  WSGIProcessGroup nova-placement-api
  WSGIApplicationGroup %{GLOBAL}
  WSGIPassAuthorization On
  WSGIDaemonProcess nova-placement-api processes=3 threads=1 user=nova group=nova
  WSGIScriptAlias / /usr/bin/nova-placement-api
  <IfVersion >= 2.4>
    ErrorLogFormat "%M"
  </IfVersion>
  ErrorLog /var/log/nova/nova-placement-api.log
  #SSLEngine On
  #SSLCertificateFile ...
  #SSLCertificateKeyFile ...
</VirtualHost>

Alias /nova-placement-api /usr/bin/nova-placement-api
<Location /nova-placement-api>
  SetHandler wsgi-script
  Options +ExecCGI
  WSGIProcessGroup nova-placement-api
  WSGIApplicationGroup %{GLOBAL}
  WSGIPassAuthorization On
</Location>

<Directory /usr/bin>
  <IfVersion >= 2.4>
    Require all granted
  </IfVersion>
  <IfVersion < 2.4 >
    Order allow,deny
    Allow from all
  </IfVersion>
</Directory>
```

```
[vagrant@localhost ~]$ sudo cp 00-nova-placement-api.conf /etc/httpd/conf.d/
```

```
[vagrant@localhost ~]$ sudo systemctl restart httpd.service
```

```
[vagrant@localhost ~]$ sudo systemctl -l status httpd.service
● httpd.service - The Apache HTTP Server
   Loaded: loaded (/usr/lib/systemd/system/httpd.service; enabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 02:35:37 UTC; 7s ago
     Docs: man:httpd(8)
           man:apachectl(8)
  Process: 31331 ExecStop=/bin/kill -WINCH ${MAINPID} (code=exited, status=0/SUCCESS)
 Main PID: 31345 (httpd)
   Status: "Processing requests..."
   CGroup: /system.slice/httpd.service
           ├─31345 /usr/sbin/httpd -DFOREGROUND
           ├─31346 (wsgi:keystone- -DFOREGROUND
           ├─31347 (wsgi:keystone- -DFOREGROUND
           ├─31348 (wsgi:keystone- -DFOREGROUND
           ├─31349 (wsgi:keystone- -DFOREGROUND
           ├─31350 (wsgi:keystone- -DFOREGROUND
           ├─31351 (wsgi:keystone- -DFOREGROUND
           ├─31352 (wsgi:keystone- -DFOREGROUND
           ├─31353 (wsgi:keystone- -DFOREGROUND
           ├─31354 (wsgi:keystone- -DFOREGROUND
           ├─31355 (wsgi:keystone- -DFOREGROUND
           ├─31356 /usr/sbin/httpd -DFOREGROUND
           ├─31357 /usr/sbin/httpd -DFOREGROUND
           ├─31358 /usr/sbin/httpd -DFOREGROUND
           ├─31359 /usr/sbin/httpd -DFOREGROUND
           └─31360 /usr/sbin/httpd -DFOREGROUND

Oct 21 02:35:37 localhost.localdomain systemd[1]: Starting The Apache HTTP Server...
Oct 21 02:35:37 localhost.localdomain systemd[1]: Started The Apache HTTP Server.
```

```
[vagrant@localhost ~]$ sudo su -s /bin/sh -c "nova-manage api_db sync" nova
[vagrant@localhost ~]$ mysql -u root -e "show tables in nova_api;"
+------------------------------+
| Tables_in_nova_api           |
+------------------------------+
| aggregate_hosts              |
| aggregate_metadata           |
| aggregates                   |
| allocations                  |
| build_requests               |
| cell_mappings                |
| consumers                    |
| flavor_extra_specs           |
| flavor_projects              |
| flavors                      |
| host_mappings                |
| instance_group_member        |
| instance_group_policy        |
| instance_groups              |
| instance_mappings            |
| inventories                  |
| key_pairs                    |
| migrate_version              |
| placement_aggregates         |
| project_user_quotas          |
| projects                     |
| quota_classes                |
| quota_usages                 |
| quotas                       |
| request_specs                |
| reservations                 |
| resource_classes             |
| resource_provider_aggregates |
| resource_provider_traits     |
| resource_providers           |
| traits                       |
| users                        |
+------------------------------+
```

```
[vagrant@localhost ~]$ sudo su -s /bin/sh -c "nova-manage cell_v2 map_cell0" nova
[vagrant@localhost ~]$ sudo su -s /bin/sh -c "nova-manage cell_v2 create_cell --name=cell1 --verbose" nova
e85d2a01-4637-4b5c-abad-3345435fdcbd
```

```
[vagrant@localhost ~]$ mysql -u root -e "show tables in nova_cell0;"
+--------------------------------------------+
| Tables_in_nova_cell0                       |
+--------------------------------------------+
| agent_builds                               |
| aggregate_hosts                            |
| aggregate_metadata                         |
| aggregates                                 |
| allocations                                |
| block_device_mapping                       |
| bw_usage_cache                             |
| cells                                      |
| certificates                               |
| compute_nodes                              |
| console_auth_tokens                        |
| console_pools                              |
| consoles                                   |
| dns_domains                                |
| fixed_ips                                  |
| floating_ips                               |
| instance_actions                           |
| instance_actions_events                    |
| instance_extra                             |
| instance_faults                            |
| instance_group_member                      |
| instance_group_policy                      |
| instance_groups                            |
| instance_id_mappings                       |
| instance_info_caches                       |
| instance_metadata                          |
| instance_system_metadata                   |
| instance_type_extra_specs                  |
| instance_type_projects                     |
| instance_types                             |
| instances                                  |
| inventories                                |
| key_pairs                                  |
| migrate_version                            |
| migrations                                 |
| networks                                   |
| pci_devices                                |
| project_user_quotas                        |
| provider_fw_rules                          |
| quota_classes                              |
| quota_usages                               |
| quotas                                     |
| reservations                               |
| resource_provider_aggregates               |
| resource_providers                         |
| s3_images                                  |
| security_group_default_rules               |
| security_group_instance_association        |
| security_group_rules                       |
| security_groups                            |
| services                                   |
| shadow_agent_builds                        |
| shadow_aggregate_hosts                     |
| shadow_aggregate_metadata                  |
| shadow_aggregates                          |
| shadow_block_device_mapping                |
| shadow_bw_usage_cache                      |
| shadow_cells                               |
| shadow_certificates                        |
| shadow_compute_nodes                       |
| shadow_console_pools                       |
| shadow_consoles                            |
| shadow_dns_domains                         |
| shadow_fixed_ips                           |
| shadow_floating_ips                        |
| shadow_instance_actions                    |
| shadow_instance_actions_events             |
| shadow_instance_extra                      |
| shadow_instance_faults                     |
| shadow_instance_group_member               |
| shadow_instance_group_policy               |
| shadow_instance_groups                     |
| shadow_instance_id_mappings                |
| shadow_instance_info_caches                |
| shadow_instance_metadata                   |
| shadow_instance_system_metadata            |
| shadow_instance_type_extra_specs           |
| shadow_instance_type_projects              |
| shadow_instance_types                      |
| shadow_instances                           |
| shadow_key_pairs                           |
| shadow_migrate_version                     |
| shadow_migrations                          |
| shadow_networks                            |
| shadow_pci_devices                         |
| shadow_project_user_quotas                 |
| shadow_provider_fw_rules                   |
| shadow_quota_classes                       |
| shadow_quota_usages                        |
| shadow_quotas                              |
| shadow_reservations                        |
| shadow_s3_images                           |
| shadow_security_group_default_rules        |
| shadow_security_group_instance_association |
| shadow_security_group_rules                |
| shadow_security_groups                     |
| shadow_services                            |
| shadow_snapshot_id_mappings                |
| shadow_snapshots                           |
| shadow_task_log                            |
| shadow_virtual_interfaces                  |
| shadow_volume_id_mappings                  |
| shadow_volume_usage_cache                  |
| snapshot_id_mappings                       |
| snapshots                                  |
| tags                                       |
| task_log                                   |
| virtual_interfaces                         |
| volume_id_mappings                         |
| volume_usage_cache                         |
+--------------------------------------------+
```

```
[vagrant@localhost ~]$ sudo su -s /bin/sh -c "nova-manage db sync" nova
/usr/lib/python2.7/site-packages/pymysql/cursors.py:166: Warning: (1831, u'Duplicate index `block_device_mapping_instance_uuid_virtual_name_device_name_idx`. This is deprecated and will be disallowed in a future release.')
  result = self._query(query)
/usr/lib/python2.7/site-packages/pymysql/cursors.py:166: Warning: (1831, u'Duplicate index `uniq_instances0uuid`. This is deprecated and will be disallowed in a future release.')
  result = self._query(query)
```

```
[vagrant@localhost ~]$ mysql -u root -e "show tables in nova;"
+--------------------------------------------+
| Tables_in_nova                             |
+--------------------------------------------+
| agent_builds                               |
| aggregate_hosts                            |
| aggregate_metadata                         |
| aggregates                                 |
| allocations                                |
| block_device_mapping                       |
| bw_usage_cache                             |
| cells                                      |
| certificates                               |
| compute_nodes                              |
| console_auth_tokens                        |
| console_pools                              |
| consoles                                   |
| dns_domains                                |
| fixed_ips                                  |
| floating_ips                               |
| instance_actions                           |
| instance_actions_events                    |
| instance_extra                             |
| instance_faults                            |
| instance_group_member                      |
| instance_group_policy                      |
| instance_groups                            |
| instance_id_mappings                       |
| instance_info_caches                       |
| instance_metadata                          |
| instance_system_metadata                   |
| instance_type_extra_specs                  |
| instance_type_projects                     |
| instance_types                             |
| instances                                  |
| inventories                                |
| key_pairs                                  |
| migrate_version                            |
| migrations                                 |
| networks                                   |
| pci_devices                                |
| project_user_quotas                        |
| provider_fw_rules                          |
| quota_classes                              |
| quota_usages                               |
| quotas                                     |
| reservations                               |
| resource_provider_aggregates               |
| resource_providers                         |
| s3_images                                  |
| security_group_default_rules               |
| security_group_instance_association        |
| security_group_rules                       |
| security_groups                            |
| services                                   |
| shadow_agent_builds                        |
| shadow_aggregate_hosts                     |
| shadow_aggregate_metadata                  |
| shadow_aggregates                          |
| shadow_block_device_mapping                |
| shadow_bw_usage_cache                      |
| shadow_cells                               |
| shadow_certificates                        |
| shadow_compute_nodes                       |
| shadow_console_pools                       |
| shadow_consoles                            |
| shadow_dns_domains                         |
| shadow_fixed_ips                           |
| shadow_floating_ips                        |
| shadow_instance_actions                    |
| shadow_instance_actions_events             |
| shadow_instance_extra                      |
| shadow_instance_faults                     |
| shadow_instance_group_member               |
| shadow_instance_group_policy               |
| shadow_instance_groups                     |
| shadow_instance_id_mappings                |
| shadow_instance_info_caches                |
| shadow_instance_metadata                   |
| shadow_instance_system_metadata            |
| shadow_instance_type_extra_specs           |
| shadow_instance_type_projects              |
| shadow_instance_types                      |
| shadow_instances                           |
| shadow_key_pairs                           |
| shadow_migrate_version                     |
| shadow_migrations                          |
| shadow_networks                            |
| shadow_pci_devices                         |
| shadow_project_user_quotas                 |
| shadow_provider_fw_rules                   |
| shadow_quota_classes                       |
| shadow_quota_usages                        |
| shadow_quotas                              |
| shadow_reservations                        |
| shadow_s3_images                           |
| shadow_security_group_default_rules        |
| shadow_security_group_instance_association |
| shadow_security_group_rules                |
| shadow_security_groups                     |
| shadow_services                            |
| shadow_snapshot_id_mappings                |
| shadow_snapshots                           |
| shadow_task_log                            |
| shadow_virtual_interfaces                  |
| shadow_volume_id_mappings                  |
| shadow_volume_usage_cache                  |
| snapshot_id_mappings                       |
| snapshots                                  |
| tags                                       |
| task_log                                   |
| virtual_interfaces                         |
| volume_id_mappings                         |
| volume_usage_cache                         |
+--------------------------------------------+
```

```
[vagrant@localhost ~]$ nova-manage cell_v2 list_cells
Traceback (most recent call last):
  File "/usr/bin/nova-manage", line 10, in <module>
    sys.exit(main())
  File "/usr/lib/python2.7/site-packages/nova/cmd/manage.py", line 1682, in main
    config.parse_args(sys.argv)
  File "/usr/lib/python2.7/site-packages/nova/config.py", line 52, in parse_args
    default_config_files=default_config_files)
  File "/usr/lib/python2.7/site-packages/oslo_config/cfg.py", line 2473, in __call__
    self._namespace._files_permission_denied)
oslo_config.cfg.ConfigFilesPermissionDeniedError: Failed to open some config files: /usr/share/nova/nova-dist.conf,/etc/nova/nova.conf
[vagrant@localhost ~]$ sudo nova-manage cell_v2 list_cells
+-------+--------------------------------------+-------------------------------------+--------------------------------------------------+
|  Name |                 UUID                 |            Transport URL            |               Database Connection                |
+-------+--------------------------------------+-------------------------------------+--------------------------------------------------+
| cell0 | 00000000-0000-0000-0000-000000000000 |                none:/               | mysql+pymysql://nova:****@10.64.33.64/nova_cell0 |
| cell1 | e85d2a01-4637-4b5c-abad-3345435fdcbd | rabbit://openstack:****@10.64.33.64 |    mysql+pymysql://nova:****@10.64.33.64/nova    |
+-------+--------------------------------------+-------------------------------------+--------------------------------------------------+
```


Start
```
[vagrant@localhost ~]$ sudo systemctl start openstack-nova-api.service
[vagrant@localhost ~]$ systemctl -l status openstack-nova-api.service
● openstack-nova-api.service - OpenStack Nova API Server
   Loaded: loaded (/usr/lib/systemd/system/openstack-nova-api.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 02:46:08 UTC; 12s ago
 Main PID: 31607 (nova-api)
   CGroup: /system.slice/openstack-nova-api.service
           ├─31607 /usr/bin/python2 /usr/bin/nova-api
           ├─31617 /usr/bin/python2 /usr/bin/nova-api
           └─31619 /usr/bin/python2 /usr/bin/nova-api
```

```
[vagrant@localhost ~]$ sudo systemctl start openstack-nova-consoleauth.service
[vagrant@localhost ~]$ systemctl -l status openstack-nova-consoleauth.service
● openstack-nova-consoleauth.service - OpenStack Nova VNC console auth Server
   Loaded: loaded (/usr/lib/systemd/system/openstack-nova-consoleauth.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 02:46:57 UTC; 8s ago
 Main PID: 31639 (nova-consoleaut)
   CGroup: /system.slice/openstack-nova-consoleauth.service
           └─31639 /usr/bin/python2 /usr/bin/nova-consoleauth
```

```
[vagrant@localhost ~]$ sudo systemctl start openstack-nova-scheduler.service
[vagrant@localhost ~]$ systemctl -l status openstack-nova-scheduler.service
● openstack-nova-scheduler.service - OpenStack Nova Scheduler Server
   Loaded: loaded (/usr/lib/systemd/system/openstack-nova-scheduler.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 02:47:19 UTC; 6s ago
 Main PID: 31661 (nova-scheduler)
   CGroup: /system.slice/openstack-nova-scheduler.service
           └─31661 /usr/bin/python2 /usr/bin/nova-scheduler
```

```
[vagrant@localhost ~]$ sudo systemctl start openstack-nova-conductor.service
[vagrant@localhost ~]$ systemctl -l status openstack-nova-conductor.service
● openstack-nova-conductor.service - OpenStack Nova Conductor Server
   Loaded: loaded (/usr/lib/systemd/system/openstack-nova-conductor.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 02:47:46 UTC; 7s ago
 Main PID: 31688 (nova-conductor)
   CGroup: /system.slice/openstack-nova-conductor.service
           └─31688 /usr/bin/python2 /usr/bin/nova-conductor
```

```
[vagrant@localhost ~]$ sudo systemctl start openstack-nova-novncproxy.service
[vagrant@localhost ~]$ systemctl -l status openstack-nova-novncproxy.service
● openstack-nova-novncproxy.service - OpenStack Nova NoVNC Proxy Server
   Loaded: loaded (/usr/lib/systemd/system/openstack-nova-novncproxy.service; disabled; vendor preset: disabled)
   Active: active (running) since Sat 2017-10-21 02:48:06 UTC; 9s ago
 Main PID: 31710 (nova-novncproxy)
   CGroup: /system.slice/openstack-nova-novncproxy.service
           └─31710 /usr/bin/python2 /usr/bin/nova-novncproxy --web /usr/share/novnc/
```

```
[vagrant@localhost ~]$ openstack compute service list
+----+------------------+-----------------------+----------+---------+-------+----------------------------+
| ID | Binary           | Host                  | Zone     | Status  | State | Updated At                 |
+----+------------------+-----------------------+----------+---------+-------+----------------------------+
|  3 | nova-consoleauth | localhost.localdomain | internal | enabled | up    | 2017-10-21T02:50:03.000000 |
|  4 | nova-scheduler   | localhost.localdomain | internal | enabled | up    | 2017-10-21T02:50:04.000000 |
|  5 | nova-conductor   | localhost.localdomain | internal | enabled | up    | 2017-10-21T02:50:01.000000 |
+----+------------------+-----------------------+----------+---------+-------+----------------------------+
```

Enable auto-starting
```
[vagrant@localhost ~]$ sudo systemctl enable openstack-nova-api.service openstack-nova-consoleauth.service openstack-nova-scheduler.service openstack-nova-conductor.service openstack-nova-novncproxy.service
Created symlink from /etc/systemd/system/multi-user.target.wants/openstack-nova-api.service to /usr/lib/systemd/system/openstack-nova-api.service.
Created symlink from /etc/systemd/system/multi-user.target.wants/openstack-nova-consoleauth.service to /usr/lib/systemd/system/openstack-nova-consoleauth.service.
Created symlink from /etc/systemd/system/multi-user.target.wants/openstack-nova-scheduler.service to /usr/lib/systemd/system/openstack-nova-scheduler.service.
Created symlink from /etc/systemd/system/multi-user.target.wants/openstack-nova-conductor.service to /usr/lib/systemd/system/openstack-nova-conductor.service.
Created symlink from /etc/systemd/system/multi-user.target.wants/openstack-nova-novncproxy.service to /usr/lib/systemd/system/openstack-nova-novncproxy.service.
```

Log
```
[vagrant@localhost ~]$ sudo tail /var/log/nova/nova-api.log
2017-10-21 02:50:11.207 31617 DEBUG nova.api.openstack.wsgi [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] Calling method '<bound method ServiceController.index of <nova.api.openstack.compute.services.ServiceController object at 0x76fd5d0>>' _process_stack /usr/lib/python2.7/site-packages/nova/api/openstack/wsgi.py:612
2017-10-21 02:50:11.217 31617 DEBUG oslo_db.sqlalchemy.engines [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] MySQL server mode set to STRICT_TRANS_TABLES,STRICT_ALL_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,TRADITIONAL,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION _check_effective_sql_mode /usr/lib/python2.7/site-packages/oslo_db/sqlalchemy/engines.py:285
2017-10-21 02:50:11.221 31617 DEBUG nova.compute.api [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] Found 2 cells: 00000000-0000-0000-0000-000000000000(cell0),e85d2a01-4637-4b5c-abad-3345435fdcbd(cell1) load_cells /usr/lib/python2.7/site-packages/nova/compute/api.py:236
2017-10-21 02:50:11.222 31617 DEBUG oslo_concurrency.lockutils [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] Lock "00000000-0000-0000-0000-000000000000" acquired by "nova.context.get_or_set_cached_cell_and_set_connections" :: waited 0.000s inner /usr/lib/python2.7/site-packages/oslo_concurrency/lockutils.py:270
2017-10-21 02:50:11.222 31617 DEBUG oslo_concurrency.lockutils [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] Lock "00000000-0000-0000-0000-000000000000" released by "nova.context.get_or_set_cached_cell_and_set_connections" :: held 0.001s inner /usr/lib/python2.7/site-packages/oslo_concurrency/lockutils.py:282
2017-10-21 02:50:11.229 31617 DEBUG oslo_db.sqlalchemy.engines [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] MySQL server mode set to STRICT_TRANS_TABLES,STRICT_ALL_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,TRADITIONAL,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION _check_effective_sql_mode /usr/lib/python2.7/site-packages/oslo_db/sqlalchemy/engines.py:285
2017-10-21 02:50:11.247 31617 DEBUG oslo_concurrency.lockutils [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] Lock "e85d2a01-4637-4b5c-abad-3345435fdcbd" acquired by "nova.context.get_or_set_cached_cell_and_set_connections" :: waited 0.000s inner /usr/lib/python2.7/site-packages/oslo_concurrency/lockutils.py:270
2017-10-21 02:50:11.248 31617 DEBUG oslo_concurrency.lockutils [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] Lock "e85d2a01-4637-4b5c-abad-3345435fdcbd" released by "nova.context.get_or_set_cached_cell_and_set_connections" :: held 0.001s inner /usr/lib/python2.7/site-packages/oslo_concurrency/lockutils.py:282
2017-10-21 02:50:11.255 31617 DEBUG oslo_db.sqlalchemy.engines [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] MySQL server mode set to STRICT_TRANS_TABLES,STRICT_ALL_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,TRADITIONAL,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION _check_effective_sql_mode /usr/lib/python2.7/site-packages/oslo_db/sqlalchemy/engines.py:285
2017-10-21 02:50:11.270 31617 INFO nova.osapi_compute.wsgi.server [req-45e5d1f9-ac03-479b-98ec-d138829e500c 44e6ee1df8ae436986d2d50f7b358aa0 a0be38aef8c74d4abca3e4e100ee7910 - default default] 10.64.33.64 "GET /v2.1/os-services HTTP/1.1" status: 200 len: 1006 time: 0.1474090
```

```
[vagrant@localhost ~]$ sudo tail /var/log/nova/nova-consoleauth.log
2017-10-21 02:46:58.023 31639 DEBUG oslo_service.service [req-0707be05-80c8-4a5b-aa01-4502a3854994 - - - - -] upgrade_levels.scheduler       = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 02:46:58.024 31639 DEBUG oslo_service.service [req-0707be05-80c8-4a5b-aa01-4502a3854994 - - - - -] key_manager.api_class          = castellan.key_manager.barbican_key_manager.BarbicanKeyManager log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 02:46:58.024 31639 DEBUG oslo_service.service [req-0707be05-80c8-4a5b-aa01-4502a3854994 - - - - -] key_manager.fixed_key          = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 02:46:58.024 31639 DEBUG oslo_service.service [req-0707be05-80c8-4a5b-aa01-4502a3854994 - - - - -] osapi_v21.project_id_regex     = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 02:46:58.024 31639 DEBUG oslo_service.service [req-0707be05-80c8-4a5b-aa01-4502a3854994 - - - - -] ******************************************************************************** log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2889
2017-10-21 02:46:58.025 31639 INFO nova.service [-] Starting consoleauth node (version 16.0.1-1.el7)
2017-10-21 02:46:58.039 31639 DEBUG oslo_db.sqlalchemy.engines [req-9d4252ab-f7b8-41c1-83a5-6edca4980a5b - - - - -] MySQL server mode set to STRICT_TRANS_TABLES,STRICT_ALL_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,TRADITIONAL,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION _check_effective_sql_mode /usr/lib/python2.7/site-packages/oslo_db/sqlalchemy/engines.py:285
2017-10-21 02:46:58.178 31639 DEBUG nova.service [req-9d4252ab-f7b8-41c1-83a5-6edca4980a5b - - - - -] Creating RPC server for service consoleauth start /usr/lib/python2.7/site-packages/nova/service.py:166
2017-10-21 02:46:58.213 31639 DEBUG nova.service [req-9d4252ab-f7b8-41c1-83a5-6edca4980a5b - - - - -] Join ServiceGroup membership for this service consoleauth start /usr/lib/python2.7/site-packages/nova/service.py:184
2017-10-21 02:46:58.214 31639 DEBUG nova.servicegroup.drivers.db [req-9d4252ab-f7b8-41c1-83a5-6edca4980a5b - - - - -] DB_Driver: join new ServiceGroup member localhost.localdomain to the consoleauth group, service = <Service: host=localhost.localdomain, binary=nova-consoleauth, manager_class_name=nova.consoleauth.manager.ConsoleAuthManager> join /usr/lib/python2.7/site-packages/nova/servicegroup/drivers/db.py:47
```

```
[vagrant@localhost ~]$ sudo tail /var/log/nova/nova-scheduler.log
2017-10-21 02:48:19.503 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._expire_reservations run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
2017-10-21 02:48:30.953 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._run_periodic_tasks run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
2017-10-21 02:49:21.500 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._expire_reservations run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
2017-10-21 02:49:32.950 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._run_periodic_tasks run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
2017-10-21 02:50:22.501 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._expire_reservations run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
2017-10-21 02:50:32.959 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._run_periodic_tasks run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
2017-10-21 02:51:22.502 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._expire_reservations run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
2017-10-21 02:51:34.957 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._run_periodic_tasks run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
2017-10-21 02:52:24.506 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._expire_reservations run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
2017-10-21 02:52:35.953 31661 DEBUG oslo_service.periodic_task [req-9543dfb9-26a5-4c40-8d12-ae028e0f8e43 - - - - -] Running periodic task SchedulerManager._run_periodic_tasks run_periodic_tasks /usr/lib/python2.7/site-packages/oslo_service/periodic_task.py:215
```

```
[vagrant@localhost ~]$ sudo tail /var/log/nova/nova-conductor.log
2017-10-21 02:47:46.493 31688 DEBUG oslo_service.service [req-69c441ee-b49c-45eb-ac77-c372e1d21d80 - - - - -] upgrade_levels.scheduler       = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 02:47:46.493 31688 DEBUG oslo_service.service [req-69c441ee-b49c-45eb-ac77-c372e1d21d80 - - - - -] key_manager.api_class          = castellan.key_manager.barbican_key_manager.BarbicanKeyManager log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 02:47:46.493 31688 DEBUG oslo_service.service [req-69c441ee-b49c-45eb-ac77-c372e1d21d80 - - - - -] key_manager.fixed_key          = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 02:47:46.493 31688 DEBUG oslo_service.service [req-69c441ee-b49c-45eb-ac77-c372e1d21d80 - - - - -] osapi_v21.project_id_regex     = None log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2887
2017-10-21 02:47:46.494 31688 DEBUG oslo_service.service [req-69c441ee-b49c-45eb-ac77-c372e1d21d80 - - - - -] ******************************************************************************** log_opt_values /usr/lib/python2.7/site-packages/oslo_config/cfg.py:2889
2017-10-21 02:47:46.494 31688 INFO nova.service [-] Starting conductor node (version 16.0.1-1.el7)
2017-10-21 02:47:46.507 31688 DEBUG oslo_db.sqlalchemy.engines [req-5e4eac74-c4f2-4aed-95b1-0a656e9e639f - - - - -] MySQL server mode set to STRICT_TRANS_TABLES,STRICT_ALL_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,TRADITIONAL,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION _check_effective_sql_mode /usr/lib/python2.7/site-packages/oslo_db/sqlalchemy/engines.py:285
2017-10-21 02:47:46.713 31688 DEBUG nova.service [req-5e4eac74-c4f2-4aed-95b1-0a656e9e639f - - - - -] Creating RPC server for service conductor start /usr/lib/python2.7/site-packages/nova/service.py:166
2017-10-21 02:47:46.731 31688 DEBUG nova.service [req-5e4eac74-c4f2-4aed-95b1-0a656e9e639f - - - - -] Join ServiceGroup membership for this service conductor start /usr/lib/python2.7/site-packages/nova/service.py:184
2017-10-21 02:47:46.731 31688 DEBUG nova.servicegroup.drivers.db [req-5e4eac74-c4f2-4aed-95b1-0a656e9e639f - - - - -] DB_Driver: join new ServiceGroup member localhost.localdomain to the conductor group, service = <Service: host=localhost.localdomain, binary=nova-conductor, manager_class_name=nova.conductor.manager.ConductorManager> join /usr/lib/python2.7/site-packages/nova/servicegroup/drivers/db.py:47
```

```
[vagrant@localhost ~]$ sudo tail /var/log/nova/nova-novncproxy.log
2017-10-21 02:48:08.005 31710 WARNING oslo_reports.guru_meditation_report [-] Guru meditation now registers SIGUSR1 and SIGUSR2 by default for backward compatibility. SIGUSR1 will no longer be registered in a future release, so please use SIGUSR2 to generate reports.
2017-10-21 02:48:08.006 31710 INFO nova.console.websocketproxy [-] WebSocket server settings:
2017-10-21 02:48:08.006 31710 INFO nova.console.websocketproxy [-]   - Listen on 0.0.0.0:6080
2017-10-21 02:48:08.006 31710 INFO nova.console.websocketproxy [-]   - Flash security policy server
2017-10-21 02:48:08.006 31710 INFO nova.console.websocketproxy [-]   - Web server (no directory listings). Web root: /usr/share/novnc
2017-10-21 02:48:08.006 31710 INFO nova.console.websocketproxy [-]   - No SSL/TLS support (no cert file)
2017-10-21 02:48:08.007 31710 INFO nova.console.websocketproxy [-]   - proxying from 0.0.0.0:6080 to None:None
```
