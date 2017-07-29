/*
   Inspired from
    https://github.com/helm/helm-classic/blob/master/kubectl/kubectl.go
*/
// package kubectl
package ovs

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

// Path is the path of the kubectl binary
var Path_vsctl = "ovs-vsctl"
var Path_ofctl = "ovs-ofctl"

// Runner is an interface to wrap kubectl convenience methods
type Runner_vsctl interface {
	//	// ClusterInfo returns Kubernetes cluster info
	//	ClusterInfo() ([]byte, error)
	//	// Create uploads a chart to Kubernetes
	//	Create([]byte, string) ([]byte, error)
	//	// Delete removes a chart from Kubernetes.
	//	Delete(string, string, string) ([]byte, error)
	//	// Get returns Kubernetes resources
	//	Get([]byte, string) ([]byte, error)
	Show() ([]byte, error)
	ListBR() ([]byte, error)
	ListPorts(br string) ([]byte, error)
}

type Runner_ofctl interface {
	ShowBR(br string) ([]byte, error)
}

// RealRunner implements Runner to execute kubectl commands
type RealRunner_vsctl struct{}
type RealRunner_ofctl struct{}

// PrintRunner implements Runner to return a []byte of the command to be executed
type PrintRunner_vsctl struct{}
type PrintRunner_ofctl struct{}

// Client stores the instance of Runner
var Client_vsctl Runner_vsctl = RealRunner_vsctl{}
var Client_ofctl Runner_ofctl = RealRunner_ofctl{}

func commandToString(cmd *exec.Cmd) string {
	var stdin string

	if cmd.Stdin != nil {
		b, _ := ioutil.ReadAll(cmd.Stdin)
		stdin = fmt.Sprintf("< %s", string(b))
	}

	return fmt.Sprintf("[CMD] %s %s", strings.Join(cmd.Args, " "), stdin)
}
