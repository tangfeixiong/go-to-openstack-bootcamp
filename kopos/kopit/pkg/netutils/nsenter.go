package netutils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util"
	"github.com/tangfeixiong/go-to-openstack-bootcamp/kopos/kopit/pkg/util/nsenter"
)

func ListAddresses_nsenter(target string) ([]*IPAddress, error) {
	util.Logger.Println("Go to show all IP address")

	cli := nsenter.Client
	out, err := cli.Run(target, false, false, false, true, false, "ip", "address", "show")
	if nil != err {
		util.Logger.Println(err)
		return nil, fmt.Errorf("Failed to list ip stacks: %s", err.Error())
	}

	data := make([]*IPAddress, 0)
	var resp *IPAddress
	var lftv4 bool
	scanner := bufio.NewScanner(bytes.NewBuffer(out))
	for i := 0; scanner.Scan(); i++ {
		if nil == scanner.Err() && 0 != len(scanner.Text()) {
			line := scanner.Text()
			switch {
			case strings.HasPrefix(line, "    inet "):
				lftv4 = true
				re := regexp.MustCompile(net_ipv4_regexp)
				result := re.FindAllStringSubmatch(line, -1)
				if nil != result {
					resp.IPv4 = result[0][1]
					resp.V4Mask = result[0][3]
					resp.V4Info = strings.Split(result[0][4], " ")
				}
			case strings.HasPrefix(line, "    inet6 "):
				lftv4 = false
				re := regexp.MustCompile(net_ipv6_regexp)
				result := re.FindAllStringSubmatch(line, -1)
				if nil != result {
					resp.IPv6 = result[0][1]
					resp.V6Mask = result[0][2]
					resp.V6Info = strings.Split(result[0][3], " ")
				}
			case strings.HasPrefix(line, "       "):
				if lftv4 {
					resp.V4Lifetime = strings.Split(strings.TrimLeft(line, " "), " ")
				} else {
					resp.V6Lifetime = strings.Split(strings.TrimLeft(line, " "), " ")
				}
			default:
				if 0 < i {
					data = append(data, resp)
				}
				resp = new(IPAddress)
				re := regexp.MustCompile(datalink_Line1_regexp)
				result := re.FindAllStringSubmatch(line, -1)
				if nil != result {
					resp.Id = result[0][1]
					resp.Name = result[0][2]
					resp.DataLinkStatus = strings.Split(result[0][3], ",")
					//				resp.DataLinkConf = make(map[string]string)
					conf := strings.Split(result[0][4], " ")
					//				for i := 0; i < len(conf)-1; i += 2 {
					//					resp.DataLinkConf[conf[i]] = conf[i+1]
					//				}
					resp.DataLinkConf = conf
					if scanner.Scan() && nil == scanner.Err() {
						line = scanner.Text()
						re = regexp.MustCompile(datalink_line2_regexp)
						result = re.FindAllStringSubmatch(line, -1)
						resp.DataLinkFrame = result[0][1]
						if "ether" == resp.DataLinkFrame || "loopback" == resp.DataLinkFrame {
							resp.DataLinkEtherMAC = result[0][2]
							resp.DataLinkEtherBRD = result[0][4]
							if len(result[0]) == 8 {
								resp.DataLinkNetnsID = result[0][7]
							}
						}
					}
				}
			}
		} else if io.EOF == scanner.Err() {
			break
		} else {
			return nil, fmt.Errorf("Failed to list ip stacks detailed: %s", scanner.Err().Error())
		}
	}
	data = append(data, resp)
	return data, nil
}
