package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/digitalocean/go-qemu/hypervisor"
)

func main() {
	driver := hypervisor.NewRPCDriver(func() (net.Conn, error) {
		return net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", 2*time.Second)
	})

	hv := hypervisor.New(driver)

	fmt.Println("Domain\t\tQEMU Version")
	fmt.Println("--------------------------------------")
	domains, err := hv.Domains()
	if err != nil {
		log.Fatal(err)
	}

	for _, dom := range domains {
		version, err := dom.Version()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\t\t%s\n", dom.Name, version)
		dom.Close()
	}
}
