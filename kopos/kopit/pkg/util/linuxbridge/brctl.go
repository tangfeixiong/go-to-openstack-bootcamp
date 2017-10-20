/*
   Inspired from
    https://github.com/helm/helm-classic/blob/master/kubectl/get.go
*/
// package kubectl
package linuxbridge

import "errors"

func (r RealRunner) Show(bridge string) ([]byte, error) {
	args := []string{"show"}

	if 0 != len(bridge) {
		args = append(args, bridge)
	}
	cmd := command(args...)

	return cmd.CombinedOutput()
}

func (r RealRunner) ShowMACs(bridge string) ([]byte, error) {
	if 0 == len(bridge) {
		return []byte{}, errors.New("bridge required")
	}

	args := []string{"showmacs", bridge}
	cmd := command(args...)

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
