#! /bin/bash

cd /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos

# gogo/protobuf
for i in $(ls /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/gogo/protobuf/); do rsync -avz /Users/fanhongling/go/src/github.com/gogo/protobuf/${i} /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/gogo/protobuf/; done

# golang/glog
for i in $(ls /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/golang/glog/); do rsync -avz /Users/fanhongling/go/src/github.com/golang/glog/${i} /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/golang/glog/; done

# golang/protobuf
rm -rf /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/golang/protobuf; for i in $(ls /Users/fanhongling/go/src/github.com/golang/protobuf/); do rsync -avz /Users/fanhongling/go/src/github.com/golang/protobuf/${i} /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/golang/protobuf/; done

# gophercloud/gophercloud
rm -rf /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/gophercloud/gophercloud; for i in $(ls /Users/fanhongling/go/src/github.com/gophercloud/gophercloud/); do rsync -avz /Users/fanhongling/go/src/github.com/gophercloud/gophercloud/${i} /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/gophercloud/gophercloud/; done

# grpc-ecosystem/grpc-gateway
rm -rf /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/grpc-ecosystem/grpc-gateway; for i in $(ls /Users/fanhongling/go/src/github.com/grpc-ecosystem/grpc-gateway/); do rsync -avz /Users/fanhongling/go/src/github.com/grpc-ecosystem/grpc-gateway/${i} /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/github.com/grpc-ecosystem/grpc-gateway/; done

# golang.org/x/net
for i in $(ls /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/golang.org/x/net/); do rsync -avz /Users/fanhongling/go/src/golang.org/x/net/${i} /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/golang.org/x/net/; done

# google.golang.org/grpc
for i in $(ls /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/google.golang.org/grpc/); do rm -rf /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/google.golang.org/grpc/${i}; rsync -avz /Users/fanhongling/go/src/google.golang.org/grpc/${i} /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/vendor/google.golang.org/grpc/; done
