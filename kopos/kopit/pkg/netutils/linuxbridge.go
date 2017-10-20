package netutils

import (
	"bufio"
	"bytes"
	"fmt"
	// "strconv"
	// "time"
	// "io"
	// "regexp"
	// "strconv"
	"strings"

	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util"
	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util/linuxbridge"
)

type LinuxBridgeInfo struct {
	Id         string
	Name       string
	STPEnabled string
	Interfaces []string
	MacInfo    []LinuxBridgeLearnedMac
}

type LinuxBridgeLearnedMac struct {
	PortNo      string
	MacAddr     string
	IsLocal     string
	AgeingTimer string
}

/*
[vagrant@localhost go-to-docker]$ brctl show
bridge name	bridge id		STP enabled	interfaces
br-3ba1465afaf1		8000.024214f071ac	no
br-ec9fc4fb4445		8000.0242c5ec27dc	no		veth737a0bd
							vethb7c3893
							vethfc520c8
docker0		8000.02426cc5b36c	no		veth0d22594
							veth209da7c
							veth56d65fa
							veth9985e1b
							vethb43280b
							vethc079921
							vethfba224c
*/
func Execute_brctl_show() ([]*LinuxBridgeInfo, error) {
	util.Logger.Println("Go to show all linux bridge")

	cli := linuxbridge.Client
	out, err := cli.Show("")
	if nil != err {
		util.Logger.Println(err)
		return nil, fmt.Errorf("Failed to list linux bridge: %s", err.Error())
	}

	data := make([]*LinuxBridgeInfo, 0)
	var br *LinuxBridgeInfo
	scanner := bufio.NewScanner(bytes.NewBuffer(out))
	for i := 0; scanner.Scan(); i++ {
		if nil == scanner.Err() && 0 != len(scanner.Text()) {
			line := scanner.Text()
			switch {
			default:
				info := strings.Split(line, "\t")
				fmt.Printf("%q\n", info)
				br = &LinuxBridgeInfo{
					Id:         info[2],
					Name:       info[0],
					STPEnabled: info[3],
					Interfaces: make([]string, 0),
				}
				if 5 < len(info) {
					br.Interfaces = append(br.Interfaces, info[5])
				}
				data = append(data, br)
			case strings.Contains(line, "bridge name") && strings.Contains(line, "bridge id"):
				continue
			case strings.HasPrefix(line, "\t") || strings.HasPrefix(line, "    "):
				iface := strings.TrimSpace(line)
				br.Interfaces = append(br.Interfaces, iface)
			}
		}
	}
	return data, nil
}

/*
[vagrant@localhost go-to-docker]$ brctl showmacs docker0
port no	mac addr		is local?	ageing timer
  3	02:42:ac:11:00:04	no		 124.37
  6	0a:69:70:9b:0f:99	yes		   0.00
  6	0a:69:70:9b:0f:99	yes		   0.00
  7	26:2c:07:d4:d2:e5	yes		   0.00
  7	26:2c:07:d4:d2:e5	yes		   0.00
  5	3e:6b:12:98:bb:94	yes		   0.00
  5	3e:6b:12:98:bb:94	yes		   0.00
  2	5e:db:ba:23:b1:a5	yes		   0.00
  2	5e:db:ba:23:b1:a5	yes		   0.00
  1	72:ef:8c:7b:b3:ab	yes		   0.00
  1	72:ef:8c:7b:b3:ab	yes		   0.00
  4	82:df:75:b5:e7:3f	yes		   0.00
  4	82:df:75:b5:e7:3f	yes		   0.00
  4	96:ab:2b:58:49:27	no		 223.99
  3	a2:69:5c:18:1f:2c	yes		   0.00
  3	a2:69:5c:18:1f:2c	yes		   0.00
*/
func Execute_brctl_showmacs(bridge string) ([]*LinuxBridgeLearnedMac, error) {
	if bridge == "" {
		return []*LinuxBridgeLearnedMac{}, fmt.Errorf("bridge required")
	}
	util.Logger.Println("Go to show all learned MACs from linux bridge", bridge)

	cli := linuxbridge.Client
	out, err := cli.ShowMACs(bridge)
	if nil != err {
		util.Logger.Println(err)
		return nil, fmt.Errorf("Failed to list MAC addr in linux bridge %s: %s", bridge, err.Error())
	}

	data := make([]*LinuxBridgeLearnedMac, 0)
	var mac *LinuxBridgeLearnedMac
	scanner := bufio.NewScanner(bytes.NewBuffer(out))
	for i := 0; scanner.Scan(); i++ {
		if nil == scanner.Err() && 0 != len(scanner.Text()) {
			line := scanner.Text()
			switch {
			default:
				info := strings.Split(line, "\t")
				fmt.Printf("%q\n", info)
				mac = &LinuxBridgeLearnedMac{
					PortNo:      strings.Trim(info[0], " "),
					MacAddr:     info[1],
					IsLocal:     info[2],
					AgeingTimer: strings.Trim(info[4], " "),
				}
				data = append(data, mac)
			case strings.Contains(line, "port no") && strings.Contains(line, "mac addr"):
				continue
			}
		}
	}
	return data, nil
}
