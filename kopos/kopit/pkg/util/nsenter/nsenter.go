/*
   Inspired from
    https://github.com/helm/helm-classic/blob/master/kubectl/get.go
*/
// package kubectl
package nsenter

import "errors"

func (r RealRunner) Run(target string, mount, uts, ipc, net, pid bool, cmds ...string) ([]byte, error) {
	if 0 == len(target) {
		return []byte{}, errors.New("Target required")
	}

	args := append([]string{"--"}, cmds...)
	if mount {
		args = append([]string{"--mount"}, args...)
	}
	if uts {
		args = append([]string{"--uts"}, args...)
	}
	if ipc {
		args = append([]string{"--ipc"}, args...)
	}
	if net {
		args = append([]string{"--net"}, args...)
	}
	if pid {
		args = append([]string{"--pid"}, args...)
	}
	args = append([]string{"--target", target}, args...)

	cmd := command(args...)
	if mount || len(cmds) == 0 {
		assignStdin(cmd, []byte("exit"))
	}

	return cmd.CombinedOutput()
}

//// Get returns Kubernetes resources
//func (r RealRunner) Get(stdin []byte, ns string) ([]byte, error) {
//	args := []string{"get", "-f", "-"}

//	if ns != "" {
//		args = append([]string{"--namespace=" + ns}, args...)
//	}
//	cmd := command(args...)
//	assignStdin(cmd, stdin)

//	return cmd.CombinedOutput()
//}

//// Get returns the commands to kubectl
//func (r PrintRunner) Get(stdin []byte, ns string) ([]byte, error) {
//	args := []string{"get", "-f", "-"}

//	if ns != "" {
//		args = append([]string{"--namespace=" + ns}, args...)
//	}
//	cmd := command(args...)
//	assignStdin(cmd, stdin)

//	return []byte(cmd.String()), nil
//}
