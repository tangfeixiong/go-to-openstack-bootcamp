# Architecture of OpenStack Bootcamp

## Tables of content



## Configurability of Domain

Core Technology: *Kubernetes* [ThirdPartyResource](http://kubernetes.io/docs/user-guide/thirdpartyresources/)

See also: *CoreOS* [Introducing Operators](https://coreos.com/blog/introducing-operators.html)

Prior Art: 

* Kuberntes Openstack Operators - [sapcc/kubernetes-operators](https://github.com/sapcc/kubernetes-operators)
* CoreOS Etcd Operators - [coros/etcd-operators](https://github.com/sapcc/kubernetes-operators)
* CoreOS Prometheus Operators - [coreos/prometheus-operators](https://github.com/coreos/prometheus-operator)
* kelseyhightower kubernetes certificate manager - [kelseyhightower/kubernetes-certificate-manager](https://github.com/kelseyhightower/kube-cert-manager)

Essential Technology: [gRPC serviing multiplexly for Protobuf and Http1]

* [gRPC](https://github.com/grpc/grpc-go)
* [protobuf](https://github.com/google/protobuf)
* [orotobuf for Go](https://github.com/golang/protobuf)
* [gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)

Prior Art:

* [philips/grpc-gateway-example](https://github.com/philips/grpc-gateway-example)
* [soheilhy/cmux](https://github.com/soheilhy/cmux)

Required Technology: [Openstack SDK for Go](https://github.com/gophercloud/gophercloud)

* Identity/keystone
* Image/glance
* Networking/neutron
* Compute/nova

### Features

**Create and Inspect private network for each project**

Powering up earch new *Project(or Tenant)* to provisioning internal L2 network of privacy

**Create and Inspect OpenStack available to hava a configured shared network after starting**

Powering up each new *Project(or Tenant)* and its joining *users* to provisioning VM in one L2 networking

### Proof of Concepts

**Battlefield** is identified as a shared networking created by admin