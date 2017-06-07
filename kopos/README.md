### Build Docker Image

OpenAPI

![屏幕快照 2017-06-07 上午7.15.28.png](../docs/屏幕快照%202017-06-07%20上午7.15.28.png)

Canary Docker Image
![屏幕快照 2017-06-07 上午7.31.25.png](../docs/屏幕快照%202017-06-07%20上午7.31.25.png)

Bin
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp$ CGO_ENABLED=0 go build -v -o kopos/bin/kopos --installsuffix cgo ./kopos/
runtime/internal/sys
runtime/internal/atomic
runtime
errors
internal/race
sync/atomic
unicode
sync
unicode/utf8
io
container/list
bytes
hash
math
crypto/subtle
crypto/cipher
crypto/internal/cipherhw
syscall
strconv
time
crypto
crypto/aes
reflect
os
math/rand
strings
crypto/sha512
encoding/binary
fmt
crypto/des
crypto/hmac
crypto/md5
math/big
bufio
internal/syscall/unix
crypto/rc4
crypto/sha1
crypto/sha256
encoding/hex
crypto/elliptic
encoding/asn1
crypto/rand
crypto/ecdsa
crypto/rsa
crypto/dsa
crypto/x509/pkix
encoding/base64
sort
context
encoding/pem
path/filepath
internal/nettrace
internal/singleflight
net
io/ioutil
vendor/golang_org/x/crypto/chacha20poly1305/internal/chacha20
vendor/golang_org/x/crypto/poly1305
vendor/golang_org/x/crypto/chacha20poly1305
vendor/golang_org/x/crypto/curve25519
encoding
unicode/utf16
encoding/json
crypto/x509
log
github.com/golang/protobuf/proto
crypto/tls
github.com/grpc-ecosystem/grpc-gateway/utilities
github.com/golang/protobuf/ptypes/struct
github.com/grpc-ecosystem/grpc-gateway/runtime/internal
github.com/golang/protobuf/jsonpb
golang.org/x/net/context
google.golang.org/grpc/codes
google.golang.org/grpc/grpclog
google.golang.org/grpc/metadata
github.com/golang/protobuf/ptypes/any
compress/flate
google.golang.org/genproto/googleapis/rpc/status
google.golang.org/grpc/status
hash/crc32
vendor/golang_org/x/net/http2/hpack
compress/gzip
vendor/golang_org/x/net/idna
vendor/golang_org/x/text/transform
vendor/golang_org/x/net/lex/httplex
mime
vendor/golang_org/x/text/unicode/norm
vendor/golang_org/x/text/width
mime/quotedprintable
net/textproto
net/http/httptrace
net/http/internal
mime/multipart
net/url
path
encoding/csv
net/http
flag
github.com/spf13/pflag
text/template/parse
text/template
github.com/spf13/cobra
golang.org/x/sys/unix
github.com/grpc-ecosystem/grpc-gateway/runtime
github.com/philips/go-bindata-assetfs
github.com/fsnotify/fsnotify
github.com/hashicorp/hcl/hcl/strconv
github.com/hashicorp/hcl/hcl/token
regexp/syntax
github.com/hashicorp/hcl/hcl/ast
github.com/hashicorp/hcl/hcl/scanner
github.com/hashicorp/hcl/hcl/parser
github.com/hashicorp/hcl/json/token
github.com/hashicorp/hcl/json/scanner
regexp
github.com/hashicorp/hcl/json/parser
github.com/hashicorp/hcl
github.com/magiconair/properties
github.com/mitchellh/mapstructure
github.com/pelletier/go-buffruneio
github.com/spf13/afero/mem
github.com/pelletier/go-toml
golang.org/x/text/transform
golang.org/x/text/unicode/norm
html
html/template
github.com/spf13/afero
github.com/spf13/cast
github.com/spf13/jwalterweatherman
gopkg.in/yaml.v2
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack/neutron
github.com/golang/protobuf/protoc-gen-go/descriptor
google.golang.org/genproto/googleapis/api/annotations
golang.org/x/net/http2/hpack
github.com/spf13/viper
golang.org/x/text/unicode/bidi
golang.org/x/text/secure/bidirule
golang.org/x/net/internal/timeseries
golang.org/x/net/idna
text/tabwriter
golang.org/x/net/trace
golang.org/x/net/lex/httplex
golang.org/x/net/http2
google.golang.org/grpc/credentials
google.golang.org/grpc/grpclb/grpc_lb_v1
google.golang.org/grpc/internal
google.golang.org/grpc/keepalive
google.golang.org/grpc/naming
google.golang.org/grpc/peer
google.golang.org/grpc/stats
google.golang.org/grpc/tap
net/http/httputil
google.golang.org/grpc/transport
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/insecure
os/user
github.com/golang/glog
github.com/gophercloud/gophercloud
google.golang.org/grpc
github.com/gophercloud/gophercloud/pagination
github.com/gophercloud/gophercloud/openstack/identity/v2/tenants
github.com/gophercloud/gophercloud/openstack/identity/v2/tokens
github.com/gophercloud/gophercloud/openstack/identity/v3/tokens
github.com/gophercloud/gophercloud/openstack/utils
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb
github.com/gophercloud/gophercloud/openstack
github.com/gophercloud/gophercloud/openstack/compute/v2/flavors
github.com/gophercloud/gophercloud/openstack/compute/v2/images
github.com/gophercloud/gophercloud/openstack/identity/v2/users
github.com/gophercloud/gophercloud/openstack/compute/v2/servers
github.com/gophercloud/gophercloud/openstack/identity/v3/projects
github.com/gophercloud/gophercloud/openstack/imageservice/v2/images
github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers
github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules
github.com/gophercloud/gophercloud/openstack/networking/v2/networks
github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups
github.com/gophercloud/gophercloud/openstack/networking/v2/ports
github.com/gophercloud/gophercloud/openstack/networking/v2/subnets
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/pkg/ui/data/swagger
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/pkg/osctl
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/cmd
github.com/tangfeixiong/go-to-openstack-bootcamp/kopos
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp$ kopos/bin/kopos --helpTo get started run the serve subcommand which will start a server
on localhost:10000:

    grpc-gateway-example serve

Then you can hit it with the client:

    grpc-gateway-example echo foo bar baz

Or over HTTP 1.1 with curl:

    curl -X POST -k https://localhost:10000/v1/echo -d '{"value": "foo"}'

Usage:
  kopos [command]

Available Commands:
  echo        Example echo gRPC service CLI client
  help        Help about any command
  serve       Launches the example webserver on https://localhost:10000
  test        Example echo gRPC service CLI client

Flags:
      --config string   config file (default is $HOME/.grpc-gateway-example.yaml)
  -h, --help            help for kopos
  -t, --toggle          Help message for toggle

Use "kopos [command] --help" for more information about a command.
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp$ kopos/bin/kopos serve
grpc on port: 10000
2017/06/07 14:08:45 grpc: addrConn.resetTransport failed to create client transport: connection error: desc = "transport: Error while dialing dial tcp 127.0.0.1:10000: getsockopt: connection refused"; Reconnecting to {localhost:10000 <nil>}
^C
```

Image
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp$ cd kopos/
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ docker build -t tangfeixiong/go-to-openstack-bootcamp:canary .
Sending build context to Docker daemon 17.11 MB
Step 1 : FROM busybox
 ---> c75bebcdd211
Step 2 : LABEL "maintainer" "tangfeixiong <tangfx128@gmail.com>"
 ---> Running in 2c8130c9ced3
 ---> ef908ee55f27
Removing intermediate container 2c8130c9ced3
Step 3 : COPY bin/kopos /
 ---> f8ce2a123e24
Removing intermediate container 77a9489bc3fa
Step 4 : CMD /kopos serve
 ---> Running in d7d69148adde
 ---> 10c469030d31
Removing intermediate container d7d69148adde
Successfully built 10c469030d31
```

Run
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ docker run -d name=bootcamp -p 10000:10000 tangfeixiong/go-to-openstack-bootcamp:canary
5f8b5535a4272e5d93627f34f7b4f61e239d2b96f5aa81280ecbd70d5a046b67
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ curl -X POST -k https://localhost:10000/v1/echo -H "Content-Type: text/plain" -d '{"value": "foo"}'
{"value":"foo"}
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ curl -X POST -k https://172.17.4.200:10000/v1/echo -H "Content-Type: text/plain" -d '{"value": "bar"}'
{"value":"bar"}
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ bin/kopos echo "foo bar baz"
foo bar baz
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ bin/kopos test
test
0xc42004e0f0
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ docker logs 5f8b5535a4272e5d93627f34f7b4f61e239d2b96f5aa81280ecbd70d5a046b67
grpc on port: 10000
rpc request Echo("foo")
rpc request Echo("foo")
rpc request Echo("bar")
rpc request Echo("foo bar baz")
ERROR: logging before flag.Parse: E0607 14:20:42.145876       1 admin.go:110] Could not load admin openrc: Missing input for argument [authURL]
rpc AdminSharedNetworkCreation(name:"test" )
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ docker stop 5f8b5535a4272e5d93627f34f7b4f61e239d2b96f5aa81280ecbd70d5a046b67
5f8b5535a4272e5d93627f34f7b4f61e239d2b96f5aa81280ecbd70d5a046b67
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ docker rm 5f8b5535a4272e5d93627f34f7b4f61e239d2b96f5aa81280ecbd70d5a046b67
5f8b5535a4272e5d93627f34f7b4f61e239d2b96f5aa81280ecbd70d5a046b67
```

Ship
```
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ docker login 
Username: tangfeixiong
Password: 
Email: tangfx128@gmail.com
WARNING: login credentials saved in /home/vagrant/.docker/config.json
Login Succeeded
vagrant@vagrant-ubuntu-trusty-64:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos$ docker push tangfeixiong/go-to-openstack-bootcamp:canary
The push refers to a repository [docker.io/tangfeixiong/go-to-openstack-bootcamp]
6a5537dbb4eb: Pushed 
4ac76077f2c7: Pushed 
canary: digest: sha256:de41ced4baccdfa61a903d253f7cb0e312dba8f8a1ef9d6bb80a51aa30fa168b size: 3769
```
