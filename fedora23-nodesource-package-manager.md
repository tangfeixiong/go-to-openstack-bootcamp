# Instruction

## Contents of table

* Nodejs
* Grunt
* Gulp
* Webpack

## Vagrant synced_folders into VirtulBox

Mac
```
    "sync_folders"      : 
        {
            ".":
              {
                "to"      : "/data/src/github.com/openshift/origin",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/99-mirror":
              {
                "to"      : "/Users/fanhongling/Downloads/99-mirror",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/99-mirror/linux-bin/k8s-v1.6.4/kubernetes":
              {
                "to"      : "/opt/kubernetes",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/go-kubernetes/pkg":
              {
                "to"      : "/Users/fanhongling/Downloads/go-kubernetes/pkg",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/go-kubernetes/src":
              {
                "to"      : "/Users/fanhongling/Downloads/go-kubernetes/src",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/go-openshift/pkg":
              {
                "to"      : "/Users/fanhongling/Downloads/go-openshift/pkg",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/go-openshift/src":
              {
                "to"      : "/Users/fanhongling/Downloads/go-openshift/src",
                "exclude" : null
              },
            "/Users/fanhongling/Downloads/workspace":
              {
                "to"      : "/Users/fanhongling/Downloads/workspace",
                "exclude" : null
              },
            "/Users/fanhongling/go/pkg":
              {
                "to"      : "/Users/fanhongling/go/pkg",
                "exclude" : null
              },
            "/Users/fanhongling/go/src":
              {
                "to"      : "/Users/fanhongling/go/src",
                "exclude" : null
              }
        }
```

## Nodejs

https://github.com/nodesource/distributions#rpminstall

```
[vagrant@localhost ~]$ curl -sL https://rpm.nodesource.com/setup_6.x | sudo -E bash -

## Installing the NodeSource Node.js 6.x repo...


## Inspecting system...

+ rpm -q --whatprovides redhat-release || rpm -q --whatprovides centos-release || rpm -q --whatprovides cloudlinux-release || rpm -q --whatprovides sl-release
+ uname -m

## Confirming "fc23-x86_64" is supported...

+ curl -sLf -o /dev/null 'https://rpm.nodesource.com/pub_6.x/fc/23/x86_64/nodesource-release-fc23-1.noarch.rpm'

## Downloading release setup RPM...

+ mktemp
+ curl -sL -o '/tmp/tmp.4YlQDgCe7s' 'https://rpm.nodesource.com/pub_6.x/fc/23/x86_64/nodesource-release-fc23-1.noarch.rpm'

## Installing release setup RPM...

+ rpm -i --nosignature --force '/tmp/tmp.4YlQDgCe7s'

## Cleaning up...

+ rm -f '/tmp/tmp.4YlQDgCe7s'

## Checking for existing installations...

+ rpm -qa 'node|npm' | grep -v nodesource

## Your system appears to already have Node.js installed from an alternative source.
## Run `yum remove -y nodejs npm` (as root) to remove these first.


## Run `yum install -y nodejs` (as root) to install Node.js 6.x and npm.
## You may also need development tools to build native addons:
##   `yum install -y gcc-c++ make`
```

Realy added
```
[vagrant@localhost ~]$ cat /etc/yum.repos.d/nodesource-fc.repo 
[nodesource]
name=Node.js Packages for Fedora Core 23 - $basearch
baseurl=https://rpm.nodesource.com/pub_6.x/fc/23/$basearch
failovermethod=priority
enabled=1
gpgcheck=1
gpgkey=file:///etc/pki/rpm-gpg/NODESOURCE-GPG-SIGNING-KEY-EL

[nodesource-source]
name=Node.js for Fedora Core 23 - $basearch - Source
baseurl=https://rpm.nodesource.com/pub_6.x/fc/23/SRPMS
failovermethod=priority
enabled=0
gpgkey=file:///etc/pki/rpm-gpg/NODESOURCE-GPG-SIGNING-KEY-EL
gpgcheck=1
```

Remove legacy version
```
[vagrant@localhost ~]$ sudo dnf remove -y nodejs npm
依赖关系解决。
====================================================================================================================================================
 Package                                       架构                   版本                                           仓库                      大小
====================================================================================================================================================
移除:
 compat-libuv010                               x86_64                 1:0.10.34-4.fc23                               @updates                 124 k
 compat-libuv010-devel                         x86_64                 1:0.10.34-4.fc23                               @updates                 142 k
 dwz                                           x86_64                 0.12-1.fc23                                    @fedora                  220 k
 ghc-srpm-macros                               noarch                 1.4.2-2.fc23                                   @fedora                  400  
 gnat-srpm-macros                              noarch                 2-1.fc23                                       @fedora                  650  
 gyp                                           noarch                 0.1-0.18.0bb6747git.fc23                       @fedora                  2.4 M
 http-parser                                   x86_64                 2.0-9.20121128gitcd01361.fc23                  @fedora                   56 k
 http-parser-devel                             x86_64                 2.0-9.20121128gitcd01361.fc23                  @fedora                   11 k
 keyutils-libs-devel                           x86_64                 1.5.9-7.fc23                                   @fedora                   31 k
 krb5-devel                                    x86_64                 1.14-9.fc23                                    @updates                 1.6 M
 libcom_err-devel                              x86_64                 1.42.13-3.fc23                                 @fedora                   17 k
 libuv                                         x86_64                 1:1.4.0-2.fc23                                 @fedora                  150 k
 libuv-devel                                   x86_64                 1:1.4.0-2.fc23                                 @fedora                   94 k
 libverto-devel                                x86_64                 0.2.6-5.fc23                                   @fedora                   25 k
 node-gyp                                      noarch                 0.10.6-5.fc23                                  @fedora                   66 k
 nodejs                                        x86_64                 0.10.42-4.fc23                                 @updates                 6.7 M
 nodejs-abbrev                                 noarch                 1.0.4-9.fc23                                   @fedora                  4.6 k
 nodejs-ansi                                   noarch                 0.2.1-3.fc23                                   @fedora                   17 k
 nodejs-archy                                  noarch                 0.0.2-11.fc23                                  @fedora                  4.3 k
 nodejs-asn1                                   noarch                 0.1.11-6.fc23                                  @fedora                   18 k
 nodejs-assert-plus                            noarch                 0.1.4-3.fc23                                   @fedora                  8.8 k
 nodejs-async                                  noarch                 0.2.10-3.fc23                                  @fedora                   73 k
 nodejs-aws-sign                               noarch                 0.3.0-4.fc23                                   @fedora                  4.2 k
 nodejs-block-stream                           noarch                 0.0.7-3.fc23                                   @fedora                  8.5 k
 nodejs-boom                                   noarch                 0.4.2-5.fc23                                   @fedora                   36 k
 nodejs-child-process-close                    noarch                 0.1.1-5.fc23                                   @fedora                  3.3 k
 nodejs-chmodr                                 noarch                 0.1.0-7.fc23                                   @fedora                  3.0 k
 nodejs-chownr                                 noarch                 0.0.1-12.fc23                                  @fedora                  2.8 k
 nodejs-cmd-shim                               noarch                 1.1.0-6.fc23                                   @fedora                  7.3 k
 nodejs-combined-stream                        noarch                 0.0.4-6.fc23                                   @fedora                  9.4 k
 nodejs-config-chain                           noarch                 1.1.7-4.fc23                                   @fedora                   14 k
 nodejs-cookie-jar                             noarch                 1:0.3.0-4.fc23                                 @fedora                  3.2 k
 nodejs-cryptiles                              noarch                 0.2.2-3.fc23                                   @fedora                  3.7 k
 nodejs-ctype                                  noarch                 0.5.3-6.fc23                                   @fedora                   92 k
 nodejs-delayed-stream                         noarch                 0.0.5-8.fc23                                   @fedora                  7.8 k
 nodejs-devel                                  x86_64                 0.10.42-4.fc23                                 @updates                 9.1 M
 nodejs-editor                                 noarch                 0.0.4-5.fc23                                   @fedora                  3.4 k
 nodejs-forever-agent                          noarch                 0.5.0-4.fc23                                   @fedora                  4.1 k
 nodejs-form-data                              noarch                 0.1.1-3.fc23                                   @fedora                   14 k
 nodejs-fstream                                noarch                 0.1.24-3.fc23                                  @fedora                   62 k
 nodejs-fstream-ignore                         noarch                 0.0.7-4.fc23                                   @fedora                  8.8 k
 nodejs-fstream-npm                            noarch                 0.1.5-4.fc23                                   @fedora                   14 k
 nodejs-github-url-from-git                    noarch                 1.1.1-5.fc23                                   @fedora                  3.2 k
 nodejs-glob                                   noarch                 3.2.6-3.fc23                                   @fedora                   29 k
 nodejs-graceful-fs                            noarch                 2.0.0-5.fc23                                   @fedora                   12 k
 nodejs-hawk                                   noarch                 1.0.0-3.fc23                                   @fedora                  167 k
 nodejs-hoek                                   noarch                 0.9.1-4.fc23                                   @fedora                   67 k
 nodejs-http-signature                         noarch                 0.10.0-6.fc23                                  @fedora                   35 k
 nodejs-inherits                               noarch                 2.0.1-6.fc23                                   @fedora                  2.6 k
 nodejs-ini                                    noarch                 1.1.0-6.fc23                                   @fedora                  7.7 k
 nodejs-init-package-json                      noarch                 0.0.10-4.fc23                                  @fedora                   10 k
 nodejs-json-stringify-safe                    noarch                 5.0.0-4.fc23                                   @fedora                  3.8 k
 nodejs-lockfile                               noarch                 0.4.2-3.fc23                                   @fedora                   19 k
 nodejs-lru-cache                              noarch                 2.3.0-6.fc23                                   @fedora                   11 k
 nodejs-mime                                   noarch                 1.2.11-3.fc23                                  @fedora                   61 k
 nodejs-minimatch                              noarch                 0.2.12-5.fc23                                  @fedora                   36 k
 nodejs-mkdirp                                 noarch                 0.3.5-6.fc23                                   @fedora                  5.2 k
 nodejs-mute-stream                            noarch                 0.0.4-3.fc23                                   @fedora                  6.7 k
 nodejs-node-uuid                              noarch                 1.4.1-3.fc23                                   @fedora                   14 k
 nodejs-nopt                                   noarch                 2.1.2-3.fc23                                   @fedora                   28 k
 nodejs-normalize-package-data                 noarch                 0.2.1-3.fc23                                   @fedora                   20 k
 nodejs-npm-registry-client                    noarch                 0.2.28-4.fc23                                  @fedora                   38 k
 nodejs-npm-user-validate                      noarch                 0.0.3-4.fc23                                   @fedora                  3.1 k
 nodejs-npmconf                                noarch                 0.1.3-3.fc23                                   @fedora                   27 k
 nodejs-npmlog                                 noarch                 0.0.4-4.fc23                                   @fedora                   11 k
 nodejs-oauth-sign                             noarch                 0.3.0-4.fc23                                   @fedora                  1.5 k
 nodejs-once                                   noarch                 1.1.1-8.fc23                                   @fedora                  2.7 k
 nodejs-opener                                 noarch                 1.3.0-10.fc23                                  @fedora                  3.8 k
 nodejs-osenv                                  noarch                 0.0.3-8.fc23                                   @fedora                  5.4 k
 nodejs-packaging                              noarch                 7-4.fc23                                       @fedora                   19 k
 nodejs-promzard                               noarch                 0.2.0-9.fc23                                   @fedora                   19 k
 nodejs-proto-list                             noarch                 1.2.2-8.fc23                                   @fedora                  3.7 k
 nodejs-qs                                     noarch                 0.6.6-4.fc23                                   @fedora                  9.7 k
 nodejs-read                                   noarch                 1.0.5-3.fc23                                   @fedora                  6.3 k
 nodejs-read-installed                         noarch                 0.2.4-4.fc23                                   @fedora                   11 k
 nodejs-read-package-json                      noarch                 1.1.3-3.fc23                                   @fedora                   20 k
 nodejs-request                                noarch                 2.25.0-3.fc23                                  @fedora                   63 k
 nodejs-retry                                  noarch                 0.6.0-8.fc23                                   @fedora                   12 k
 nodejs-rimraf                                 noarch                 2.2.2-3.fc23                                   @fedora                  6.6 k
 nodejs-semver                                 noarch                 2.1.0-4.fc23                                   @fedora                   35 k
 nodejs-sha                                    noarch                 1.2.1-3.fc23                                   @fedora                  7.4 k
 nodejs-sigmund                                noarch                 1.0.0-8.fc23                                   @fedora                   12 k
 nodejs-slide                                  noarch                 1.1.5-3.fc23                                   @fedora                  9.5 k
 nodejs-sntp                                   noarch                 0.2.4-4.fc23                                   @fedora                   14 k
 nodejs-tar                                    noarch                 0.1.18-3.fc23                                  @fedora                   52 k
 nodejs-tunnel-agent                           noarch                 0.3.0-4.fc23                                   @fedora                  6.7 k
 nodejs-uid-number                             noarch                 0.0.3-10.fc23                                  @fedora                  4.3 k
 nodejs-which                                  noarch                 1.0.5-11.fc23                                  @fedora                  4.9 k
 npm                                           noarch                 1.3.6-8.fc23                                   @fedora                  996 k
 ocaml-srpm-macros                             noarch                 2-3.fc23                                       @fedora                  537  
 openssl-devel                                 x86_64                 1:1.0.2f-1.fc23                                @updates                 3.1 M
 perl-srpm-macros                              noarch                 1-17.fc23                                      @fedora                  794  
 redhat-rpm-config                             noarch                 36-1.fc23                                      @fedora                   99 k
 v8                                            x86_64                 1:3.14.5.10-21.fc23                            @fedora                  9.6 M
 v8-devel                                      x86_64                 1:3.14.5.10-21.fc23                            @fedora                  197 k
 zlib-devel                                    x86_64                 1.2.8-9.fc23                                   @fedora                  134 k

事务概要
====================================================================================================================================================
移除  96 Packages

安装大小：36 M
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  删除: npm-1.3.6-8.fc23.noarch                                                                                                                1/96 
警告：文件 /usr/lib/node_modules/npm/node_modules/npmconf: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/node_modules/minimatch: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/node_modules/lru-cache: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/node_modules/chmodr: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/node_modules/child-process-close: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/node_modules/block-stream: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/node_modules/ansi: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/lib/utils/is-git-url.js: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/lib/utils/find-prefix.js: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/lib/utils/fetch.js: 移除失败: 没有那个文件或目录
警告：文件 /usr/lib/node_modules/npm/lib/submodule.js: 移除失败: 没有那个文件或目录
  删除: node-gyp-0.10.6-5.fc23.noarch                                                                                                          2/96 
  删除: nodejs-npm-registry-client-0.2.28-4.fc23.noarch                                                                                        3/96 
  删除: nodejs-request-2.25.0-3.fc23.noarch                                                                                                    4/96 
  删除: nodejs-npmconf-0.1.3-3.fc23.noarch                                                                                                     5/96 
  删除: nodejs-hawk-1.0.0-3.fc23.noarch                                                                                                        6/96 
  删除: nodejs-init-package-json-0.0.10-4.fc23.noarch                                                                                          7/96 
  删除: nodejs-form-data-0.1.1-3.fc23.noarch                                                                                                   8/96 
  删除: nodejs-http-signature-0.10.0-6.fc23.noarch                                                                                             9/96 
  删除: nodejs-tar-0.1.18-3.fc23.noarch                                                                                                       10/96 
  删除: nodejs-read-installed-0.2.4-4.fc23.noarch                                                                                             11/96 
  删除: nodejs-read-package-json-1.1.3-3.fc23.noarch                                                                                          12/96 
  删除: nodejs-glob-3.2.6-3.fc23.noarch                                                                                                       13/96 
  删除: nodejs-normalize-package-data-0.2.1-3.fc23.noarch                                                                                     14/96 
  删除: nodejs-config-chain-1.1.7-4.fc23.noarch                                                                                               15/96 
  删除: nodejs-fstream-npm-0.1.5-4.fc23.noarch                                                                                                16/96 
  删除: nodejs-fstream-ignore-0.0.7-4.fc23.noarch                                                                                             17/96 
  删除: nodejs-fstream-0.1.24-3.fc23.noarch                                                                                                   18/96 
  删除: nodejs-minimatch-0.2.12-5.fc23.noarch                                                                                                 19/96 
  删除: nodejs-block-stream-0.0.7-3.fc23.noarch                                                                                               20/96 
  删除: nodejs-combined-stream-0.0.4-6.fc23.noarch                                                                                            21/96 
  删除: nodejs-promzard-0.2.0-9.fc23.noarch                                                                                                   22/96 
  删除: nodejs-read-1.0.5-3.fc23.noarch                                                                                                       23/96 
  删除: nodejs-cryptiles-0.2.2-3.fc23.noarch                                                                                                  24/96 
  删除: nodejs-boom-0.4.2-5.fc23.noarch                                                                                                       25/96 
  删除: nodejs-sntp-0.2.4-4.fc23.noarch                                                                                                       26/96 
  删除: nodejs-nopt-2.1.2-3.fc23.noarch                                                                                                       27/96 
  删除: nodejs-npmlog-0.0.4-4.fc23.noarch                                                                                                     28/96 
  删除: nodejs-cmd-shim-1.1.0-6.fc23.noarch                                                                                                   29/96 
  删除: nodejs-sha-1.2.1-3.fc23.noarch                                                                                                        30/96 
  删除: nodejs-graceful-fs-2.0.0-5.fc23.noarch                                                                                                31/96 
  删除: nodejs-mkdirp-0.3.5-6.fc23.noarch                                                                                                     32/96 
  删除: nodejs-ansi-0.2.1-3.fc23.noarch                                                                                                       33/96 
  删除: nodejs-abbrev-1.0.4-9.fc23.noarch                                                                                                     34/96 
  删除: nodejs-hoek-0.9.1-4.fc23.noarch                                                                                                       35/96 
  删除: nodejs-mute-stream-0.0.4-3.fc23.noarch                                                                                                36/96 
  删除: nodejs-delayed-stream-0.0.5-8.fc23.noarch                                                                                             37/96 
  删除: nodejs-inherits-2.0.1-6.fc23.noarch                                                                                                   38/96 
  删除: nodejs-lru-cache-2.3.0-6.fc23.noarch                                                                                                  39/96 
  删除: nodejs-sigmund-1.0.0-8.fc23.noarch                                                                                                    40/96 
  删除: nodejs-rimraf-2.2.2-3.fc23.noarch                                                                                                     41/96 
  删除: nodejs-ini-1.1.0-6.fc23.noarch                                                                                                        42/96 
  删除: nodejs-proto-list-1.2.2-8.fc23.noarch                                                                                                 43/96 
  删除: nodejs-github-url-from-git-1.1.1-5.fc23.noarch                                                                                        44/96 
  删除: nodejs-semver-2.1.0-4.fc23.noarch                                                                                                     45/96 
  删除: nodejs-slide-1.1.5-3.fc23.noarch                                                                                                      46/96 
  删除: nodejs-asn1-0.1.11-6.fc23.noarch                                                                                                      47/96 
  删除: nodejs-assert-plus-0.1.4-3.fc23.noarch                                                                                                48/96 
  删除: nodejs-ctype-0.5.3-6.fc23.noarch                                                                                                      49/96 
  删除: nodejs-async-0.2.10-3.fc23.noarch                                                                                                     50/96 
  删除: nodejs-mime-1.2.11-3.fc23.noarch                                                                                                      51/96 
  删除: nodejs-once-1.1.1-8.fc23.noarch                                                                                                       52/96 
  删除: nodejs-osenv-0.0.3-8.fc23.noarch                                                                                                      53/96 
  删除: nodejs-aws-sign-0.3.0-4.fc23.noarch                                                                                                   54/96 
  删除: nodejs-cookie-jar-1:0.3.0-4.fc23.noarch                                                                                               55/96 
  删除: nodejs-forever-agent-0.5.0-4.fc23.noarch                                                                                              56/96 
  删除: nodejs-json-stringify-safe-5.0.0-4.fc23.noarch                                                                                        57/96 
  删除: nodejs-node-uuid-1.4.1-3.fc23.noarch                                                                                                  58/96 
  删除: nodejs-oauth-sign-0.3.0-4.fc23.noarch                                                                                                 59/96 
  删除: nodejs-qs-0.6.6-4.fc23.noarch                                                                                                         60/96 
  删除: nodejs-tunnel-agent-0.3.0-4.fc23.noarch                                                                                               61/96 
  删除: nodejs-chownr-0.0.1-12.fc23.noarch                                                                                                    62/96 
  删除: nodejs-retry-0.6.0-8.fc23.noarch                                                                                                      63/96 
  删除: http-parser-devel-2.0-9.20121128gitcd01361.fc23.x86_64                                                                                64/96 
  删除: libuv-devel-1:1.4.0-2.fc23.x86_64                                                                                                     65/96 
  删除: nodejs-which-1.0.5-11.fc23.noarch                                                                                                     66/96 
  删除: nodejs-archy-0.0.2-11.fc23.noarch                                                                                                     67/96 
  删除: nodejs-child-process-close-0.1.1-5.fc23.noarch                                                                                        68/96 
  删除: nodejs-chmodr-0.1.0-7.fc23.noarch                                                                                                     69/96 
  删除: nodejs-editor-0.0.4-5.fc23.noarch                                                                                                     70/96 
  删除: nodejs-lockfile-0.4.2-3.fc23.noarch                                                                                                   71/96 
  删除: nodejs-npm-user-validate-0.0.3-4.fc23.noarch                                                                                          72/96 
  删除: nodejs-opener-1.3.0-10.fc23.noarch                                                                                                    73/96 
  删除: nodejs-uid-number-0.0.3-10.fc23.noarch                                                                                                74/96 
  删除: nodejs-devel-0.10.42-4.fc23.x86_64                                                                                                    75/96 
  删除: nodejs-packaging-7-4.fc23.noarch                                                                                                      76/96 
  删除: redhat-rpm-config-36-1.fc23.noarch                                                                                                    77/96 
  删除: openssl-devel-1:1.0.2f-1.fc23.x86_64                                                                                                  78/96 
  删除: compat-libuv010-devel-1:0.10.34-4.fc23.x86_64                                                                                         79/96 
  删除: v8-devel-1:3.14.5.10-21.fc23.x86_64                                                                                                   80/96 
  删除: zlib-devel-1.2.8-9.fc23.x86_64                                                                                                        81/96 
  删除: ghc-srpm-macros-1.4.2-2.fc23.noarch                                                                                                   82/96 
  删除: gnat-srpm-macros-2-1.fc23.noarch                                                                                                      83/96 
  删除: ocaml-srpm-macros-2-3.fc23.noarch                                                                                                     84/96 
  删除: perl-srpm-macros-1-17.fc23.noarch                                                                                                     85/96 
  删除: gyp-0.1-0.18.0bb6747git.fc23.noarch                                                                                                   86/96 
  删除: krb5-devel-1.14-9.fc23.x86_64                                                                                                         87/96 
  删除: nodejs-0.10.42-4.fc23.x86_64                                                                                                          88/96 
  删除: keyutils-libs-devel-1.5.9-7.fc23.x86_64                                                                                               89/96 
  删除: libcom_err-devel-1.42.13-3.fc23.x86_64                                                                                                90/96 
  删除: libverto-devel-0.2.6-5.fc23.x86_64                                                                                                    91/96 
  删除: compat-libuv010-1:0.10.34-4.fc23.x86_64                                                                                               92/96 
  删除: v8-1:3.14.5.10-21.fc23.x86_64                                                                                                         93/96 
  删除: dwz-0.12-1.fc23.x86_64                                                                                                                94/96 
  删除: libuv-1:1.4.0-2.fc23.x86_64                                                                                                           95/96 
  删除: http-parser-2.0-9.20121128gitcd01361.fc23.x86_64                                                                                      96/96 
  验证: node-gyp-0.10.6-5.fc23.noarch                                                                                                          1/96 
  验证: nodejs-0.10.42-4.fc23.x86_64                                                                                                           2/96 
  验证: nodejs-abbrev-1.0.4-9.fc23.noarch                                                                                                      3/96 
  验证: nodejs-ansi-0.2.1-3.fc23.noarch                                                                                                        4/96 
  验证: nodejs-archy-0.0.2-11.fc23.noarch                                                                                                      5/96 
  验证: nodejs-asn1-0.1.11-6.fc23.noarch                                                                                                       6/96 
  验证: nodejs-assert-plus-0.1.4-3.fc23.noarch                                                                                                 7/96 
  验证: nodejs-async-0.2.10-3.fc23.noarch                                                                                                      8/96 
  验证: nodejs-aws-sign-0.3.0-4.fc23.noarch                                                                                                    9/96 
  验证: nodejs-block-stream-0.0.7-3.fc23.noarch                                                                                               10/96 
  验证: nodejs-boom-0.4.2-5.fc23.noarch                                                                                                       11/96 
  验证: nodejs-child-process-close-0.1.1-5.fc23.noarch                                                                                        12/96 
  验证: nodejs-chmodr-0.1.0-7.fc23.noarch                                                                                                     13/96 
  验证: nodejs-chownr-0.0.1-12.fc23.noarch                                                                                                    14/96 
  验证: nodejs-cmd-shim-1.1.0-6.fc23.noarch                                                                                                   15/96 
  验证: nodejs-combined-stream-0.0.4-6.fc23.noarch                                                                                            16/96 
  验证: nodejs-config-chain-1.1.7-4.fc23.noarch                                                                                               17/96 
  验证: nodejs-cookie-jar-1:0.3.0-4.fc23.noarch                                                                                               18/96 
  验证: nodejs-cryptiles-0.2.2-3.fc23.noarch                                                                                                  19/96 
  验证: nodejs-ctype-0.5.3-6.fc23.noarch                                                                                                      20/96 
  验证: nodejs-delayed-stream-0.0.5-8.fc23.noarch                                                                                             21/96 
  验证: nodejs-devel-0.10.42-4.fc23.x86_64                                                                                                    22/96 
  验证: nodejs-editor-0.0.4-5.fc23.noarch                                                                                                     23/96 
  验证: nodejs-forever-agent-0.5.0-4.fc23.noarch                                                                                              24/96 
  验证: nodejs-form-data-0.1.1-3.fc23.noarch                                                                                                  25/96 
  验证: nodejs-fstream-0.1.24-3.fc23.noarch                                                                                                   26/96 
  验证: nodejs-fstream-ignore-0.0.7-4.fc23.noarch                                                                                             27/96 
  验证: nodejs-fstream-npm-0.1.5-4.fc23.noarch                                                                                                28/96 
  验证: nodejs-github-url-from-git-1.1.1-5.fc23.noarch                                                                                        29/96 
  验证: nodejs-glob-3.2.6-3.fc23.noarch                                                                                                       30/96 
  验证: nodejs-graceful-fs-2.0.0-5.fc23.noarch                                                                                                31/96 
  验证: nodejs-hawk-1.0.0-3.fc23.noarch                                                                                                       32/96 
  验证: nodejs-hoek-0.9.1-4.fc23.noarch                                                                                                       33/96 
  验证: nodejs-http-signature-0.10.0-6.fc23.noarch                                                                                            34/96 
  验证: nodejs-inherits-2.0.1-6.fc23.noarch                                                                                                   35/96 
  验证: nodejs-ini-1.1.0-6.fc23.noarch                                                                                                        36/96 
  验证: nodejs-init-package-json-0.0.10-4.fc23.noarch                                                                                         37/96 
  验证: nodejs-json-stringify-safe-5.0.0-4.fc23.noarch                                                                                        38/96 
  验证: nodejs-lockfile-0.4.2-3.fc23.noarch                                                                                                   39/96 
  验证: nodejs-lru-cache-2.3.0-6.fc23.noarch                                                                                                  40/96 
  验证: nodejs-mime-1.2.11-3.fc23.noarch                                                                                                      41/96 
  验证: nodejs-minimatch-0.2.12-5.fc23.noarch                                                                                                 42/96 
  验证: nodejs-mkdirp-0.3.5-6.fc23.noarch                                                                                                     43/96 
  验证: nodejs-mute-stream-0.0.4-3.fc23.noarch                                                                                                44/96 
  验证: nodejs-node-uuid-1.4.1-3.fc23.noarch                                                                                                  45/96 
  验证: nodejs-nopt-2.1.2-3.fc23.noarch                                                                                                       46/96 
  验证: nodejs-normalize-package-data-0.2.1-3.fc23.noarch                                                                                     47/96 
  验证: nodejs-npm-registry-client-0.2.28-4.fc23.noarch                                                                                       48/96 
  验证: nodejs-npm-user-validate-0.0.3-4.fc23.noarch                                                                                          49/96 
  验证: nodejs-npmconf-0.1.3-3.fc23.noarch                                                                                                    50/96 
  验证: nodejs-npmlog-0.0.4-4.fc23.noarch                                                                                                     51/96 
  验证: nodejs-oauth-sign-0.3.0-4.fc23.noarch                                                                                                 52/96 
  验证: nodejs-once-1.1.1-8.fc23.noarch                                                                                                       53/96 
  验证: nodejs-opener-1.3.0-10.fc23.noarch                                                                                                    54/96 
  验证: nodejs-osenv-0.0.3-8.fc23.noarch                                                                                                      55/96 
  验证: nodejs-packaging-7-4.fc23.noarch                                                                                                      56/96 
  验证: nodejs-promzard-0.2.0-9.fc23.noarch                                                                                                   57/96 
  验证: nodejs-proto-list-1.2.2-8.fc23.noarch                                                                                                 58/96 
  验证: keyutils-libs-devel-1.5.9-7.fc23.x86_64                                                                                               59/96 
  验证: nodejs-qs-0.6.6-4.fc23.noarch                                                                                                         60/96 
  验证: nodejs-read-1.0.5-3.fc23.noarch                                                                                                       61/96 
  验证: nodejs-read-installed-0.2.4-4.fc23.noarch                                                                                             62/96 
  验证: krb5-devel-1.14-9.fc23.x86_64                                                                                                         63/96 
  验证: nodejs-read-package-json-1.1.3-3.fc23.noarch                                                                                          64/96 
  验证: nodejs-request-2.25.0-3.fc23.noarch                                                                                                   65/96 
  验证: nodejs-retry-0.6.0-8.fc23.noarch                                                                                                      66/96 
  验证: nodejs-rimraf-2.2.2-3.fc23.noarch                                                                                                     67/96 
  验证: nodejs-semver-2.1.0-4.fc23.noarch                                                                                                     68/96 
  验证: nodejs-sha-1.2.1-3.fc23.noarch                                                                                                        69/96 
  验证: nodejs-sigmund-1.0.0-8.fc23.noarch                                                                                                    70/96 
  验证: nodejs-slide-1.1.5-3.fc23.noarch                                                                                                      71/96 
  验证: nodejs-sntp-0.2.4-4.fc23.noarch                                                                                                       72/96 
  验证: nodejs-tar-0.1.18-3.fc23.noarch                                                                                                       73/96 
  验证: nodejs-tunnel-agent-0.3.0-4.fc23.noarch                                                                                               74/96 
  验证: nodejs-uid-number-0.0.3-10.fc23.noarch                                                                                                75/96 
  验证: nodejs-which-1.0.5-11.fc23.noarch                                                                                                     76/96 
  验证: npm-1.3.6-8.fc23.noarch                                                                                                               77/96 
  验证: compat-libuv010-1:0.10.34-4.fc23.x86_64                                                                                               78/96 
  验证: compat-libuv010-devel-1:0.10.34-4.fc23.x86_64                                                                                         79/96 
  验证: redhat-rpm-config-36-1.fc23.noarch                                                                                                    80/96 
  验证: ocaml-srpm-macros-2-3.fc23.noarch                                                                                                     81/96 
  验证: openssl-devel-1:1.0.2f-1.fc23.x86_64                                                                                                  82/96 
  验证: libcom_err-devel-1.42.13-3.fc23.x86_64                                                                                                83/96 
  验证: dwz-0.12-1.fc23.x86_64                                                                                                                84/96 
  验证: v8-1:3.14.5.10-21.fc23.x86_64                                                                                                         85/96 
  验证: v8-devel-1:3.14.5.10-21.fc23.x86_64                                                                                                   86/96 
  验证: ghc-srpm-macros-1.4.2-2.fc23.noarch                                                                                                   87/96 
  验证: libuv-1:1.4.0-2.fc23.x86_64                                                                                                           88/96 
  验证: libuv-devel-1:1.4.0-2.fc23.x86_64                                                                                                     89/96 
  验证: libverto-devel-0.2.6-5.fc23.x86_64                                                                                                    90/96 
  验证: zlib-devel-1.2.8-9.fc23.x86_64                                                                                                        91/96 
  验证: gnat-srpm-macros-2-1.fc23.noarch                                                                                                      92/96 
  验证: gyp-0.1-0.18.0bb6747git.fc23.noarch                                                                                                   93/96 
  验证: http-parser-2.0-9.20121128gitcd01361.fc23.x86_64                                                                                      94/96 
  验证: http-parser-devel-2.0-9.20121128gitcd01361.fc23.x86_64                                                                                95/96 
  验证: perl-srpm-macros-1-17.fc23.noarch                                                                                                     96/96 

已移除:
  compat-libuv010.x86_64 1:0.10.34-4.fc23                                compat-libuv010-devel.x86_64 1:0.10.34-4.fc23                              
  dwz.x86_64 0.12-1.fc23                                                 ghc-srpm-macros.noarch 1.4.2-2.fc23                                        
  gnat-srpm-macros.noarch 2-1.fc23                                       gyp.noarch 0.1-0.18.0bb6747git.fc23                                        
  http-parser.x86_64 2.0-9.20121128gitcd01361.fc23                       http-parser-devel.x86_64 2.0-9.20121128gitcd01361.fc23                     
  keyutils-libs-devel.x86_64 1.5.9-7.fc23                                krb5-devel.x86_64 1.14-9.fc23                                              
  libcom_err-devel.x86_64 1.42.13-3.fc23                                 libuv.x86_64 1:1.4.0-2.fc23                                                
  libuv-devel.x86_64 1:1.4.0-2.fc23                                      libverto-devel.x86_64 0.2.6-5.fc23                                         
  node-gyp.noarch 0.10.6-5.fc23                                          nodejs.x86_64 0.10.42-4.fc23                                               
  nodejs-abbrev.noarch 1.0.4-9.fc23                                      nodejs-ansi.noarch 0.2.1-3.fc23                                            
  nodejs-archy.noarch 0.0.2-11.fc23                                      nodejs-asn1.noarch 0.1.11-6.fc23                                           
  nodejs-assert-plus.noarch 0.1.4-3.fc23                                 nodejs-async.noarch 0.2.10-3.fc23                                          
  nodejs-aws-sign.noarch 0.3.0-4.fc23                                    nodejs-block-stream.noarch 0.0.7-3.fc23                                    
  nodejs-boom.noarch 0.4.2-5.fc23                                        nodejs-child-process-close.noarch 0.1.1-5.fc23                             
  nodejs-chmodr.noarch 0.1.0-7.fc23                                      nodejs-chownr.noarch 0.0.1-12.fc23                                         
  nodejs-cmd-shim.noarch 1.1.0-6.fc23                                    nodejs-combined-stream.noarch 0.0.4-6.fc23                                 
  nodejs-config-chain.noarch 1.1.7-4.fc23                                nodejs-cookie-jar.noarch 1:0.3.0-4.fc23                                    
  nodejs-cryptiles.noarch 0.2.2-3.fc23                                   nodejs-ctype.noarch 0.5.3-6.fc23                                           
  nodejs-delayed-stream.noarch 0.0.5-8.fc23                              nodejs-devel.x86_64 0.10.42-4.fc23                                         
  nodejs-editor.noarch 0.0.4-5.fc23                                      nodejs-forever-agent.noarch 0.5.0-4.fc23                                   
  nodejs-form-data.noarch 0.1.1-3.fc23                                   nodejs-fstream.noarch 0.1.24-3.fc23                                        
  nodejs-fstream-ignore.noarch 0.0.7-4.fc23                              nodejs-fstream-npm.noarch 0.1.5-4.fc23                                     
  nodejs-github-url-from-git.noarch 1.1.1-5.fc23                         nodejs-glob.noarch 3.2.6-3.fc23                                            
  nodejs-graceful-fs.noarch 2.0.0-5.fc23                                 nodejs-hawk.noarch 1.0.0-3.fc23                                            
  nodejs-hoek.noarch 0.9.1-4.fc23                                        nodejs-http-signature.noarch 0.10.0-6.fc23                                 
  nodejs-inherits.noarch 2.0.1-6.fc23                                    nodejs-ini.noarch 1.1.0-6.fc23                                             
  nodejs-init-package-json.noarch 0.0.10-4.fc23                          nodejs-json-stringify-safe.noarch 5.0.0-4.fc23                             
  nodejs-lockfile.noarch 0.4.2-3.fc23                                    nodejs-lru-cache.noarch 2.3.0-6.fc23                                       
  nodejs-mime.noarch 1.2.11-3.fc23                                       nodejs-minimatch.noarch 0.2.12-5.fc23                                      
  nodejs-mkdirp.noarch 0.3.5-6.fc23                                      nodejs-mute-stream.noarch 0.0.4-3.fc23                                     
  nodejs-node-uuid.noarch 1.4.1-3.fc23                                   nodejs-nopt.noarch 2.1.2-3.fc23                                            
  nodejs-normalize-package-data.noarch 0.2.1-3.fc23                      nodejs-npm-registry-client.noarch 0.2.28-4.fc23                            
  nodejs-npm-user-validate.noarch 0.0.3-4.fc23                           nodejs-npmconf.noarch 0.1.3-3.fc23                                         
  nodejs-npmlog.noarch 0.0.4-4.fc23                                      nodejs-oauth-sign.noarch 0.3.0-4.fc23                                      
  nodejs-once.noarch 1.1.1-8.fc23                                        nodejs-opener.noarch 1.3.0-10.fc23                                         
  nodejs-osenv.noarch 0.0.3-8.fc23                                       nodejs-packaging.noarch 7-4.fc23                                           
  nodejs-promzard.noarch 0.2.0-9.fc23                                    nodejs-proto-list.noarch 1.2.2-8.fc23                                      
  nodejs-qs.noarch 0.6.6-4.fc23                                          nodejs-read.noarch 1.0.5-3.fc23                                            
  nodejs-read-installed.noarch 0.2.4-4.fc23                              nodejs-read-package-json.noarch 1.1.3-3.fc23                               
  nodejs-request.noarch 2.25.0-3.fc23                                    nodejs-retry.noarch 0.6.0-8.fc23                                           
  nodejs-rimraf.noarch 2.2.2-3.fc23                                      nodejs-semver.noarch 2.1.0-4.fc23                                          
  nodejs-sha.noarch 1.2.1-3.fc23                                         nodejs-sigmund.noarch 1.0.0-8.fc23                                         
  nodejs-slide.noarch 1.1.5-3.fc23                                       nodejs-sntp.noarch 0.2.4-4.fc23                                            
  nodejs-tar.noarch 0.1.18-3.fc23                                        nodejs-tunnel-agent.noarch 0.3.0-4.fc23                                    
  nodejs-uid-number.noarch 0.0.3-10.fc23                                 nodejs-which.noarch 1.0.5-11.fc23                                          
  npm.noarch 1.3.6-8.fc23                                                ocaml-srpm-macros.noarch 2-3.fc23                                          
  openssl-devel.x86_64 1:1.0.2f-1.fc23                                   perl-srpm-macros.noarch 1-17.fc23                                          
  redhat-rpm-config.noarch 36-1.fc23                                     v8.x86_64 1:3.14.5.10-21.fc23                                              
  v8-devel.x86_64 1:3.14.5.10-21.fc23                                    zlib-devel.x86_64 1.2.8-9.fc23                                             

完毕！
```

Install new
```
[vagrant@localhost ~]$ sudo dnf install -y nodejs
Node.js Packages for Fedora Core 23 - x86_64                                                                        213 kB/s | 632 kB     00:02    
上次元数据过期检查在 0:00:02 前执行于 Fri Jun  9 18:19:48 2017。
依赖关系解决。
====================================================================================================================================================
 Package                      架构                         版本                                              仓库                              大小
====================================================================================================================================================
安装:
 nodejs                       x86_64                       2:6.9.5-1nodesource.fc23                          nodesource                       9.7 M

事务概要
====================================================================================================================================================
安装  1 Package

总下载：9.7 M
安装大小：33 M
下载软件包：
nodejs-6.9.5-1nodesource.fc23.x86_64.rpm                                                                            431 kB/s | 9.7 MB     00:23    
----------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                                431 kB/s | 9.7 MB     00:23     
警告：/var/cache/dnf/nodesource-2b3259932d4c93d6/packages/nodejs-6.9.5-1nodesource.fc23.x86_64.rpm: 头V4 RSA/SHA1 Signature, 密钥 ID 34fa74dd: NOKEY
导入 GPG 公钥 0x34FA74DD:
 Userid: "NodeSource <gpg-rpm@nodesource.com>"
 指纹: 2E55 207A 95D9 944B 0CC9 3261 5DDB E8D4 34FA 74DD
 来自: /etc/pki/rpm-gpg/NODESOURCE-GPG-SIGNING-KEY-EL
导入公钥成功
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
Detected old npm client, removing...
  安装: nodejs-2:6.9.5-1nodesource.fc23.x86_64                                                                                                  1/1 
  验证: nodejs-2:6.9.5-1nodesource.fc23.x86_64                                                                                                  1/1 

已安装:
  nodejs.x86_64 2:6.9.5-1nodesource.fc23                                                                                                            

完毕！
```

Verify
```
[vagrant@localhost ~]$ node --version
v6.9.5
[vagrant@localhost ~]$ npm --version
3.10.10
```

Modules
```
[vagrant@localhost ~]$ ls /usr/lib/node_modules/
bower  grunt-cli  npm
```

Already existed local repository
```
[vagrant@localhost ~]$ ls .npm/
abbrev                 decompress-zip            hash.js                        merge-descriptors               semver
accepts                deep-equal                hawk                           merge-stream                    semver-diff
acorn                  deep-extend               he                             method-override                 semver-regex
acorn-dynamic-import   deep-is                   hmac-drbg                      methods                         semver-truncate
active-x-obfuscator    defaults                  hoek                           micromatch                      send
adm-zip                defined                   hooker                         miller-rabin                    sentence-case
after                  delayed-stream            hosted-git-info                mime                            serve-index
ajv                    depd                      htmlhint                       mime-db                         serve-static
ajv-keywords           des.js                    html-minifier                  mime-types                      set-blocking
align-text             destroy                   htmlparser2                    minijasminenode                 setimmediate
alter                  di                        http2                          minimalistic-assert             set-immediate-shim
amdefine               diff                      http-errors                    minimalistic-crypto-utils       setprototypeof
ansi-regex             diffie-hellman            http-proxy                     minimatch                       sha.js
ansi-styles            domain-browser            https-browserify               minimist                        shelljs
anymatch               domelementtype            http-signature                 mkdirp                          shell-quote
archive-type           domhandler                i                              mkpath                          sigmund
archy                  dom-serializer            ibrik                          morgan                          signal-exit
argparse               domutils                  iconv-lite                     mout                            simple-fmt
arraybuffer.slice      dot-case                  ieee754                        ms                              simple-is
array-differ           download                  imagemin                       multimatch                      snake-case
array-filter           download-status           imagemin-gifsicle              multiparty                      sntp
array-find-index       duplexer                  imagemin-jpegtran              multipipe                       socket.io
array-map              duplexer2                 imagemin-optipng               mute-stream                     socket.io-adapter
array-reduce           duplexify                 imagemin-pngquant              nan                             socket.io-client
array-union            each-async                imagemin-svgo                  natives                         socket.io-parser
array-uniq             ecc-jsbn                  image-size                     ncname                          source-list-map
array-unique           ee-first                  indent-string                  ncp                             source-map
arr-diff               elliptic                  indexof                        negotiator                      source-map-support
arr-flatten            emojis-list               inflight                       next-tick                       sparkles
arrify                 end-of-stream             inherits                       ng-annotate                     spawn-sync
asap                   engine.io                 ini                            node-libs-browser               spdx-correct
asn1                   engine.io-client          inquirer                       node-status-codes               spdx-exceptions
asn1.js                engine.io-parser          insight                        node-uuid                       spdx-expression-parse
assert                 enhanced-resolve          interpret                      nopt                            spdx-license-ids
assert-plus            entities                  intersect                      noptify                         split
async                  errno                     invert-kv                      normalize-package-data          sprintf-js
async-each             error-ex                  ip-regex                       normalize-path                  squeak
async-each-series      errorhandler              is-absolute                    npmconf                         sshpk
asynckit               es5-ext                   isarray                        npm-path                        stable
autoprefixer           es6-iterator              is-arrayish                    npm-run-path                    static-favicon
aws4                   es6-symbol                is-binary-path                 npm-which                       stat-mode
aws-sign2              es6-weak-map              is-buffer                      number-is-nan                   statuses
backo2                 escape-html               is-builtin-module              oauth-sign                      stream-browserify
balanced-match         escape-string-regexp      is-bzip2                       object-assign                   stream-combiner
base64-arraybuffer     escodegen                 is-dotfile                     object-component                stream-combiner2
base64id               escope                    is-equal-shallow               object.omit                     stream-counter
base64-js              esmangle                  isexe                          once                            stream-http
basic-auth             esprima                   is-extendable                  onetime                         stream-shift
basic-auth-connect     esshorten                 is-extglob                     on-finished                     string_decoder
batch                  estraverse                is-finite                      on-headers                      stringify-object
bcrypt-pbkdf           esutils                   is-fullwidth-code-point        open                            string-length
beeper                 etag                      is-gif                         opn                             stringmap
benchmark              event-emitter             is-glob                        optimist                        StringScanner
better-assert          eventemitter2             is-gzip                        optional                        stringset
big.js                 events                    is-integer                     optionator                      stringstream
binary                 evp_bytestokey            is-jpg                         options                         string-width
binary-extensions      exec-buffer               is-lower-case                  optipng-bin                     strip-ansi
bin-build              exec-series               is-my-json-valid               ordered-ast-traverse            strip-bom
bin-check              executable                is-natural-number              ordered-esprima-props           strip-bom-stream
bin-version            exit                      is-number                      ordered-read-streams            strip-dirs
bin-version-check      expand-brackets           is-obj                         os-browserify                   strip-indent
bin-wrapper            expand-range              isobject                       osenv                           strip-json-comments
bl                     express                   is-plain-obj                   os-filter-obj                   strip-outer
blob                   express-session           is-png                         os-homedir                      sum-up
bn.js                  extend                    is-posix-bracket               os-locale                       supports-color
body-parser            extend-shallow            is-primitive                   os-name                         svgo
boom                   extglob                   is-property                    os-shim                         swap-case
bower                  extract-zip               is-redirect                    os-tmpdir                       tapable
bower-config           extsprintf                is-relative                    osx-release                     tape
bower-endpoint-parser  fancy-log                 is-retry-allowed               package-json                    tar-fs
bower-json             fast-levenshtein          is-root                        pad-stdio                       tar-stream
bower-logger           faye-websocket            isstream                       pako                            tempfile
bower-registry-client  fd-slicer                 is-stream                      param-case                      text-table
brace-expansion        figures                   is-svg                         parse-asn1                      thread-sleep
braces                 filename-regex            istanbul                       parse-glob                      throttleit
brorand                filename-reserved-regex   is-tar                         parsejson                       through
browserify-aes         filenamify                is-typedarray                  parse-json                      through2
browserify-cipher      fileset                   is-upper-case                  parseqs                         through2-filter
browserify-des         file-sync-cmp             is-url                         parserlib                       timed-out
browserify-rsa         file-type                 is-utf8                        parseuri                        time-grunt
browserify-sign        fill-range                is-valid-glob                  parseurl                        timers-browserify
browserify-zlib        finalhandler              is-zip                         pascal-case                     timers-ext
buffer                 find-file                 jasmine                        path-browserify                 time-stamp
buffer-crc32           find-index                jasmine-beforeall              path-case                       tinycolor
buffers                find-up                   jasmine-core                   path-exists                     tiny-lr
buffer-shims           findup-sync               jasmine-spec-reporter          path-is-absolute                tiny-lr-fork
buffer-to-vinyl        find-versions             jasminewd                      path-key                        title-case
buffer-xor             first-chunk-stream        jasminewd2                     path-type                       tmp
builtin-modules        flopmang                  jodid25519                     pause                           to-absolute-glob
builtin-status-codes   forever-agent             jpegtran-bin                   pbkdf2                          to-array
bytes                  for-in                    jsbn                           peakle                          to-arraybuffer
callsite               form-data                 jshint                         pend                            touch
camelcase              for-own                   jshint-stylish                 phantomjs                       tough-cookie
camel-case             fresh                     json3                          pify                            traverse
camelcase-keys         fs-extra                  json5                          pinkie                          trim-newlines
caniuse-db             fs.realpath               jsonfile                       pinkie-promise                  trim-repeated
capture-stack-trace    fstream                   jsonify                        pkginfo                         tryor
cardinal               fstream-ignore            json-loader                    pkg-up                          try-thread-sleep
caseless               gaze                      jsonpointer                    pngquant-bin                    tty-browserify
caw                    generate-function         json-schema                    policyfile                      tunnel-agent
cdnjs-cdn-data         generate-object-property  json-stable-stringify          portscanner                     tweetnacl
center-align           get-caller-file           json-stringify-safe            postcss                         type-check
chainsaw               getobject                 jsprim                         prelude-ls                      typedarray
chalk                  getpass                   js-yaml                        prepend-http                    type-is
change-case            get-proxy                 junk                           preserve                        uglify-js
chmodr                 get-stdin                 karma                          pretty-bytes                    uglify-to-browserify
chokidar               gifsicle                  karma-coverage                 pretty-ms                       uid2
cipher-base            _git-remotes              karma-firefox-launcher         process                         uid-number
clap                   glob                      karma-jasmine                  process-nextick-args            ultron
clean-css              glob2base                 karma-ng-html2js-preprocessor  progress                        underscore
cli                    glob-base                 karma-phantomjs-launcher       promise                         underscore.string
cli-color              glob-parent               kew                            promptly                        unique-stream
cliui                  glob-stream               keypress                       propprop                        unpipe
clone                  globule                   kind-of                        proto-list                      untildify
clone-stats            glob-watcher              klaw                           protractor                      unzip-response
co                     glogg                     latest-version                 protractor-screenshot-reporter  update-notifier
coa                    google-cdn                lazy-cache                     prr                             upper-case
code-point-at          google-cdn-data           lazy-req                       pseudomap                       upper-case-first
coffee-script          got                       lazystream                     p-throttler                     uri-path
coffee-script-redux    graceful-fs               lcid                           pty.js                          url
color-convert          graceful-readlink         less                           public-encrypt                  url-parse-lax
colors                 grunt                     levn                           pump                            url-regex
combined-stream        grunt-angular-templates   livereload-js                  punycode                        useragent
commander              grunt-autoprefixer        loader-runner                  q                               user-home
component-bind         grunt-cli                 loader-utils                   qs                              utf8
component-emitter      grunt-concurrent          load-grunt-tasks               querystring                     util
component-inherit      grunt-contrib-clean       load-json-file                 querystring-es3                 util-deprecate
compressible           grunt-contrib-concat      lockfile                       randomatic                      utile
compression            grunt-contrib-connect     _locks                         randombytes                     utils-merge
concat-map             grunt-contrib-copy        lodash                         range-parser                    uuid
concat-stream          grunt-contrib-cssmin      lodash._basecopy               raw-body                        vali-date
config-chain           grunt-contrib-htmlmin     lodash._basetostring           rc                              validate-npm-package-license
configstore            grunt-contrib-imagemin    lodash._basevalues             read                            verror
connect                grunt-contrib-jshint      lodash.debounce                readable-stream                 vhost
connect-livereload     grunt-contrib-less        lodash.escape                  read-all-stream                 vinyl
connect-modrewrite     grunt-contrib-uglify      lodash._getnative              readdirp                        vinyl-assign
connect-timeout        grunt-contrib-watch       lodash.isarguments             readline2                       vinyl-fs
console-browserify     grunt-filerev             lodash.isarray                 read-pkg                        vm-browserify
console-stream         grunt-git                 lodash.isequal                 read-pkg-up                     ware
constant-case          grunt-google-cdn          lodash.isfunction              redent                          watchpack
constants-browserify   grunt-htmlhint            lodash._isiterateecall         redeyed                         webpack
content-type           grunt-istanbul-coverage   lodash._isnative               redis                           webpack-sources
convert-source-map     grunt-karma               lodash.isobject                regex-cache                     websocket-driver
cookie                 grunt-known-options       lodash.keys                    regexp-quote                    websocket-extensions
cookie-parser          grunt-legacy-log          lodash.now                     registry.npmjs.org              whet.extend
cookie-signature       grunt-legacy-log-utils    lodash._objecttypes            registry-url                    which
core-util-is           grunt-legacy-util         lodash._reescape               relateurl                       which-module
create-ecdh            grunt-mkdir               lodash._reevaluate             remove-trailing-separator       window-size
create-error-class     grunt-newer               lodash._reinterpolate          repeat-element                  win-release
create-hash            grunt-ng-annotate         lodash.restparam               repeating                       wiredep
create-hmac            grunt-protractor-runner   lodash._root                   repeat-string                   wiredep-cli
cryptiles              grunt-shell               lodash.template                replace-ext                     wordwrap
crypto-browserify      grunt-svgmin              lodash.templatesettings        request                         wrap-ansi
cscodegen              grunt-usemin              log4js                         request-progress                wrap-fn
csslint                grunt-wiredep             logalot                        request-replay                  wrappy
csso                   gulp-decompress           log-symbols                    require-directory               ws
csurf                  gulplog                   longest                        require-main-filename           xdg-basedir
ctype                  gulp-rename               loud-rejection                 resolve                         xml2js
currently-unhandled    gulp-sourcemaps           lower-case                     resolve-from                    xmlbuilder
d                      gulp-util                 lower-case-first               resolve-pkg                     xml-char-classes
dashdash               gzip-size                 lowercase-keys                 response-time                   xmlhttprequest
dateformat             handlebars                lpad                           retry                           xmlhttprequest-ssl
date-now               har-validator             lpad-align                     right-align                     xtend
date-time              has-ansi                  lru-cache                      rimraf                          y18n
debug                  has-binary                lru-queue                      ripemd160                       yallist
decamelize             has-color                 map-obj                        rx                              yargs
decompress             has-cors                  maxmin                         safe-buffer                     yargs-parser
decompress-tar         has-flag                  media-typer                    saucelabs                       yauzl
decompress-tarbz2      has-gulplog               memoizee                       sax                             yeast
decompress-targz       hasha                     memory-fs                      seek-bzip                       zeparser
decompress-unzip       hash-base                 meow                           selenium-webdriver              zlib-browserify
```

## Grunt

install
```
[vagrant@localhost ~]$ sudo npm install -g grunt-cli bower
npm WARN deprecated bower@1.8.0: ..psst! While Bower is maintained, we recommend Yarn and Webpack for *new* front-end projects! Yarn's advantage is security and reliability, and Webpack's is support for both CommonJS and AMD projects. Currently there's no migration path but we hope you'll help us figure out one.
npm WARN gentlyRm not removing /usr/bin/grunt as it wasn't installed by /usr/lib/node_modules/grunt-cli
npm WARN gentlyRm not removing /usr/bin/bower as it wasn't installed by /usr/lib/node_modules/bower
/usr/bin/bower -> /usr/lib/node_modules/bower/bin/bower
/usr/bin/grunt -> /usr/lib/node_modules/grunt-cli/bin/grunt
/usr/lib
├── bower@1.8.0 
└─┬ grunt-cli@1.2.0 
  ├─┬ findup-sync@0.3.0
  │ └─┬ glob@5.0.15
  │   ├── inflight@1.0.6 
  │   ├─┬ minimatch@3.0.4 
  │   │ └── brace-expansion@1.1.7 
  │   └── path-is-absolute@1.0.1 
  └─┬ nopt@3.0.6
    └── abbrev@1.1.0 

```

Verify
```
[vagrant@localhost ~]$ grunt --version
grunt-cli v1.2.0
[vagrant@localhost ~]$ bower --version
1.8.0
[vagrant@localhost ~]$ ls /usr/lib/node_modules/
bower  grunt-cli  gulp  gulp-cli  npm
```

## Gulp

Gulp-cli
```
[vagrant@localhost ~]$ sudo npm install --global gulp-cli
/usr/bin/gulp -> /usr/lib/node_modules/gulp-cli/bin/gulp.js
/usr/lib
└─┬ gulp-cli@1.3.0 
  ├── archy@1.0.0 
  ├─┬ chalk@1.1.3 
  │ ├── ansi-styles@2.2.1 
  │ ├── escape-string-regexp@1.0.5 
  │ ├─┬ has-ansi@2.0.0 
  │ │ └── ansi-regex@2.1.1 
  │ ├── strip-ansi@3.0.1 
  │ └── supports-color@2.0.0 
  ├─┬ copy-props@1.6.0 
  │ ├─┬ each-props@1.3.0 
  │ │ └── object-assign@4.1.1 
  │ └─┬ is-plain-object@2.0.3 
  │   └── isobject@3.0.0 
  ├─┬ fancy-log@1.3.0 
  │ └── time-stamp@1.1.0 
  ├─┬ gulplog@1.0.0 
  │ └─┬ glogg@1.0.0 
  │   └── sparkles@1.0.0 
  ├── interpret@1.0.3 
  ├─┬ liftoff@2.3.0 
  │ ├── extend@3.0.1 
  │ ├─┬ findup-sync@0.4.3 
  │ │ ├─┬ detect-file@0.1.0 
  │ │ │ └── fs-exists-sync@0.1.0 
  │ │ ├── is-glob@2.0.1 
  │ │ └─┬ resolve-dir@0.1.1 
  │ │   └─┬ global-modules@0.2.3 
  │ │     ├─┬ global-prefix@0.1.5 
  │ │     │ ├─┬ homedir-polyfill@1.0.1 
  │ │     │ │ └── parse-passwd@1.0.0 
  │ │     │ ├── ini@1.3.4 
  │ │     │ └─┬ which@1.2.14 
  │ │     │   └── isexe@2.0.0 
  │ │     └── is-windows@0.2.0 
  │ ├─┬ fined@1.0.2 
  │ │ ├── expand-tilde@1.2.2 
  │ │ ├── lodash.assignwith@4.2.0 
  │ │ ├── lodash.isempty@4.4.0 
  │ │ ├── lodash.pick@4.4.0 
  │ │ └─┬ parse-filepath@1.0.1 
  │ │   ├─┬ is-absolute@0.2.6 
  │ │   │ └─┬ is-relative@0.2.1 
  │ │   │   └─┬ is-unc-path@0.1.2 
  │ │   │     └── unc-path-regex@0.1.2 
  │ │   ├── map-cache@0.2.2 
  │ │   └─┬ path-root@0.1.1 
  │ │     └── path-root-regex@0.1.2 
  │ ├── flagged-respawn@0.3.2 
  │ ├── lodash.isstring@4.0.1 
  │ ├── lodash.mapvalues@4.6.0 
  │ ├── rechoir@0.6.2 
  │ └─┬ resolve@1.3.3 
  │   └── path-parse@1.0.5 
  ├── lodash.isfunction@3.0.8 
  ├── lodash.isplainobject@4.0.6 
  ├── lodash.sortby@4.7.0 
  ├─┬ matchdep@1.0.1 
  │ ├─┬ findup-sync@0.3.0 
  │ │ └─┬ glob@5.0.15 
  │ │   ├─┬ inflight@1.0.6 
  │ │   │ └── wrappy@1.0.2 
  │ │   ├── inherits@2.0.3 
  │ │   ├─┬ minimatch@3.0.4 
  │ │   │ └─┬ brace-expansion@1.1.7 
  │ │   │   ├── balanced-match@0.4.2 
  │ │   │   └── concat-map@0.0.1 
  │ │   ├── once@1.4.0 
  │ │   └── path-is-absolute@1.0.1 
  │ ├─┬ micromatch@2.3.11 
  │ │ ├─┬ arr-diff@2.0.0 
  │ │ │ └── arr-flatten@1.0.3 
  │ │ ├── array-unique@0.2.1 
  │ │ ├─┬ braces@1.8.5 
  │ │ │ ├─┬ expand-range@1.8.2 
  │ │ │ │ └─┬ fill-range@2.2.3 
  │ │ │ │   ├── is-number@2.1.0 
  │ │ │ │   ├─┬ isobject@2.1.0 
  │ │ │ │   │ └── isarray@1.0.0 
  │ │ │ │   ├── randomatic@1.1.6 
  │ │ │ │   └── repeat-string@1.6.1 
  │ │ │ ├── preserve@0.2.0 
  │ │ │ └── repeat-element@1.1.2 
  │ │ ├─┬ expand-brackets@0.1.5 
  │ │ │ └── is-posix-bracket@0.1.1 
  │ │ ├── extglob@0.3.2 
  │ │ ├── filename-regex@2.0.1 
  │ │ ├── is-extglob@1.0.0 
  │ │ ├─┬ kind-of@3.2.2 
  │ │ │ └── is-buffer@1.1.5 
  │ │ ├─┬ normalize-path@2.1.1 
  │ │ │ └── remove-trailing-separator@1.0.2 
  │ │ ├─┬ object.omit@2.0.1 
  │ │ │ ├─┬ for-own@0.1.5 
  │ │ │ │ └── for-in@1.0.2 
  │ │ │ └── is-extendable@0.1.1 
  │ │ ├─┬ parse-glob@3.0.4 
  │ │ │ ├─┬ glob-base@0.3.0 
  │ │ │ │ └── glob-parent@2.0.0 
  │ │ │ └── is-dotfile@1.0.3 
  │ │ └─┬ regex-cache@0.4.3 
  │ │   ├── is-equal-shallow@0.1.3 
  │ │   └── is-primitive@2.0.0 
  │ ├── resolve@1.1.7 
  │ └── stack-trace@0.0.9 
  ├── mute-stdout@1.0.0 
  ├── pretty-hrtime@1.0.3 
  ├─┬ semver-greatest-satisfied-range@1.0.0 
  │ ├── semver@4.3.6 
  │ └── semver-regex@1.0.0 
  ├─┬ tildify@1.2.0 
  │ └── os-homedir@1.0.2 
  ├─┬ v8flags@2.1.1 
  │ └── user-home@1.1.1 
  ├─┬ wreck@6.3.0 
  │ ├── boom@2.10.1 
  │ └── hoek@2.16.3 
  └─┬ yargs@3.32.0 
    ├── camelcase@2.1.1 
    ├─┬ cliui@3.2.0 
    │ └── wrap-ansi@2.1.0 
    ├── decamelize@1.2.0 
    ├─┬ os-locale@1.4.0 
    │ └─┬ lcid@1.0.0 
    │   └── invert-kv@1.0.0 
    ├─┬ string-width@1.0.2 
    │ ├── code-point-at@1.1.0 
    │ └─┬ is-fullwidth-code-point@1.0.0 
    │   └── number-is-nan@1.0.1 
    ├── window-size@0.1.4 
    └── y18n@3.2.1 
```

Gulp
```
[vagrant@localhost ~]$ sudo npm install --global gulp
npm WARN deprecated minimatch@2.0.10: Please update to minimatch 3.0.2 or higher to avoid a RegExp DoS issue
npm WARN deprecated minimatch@0.2.14: Please update to minimatch 3.0.2 or higher to avoid a RegExp DoS issue
npm WARN deprecated graceful-fs@1.2.3: graceful-fs v3.0.0 and before will fail on node releases >= v7.0. Please update to graceful-fs@^4.0.0 as soon as possible. Use 'npm ls graceful-fs' to find it in the tree.
/usr/bin/gulp -> /usr/lib/node_modules/gulp/bin/gulp.js
/usr/lib
└─┬ gulp@3.9.1 
  ├── archy@1.0.0 
  ├─┬ chalk@1.1.3 
  │ ├── ansi-styles@2.2.1 
  │ ├── escape-string-regexp@1.0.5 
  │ ├─┬ has-ansi@2.0.0 
  │ │ └── ansi-regex@2.1.1 
  │ ├── strip-ansi@3.0.1 
  │ └── supports-color@2.0.0 
  ├── deprecated@0.0.1 
  ├─┬ gulp-util@3.0.8 
  │ ├── array-differ@1.0.0 
  │ ├── array-uniq@1.0.3 
  │ ├── beeper@1.1.1 
  │ ├── dateformat@2.0.0 
  │ ├─┬ fancy-log@1.3.0 
  │ │ └── time-stamp@1.1.0 
  │ ├─┬ gulplog@1.0.0 
  │ │ └── glogg@1.0.0 
  │ ├─┬ has-gulplog@0.1.0 
  │ │ └── sparkles@1.0.0 
  │ ├── lodash._reescape@3.0.0 
  │ ├── lodash._reevaluate@3.0.0 
  │ ├── lodash._reinterpolate@3.0.0 
  │ ├─┬ lodash.template@3.6.2 
  │ │ ├── lodash._basecopy@3.0.1 
  │ │ ├── lodash._basetostring@3.0.1 
  │ │ ├── lodash._basevalues@3.0.0 
  │ │ ├── lodash._isiterateecall@3.0.9 
  │ │ ├─┬ lodash.escape@3.2.0 
  │ │ │ └── lodash._root@3.0.1 
  │ │ ├─┬ lodash.keys@3.1.2 
  │ │ │ ├── lodash._getnative@3.9.1 
  │ │ │ ├── lodash.isarguments@3.1.0 
  │ │ │ └── lodash.isarray@3.0.4 
  │ │ ├── lodash.restparam@3.6.1 
  │ │ └── lodash.templatesettings@3.1.1 
  │ ├─┬ multipipe@0.1.2 
  │ │ └─┬ duplexer2@0.0.2 
  │ │   └── readable-stream@1.1.14 
  │ ├── object-assign@3.0.0 
  │ ├── replace-ext@0.0.1 
  │ ├─┬ through2@2.0.3 
  │ │ ├─┬ readable-stream@2.2.11 
  │ │ │ ├── core-util-is@1.0.2 
  │ │ │ ├── inherits@2.0.3 
  │ │ │ ├── isarray@1.0.0 
  │ │ │ ├── process-nextick-args@1.0.7 
  │ │ │ ├── safe-buffer@5.0.1 
  │ │ │ ├── string_decoder@1.0.2 
  │ │ │ └── util-deprecate@1.0.2 
  │ │ └── xtend@4.0.1 
  │ └─┬ vinyl@0.5.3 
  │   ├── clone@1.0.2 
  │   └── clone-stats@0.0.1 
  ├── interpret@1.0.3 
  ├─┬ liftoff@2.3.0 
  │ ├── extend@3.0.1 
  │ ├─┬ findup-sync@0.4.3 
  │ │ ├─┬ detect-file@0.1.0 
  │ │ │ └── fs-exists-sync@0.1.0 
  │ │ ├─┬ is-glob@2.0.1 
  │ │ │ └── is-extglob@1.0.0 
  │ │ ├─┬ micromatch@2.3.11 
  │ │ │ ├─┬ arr-diff@2.0.0 
  │ │ │ │ └── arr-flatten@1.0.3 
  │ │ │ ├── array-unique@0.2.1 
  │ │ │ ├─┬ braces@1.8.5 
  │ │ │ │ ├─┬ expand-range@1.8.2 
  │ │ │ │ │ └─┬ fill-range@2.2.3 
  │ │ │ │ │   ├── is-number@2.1.0 
  │ │ │ │ │   ├─┬ isobject@2.1.0 
  │ │ │ │ │   │ └── isarray@1.0.0 
  │ │ │ │ │   ├── randomatic@1.1.6 
  │ │ │ │ │   └── repeat-string@1.6.1 
  │ │ │ │ ├── preserve@0.2.0 
  │ │ │ │ └── repeat-element@1.1.2 
  │ │ │ ├─┬ expand-brackets@0.1.5 
  │ │ │ │ └── is-posix-bracket@0.1.1 
  │ │ │ ├── extglob@0.3.2 
  │ │ │ ├── filename-regex@2.0.1 
  │ │ │ ├─┬ kind-of@3.2.2 
  │ │ │ │ └── is-buffer@1.1.5 
  │ │ │ ├─┬ normalize-path@2.1.1 
  │ │ │ │ └── remove-trailing-separator@1.0.2 
  │ │ │ ├─┬ object.omit@2.0.1 
  │ │ │ │ ├─┬ for-own@0.1.5 
  │ │ │ │ │ └── for-in@1.0.2 
  │ │ │ │ └── is-extendable@0.1.1 
  │ │ │ ├─┬ parse-glob@3.0.4 
  │ │ │ │ ├─┬ glob-base@0.3.0 
  │ │ │ │ │ └── glob-parent@2.0.0 
  │ │ │ │ └── is-dotfile@1.0.3 
  │ │ │ └─┬ regex-cache@0.4.3 
  │ │ │   ├── is-equal-shallow@0.1.3 
  │ │ │   └── is-primitive@2.0.0 
  │ │ └─┬ resolve-dir@0.1.1 
  │ │   └─┬ global-modules@0.2.3 
  │ │     ├─┬ global-prefix@0.1.5 
  │ │     │ ├─┬ homedir-polyfill@1.0.1 
  │ │     │ │ └── parse-passwd@1.0.0 
  │ │     │ ├── ini@1.3.4 
  │ │     │ └─┬ which@1.2.14 
  │ │     │   └── isexe@2.0.0 
  │ │     └── is-windows@0.2.0 
  │ ├─┬ fined@1.0.2 
  │ │ ├── expand-tilde@1.2.2 
  │ │ ├── lodash.assignwith@4.2.0 
  │ │ ├── lodash.isempty@4.4.0 
  │ │ ├── lodash.pick@4.4.0 
  │ │ └─┬ parse-filepath@1.0.1 
  │ │   ├─┬ is-absolute@0.2.6 
  │ │   │ └─┬ is-relative@0.2.1 
  │ │   │   └─┬ is-unc-path@0.1.2 
  │ │   │     └── unc-path-regex@0.1.2 
  │ │   ├── map-cache@0.2.2 
  │ │   └─┬ path-root@0.1.1 
  │ │     └── path-root-regex@0.1.2 
  │ ├── flagged-respawn@0.3.2 
  │ ├── lodash.isplainobject@4.0.6 
  │ ├── lodash.isstring@4.0.1 
  │ ├── lodash.mapvalues@4.6.0 
  │ ├── rechoir@0.6.2 
  │ └─┬ resolve@1.3.3 
  │   └── path-parse@1.0.5 
  ├── minimist@1.2.0 
  ├─┬ orchestrator@0.3.8 
  │ ├─┬ end-of-stream@0.1.5 
  │ │ └─┬ once@1.3.3 
  │ │   └── wrappy@1.0.2 
  │ ├── sequencify@0.0.7 
  │ └── stream-consume@0.1.0 
  ├── pretty-hrtime@1.0.3 
  ├── semver@4.3.6 
  ├─┬ tildify@1.2.0 
  │ └── os-homedir@1.0.2 
  ├─┬ v8flags@2.1.1 
  │ └── user-home@1.1.1 
  └─┬ vinyl-fs@0.3.14 
    ├── defaults@1.0.3 
    ├─┬ glob-stream@3.1.18 
    │ ├─┬ glob@4.5.3 
    │ │ └── inflight@1.0.6 
    │ ├─┬ glob2base@0.0.12 
    │ │ └── find-index@0.1.1 
    │ ├─┬ minimatch@2.0.10 
    │ │ └─┬ brace-expansion@1.1.7 
    │ │   ├── balanced-match@0.4.2 
    │ │   └── concat-map@0.0.1 
    │ ├── ordered-read-streams@0.1.0 
    │ ├─┬ through2@0.6.5 
    │ │ └── readable-stream@1.0.34 
    │ └── unique-stream@1.0.0 
    ├─┬ glob-watcher@0.0.6 
    │ └─┬ gaze@0.5.2 
    │   └─┬ globule@0.1.0 
    │     ├─┬ glob@3.1.21 
    │     │ ├── graceful-fs@1.2.3 
    │     │ └── inherits@1.0.2 
    │     ├── lodash@1.0.2 
    │     └─┬ minimatch@0.2.14 
    │       ├── lru-cache@2.7.3 
    │       └── sigmund@1.0.1 
    ├─┬ graceful-fs@3.0.11 
    │ └── natives@1.1.0 
    ├─┬ mkdirp@0.5.1 
    │ └── minimist@0.0.8 
    ├─┬ strip-bom@1.0.0 
    │ ├── first-chunk-stream@1.0.0 
    │ └── is-utf8@0.2.1 
    ├─┬ through2@0.6.5 
    │ └─┬ readable-stream@1.0.34 
    │   ├── isarray@0.0.1 
    │   └── string_decoder@0.10.31 
    └─┬ vinyl@0.4.6 
      └── clone@0.2.0 

```

Verify
```
[vagrant@localhost ~]$ gulp -v
[18:35:34] CLI version 3.9.1
[vagrant@localhost ~]$ ls /usr/lib/node_modules/
bower  grunt-cli  gulp  gulp-cli  npm
```

## Webpack 

### Source

clone (Mac)
```
fanhonglingdeMacBook-Pro:github.com fanhongling$ git clone https://github.com/webpack/webpack webpack/webpack
Cloning into 'webpack/webpack'...
remote: Counting objects: 28733, done.
remote: Compressing objects: 100% (43/43), done.
remote: Total 28733 (delta 17), reused 45 (delta 11), pack-reused 28675
Receiving objects: 100% (28733/28733), 7.52 MiB | 138.00 KiB/s, done.
Resolving deltas: 100% (18752/18752), done.
Checking connectivity... done.
```

### Package

install
```
[vagrant@localhost go-starter-kit]$ npm install --save-dev webpack
go-starter-kit@0.1.0 /Users/fanhongling/Downloads/go-kubernetes/src/github.com/olebedev/go-starter-kit
└─┬ webpack@1.15.0 
  ├─┬ enhanced-resolve@0.9.1 
  │ └── memory-fs@0.2.0 
  ├── interpret@0.6.6 
  ├── loader-utils@0.2.17 
  ├── memory-fs@0.3.0 
  ├─┬ node-libs-browser@0.7.0 
  │ ├── assert@1.4.1 
  │ ├─┬ browserify-zlib@0.1.4 
  │ │ └── pako@0.2.9 
  │ ├─┬ buffer@4.9.1 
  │ │ ├── base64-js@1.2.0 
  │ │ └── ieee754@1.1.8 
  │ ├─┬ console-browserify@1.1.0 
  │ │ └── date-now@0.1.4 
  │ ├── constants-browserify@1.0.0 
  │ ├─┬ crypto-browserify@3.3.0 
  │ │ ├── browserify-aes@0.4.0 
  │ │ ├── pbkdf2-compat@2.0.1 
  │ │ ├── ripemd160@0.2.0 
  │ │ └── sha.js@2.2.6 
  │ ├── domain-browser@1.1.7 
  │ ├── events@1.1.1 
  │ ├── https-browserify@0.0.1 
  │ ├── os-browserify@0.2.1 
  │ ├── path-browserify@0.0.0 
  │ ├── process@0.11.10 
  │ ├── punycode@1.4.1 
  │ ├── querystring-es3@0.2.1 
  │ ├── stream-browserify@2.0.1 
  │ ├─┬ stream-http@2.7.1 
  │ │ ├── builtin-status-codes@3.0.0 
  │ │ └── to-arraybuffer@1.0.1 
  │ ├── string_decoder@0.10.31 
  │ ├── timers-browserify@2.0.2 
  │ ├── tty-browserify@0.0.0 
  │ ├─┬ url@0.11.0 
  │ │ └── punycode@1.3.2 
  │ ├─┬ util@0.10.3 
  │ │ └── inherits@2.0.1 
  │ └─┬ vm-browserify@0.0.4 
  │   └── indexof@0.0.1 
  ├─┬ optimist@0.6.1 
  │ └── wordwrap@0.0.3 
  ├── tapable@0.1.10 
  ├─┬ uglify-js@2.7.5 
  │ ├── async@0.2.10 
  │ ├── uglify-to-browserify@1.0.2 
  │ └─┬ yargs@3.10.0 
  │   ├── camelcase@1.2.1 
  │   ├─┬ cliui@2.1.0 
  │   │ ├─┬ center-align@0.1.3 
  │   │ │ ├─┬ align-text@0.1.4 
  │   │ │ │ ├── longest@1.0.1 
  │   │ │ │ └── repeat-string@1.6.1 
  │   │ │ └── lazy-cache@1.0.4 
  │   │ ├── right-align@0.1.3 
  │   │ └── wordwrap@0.0.2 
  │   └── window-size@0.1.0 
  ├─┬ watchpack@0.2.9 
  │ ├── async@0.9.2 
  │ └─┬ chokidar@1.7.0 
  │   ├─┬ anymatch@1.3.0 
  │   │ └─┬ micromatch@2.3.11 
  │   │   ├─┬ arr-diff@2.0.0 
  │   │   │ └── arr-flatten@1.0.3 
  │   │   ├── array-unique@0.2.1 
  │   │   ├─┬ braces@1.8.5 
  │   │   │ ├─┬ expand-range@1.8.2 
  │   │   │ │ └─┬ fill-range@2.2.3 
  │   │   │ │   ├── is-number@2.1.0 
  │   │   │ │   ├── isobject@2.1.0 
  │   │   │ │   └─┬ randomatic@1.1.7 
  │   │   │ │     ├─┬ is-number@3.0.0 
  │   │   │ │     │ └── kind-of@3.2.2 
  │   │   │ │     └── kind-of@4.0.0 
  │   │   │ ├── preserve@0.2.0 
  │   │   │ └── repeat-element@1.1.2 
  │   │   ├─┬ expand-brackets@0.1.5 
  │   │   │ └── is-posix-bracket@0.1.1 
  │   │   ├── extglob@0.3.2 
  │   │   ├── filename-regex@2.0.1 
  │   │   ├─┬ kind-of@3.2.2 
  │   │   │ └── is-buffer@1.1.5 
  │   │   ├─┬ normalize-path@2.1.1 
  │   │   │ └── remove-trailing-separator@1.0.2 
  │   │   ├─┬ object.omit@2.0.1 
  │   │   │ ├─┬ for-own@0.1.5 
  │   │   │ │ └── for-in@1.0.2 
  │   │   │ └── is-extendable@0.1.1 
  │   │   ├─┬ parse-glob@3.0.4 
  │   │   │ ├── glob-base@0.3.0 
  │   │   │ └── is-dotfile@1.0.3 
  │   │   └─┬ regex-cache@0.4.3 
  │   │     ├── is-equal-shallow@0.1.3 
  │   │     └── is-primitive@2.0.0 
  │   ├── async-each@1.0.1 
  │   ├── glob-parent@2.0.0 
  │   ├─┬ is-binary-path@1.0.1 
  │   │ └── binary-extensions@1.8.0 
  │   ├─┬ is-glob@2.0.1 
  │   │ └── is-extglob@1.0.0 
  │   └─┬ readdirp@2.1.0 
  │     └── set-immediate-shim@1.0.1 
  └─┬ webpack-core@0.6.9 
    └─┬ source-map@0.4.4 
      └── amdefine@1.0.1 

npm WARN optional SKIPPING OPTIONAL DEPENDENCY: fsevents@^1.0.0 (node_modules/chokidar/node_modules/fsevents):
npm WARN notsup SKIPPING OPTIONAL DEPENDENCY: Unsupported platform for fsevents@1.1.1: wanted {"os":"darwin","arch":"any"} (current: {"os":"linux","arch":"x64"})
```

install
```
[vagrant@localhost web-starter-kit]$ npm install --save-dev webpack
- acorn@4.0.13 node_modules/acorn-dynamic-import/node_modules/acorn
- ansi-regex@2.1.1 node_modules/ansi-regex
- bn.js@4.11.6 node_modules/bn.js
- brorand@1.1.0 node_modules/brorand
- buffer-xor@1.0.3 node_modules/buffer-xor
- builtin-modules@1.1.1 node_modules/builtin-modules
- cipher-base@1.0.3 node_modules/cipher-base
- co@4.6.0 node_modules/co
- code-point-at@1.1.0 node_modules/code-point-at
- create-hash@1.1.3 node_modules/create-hash
- create-hmac@1.1.6 node_modules/create-hmac
- evp_bytestokey@1.0.0 node_modules/evp_bytestokey
- get-caller-file@1.0.2 node_modules/get-caller-file
- hash-base@2.0.2 node_modules/hash-base
- hash.js@1.0.3 node_modules/hash.js
- hosted-git-info@2.4.2 node_modules/hosted-git-info
- invert-kv@1.0.0 node_modules/invert-kv
- is-arrayish@0.2.1 node_modules/is-arrayish
- error-ex@1.3.1 node_modules/error-ex
- is-builtin-module@1.0.0 node_modules/is-builtin-module
- is-utf8@0.2.1 node_modules/is-utf8
- jsonify@0.0.0 node_modules/jsonify
- json-stable-stringify@1.0.1 node_modules/json-stable-stringify
- lcid@1.0.0 node_modules/lcid
- miller-rabin@4.0.0 node_modules/miller-rabin
- minimalistic-assert@1.0.0 node_modules/minimalistic-assert
- asn1.js@4.9.1 node_modules/asn1.js
- des.js@1.0.0 node_modules/des.js
- browserify-des@1.0.0 node_modules/browserify-des
- browserify-cipher@1.0.0 node_modules/browserify-cipher
- minimalistic-crypto-utils@1.0.1 node_modules/minimalistic-crypto-utils
- hmac-drbg@1.0.1 node_modules/hmac-drbg
- elliptic@6.4.0 node_modules/elliptic
- create-ecdh@4.0.0 node_modules/create-ecdh
- number-is-nan@1.0.1 node_modules/number-is-nan
- is-fullwidth-code-point@1.0.0 node_modules/is-fullwidth-code-point
- os-locale@1.4.0 node_modules/os-locale
- parse-json@2.2.0 node_modules/parse-json
- pbkdf2@3.0.12 node_modules/pbkdf2
- parse-asn1@5.1.0 node_modules/parse-asn1
- pify@2.3.0 node_modules/pify
- pinkie@2.0.4 node_modules/pinkie
- pinkie-promise@2.0.1 node_modules/pinkie-promise
- path-exists@2.1.0 node_modules/path-exists
- find-up@1.1.2 node_modules/find-up
- path-type@1.1.0 node_modules/path-type
- safe-buffer@5.1.0 node_modules/randombytes/node_modules/safe-buffer
- randombytes@2.0.5 node_modules/randombytes
- browserify-rsa@4.0.1 node_modules/browserify-rsa
- browserify-sign@4.0.4 node_modules/browserify-sign
- diffie-hellman@5.0.2 node_modules/diffie-hellman
- public-encrypt@4.0.0 node_modules/public-encrypt
- require-directory@2.1.1 node_modules/require-directory
- require-main-filename@1.0.1 node_modules/require-main-filename
- semver@5.3.0 node_modules/semver
- set-blocking@2.0.0 node_modules/set-blocking
- spdx-expression-parse@1.0.4 node_modules/spdx-expression-parse
- spdx-license-ids@1.2.2 node_modules/spdx-license-ids
- spdx-correct@1.0.2 node_modules/spdx-correct
- strip-ansi@3.0.1 node_modules/strip-ansi
- string-width@1.0.2 node_modules/string-width
- strip-bom@2.0.0 node_modules/strip-bom
- load-json-file@1.1.0 node_modules/load-json-file
- yargs@3.10.0 node_modules/uglify-js/node_modules/yargs
- validate-npm-package-license@3.0.1 node_modules/validate-npm-package-license
- normalize-package-data@2.3.8 node_modules/normalize-package-data
- read-pkg@1.1.0 node_modules/read-pkg
- read-pkg-up@1.0.1 node_modules/read-pkg-up
- which-module@1.0.0 node_modules/which-module
- wrap-ansi@2.1.0 node_modules/wrap-ansi
- y18n@3.2.1 node_modules/y18n
- camelcase@3.0.0 node_modules/yargs-parser/node_modules/camelcase
- yargs-parser@4.2.1 node_modules/yargs-parser
- camelcase@3.0.0 node_modules/yargs/node_modules/camelcase
- cliui@3.2.0 node_modules/yargs/node_modules/cliui
- acorn-dynamic-import@2.0.2 node_modules/acorn-dynamic-import
- ajv@4.11.8 node_modules/ajv
- ajv-keywords@1.5.1 node_modules/ajv-keywords
- json-loader@0.5.4 node_modules/json-loader
- loader-runner@2.3.0 node_modules/loader-runner
- webpack-sources@0.2.3 node_modules/webpack-sources
web-starter-kit@0.1.0 /Users/fanhongling/Downloads/go-kubernetes/src/github.com/mijia/web-starter-kit
└─┬ webpack@1.15.0 
  ├── acorn@3.3.0 
  ├── async@1.5.2 
  ├── clone@1.0.2 
  ├─┬ enhanced-resolve@0.9.1 
  │ └── memory-fs@0.2.0 
  ├── interpret@0.6.6 
  ├── memory-fs@0.3.0 
  ├─┬ node-libs-browser@0.7.0 
  │ └─┬ crypto-browserify@3.3.0 
  │   ├── browserify-aes@0.4.0 
  │   ├── pbkdf2-compat@2.0.1 
  │   ├── ripemd160@0.2.0 
  │   └── sha.js@2.2.6 
  ├─┬ optimist@0.6.1 
  │ └── wordwrap@0.0.3 
  ├── tapable@0.1.10 
  ├─┬ uglify-js@2.7.5 
  │ ├── async@0.2.10 
  │ └─┬ yargs@3.10.0 
  │   └─┬ cliui@2.1.0
  │     └── wordwrap@0.0.2 
  ├─┬ watchpack@0.2.9 
  │ ├── async@0.9.2 
  │ └─┬ chokidar@1.7.0
  │   └─┬ anymatch@1.3.0
  │     └─┬ micromatch@2.3.11
  │       └─┬ braces@1.8.5
  │         └─┬ expand-range@1.8.2
  │           └─┬ fill-range@2.2.3
  │             └─┬ randomatic@1.1.7 
  │               ├─┬ is-number@3.0.0 
  │               │ └── kind-of@3.2.2 
  │               └── kind-of@4.0.0 
  └─┬ webpack-core@0.6.9 
    ├── source-list-map@0.1.8 
    └─┬ source-map@0.4.4 
      └── amdefine@1.0.1 

npm WARN optional SKIPPING OPTIONAL DEPENDENCY: fsevents@^1.0.0 (node_modules/chokidar/node_modules/fsevents):
npm WARN notsup SKIPPING OPTIONAL DEPENDENCY: Unsupported platform for fsevents@1.1.1: wanted {"os":"darwin","arch":"any"} (current: {"os":"linux","arch":"x64"})
npm WARN web-starter-kit@0.1.0 No repository field.
```

## yarn

### repo

repo
```
[vagrant@localhost ~]$ sudo wget https://dl.yarnpkg.com/rpm/yarn.repo -O /etc/yum.repos.d/yarn.repo
--2017-06-11 23:17:31--  https://dl.yarnpkg.com/rpm/yarn.repo
正在解析主机 dl.yarnpkg.com (dl.yarnpkg.com)... 104.16.62.173, 104.16.63.173, 104.16.61.173, ...
正在连接 dl.yarnpkg.com (dl.yarnpkg.com)|104.16.62.173|:443... 已连接。
已发出 HTTP 请求，正在等待回应... 200 OK
长度：130 [application/octet-stream]
正在保存至: “/etc/yum.repos.d/yarn.repo”

/etc/yum.repos.d/yarn.repo            100%[=========================================================================>]     130  --.-KB/s    in 0s      

2017-06-11 23:17:32 (9.26 MB/s) - 已保存 “/etc/yum.repos.d/yarn.repo” [130/130])
```

Packages
```
[vagrant@localhost ~]$ sudo dnf list --showduplicates yarn
上次元数据过期检查在 0:00:24 前执行于 Sun Jun 11 23:19:53 2017。
可安装的软件包
yarn.noarch                                                                0.15.0-1                                                                 yarn
yarn.noarch                                                                0.16.0-1                                                                 yarn
yarn.noarch                                                                0.16.1-1                                                                 yarn
yarn.noarch                                                                0.17.6-1                                                                 yarn
yarn.noarch                                                                0.17.8-1                                                                 yarn
yarn.noarch                                                                0.17.9-1                                                                 yarn
yarn.noarch                                                                0.17.10-1                                                                yarn
yarn.noarch                                                                0.18.1-1                                                                 yarn
yarn.noarch                                                                0.19.1-1                                                                 yarn
yarn.noarch                                                                0.20.3-1                                                                 yarn
yarn.noarch                                                                0.21.3-1                                                                 yarn
yarn.noarch                                                                0.22.0-1                                                                 yarn
yarn.noarch                                                                0.23.2-1                                                                 yarn
yarn.noarch                                                                0.23.3-1                                                                 yarn
yarn.noarch                                                                0.23.4-1                                                                 yarn
yarn.noarch                                                                0.24.3-1                                                                 yarn
yarn.noarch                                                                0.24.4-1                                                                 yarn
yarn.noarch                                                                0.24.5-1                                                                 yarn
yarn.noarch                                                                0.24.6-1                                                                 yarn
```

install
```
[vagrant@localhost ~]$ sudo dnf install -y yarn
上次元数据过期检查在 0:00:21 前执行于 Sun Jun 11 23:21:21 2017。
依赖关系解决。
========================================================================================================================================================
 Package                           架构                                版本                                     仓库                               大小
========================================================================================================================================================
安装:
 yarn                              noarch                              0.24.6-1                                 yarn                              758 k

事务概要
========================================================================================================================================================
安装  1 Package

总下载：758 k
安装大小：3.4 M
下载软件包：
yarn-0.24.6-1.noarch.rpm                                                                                                 95 kB/s | 758 kB     00:07    
--------------------------------------------------------------------------------------------------------------------------------------------------------
总计                                                                                                                     95 kB/s | 758 kB     00:07     
警告：/var/cache/dnf/yarn-39eb54bf733530c3/packages/yarn-0.24.6-1.noarch.rpm: 头V4 RSA/SHA256 Signature, 密钥 ID 6963f07f: NOKEY
导入 GPG 公钥 0x6963F07F:
 Userid: "Yarn RPM Packaging <yarn@dan.cx>"
 指纹: 9A6F 73F3 4BEB 7473 4D8C 6914 9CBB B558 6963 F07F
 来自: https://dl.yarnpkg.com/rpm/pubkey.gpg
导入公钥成功
运行事务检查
事务检查成功。
运行事务测试
事务测试成功。
运行事务
  安装: yarn-0.24.6-1.noarch                                                                                                                        1/1 
  验证: yarn-0.24.6-1.noarch                                                                                                                        1/1 

已安装:
  yarn.noarch 0.24.6-1                                                                                                                                  

完毕！
[vagrant@localhost ~]$ yarn --version
0.24.6
```

### Golang-Webpack

__Search__

* https://golanglibs.com/top?q=webpack

__List__

* https://github.com/huoy/moon
* https://github.com/mijia/web-starter-kit

__Other__

* https://github.com/olebedev/go-starter-kit