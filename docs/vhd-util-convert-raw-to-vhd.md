
https://www.citrix.com/blogs/2012/10/04/convert-a-raw-image-to-xenserver-vhd/
https://stackoverflow.com/questions/14436184/how-to-install-as86-in-debian-6-0
https://github.com/xapi-project/blktap/blob/master/vhd/vhd-util.c


vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen$ curl -jkSLO https://downloads.xenproject.org/release/xen/4.2.0/xen-4.2.0.tar.gz
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
 62 14.8M   62 9568k    0     0   110k      0  0:02:17  0:01:26  0:00:51  128k



vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ curl -jkSL https://github.com/citrix-openstack/xenserver-utils/raw/master/blktap2.patch -o - | patch -p0
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   153  100   153    0     0    132      0  0:00:01  0:00:01 --:--:--   132
100 29932  100 29932    0     0   9012      0  0:00:03  0:00:03 --:--:-- 25539
patching file blktap2/include/libvhd.h
patching file blktap2/include/vhd-util.h
patching file blktap2/vhd/lib/libvhd.c
patching file blktap2/vhd/lib/Makefile
patching file blktap2/vhd/lib/vhd-util-check.c
patching file blktap2/vhd/lib/vhd-util-convert.c
patching file blktap2/vhd/vhd-util.c



vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ ./configure --disable-monitors --disable-ocamltools --disable-rombios --disable-seabios
checking build system type... x86_64-unknown-linux-gnu
checking host system type... x86_64-unknown-linux-gnu
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether make sets $(MAKE)... yes
checking for a BSD-compatible install... /usr/bin/install -c
checking for bison... /usr/bin/bison
checking for flex... /usr/bin/flex
checking for perl... /usr/bin/perl
checking for bash... /bin/bash
checking for python... /usr/bin/python
checking for python version >= 2.3 ... yes
checking how to run the C preprocessor... gcc -E
checking for grep that handles long lines and -e... /bin/grep
checking for egrep... /bin/grep -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for python-config... /usr/bin/python-config
checking Python.h usability... yes
checking Python.h presence... yes
checking for Python.h... yes
checking for PyArg_ParseTuple in -lpython2.7... yes
checking for xgettext... /usr/bin/xgettext
checking for as86... no
configure: error: Unable to find as86, please install as86


vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ apt-cache search as86
bin86 - 16-bit x86 assembler and loader


vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ sudo apt-get install bin86
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  fontconfig fonts-font-awesome libdatrie1 libdbi1 libhiredis0.13 libjemalloc1 libjs-bootstrap libjs-d3 libjs-jquery-form libjs-jquery-metadata
  libjs-jquery-tablesorter libjs-rickshaw libmysqlclient20 libndpi5 libnorm1 libpango-1.0-0 libpangoft2-1.0-0 libpgm-5.2-0 libsodium23 libthai-data
  libthai0 libwireshark-data libxcb-render0 libxcb-shm0 libxrender1 libzmq5 mysql-common ntopng-data redis-server redis-tools
Use 'sudo apt autoremove' to remove them.
The following NEW packages will be installed:
  bin86
0 upgraded, 1 newly installed, 0 to remove and 143 not upgraded.
Need to get 81.2 kB of archives.
After this operation, 249 kB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu bionic/universe amd64 bin86 amd64 0.16.17-3.3 [81.2 kB]
Fetched 81.2 kB in 2s (34.7 kB/s)
Selecting previously unselected package bin86.
(Reading database ... 141292 files and directories currently installed.)
Preparing to unpack .../bin86_0.16.17-3.3_amd64.deb ...
Unpacking bin86 (0.16.17-3.3) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...
Setting up bin86 (0.16.17-3.3) ...


vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ ./configure --disable-monitors --disable-ocamltools --disable-rombios --disable-seabios
checking build system type... x86_64-unknown-linux-gnu
checking host system type... x86_64-unknown-linux-gnu
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether make sets $(MAKE)... yes
checking for a BSD-compatible install... /usr/bin/install -c
checking for bison... /usr/bin/bison
checking for flex... /usr/bin/flex
checking for perl... /usr/bin/perl
checking for bash... /bin/bash
checking for python... /usr/bin/python
checking for python version >= 2.3 ... yes
checking how to run the C preprocessor... gcc -E
checking for grep that handles long lines and -e... /bin/grep
checking for egrep... /bin/grep -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for python-config... /usr/bin/python-config
checking Python.h usability... yes
checking Python.h presence... yes
checking for Python.h... yes
checking for PyArg_ParseTuple in -lpython2.7... yes
checking for xgettext... /usr/bin/xgettext
checking for as86... /usr/bin/as86
checking for ld86... /usr/bin/ld86
checking for bcc... no
configure: error: Unable to find bcc, please install bcc




vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ sudo apt-get install bcc
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  fontconfig fonts-font-awesome libdatrie1 libdbi1 libhiredis0.13 libjemalloc1 libjs-bootstrap libjs-d3 libjs-jquery-form libjs-jquery-metadata
  libjs-jquery-tablesorter libjs-rickshaw libmysqlclient20 libndpi5 libnorm1 libpango-1.0-0 libpangoft2-1.0-0 libpgm-5.2-0 libsodium23 libthai-data
  libthai0 libwireshark-data libxcb-render0 libxcb-shm0 libxrender1 libzmq5 mysql-common ntopng-data redis-server redis-tools
Use 'sudo apt autoremove' to remove them.
The following additional packages will be installed:
  elks-libc
The following NEW packages will be installed:
  bcc elks-libc
0 upgraded, 2 newly installed, 0 to remove and 143 not upgraded.
Need to get 232 kB of archives.
After this operation, 954 kB of additional disk space will be used.
Do you want to continue? [Y/n] y
Get:1 http://archive.ubuntu.com/ubuntu bionic/universe amd64 bcc amd64 0.16.17-3.3 [109 kB]
Get:2 http://archive.ubuntu.com/ubuntu bionic/universe amd64 elks-libc all 0.16.17-3.3 [123 kB]
Fetched 232 kB in 5s (47.3 kB/s)    
Selecting previously unselected package bcc.
(Reading database ... 141311 files and directories currently installed.)
Preparing to unpack .../bcc_0.16.17-3.3_amd64.deb ...
Unpacking bcc (0.16.17-3.3) ...
Selecting previously unselected package elks-libc.
Preparing to unpack .../elks-libc_0.16.17-3.3_all.deb ...
Unpacking elks-libc (0.16.17-3.3) ...
Setting up elks-libc (0.16.17-3.3) ...
Setting up bcc (0.16.17-3.3) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...



vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ ./configure --disable-monitors --disable-ocamltools --disable-rombios --disable-seabios
checking build system type... x86_64-unknown-linux-gnu
checking host system type... x86_64-unknown-linux-gnu
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether make sets $(MAKE)... yes
checking for a BSD-compatible install... /usr/bin/install -c
checking for bison... /usr/bin/bison
checking for flex... /usr/bin/flex
checking for perl... /usr/bin/perl
checking for bash... /bin/bash
checking for python... /usr/bin/python
checking for python version >= 2.3 ... yes
checking how to run the C preprocessor... gcc -E
checking for grep that handles long lines and -e... /bin/grep
checking for egrep... /bin/grep -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for python-config... /usr/bin/python-config
checking Python.h usability... yes
checking Python.h presence... yes
checking for Python.h... yes
checking for PyArg_ParseTuple in -lpython2.7... yes
checking for xgettext... /usr/bin/xgettext
checking for as86... /usr/bin/as86
checking for ld86... /usr/bin/ld86
checking for bcc... /usr/bin/bcc
checking for iasl... no
configure: error: Unable to find iasl, please install iasl




vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ sudo apt-get install iasl
Reading package lists... Done
Building dependency tree       
Reading state information... Done
Note, selecting 'acpica-tools' instead of 'iasl'
The following packages were automatically installed and are no longer required:
  fontconfig fonts-font-awesome libdatrie1 libdbi1 libhiredis0.13 libjemalloc1 libjs-bootstrap libjs-d3 libjs-jquery-form libjs-jquery-metadata
  libjs-jquery-tablesorter libjs-rickshaw libmysqlclient20 libndpi5 libnorm1 libpango-1.0-0 libpangoft2-1.0-0 libpgm-5.2-0 libsodium23 libthai-data
  libthai0 libwireshark-data libxcb-render0 libxcb-shm0 libxrender1 libzmq5 mysql-common ntopng-data redis-server redis-tools
Use 'sudo apt autoremove' to remove them.
The following NEW packages will be installed:
  acpica-tools
0 upgraded, 1 newly installed, 0 to remove and 143 not upgraded.
Need to get 841 kB of archives.
After this operation, 2758 kB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu bionic/universe amd64 acpica-tools amd64 20180105-1 [841 kB]
Fetched 841 kB in 5s (156 kB/s)       
Selecting previously unselected package acpica-tools.
(Reading database ... 141469 files and directories currently installed.)
Preparing to unpack .../acpica-tools_20180105-1_amd64.deb ...
Unpacking acpica-tools (20180105-1) ...
Setting up acpica-tools (20180105-1) ...
update-alternatives: using /usr/bin/acpixtract-acpica to provide /usr/bin/acpixtract (acpixtract) in auto mode
update-alternatives: warning: skip creation of /usr/share/man/man1/acpixtract.1.gz because associated file /usr/share/man/man1/acpixtract-acpica.1 (of link group acpixtract) doesn't exist
update-alternatives: using /usr/bin/acpidump-acpica to provide /usr/bin/acpidump (acpidump) in auto mode
update-alternatives: warning: skip creation of /usr/share/man/man1/acpidump.1.gz because associated file /usr/share/man/man1/acpidump-acpica.1 (of link group acpidump) doesn't exist
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...



vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ ./configure --disable-monitors --disable-ocamltools --disable-rombios --disable-seabios
checking build system type... x86_64-unknown-linux-gnu
checking host system type... x86_64-unknown-linux-gnu
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether make sets $(MAKE)... yes
checking for a BSD-compatible install... /usr/bin/install -c
checking for bison... /usr/bin/bison
checking for flex... /usr/bin/flex
checking for perl... /usr/bin/perl
checking for bash... /bin/bash
checking for python... /usr/bin/python
checking for python version >= 2.3 ... yes
checking how to run the C preprocessor... gcc -E
checking for grep that handles long lines and -e... /bin/grep
checking for egrep... /bin/grep -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for python-config... /usr/bin/python-config
checking Python.h usability... yes
checking Python.h presence... yes
checking for Python.h... yes
checking for PyArg_ParseTuple in -lpython2.7... yes
checking for xgettext... /usr/bin/xgettext
checking for as86... /usr/bin/as86
checking for ld86... /usr/bin/ld86
checking for bcc... /usr/bin/bcc
checking for iasl... /usr/bin/iasl
checking uuid/uuid.h usability... no
checking uuid/uuid.h presence... no
checking for uuid/uuid.h... no
checking uuid.h usability... no
checking uuid.h presence... no
checking for uuid.h... no
configure: error: cannot find a valid uuid library



vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ sudo apt-get install uuid
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  fontconfig fonts-font-awesome libdatrie1 libdbi1 libhiredis0.13 libjemalloc1 libjs-bootstrap libjs-d3 libjs-jquery-form libjs-jquery-metadata
  libjs-jquery-tablesorter libjs-rickshaw libmysqlclient20 libndpi5 libnorm1 libpango-1.0-0 libpangoft2-1.0-0 libpgm-5.2-0 libsodium23 libthai-data
  libthai0 libwireshark-data libxcb-render0 libxcb-shm0 libxrender1 libzmq5 mysql-common ntopng-data redis-server redis-tools
Use 'sudo apt autoremove' to remove them.
The following additional packages will be installed:
  libossp-uuid16
The following NEW packages will be installed:
  libossp-uuid16 uuid
0 upgraded, 2 newly installed, 0 to remove and 143 not upgraded.
Need to get 40.0 kB of archives.
After this operation, 158 kB of additional disk space will be used.
Do you want to continue? [Y/n] y
Get:1 http://archive.ubuntu.com/ubuntu bionic/universe amd64 libossp-uuid16 amd64 1.6.2-1.5build4 [29.0 kB]
Get:2 http://archive.ubuntu.com/ubuntu bionic/universe amd64 uuid amd64 1.6.2-1.5build4 [10.9 kB]                                                     
Fetched 40.0 kB in 12s (3325 B/s)                                                                                                                     
Selecting previously unselected package libossp-uuid16:amd64.
(Reading database ... 141491 files and directories currently installed.)
Preparing to unpack .../libossp-uuid16_1.6.2-1.5build4_amd64.deb ...
Unpacking libossp-uuid16:amd64 (1.6.2-1.5build4) ...
Selecting previously unselected package uuid.
Preparing to unpack .../uuid_1.6.2-1.5build4_amd64.deb ...
Unpacking uuid (1.6.2-1.5build4) ...
Setting up libossp-uuid16:amd64 (1.6.2-1.5build4) ...
Setting up uuid (1.6.2-1.5build4) ...
Processing triggers for libc-bin (2.27-3ubuntu1) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...



vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ sudo apt-get install uuid-dev 
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  fontconfig fonts-font-awesome libdatrie1 libdbi1 libhiredis0.13 libjemalloc1 libjs-bootstrap libjs-d3 libjs-jquery-form libjs-jquery-metadata
  libjs-jquery-tablesorter libjs-rickshaw libmysqlclient20 libndpi5 libnorm1 libpango-1.0-0 libpangoft2-1.0-0 libpgm-5.2-0 libsodium23 libthai-data
  libthai0 libwireshark-data libxcb-render0 libxcb-shm0 libxrender1 libzmq5 mysql-common ntopng-data redis-server redis-tools
Use 'sudo apt autoremove' to remove them.
The following additional packages will be installed:
  libuuid1
The following NEW packages will be installed:
  uuid-dev
The following packages will be upgraded:
  libuuid1
1 upgraded, 1 newly installed, 0 to remove and 142 not upgraded.
Need to get 53.2 kB of archives.
After this operation, 166 kB of additional disk space will be used.
Do you want to continue? [Y/n] y
Get:1 http://archive.ubuntu.com/ubuntu bionic-updates/main amd64 libuuid1 amd64 2.31.1-0.4ubuntu3.4 [20.0 kB]
Get:2 http://archive.ubuntu.com/ubuntu bionic-updates/main amd64 uuid-dev amd64 2.31.1-0.4ubuntu3.4 [33.2 kB]
Fetched 53.2 kB in 3s (16.2 kB/s)   
(Reading database ... 141503 files and directories currently installed.)
Preparing to unpack .../libuuid1_2.31.1-0.4ubuntu3.4_amd64.deb ...
Unpacking libuuid1:amd64 (2.31.1-0.4ubuntu3.4) over (2.31.1-0.4ubuntu3.2) ...
Setting up libuuid1:amd64 (2.31.1-0.4ubuntu3.4) ...
Selecting previously unselected package uuid-dev:amd64.
(Reading database ... 141503 files and directories currently installed.)
Preparing to unpack .../uuid-dev_2.31.1-0.4ubuntu3.4_amd64.deb ...
Unpacking uuid-dev:amd64 (2.31.1-0.4ubuntu3.4) ...
Setting up uuid-dev:amd64 (2.31.1-0.4ubuntu3.4) ...
Processing triggers for libc-bin (2.27-3ubuntu1) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...




vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ ./configure --disable-monitors --disable-ocamltools --disable-rombios --disable-seabios
checking build system type... x86_64-unknown-linux-gnu
checking host system type... x86_64-unknown-linux-gnu
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether make sets $(MAKE)... yes
checking for a BSD-compatible install... /usr/bin/install -c
checking for bison... /usr/bin/bison
checking for flex... /usr/bin/flex
checking for perl... /usr/bin/perl
checking for bash... /bin/bash
checking for python... /usr/bin/python
checking for python version >= 2.3 ... yes
checking how to run the C preprocessor... gcc -E
checking for grep that handles long lines and -e... /bin/grep
checking for egrep... /bin/grep -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for python-config... /usr/bin/python-config
checking Python.h usability... yes
checking Python.h presence... yes
checking for Python.h... yes
checking for PyArg_ParseTuple in -lpython2.7... yes
checking for xgettext... /usr/bin/xgettext
checking for as86... /usr/bin/as86
checking for ld86... /usr/bin/ld86
checking for bcc... /usr/bin/bcc
checking for iasl... /usr/bin/iasl
checking uuid/uuid.h usability... yes
checking uuid/uuid.h presence... yes
checking for uuid/uuid.h... yes
checking for uuid_clear in -luuid... yes
checking uuid.h usability... no
checking uuid.h presence... no
checking for uuid.h... no
checking curses.h usability... no
checking curses.h presence... no
checking for curses.h... no
checking ncurses.h usability... no
checking ncurses.h presence... no
checking for ncurses.h... no
configure: error: Unable to find a suitable curses library



vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ sudo apt-get install libncurses5-dev 
Reading package lists... Done
Building dependency tree       
Reading state information... Done
The following packages were automatically installed and are no longer required:
  fontconfig fonts-font-awesome libdatrie1 libdbi1 libhiredis0.13 libjemalloc1 libjs-bootstrap libjs-d3 libjs-jquery-form libjs-jquery-metadata
  libjs-jquery-tablesorter libjs-rickshaw libmysqlclient20 libndpi5 libnorm1 libpango-1.0-0 libpangoft2-1.0-0 libpgm-5.2-0 libsodium23 libthai-data
  libthai0 libwireshark-data libxcb-render0 libxcb-shm0 libxrender1 libzmq5 mysql-common ntopng-data redis-server redis-tools
Use 'sudo apt autoremove' to remove them.
Suggested packages:
  ncurses-doc
The following NEW packages will be installed:
  libncurses5-dev
0 upgraded, 1 newly installed, 0 to remove and 142 not upgraded.
Need to get 174 kB of archives.
After this operation, 1017 kB of additional disk space will be used.
Get:1 http://archive.ubuntu.com/ubuntu bionic-updates/main amd64 libncurses5-dev amd64 6.1-1ubuntu1.18.04 [174 kB]
Fetched 174 kB in 18s (9883 B/s)                                                                                                                      
Selecting previously unselected package libncurses5-dev:amd64.
(Reading database ... 141523 files and directories currently installed.)
Preparing to unpack .../libncurses5-dev_6.1-1ubuntu1.18.04_amd64.deb ...
Unpacking libncurses5-dev:amd64 (6.1-1ubuntu1.18.04) ...
Setting up libncurses5-dev:amd64 (6.1-1ubuntu1.18.04) ...
Processing triggers for man-db (2.8.3-2ubuntu0.1) ...



vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ ./configure --disable-monitors --disable-ocamltools --disable-rombios --disable-seabios
checking build system type... x86_64-unknown-linux-gnu
checking host system type... x86_64-unknown-linux-gnu
checking for gcc... gcc
checking whether the C compiler works... yes
checking for C compiler default output file name... a.out
checking for suffix of executables... 
checking whether we are cross compiling... no
checking for suffix of object files... o
checking whether we are using the GNU C compiler... yes
checking whether gcc accepts -g... yes
checking for gcc option to accept ISO C89... none needed
checking whether make sets $(MAKE)... yes
checking for a BSD-compatible install... /usr/bin/install -c
checking for bison... /usr/bin/bison
checking for flex... /usr/bin/flex
checking for perl... /usr/bin/perl
checking for bash... /bin/bash
checking for python... /usr/bin/python
checking for python version >= 2.3 ... yes
checking how to run the C preprocessor... gcc -E
checking for grep that handles long lines and -e... /bin/grep
checking for egrep... /bin/grep -E
checking for ANSI C header files... yes
checking for sys/types.h... yes
checking for sys/stat.h... yes
checking for stdlib.h... yes
checking for string.h... yes
checking for memory.h... yes
checking for strings.h... yes
checking for inttypes.h... yes
checking for stdint.h... yes
checking for unistd.h... yes
checking for python-config... /usr/bin/python-config
checking Python.h usability... yes
checking Python.h presence... yes
checking for Python.h... yes
checking for PyArg_ParseTuple in -lpython2.7... yes
checking for xgettext... /usr/bin/xgettext
checking for as86... /usr/bin/as86
checking for ld86... /usr/bin/ld86
checking for bcc... /usr/bin/bcc
checking for iasl... /usr/bin/iasl
checking uuid/uuid.h usability... yes
checking uuid/uuid.h presence... yes
checking for uuid/uuid.h... yes
checking for uuid_clear in -luuid... yes
checking uuid.h usability... no
checking uuid.h presence... no
checking for uuid.h... no
checking curses.h usability... yes
checking curses.h presence... yes
checking for curses.h... yes
checking for clear in -lcurses... yes
checking ncurses.h usability... yes
checking ncurses.h presence... yes
checking for ncurses.h... yes
checking for clear in -lncurses... yes
checking for pkg-config... /usr/bin/pkg-config
checking pkg-config is at least version 0.9.0... yes
checking for glib... yes
checking bzlib.h usability... no
checking bzlib.h presence... no
checking for bzlib.h... no
checking lzma.h usability... yes
checking lzma.h presence... yes
checking for lzma.h... yes
checking for lzma_stream_decoder in -llzma... yes
checking lzo/lzo1x.h usability... no
checking lzo/lzo1x.h presence... no
checking for lzo/lzo1x.h... no
checking for io_setup in -laio... no
checking for MD5 in -lcrypto... yes
checking for ext2fs_open2 in -lext2fs... no
checking for gcry_md_hash_buffer in -lgcrypt... no
checking for pthread flag... -pthread
checking libutil.h usability... no
checking libutil.h presence... no
checking for libutil.h... no
checking for openpty et al... -lutil
checking for yajl_alloc in -lyajl... yes
checking for deflateCopy in -lz... yes
checking for libiconv_open in -liconv... no
checking yajl/yajl_version.h usability... yes
checking yajl/yajl_version.h presence... yes
checking for yajl/yajl_version.h... yes
configure: creating ./config.status
config.status: creating ../config/Tools.mk
config.status: creating config.h




vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ make
make[1]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools'
make -C include all
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include'
make -C xen-foreign
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign'
python mkheader.py x86_32 x86_32.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign/../../../xen/include/public/arch-x86/xen-x86_32.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign/../../../xen/include/public/arch-x86/xen.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign/../../../xen/include/public/xen.h
python mkheader.py x86_64 x86_64.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign/../../../xen/include/public/arch-x86/xen-x86_64.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign/../../../xen/include/public/arch-x86/xen.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign/../../../xen/include/public/xen.h
python mkheader.py ia64 ia64.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign/../../../xen/include/public/arch-ia64.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign/../../../xen/include/public/xen.h
python mkchecker.py checker.c x86_32 x86_64 ia64
gcc -Wall -Werror -Wstrict-prototypes -O2 -fomit-frame-pointer -fno-strict-aliasing -Wdeclaration-after-statement -o checker checker.c
./checker > tmp.size
diff -u reference.size tmp.size
rm tmp.size
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/xen-foreign'
mkdir -p xen/libelf
ln -sf /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/COPYING xen
ln -sf /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/memory.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/sysctl.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/kexec.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/arch-ia64.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/nmi.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/xencomm.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/vcpu.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/arch-arm.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/elfnote.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/xenoprof.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/event_channel.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/version.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/mem_event.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/dom0_ops.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/arch-x86_32.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/xen.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/features.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/sched.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/xen-compat.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/callback.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/grant_table.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/physdev.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/arch-x86_64.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/platform.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/tmem.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/trace.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/domctl.h xen
ln -sf /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/arch-ia64 /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/arch-x86 /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/hvm /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/io /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/public/xsm xen
ln -sf ../xen-sys/Linux xen/sys
ln -sf /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/xen/libelf.h /Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include/../../xen/include/xen/elfstructs.h xen/libelf/
ln -s ../xen-foreign xen/foreign
touch xen/.dir
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/include'
make[1]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools'
make[1]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools'
make -C libxc all
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc'
make libs
make[3]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc'
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_core.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_core.o xc_core.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_core_x86.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_core_x86.o xc_core_x86.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_cpupool.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_cpupool.o xc_cpupool.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_domain.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_domain.o xc_domain.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_evtchn.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_evtchn.o xc_evtchn.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_gnttab.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_gnttab.o xc_gnttab.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_misc.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_misc.o xc_misc.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_flask.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_flask.o xc_flask.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_physdev.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_physdev.o xc_physdev.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_private.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_private.o xc_private.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_sedf.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_sedf.o xc_sedf.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_csched.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_csched.o xc_csched.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_csched2.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_csched2.o xc_csched2.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_arinc653.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_arinc653.o xc_arinc653.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_tbuf.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_tbuf.o xc_tbuf.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_pm.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_pm.o xc_pm.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_cpu_hotplug.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_cpu_hotplug.o xc_cpu_hotplug.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_resume.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_resume.o xc_resume.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_tmem.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_tmem.o xc_tmem.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_mem_event.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_mem_event.o xc_mem_event.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_mem_paging.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_mem_paging.o xc_mem_paging.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_mem_access.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_mem_access.o xc_mem_access.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_memshr.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_memshr.o xc_memshr.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_hcall_buf.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_hcall_buf.o xc_hcall_buf.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_foreign_memory.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_foreign_memory.o xc_foreign_memory.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xtl_core.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xtl_core.o xtl_core.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xtl_logger_stdio.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xtl_logger_stdio.o xtl_logger_stdio.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_pagetab.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_pagetab.o xc_pagetab.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_linux.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_linux.o xc_linux.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_linux_osdep.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_linux_osdep.o xc_linux_osdep.c 
ar rc libxenctrl.a xc_core.o xc_core_x86.o xc_cpupool.o xc_domain.o xc_evtchn.o xc_gnttab.o xc_misc.o xc_flask.o xc_physdev.o xc_private.o xc_sedf.o xc_csched.o xc_csched2.o xc_arinc653.o xc_tbuf.o xc_pm.o xc_cpu_hotplug.o xc_resume.o xc_tmem.o xc_mem_event.o xc_mem_paging.o xc_mem_access.o xc_memshr.o xc_hcall_buf.o xc_foreign_memory.o xtl_core.o xtl_logger_stdio.o xc_pagetab.o xc_linux.o xc_linux_osdep.o
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_core.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_core.opic xc_core.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_core_x86.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_core_x86.opic xc_core_x86.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_cpupool.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_cpupool.opic xc_cpupool.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_domain.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_domain.opic xc_domain.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_evtchn.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_evtchn.opic xc_evtchn.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_gnttab.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_gnttab.opic xc_gnttab.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_misc.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_misc.opic xc_misc.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_flask.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_flask.opic xc_flask.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_physdev.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_physdev.opic xc_physdev.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_private.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_private.opic xc_private.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_sedf.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_sedf.opic xc_sedf.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_csched.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_csched.opic xc_csched.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_csched2.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_csched2.opic xc_csched2.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_arinc653.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_arinc653.opic xc_arinc653.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_tbuf.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_tbuf.opic xc_tbuf.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_pm.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_pm.opic xc_pm.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_cpu_hotplug.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_cpu_hotplug.opic xc_cpu_hotplug.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_resume.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_resume.opic xc_resume.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_tmem.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_tmem.opic xc_tmem.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_mem_event.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_mem_event.opic xc_mem_event.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_mem_paging.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_mem_paging.opic xc_mem_paging.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_mem_access.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_mem_access.opic xc_mem_access.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_memshr.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_memshr.opic xc_memshr.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_hcall_buf.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_hcall_buf.opic xc_hcall_buf.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_foreign_memory.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_foreign_memory.opic xc_foreign_memory.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xtl_core.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xtl_core.opic xtl_core.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xtl_logger_stdio.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xtl_logger_stdio.opic xtl_logger_stdio.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_pagetab.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_pagetab.opic xc_pagetab.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_linux.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_linux.opic xc_linux.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_linux_osdep.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -fPIC -c -o xc_linux_osdep.opic xc_linux_osdep.c 
gcc     -pthread -Wl,-soname -Wl,libxenctrl.so.4.2 -shared -o libxenctrl.so.4.2.0 xc_core.opic xc_core_x86.opic xc_cpupool.opic xc_domain.opic xc_evtchn.opic xc_gnttab.opic xc_misc.opic xc_flask.opic xc_physdev.opic xc_private.opic xc_sedf.opic xc_csched.opic xc_csched2.opic xc_arinc653.opic xc_tbuf.opic xc_pm.opic xc_cpu_hotplug.opic xc_resume.opic xc_tmem.opic xc_mem_event.opic xc_mem_paging.opic xc_mem_access.opic xc_memshr.opic xc_hcall_buf.opic xc_foreign_memory.opic xtl_core.opic xtl_logger_stdio.opic xc_pagetab.opic xc_linux.opic xc_linux_osdep.opic -ldl  
ln -sf libxenctrl.so.4.2.0 libxenctrl.so.4.2
ln -sf libxenctrl.so.4.2 libxenctrl.so
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xg_private.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xg_private.o xg_private.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_suspend.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_suspend.o xc_suspend.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_domain_restore.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_domain_restore.o xc_domain_restore.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_domain_save.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_domain_save.o xc_domain_save.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_offline_page.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_offline_page.o xc_offline_page.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_compression.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_compression.o xc_compression.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .libelf-tools.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o libelf-tools.o ../../xen/common/libelf/libelf-tools.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .libelf-loader.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o libelf-loader.o ../../xen/common/libelf/libelf-loader.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .libelf-dominfo.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o libelf-dominfo.o ../../xen/common/libelf/libelf-dominfo.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .libelf-relocate.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o libelf-relocate.o ../../xen/common/libelf/libelf-relocate.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_dom_core.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_dom_core.o xc_dom_core.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .xc_dom_boot.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls -D_GNU_SOURCE  -I../../xen/common/libelf -Werror -Wmissing-prototypes -I. -I/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/include -pthread  -c -o xc_dom_boot.o xc_dom_boot.c 
xc_dom_boot.c: In function ‘xc_dom_boot_image’:
xc_dom_boot.c:269:27: error: argument to ‘sizeof’ in ‘memset’ call is the same expression as the destination; did you mean to dereference it? [-Werror=sizeof-pointer-memaccess]
     memset(ctxt, 0, sizeof(ctxt));
                           ^
cc1: all warnings being treated as errors
/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc/../../tools/Rules.mk:93: recipe for target 'xc_dom_boot.o' failed
make[3]: *** [xc_dom_boot.o] Error 1
make[3]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc'
Makefile:116: recipe for target 'build' failed
make[2]: *** [build] Error 2
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/libxc'
/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/../tools/Rules.mk:109: recipe for target 'subdir-all-libxc' failed
make[1]: *** [subdir-all-libxc] Error 2
make[1]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools'
/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/../tools/Rules.mk:104: recipe for target 'subdirs-all' failed
make: *** [subdirs-all] Error 2




vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools$ cd blktap2/vhd/
vagrant@ubuntu-bionic:/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/blktap2/vhd$ makemake[1]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/blktap2/vhd'
make -C lib all
make[2]: Entering directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/blktap2/vhd/lib'
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .libvhd.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o libvhd.o libvhd.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .libvhd-journal.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o libvhd-journal.o libvhd-journal.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-coalesce.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-coalesce.o vhd-util-coalesce.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-create.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-create.o vhd-util-create.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-fill.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-fill.o vhd-util-fill.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-modify.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-modify.o vhd-util-modify.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-query.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-query.o vhd-util-query.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-read.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-read.o vhd-util-read.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-repair.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-repair.o vhd-util-repair.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-resize.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-resize.o vhd-util-resize.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-revert.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-revert.o vhd-util-revert.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-set-field.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-set-field.o vhd-util-set-field.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-snapshot.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-snapshot.o vhd-util-snapshot.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-scan.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-scan.o vhd-util-scan.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-check.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-check.o vhd-util-check.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-convert.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-convert.o vhd-util-convert.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-uuid.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o vhd-util-uuid.o vhd-util-uuid.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .relative-path.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o relative-path.o relative-path.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .atomicio.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o atomicio.o atomicio.c 
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .lvm-util.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -c -o ../../lvm/lvm-util.o ../../lvm/lvm-util.c 
ar rc libvhd.a libvhd.o libvhd-journal.o vhd-util-coalesce.o vhd-util-create.o vhd-util-fill.o vhd-util-modify.o vhd-util-query.o vhd-util-read.o vhd-util-repair.o vhd-util-resize.o vhd-util-revert.o vhd-util-set-field.o vhd-util-snapshot.o vhd-util-scan.o vhd-util-check.o vhd-util-convert.o vhd-util-uuid.o relative-path.o atomicio.o ../../lvm/lvm-util.o
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .libvhd.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o libvhd.opic libvhd.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .libvhd-journal.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o libvhd-journal.opic libvhd-journal.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-coalesce.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-coalesce.opic vhd-util-coalesce.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-create.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-create.opic vhd-util-create.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-fill.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-fill.opic vhd-util-fill.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-modify.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-modify.opic vhd-util-modify.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-query.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-query.opic vhd-util-query.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-read.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-read.opic vhd-util-read.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-repair.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-repair.opic vhd-util-repair.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-resize.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-resize.opic vhd-util-resize.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-revert.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-revert.opic vhd-util-revert.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-set-field.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-set-field.opic vhd-util-set-field.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-snapshot.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-snapshot.opic vhd-util-snapshot.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-scan.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-scan.opic vhd-util-scan.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-check.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-check.opic vhd-util-check.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-convert.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-convert.opic vhd-util-convert.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util-uuid.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o vhd-util-uuid.opic vhd-util-uuid.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .relative-path.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o relative-path.opic relative-path.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .atomicio.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o atomicio.opic atomicio.c 
gcc  -DPIC -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .lvm-util.opic.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../../include -D_GNU_SOURCE -fPIC -g  -fPIC -c -o ../../lvm/lvm-util.opic ../../lvm/lvm-util.c 
gcc -Wl,-soname,libvhd.so.1.0 -shared \
	    -o libvhd.so.1.0.0 libvhd.opic libvhd-journal.opic vhd-util-coalesce.opic vhd-util-create.opic vhd-util-fill.opic vhd-util-modify.opic vhd-util-query.opic vhd-util-read.opic vhd-util-repair.opic vhd-util-resize.opic vhd-util-revert.opic vhd-util-set-field.opic vhd-util-snapshot.opic vhd-util-scan.opic vhd-util-check.opic vhd-util-convert.opic vhd-util-uuid.opic relative-path.opic atomicio.opic ../../lvm/lvm-util.opic -luuid
ln -sf libvhd.so.1.0.0 libvhd.so.1.0
ln -sf libvhd.so.1.0 libvhd.so
make[2]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/blktap2/vhd/lib'
make[1]: Leaving directory '/Users/fanhongling/Downloads/workspace/src/github.com/tangfeixiong/go-to-openstack-bootcamp/xen/xen-4.2.0/tools/blktap2/vhd'
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-util.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../include -D_GNU_SOURCE -fPIC  -c -o vhd-util.o vhd-util.c 
make: Warning: File 'lib/libvhd.so' has modification time 1.7 s in the future
gcc     -o vhd-util vhd-util.o -Llib -lvhd
gcc  -O1 -fno-omit-frame-pointer -m64 -g -fno-strict-aliasing -std=gnu99 -Wall -Wstrict-prototypes -Wdeclaration-after-statement -Wno-unused-but-set-variable   -D__XEN_TOOLS__ -MMD -MF .vhd-update.o.d  -D_LARGEFILE_SOURCE -D_LARGEFILE64_SOURCE -fno-optimize-sibling-calls  -Werror -Wno-unused -I../include -D_GNU_SOURCE -fPIC  -c -o vhd-update.o vhd-update.c 
gcc     -o vhd-update vhd-update.o -Llib -lvhd
make: warning:  Clock skew detected.  Your build may be incomplete.
