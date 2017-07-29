# Inspired from https://github.com/philips/grpc-gateway-example

GOPATHP:=/Users/fanhongling/Downloads/workspace
GOPATHD:=/home/vagrant/go

all: openstack
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mopenstack/data.proto=github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack,Mopenstack/image.proto=github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack,Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
		service.proto workerservice.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=Mopenstack/data.proto=github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack,Mopenstack/image.proto=github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack,logtostderr=true:. \
		service.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		service.proto
	go generate .

openstack: openstack/identity openstack/nova openstack/neutron
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,Mopenstack/neutron/neutron.proto=github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack/neutron,Mopenstack/nova/nova.proto=github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack/nova,Mopenstack/identity/identity.proto=github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack/identity:. \
        openstack/data.proto openstack/image.proto
        
openstack/neutron:
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any,plugins=grpc:. \
		openstack/neutron/neutron.proto
        
openstack/nova:
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,plugins=grpc:. \
		openstack/nova/nova.proto
        
openstack/identity:
	protoc -I/usr/local/include -I. \
		-I${GOPATHP}/src \
		-I${GOPATHD}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,plugins=grpc:. \
		openstack/identity/identity.proto

.PHONY: all openstack openstack/identity openstack/nova openstack/neutron