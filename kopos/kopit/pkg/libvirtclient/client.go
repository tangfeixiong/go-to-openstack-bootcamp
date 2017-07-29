package libvirtclient

import (
	"encoding/xml"
	"fmt"

	libvirt "github.com/libvirt/libvirt-go"
	libvirtxml "github.com/libvirt/libvirt-go-xml"

	pbos "github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/echopb/openstack"
)

func Execute() (*pbos.LibvirtDomainInfo, error) {
	resp := new(pbos.LibvirtDomainInfo)

	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		return resp, err
	}

	dom, err := conn.LookupDomainByName("demo")
	if err != nil {
		return resp, err
	}
	xmldoc, err := dom.GetXMLDesc(0)
	if err != nil {
		return resp, err
	}

	domcfg := &libvirtxml.Domain{}
	if err := xml.Unmarshal([]byte(xmldoc), domcfg); err != nil {
		return resp, err
	}
	fmt.Printf("Virt type %s", domcfg.Type)

	return resp, nil
}
