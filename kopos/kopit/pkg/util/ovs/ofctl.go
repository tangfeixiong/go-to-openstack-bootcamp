package ovs

/*
# ovs-ofctl show br-int
OFPT_FEATURES_REPLY (xid=0x2): dpid:0000bacba8eb4f43
n_tables:254, n_buffers:256
capabilities: FLOW_STATS TABLE_STATS PORT_STATS QUEUE_STATS ARP_MATCH_IP
actions: OUTPUT SET_VLAN_VID SET_VLAN_PCP STRIP_VLAN SET_DL_SRC SET_DL_DST SET_NW_SRC SET_NW_DST SET_NW_TOS SET_TP_SRC SET_TP_DST ENQUEUE
 3(int-br-eth1): addr:76:d7:1d:3f:4e:c4
     config:     0
     state:      0
     speed: 0 Mbps now, 0 Mbps max
 4(int-br-ex2): addr:7e:d9:87:cc:ea:57
     config:     0
     state:      0
     speed: 0 Mbps now, 0 Mbps max
 LOCAL(br-int): addr:ba:cb:a8:eb:4f:43
     config:     PORT_DOWN
     state:      LINK_DOWN
     speed: 0 Mbps now, 0 Mbps max
OFPT_GET_CONFIG_REPLY (xid=0x4): frags=normal miss_send_len=0
int-br-eth1
int-br-ex2
patch-tun
*/
func (r RealRunner_ofctl) ShowBR(br string) ([]byte, error) {
	args := []string{"show", br}

	cmd := command_ofctl(args...)

	return cmd.CombinedOutput()

}
