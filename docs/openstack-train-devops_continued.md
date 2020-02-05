# Continued DevOps

## Configure Linux router in Ubuntu 18.04

VM is required to access internet. But the physical machine Lenovo Thinkpad x1 only hasing actived wifi to work as OpenStack mgmt net, 
and its another ethernet interface currently is n/a

Refer to:
- https://askubuntu.com/questions/1050816/ubuntu-18-04-as-a-router

Ops
```
root@wei-ThinkPad-X1-Extreme:~# ufw enable
Firewall is active and enabled on system startup
```

```
root@wei-ThinkPad-X1-Extreme:~# ufw logging on
Logging enabled
```

```
root@wei-ThinkPad-X1-Extreme:~# cat /etc/default/ufw | grep DEFAULT_FORWARD_POLICY
DEFAULT_FORWARD_POLICY="DROP"
```

```
root@wei-ThinkPad-X1-Extreme:~# sed -i "s/DEFAULT_FORWARD_POLICY=\"DROP\"/# &\nDEFAULT_FORWARD_POLICY=\"ACCEPT\"/" /etc/default/ufw 
```

```
root@wei-ThinkPad-X1-Extreme:~# cat /etc/default/ufw | grep DEFAULT_FORWARD_POLICY
# DEFAULT_FORWARD_POLICY="DROP"
DEFAULT_FORWARD_POLICY="ACCEPT"
```

```
root@wei-ThinkPad-X1-Extreme:~# cat /etc/sysctl.conf | grep "net.ipv4.ip_forward"
#net.ipv4.ip_forward=1
```

```
root@wei-ThinkPad-X1-Extreme:~# cat /etc/sysctl.conf | grep "net.ipv4.conf.all.forwarding"
```

```
root@wei-ThinkPad-X1-Extreme:~# cat /etc/sysctl.conf | grep "net.ipv6.conf.default.forwarding"
```

```
root@wei-ThinkPad-X1-Extreme:~# sed -i "1 i net.ipv4.ip_forward=1\nnet.ipv4.conf.all.forwarding=1\nnet.ipv6.conf.default.forwarding=1\n" /etc/sysctl.conf 
```

```
root@wei-ThinkPad-X1-Extreme:~# cat >snat.rules <<EOF
> # nat Table rules
> *nat
> :POSTROUTING ACCEPT [0:0]
> # Forward traffic from tap1 through wifi
> -A POSTROUTING -s 192.168.192.0/24 -o wlp0s20f3 -j MASQUERADE
> 
> COMMIT
> EOF
```

Copy upon file content into
```
root@wei-ThinkPad-X1-Extreme:~# vi /etc/ufw/before.rules 
```

Now restart
```
root@wei-ThinkPad-X1-Extreme:~# ufw disable && ufw enable
Firewall stopped and disabled on system startup
Firewall is active and enabled on system startup
```

```
root@wei-ThinkPad-X1-Extreme:~# iptables -t nat -L POSTROUTING -v -n
Chain POSTROUTING (policy ACCEPT 59 packets, 3731 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 MASQUERADE  all  --  *      wlp0s20f3  192.168.192.0/24     0.0.0.0/0         
```


### Test

Launch _cirros_ VM from OpenStack Dashboard

Net ns
```
wei@wei-ThinkPad-X1-Extreme:~/go-to-openstack-bootcamp/docs$ ip netns
qrouter-f46292f9-3e0d-427e-83dc-13d803537360 (id: 2)
qdhcp-13d1dd1d-71d5-4b2e-8876-45158bdfe7f8 (id: 0)
qdhcp-cb0c8be8-b36b-4bb6-935d-4024d44c38bc (id: 1)
```

The password is _gocubsgo_ from launch log from OpenStack Dashboard
```  
root@wei-ThinkPad-X1-Extreme:~# ip netns exec qrouter-f46292f9-3e0d-427e-83dc-13d803537360 ssh cirros@10.20.30.146
The authenticity of host '10.20.30.146 (10.20.30.146)' can't be established.
ECDSA key fingerprint is SHA256:emZ7n+O6y3FvNm+Gm1KCNxu08ZgqK9H+FxOS0oHG38w.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '10.20.30.146' (ECDSA) to the list of known hosts.
cirros@10.20.30.146's password: 
```

Response from physical machine
```
$ ping -c3 192.168.0.106
PING 192.168.0.106 (192.168.0.106): 56 data bytes
64 bytes from 192.168.0.106: seq=0 ttl=63 time=0.479 ms
64 bytes from 192.168.0.106: seq=1 ttl=63 time=0.435 ms
64 bytes from 192.168.0.106: seq=2 ttl=63 time=0.445 ms

--- 192.168.0.106 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 0.435/0.453/0.479 ms
```

Access internet is OK
```
$ ping -c3 114.114.114.114
PING 114.114.114.114 (114.114.114.114): 56 data bytes
64 bytes from 114.114.114.114: seq=0 ttl=78 time=21.120 ms
64 bytes from 114.114.114.114: seq=1 ttl=84 time=18.630 ms
64 bytes from 114.114.114.114: seq=2 ttl=73 time=18.824 ms

--- 114.114.114.114 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 18.630/19.524/21.120 ms
```


## ISSUE

### 1st

After launching VM, the DNS in correct

In `/etc/resolv.conf`
```
nameserver 10.20.30.2
search openstacklocal
```

Reference
+ https://docs.openstack.org/newton/networking-guide/config-dns-res.html

Let dhcp know internet dns
```
wei@wei-ThinkPad-X1-Extreme:~$ sudo sed -i "s/^\[DEFAULT\]$/&\ndnsmasq_dns_servers=114.114.114.114,8.8.8.8\n/" /etc/neutron/neutron-dhcp-agent.ini
```

Then restart _neutron-dhcp-agent_ service
```
wei@wei-ThinkPad-X1-Extreme:~$ sudo systemctl restart neuron-dhcp-agent
```

### 2nd

After launching VM, The incoming traffic is worked, but outgoing is not

Reference
+ https://www.howtoforge.com/nat_iptables
+ https://askubuntu.com/questions/493534/virtualbox-server-client-topology


Clean iptables
```
root@wei-ThinkPad-X1-Extreme:~# iptables -F
root@wei-ThinkPad-X1-Extreme:~# iptables -t nat -F
root@wei-ThinkPad-X1-Extreme:~# iptables -t mangle -F
```

Add FORWARD rule
```
root@wei-ThinkPad-X1-Extreme:~# vi /etc/ufw/before.rules 
```

```
root@wei-ThinkPad-X1-Extreme:~# cat /etc/ufw/before.rules | more


# nat Table rules
*nat
:POSTROUTING ACCEPT [0:0]
# Forward traffic from tap1 through wifi
-A POSTROUTING -s 192.168.192.0/24 -o wlp0s20f3 -j MASQUERADE

COMMIT

#
# rules.before
#
# Rules that should be run before the ufw command line added rules. Custom
# rules should be added to one of these chains:
#   ufw-before-input
#   ufw-before-output
#   ufw-before-forward
#

# Don't delete these required lines, otherwise there will be errors
*filter
:ufw-before-input - [0:0]
:ufw-before-output - [0:0]
:ufw-before-forward - [0:0]
:ufw-not-local - [0:0]
# End required lines

# --append FORWARD --in-interface eth1 -j ACCEPT
-A ufw-before-forward -i brq13d1dd1d-71 -j ACCEPT

# allow all on loopback
```

Reset all rules
```
root@wei-ThinkPad-X1-Extreme:~# ufw disable && ufw enable
```

Now look like
```
root@wei-ThinkPad-X1-Extreme:~# iptables -t nat -L POSTROUTING -vn
Chain POSTROUTING (policy ACCEPT 26 packets, 1756 bytes)
 pkts bytes target     prot opt in     out     source               destination         
    2   144 MASQUERADE  all  --  *      wlp0s20f3  192.168.192.0/24     0.0.0.0/0           

root@wei-ThinkPad-X1-Extreme:~# iptables -t filter -L ufw-before-forward -vn
Chain ufw-before-forward (0 references)
 pkts bytes target     prot opt in     out     source               destination         
    0     0 ACCEPT     all  --  brq13d1dd1d-71 *       0.0.0.0/0            0.0.0.0/0           
    0     0 ACCEPT     all  --  *      *       0.0.0.0/0            0.0.0.0/0            ctstate RELATED,ESTABLISHED
    0     0 ACCEPT     icmp --  *      *       0.0.0.0/0            0.0.0.0/0            icmptype 3
    0     0 ACCEPT     icmp --  *      *       0.0.0.0/0            0.0.0.0/0            icmptype 11
    0     0 ACCEPT     icmp --  *      *       0.0.0.0/0            0.0.0.0/0            icmptype 12
    0     0 ACCEPT     icmp --  *      *       0.0.0.0/0            0.0.0.0/0            icmptype 8
    0     0 ufw-user-forward  all  --  *      *       0.0.0.0/0            0.0.0.0/0           

```

### 3rd

All-in-one OpenStack resize VM failed

```
wei@wei-ThinkPad-X1-Extreme:~$ openstack server resize --flavor=m1.medium.disk15gb masterk8s
```

Reference
+ https://computingforgeeks.com/how-to-resize-openstack-instance-virtual-machine/

```
wei@wei-ThinkPad-X1-Extreme:~$ sudo sed -i "s/^\[DEFAULT\]$/&\nallow_resize_to_same_host=true\n/" /etc/nova/nova.conf
```

```
wei@wei-ThinkPad-X1-Extreme:~$ for svc in compute api scheduler; do sudo systemctl restart nova-${service}.service; done
```

```
wei@wei-ThinkPad-X1-Extreme:~$ openstack server resize --flavor m1.medium.disk15gb masterk8s
```

```
wei@wei-ThinkPad-X1-Extreme:~$ openstack server resize confirm masterk8s
```
