package netutils

import (
	"fmt"
	"testing"
)

func TestLinuxBridge_show(t *testing.T) {
	out, err := Execute_brctl_show()
	if nil != err {
		t.Fail()
	}

	for _, v := range out {
		fmt.Println(v)
	}
}

func TestLinuxBridge_showmacs(t *testing.T) {
	out, err := Execute_brctl_showmacs("docker0")
	if nil != err {
		t.Fail()
	}

	for _, v := range out {
		fmt.Println(v)
	}
}
