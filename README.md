OpenStack Boot Camp
====================

Kopos
-----

Kubernetes OPerate OpenStack

### Prerequisites to build *protoc*, *protoc-gen-go*, *protoc-gen-grpc-gateway*

protoc - see also https://github.com/tangfeixiong/go-to-Kubernetes/blob/master/docs/protobuf
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb$ protoc --version
libprotoc 3.3.0
```

protoc-gen-go
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/go-kubernetes/src/github.com/golang/protobuf$ go install -v ./protoc-gen-go/
github.com/golang/protobuf/proto
github.com/golang/protobuf/protoc-gen-go/descriptor
github.com/golang/protobuf/protoc-gen-go/plugin
github.com/golang/protobuf/protoc-gen-go/generator
github.com/golang/protobuf/protoc-gen-go/grpc
```

protoc-gen-grpc-gateway
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/go-kubernetes/src/github.com/grpc-ecosystem/grpc-gateway$ GOPATH=/Users/fanhongling/Downloads/go-kubernetes go build -v -o /home/vagrant/go/bin/protoc-gen-grpc-gateway ./protoc-gen-grpc-gateway/
github.com/golang/glog
github.com/golang/protobuf/proto
github.com/grpc-ecosystem/grpc-gateway/utilities
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/httprule
github.com/golang/protobuf/protoc-gen-go/descriptor
github.com/golang/protobuf/protoc-gen-go/plugin
google.golang.org/genproto/googleapis/api/annotations
github.com/golang/protobuf/protoc-gen-go/generator
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/generator
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/gengateway
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
```

protoc-gen-swagger
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/go-kubernetes/src/github.com/grpc-ecosystem/grpc-gateway$ GOPATH=/Users/fanhongling/Downloads/go-kubernetes go build -v -o /home/vagrant/go/bin/protoc-gen-swagger ./protoc-gen-swagger/
github.com/golang/protobuf/proto
github.com/golang/glog
github.com/grpc-ecosystem/grpc-gateway/utilities
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/httprule
github.com/golang/protobuf/protoc-gen-go/descriptor
github.com/golang/protobuf/protoc-gen-go/plugin
google.golang.org/genproto/googleapis/api/annotations
github.com/golang/protobuf/protoc-gen-go/generator
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/generator
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/genswagger
github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

### Generate stub of *Protobuf*, *gRPC* and *gRPC-gateway*

Vagrant up (with synced-folder)
```
fanhonglingdeMacBook-Pro:trusty64 fanhongling$ vagrant up
Bringing machine 'default' up with 'virtualbox' provider...
==> default: Checking if box 'ubuntu/trusty64' is up to date...
==> default: A newer version of the box 'ubuntu/trusty64' is available! You currently
==> default: have version '20160122.0.0'. The latest is version '20170602.0.0'. Run
==> default: `vagrant box update` to update.
==> default: Clearing any previously set forwarded ports...
==> default: Clearing any previously set network interfaces...
==> default: Preparing network interfaces based on configuration...
    default: Adapter 1: nat
    default: Adapter 2: hostonly
==> default: Forwarding ports...
    default: 6443 (guest) => 6443 (host) (adapter 1)
    default: 22 (guest) => 2222 (host) (adapter 1)
==> default: Running 'pre-boot' VM customizations...
==> default: Booting VM...
==> default: Waiting for machine to boot. This may take a few minutes...
    default: SSH address: 127.0.0.1:2222
    default: SSH username: vagrant
    default: SSH auth method: private key
    default: Warning: Remote connection disconnect. Retrying...
    default: Warning: Remote connection disconnect. Retrying...
==> default: Machine booted and ready!
==> default: Checking for guest additions in VM...
==> default: Configuring and enabling network interfaces...
==> default: Mounting shared folders...
    default: /vagrant => /Users/fanhongling/https%3A%2F%2Fwww.vagrantup.com/https%3A%2F%2Fatlas.hashicorp.com/boxes/ubuntu/trusty64
    default: /Users/fanhongling/go/src => /Users/fanhongling/go/src
    default: /Users/fanhongling/go/pkg => /Users/fanhongling/go/pkg
    default: /Users/fanhongling/Downloads/workspace => /Users/fanhongling/Downloads/workspace
    default: /Users/fanhongling/Downloads/99-mirror => /Users/fanhongling/Downloads/99-mirror
    default: /Users/fanhongling/Downloads/go-openshift/pkg => /Users/fanhongling/Downloads/go-openshift/pkg
    default: /Users/fanhongling/Downloads/go-openshift/src => /Users/fanhongling/Downloads/go-openshift/src
    default: /Users/fanhongling/Downloads/go-kubernetes/src => /Users/fanhongling/Downloads/go-kubernetes/src
    default: /Users/fanhongling/Downloads/go-kubernetes/pkg => /Users/fanhongling/Downloads/go-kubernetes/pkg
==> default: Machine already provisioned. Run `vagrant provision` or use the `--provision`
==> default: flag to force provisioning. Provisioners marked to run always will still run.
fanhonglingdeMacBook-Pro:trusty64 fanhongling$ ssh vagrant@172.17.4.200
vagrant@172.17.4.200's password: 
Welcome to Ubuntu 14.04.5 LTS (GNU/Linux 3.13.0-76-generic x86_64)

 * Documentation:  https://help.ubuntu.com/

  System information as of Wed Jun  7 12:40:07 UTC 2017

  System load:  0.55               Users logged in:        0
  Usage of /:   93.0% of 39.34GB   IP address for eth0:    10.0.2.15
  Memory usage: 3%                 IP address for eth1:    172.17.4.200
  Swap usage:   0%                 IP address for docker0: 172.18.0.1
  Processes:    120

  => / is using 93.0% of 39.34GB

  Graph this data and manage this system at:
    https://landscape.canonical.com/

  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

51 packages can be updated.
31 updates are security updates.

New release '16.04.2 LTS' available.
Run 'do-release-upgrade' to upgrade to it.


Last login: Sun Jun  4 06:27:37 2017 from 172.17.4.1
vagrant@vagrant-ubuntu-trusty-64:~$ cd /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/
```

Protocol buffer
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb$ make openstack/neutron
protoc -I/usr/local/include -I. \
		-I/Users/fanhongling/Downloads/workspace/src \
		-I/Users/fanhongling/Downloads/workspace/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		openstack/neutron/neutron.proto
go generate .
```

For gRPC and gRPC-gateway
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb$ make all
protoc -I/usr/local/include -I. \
		-I/Users/fanhongling/Downloads/go-kubernetes:/Users/fanhongling/Downloads/go-openshift:/Users/fanhongling/Downloads/workspace/src \
		-I/Users/fanhongling/Downloads/go-kubernetes:/Users/fanhongling/Downloads/go-openshift:/Users/fanhongling/Downloads/workspace/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		service.proto data.proto
protoc -I/usr/local/include -I. \
		-I/Users/fanhongling/Downloads/go-kubernetes:/Users/fanhongling/Downloads/go-openshift:/Users/fanhongling/Downloads/workspace/src \
		-I/Users/fanhongling/Downloads/go-kubernetes:/Users/fanhongling/Downloads/go-openshift:/Users/fanhongling/Downloads/workspace/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		service.proto
protoc -I/usr/local/include -I. \
		-I/Users/fanhongling/Downloads/go-kubernetes:/Users/fanhongling/Downloads/go-openshift:/Users/fanhongling/Downloads/workspace/src \
		-I/Users/fanhongling/Downloads/go-kubernetes:/Users/fanhongling/Downloads/go-openshift:/Users/fanhongling/Downloads/workspace/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		service.proto
go generate .
```

### Build and Run

Build
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ go build -o bin/kopos -v ./kopos/
golang.org/x/net/context
github.com/grpc-ecosystem/grpc-gateway/utilities
github.com/golang/protobuf/proto
google.golang.org/grpc/codes
google.golang.org/grpc/grpclog
google.golang.org/grpc/metadata
github.com/spf13/pflag
golang.org/x/sys/unix
github.com/magiconair/properties
gopkg.in/yaml.v2
github.com/fsnotify/fsnotify
golang.org/x/net/http2/hpack
github.com/spf13/cobra
golang.org/x/net/idna
golang.org/x/net/internal/timeseries
golang.org/x/net/lex/httplex
golang.org/x/net/http2
golang.org/x/net/trace
github.com/golang/protobuf/ptypes/struct
github.com/golang/protobuf/jsonpb
github.com/grpc-ecosystem/grpc-gateway/runtime/internal
github.com/golang/protobuf/ptypes/any
google.golang.org/genproto/googleapis/rpc/status
google.golang.org/grpc/status
github.com/spf13/viper
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack/neutron
github.com/golang/protobuf/protoc-gen-go/descriptor
github.com/grpc-ecosystem/grpc-gateway/runtime
google.golang.org/genproto/googleapis/api/annotations
google.golang.org/grpc/credentials
google.golang.org/grpc/grpclb/grpc_lb_v1
google.golang.org/grpc/internal
google.golang.org/grpc/keepalive
google.golang.org/grpc/naming
google.golang.org/grpc/peer
google.golang.org/grpc/stats
google.golang.org/grpc/tap
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/insecure
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/pkg/ui/data/swagger
google.golang.org/grpc/transport
google.golang.org/grpc
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/cmd
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos
```

Server
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ bin/kopos serve
2017/06/04 03:36:19 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: dial tcp [::1]:10000: getsockopt: connection refused"; Reconnecting to {localhost:10000 <nil>}
grpc on port: 10000
```

Client test
```
fanhonglingdeMacBook-Pro:go-to-openstack-bootcamp fanhongling$ bin/kopos test create
test create
0xc420176af0
```

OpenStack DevOps
----------------

### OpenStack Pike

General [doc](./docs)
* ./docs/openstack-pike-install-1st-controller.md
* ./docs/openstack-pike-install-2nd-identity-keystone.md
* ./docs/openstack-pike-install-3rd-image-glance.md
* ./docs/openstack-pike-install-4th-controller-nova.md
* ./docs/openstack-pike-install-5th-hypervisor-nova.md

Provider networks
* ./docs/openstack-pike-install-6th-controller-provider-networks.md
* ./docs/openstack-pike-install-7th-hypervisor-provider-networks.md
* ./docs/openstack-pike-install-8th-dashboard.md


### OpenStack Grizzly版在Ubuntu Precise (12.04.3 LTS)的安装实验

参考[mseknibilel](https://github.com/mseknibilel/OpenStack-Grizzly-Install-Guide/tree/OVS_SingleNode)的安装指南，向作者致谢

#### 准备

1. 切换到超级用户root
su是linux命令，Switch User的意思，但当前用户没有执行该命令的权限
sudo命令即Super User Do，使当前用户可以执行root命令
    sudo su或sudo -i

2. 添加Grizzly的apt-get在线安装仓库
apt-get的在线安装仓库配置文件在/etc/apt-get目录下
source.list配置文件是Ubuntu系统的在线更新或升级配置文件
使用国内镜像mirrors.163.com
而第三方的在线更新配置文件可以放置在/etc/apt/source.list.d/目录下
这里将OpenStack的配置文件取名为grizzly.list
可以用vi或vim来创建和编辑
cd /etc/apt/source.list.d/
vi grizzly.list
按insert键
输入deb http://ubuntu-cloud.archive.canonical.com/ubuntu precise-updates/grizzly main
按escape键
按wq键
mseknibilel的文章中采用了重定向stdout到grizzly.list文件的写法：
echo deb http://ubuntu-cloud.archive.canonical.com/ubuntu precise-updates/grizzly main >> /etc/apt/sources.list.d/grizzly.list

3. 添加在Ubuntu Precise上安装OpenStack必要的keyring和python等工具
    apt-get install ubuntu-cloud-keyring python-software-properties software-properties-common python-keyring

4. 更新Ubuntu
    apt-get update
    apt-get upgrade
    apt-get dist-upgrade

Appendix
---------

[Markdown](https://daringfireball.net/projects/markdown/)

[Sphinx - Python Documentation Generator](http://sphinx-doc.org/contents.html)
