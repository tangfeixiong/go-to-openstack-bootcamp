# Iteration 01: OpenStack开发环境                                       
## Sprint 01: OpenStack开发平台                                        
### 学习                                                            

* 安装Ubuntu或CentOS                                                   

  windows下可以使用虚拟机，虚拟化平台可以使用VirtualBox或VMware Player（或VMware Workstation）
  - [Ubuntu Mirrors](https://launchpad.net/ubuntu/+archivemirrors)
    ![domestic](/image/domestic-ubuntu-mirrors.png)
    For example:
    ![isos-in-aliyun](/image/aliyun-ubuntu-14.04.1.png)
  - [CentOS Mirrors](http://www.centos.org/download/mirrors/)
    ![domestic](/image/domestic-centos-mirrors.png)
    For example:
    ![isos-in-aliyun](/image/aliyun-centos-6.5-isos-x86_64.png)
  - [Virtual Box](https://www.virtualbox.org/wiki/Downloads) 
    Tutorial: [How to Install Ubuntu on VirtualBox (with Pictures) - wikiHow](http://www.wikihow.com/Install-Ubuntu-on-VirtualBox)
  - [VMware Player](https://my.vmware.com/web/vmware/free#desktop_end_user_computing/vmware_player/6_0)
    Tutorial: [CentOS 6.3 Step by Step Installation Guide with Screenshots](http://www.tecmint.com/centos-6-3-step-by-step-installation-guide-with-screenshots/)
  - http://www.baidu.com 

* 设置OpenStack开发环境                                                  
  1. 使用devstack安装icehouse或juno，也可以安装havana

    devstack是社区推荐的openstack单机(or multi-node)开发环境
    - [DevStack developer](https://wiki.openstack.org/wiki/DevStack)
    - [Source code in Github](https://github.com/openstack-dev/devstack)
    - [Stack Tutorial](http://docs.openstack.org/developer/devstack/)
    
  2. 掌握bash命令screen

    devstack在用户会话模式下安装openstack，通过screen命令切换会话窗口，可以检查openstack各组件的控制台日志
    - [screen GNU user manual](http://www.gnu.org/software/screen/manual/screen.html)

* 练习OpenStack操作                                                     
  - [End User Guide](http://docs.openstack.org/user-guide/content/)
  - [Admin User Guide](http://docs.openstack.org/user-guide-admin/content/)
  - [OpenStack Documentation](http://docs.openstack.org/)

* 掌握Apache HTTPD在Linux下的管理和OpenStack下的使用                           
  - [for Ubuntu](https://help.ubuntu.com/14.04/serverguide/httpd.html)
  - [for CentOS](http://www.centos.org/docs/5/html/Deployment_Guide-en-US/ch-httpd.html)
  - [for OpenStack](http://docs.openstack.org/juno/install-guide/install/apt/content/install_dashboard.html)
		
	
## Sprint 02: Python开发环境                                            
### 学习                                                                
  * Python2.6/2.7及开发工具                                           
    - [Python 2.7](https://docs.python.org/2.7/)
    - [调试器](https://docs.python.org/2.7/library/debug.html)
		
  * Python SDK (PyPI和Pip)                                            
    - [SDK目录](https://pypi.python.org/pypi)
    - [SDK在线安装工具](https://pip.pypa.io/en/latest/installing.html#install-pip)
    - [pip vs easy_install](https://packaging.python.org/en/latest/technical.html#pip-vs-easy-install) 
		
  * 开发环境和生产环境                                                
    - [virtualenv](https://virtualenv.pypa.io/en/latest/virtualenv.html#installation)
    - [staging](https://docs.python.org/2.7/using/index.html)
    - [production](https://docs.python.org/2/library/site.html)
		
  * OpenStack开发环境                                                 
    1 Nova及相关项目
      - [源代码](https://github.com/openstack/nova)
      - [wiki](https://wiki.openstack.org/wiki/Nova)
      - [developer](http://docs.openstack.org/developer/nova/)
        a. Setup Environment
        b. Unit Tests
        ...
    2 所有项目
      - [Developer Documentation](http://docs.openstack.org/developer/openstack-projects.html)
      - [源代码](https://github.com/openstack)
      - [单元测试] (https://wiki.openstack.org/wiki/Testing#Unit_Tests)

### Written
  * [manually setup python 2.7.8 developer environment](/devel-env.rst)     
	

# Iteration 02: 需求分析                                                12月份~1月份
## Sprint 01: 租户账户管理                                              1周
### 学习                                                                1d
    * Keystone
	    - [wiki](https://wiki.openstack.org/wiki/Keystone)
		- [developer doc](http://docs.openstack.org/developer/keystone/)
	* 工具
        - [curl](http://curl.haxx.se/docs/manpage.html)
	    - [soap ui](http://www.soapui.org/)
		
### 实验                                                                2d
    * MySQL和Keystone
	    准备一个vm，安装MySQL和Keystone，使用MySQL client工具，keystone源代码掌握keystone数据库
		- [keystone service installation](http://docs.openstack.org/juno/install-guide/install/yum/content/ch_keystone.html)
		- [MySQL](http://dev.mysql.com/doc/)
	* MySQL的分发
	    在线分发
	        - [yum repository参考](http://dev.mysql.com/downloads/repo/yum/)
		    - [apt repository参考](http://dev.mysql.com/downloads/repo/apt/)
		制作本地分发仓库
		    在CentOS，在线分发的RPM包在/etc/cache下缓存，以该RPM用于制作本地分发包
		    - [for CentOS 6](http://wiki.centos.org/HowTos/CreateLocalRepos)
			- [RedHat yum doc](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/6/html/Deployment_Guide/sec-Configuring_Yum_and_Yum_Repositories.html)
			- [RedHat httpd doc](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/6/html/Deployment_Guide/ch-Web_Servers.html)
			- [using Apache Httpd for Ubuntu](http://www.unixmen.com/setup-local-repository-ubuntu-14-0413-1013-04-server/)




	

## Sprint 02: 虚拟机申请（X86_64）                                      1周
### 学习                                                                1d
    * Glance
	    - 掌握Glance数据库
		- 掌握Glance API，Glance Registry 服务
		- 掌握python glance client
	    - 掌握Glance支持Local File System与Object Storage的技术
	
### 实验                                                                2d
    * Glance的分发
	    Glance在RDO的在线分发
		    - [RDO](https://openstack.redhat.com/Main_Page)
			- [RDO Repo](https://repos.fedorapeople.org/repos/openstack/)
			- [RedHat OpenStack Production](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux_OpenStack_Platform/)
		制作Glance的本地分发仓库
            - RPM或Debian			
    


	

## Sprint 03: 虚拟机管理（X86_64）                                      2周
### 学习                                                                3d
    * Nova
	    - 掌握Nova数据库
		- 掌握Nova API, Nova Scheduler, Nova Conductor,	Nova Cert, Nova ConsoleAuth, Nova NoVNCProxy, 和Nova Compute服务
		- 掌握Nova Libvirt driver, Nova VMware vSphere Driver
		- 掌握python nova client
	    - 掌握Nova支持虚拟机实例的存储系统的技术
		- 研究VMware vSphere 5.1以下版本和5.5版本的支持区别
		- 掌握多Region，多Cell，以及多Zone，多Aggregate的原理
		- [Tenant](http://docs.openstack.org/openstack-ops/content/projects_users.html)
		- [User facing operation](http://docs.openstack.org/openstack-ops/content/user_facing_operations.html)
	
### 实验                                                                3d
    * Nova的分发
	    Nova在RDO的在线分发
		    - 控制节点上分发
			- 计算节点上分发（单计算节点和多计算节点）
		制作Nova的本地分发仓库
            - RPM或Debian		
    

	

	

## Sprint 04: 租户网络管理（X86_64）                                    2周
### 学习                                                                3d
    * Neutron
	    - 掌握Neutron数据库
		- 掌握Neutron Server, Modular Layer 2, Neutron OpenVSwitch Agent, Neutron L3 Agent, Neutron DHCP Agent服务
		- 掌握OpenVSwitch, VMware vSwitch, VMware vDSwitch, VMware NSX, Linux Bridge
		- 掌握python neutron client
	    - 掌握neutron支持flat和vlan的技术
		- 研究VMware vSphere 5.1以下版本和5.5版本的支持区别
		- 掌握多Region，多Cell，以及多Zone，多Aggregate的原理
	* Nova-Network服务
		- [Legacy Networking和OpenStack Networking](http://docs.openstack.org/openstack-ops/content/example_architecture.html)
	
### 实验                                                                3d
    * Neutron和Nova-Network的分发
	    Neutron在RDO的在线分发
		    - 控制节点上分发
			- 网络节点上分发
			- 计算节点上分发
		制作Neutron和Nova-Network的本地分发仓库
            - RPM或Debian		
    

	

## Sprint 05: 计量管理                                                  1周
### 学习                                                                1.5d
    * Ceilometer
	    - (http://docs.openstack.org/juno/install-guide/install/yum/content/ch_ceilometer.html)
	    - (http://docs.openstack.org/developer/ceilometer/)
		- (http://docs.openstack.org/admin-guide-cloud/content/ch_admin-openstack-telemetry.html)
	
### 实验                                                                1.5d
    * Ceilometer的分发
	    Ceilometer在RDO的在线分发
		    - 控制节点上分发
			- 网络节点上分发
			- 计算节点上分发
		制作Ceilometer的本地分发仓库
            - RPM或Debian		
    

	

## Sprint 06: 存储管理，协调管理，数据库管理等                          1周
### 学习                                                                1.5d
    * Cinder
	    - Cinder api和Cinder Scheduler
        - Open iSCSI和NFS为后端存储
	    - 分布式存储为后端
	* Swift
	* Heat
	* Trove
	
### 实验                                                                1.5d
    * Cinder的分发
	    官方在线分发
		    - 控制节点上分发
			- 块存储节点上分发
		制作Cinder的本地分发仓库
            - RPM或Debian		   
	* Swift的分发
	* Heat的分发
	* Trove的分发
	

	
# Iteration 03: 功能设计                                                2月份
## Sprint 01：                                                          2周
### 用户自服务功能                                                      2d

### 管理仪表盘功能                                                      3d

### 统计和报告接口                                                      2d

	
### 开发语言和工具                                                      1d



	
## Sprint 02：                                                          0.5周
### 自动化部署



# Iteration 04: 软件开发                                                3月份~

	
# Iteration 05：Staging
	
	
# Iteration 06：CI	
## Sprint 01: 告警监控（物理机）                                          
### 研究Zabbix                                                  
    * Zabbix
	    - [官方文档](http://www.zabbix.com/documentation.php)
	    - [社区](http://www.zabbix.org/wiki/Main_Page)
	* IPMI
	* Nagios

	
## Sprint 02: 服务自动化（应用软件）                                          
### Puppet或Salt Stack


## Sprint 03：服务自动化（Oracle，WebLogic等）
### 集群自动化部署


### 补丁


### 升级


## Sprint 04：服务自动化（操作系统）
### PXE


### 集群


### 补丁







  
软件分发
====================================

标准介质
------------------------------------

	
在线安装
------------------------------------



资源监控
====================================




资源配置
====================================



资源服务
====================================




开放接口
====================================
