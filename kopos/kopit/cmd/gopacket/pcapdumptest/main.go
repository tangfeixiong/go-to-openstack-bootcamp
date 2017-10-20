package main

import (
	"fmt"

	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/gopacketctl"
)

func main() {
	fmt.Println("go to pacapdump")
	err := gopacketctl.Pcapdump("docker0")
	if err != nil {
		panic(err)
	}
	println("terminate dump")
}
