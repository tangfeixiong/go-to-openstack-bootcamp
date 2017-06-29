# Installing Openstack command client 

## Contents of table

* CentOS 7
* Fedora 23

## CentOS 7


## Feodra 23

For example
```
[vagrant@localhost ~]$ sudo dnf install python-openstackclient
google-chrome                                                                                                            897  B/s | 3.8 kB     00:04    
added from: http://dl.google.com/linux/chrome/rpm/stable/x86_64                                                          3.3 kB/s | 3.8 kB     00:01    
Yarn Repository                                                                                                          2.9 kB/s | 361 kB     02:02    
Node.js Packages for Fedora Core 23 - x86_64                                                                             7.8 kB/s | 632 kB     01:20    
上次元数据过期检查在 0:01:17 前执行于 Wed Jun 21 04:21:47 2017。
依赖关系解决。
=========================================================================================================================================================
 Package                                         架构                         版本                                   仓库                           大小
=========================================================================================================================================================
安装:
 pyOpenSSL                                       noarch                       0.15.1-1.fc23                          fedora                        186 k
 pyparsing                                       noarch                       2.1.5-1.fc23                           updates                        11 k
 python-babel                                    noarch                       1.3-8.fc23                             fedora                        2.5 M
 python-cffi                                     x86_64                       1.4.2-1.fc23                           updates                       209 k
 python-cinderclient                             noarch                       1.2.1-1.fc23                           fedora                        185 k
 python-cliff                                    noarch                       1.13.0-1.fc23                          fedora                         87 k
 python-cmd2                                     noarch                       0.6.8-8.fc23                           updates                        45 k
 python-enum34                                   noarch                       1.0.4-2.fc23                           fedora                         56 k
 python-extras                                   noarch                       0.0.3-7.fc23                           fedora                         18 k
 python-fixtures                                 noarch                       0.3.14-4.fc23                          fedora                         74 k
 python-glanceclient                             noarch                       1:0.17.0-3.fc23                        fedora                        138 k
 python-httplib2                                 noarch                       0.9.1-2.fc23                           fedora                        119 k
 python-idna                                     noarch                       2.0-1.fc23                             fedora                         97 k
 python-ipaddress                                noarch                       1.0.7-4.fc23                           fedora                         36 k
 python-jsonpatch                                noarch                       1.2-6.fc23                             fedora                         19 k
 python-jsonpointer                              noarch                       1.10-1.fc23                            updates                        17 k
 python-jsonschema                               noarch                       2.4.0-2.fc23                           fedora                         75 k
 python-keyring                                  noarch                       5.0-5.fc23                             updates                       119 k
 python-keystoneclient                           noarch                       1:1.3.0-2.fc23                         fedora                        575 k
 python-linecache2                               noarch                       1.0.0-1.fc23                           fedora                         16 k
 python-mimeparse                                noarch                       0.1.4-6.fc23                           updates                        14 k
 python-netaddr                                  noarch                       0.7.18-1.fc23                          fedora                        1.3 M
 python-netifaces                                x86_64                       0.10.4-2.fc23                          fedora                         22 k
 python-neutronclient                            noarch                       2.4.0-2.fc23                           fedora                        178 k
 python-novaclient                               noarch                       1:2.23.3-1.fc23                        updates                       245 k
 python-openstackclient                          noarch                       1.0.3-3.fc23                           fedora                        213 k
 python-oslo-config                              noarch                       2:1.9.3-2.fc23                         fedora                        130 k
 python-oslo-i18n                                noarch                       1.5.0-4.fc23                           fedora                         58 k
 python-oslo-serialization                       noarch                       1.4.0-2.fc23                           fedora                         30 k
 python-oslo-utils                               noarch                       1.4.0-2.fc23                           fedora                         90 k
 python-ply                                      noarch                       3.6-2.fc23                             fedora                        149 k
 python-prettytable                              noarch                       0.7.2-5.fc23                           fedora                         42 k
 python-pyasn1                                   noarch                       0.1.8-1.fc23                           fedora                        104 k
 python-pycparser                                noarch                       2.14-3.fc23                            fedora                        108 k
 python-simplejson                               x86_64                       3.5.3-3.fc23                           fedora                        190 k
 python-stevedore                                noarch                       1.5.0-1.fc23                           fedora                         58 k
 python-testtools                                noarch                       1.8.0-3.fc23                           updates                       305 k
 python-warlock                                  noarch                       1.0.1-3.fc23                           fedora                         19 k
 python-webob                                    noarch                       1.4.2-1.fc23                           updates                       216 k
 python2-cryptography                            x86_64                       1.5.3-3.fc23                           updates                       485 k
 python2-iso8601                                 noarch                       0.1.11-1.fc23                          updates                        23 k
 python2-msgpack                                 x86_64                       0.4.8-1.fc23                           updates                        89 k
 python2-pbr                                     noarch                       1.8.1-3.fc23                           updates                       176 k
 python2-pyparsing                               noarch                       2.1.5-1.fc23                           updates                       118 k
 python2-traceback2                              noarch                       1.4.0-7.fc23                           updates                        23 k
 python2-unittest2                               noarch                       1.1.0-5.fc23                           updates                       178 k
 pytz                                            noarch                       2015.4-1.fc23                          fedora                         60 k

事务概要
=========================================================================================================================================================
安装  47 Packages

总下载：9.1 M
安装大小：38 M
确定吗？[y/N]： y
下载软件包：
(1/47): python-cinderclient-1.2.1-1.fc23.noarch.rpm                                                                      8.7 kB/s | 185 kB     00:21    
(2/47): python-openstackclient-1.0.3-3.fc23.noarch.rpm                                                                   9.1 kB/s | 213 kB     00:23    
(3/47): python-cliff-1.13.0-1.fc23.noarch.rpm                                                                            9.1 kB/s |  87 kB     00:09    
(4/47): python-glanceclient-0.17.0-3.fc23.noarch.rpm                                                                      13 kB/s | 138 kB     00:10    
(5/47): python-neutronclient-2.4.0-2.fc23.noarch.rpm                                                                      18 kB/s | 178 kB     00:09    
(6/47): python-oslo-i18n-1.5.0-4.fc23.noarch.rpm                                                                          11 kB/s |  58 kB     00:05    
(7/47): python-oslo-serialization-1.4.0-2.fc23.noarch.rpm                                                                 10 kB/s |  30 kB     00:03    
(8/47): python-oslo-utils-1.4.0-2.fc23.noarch.rpm                                                                         13 kB/s |  90 kB     00:06    
(9/47): python-stevedore-1.5.0-1.fc23.noarch.rpm                                                                         9.9 kB/s |  58 kB     00:05    
(10/47): python-keystoneclient-1.3.0-2.fc23.noarch.rpm                                                                    17 kB/s | 575 kB     00:34    
(11/47): python-prettytable-0.7.2-5.fc23.noarch.rpm                                                                      6.0 kB/s |  42 kB     00:06    
(12/47): pytz-2015.4-1.fc23.noarch.rpm                                                                                   6.5 kB/s |  60 kB     00:09    
(13/47): python-simplejson-3.5.3-3.fc23.x86_64.rpm                                                                        11 kB/s | 190 kB     00:17    
(14/47): pyOpenSSL-0.15.1-1.fc23.noarch.rpm                                                                              9.1 kB/s | 186 kB     00:20    
(15/47): python-warlock-1.0.1-3.fc23.noarch.rpm                                                                          7.9 kB/s |  19 kB     00:02    
(16/47): python-httplib2-0.9.1-2.fc23.noarch.rpm                                                                         9.4 kB/s | 119 kB     00:12    
(17/47): python-oslo-config-1.9.3-2.fc23.noarch.rpm                                                                       19 kB/s | 130 kB     00:06    
(18/47): python-fixtures-0.3.14-4.fc23.noarch.rpm                                                                        8.9 kB/s |  74 kB     00:08    
(19/47): python-netifaces-0.10.4-2.fc23.x86_64.rpm                                                                       4.2 kB/s |  22 kB     00:05    
(20/47): python-jsonpatch-1.2-6.fc23.noarch.rpm                                                                          6.6 kB/s |  19 kB     00:02    
(21/47): python-jsonschema-2.4.0-2.fc23.noarch.rpm                                                                       7.2 kB/s |  75 kB     00:10    
(22/47): python2-pbr-1.8.1-3.fc23.noarch.rpm                                                                             6.8 kB/s | 176 kB     00:25    
(23/47): python2-iso8601-0.1.11-1.fc23.noarch.rpm                                                                        6.1 kB/s |  23 kB     00:03    
(24/47): python2-msgpack-0.4.8-1.fc23.x86_64.rpm                                                                          10 kB/s |  89 kB     00:08    
(25/47): python-babel-1.3-8.fc23.noarch.rpm                                                                               14 kB/s | 2.5 MB     03:02    
(26/47): pyparsing-2.1.5-1.fc23.noarch.rpm                                                                               4.6 kB/s |  11 kB     00:02    
(27/47): python-novaclient-2.23.3-1.fc23.noarch.rpm                                                                       15 kB/s | 245 kB     00:16    
(28/47): python-netaddr-0.7.18-1.fc23.noarch.rpm                                                                          14 kB/s | 1.3 MB     01:35    
(29/47): python2-pyparsing-2.1.5-1.fc23.noarch.rpm                                                                        12 kB/s | 118 kB     00:09    
(30/47): python-jsonpointer-1.10-1.fc23.noarch.rpm                                                                       4.2 kB/s |  17 kB     00:04    
(31/47): python-webob-1.4.2-1.fc23.noarch.rpm                                                                             13 kB/s | 216 kB     00:16    
(32/47): python-keyring-5.0-5.fc23.noarch.rpm                                                                            6.7 kB/s | 119 kB     00:17    
(33/47): python-extras-0.0.3-7.fc23.noarch.rpm                                                                           6.7 kB/s |  18 kB     00:02    
(34/47): python-cmd2-0.6.8-8.fc23.noarch.rpm                                                                             1.5 kB/s |  45 kB     00:30    
(35/47): python-testtools-1.8.0-3.fc23.noarch.rpm                                                                        1.8 kB/s | 305 kB     02:46    
(36/47): python-enum34-1.0.4-2.fc23.noarch.rpm                                                                           8.4 kB/s |  56 kB     00:06    
(37/47): python-idna-2.0-1.fc23.noarch.rpm                                                                               4.1 kB/s |  97 kB     00:23    
(38/47): python2-unittest2-1.1.0-5.fc23.noarch.rpm                                                                       958  B/s | 178 kB     03:10    
(39/47): python-ipaddress-1.0.7-4.fc23.noarch.rpm                                                                        8.8 kB/s |  36 kB     00:04    
(40/47): python-pyasn1-0.1.8-1.fc23.noarch.rpm                                                                            15 kB/s | 104 kB     00:07    
[MIRROR] python2-cryptography-1.5.3-3.fc23.x86_64.rpm: Curl error (18): Transferred a partial file for https://ftp.yz.yamagata-u.ac.jp/pub/linux/fedora-projects/fedora/linux/updates/23/x86_64/p/python2-cryptography-1.5.3-3.fc23.x86_64.rpm [transfer closed with 136170 bytes remaining to read]
(41/47): python-pycparser-2.14-3.fc23.noarch.rpm                                                                         9.0 kB/s | 108 kB     00:12    
(42/47): python-cffi-1.4.2-1.fc23.x86_64.rpm                                                                             9.2 kB/s | 209 kB     00:22    
(43/47): python2-traceback2-1.4.0-7.fc23.noarch.rpm                                                                      6.8 kB/s |  23 kB     00:03    
(44/47): python-linecache2-1.0.0-1.fc23.noarch.rpm                                                                       6.6 kB/s |  16 kB     00:02    
(45/47): python-mimeparse-0.1.4-6.fc23.noarch.rpm                                                                         15 kB/s |  14 kB     00:00    
(46/47): python-ply-3.6-2.fc23.noarch.rpm                                                                                9.3 kB/s | 149 kB     00:16    
(47/47): python2-cryptography-1.5.3-3.fc23.x86_64.rpm                                                                    2.2 kB/s | 485 kB     03:40    
---------------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                                      20 kB/s | 9.1 MB     07:49     
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: python2-pbr-1.8.1-3.fc23.noarch                                                                                                             1/47 
  安装: python-prettytable-0.7.2-5.fc23.noarch                                                                                                      2/47 
  安装: python-stevedore-1.5.0-1.fc23.noarch                                                                                                        3/47 
  安装: python2-iso8601-0.1.11-1.fc23.noarch                                                                                                        4/47 
  安装: python-netaddr-0.7.18-1.fc23.noarch                                                                                                         5/47 
  安装: python-simplejson-3.5.3-3.fc23.x86_64                                                                                                       6/47 
  安装: python-oslo-config-2:1.9.3-2.fc23.noarch                                                                                                    7/47 
  安装: python-keyring-5.0-5.fc23.noarch                                                                                                            8/47 
  安装: python-netifaces-0.10.4-2.fc23.x86_64                                                                                                       9/47 
  安装: pytz-2015.4-1.fc23.noarch                                                                                                                  10/47 
  安装: python-babel-1.3-8.fc23.noarch                                                                                                             11/47 
  安装: python-mimeparse-0.1.4-6.fc23.noarch                                                                                                       12/47 
  安装: python-linecache2-1.0.0-1.fc23.noarch                                                                                                      13/47 
  安装: python2-traceback2-1.4.0-7.fc23.noarch                                                                                                     14/47 
  安装: python2-unittest2-1.1.0-5.fc23.noarch                                                                                                      15/47 
  安装: python-ply-3.6-2.fc23.noarch                                                                                                               16/47 
  安装: python-pycparser-2.14-3.fc23.noarch                                                                                                        17/47 
  安装: python-cffi-1.4.2-1.fc23.x86_64                                                                                                            18/47 
  安装: python-pyasn1-0.1.8-1.fc23.noarch                                                                                                          19/47 
  安装: python-ipaddress-1.0.7-4.fc23.noarch                                                                                                       20/47 
  安装: python-idna-2.0-1.fc23.noarch                                                                                                              21/47 
  安装: python-enum34-1.0.4-2.fc23.noarch                                                                                                          22/47 
  安装: python2-cryptography-1.5.3-3.fc23.x86_64                                                                                                   23/47 
  安装: pyOpenSSL-0.15.1-1.fc23.noarch                                                                                                             24/47 
  安装: python-extras-0.0.3-7.fc23.noarch                                                                                                          25/47 
  安装: python-testtools-1.8.0-3.fc23.noarch                                                                                                       26/47 
  安装: python-fixtures-0.3.14-4.fc23.noarch                                                                                                       27/47 
  安装: python-oslo-i18n-1.5.0-4.fc23.noarch                                                                                                       28/47 
  安装: python-oslo-utils-1.4.0-2.fc23.noarch                                                                                                      29/47 
  安装: python-jsonpointer-1.10-1.fc23.noarch                                                                                                      30/47 
  安装: python-jsonpatch-1.2-6.fc23.noarch                                                                                                         31/47 
  安装: python-webob-1.4.2-1.fc23.noarch                                                                                                           32/47 
  安装: python2-pyparsing-2.1.5-1.fc23.noarch                                                                                                      33/47 
  安装: pyparsing-2.1.5-1.fc23.noarch                                                                                                              34/47 
  安装: python-cmd2-0.6.8-8.fc23.noarch                                                                                                            35/47 
  安装: python-cliff-1.13.0-1.fc23.noarch                                                                                                          36/47 
  安装: python2-msgpack-0.4.8-1.fc23.x86_64                                                                                                        37/47 
  安装: python-oslo-serialization-1.4.0-2.fc23.noarch                                                                                              38/47 
  安装: python-keystoneclient-1:1.3.0-2.fc23.noarch                                                                                                39/47 
  安装: python-cinderclient-1.2.1-1.fc23.noarch                                                                                                    40/47 
  安装: python-neutronclient-2.4.0-2.fc23.noarch                                                                                                   41/47 
  安装: python-novaclient-1:2.23.3-1.fc23.noarch                                                                                                   42/47 
  安装: python-jsonschema-2.4.0-2.fc23.noarch                                                                                                      43/47 
  安装: python-warlock-1.0.1-3.fc23.noarch                                                                                                         44/47 
  安装: python-httplib2-0.9.1-2.fc23.noarch                                                                                                        45/47 
  安装: python-glanceclient-1:0.17.0-3.fc23.noarch                                                                                                 46/47 
  安装: python-openstackclient-1.0.3-3.fc23.noarch                                                                                                 47/47 
  验证: python-openstackclient-1.0.3-3.fc23.noarch                                                                                                  1/47 
  验证: python-babel-1.3-8.fc23.noarch                                                                                                              2/47 
  验证: python-cinderclient-1.2.1-1.fc23.noarch                                                                                                     3/47 
  验证: python-cliff-1.13.0-1.fc23.noarch                                                                                                           4/47 
  验证: python-glanceclient-1:0.17.0-3.fc23.noarch                                                                                                  5/47 
  验证: python-keystoneclient-1:1.3.0-2.fc23.noarch                                                                                                 6/47 
  验证: python-neutronclient-2.4.0-2.fc23.noarch                                                                                                    7/47 
  验证: python-oslo-i18n-1.5.0-4.fc23.noarch                                                                                                        8/47 
  验证: python-oslo-serialization-1.4.0-2.fc23.noarch                                                                                               9/47 
  验证: python-oslo-utils-1.4.0-2.fc23.noarch                                                                                                      10/47 
  验证: python-stevedore-1.5.0-1.fc23.noarch                                                                                                       11/47 
  验证: pytz-2015.4-1.fc23.noarch                                                                                                                  12/47 
  验证: python-prettytable-0.7.2-5.fc23.noarch                                                                                                     13/47 
  验证: python-simplejson-3.5.3-3.fc23.x86_64                                                                                                      14/47 
  验证: pyOpenSSL-0.15.1-1.fc23.noarch                                                                                                             15/47 
  验证: python-httplib2-0.9.1-2.fc23.noarch                                                                                                        16/47 
  验证: python-warlock-1.0.1-3.fc23.noarch                                                                                                         17/47 
  验证: python-netaddr-0.7.18-1.fc23.noarch                                                                                                        18/47 
  验证: python-oslo-config-2:1.9.3-2.fc23.noarch                                                                                                   19/47 
  验证: python-fixtures-0.3.14-4.fc23.noarch                                                                                                       20/47 
  验证: python-netifaces-0.10.4-2.fc23.x86_64                                                                                                      21/47 
  验证: python-jsonpatch-1.2-6.fc23.noarch                                                                                                         22/47 
  验证: python-jsonschema-2.4.0-2.fc23.noarch                                                                                                      23/47 
  验证: python2-pbr-1.8.1-3.fc23.noarch                                                                                                            24/47 
  验证: python2-iso8601-0.1.11-1.fc23.noarch                                                                                                       25/47 
  验证: python2-msgpack-0.4.8-1.fc23.x86_64                                                                                                        26/47 
  验证: python-novaclient-1:2.23.3-1.fc23.noarch                                                                                                   27/47 
  验证: pyparsing-2.1.5-1.fc23.noarch                                                                                                              28/47 
  验证: python2-pyparsing-2.1.5-1.fc23.noarch                                                                                                      29/47 
  验证: python-keyring-5.0-5.fc23.noarch                                                                                                           30/47 
  验证: python-webob-1.4.2-1.fc23.noarch                                                                                                           31/47 
  验证: python-jsonpointer-1.10-1.fc23.noarch                                                                                                      32/47 
  验证: python-testtools-1.8.0-3.fc23.noarch                                                                                                       33/47 
  验证: python-extras-0.0.3-7.fc23.noarch                                                                                                          34/47 
  验证: python2-unittest2-1.1.0-5.fc23.noarch                                                                                                      35/47 
  验证: python-cmd2-0.6.8-8.fc23.noarch                                                                                                            36/47 
  验证: python2-cryptography-1.5.3-3.fc23.x86_64                                                                                                   37/47 
  验证: python-enum34-1.0.4-2.fc23.noarch                                                                                                          38/47 
  验证: python-idna-2.0-1.fc23.noarch                                                                                                              39/47 
  验证: python-ipaddress-1.0.7-4.fc23.noarch                                                                                                       40/47 
  验证: python-pyasn1-0.1.8-1.fc23.noarch                                                                                                          41/47 
  验证: python-cffi-1.4.2-1.fc23.x86_64                                                                                                            42/47 
  验证: python-pycparser-2.14-3.fc23.noarch                                                                                                        43/47 
  验证: python-ply-3.6-2.fc23.noarch                                                                                                               44/47 
  验证: python2-traceback2-1.4.0-7.fc23.noarch                                                                                                     45/47 
  验证: python-linecache2-1.0.0-1.fc23.noarch                                                                                                      46/47 
  验证: python-mimeparse-0.1.4-6.fc23.noarch                                                                                                       47/47 

已安装:
  pyOpenSSL.noarch 0.15.1-1.fc23                     pyparsing.noarch 2.1.5-1.fc23                        python-babel.noarch 1.3-8.fc23                 
  python-cffi.x86_64 1.4.2-1.fc23                    python-cinderclient.noarch 1.2.1-1.fc23              python-cliff.noarch 1.13.0-1.fc23              
  python-cmd2.noarch 0.6.8-8.fc23                    python-enum34.noarch 1.0.4-2.fc23                    python-extras.noarch 0.0.3-7.fc23              
  python-fixtures.noarch 0.3.14-4.fc23               python-glanceclient.noarch 1:0.17.0-3.fc23           python-httplib2.noarch 0.9.1-2.fc23            
  python-idna.noarch 2.0-1.fc23                      python-ipaddress.noarch 1.0.7-4.fc23                 python-jsonpatch.noarch 1.2-6.fc23             
  python-jsonpointer.noarch 1.10-1.fc23              python-jsonschema.noarch 2.4.0-2.fc23                python-keyring.noarch 5.0-5.fc23               
  python-keystoneclient.noarch 1:1.3.0-2.fc23        python-linecache2.noarch 1.0.0-1.fc23                python-mimeparse.noarch 0.1.4-6.fc23           
  python-netaddr.noarch 0.7.18-1.fc23                python-netifaces.x86_64 0.10.4-2.fc23                python-neutronclient.noarch 2.4.0-2.fc23       
  python-novaclient.noarch 1:2.23.3-1.fc23           python-openstackclient.noarch 1.0.3-3.fc23           python-oslo-config.noarch 2:1.9.3-2.fc23       
  python-oslo-i18n.noarch 1.5.0-4.fc23               python-oslo-serialization.noarch 1.4.0-2.fc23        python-oslo-utils.noarch 1.4.0-2.fc23          
  python-ply.noarch 3.6-2.fc23                       python-prettytable.noarch 0.7.2-5.fc23               python-pyasn1.noarch 0.1.8-1.fc23              
  python-pycparser.noarch 2.14-3.fc23                python-simplejson.x86_64 3.5.3-3.fc23                python-stevedore.noarch 1.5.0-1.fc23           
  python-testtools.noarch 1.8.0-3.fc23               python-warlock.noarch 1.0.1-3.fc23                   python-webob.noarch 1.4.2-1.fc23               
  python2-cryptography.x86_64 1.5.3-3.fc23           python2-iso8601.noarch 0.1.11-1.fc23                 python2-msgpack.x86_64 0.4.8-1.fc23            
  python2-pbr.noarch 1.8.1-3.fc23                    python2-pyparsing.noarch 2.1.5-1.fc23                python2-traceback2.noarch 1.4.0-7.fc23         
  python2-unittest2.noarch 1.1.0-5.fc23              pytz.noarch 2015.4-1.fc23                           

完毕！
```

### Before 

Environment variables
```
[vagrant@localhost ~]$ . openrc 
```

Run bin `openstack`
```
[vagrant@localhost ~]$ openstack network list
+--------------------------------------+-------------+--------------------------------------+
| ID                                   | Name        | Subnets                              |
+--------------------------------------+-------------+--------------------------------------+
| 830548b8-c7fc-435e-b144-b81f29b1e312 | private     | 931c8e3a-47df-42a6-aecd-9d81789b5fb7 |
| 92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d | public      | 27147d15-ce56-4913-9097-25d24a6d590e |
| a441b44a-866c-46ac-b454-ad5d43632e13 | private_net | 9d65712d-adb3-4e2c-a21f-3b8020b0e217 |
| cbb513ec-dd2a-4c0d-bf95-96e720223a04 | private_net | 52b1471f-55d2-44c4-9078-fe5ea16b3bdc |
| dc5a5abf-a581-492e-9dab-6fbdcdbc65bd | private_net | 69e6e877-3020-4874-86d9-75bfe6a6902d |
+--------------------------------------+-------------+--------------------------------------+
[vagrant@localhost ~]$ openstack network show public
+---------------------------+--------------------------------------+
| Field                     | Value                                |
+---------------------------+--------------------------------------+
| id                        | 92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d |
| name                      | public                               |
| project_id                | 8907b30a998647d5991547e9bbffa69a     |
| provider:network_type     | gre                                  |
| provider:physical_network | None                                 |
| provider:segmentation_id  | 3604                                 |
| router_type               | External                             |
| shared                    | False                                |
| state                     | UP                                   |
| status                    | ACTIVE                               |
| subnets                   | 27147d15-ce56-4913-9097-25d24a6d590e |
+---------------------------+--------------------------------------+
```

For more information, input `openstack --help`

Or run 'neutron'
```
[vagrant@localhost ~]$ neutron net-list
+--------------------------------------+-------------+------------------------------------------------------+
| id                                   | name        | subnets                                              |
+--------------------------------------+-------------+------------------------------------------------------+
| 830548b8-c7fc-435e-b144-b81f29b1e312 | private     | 931c8e3a-47df-42a6-aecd-9d81789b5fb7 192.168.0.0/24  |
| 92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d | public      | 27147d15-ce56-4913-9097-25d24a6d590e 10.100.151.0/24 |
| a441b44a-866c-46ac-b454-ad5d43632e13 | private_net | 9d65712d-adb3-4e2c-a21f-3b8020b0e217 10.20.30.0/24   |
| cbb513ec-dd2a-4c0d-bf95-96e720223a04 | private_net | 52b1471f-55d2-44c4-9078-fe5ea16b3bdc 10.20.30.0/24   |
| dc5a5abf-a581-492e-9dab-6fbdcdbc65bd | private_net | 69e6e877-3020-4874-86d9-75bfe6a6902d 10.20.30.0/24   |
+--------------------------------------+-------------+------------------------------------------------------+
[vagrant@localhost ~]$ neutron net-show public
+---------------------------+--------------------------------------+
| Field                     | Value                                |
+---------------------------+--------------------------------------+
| admin_state_up            | True                                 |
| id                        | 92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d |
| name                      | public                               |
| provider:network_type     | gre                                  |
| provider:physical_network |                                      |
| provider:segmentation_id  | 3604                                 |
| router:external           | True                                 |
| shared                    | False                                |
| status                    | ACTIVE                               |
| subnets                   | 27147d15-ce56-4913-9097-25d24a6d590e |
| tenant_id                 | 8907b30a998647d5991547e9bbffa69a     |
+---------------------------+--------------------------------------+
[vagrant@localhost ~]$ neutron subnet-list
+--------------------------------------+-----------------+-----------------+-----------------------------------------------------+
| id                                   | name            | cidr            | allocation_pools                                    |
+--------------------------------------+-----------------+-----------------+-----------------------------------------------------+
| 27147d15-ce56-4913-9097-25d24a6d590e | 10.100.151.0/24 | 10.100.151.0/24 | {"start": "10.100.151.50", "end": "10.100.151.240"} |
| 52b1471f-55d2-44c4-9078-fe5ea16b3bdc | private_subnet  | 10.20.30.0/24   | {"start": "10.20.30.2", "end": "10.20.30.254"}      |
| 69e6e877-3020-4874-86d9-75bfe6a6902d | private_subnet  | 10.20.30.0/24   | {"start": "10.20.30.2", "end": "10.20.30.254"}      |
| 931c8e3a-47df-42a6-aecd-9d81789b5fb7 | 192.168.0.0/24  | 192.168.0.0/24  | {"start": "192.168.0.2", "end": "192.168.0.254"}    |
| 9d65712d-adb3-4e2c-a21f-3b8020b0e217 | private_subnet  | 10.20.30.0/24   | {"start": "10.20.30.2", "end": "10.20.30.254"}      |
+--------------------------------------+-----------------+-----------------+-----------------------------------------------------+
[vagrant@localhost ~]$ neutron subnet-show 27147d15-ce56-4913-9097-25d24a6d590e
+-------------------+-----------------------------------------------------+
| Field             | Value                                               |
+-------------------+-----------------------------------------------------+
| allocation_pools  | {"start": "10.100.151.50", "end": "10.100.151.240"} |
| cidr              | 10.100.151.0/24                                     |
| dns_nameservers   | 114.114.114.114                                     |
| enable_dhcp       | False                                               |
| gateway_ip        | 10.100.151.254                                      |
| host_routes       |                                                     |
| id                | 27147d15-ce56-4913-9097-25d24a6d590e                |
| ip_version        | 4                                                   |
| ipv6_address_mode |                                                     |
| ipv6_ra_mode      |                                                     |
| name              | 10.100.151.0/24                                     |
| network_id        | 92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d                |
| tenant_id         | 8907b30a998647d5991547e9bbffa69a                    |
+-------------------+-----------------------------------------------------+
[vagrant@localhost ~]$ neutron router-list
+--------------------------------------+------------------------------------------------------------+--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------+-------+
| id                                   | name                                                       | external_gateway_info                                                                                                                                                                      | distributed | ha    |
+--------------------------------------+------------------------------------------------------------+--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------+-------+
| 76b77f20-4a0e-4c1c-bf4a-00fd00e315fc | jingqi621383039a6dc898e1b44713a72768d1-router-tqlgkitu5mbi | {"network_id": "92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d", "enable_snat": true, "external_fixed_ips": [{"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.101"}]} | False       | False |
| 951d2d2f-5417-413e-a18d-b0627d09b8dd | test                                                       | {"network_id": "92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d", "enable_snat": true, "external_fixed_ips": [{"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.50"}]}  | False       | False |
+--------------------------------------+------------------------------------------------------------+--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+-------------+-------+
[vagrant@localhost ~]$ neutron router-show 76b77f20-4a0e-4c1c-bf4a-00fd00e315fc
+-----------------------+--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Field                 | Value                                                                                                                                                                                      |
+-----------------------+--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| admin_state_up        | True                                                                                                                                                                                       |
| distributed           | False                                                                                                                                                                                      |
| external_gateway_info | {"network_id": "92c0f0c8-7c49-4b56-ae8b-86aa6c2f621d", "enable_snat": true, "external_fixed_ips": [{"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.101"}]} |
| ha                    | False                                                                                                                                                                                      |
| id                    | 76b77f20-4a0e-4c1c-bf4a-00fd00e315fc                                                                                                                                                       |
| name                  | jingqi621383039a6dc898e1b44713a72768d1-router-tqlgkitu5mbi                                                                                                                                 |
| routes                |                                                                                                                                                                                            |
| status                | ACTIVE                                                                                                                                                                                     |
| tenant_id             | a2a01453f7ed456a8d0d270ed5207697                                                                                                                                                           |
+-----------------------+--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
[vagrant@localhost ~]$ neutron router-port-list 76b77f20-4a0e-4c1c-bf4a-00fd00e315fc
+--------------------------------------+------+-------------------+---------------------------------------------------------------------------------------+
| id                                   | name | mac_address       | fixed_ips                                                                             |
+--------------------------------------+------+-------------------+---------------------------------------------------------------------------------------+
| 09548a18-27dd-4689-b20c-2504b43a63d0 |      | fa:16:3e:99:48:8a | {"subnet_id": "27147d15-ce56-4913-9097-25d24a6d590e", "ip_address": "10.100.151.101"} |
| 20917005-bd66-44d7-9764-c77c0688f1d4 |      | fa:16:3e:a5:9b:5c | {"subnet_id": "69e6e877-3020-4874-86d9-75bfe6a6902d", "ip_address": "10.20.30.1"}     |
+--------------------------------------+------+-------------------+---------------------------------------------------------------------------------------+
```
For more information, input `neutron --help`
