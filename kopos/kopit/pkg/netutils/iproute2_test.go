package netutils

import (
	"fmt"
	"testing"
)

func TestIPAddrList(t *testing.T) {
	fmt.Println("Go to show all IP address")

	result, err := Execute_ip_addr_show()

	if err != nil {
		t.Fail()
	}

	for _, item := range result {
		fmt.Println(item)
	}

	t.Log(result)
}
