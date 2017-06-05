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
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb$ make
protoc -I/usr/local/include -I. \
		-I/Users/fanhongling/Downloads/workspace/src \
		-I/Users/fanhongling/Downloads/workspace/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		service.proto data.proto
protoc -I/usr/local/include -I. \
		-I/Users/fanhongling/Downloads/workspace/src \
		-I/Users/fanhongling/Downloads/workspace/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:. \
		service.proto
protoc -I/usr/local/include -I. \
		-I/Users/fanhongling/Downloads/workspace/src \
		-I/Users/fanhongling/Downloads/workspace/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
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

[Markdown](https://daringfireball.net/projects/markdown/)

[Sphinx - Python Documentation Generator](http://sphinx-doc.org/contents.html)

### OpenStack-Grizzly版在Ubuntu-Precise即12.04.3-LTS的安装实验-

参考[mseknibilel](https://github.com/mseknibilel/OpenStack-Grizzly-Install-Guide/tree/OVS_SingleNode)的安装指南，向作者致谢

#### 准备

1. 切换到超级用户root
    sudo su或sudo -i
su是linux命令，Switch User的意思，但当前用户没有执行该命令的权限
sudo命令即Super User Do，使当前用户可以执行root命令

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
