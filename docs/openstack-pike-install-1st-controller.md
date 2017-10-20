# OpenStack Pike Installation

## Table of content

控制节点
* [网络配置](#network)
* [SSH客户端排错](#ssh-trouble-shooting)
* [防火墙](#firewall)
* [NTP服务](#chrony)
* [Openstack Pike版本YUM仓库](#openstack-Repository)
* [Openstack 命令工具包](#openstack-client)
* [MariaDB数据库](#database)
* [RabbitMQ消息队列](#queue)
* [Memcached缓存](#cache)

## Controller Node

### Network

Before
```
[vagrant@localhost ~]$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: eth0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 52:54:00:7a:69:5a brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global dynamic eth0
       valid_lft 86137sec preferred_lft 86137sec
    inet6 fe80::5054:ff:fe7a:695a/64 scope link
       valid_lft forever preferred_lft forever
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 08:00:27:68:ae:d9 brd ff:ff:ff:ff:ff:ff
```

```
[vagrant@localhost ~]$ ls /etc/sysconfig/network-scripts/ifcfg*
/etc/sysconfig/network-scripts/ifcfg-eth0
/etc/sysconfig/network-scripts/ifcfg-lo
```

Note, it is VirtualBox, the eth0 must be preserved by Vagrant

Using `nmcli`
```
[vagrant@localhost ~]$ nmcli c
NAME                UUID                                  TYPE            DEVICE
System eth0         5fb06bd0-0bb0-7ffb-45f1-d6edd65f3e03  802-3-ethernet  eth0  
Wired connection 1  3c55af49-6222-3b6b-b91d-eb1b82b6005e  802-3-ethernet  eth1  
```
```
[vagrant@localhost ~]$ nmcli d
DEVICE  TYPE      STATE                                  CONNECTION
eth0    ethernet  connected                              System eth0
eth1    ethernet  connecting (getting IP configuration)  Wired connection 1
lo      loopback  unmanaged                              --
```

```
[vagrant@localhost ~]$ sudo nmcli con mod "Wired connection 1" ipv4.method manual ipv4.addr "10.64.33.64/24" ipv4.dns 8.8.4.4
```

```
[vagrant@localhost ~]$ cat /etc/sysconfig/network-scripts/ifcfg-Wired_connection_1
HWADDR=08:00:27:68:AE:D9
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=none
IPADDR=10.64.33.64
PREFIX=24
DNS1=8.8.4.4
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=stable-privacy
NAME="Wired connection 1"
UUID=3c55af49-6222-3b6b-b91d-eb1b82b6005e
ONBOOT=yes
AUTOCONNECT_PRIORITY=-999
```

```
[vagrant@localhost ~]$ ip a show eth1
3: eth1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP qlen 1000
    link/ether 08:00:27:68:ae:d9 brd ff:ff:ff:ff:ff:ff
    inet 10.64.33.64/24 brd 10.64.33.255 scope global eth1
       valid_lft forever preferred_lft forever
    inet6 fe80::a1f1:2c18:bc74:7a3d/64 scope link
       valid_lft forever preferred_lft forever
```

```
[vagrant@localhost ~]$ sudo nmcli con reload
```

### SSH trouble shooting

trouble shooting
```
$ ssh vagrant@10.64.33.64
The authenticity of host '10.64.33.64 (10.64.33.64)' can't be established.
ECDSA key fingerprint is SHA256:rmUJwM3oz/L+uushIeJGRatrsRIxxxh79M5JMA+Ljzc.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '10.64.33.64' (ECDSA) to the list of known hosts.
vagrant@10.64.33.64: Permission denied (publickey,gssapi-keyex,gssapi-with-mic).
```

```
[vagrant@localhost ~]$ sudo cat /etc/ssh/sshd_config  | egrep "^#?Password"
#PasswordAuthentication yes
PasswordAuthentication no
```

```
tangf@DESKTOP-H68OQDV ~
$ cat ~/.ssh/id_rsa.pub
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCvfIdGduA81WVgf1F5DikDG+1qJEPk0FBYtMPk7WTEkb4p8KkqMKKdrt7Sy7Ig4ZIBwFCCU4rtHiaLeslNxwdjT1l1sH18uiNxjDtP/8RyDrGeED5id84RvIdcqZlS17mtxXg1KcALUOBm8EeRqT5yT1q6/DQWN0Q8aHP5XbVYZ9yotzoU0+uaHqjkf7lwATES/+4NpC/BlRF6uNd2oFC7pymhOhb/FbeJWJpLTHRFtdHVPQm/2VY6UH4auCaz3rDZP5Zd1sT1nsUnExII2y5NIMi7N/PNbU2vPPnXYwOrZiY7I/pGmu95r6oo3DkTyE3VdOaiXX6El6DAeNL1DRo5 tangf@DESKTOP-H68OQDV
```

```
[vagrant@localhost ~]$ echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCvfIdGduA81WVgf1F5DikDG+1qJEPk0FBYtMPk7WTEkb4p8KkqMKKdrt7Sy7Ig4ZIBwFCCU4rtHiaLeslNxwdjT1l1sH18uiNxjDtP/8RyDrGeED5id84RvIdcqZlS17mtxXg1KcALUOBm8EeRqT5yT1q6/DQWN0Q8aHP5XbVYZ9yotzoU0+uaHqjkf7lwATES/+4NpC/BlRF6uNd2oFC7pymhOhb/FbeJWJpLTHRFtdHVPQm/2VY6UH4auCaz3rDZP5Zd1sT1nsUnExII2y5NIMi7N/PNbU2vPPnXYwOrZiY7I/pGmu95r6oo3DkTyE3VdOaiXX6El6DAeNL1DRo5 tangf@DESKTOP-H68OQDV" >> .ssh/authorized_keys
```

```
[vagrant@localhost ~]$ cat .ssh/authorized_keys
ssh-rsa AAAAB3NzaC1yc2EAAAABIwAAAQEA6NF8iallvQVp22WDkTkyrtvp9eWW6A8YVr+kz4TjGYe7gHzIw+niNltGEFHzD8+v1I2YJ6oXevct1YeS0o9HZyN1Q9qgCgzUFtdOKLv6IedplqoPkcmF0aYet2PkEDo3MlTBckFXPITAMzF8dJSIFo9D8HfdOV0IAdx4O7PtixWKn5y2hMNG0zQPyUecp4pzC6kivAIhyfHilFR61RGL+GPXQ2MWZWFYbAGjyiYJnAmCP3NOTd0jMZEnDkbUvxhMmBYSdETk1rRgm+R4LOzFUGaHqHDLKLX+FIPKcF96hrucXzcWyLbIbEgE98OHlnVYCzRdK8jlqm8tehUc9c9WhQ== vagrant insecure public key
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCvfIdGduA81WVgf1F5DikDG+1qJEPk0FBYtMPk7WTEkb4p8KkqMKKdrt7Sy7Ig4ZIBwFCCU4rtHiaLeslNxwdjT1l1sH18uiNxjDtP/8RyDrGeED5id84RvIdcqZlS17mtxXg1KcALUOBm8EeRqT5yT1q6/DQWN0Q8aHP5XbVYZ9yotzoU0+uaHqjkf7lwATES/+4NpC/BlRF6uNd2oFC7pymhOhb/FbeJWJpLTHRFtdHVPQm/2VY6UH4auCaz3rDZP5Zd1sT1nsUnExII2y5NIMi7N/PNbU2vPPnXYwOrZiY7I/pGmu95r6oo3DkTyE3VdOaiXX6El6DAeNL1DRo5 tangf@DESKTOP-H68OQDV
```


Trouble shooting
```
$ ssh vagrant@10.64.33.64
The authenticity of host '10.64.33.64 (10.64.33.64)' can't be established.
ECDSA key fingerprint is SHA256:rmUJwM3oz/L+uushIeJGRatrsRIxxxh79M5JMA+Ljzc.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '10.64.33.64' (ECDSA) to the list of known hosts.
vagrant@10.64.33.64: Permission denied (publickey,gssapi-keyex,gssapi-with-mic).
```

```
tangf@DESKTOP-H68OQDV /cygdrive/g/https0x3A0x2F0x2Fapp.vagrantup.com0x2Fboxes0x2Fsearch/centos0x2Fboxes0x2F70x3A7.4
$ ssh -v vagrant@10.64.33.64
OpenSSH_7.6p1, OpenSSL 1.0.2k  26 Jan 2017
debug1: Reading configuration data /etc/ssh_config
debug1: Connecting to 10.64.33.64 [10.64.33.64] port 22.
debug1: Connection established.
debug1: identity file /home/tangf/.ssh/id_rsa type 0
debug1: key_load_public: No such file or directory
debug1: identity file /home/tangf/.ssh/id_rsa-cert type -1
debug1: key_load_public: No such file or directory
debug1: identity file /home/tangf/.ssh/id_dsa type -1
debug1: key_load_public: No such file or directory
debug1: identity file /home/tangf/.ssh/id_dsa-cert type -1
debug1: key_load_public: No such file or directory
debug1: identity file /home/tangf/.ssh/id_ecdsa type -1
debug1: key_load_public: No such file or directory
debug1: identity file /home/tangf/.ssh/id_ecdsa-cert type -1
debug1: key_load_public: No such file or directory
debug1: identity file /home/tangf/.ssh/id_ed25519 type -1
debug1: key_load_public: No such file or directory
debug1: identity file /home/tangf/.ssh/id_ed25519-cert type -1
debug1: Local version string SSH-2.0-OpenSSH_7.6
debug1: Remote protocol version 2.0, remote software version OpenSSH_7.4
debug1: match: OpenSSH_7.4 pat OpenSSH* compat 0x04000000
debug1: Authenticating to 10.64.33.64:22 as 'vagrant'
debug1: SSH2_MSG_KEXINIT sent
debug1: SSH2_MSG_KEXINIT received
debug1: kex: algorithm: curve25519-sha256
debug1: kex: host key algorithm: ecdsa-sha2-nistp256
debug1: kex: server->client cipher: chacha20-poly1305@openssh.com MAC: <implicit> compression: none
debug1: kex: client->server cipher: chacha20-poly1305@openssh.com MAC: <implicit> compression: none
debug1: expecting SSH2_MSG_KEX_ECDH_REPLY
debug1: Server host key: ecdsa-sha2-nistp256 SHA256:rmUJwM3oz/L+uushIeJGRatrsRIxxxh79M5JMA+Ljzc
debug1: Host '10.64.33.64' is known and matches the ECDSA host key.
debug1: Found key in /home/tangf/.ssh/known_hosts:3
debug1: rekey after 134217728 blocks
debug1: SSH2_MSG_NEWKEYS sent
debug1: expecting SSH2_MSG_NEWKEYS
debug1: SSH2_MSG_NEWKEYS received
debug1: rekey after 134217728 blocks
debug1: SSH2_MSG_EXT_INFO received
debug1: kex_input_ext_info: server-sig-algs=<rsa-sha2-256,rsa-sha2-512>
debug1: SSH2_MSG_SERVICE_ACCEPT received
debug1: Authentications that can continue: publickey,gssapi-keyex,gssapi-with-mic
debug1: Next authentication method: publickey
debug1: Offering public key: RSA SHA256:aQKJuefmty0v2isbhCq6QKdZUUhlVo8AVQ/pTaY1v2Q /home/tangf/.ssh/id_rsa
debug1: Authentications that can continue: publickey,gssapi-keyex,gssapi-with-mic
debug1: Trying private key: /home/tangf/.ssh/id_dsa
debug1: Trying private key: /home/tangf/.ssh/id_ecdsa
debug1: Trying private key: /home/tangf/.ssh/id_ed25519
debug1: No more authentication methods to try.
vagrant@10.64.33.64: Permission denied (publickey,gssapi-keyex,gssapi-with-mic).
```




```
tangf@DESKTOP-H68OQDV /cygdrive/g/https0x3A0x2F0x2Fapp.vagrantup.com0x2Fboxes0x2Fsearch/centos0x2Fboxes0x2F70x3A7.4
$ ssh vagrant@10.64.33.64
Last login: Fri Oct 20 17:33:00 2017 from 10.0.2.2
```

### Firewall

Check
```
[vagrant@localhost ~]$ systemctl is-active firewalld.service
unknown
```


### Chrony
__安装网络时间协议服务__

[NTP](https://docs.openstack.org/install-guide/environment-ntp-controller.html)
```
[vagrant@localhost ~]$ yum search --verbose chrony
Loading "fastestmirror" plugin
Config time: 0.004
Yum version: 3.4.3
Setting up Package Sacks
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
pkgsack time: 0.008
rpmdb time: 0.000
tags time: 0.000
============================================================== N/S matched: chrony ===============================================================
chrony.x86_64 : An NTP client/server
Repo        : @anaconda




  Name and summary matches only, use "search all" for everything.
```

```
[vagrant@localhost ~]$ systemctl -l status chronyd.service
● chronyd.service - NTP client/server
   Loaded: loaded (/usr/lib/systemd/system/chronyd.service; enabled; vendor preset: enabled)
   Active: active (running) since Fri 2017-10-20 17:28:38 UTC; 1h 33min ago
     Docs: man:chronyd(8)
           man:chrony.conf(5)
  Process: 595 ExecStartPost=/usr/libexec/chrony-helper update-daemon (code=exited, status=0/SUCCESS)
  Process: 570 ExecStart=/usr/sbin/chronyd $OPTIONS (code=exited, status=0/SUCCESS)
 Main PID: 594 (chronyd)
   CGroup: /system.slice/chronyd.service
           └─594 /usr/sbin/chronyd
```

Default _conf_
```
[vagrant@localhost ~]$ cat /etc/chrony.conf | egrep "server|allow"
# Use public servers from the pool.ntp.org project.
server 0.centos.pool.ntp.org iburst
server 1.centos.pool.ntp.org iburst
server 2.centos.pool.ntp.org iburst
server 3.centos.pool.ntp.org iburst
#allow 192.168.0.0/16
```

### Openstack Repository

[Pike Repo](https://docs.openstack.org/install-guide/environment-packages-rdo.html)
```
[vagrant@localhost ~]$ yum search openstack
Loaded plugins: fastestmirror
Determining fastest mirrors
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
============================================================= N/S matched: openstack =============================================================
centos-release-openstack-newton.noarch : OpenStack from the CentOS Cloud SIG repo configs
centos-release-openstack-ocata.noarch : OpenStack from the CentOS Cloud SIG repo configs
centos-release-openstack-pike.x86_64 : OpenStack from the CentOS Cloud SIG repo configs

  Name and summary matches only, use "search all" for everything.
```

Choose _pike_
```
[vagrant@localhost ~]$ sudo yum install -y centos-release-openstack-pike
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
Resolving Dependencies
--> Running transaction check
---> Package centos-release-openstack-pike.x86_64 0:1-1.el7 will be installed
--> Processing Dependency: centos-release-qemu-ev for package: centos-release-openstack-pike-1-1.el7.x86_64
--> Processing Dependency: centos-release-ceph-jewel for package: centos-release-openstack-pike-1-1.el7.x86_64
--> Running transaction check
---> Package centos-release-ceph-jewel.noarch 0:1.0-1.el7.centos will be installed
--> Processing Dependency: centos-release-storage-common for package: centos-release-ceph-jewel-1.0-1.el7.centos.noarch
---> Package centos-release-qemu-ev.noarch 0:1.0-2.el7 will be installed
--> Processing Dependency: centos-release-virt-common for package: centos-release-qemu-ev-1.0-2.el7.noarch
--> Running transaction check
---> Package centos-release-storage-common.noarch 0:1-2.el7.centos will be installed
---> Package centos-release-virt-common.noarch 0:1-1.el7.centos will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                                          Arch                      Version                               Repository                 Size
==================================================================================================================================================
Installing:
 centos-release-openstack-pike                    x86_64                    1-1.el7                               extras                    5.3 k
Installing for dependencies:
 centos-release-ceph-jewel                        noarch                    1.0-1.el7.centos                      extras                    4.1 k
 centos-release-qemu-ev                           noarch                    1.0-2.el7                             extras                     11 k
 centos-release-storage-common                    noarch                    1-2.el7.centos                        extras                    4.5 k
 centos-release-virt-common                       noarch                    1-1.el7.centos                        extras                    4.5 k

Transaction Summary
==================================================================================================================================================
Install  1 Package (+4 Dependent packages)

Total download size: 29 k
Installed size: 23 k
Downloading packages:
warning: /var/cache/yum/x86_64/7/extras/packages/centos-release-ceph-jewel-1.0-1.el7.centos.noarch.rpm: Header V3 RSA/SHA256 Signature, key ID f4a80eb5: NOKEY
Public key for centos-release-ceph-jewel-1.0-1.el7.centos.noarch.rpm is not installed
(1/5): centos-release-ceph-jewel-1.0-1.el7.centos.noarch.rpm                                                               | 4.1 kB  00:00:00     
(2/5): centos-release-openstack-pike-1-1.el7.x86_64.rpm                                                                    | 5.3 kB  00:00:00     
(3/5): centos-release-storage-common-1-2.el7.centos.noarch.rpm                                                             | 4.5 kB  00:00:00     
(4/5): centos-release-qemu-ev-1.0-2.el7.noarch.rpm                                                                         |  11 kB  00:00:00     
(5/5): centos-release-virt-common-1-1.el7.centos.noarch.rpm                                                                | 4.5 kB  00:00:00     
--------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                              83 kB/s |  29 kB  00:00:00     
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Importing GPG key 0xF4A80EB5:
 Userid     : "CentOS-7 Key (CentOS 7 Official Signing Key) <security@centos.org>"
 Fingerprint: 6341 ab27 53d7 8a78 a7c2 7bb1 24c6 a8a7 f4a8 0eb5
 Package    : centos-release-7-4.1708.el7.centos.x86_64 (@anaconda)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-7
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : centos-release-virt-common-1-1.el7.centos.noarch                                                                               1/5 
  Installing : centos-release-qemu-ev-1.0-2.el7.noarch                                                                                        2/5 
  Installing : centos-release-storage-common-1-2.el7.centos.noarch                                                                            3/5 
  Installing : centos-release-ceph-jewel-1.0-1.el7.centos.noarch                                                                              4/5 
  Installing : centos-release-openstack-pike-1-1.el7.x86_64                                                                                   5/5 
  Verifying  : centos-release-storage-common-1-2.el7.centos.noarch                                                                            1/5 
  Verifying  : centos-release-openstack-pike-1-1.el7.x86_64                                                                                   2/5 
  Verifying  : centos-release-ceph-jewel-1.0-1.el7.centos.noarch                                                                              3/5 
  Verifying  : centos-release-virt-common-1-1.el7.centos.noarch                                                                               4/5 
  Verifying  : centos-release-qemu-ev-1.0-2.el7.noarch                                                                                        5/5 

Installed:
  centos-release-openstack-pike.x86_64 0:1-1.el7                                                                                                  

Dependency Installed:
  centos-release-ceph-jewel.noarch 0:1.0-1.el7.centos                       centos-release-qemu-ev.noarch 0:1.0-2.el7                             
  centos-release-storage-common.noarch 0:1-2.el7.centos                     centos-release-virt-common.noarch 0:1-1.el7.centos                    

Complete!
```

```
[vagrant@localhost ~]$ ls -1 /etc/yum.repos.d/ | egrep 'OpenStack|Ceph'
CentOS-Ceph-Jewel.repo
CentOS-OpenStack-pike.repo
```

```
[vagrant@localhost ~]$ head /etc/yum.repos.d/CentOS-OpenStack-pike.repo 
# CentOS-OpenStack-pike.repo
#
# Please see http://wiki.centos.org/SpecialInterestGroup/Cloud for more
# information

[centos-openstack-pike]
name=CentOS-7 - OpenStack pike
baseurl=http://mirror.centos.org/centos/7/cloud/$basearch/openstack-pike/
gpgcheck=1
enabled=1
```

### Openstack client

CLI
```
[vagrant@localhost ~]$ sudo yum install -y python-openstackclient
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
Resolving Dependencies
--> Running transaction check
---> Package python2-openstackclient.noarch 0:3.12.0-1.el7 will be installed
--> Processing Dependency: python-openstackclient-lang = 3.12.0-1.el7 for package: python2-openstackclient-3.12.0-1.el7.noarch
--> Processing Dependency: python-osc-lib >= 1.7.0 for package: python2-openstackclient-3.12.0-1.el7.noarch
--> Processing Dependency: python-openstacksdk >= 0.9.17 for package: python2-openstackclient-3.12.0-1.el7.noarch
--> Processing Dependency: python-novaclient >= 1:9.0.0 for package: python2-openstackclient-3.12.0-1.el7.noarch
--> Processing Dependency: python-neutronclient >= 6.3.0 for package: python2-openstackclient-3.12.0-1.el7.noarch
--> Processing Dependency: python-glanceclient >= 1:2.8.0 for package: python2-openstackclient-3.12.0-1.el7.noarch
--> Processing Dependency: python-cinderclient >= 3.1.0 for package: python2-openstackclient-3.12.0-1.el7.noarch
--> Processing Dependency: python-cliff for package: python2-openstackclient-3.12.0-1.el7.noarch
--> Running transaction check
---> Package python-cliff.noarch 0:2.8.0-1.el7 will be installed
--> Processing Dependency: python-cmd2 >= 0.6.7 for package: python-cliff-2.8.0-1.el7.noarch
--> Processing Dependency: python-unicodecsv for package: python-cliff-2.8.0-1.el7.noarch
---> Package python-openstackclient-lang.noarch 0:3.12.0-1.el7 will be installed
---> Package python2-cinderclient.noarch 0:3.1.0-1.el7 will be installed
--> Processing Dependency: python-simplejson for package: python2-cinderclient-3.1.0-1.el7.noarch
---> Package python2-glanceclient.noarch 1:2.8.0-1.el7 will be installed
--> Processing Dependency: python-warlock for package: 1:python2-glanceclient-2.8.0-1.el7.noarch
---> Package python2-neutronclient.noarch 0:6.5.0-1.el7 will be installed
--> Processing Dependency: python-os-client-config >= 1.28.0 for package: python2-neutronclient-6.5.0-1.el7.noarch
---> Package python2-novaclient.noarch 1:9.1.1-1.el7 will be installed
---> Package python2-openstacksdk.noarch 0:0.9.17-1.el7 will be installed
--> Processing Dependency: python-jsonpatch >= 1.1 for package: python2-openstacksdk-0.9.17-1.el7.noarch
--> Processing Dependency: python-deprecation for package: python2-openstacksdk-0.9.17-1.el7.noarch
---> Package python2-osc-lib.noarch 0:1.7.0-1.el7 will be installed
--> Running transaction check
---> Package python-cmd2.noarch 0:0.6.8-8.el7 will be installed
---> Package python-simplejson.x86_64 0:3.5.3-5.el7 will be installed
---> Package python-unicodecsv.noarch 0:0.14.1-1.el7 will be installed
---> Package python-warlock.noarch 0:1.0.1-1.el7 will be installed
---> Package python2-deprecation.noarch 0:1.0-3.el7 will be installed
---> Package python2-jsonpatch.noarch 0:1.14-1.el7 will be installed
--> Processing Dependency: python-jsonpointer for package: python2-jsonpatch-1.14-1.el7.noarch
---> Package python2-os-client-config.noarch 0:1.28.0-1.el7 will be installed
--> Processing Dependency: python-requestsexceptions >= 1.2.0 for package: python2-os-client-config-1.28.0-1.el7.noarch
--> Processing Dependency: python-appdirs >= 1.3.0 for package: python2-os-client-config-1.28.0-1.el7.noarch
--> Running transaction check
---> Package python2-appdirs.noarch 0:1.4.0-4.el7 will be installed
---> Package python2-jsonpointer.noarch 0:1.10-4.el7 will be installed
---> Package python2-requestsexceptions.noarch 0:1.3.0-1.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                                      Arch                    Version                        Repository                              Size
==================================================================================================================================================
Installing:
 python2-openstackclient                      noarch                  3.12.0-1.el7                   centos-openstack-pike                  1.0 M
Installing for dependencies:
 python-cliff                                 noarch                  2.8.0-1.el7                    centos-openstack-pike                   81 k
 python-cmd2                                  noarch                  0.6.8-8.el7                    centos-openstack-pike                   41 k
 python-openstackclient-lang                  noarch                  3.12.0-1.el7                   centos-openstack-pike                   14 k
 python-simplejson                            x86_64                  3.5.3-5.el7                    centos-openstack-pike                  185 k
 python-unicodecsv                            noarch                  0.14.1-1.el7                   centos-openstack-pike                   25 k
 python-warlock                               noarch                  1.0.1-1.el7                    centos-openstack-pike                   14 k
 python2-appdirs                              noarch                  1.4.0-4.el7                    centos-openstack-pike                   16 k
 python2-cinderclient                         noarch                  3.1.0-1.el7                    centos-openstack-pike                  238 k
 python2-deprecation                          noarch                  1.0-3.el7                      centos-openstack-pike                   15 k
 python2-glanceclient                         noarch                  1:2.8.0-1.el7                  centos-openstack-pike                  134 k
 python2-jsonpatch                            noarch                  1.14-1.el7                     centos-openstack-pike                   22 k
 python2-jsonpointer                          noarch                  1.10-4.el7                     centos-openstack-pike                   14 k
 python2-neutronclient                        noarch                  6.5.0-1.el7                    centos-openstack-pike                  269 k
 python2-novaclient                           noarch                  1:9.1.1-1.el7                  centos-openstack-pike                  211 k
 python2-openstacksdk                         noarch                  0.9.17-1.el7                   centos-openstack-pike                  334 k
 python2-os-client-config                     noarch                  1.28.0-1.el7                   centos-openstack-pike                   79 k
 python2-osc-lib                              noarch                  1.7.0-1.el7                    centos-openstack-pike                   65 k
 python2-requestsexceptions                   noarch                  1.3.0-1.el7                    centos-openstack-pike                   11 k

Transaction Summary
==================================================================================================================================================
Install  1 Package (+18 Dependent packages)

Total download size: 2.7 M
Installed size: 15 M
Downloading packages:
(1/19): python-cmd2-0.6.8-8.el7.noarch.rpm                                                                                 |  41 kB  00:00:01     
(2/19): python-openstackclient-lang-3.12.0-1.el7.noarch.rpm                                                                |  14 kB  00:00:00     
(3/19): python-cliff-2.8.0-1.el7.noarch.rpm                                                                                |  81 kB  00:00:01     
(4/19): python-unicodecsv-0.14.1-1.el7.noarch.rpm                                                                          |  25 kB  00:00:00     
(5/19): python-warlock-1.0.1-1.el7.noarch.rpm                                                                              |  14 kB  00:00:00     
(6/19): python2-appdirs-1.4.0-4.el7.noarch.rpm                                                                             |  16 kB  00:00:00     
(7/19): python-simplejson-3.5.3-5.el7.x86_64.rpm                                                                           | 185 kB  00:00:01     
(8/19): python2-deprecation-1.0-3.el7.noarch.rpm                                                                           |  15 kB  00:00:00     
(9/19): python2-glanceclient-2.8.0-1.el7.noarch.rpm                                                                        | 134 kB  00:00:00     
(10/19): python2-cinderclient-3.1.0-1.el7.noarch.rpm                                                                       | 238 kB  00:00:00     
(11/19): python2-jsonpatch-1.14-1.el7.noarch.rpm                                                                           |  22 kB  00:00:00     
(12/19): python2-jsonpointer-1.10-4.el7.noarch.rpm                                                                         |  14 kB  00:00:00     
(13/19): python2-neutronclient-6.5.0-1.el7.noarch.rpm                                                                      | 269 kB  00:00:00     
(14/19): python2-novaclient-9.1.1-1.el7.noarch.rpm                                                                         | 211 kB  00:00:01     
(15/19): python2-openstackclient-3.12.0-1.el7.noarch.rpm                                                                   | 1.0 MB  00:00:01     
(16/19): python2-os-client-config-1.28.0-1.el7.noarch.rpm                                                                  |  79 kB  00:00:00     
(17/19): python2-osc-lib-1.7.0-1.el7.noarch.rpm                                                                            |  65 kB  00:00:00     
(18/19): python2-requestsexceptions-1.3.0-1.el7.noarch.rpm                                                                 |  11 kB  00:00:00     
(19/19): python2-openstacksdk-0.9.17-1.el7.noarch.rpm                                                                      | 334 kB  00:00:01     
--------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                             432 kB/s | 2.7 MB  00:00:06     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : python-simplejson-3.5.3-5.el7.x86_64                                                                                          1/19 
  Installing : 1:python2-novaclient-9.1.1-1.el7.noarch                                                                                       2/19 
  Installing : python2-cinderclient-3.1.0-1.el7.noarch                                                                                       3/19 
  Installing : python2-jsonpointer-1.10-4.el7.noarch                                                                                         4/19 
  Installing : python2-jsonpatch-1.14-1.el7.noarch                                                                                           5/19 
  Installing : python-warlock-1.0.1-1.el7.noarch                                                                                             6/19 
  Installing : 1:python2-glanceclient-2.8.0-1.el7.noarch                                                                                     7/19 
  Installing : python2-appdirs-1.4.0-4.el7.noarch                                                                                            8/19 
  Installing : python-openstackclient-lang-3.12.0-1.el7.noarch                                                                               9/19 
  Installing : python-cmd2-0.6.8-8.el7.noarch                                                                                               10/19 
  Installing : python2-requestsexceptions-1.3.0-1.el7.noarch                                                                                11/19 
  Installing : python2-os-client-config-1.28.0-1.el7.noarch                                                                                 12/19 
  Installing : python-unicodecsv-0.14.1-1.el7.noarch                                                                                        13/19 
  Installing : python-cliff-2.8.0-1.el7.noarch                                                                                              14/19 
  Installing : python2-osc-lib-1.7.0-1.el7.noarch                                                                                           15/19 
  Installing : python2-neutronclient-6.5.0-1.el7.noarch                                                                                     16/19 
  Installing : python2-deprecation-1.0-3.el7.noarch                                                                                         17/19 
  Installing : python2-openstacksdk-0.9.17-1.el7.noarch                                                                                     18/19 
  Installing : python2-openstackclient-3.12.0-1.el7.noarch                                                                                  19/19 
  Verifying  : python2-deprecation-1.0-3.el7.noarch                                                                                          1/19 
  Verifying  : python2-os-client-config-1.28.0-1.el7.noarch                                                                                  2/19 
  Verifying  : python-simplejson-3.5.3-5.el7.x86_64                                                                                          3/19 
  Verifying  : 1:python2-novaclient-9.1.1-1.el7.noarch                                                                                       4/19 
  Verifying  : python-cliff-2.8.0-1.el7.noarch                                                                                               5/19 
  Verifying  : python2-openstacksdk-0.9.17-1.el7.noarch                                                                                      6/19 
  Verifying  : python2-jsonpatch-1.14-1.el7.noarch                                                                                           7/19 
  Verifying  : python2-cinderclient-3.1.0-1.el7.noarch                                                                                       8/19 
  Verifying  : python2-openstackclient-3.12.0-1.el7.noarch                                                                                   9/19 
  Verifying  : python-unicodecsv-0.14.1-1.el7.noarch                                                                                        10/19 
  Verifying  : python2-requestsexceptions-1.3.0-1.el7.noarch                                                                                11/19 
  Verifying  : python2-neutronclient-6.5.0-1.el7.noarch                                                                                     12/19 
  Verifying  : python-cmd2-0.6.8-8.el7.noarch                                                                                               13/19 
  Verifying  : 1:python2-glanceclient-2.8.0-1.el7.noarch                                                                                    14/19 
  Verifying  : python-warlock-1.0.1-1.el7.noarch                                                                                            15/19 
  Verifying  : python-openstackclient-lang-3.12.0-1.el7.noarch                                                                              16/19 
  Verifying  : python2-appdirs-1.4.0-4.el7.noarch                                                                                           17/19 
  Verifying  : python2-jsonpointer-1.10-4.el7.noarch                                                                                        18/19 
  Verifying  : python2-osc-lib-1.7.0-1.el7.noarch                                                                                           19/19 

Installed:
  python2-openstackclient.noarch 0:3.12.0-1.el7                                                                                                   

Dependency Installed:
  python-cliff.noarch 0:2.8.0-1.el7                python-cmd2.noarch 0:0.6.8-8.el7            python-openstackclient-lang.noarch 0:3.12.0-1.el7  
  python-simplejson.x86_64 0:3.5.3-5.el7           python-unicodecsv.noarch 0:0.14.1-1.el7     python-warlock.noarch 0:1.0.1-1.el7                
  python2-appdirs.noarch 0:1.4.0-4.el7             python2-cinderclient.noarch 0:3.1.0-1.el7   python2-deprecation.noarch 0:1.0-3.el7             
  python2-glanceclient.noarch 1:2.8.0-1.el7        python2-jsonpatch.noarch 0:1.14-1.el7       python2-jsonpointer.noarch 0:1.10-4.el7            
  python2-neutronclient.noarch 0:6.5.0-1.el7       python2-novaclient.noarch 1:9.1.1-1.el7     python2-openstacksdk.noarch 0:0.9.17-1.el7         
  python2-os-client-config.noarch 0:1.28.0-1.el7   python2-osc-lib.noarch 0:1.7.0-1.el7        python2-requestsexceptions.noarch 0:1.3.0-1.el7    

Complete!
```

```
[vagrant@localhost ~]$ openstack --version
openstack 3.12.0
```

### Database



[MariaDB](https://docs.openstack.org/install-guide/environment-sql-database-rdo.html)
```
[vagrant@localhost ~]$ yum list | egrep '^mariadb\.'
mariadb.x86_64                     3:10.1.20-1.el7         centos-openstack-pike
```

Install
```
[vagrant@localhost ~]$ sudo yum install -y mariadb mariadb-server python2-PyMySQL
Loaded plugins: fastestmirror
centos-ceph-jewel                                                                                                          | 2.9 kB  00:00:00     
centos-openstack-pike                                                                                                      | 2.9 kB  00:00:00     
centos-qemu-ev                                                                                                             | 2.9 kB  00:00:00     
(1/3): centos-qemu-ev/7/x86_64/primary_db                                                                                  |  23 kB  00:00:01     
(2/3): centos-ceph-jewel/7/x86_64/primary_db                                                                               |  42 kB  00:00:01     
(3/3): centos-openstack-pike/x86_64/primary_db                                                                             | 758 kB  00:00:03     
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
Resolving Dependencies
--> Running transaction check
---> Package mariadb.x86_64 3:10.1.20-1.el7 will be installed
--> Processing Dependency: mariadb-libs(x86-64) = 3:10.1.20-1.el7 for package: 3:mariadb-10.1.20-1.el7.x86_64
--> Processing Dependency: mariadb-common(x86-64) = 3:10.1.20-1.el7 for package: 3:mariadb-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(Sys::Hostname) for package: 3:mariadb-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(IPC::Open3) for package: 3:mariadb-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(Getopt::Long) for package: 3:mariadb-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(File::Temp) for package: 3:mariadb-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(Fcntl) for package: 3:mariadb-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(Exporter) for package: 3:mariadb-10.1.20-1.el7.x86_64
--> Processing Dependency: /usr/bin/perl for package: 3:mariadb-10.1.20-1.el7.x86_64
---> Package mariadb-server.x86_64 3:10.1.20-1.el7 will be installed
--> Processing Dependency: mariadb-errmsg(x86-64) = 3:10.1.20-1.el7 for package: 3:mariadb-server-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(File::Path) for package: 3:mariadb-server-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(Data::Dumper) for package: 3:mariadb-server-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(DBI) for package: 3:mariadb-server-10.1.20-1.el7.x86_64
--> Processing Dependency: perl(DBD::mysql) for package: 3:mariadb-server-10.1.20-1.el7.x86_64
--> Processing Dependency: net-tools for package: 3:mariadb-server-10.1.20-1.el7.x86_64
--> Processing Dependency: lsof for package: 3:mariadb-server-10.1.20-1.el7.x86_64
--> Processing Dependency: /etc/my.cnf.d for package: 3:mariadb-server-10.1.20-1.el7.x86_64
--> Processing Dependency: /etc/my.cnf for package: 3:mariadb-server-10.1.20-1.el7.x86_64
---> Package python2-PyMySQL.noarch 0:0.7.11-1.el7 will be installed
--> Running transaction check
---> Package lsof.x86_64 0:4.87-4.el7 will be installed
---> Package mariadb-common.x86_64 3:10.1.20-1.el7 will be installed
---> Package mariadb-config.x86_64 3:10.1.20-1.el7 will be installed
---> Package mariadb-errmsg.x86_64 3:10.1.20-1.el7 will be installed
---> Package mariadb-libs.x86_64 1:5.5.56-2.el7 will be updated
---> Package mariadb-libs.x86_64 1:5.5.56-2.el7 will be updated
---> Package mariadb-libs.x86_64 3:10.1.20-1.el7 will be an update
---> Package net-tools.x86_64 0:2.0-0.22.20131004git.el7 will be installed
---> Package perl.x86_64 4:5.16.3-292.el7 will be installed
--> Processing Dependency: perl-libs = 4:5.16.3-292.el7 for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Socket) >= 1.3 for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Scalar::Util) >= 1.10 for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl-macros for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl-libs for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(threads::shared) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(threads) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(constant) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Time::Local) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Time::HiRes) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Storable) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Socket) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Scalar::Util) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Pod::Simple::XHTML) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Pod::Simple::Search) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Filter::Util::Call) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(File::Spec::Unix) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(File::Spec::Functions) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(File::Spec) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Cwd) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: perl(Carp) for package: 4:perl-5.16.3-292.el7.x86_64
--> Processing Dependency: libperl.so()(64bit) for package: 4:perl-5.16.3-292.el7.x86_64
---> Package perl-DBD-MySQL.x86_64 0:4.023-5.el7 will be installed
---> Package perl-DBI.x86_64 0:1.627-4.el7 will be installed
--> Processing Dependency: perl(RPC::PlServer) >= 0.2001 for package: perl-DBI-1.627-4.el7.x86_64
--> Processing Dependency: perl(RPC::PlClient) >= 0.2000 for package: perl-DBI-1.627-4.el7.x86_64
---> Package perl-Data-Dumper.x86_64 0:2.145-3.el7 will be installed
---> Package perl-Exporter.noarch 0:5.68-3.el7 will be installed
---> Package perl-File-Path.noarch 0:2.09-2.el7 will be installed
---> Package perl-File-Temp.noarch 0:0.23.01-3.el7 will be installed
---> Package perl-Getopt-Long.noarch 0:2.40-2.el7 will be installed
--> Processing Dependency: perl(Pod::Usage) >= 1.14 for package: perl-Getopt-Long-2.40-2.el7.noarch
--> Processing Dependency: perl(Text::ParseWords) for package: perl-Getopt-Long-2.40-2.el7.noarch
--> Running transaction check
---> Package perl-Carp.noarch 0:1.26-244.el7 will be installed
---> Package perl-Filter.x86_64 0:1.49-3.el7 will be installed
---> Package perl-PathTools.x86_64 0:3.40-5.el7 will be installed
---> Package perl-PlRPC.noarch 0:0.2020-14.el7 will be installed
--> Processing Dependency: perl(Net::Daemon) >= 0.13 for package: perl-PlRPC-0.2020-14.el7.noarch
--> Processing Dependency: perl(Net::Daemon::Test) for package: perl-PlRPC-0.2020-14.el7.noarch
--> Processing Dependency: perl(Net::Daemon::Log) for package: perl-PlRPC-0.2020-14.el7.noarch
--> Processing Dependency: perl(Compress::Zlib) for package: perl-PlRPC-0.2020-14.el7.noarch
---> Package perl-Pod-Simple.noarch 1:3.28-4.el7 will be installed
--> Processing Dependency: perl(Pod::Escapes) >= 1.04 for package: 1:perl-Pod-Simple-3.28-4.el7.noarch
--> Processing Dependency: perl(Encode) for package: 1:perl-Pod-Simple-3.28-4.el7.noarch
---> Package perl-Pod-Usage.noarch 0:1.63-3.el7 will be installed
--> Processing Dependency: perl(Pod::Text) >= 3.15 for package: perl-Pod-Usage-1.63-3.el7.noarch
--> Processing Dependency: perl-Pod-Perldoc for package: perl-Pod-Usage-1.63-3.el7.noarch
---> Package perl-Scalar-List-Utils.x86_64 0:1.27-248.el7 will be installed
---> Package perl-Socket.x86_64 0:2.010-4.el7 will be installed
---> Package perl-Storable.x86_64 0:2.45-3.el7 will be installed
---> Package perl-Text-ParseWords.noarch 0:3.29-4.el7 will be installed
---> Package perl-Time-HiRes.x86_64 4:1.9725-3.el7 will be installed
---> Package perl-Time-Local.noarch 0:1.2300-2.el7 will be installed
---> Package perl-constant.noarch 0:1.27-2.el7 will be installed
---> Package perl-libs.x86_64 4:5.16.3-292.el7 will be installed
---> Package perl-macros.x86_64 4:5.16.3-292.el7 will be installed
---> Package perl-threads.x86_64 0:1.87-4.el7 will be installed
---> Package perl-threads-shared.x86_64 0:1.43-6.el7 will be installed
--> Running transaction check
---> Package perl-Encode.x86_64 0:2.51-7.el7 will be installed
---> Package perl-IO-Compress.noarch 0:2.061-2.el7 will be installed
--> Processing Dependency: perl(Compress::Raw::Zlib) >= 2.061 for package: perl-IO-Compress-2.061-2.el7.noarch
--> Processing Dependency: perl(Compress::Raw::Bzip2) >= 2.061 for package: perl-IO-Compress-2.061-2.el7.noarch
---> Package perl-Net-Daemon.noarch 0:0.48-5.el7 will be installed
---> Package perl-Pod-Escapes.noarch 1:1.04-292.el7 will be installed
---> Package perl-Pod-Perldoc.noarch 0:3.20-4.el7 will be installed
--> Processing Dependency: perl(parent) for package: perl-Pod-Perldoc-3.20-4.el7.noarch
--> Processing Dependency: perl(HTTP::Tiny) for package: perl-Pod-Perldoc-3.20-4.el7.noarch
---> Package perl-podlators.noarch 0:2.5.1-3.el7 will be installed
--> Running transaction check
---> Package perl-Compress-Raw-Bzip2.x86_64 0:2.061-3.el7 will be installed
---> Package perl-Compress-Raw-Zlib.x86_64 1:2.061-4.el7 will be installed
---> Package perl-HTTP-Tiny.noarch 0:0.033-3.el7 will be installed
---> Package perl-parent.noarch 1:0.225-244.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                                Arch                  Version                                  Repository                            Size
==================================================================================================================================================
Installing:
 mariadb                                x86_64                3:10.1.20-1.el7                          centos-openstack-pike                6.3 M
 mariadb-server                         x86_64                3:10.1.20-1.el7                          centos-openstack-pike                 19 M
 python2-PyMySQL                        noarch                0.7.11-1.el7                             centos-openstack-pike                150 k
Installing for dependencies:
 lsof                                   x86_64                4.87-4.el7                               base                                 331 k
 mariadb-common                         x86_64                3:10.1.20-1.el7                          centos-openstack-pike                 63 k
 mariadb-config                         x86_64                3:10.1.20-1.el7                          centos-openstack-pike                 26 k
 mariadb-errmsg                         x86_64                3:10.1.20-1.el7                          centos-openstack-pike                200 k
 net-tools                              x86_64                2.0-0.22.20131004git.el7                 base                                 305 k
 perl                                   x86_64                4:5.16.3-292.el7                         base                                 8.0 M
 perl-Carp                              noarch                1.26-244.el7                             base                                  19 k
 perl-Compress-Raw-Bzip2                x86_64                2.061-3.el7                              base                                  32 k
 perl-Compress-Raw-Zlib                 x86_64                1:2.061-4.el7                            base                                  57 k
 perl-DBD-MySQL                         x86_64                4.023-5.el7                              base                                 140 k
 perl-DBI                               x86_64                1.627-4.el7                              base                                 802 k
 perl-Data-Dumper                       x86_64                2.145-3.el7                              base                                  47 k
 perl-Encode                            x86_64                2.51-7.el7                               base                                 1.5 M
 perl-Exporter                          noarch                5.68-3.el7                               base                                  28 k
 perl-File-Path                         noarch                2.09-2.el7                               base                                  26 k
 perl-File-Temp                         noarch                0.23.01-3.el7                            base                                  56 k
 perl-Filter                            x86_64                1.49-3.el7                               base                                  76 k
 perl-Getopt-Long                       noarch                2.40-2.el7                               base                                  56 k
 perl-HTTP-Tiny                         noarch                0.033-3.el7                              base                                  38 k
 perl-IO-Compress                       noarch                2.061-2.el7                              base                                 260 k
 perl-Net-Daemon                        noarch                0.48-5.el7                               base                                  51 k
 perl-PathTools                         x86_64                3.40-5.el7                               base                                  82 k
 perl-PlRPC                             noarch                0.2020-14.el7                            base                                  36 k
 perl-Pod-Escapes                       noarch                1:1.04-292.el7                           base                                  51 k
 perl-Pod-Perldoc                       noarch                3.20-4.el7                               base                                  87 k
 perl-Pod-Simple                        noarch                1:3.28-4.el7                             base                                 216 k
 perl-Pod-Usage                         noarch                1.63-3.el7                               base                                  27 k
 perl-Scalar-List-Utils                 x86_64                1.27-248.el7                             base                                  36 k
 perl-Socket                            x86_64                2.010-4.el7                              base                                  49 k
 perl-Storable                          x86_64                2.45-3.el7                               base                                  77 k
 perl-Text-ParseWords                   noarch                3.29-4.el7                               base                                  14 k
 perl-Time-HiRes                        x86_64                4:1.9725-3.el7                           base                                  45 k
 perl-Time-Local                        noarch                1.2300-2.el7                             base                                  24 k
 perl-constant                          noarch                1.27-2.el7                               base                                  19 k
 perl-libs                              x86_64                4:5.16.3-292.el7                         base                                 688 k
 perl-macros                            x86_64                4:5.16.3-292.el7                         base                                  43 k
 perl-parent                            noarch                1:0.225-244.el7                          base                                  12 k
 perl-podlators                         noarch                2.5.1-3.el7                              base                                 112 k
 perl-threads                           x86_64                1.87-4.el7                               base                                  49 k
 perl-threads-shared                    x86_64                1.43-6.el7                               base                                  39 k
Updating for dependencies:
 mariadb-libs                           x86_64                3:10.1.20-1.el7                          centos-openstack-pike                643 k

Transaction Summary
==================================================================================================================================================
Install  3 Packages (+40 Dependent packages)
Upgrade             (  1 Dependent package)

Total download size: 40 M
Downloading packages:
No Presto metadata available for centos-openstack-pike
(1/44): lsof-4.87-4.el7.x86_64.rpm                                                                                         | 331 kB  00:00:00     
warning: /var/cache/yum/x86_64/7/centos-openstack-pike/packages/mariadb-common-10.1.20-1.el7.x86_64.rpm: Header V4 RSA/SHA1 Signature, key ID 764429e6: NOKEY
Public key for mariadb-common-10.1.20-1.el7.x86_64.rpm is not installed
(2/44): mariadb-common-10.1.20-1.el7.x86_64.rpm                                                                            |  63 kB  00:00:02     
(3/44): mariadb-config-10.1.20-1.el7.x86_64.rpm                                                                            |  26 kB  00:00:00     
(4/44): mariadb-errmsg-10.1.20-1.el7.x86_64.rpm                                                                            | 200 kB  00:00:01     
(5/44): mariadb-libs-10.1.20-1.el7.x86_64.rpm                                                                              | 643 kB  00:00:01     
(6/44): perl-Carp-1.26-244.el7.noarch.rpm                                                                                  |  19 kB  00:00:00     
(7/44): perl-Compress-Raw-Bzip2-2.061-3.el7.x86_64.rpm                                                                     |  32 kB  00:00:00     
(8/44): perl-Compress-Raw-Zlib-2.061-4.el7.x86_64.rpm                                                                      |  57 kB  00:00:00     
(9/44): net-tools-2.0-0.22.20131004git.el7.x86_64.rpm                                                                      | 305 kB  00:00:00     
(10/44): perl-DBD-MySQL-4.023-5.el7.x86_64.rpm                                                                             | 140 kB  00:00:00     
(11/44): perl-Data-Dumper-2.145-3.el7.x86_64.rpm                                                                           |  47 kB  00:00:00     
(12/44): perl-Encode-2.51-7.el7.x86_64.rpm                                                                                 | 1.5 MB  00:00:00     
(13/44): perl-Exporter-5.68-3.el7.noarch.rpm                                                                               |  28 kB  00:00:00     
(14/44): perl-File-Path-2.09-2.el7.noarch.rpm                                                                              |  26 kB  00:00:00     
(15/44): perl-File-Temp-0.23.01-3.el7.noarch.rpm                                                                           |  56 kB  00:00:00     
(16/44): perl-Filter-1.49-3.el7.x86_64.rpm                                                                                 |  76 kB  00:00:00     
(17/44): perl-Getopt-Long-2.40-2.el7.noarch.rpm                                                                            |  56 kB  00:00:00     
(18/44): perl-HTTP-Tiny-0.033-3.el7.noarch.rpm                                                                             |  38 kB  00:00:00     
(19/44): perl-IO-Compress-2.061-2.el7.noarch.rpm                                                                           | 260 kB  00:00:00     
(20/44): perl-Net-Daemon-0.48-5.el7.noarch.rpm                                                                             |  51 kB  00:00:00     
(21/44): perl-PathTools-3.40-5.el7.x86_64.rpm                                                                              |  82 kB  00:00:00     
(22/44): perl-PlRPC-0.2020-14.el7.noarch.rpm                                                                               |  36 kB  00:00:00     
(23/44): perl-Pod-Escapes-1.04-292.el7.noarch.rpm                                                                          |  51 kB  00:00:00     
(24/44): perl-Pod-Perldoc-3.20-4.el7.noarch.rpm                                                                            |  87 kB  00:00:00     
(25/44): perl-DBI-1.627-4.el7.x86_64.rpm                                                                                   | 802 kB  00:00:02     
(26/44): perl-Pod-Usage-1.63-3.el7.noarch.rpm                                                                              |  27 kB  00:00:00     
(27/44): perl-Pod-Simple-3.28-4.el7.noarch.rpm                                                                             | 216 kB  00:00:00     
(28/44): perl-Socket-2.010-4.el7.x86_64.rpm                                                                                |  49 kB  00:00:00     
(29/44): perl-Scalar-List-Utils-1.27-248.el7.x86_64.rpm                                                                    |  36 kB  00:00:00     
(30/44): perl-5.16.3-292.el7.x86_64.rpm                                                                                    | 8.0 MB  00:00:02     
(31/44): perl-Text-ParseWords-3.29-4.el7.noarch.rpm                                                                        |  14 kB  00:00:00     
(32/44): perl-Storable-2.45-3.el7.x86_64.rpm                                                                               |  77 kB  00:00:00     
(33/44): perl-Time-Local-1.2300-2.el7.noarch.rpm                                                                           |  24 kB  00:00:00     
(34/44): perl-Time-HiRes-1.9725-3.el7.x86_64.rpm                                                                           |  45 kB  00:00:00     
(35/44): perl-constant-1.27-2.el7.noarch.rpm                                                                               |  19 kB  00:00:00     
(36/44): perl-macros-5.16.3-292.el7.x86_64.rpm                                                                             |  43 kB  00:00:00     
(37/44): perl-parent-0.225-244.el7.noarch.rpm                                                                              |  12 kB  00:00:00     
(38/44): perl-podlators-2.5.1-3.el7.noarch.rpm                                                                             | 112 kB  00:00:00     
(39/44): perl-libs-5.16.3-292.el7.x86_64.rpm                                                                               | 688 kB  00:00:00     
(40/44): perl-threads-1.87-4.el7.x86_64.rpm                                                                                |  49 kB  00:00:00     
(41/44): perl-threads-shared-1.43-6.el7.x86_64.rpm                                                                         |  39 kB  00:00:00     
(42/44): mariadb-10.1.20-1.el7.x86_64.rpm                                                                                  | 6.3 MB  00:00:31     
(43/44): python2-PyMySQL-0.7.11-1.el7.noarch.rpm                                                                           | 150 kB  00:00:00     
(44/44): mariadb-server-10.1.20-1.el7.x86_64.rpm                                                                           |  19 MB  00:00:37     
--------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                             956 kB/s |  40 MB  00:00:42     
Retrieving key from file:///etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-SIG-Cloud
Importing GPG key 0x764429E6:
 Userid     : "CentOS Cloud SIG (http://wiki.centos.org/SpecialInterestGroup/Cloud) <security@centos.org>"
 Fingerprint: 736a f511 6d9c 40e2 af6b 074b f9b9 fee7 7644 29e6
 Package    : centos-release-openstack-pike-1-1.el7.x86_64 (@extras)
 From       : /etc/pki/rpm-gpg/RPM-GPG-KEY-CentOS-SIG-Cloud
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : 3:mariadb-config-10.1.20-1.el7.x86_64                                                                                         1/45 
  Installing : 3:mariadb-common-10.1.20-1.el7.x86_64                                                                                         2/45 
  Updating   : 3:mariadb-libs-10.1.20-1.el7.x86_64                                                                                           3/45 
  Installing : 3:mariadb-errmsg-10.1.20-1.el7.x86_64                                                                                         4/45 
  Installing : 1:perl-parent-0.225-244.el7.noarch                                                                                            5/45 
  Installing : perl-HTTP-Tiny-0.033-3.el7.noarch                                                                                             6/45 
  Installing : perl-podlators-2.5.1-3.el7.noarch                                                                                             7/45 
  Installing : perl-Pod-Perldoc-3.20-4.el7.noarch                                                                                            8/45 
  Installing : 1:perl-Pod-Escapes-1.04-292.el7.noarch                                                                                        9/45 
  Installing : perl-Text-ParseWords-3.29-4.el7.noarch                                                                                       10/45 
  Installing : perl-Encode-2.51-7.el7.x86_64                                                                                                11/45 
  Installing : perl-Pod-Usage-1.63-3.el7.noarch                                                                                             12/45 
  Installing : 4:perl-macros-5.16.3-292.el7.x86_64                                                                                          13/45 
  Installing : 4:perl-libs-5.16.3-292.el7.x86_64                                                                                            14/45 
  Installing : perl-threads-1.87-4.el7.x86_64                                                                                               15/45 
  Installing : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                                                                                        16/45 
  Installing : perl-Exporter-5.68-3.el7.noarch                                                                                              17/45 
  Installing : perl-constant-1.27-2.el7.noarch                                                                                              18/45 
  Installing : perl-Time-Local-1.2300-2.el7.noarch                                                                                          19/45 
  Installing : perl-Socket-2.010-4.el7.x86_64                                                                                               20/45 
  Installing : perl-Carp-1.26-244.el7.noarch                                                                                                21/45 
  Installing : perl-Storable-2.45-3.el7.x86_64                                                                                              22/45 
  Installing : perl-Filter-1.49-3.el7.x86_64                                                                                                23/45 
  Installing : perl-threads-shared-1.43-6.el7.x86_64                                                                                        24/45 
  Installing : perl-File-Path-2.09-2.el7.noarch                                                                                             25/45 
  Installing : perl-PathTools-3.40-5.el7.x86_64                                                                                             26/45 
  Installing : perl-File-Temp-0.23.01-3.el7.noarch                                                                                          27/45 
  Installing : perl-Scalar-List-Utils-1.27-248.el7.x86_64                                                                                   28/45 
  Installing : 1:perl-Pod-Simple-3.28-4.el7.noarch                                                                                          29/45 
  Installing : perl-Getopt-Long-2.40-2.el7.noarch                                                                                           30/45 
  Installing : 4:perl-5.16.3-292.el7.x86_64                                                                                                 31/45 
  Installing : perl-Data-Dumper-2.145-3.el7.x86_64                                                                                          32/45 
  Installing : perl-Compress-Raw-Bzip2-2.061-3.el7.x86_64                                                                                   33/45 
  Installing : perl-Net-Daemon-0.48-5.el7.noarch                                                                                            34/45 
  Installing : 3:mariadb-10.1.20-1.el7.x86_64                                                                                               35/45 
  Installing : 1:perl-Compress-Raw-Zlib-2.061-4.el7.x86_64                                                                                  36/45 
  Installing : perl-IO-Compress-2.061-2.el7.noarch                                                                                          37/45 
  Installing : perl-PlRPC-0.2020-14.el7.noarch                                                                                              38/45 
  Installing : perl-DBI-1.627-4.el7.x86_64                                                                                                  39/45 
  Installing : perl-DBD-MySQL-4.023-5.el7.x86_64                                                                                            40/45 
  Installing : lsof-4.87-4.el7.x86_64                                                                                                       41/45 
  Installing : net-tools-2.0-0.22.20131004git.el7.x86_64                                                                                    42/45 
  Installing : 3:mariadb-server-10.1.20-1.el7.x86_64                                                                                        43/45 
  Installing : python2-PyMySQL-0.7.11-1.el7.noarch                                                                                          44/45 
  Cleanup    : 1:mariadb-libs-5.5.56-2.el7.x86_64                                                                                           45/45 
  Verifying  : perl-HTTP-Tiny-0.033-3.el7.noarch                                                                                             1/45 
  Verifying  : 3:mariadb-config-10.1.20-1.el7.x86_64                                                                                         2/45 
  Verifying  : net-tools-2.0-0.22.20131004git.el7.x86_64                                                                                     3/45 
  Verifying  : perl-threads-shared-1.43-6.el7.x86_64                                                                                         4/45 
  Verifying  : 4:perl-Time-HiRes-1.9725-3.el7.x86_64                                                                                         5/45 
  Verifying  : perl-threads-1.87-4.el7.x86_64                                                                                                6/45 
  Verifying  : perl-Exporter-5.68-3.el7.noarch                                                                                               7/45 
  Verifying  : perl-constant-1.27-2.el7.noarch                                                                                               8/45 
  Verifying  : perl-PathTools-3.40-5.el7.x86_64                                                                                              9/45 
  Verifying  : 4:perl-macros-5.16.3-292.el7.x86_64                                                                                          10/45 
  Verifying  : perl-Compress-Raw-Bzip2-2.061-3.el7.x86_64                                                                                   11/45 
  Verifying  : 1:perl-parent-0.225-244.el7.noarch                                                                                           12/45 
  Verifying  : perl-Net-Daemon-0.48-5.el7.noarch                                                                                            13/45 
  Verifying  : 4:perl-5.16.3-292.el7.x86_64                                                                                                 14/45 
  Verifying  : 3:mariadb-libs-10.1.20-1.el7.x86_64                                                                                          15/45 
  Verifying  : perl-File-Temp-0.23.01-3.el7.noarch                                                                                          16/45 
  Verifying  : 1:perl-Pod-Simple-3.28-4.el7.noarch                                                                                          17/45 
  Verifying  : 3:mariadb-10.1.20-1.el7.x86_64                                                                                               18/45 
  Verifying  : perl-Time-Local-1.2300-2.el7.noarch                                                                                          19/45 
  Verifying  : 4:perl-libs-5.16.3-292.el7.x86_64                                                                                            20/45 
  Verifying  : perl-Pod-Perldoc-3.20-4.el7.noarch                                                                                           21/45 
  Verifying  : perl-DBI-1.627-4.el7.x86_64                                                                                                  22/45 
  Verifying  : perl-Socket-2.010-4.el7.x86_64                                                                                               23/45 
  Verifying  : perl-Carp-1.26-244.el7.noarch                                                                                                24/45 
  Verifying  : perl-Data-Dumper-2.145-3.el7.x86_64                                                                                          25/45 
  Verifying  : perl-Storable-2.45-3.el7.x86_64                                                                                              26/45 
  Verifying  : python2-PyMySQL-0.7.11-1.el7.noarch                                                                                          27/45 
  Verifying  : 1:perl-Compress-Raw-Zlib-2.061-4.el7.x86_64                                                                                  28/45 
  Verifying  : 1:perl-Pod-Escapes-1.04-292.el7.noarch                                                                                       29/45 
  Verifying  : perl-PlRPC-0.2020-14.el7.noarch                                                                                              30/45 
  Verifying  : perl-IO-Compress-2.061-2.el7.noarch                                                                                          31/45 
  Verifying  : perl-Pod-Usage-1.63-3.el7.noarch                                                                                             32/45 
  Verifying  : perl-DBD-MySQL-4.023-5.el7.x86_64                                                                                            33/45 
  Verifying  : perl-Encode-2.51-7.el7.x86_64                                                                                                34/45 
  Verifying  : perl-podlators-2.5.1-3.el7.noarch                                                                                            35/45 
  Verifying  : perl-Getopt-Long-2.40-2.el7.noarch                                                                                           36/45 
  Verifying  : 3:mariadb-server-10.1.20-1.el7.x86_64                                                                                        37/45 
  Verifying  : perl-File-Path-2.09-2.el7.noarch                                                                                             38/45 
  Verifying  : lsof-4.87-4.el7.x86_64                                                                                                       39/45 
  Verifying  : 3:mariadb-errmsg-10.1.20-1.el7.x86_64                                                                                        40/45 
  Verifying  : perl-Filter-1.49-3.el7.x86_64                                                                                                41/45 
  Verifying  : perl-Scalar-List-Utils-1.27-248.el7.x86_64                                                                                   42/45 
  Verifying  : perl-Text-ParseWords-3.29-4.el7.noarch                                                                                       43/45 
  Verifying  : 3:mariadb-common-10.1.20-1.el7.x86_64                                                                                        44/45 
  Verifying  : 1:mariadb-libs-5.5.56-2.el7.x86_64                                                                                           45/45 

Installed:
  mariadb.x86_64 3:10.1.20-1.el7             mariadb-server.x86_64 3:10.1.20-1.el7             python2-PyMySQL.noarch 0:0.7.11-1.el7            

Dependency Installed:
  lsof.x86_64 0:4.87-4.el7                        mariadb-common.x86_64 3:10.1.20-1.el7           mariadb-config.x86_64 3:10.1.20-1.el7         
  mariadb-errmsg.x86_64 3:10.1.20-1.el7           net-tools.x86_64 0:2.0-0.22.20131004git.el7     perl.x86_64 4:5.16.3-292.el7                  
  perl-Carp.noarch 0:1.26-244.el7                 perl-Compress-Raw-Bzip2.x86_64 0:2.061-3.el7    perl-Compress-Raw-Zlib.x86_64 1:2.061-4.el7   
  perl-DBD-MySQL.x86_64 0:4.023-5.el7             perl-DBI.x86_64 0:1.627-4.el7                   perl-Data-Dumper.x86_64 0:2.145-3.el7         
  perl-Encode.x86_64 0:2.51-7.el7                 perl-Exporter.noarch 0:5.68-3.el7               perl-File-Path.noarch 0:2.09-2.el7            
  perl-File-Temp.noarch 0:0.23.01-3.el7           perl-Filter.x86_64 0:1.49-3.el7                 perl-Getopt-Long.noarch 0:2.40-2.el7          
  perl-HTTP-Tiny.noarch 0:0.033-3.el7             perl-IO-Compress.noarch 0:2.061-2.el7           perl-Net-Daemon.noarch 0:0.48-5.el7           
  perl-PathTools.x86_64 0:3.40-5.el7              perl-PlRPC.noarch 0:0.2020-14.el7               perl-Pod-Escapes.noarch 1:1.04-292.el7        
  perl-Pod-Perldoc.noarch 0:3.20-4.el7            perl-Pod-Simple.noarch 1:3.28-4.el7             perl-Pod-Usage.noarch 0:1.63-3.el7            
  perl-Scalar-List-Utils.x86_64 0:1.27-248.el7    perl-Socket.x86_64 0:2.010-4.el7                perl-Storable.x86_64 0:2.45-3.el7             
  perl-Text-ParseWords.noarch 0:3.29-4.el7        perl-Time-HiRes.x86_64 4:1.9725-3.el7           perl-Time-Local.noarch 0:1.2300-2.el7         
  perl-constant.noarch 0:1.27-2.el7               perl-libs.x86_64 4:5.16.3-292.el7               perl-macros.x86_64 4:5.16.3-292.el7           
  perl-parent.noarch 1:0.225-244.el7              perl-podlators.noarch 0:2.5.1-3.el7             perl-threads.x86_64 0:1.87-4.el7              
  perl-threads-shared.x86_64 0:1.43-6.el7        

Dependency Updated:
  mariadb-libs.x86_64 3:10.1.20-1.el7                                                                                                             

Complete!
```

Configure
```
[vagrant@localhost ~]$ IPADDR=10.64.33.64 && echo -e "[mysqld]\nbind-address = $IPADDR\n\ndefault-storage-engine = innodb\ninnodb_file_per_table = on\nmax_connections = 4096\ncollation-server = utf8_general_ci\ncharacter-set-server = utf8" | sudo tee /etc/my.cnf.d/openstack.cnf
[mysqld]
bind-address = 10.64.33.64

default-storage-engine = innodb
innodb_file_per_table = on
max_connections = 4096
collation-server = utf8_general_ci
character-set-server = utf8
```

```
[vagrant@localhost ~]$ systemctl is-active mariadb.service
unknown
```

```
[vagrant@localhost ~]$ sudo systemctl start mariadb.service
```

```
[vagrant@localhost ~]$ systemctl is-active mariadb.service
active
```

```
[vagrant@localhost ~]$ systemctl -l status --no-pager  mariadb.service
● mariadb.service - MariaDB 10.1 database server
   Loaded: loaded (/usr/lib/systemd/system/mariadb.service; disabled; vendor preset: disabled)
   Active: active (running) since Fri 2017-10-20 19:41:13 UTC; 1min 7s ago
  Process: 22927 ExecStartPost=/usr/libexec/mysql-check-upgrade (code=exited, status=0/SUCCESS)
  Process: 22740 ExecStartPre=/usr/libexec/mysql-prepare-db-dir %n (code=exited, status=0/SUCCESS)
  Process: 22718 ExecStartPre=/usr/libexec/mysql-check-socket (code=exited, status=0/SUCCESS)
 Main PID: 22899 (mysqld)
   Status: "Taking your SQL requests now..."
   CGroup: /system.slice/mariadb.service
           └─22899 /usr/libexec/mysqld --basedir=/usr
```

```
[vagrant@localhost ~]$ sudo systemctl enable mariadb.service
Created symlink from /etc/systemd/system/multi-user.target.wants/mariadb.service to /usr/lib/systemd/system/mariadb.service.
```

Test
```
[vagrant@localhost ~]$ mysql -u root -e "show databases;"
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| test               |
+--------------------+
[vagrant@localhost ~]$ mysql -u root -e "show tables in test;"
[vagrant@localhost ~]$ mysql -u root -e "CREATE TABLE test.test(test varchar(20));"
[vagrant@localhost ~]$ mysql -u root -e "show tables in test;"
+----------------+
| Tables_in_test |
+----------------+
| test           |
+----------------+
[vagrant@localhost ~]$ mysql -u root -e "INSERT test.test VALUES ('test');"
[vagrant@localhost ~]$ mysql -u root -e "SELECT * FROM test.test;"
+------+
| test |
+------+
| test |
+------+
```

```
[vagrant@localhost ~]$ mysql -u root -e "select host, user from mysql.user;"
+-----------------------+------+
| host                  | user |
+-----------------------+------+
| 127.0.0.1             | root |
| ::1                   | root |
| localhost             |      |
| localhost             | root |
| localhost.localdomain |      |
| localhost.localdomain | root |
+-----------------------+------+
```

Password for _root_
```
mysql -u root -e "UPDATE mysql.user SET Password=PASSWORD('$MARIADB_PASS') WHERE User='root';"
```

### Queue

[Rabbit](https://docs.openstack.org/install-guide/environment-messaging-rdo.html)
```
[vagrant@localhost ~]$ yum search --verbose rabbitmq-server
Loading "fastestmirror" plugin
Config time: 0.004
Yum version: 3.4.3
Setting up Package Sacks
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
pkgsack time: 0.008
rpmdb time: 0.000
tags time: 0.000
========================================================== N/S matched: rabbitmq-server ==========================================================
rabbitmq-server.noarch : The RabbitMQ server
Repo        : centos-openstack-pike




  Name and summary matches only, use "search all" for everything.
```

```
[vagrant@localhost ~]$ sudo yum install -y rabbitmq-server
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
Resolving Dependencies
--> Running transaction check
---> Package rabbitmq-server.noarch 0:3.6.5-1.el7 will be installed
--> Processing Dependency: erlang-xmerl >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-tools >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-stdlib >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-ssl >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-sasl >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-public_key >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-os_mon >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-mnesia >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-kernel >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-erts >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-eldap >= R12B-3 for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Processing Dependency: erlang-sd_notify for package: rabbitmq-server-3.6.5-1.el7.noarch
--> Running transaction check
---> Package erlang-eldap.x86_64 0:18.3.4.4-2.el7 will be installed
--> Processing Dependency: erlang-asn1(x86-64) = 18.3.4.4-2.el7 for package: erlang-eldap-18.3.4.4-2.el7.x86_64
---> Package erlang-erts.x86_64 0:18.3.4.4-2.el7 will be installed
--> Processing Dependency: lksctp-tools for package: erlang-erts-18.3.4.4-2.el7.x86_64
---> Package erlang-kernel.x86_64 0:18.3.4.4-2.el7 will be installed
---> Package erlang-mnesia.x86_64 0:18.3.4.4-2.el7 will be installed
---> Package erlang-os_mon.x86_64 0:18.3.4.4-2.el7 will be installed
--> Processing Dependency: erlang-snmp(x86-64) = 18.3.4.4-2.el7 for package: erlang-os_mon-18.3.4.4-2.el7.x86_64
--> Processing Dependency: erlang-otp_mibs(x86-64) = 18.3.4.4-2.el7 for package: erlang-os_mon-18.3.4.4-2.el7.x86_64
---> Package erlang-public_key.x86_64 0:18.3.4.4-2.el7 will be installed
--> Processing Dependency: erlang-crypto(x86-64) = 18.3.4.4-2.el7 for package: erlang-public_key-18.3.4.4-2.el7.x86_64
---> Package erlang-sasl.x86_64 0:18.3.4.4-2.el7 will be installed
---> Package erlang-sd_notify.x86_64 0:0.1-9.el7 will be installed
---> Package erlang-ssl.x86_64 0:18.3.4.4-2.el7 will be installed
--> Processing Dependency: erlang-inets(x86-64) = 18.3.4.4-2.el7 for package: erlang-ssl-18.3.4.4-2.el7.x86_64
---> Package erlang-stdlib.x86_64 0:18.3.4.4-2.el7 will be installed
--> Processing Dependency: erlang-compiler(x86-64) = 18.3.4.4-2.el7 for package: erlang-stdlib-18.3.4.4-2.el7.x86_64
---> Package erlang-tools.x86_64 0:18.3.4.4-2.el7 will be installed
--> Processing Dependency: erlang-runtime_tools(x86-64) = 18.3.4.4-2.el7 for package: erlang-tools-18.3.4.4-2.el7.x86_64
---> Package erlang-xmerl.x86_64 0:18.3.4.4-2.el7 will be installed
--> Running transaction check
---> Package erlang-asn1.x86_64 0:18.3.4.4-2.el7 will be installed
---> Package erlang-compiler.x86_64 0:18.3.4.4-2.el7 will be installed
--> Processing Dependency: erlang-hipe(x86-64) = 18.3.4.4-2.el7 for package: erlang-compiler-18.3.4.4-2.el7.x86_64
---> Package erlang-crypto.x86_64 0:18.3.4.4-2.el7 will be installed
---> Package erlang-inets.x86_64 0:18.3.4.4-2.el7 will be installed
---> Package erlang-otp_mibs.x86_64 0:18.3.4.4-2.el7 will be installed
---> Package erlang-runtime_tools.x86_64 0:18.3.4.4-2.el7 will be installed
---> Package erlang-snmp.x86_64 0:18.3.4.4-2.el7 will be installed
---> Package lksctp-tools.x86_64 0:1.0.17-2.el7 will be installed
--> Running transaction check
---> Package erlang-hipe.x86_64 0:18.3.4.4-2.el7 will be installed
--> Processing Dependency: erlang-syntax_tools(x86-64) = 18.3.4.4-2.el7 for package: erlang-hipe-18.3.4.4-2.el7.x86_64
--> Running transaction check
---> Package erlang-syntax_tools.x86_64 0:18.3.4.4-2.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                                Arch                     Version                            Repository                               Size
==================================================================================================================================================
Installing:
 rabbitmq-server                        noarch                   3.6.5-1.el7                        centos-openstack-pike                   5.1 M
Installing for dependencies:
 erlang-asn1                            x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   736 k
 erlang-compiler                        x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   1.1 M
 erlang-crypto                          x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   129 k
 erlang-eldap                           x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   117 k
 erlang-erts                            x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   2.8 M
 erlang-hipe                            x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   2.6 M
 erlang-inets                           x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   786 k
 erlang-kernel                          x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   1.1 M
 erlang-mnesia                          x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   712 k
 erlang-os_mon                          x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   115 k
 erlang-otp_mibs                        x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                    35 k
 erlang-public_key                      x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   571 k
 erlang-runtime_tools                   x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   191 k
 erlang-sasl                            x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   298 k
 erlang-sd_notify                       x86_64                   0.1-9.el7                          centos-openstack-pike                   8.8 k
 erlang-snmp                            x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   1.6 M
 erlang-ssl                             x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   702 k
 erlang-stdlib                          x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   2.4 M
 erlang-syntax_tools                    x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   382 k
 erlang-tools                           x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   584 k
 erlang-xmerl                           x86_64                   18.3.4.4-2.el7                     centos-openstack-pike                   993 k
 lksctp-tools                           x86_64                   1.0.17-2.el7                       base                                     88 k

Transaction Summary
==================================================================================================================================================
Install  1 Package (+22 Dependent packages)

Total download size: 23 M
Installed size: 39 M
Downloading packages:
(1/23): erlang-asn1-18.3.4.4-2.el7.x86_64.rpm                                                                              | 736 kB  00:00:03     
(2/23): erlang-crypto-18.3.4.4-2.el7.x86_64.rpm                                                                            | 129 kB  00:00:00     
(3/23): erlang-eldap-18.3.4.4-2.el7.x86_64.rpm                                                                             | 117 kB  00:00:01     
(4/23): erlang-compiler-18.3.4.4-2.el7.x86_64.rpm                                                                          | 1.1 MB  00:00:06     
(5/23): erlang-hipe-18.3.4.4-2.el7.x86_64.rpm                                                                              | 2.6 MB  00:00:08     
(6/23): erlang-erts-18.3.4.4-2.el7.x86_64.rpm                                                                              | 2.8 MB  00:00:09     
(7/23): erlang-inets-18.3.4.4-2.el7.x86_64.rpm                                                                             | 786 kB  00:00:02     
(8/23): erlang-mnesia-18.3.4.4-2.el7.x86_64.rpm                                                                            | 712 kB  00:00:03     
(9/23): erlang-kernel-18.3.4.4-2.el7.x86_64.rpm                                                                            | 1.1 MB  00:00:05     
(10/23): erlang-otp_mibs-18.3.4.4-2.el7.x86_64.rpm                                                                         |  35 kB  00:00:00     
(11/23): erlang-os_mon-18.3.4.4-2.el7.x86_64.rpm                                                                           | 115 kB  00:00:00     
(12/23): erlang-public_key-18.3.4.4-2.el7.x86_64.rpm                                                                       | 571 kB  00:00:02     
(13/23): erlang-runtime_tools-18.3.4.4-2.el7.x86_64.rpm                                                                    | 191 kB  00:00:03     
(14/23): erlang-sd_notify-0.1-9.el7.x86_64.rpm                                                                             | 8.8 kB  00:00:00     
(15/23): erlang-sasl-18.3.4.4-2.el7.x86_64.rpm                                                                             | 298 kB  00:00:01     
(16/23): erlang-ssl-18.3.4.4-2.el7.x86_64.rpm                                                                              | 702 kB  00:00:02     
(17/23): erlang-stdlib-18.3.4.4-2.el7.x86_64.rpm                                                                           | 2.4 MB  00:00:07     
(18/23): erlang-snmp-18.3.4.4-2.el7.x86_64.rpm                                                                             | 1.6 MB  00:00:10     
(19/23): erlang-syntax_tools-18.3.4.4-2.el7.x86_64.rpm                                                                     | 382 kB  00:00:01     
(20/23): lksctp-tools-1.0.17-2.el7.x86_64.rpm                                                                              |  88 kB  00:00:00     
(21/23): erlang-tools-18.3.4.4-2.el7.x86_64.rpm                                                                            | 584 kB  00:00:02     
(22/23): erlang-xmerl-18.3.4.4-2.el7.x86_64.rpm                                                                            | 993 kB  00:00:03     
(23/23): rabbitmq-server-3.6.5-1.el7.noarch.rpm                                                                            | 5.1 MB  00:00:12     
--------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                             464 kB/s |  23 MB  00:00:50     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : lksctp-tools-1.0.17-2.el7.x86_64                                                                                              1/23 
  Installing : erlang-kernel-18.3.4.4-2.el7.x86_64                                                                                           2/23 
  Installing : erlang-crypto-18.3.4.4-2.el7.x86_64                                                                                           3/23 
  Installing : erlang-erts-18.3.4.4-2.el7.x86_64                                                                                             4/23 
  Installing : erlang-stdlib-18.3.4.4-2.el7.x86_64                                                                                           5/23 
  Installing : erlang-syntax_tools-18.3.4.4-2.el7.x86_64                                                                                     6/23 
  Installing : erlang-compiler-18.3.4.4-2.el7.x86_64                                                                                         7/23 
  Installing : erlang-hipe-18.3.4.4-2.el7.x86_64                                                                                             8/23 
  Installing : erlang-mnesia-18.3.4.4-2.el7.x86_64                                                                                           9/23 
  Installing : erlang-runtime_tools-18.3.4.4-2.el7.x86_64                                                                                   10/23 
  Installing : erlang-snmp-18.3.4.4-2.el7.x86_64                                                                                            11/23 
  Installing : erlang-asn1-18.3.4.4-2.el7.x86_64                                                                                            12/23 
  Installing : erlang-public_key-18.3.4.4-2.el7.x86_64                                                                                      13/23 
  Installing : erlang-ssl-18.3.4.4-2.el7.x86_64                                                                                             14/23 
  Installing : erlang-inets-18.3.4.4-2.el7.x86_64                                                                                           15/23 
  Installing : erlang-tools-18.3.4.4-2.el7.x86_64                                                                                           16/23 
  Installing : erlang-sasl-18.3.4.4-2.el7.x86_64                                                                                            17/23 
  Installing : erlang-eldap-18.3.4.4-2.el7.x86_64                                                                                           18/23 
  Installing : erlang-otp_mibs-18.3.4.4-2.el7.x86_64                                                                                        19/23 
  Installing : erlang-os_mon-18.3.4.4-2.el7.x86_64                                                                                          20/23 
  Installing : erlang-xmerl-18.3.4.4-2.el7.x86_64                                                                                           21/23 
  Installing : erlang-sd_notify-0.1-9.el7.x86_64                                                                                            22/23 
  Installing : rabbitmq-server-3.6.5-1.el7.noarch                                                                                           23/23 
  Verifying  : erlang-eldap-18.3.4.4-2.el7.x86_64                                                                                            1/23 
  Verifying  : rabbitmq-server-3.6.5-1.el7.noarch                                                                                            2/23 
  Verifying  : erlang-mnesia-18.3.4.4-2.el7.x86_64                                                                                           3/23 
  Verifying  : erlang-hipe-18.3.4.4-2.el7.x86_64                                                                                             4/23 
  Verifying  : erlang-kernel-18.3.4.4-2.el7.x86_64                                                                                           5/23 
  Verifying  : erlang-otp_mibs-18.3.4.4-2.el7.x86_64                                                                                         6/23 
  Verifying  : erlang-crypto-18.3.4.4-2.el7.x86_64                                                                                           7/23 
  Verifying  : erlang-xmerl-18.3.4.4-2.el7.x86_64                                                                                            8/23 
  Verifying  : erlang-stdlib-18.3.4.4-2.el7.x86_64                                                                                           9/23 
  Verifying  : erlang-sasl-18.3.4.4-2.el7.x86_64                                                                                            10/23 
  Verifying  : lksctp-tools-1.0.17-2.el7.x86_64                                                                                             11/23 
  Verifying  : erlang-tools-18.3.4.4-2.el7.x86_64                                                                                           12/23 
  Verifying  : erlang-snmp-18.3.4.4-2.el7.x86_64                                                                                            13/23 
  Verifying  : erlang-sd_notify-0.1-9.el7.x86_64                                                                                            14/23 
  Verifying  : erlang-compiler-18.3.4.4-2.el7.x86_64                                                                                        15/23 
  Verifying  : erlang-ssl-18.3.4.4-2.el7.x86_64                                                                                             16/23 
  Verifying  : erlang-erts-18.3.4.4-2.el7.x86_64                                                                                            17/23 
  Verifying  : erlang-public_key-18.3.4.4-2.el7.x86_64                                                                                      18/23 
  Verifying  : erlang-runtime_tools-18.3.4.4-2.el7.x86_64                                                                                   19/23 
  Verifying  : erlang-syntax_tools-18.3.4.4-2.el7.x86_64                                                                                    20/23 
  Verifying  : erlang-os_mon-18.3.4.4-2.el7.x86_64                                                                                          21/23 
  Verifying  : erlang-asn1-18.3.4.4-2.el7.x86_64                                                                                            22/23 
  Verifying  : erlang-inets-18.3.4.4-2.el7.x86_64                                                                                           23/23 

Installed:
  rabbitmq-server.noarch 0:3.6.5-1.el7                                                                                                            

Dependency Installed:
  erlang-asn1.x86_64 0:18.3.4.4-2.el7                erlang-compiler.x86_64 0:18.3.4.4-2.el7       erlang-crypto.x86_64 0:18.3.4.4-2.el7          
  erlang-eldap.x86_64 0:18.3.4.4-2.el7               erlang-erts.x86_64 0:18.3.4.4-2.el7           erlang-hipe.x86_64 0:18.3.4.4-2.el7            
  erlang-inets.x86_64 0:18.3.4.4-2.el7               erlang-kernel.x86_64 0:18.3.4.4-2.el7         erlang-mnesia.x86_64 0:18.3.4.4-2.el7          
  erlang-os_mon.x86_64 0:18.3.4.4-2.el7              erlang-otp_mibs.x86_64 0:18.3.4.4-2.el7       erlang-public_key.x86_64 0:18.3.4.4-2.el7      
  erlang-runtime_tools.x86_64 0:18.3.4.4-2.el7       erlang-sasl.x86_64 0:18.3.4.4-2.el7           erlang-sd_notify.x86_64 0:0.1-9.el7            
  erlang-snmp.x86_64 0:18.3.4.4-2.el7                erlang-ssl.x86_64 0:18.3.4.4-2.el7            erlang-stdlib.x86_64 0:18.3.4.4-2.el7          
  erlang-syntax_tools.x86_64 0:18.3.4.4-2.el7        erlang-tools.x86_64 0:18.3.4.4-2.el7          erlang-xmerl.x86_64 0:18.3.4.4-2.el7           
  lksctp-tools.x86_64 0:1.0.17-2.el7                

Complete!
```

Start
```
[vagrant@localhost ~]$ [ 'unknown' == $(systemctl is-active rabbitmq-server.service) ] && sudo systemctl start rabbitmq-server.service
```

Auto-Start
```
[vagrant@localhost ~]$ [ 'active' == $(systemctl is-active rabbitmq-server.service) ] && sudo systemctl enable rabbitmq-server.service
Created symlink from /etc/systemd/system/multi-user.target.wants/rabbitmq-server.service to /usr/lib/systemd/system/rabbitmq-server.service.
```

```
[vagrant@localhost ~]$ systemctl -l status --no-pager rabbitmq-server.service
● rabbitmq-server.service - RabbitMQ broker
   Loaded: loaded (/usr/lib/systemd/system/rabbitmq-server.service; disabled; vendor preset: disabled)
   Active: active (running) since Fri 2017-10-20 20:22:22 UTC; 43s ago
 Main PID: 23134 (beam)
   Status: "Initialized"
   CGroup: /system.slice/rabbitmq-server.service
           ├─23134 /usr/lib64/erlang/erts-7.3.1.2/bin/beam -W w -A 64 -P 1048576 -t 5000000 -stbt db -K true -- -root /usr/lib64/erlang -progname erl -- -home /var/lib/rabbitmq -- -pa /usr/lib/rabbitmq/lib/rabbitmq_server-3.6.5/ebin -noshell -noinput -s rabbit boot -sname rabbit@localhost -boot start_sasl -config /etc/rabbitmq/rabbitmq -kernel inet_default_connect_options [{nodelay,true}] -sasl errlog_type error -sasl sasl_error_logger false -rabbit error_logger {file,"/var/log/rabbitmq/rabbit@localhost.log"} -rabbit sasl_error_logger {file,"/var/log/rabbitmq/rabbit@localhost-sasl.log"} -rabbit enabled_plugins_file "/etc/rabbitmq/enabled_plugins" -rabbit plugins_dir "/usr/lib/rabbitmq/lib/rabbitmq_server-3.6.5/plugins" -rabbit plugins_expand_dir "/var/lib/rabbitmq/mnesia/rabbit@localhost-plugins-expand" -os_mon start_cpu_sup false -os_mon start_disksup false -os_mon start_memsup false -mnesia dir "/var/lib/rabbitmq/mnesia/rabbit@localhost" -kernel inet_dist_listen_min 25672 -kernel inet_dist_listen_max 25672
           ├─23310 inet_gethost 4
           └─23311 inet_gethost 4

Oct 20 20:22:21 localhost.localdomain systemd[1]: Starting RabbitMQ broker...
Oct 20 20:22:21 localhost.localdomain rabbitmq-server[23134]: RabbitMQ 3.6.5. Copyright (C) 2007-2016 Pivotal Software, Inc.
Oct 20 20:22:21 localhost.localdomain rabbitmq-server[23134]: ##  ##      Licensed under the MPL.  See http://www.rabbitmq.com/
Oct 20 20:22:21 localhost.localdomain rabbitmq-server[23134]: ##  ##
Oct 20 20:22:21 localhost.localdomain rabbitmq-server[23134]: ##########  Logs: /var/log/rabbitmq/rabbit@localhost.log
Oct 20 20:22:21 localhost.localdomain rabbitmq-server[23134]: ######  ##        /var/log/rabbitmq/rabbit@localhost-sasl.log
Oct 20 20:22:21 localhost.localdomain rabbitmq-server[23134]: ##########
Oct 20 20:22:21 localhost.localdomain rabbitmq-server[23134]: Starting broker...
Oct 20 20:22:22 localhost.localdomain systemd[1]: Started RabbitMQ broker.
Oct 20 20:22:22 localhost.localdomain rabbitmq-server[23134]: completed with 0 plugins.
```

```
[vagrant@localhost ~]$ sudo rabbitmqctl status
Status of node rabbit@localhost ...
[{pid,23134},
 {running_applications,[{rabbit,"RabbitMQ","3.6.5"},
                        {mnesia,"MNESIA  CXC 138 12","4.13.4"},
                        {os_mon,"CPO  CXC 138 46","2.4"},
                        {ranch,"Socket acceptor pool for TCP protocols.",
                               "1.2.1"},
                        {rabbit_common,[],"3.6.5"},
                        {xmerl,"XML parser","1.3.10"},
                        {sasl,"SASL  CXC 138 11","2.7"},
                        {stdlib,"ERTS  CXC 138 10","2.8"},
                        {kernel,"ERTS  CXC 138 10","4.2"}]},
 {os,{unix,linux}},
 {erlang_version,"Erlang/OTP 18 [erts-7.3.1.2] [source] [64-bit] [async-threads:64] [hipe] [kernel-poll:true]\n"},
 {memory,[{total,43675752},
          {connection_readers,0},
          {connection_writers,0},
          {connection_channels,0},
          {connection_other,0},
          {queue_procs,2680},
          {queue_slave_procs,0},
          {plugins,0},
          {other_proc,18491544},
          {mnesia,58168},
          {mgmt_db,0},
          {msg_index,49744},
          {other_ets,902960},
          {binary,19080},
          {code,19689812},
          {atom,719761},
          {other_system,3742003}]},
 {alarms,[]},
 {listeners,[{clustering,25672,"::"},{amqp,5672,"::"}]},
 {vm_memory_high_watermark,0.4},
 {vm_memory_limit,771506176},
 {disk_free_limit,50000000},
 {disk_free,38832476160},
 {file_descriptors,[{total_limit,924},
                    {total_used,2},
                    {sockets_limit,829},
                    {sockets_used,0}]},
 {processes,[{limit,1048576},{used,137}]},
 {run_queue,0},
 {uptime,173},
 {kernel,{net_ticktime,60}}]
```

```
[vagrant@localhost ~]$ sudo rabbitmqctl add_user openstack RABBIT_PASS && sudo rabbitmq set_permissions openstack ".*" ".*" ".*"
Creating user "openstack" ...
Setting permissions for user "openstack" in vhost "/" ...
```

```
[vagrant@localhost ~]$ sudo rabbitmqctl list_users && sudo rabbitmqctl list_user_permissions openstack
Listing users ...
openstack	[]
guest	[administrator]
Listing permissions for user "openstack" ...
/	.*	.*	.*
```

### Cache

[Memcached](https://docs.openstack.org/install-guide/environment-memcached-rdo.html)
```
[vagrant@localhost ~]$ yum list | egrep '^memcached\.'
memcached.x86_64                   1.4.39-1.el7            centos-openstack-pike
```

```
[vagrant@localhost ~]$ sudo yum install -y memcached
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
Resolving Dependencies
--> Running transaction check
---> Package memcached.x86_64 0:1.4.39-1.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                        Arch                        Version                              Repository                                  Size
==================================================================================================================================================
Installing:
 memcached                      x86_64                      1.4.39-1.el7                         centos-openstack-pike                      118 k

Transaction Summary
==================================================================================================================================================
Install  1 Package

Total download size: 118 k
Installed size: 255 k
Downloading packages:
memcached-1.4.39-1.el7.x86_64.rpm                                                                                          | 118 kB  00:00:01     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : memcached-1.4.39-1.el7.x86_64                                                                                                  1/1 
  Verifying  : memcached-1.4.39-1.el7.x86_64                                                                                                  1/1 

Installed:
  memcached.x86_64 0:1.4.39-1.el7                                                                                                                 

Complete!
```

Start
```
[vagrant@localhost ~]$ sudo systemctl start memcached.service && sudo systemctl enable memcached.service
Created symlink from /etc/systemd/system/multi-user.target.wants/memcached.service to /usr/lib/systemd/system/memcached.service.
```

```
[vagrant@localhost ~]$ systemctl -l status --no-pager memcached.service
● memcached.service - memcached daemon
   Loaded: loaded (/usr/lib/systemd/system/memcached.service; enabled; vendor preset: disabled)
   Active: active (running) since Fri 2017-10-20 20:37:22 UTC; 13s ago
 Main PID: 25037 (memcached)
   CGroup: /system.slice/memcached.service
           └─25037 /usr/bin/memcached -p 11211 -u memcached -m 64 -c 1024 -l 127.0.0.1,::1
```

