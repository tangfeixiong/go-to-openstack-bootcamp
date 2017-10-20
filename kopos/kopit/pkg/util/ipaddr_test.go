package util

import (
	"fmt"
	"testing"
)

func TestIPRoute2_AddrList(t *testing.T) {
	Logger.Println("Go to show all IP address")

	out, err := Client.AddrShow("")
	if nil != err {
		Logger.Println(err)
		t.Fail()
	}

	fmt.Println(string(out))
}
