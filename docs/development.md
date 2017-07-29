
## OpenStack Development

### Provider

![openstack-dev-architect.png](./OpenStack-dev-architect.png)

Overview

* 192.168.1.0/24 办公网络
* 其它网络 datacenter网络

DataCenter networking

1. 10.121.0.0/24 为Openstack的data network（作为租户网络的隧道，即tun）
1. 10.121.198.0/24 为Openstack的management network，内部ReST API网络
1. 10.100.151.0/24是数据中心网络中与外部通信的datacenter公共网络，是物理网络，租户网络是虚拟网络，物理世界与虚拟世界的连通通过第三块网卡，This NIC would work in promiscuous mode(No its own ip address), 如果第三块网卡设置了ip地址，就成了终端网络单元（NU），而不是中继网络单元
1. 为方便开发，10.100.151.0/24地址范围太小，需要扩展为B类子网10.100.0.0/16
1. ReST API在10.121.198.0/24，不能被datacenter公共网络访问，由于controller node不需要隧道，因此，The 2nd NIC可以配置到10.100.0.0/16公共网络（例如10.100.12.34），再安装L4或L7代理，如nginx或HAProxy，将5000，9292等端口服务代理到10.121.198.2的对应端口，此即IaaS API网络
1. 开发人员在192.168.1.0/24办公网络，还是无法访问datacenter公共网络，controller node的第三网口配置到192.168.1.0/24，同样进行L4/7代理，如果没有多余的网口，可以在第二网口上设置多IP地址，例如：192.168.1.100，则开发者的openstack service endpoint变为192.168.1.100:5000，:9292等等
1. 这里controller node按5、6两条，在扮演俗话所说的跳板机，此情况下，该跳板机配置尽量多的网口必要，可以买[2端口网卡](http://re.jd.com/search?keyword=%E5%8F%8C%E5%8F%A3%E5%8D%83%E5%85%86%E7%BD%91%E5%8D%A1pcie&keywordid=28125493460&re_dcp=202m0QjIIg==&traffic_source=1004&test=1&enc=utf8&cu=true&utm_source=baidu-search&utm_medium=cpc&utm_campaign=t_262767352_baidusearch&utm_term=28125493460_0_928f830849a5451280398a31328e5ecd&ad_od=1)。而公有云场景下，需要独立的跳板机，至少两个端口，一个配置公网ip地址，一个配置到10.121.198.0/24网段。
1. controller node在datacenter的public网络上，因此其网关是10.100.151.253，方便openstack devops在线安装和升级软件
1. 但其它network node和compute node没有在public网络上，如果网络端口足够，也可以分配一个public网络的地址，当然也可以在一个网口上配置多个地址。这样，每个节点的网关均为10.100.151.253
1. 以上，对交换机端口数量要求较多，因此，也可将controller node设置为linux router，其可作为network node和computer node的网关（10.121.198.2），当它们需要对外请求，ip packet会发给controller node，controller node会选择默认路由的端口（10.100.12.34）转发，upstream给10.100.151.253
1. K版本后的openstack架构，为了减少设备投入，没有了network node，因此，controller node的端口中还需额外的tun和promiscuous端口
1. 如果现场实施采用network boot进行PXE，应事先规划
1. 在demo/simple/general网络设计上，没必要采用如此的数据中心开发环境架构，尽量简单些. For example, ![openstack-deploy-sample.png](./openstack-deploy-sample.png)

### Hack

http://www.tcpdump.org/

/Users/fanhongling/Downloads/go-kubernetes/src/github.com/ntop/PF_RING

Trouble shooting

MacOS
```
github.com/libvirt/libvirt-go
# pkg-config --cflags libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt
pkg-config: exec: "pkg-config": executable file not found in $PATH
github.com/libvirt/libvirt-go-xml
```

Fedora32
```
[vagrant@localhost go-to-openstack-bootcamp]$ go install -v ./kopos/kopit/
...snip
github.com/google/gopacket/pcap
# github.com/google/gopacket/pcap
/home/vagrant/go/src/github.com/google/gopacket/pcap/pcap.go:21:18: fatal error: pcap.h: No such file or directory
compilation terminated.
...snip
github.com/libvirt/libvirt-go
github.com/libvirt/libvirt-go-xml
# pkg-config --cflags libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt libvirt
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
Package libvirt was not found in the pkg-config search path.
Perhaps you should add the directory containing `libvirt.pc'
to the PKG_CONFIG_PATH environment variable
No package 'libvirt' found
pkg-config: exit status 1
...snip
```

Need __libpcap__
```
[vagrant@localhost go-to-openstack-bootcamp]$ sudo dnf search libpcap-devel
上次元数据过期检查在 2:06:05 前执行于 Sat Jul  8 08:59:29 2017。
============================================================ N/S 匹配：libpcap-devel ============================================================
libpcap-devel.i686 : Libraries and header files for the libpcap library
libpcap-devel.x86_64 : Libraries and header files for the libpcap library
[vagrant@localhost go-to-openstack-bootcamp]$ sudo dnf install -y libpcap-devel
上次元数据过期检查在 2:06:14 前执行于 Sat Jul  8 08:59:29 2017。
依赖关系解决。
=================================================================================================================================================
 Package                              架构                          版本                                     仓库                           大小
=================================================================================================================================================
安装:
 libpcap-devel                        x86_64                        14:1.7.4-1.fc23                          fedora                        122 k

事务概要
=================================================================================================================================================
安装  1 Package

总下载：122 k
安装大小：168 k
下载软件包：
libpcap-devel-1.7.4-1.fc23.x86_64.rpm                                                                             78 kB/s | 122 kB     00:01    
-------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                              42 kB/s | 122 kB     00:02     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: libpcap-devel-14:1.7.4-1.fc23.x86_64                                                                                                 1/1 
  验证: libpcap-devel-14:1.7.4-1.fc23.x86_64                                                                                                 1/1 

已安装:
  libpcap-devel.x86_64 14:1.7.4-1.fc23                                                                                                           

完毕！
```

Need __libvirt__
```
[vagrant@localhost go-to-openstack-bootcamp]$ sudo dnf search libvirt-devel
上次元数据过期检查在 1:54:49 前执行于 Sat Jul  8 08:59:29 2017。
============================================================ N/S 匹配：libvirt-devel ============================================================
libvirt-devel.i686 : Libraries, includes, etc. to compile with the libvirt library
libvirt-devel.x86_64 : Libraries, includes, etc. to compile with the libvirt library
ocaml-libvirt-devel.i686 : Development files for ocaml-libvirt
ocaml-libvirt-devel.x86_64 : Development files for ocaml-libvirt
[vagrant@localhost go-to-openstack-bootcamp]$ sudo dnf install -y libvirt-devel
上次元数据过期检查在 1:55:18 前执行于 Sat Jul  8 08:59:29 2017。
依赖关系解决。
=================================================================================================================================================
 Package                              架构                         版本                                      仓库                           大小
=================================================================================================================================================
安装:
 cyrus-sasl                           x86_64                       2.1.26-25.2.fc23                          fedora                         92 k
 cyrus-sasl-md5                       x86_64                       2.1.26-25.2.fc23                          fedora                         61 k
 gnutls-dane                          x86_64                       3.4.15-1.fc23                             updates                        39 k
 gnutls-utils                         x86_64                       3.4.15-1.fc23                             updates                       267 k
 libvirt-client                       x86_64                       1.2.18.4-1.fc23                           updates                       4.4 M
 libvirt-devel                        x86_64                       1.2.18.4-1.fc23                           updates                       178 k
 libwsman1                            x86_64                       2.6.0-2.fc23                              updates                       131 k
 nmap-ncat                            x86_64                       2:7.12-1.fc23                             updates                       218 k
 numactl-libs                         x86_64                       2.0.10-3.fc23                             fedora                         33 k
 trousers                             x86_64                       0.3.13-5.fc23                             fedora                        151 k
 trousers-lib                         x86_64                       0.3.13-5.fc23                             fedora                        158 k
 unbound-libs                         x86_64                       1.5.8-2.fc23                              updates                       372 k
 xen-libs                             x86_64                       4.5.5-5.fc23                              updates                       504 k
 xen-licenses                         x86_64                       4.5.5-5.fc23                              updates                       104 k
 yajl                                 x86_64                       2.1.0-4.fc23                              fedora                         37 k
升级:
 gnutls                               x86_64                       3.4.15-1.fc23                             updates                       663 k

事务概要
=================================================================================================================================================
安装  15 Packages
升级   1 Package

总下载：7.3 M
下载软件包：
(1/16): cyrus-sasl-2.1.26-25.2.fc23.x86_64.rpm                                                                    94 kB/s |  92 kB     00:00    
(2/16): libvirt-devel-1.2.18.4-1.fc23.x86_64.rpm                                                                 180 kB/s | 178 kB     00:00    
(3/16): yajl-2.1.0-4.fc23.x86_64.rpm                                                                              35 kB/s |  37 kB     00:01    
(4/16): numactl-libs-2.0.10-3.fc23.x86_64.rpm                                                                     28 kB/s |  33 kB     00:01    
(5/16): gnutls-dane-3.4.15-1.fc23.x86_64.rpm                                                                      44 kB/s |  39 kB     00:00    
(6/16): gnutls-utils-3.4.15-1.fc23.x86_64.rpm                                                                    251 kB/s | 267 kB     00:01    
(7/16): cyrus-sasl-md5-2.1.26-25.2.fc23.x86_64.rpm                                                               159 kB/s |  61 kB     00:00    
(8/16): libwsman1-2.6.0-2.fc23.x86_64.rpm                                                                        334 kB/s | 131 kB     00:00    
(9/16): nmap-ncat-7.12-1.fc23.x86_64.rpm                                                                         275 kB/s | 218 kB     00:00    
(10/16): unbound-libs-1.5.8-2.fc23.x86_64.rpm                                                                    161 kB/s | 372 kB     00:02    
(11/16): xen-licenses-4.5.5-5.fc23.x86_64.rpm                                                                    282 kB/s | 104 kB     00:00    
(12/16): trousers-0.3.13-5.fc23.x86_64.rpm                                                                       129 kB/s | 151 kB     00:01    
(13/16): trousers-lib-0.3.13-5.fc23.x86_64.rpm                                                                   212 kB/s | 158 kB     00:00    
(14/16): xen-libs-4.5.5-5.fc23.x86_64.rpm                                                                         93 kB/s | 504 kB     00:05    
(15/16): gnutls-3.4.15-1.fc23.x86_64.rpm                                                                         151 kB/s | 663 kB     00:04    
(16/16): libvirt-client-1.2.18.4-1.fc23.x86_64.rpm                                                               245 kB/s | 4.4 MB     00:18    
-------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                             314 kB/s | 7.3 MB     00:23     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  升级: gnutls-3.4.15-1.fc23.x86_64                                                                                                         1/17 
  安装: yajl-2.1.0-4.fc23.x86_64                                                                                                            2/17 
  安装: trousers-lib-0.3.13-5.fc23.x86_64                                                                                                   3/17 
  安装: xen-licenses-4.5.5-5.fc23.x86_64                                                                                                    4/17 
  安装: xen-libs-4.5.5-5.fc23.x86_64                                                                                                        5/17 
  安装: unbound-libs-1.5.8-2.fc23.x86_64                                                                                                    6/17 
  安装: gnutls-dane-3.4.15-1.fc23.x86_64                                                                                                    7/17 
  安装: gnutls-utils-3.4.15-1.fc23.x86_64                                                                                                   8/17 
  安装: nmap-ncat-2:7.12-1.fc23.x86_64                                                                                                      9/17 
  安装: libwsman1-2.6.0-2.fc23.x86_64                                                                                                      10/17 
  安装: cyrus-sasl-md5-2.1.26-25.2.fc23.x86_64                                                                                             11/17 
  安装: numactl-libs-2.0.10-3.fc23.x86_64                                                                                                  12/17 
  安装: cyrus-sasl-2.1.26-25.2.fc23.x86_64                                                                                                 13/17 
  安装: libvirt-client-1.2.18.4-1.fc23.x86_64                                                                                              14/17 
  安装: libvirt-devel-1.2.18.4-1.fc23.x86_64                                                                                               15/17 
  安装: trousers-0.3.13-5.fc23.x86_64                                                                                                      16/17 
  清理: gnutls-3.4.8-1.fc23.x86_64                                                                                                         17/17 
  验证: libvirt-devel-1.2.18.4-1.fc23.x86_64                                                                                                1/17 
  验证: libvirt-client-1.2.18.4-1.fc23.x86_64                                                                                               2/17 
  验证: cyrus-sasl-2.1.26-25.2.fc23.x86_64                                                                                                  3/17 
  验证: numactl-libs-2.0.10-3.fc23.x86_64                                                                                                   4/17 
  验证: yajl-2.1.0-4.fc23.x86_64                                                                                                            5/17 
  验证: gnutls-utils-3.4.15-1.fc23.x86_64                                                                                                   6/17 
  验证: gnutls-dane-3.4.15-1.fc23.x86_64                                                                                                    7/17 
  验证: cyrus-sasl-md5-2.1.26-25.2.fc23.x86_64                                                                                              8/17 
  验证: libwsman1-2.6.0-2.fc23.x86_64                                                                                                       9/17 
  验证: nmap-ncat-2:7.12-1.fc23.x86_64                                                                                                     10/17 
  验证: xen-libs-4.5.5-5.fc23.x86_64                                                                                                       11/17 
  验证: unbound-libs-1.5.8-2.fc23.x86_64                                                                                                   12/17 
  验证: xen-licenses-4.5.5-5.fc23.x86_64                                                                                                   13/17 
  验证: trousers-0.3.13-5.fc23.x86_64                                                                                                      14/17 
  验证: trousers-lib-0.3.13-5.fc23.x86_64                                                                                                  15/17 
  验证: gnutls-3.4.15-1.fc23.x86_64                                                                                                        16/17 
  验证: gnutls-3.4.8-1.fc23.x86_64                                                                                                         17/17 

已安装:
  cyrus-sasl.x86_64 2.1.26-25.2.fc23            cyrus-sasl-md5.x86_64 2.1.26-25.2.fc23            gnutls-dane.x86_64 3.4.15-1.fc23               
  gnutls-utils.x86_64 3.4.15-1.fc23             libvirt-client.x86_64 1.2.18.4-1.fc23             libvirt-devel.x86_64 1.2.18.4-1.fc23           
  libwsman1.x86_64 2.6.0-2.fc23                 nmap-ncat.x86_64 2:7.12-1.fc23                    numactl-libs.x86_64 2.0.10-3.fc23              
  trousers.x86_64 0.3.13-5.fc23                 trousers-lib.x86_64 0.3.13-5.fc23                 unbound-libs.x86_64 1.5.8-2.fc23               
  xen-libs.x86_64 4.5.5-5.fc23                  xen-licenses.x86_64 4.5.5-5.fc23                  yajl.x86_64 2.1.0-4.fc23                       

已升级:
  gnutls.x86_64 3.4.15-1.fc23                                                                                                                    

完毕！
```


Need
```
apt-get install -y pkg-config lxc-dev
```



### DevOps

Ocata on CentOS7.3

