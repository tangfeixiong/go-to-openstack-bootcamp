/*
   Inspired from
    https://github.com/helm/helm-classic/blob/master/kubectl/get.go
*/
// package kubectl
package ovs

/*
# ovs-vsctl show
5ee56fdf-09b6-4f38-8de2-a699670bf8f7
    Bridge "br-ex2"
        Port "phy-br-ex2"
            Interface "phy-br-ex2"
                type: patch
                options: {peer="int-br-ex2"}
        Port "br-ex2"
            Interface "br-ex2"
                type: internal
        Port "eth2"
            Interface "eth2"
    Bridge "br-eth1"
        Port "br-eth1"
            Interface "br-eth1"
                type: internal
        Port "phy-br-eth1"
            Interface "phy-br-eth1"
                type: patch
                options: {peer="int-br-eth1"}
        Port "eth1"
            Interface "eth1"
    Bridge br-int
        fail_mode: secure
        Port br-int
            Interface br-int
                type: internal
        Port "qr-476984a0-53"
            tag: 1
            Interface "qr-476984a0-53"
                type: internal
        Port "int-br-ex2"
            Interface "int-br-ex2"
                type: patch
                options: {peer="phy-br-ex2"}
        Port "int-br-eth1"
            Interface "int-br-eth1"
                type: patch
                options: {peer="phy-br-eth1"}
        Port "tapa3281257-92"
            tag: 1
            Interface "tapa3281257-92"
                type: internal
     ovs_version: "2.3.0"
*/
func (r RealRunner_vsctl) Show() ([]byte, error) {
	args := []string{"show"}

	cmd := command_vsctl(args...)

	return cmd.CombinedOutput()
}

/*
# ovs-vsctl list-br
br-eth1
br-ex2
br-int
br-tun
*/
func (r RealRunner_vsctl) ListBR() ([]byte, error) {
	args := []string{"list-br"}

	cmd := command_vsctl(args...)

	return cmd.CombinedOutput()
}

/*
# ovs-vsctl list-ports br-int
int-br-eth1
int-br-ex2
patch-tun
*/
func (r RealRunner_vsctl) ListPorts(br string) ([]byte, error) {
	args := []string{"list-ports", br}

	cmd := command_vsctl(args...)

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
