package gopacketctl

import (
	"fmt"
	"testing"
)

func TestPcap_dumpcommand(t *testing.T) {
	fmt.Println("go to pacapdump")
	err := Pcapdump("docker0")
	if err != nil {
		t.Fail()
	}
	t.Log("terminate dump")
}
