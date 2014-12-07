


Administrator@lenovo-9d779749 ~
$ python --version
Python 2.7.8

Administrator@lenovo-9d779749 ~
$ pip --version
pip 1.2.1 from /usr/lib/python2.7/site-packages/pip-1.2.1-py2.7.egg (python 2.7)

Administrator@lenovo-9d779749 ~
$ pip install --upgrade python-keystoneclient
Downloading/unpacking python-keystoneclient
  Running setup.py egg_info for package python-keystoneclient
    [pbr] Reusing existing SOURCES.txt
Cannot fetch index base URL http://pypi.python.org/simple/
Requirement already up-to-date: pbr>=0.6,!=0.7,<1.0 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Downloading/unpacking argparse (from python-keystoneclient)
  Downloading argparse-1.2.2.tar.gz (68kB): 68kB downloaded
  Running setup.py egg_info for package argparse

    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files matching '*.pyo' found anywhere in distribution
    warning: no previously-included files matching '*.orig' found anywhere in distribution
    warning: no previously-included files matching '*.rej' found anywhere in distribution
    no previously-included directories found matching 'doc/_build'
    no previously-included directories found matching 'env24'
    no previously-included directories found matching 'env25'
    no previously-included directories found matching 'env26'
    no previously-included directories found matching 'env27'
Downloading/unpacking Babel>=1.3 (from python-keystoneclient)
  Downloading Babel-1.3.tar.gz (3.4MB): 3.4MB downloaded
  Running setup.py egg_info for package Babel

    warning: no previously-included files matching '*' found under directory 'docs/_build'
    warning: no previously-included files matching '*.pyc' found under directory 'tests'
    warning: no previously-included files matching '*.pyo' found under directory 'tests'
Downloading/unpacking iso8601>=0.1.9 (from python-keystoneclient)
  Downloading iso8601-0.1.10.tar.gz
  Running setup.py egg_info for package iso8601

Downloading/unpacking netaddr>=0.7.12 (from python-keystoneclient)
  Cannot fetch index base URL http://pypi.python.org/simple/
  Downloading netaddr-0.7.12.tar.gz (1.5MB): 1.5MB downloaded
  Running setup.py egg_info for package netaddr

    warning: no previously-included files matching '*.svn*' found anywhere in distribution
    warning: no previously-included files matching '*.git*' found anywhere in distribution
Downloading/unpacking oslo.config>=1.4.0 (from python-keystoneclient)
  Downloading oslo.config-1.4.0.tar.gz (66kB): 66kB downloaded
  Running setup.py egg_info for package oslo.config
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found

    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
Downloading/unpacking oslo.serialization>=1.0.0 (from python-keystoneclient)
  Downloading oslo.serialization-1.0.0.tar.gz
  Running setup.py egg_info for package oslo.serialization
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found

    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
Downloading/unpacking oslo.utils>=1.0.0 (from python-keystoneclient)
  Downloading oslo.utils-1.0.0.tar.gz
  Running setup.py egg_info for package oslo.utils
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found

    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
Downloading/unpacking PrettyTable>=0.7,<0.8 (from python-keystoneclient)
  Downloading prettytable-0.7.2.tar.bz2
  Running setup.py egg_info for package PrettyTable

Downloading/unpacking requests>=2.2.0,!=2.4.0 (from python-keystoneclient)
  Downloading requests-2.4.3.tar.gz (438kB): 438kB downloaded
  Running setup.py egg_info for package requests

Downloading/unpacking six>=1.7.0 (from python-keystoneclient)
  Downloading six-1.8.0.tar.gz
Exception:
Traceback (most recent call last):
  File "/usr/lib/python2.7/site-packages/pip-1.2.1-py2.7.egg/pip/basecommand.py", line 107, in main
    status = self.run(options, args)
  File "/usr/lib/python2.7/site-packages/pip-1.2.1-py2.7.egg/pip/commands/install.py", line 256, in run
    requirement_set.prepare_files(finder, force_root_egg_info=self.bundle, bundle=self.bundle)
  File "/usr/lib/python2.7/site-packages/pip-1.2.1-py2.7.egg/pip/req.py", line 1018, in prepare_files
    self.unpack_url(url, location, self.is_download)
  File "/usr/lib/python2.7/site-packages/pip-1.2.1-py2.7.egg/pip/req.py", line 1142, in unpack_url
    retval = unpack_http_url(link, location, self.download_cache, self.download_dir)
  File "/usr/lib/python2.7/site-packages/pip-1.2.1-py2.7.egg/pip/download.py", line 463, in unpack_http_url
    download_hash = _download_url(resp, link, temp_location)
  File "/usr/lib/python2.7/site-packages/pip-1.2.1-py2.7.egg/pip/download.py", line 380, in _download_url
    chunk = resp.read(4096)
  File "/usr/lib/python2.7/socket.py", line 380, in read
    data = self._sock.recv(left)
  File "/usr/lib/python2.7/httplib.py", line 567, in read
    s = self.fp.read(amt)
  File "/usr/lib/python2.7/socket.py", line 380, in read
    data = self._sock.recv(left)
  File "/usr/lib/python2.7/ssl.py", line 246, in recv
    return self.read(buflen)
  File "/usr/lib/python2.7/ssl.py", line 165, in read
    return self._sslobj.read(len)
SSLError: The read operation timed out

Storing complete log in /home/Administrator/.pip/pip.log

### python nova client

Administrator@lenovo-9d779749 ~
$ pip install --upgrade python-novaclient
Downloading/unpacking python-novaclient
  Downloading python-novaclient-2.20.0.tar.gz (267kB): 267kB downloaded
  Running setup.py egg_info for package python-novaclient
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found

    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
Requirement already up-to-date: pbr>=0.6,!=0.7,<1.0 in /usr/lib/python2.7/site-packages (from python-novaclient)
Downloading/unpacking argparse (from python-novaclient)
  Running setup.py egg_info for package argparse

    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files matching '*.pyo' found anywhere in distribution
    warning: no previously-included files matching '*.orig' found anywhere in distribution
    warning: no previously-included files matching '*.rej' found anywhere in distribution
    no previously-included directories found matching 'doc/_build'
    no previously-included directories found matching 'env24'
    no previously-included directories found matching 'env25'
    no previously-included directories found matching 'env26'
    no previously-included directories found matching 'env27'
Downloading/unpacking iso8601>=0.1.9 (from python-novaclient)
  Running setup.py egg_info for package iso8601

Downloading/unpacking oslo.utils>=1.0.0 (from python-novaclient)
  Running setup.py egg_info for package oslo.utils
    [pbr] Reusing existing SOURCES.txt
Downloading/unpacking PrettyTable>=0.7,<0.8 (from python-novaclient)
  Running setup.py egg_info for package PrettyTable

Downloading/unpacking requests>=1.2.1,!=2.4.0 (from python-novaclient)
  Running setup.py egg_info for package requests

Downloading/unpacking simplejson>=2.2.0 (from python-novaclient)
  Downloading simplejson-3.6.5.tar.gz (73kB): 73kB downloaded
  Running setup.py egg_info for package simplejson

Downloading/unpacking six>=1.7.0 (from python-novaclient)
  Downloading six-1.8.0.tar.gz
  Running setup.py egg_info for package six

    no previously-included directories found matching 'documentation/_build'
Downloading/unpacking Babel>=1.3 (from python-novaclient)
  Running setup.py egg_info for package Babel

    warning: no previously-included files matching '*' found under directory 'docs/_build'
    warning: no previously-included files matching '*.pyc' found under directory 'tests'
    warning: no previously-included files matching '*.pyo' found under directory 'tests'
Downloading/unpacking python-keystoneclient>=0.10.0 (from python-novaclient)
  Running setup.py egg_info for package python-keystoneclient
    [pbr] Reusing existing SOURCES.txt
Downloading/unpacking pip from https://pypi.python.org/packages/source/p/pip/pip-1.5.6.tar.gz#md5=01026f87978932060cc86c1dc527903e (from pbr>=0.6,!=0.7,<1.0->python-novaclient)
  Downloading pip-1.5.6.tar.gz (938kB): 938kB downloaded
  Running setup.py egg_info for package pip

    warning: no files found matching 'pip/cacert.pem'
    warning: no files found matching '*.html' under directory 'docs'
    warning: no previously-included files matching '*.rst' found under directory 'docs/_build'
    no previously-included directories found matching 'docs/_build/_sources'
Downloading/unpacking oslo.i18n>=0.2.0 (from oslo.utils>=1.0.0->python-novaclient)
  Downloading oslo.i18n-1.0.0.tar.gz
  Running setup.py egg_info for package oslo.i18n
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found

    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
Downloading/unpacking pytz>=0a (from Babel>=1.3->python-novaclient)
  Downloading pytz-2014.9.tar.bz2 (166kB): 166kB downloaded
  Running setup.py egg_info for package pytz

Downloading/unpacking netaddr>=0.7.12 (from python-keystoneclient>=0.10.0->python-novaclient)
  Running setup.py egg_info for package netaddr

    warning: no previously-included files matching '*.svn*' found anywhere in distribution
    warning: no previously-included files matching '*.git*' found anywhere in distribution
Downloading/unpacking oslo.config>=1.4.0 (from python-keystoneclient>=0.10.0->python-novaclient)
  Running setup.py egg_info for package oslo.config
    [pbr] Reusing existing SOURCES.txt
Downloading/unpacking oslo.serialization>=1.0.0 (from python-keystoneclient>=0.10.0->python-novaclient)
  Running setup.py egg_info for package oslo.serialization
    [pbr] Reusing existing SOURCES.txt
Downloading/unpacking stevedore>=1.0.0 (from python-keystoneclient>=0.10.0->python-novaclient)
  Downloading stevedore-1.1.0.tar.gz (352kB): 352kB downloaded
  Running setup.py egg_info for package stevedore
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found

    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no files found matching '*.html' under directory 'docs'
    warning: no files found matching '*.css' under directory 'docs'
    warning: no files found matching '*.js' under directory 'docs'
    warning: no files found matching '*.png' under directory 'docs'
    warning: no files found matching '*.py' under directory 'tests'
Installing collected packages: python-novaclient, argparse, iso8601, oslo.utils, PrettyTable, requests, simplejson, six, Babel, python-keystoneclient, pip, oslo.i18n, pytz, netaddr, oslo.config, oslo.serialization, stevedore
  Running setup.py install for python-novaclient
    [pbr] Reusing existing SOURCES.txt
    Installing nova script to /usr/bin
  Running setup.py install for argparse

    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files matching '*.pyo' found anywhere in distribution
    warning: no previously-included files matching '*.orig' found anywhere in distribution
    warning: no previously-included files matching '*.rej' found anywhere in distribution
    no previously-included directories found matching 'doc/_build'
    no previously-included directories found matching 'env24'
    no previously-included directories found matching 'env25'
    no previously-included directories found matching 'env26'
    no previously-included directories found matching 'env27'
  Running setup.py install for iso8601

  Running setup.py install for oslo.utils
    [pbr] Reusing existing SOURCES.txt
    Skipping installation of /usr/lib/python2.7/site-packages/oslo/__init__.py (namespace package)
    Installing /usr/lib/python2.7/site-packages/oslo.utils-1.0.0-py2.7-nspkg.pth
  Running setup.py install for PrettyTable

  Running setup.py install for requests

  Running setup.py install for simplejson
    building 'simplejson._speedups' extension
    gcc -fno-strict-aliasing -ggdb -O2 -pipe -Wimplicit-function-declaration -fdebug-prefix-map=/usr/src/ports/python/python-2.7.8-1.i686/build=/usr/src/debug/python-2.7.8-1 -fdebug-prefix-map=/usr/src/ports/python/python-2.7.8-1.i686/src/Python-2.7.8=/usr/src/debug/python-2.7.8-1 -DNDEBUG -g -fwrapv -O3 -Wall -Wstrict-prototypes -I/usr/include/python2.7 -c simplejson/_speedups.c -o build/temp.cygwin-1.7.32-i686-2.7/simplejson/_speedups.o
    gcc -shared -Wl,--enable-auto-image-base -L. build/temp.cygwin-1.7.32-i686-2.7/simplejson/_speedups.o -L/usr/lib/python2.7/config -L/usr/lib -lpython2.7 -o build/lib.cygwin-1.7.32-i686-2.7/simplejson/_speedups.dll

  Running setup.py install for six

    no previously-included directories found matching 'documentation/_build'
  Running setup.py install for Babel

    warning: no previously-included files matching '*' found under directory 'docs/_build'
    warning: no previously-included files matching '*.pyc' found under directory 'tests'
    warning: no previously-included files matching '*.pyo' found under directory 'tests'
    Installing pybabel script to /usr/bin
  Running setup.py install for python-keystoneclient
    [pbr] Reusing existing SOURCES.txt
    Installing keystone script to /usr/bin
  Found existing installation: pip 1.2.1
    Uninstalling pip:
      Successfully uninstalled pip
  Running setup.py install for pip

    warning: no files found matching 'pip/cacert.pem'
    warning: no files found matching '*.html' under directory 'docs'
    warning: no previously-included files matching '*.rst' found under directory 'docs/_build'
    no previously-included directories found matching 'docs/_build/_sources'
    Installing pip script to /usr/bin
    Installing pip2.7 script to /usr/bin
    Installing pip2 script to /usr/bin
  Running setup.py install for oslo.i18n
    [pbr] Reusing existing SOURCES.txt
    Skipping installation of /usr/lib/python2.7/site-packages/oslo/__init__.py (namespace package)
    Installing /usr/lib/python2.7/site-packages/oslo.i18n-1.0.0-py2.7-nspkg.pth
  Running setup.py install for pytz

  Running setup.py install for netaddr
    changing mode of build/scripts-2.7/netaddr from 644 to 755

    warning: no previously-included files matching '*.svn*' found anywhere in distribution
    warning: no previously-included files matching '*.git*' found anywhere in distribution
    changing mode of /usr/bin/netaddr to 755
  Running setup.py install for oslo.config
    [pbr] Reusing existing SOURCES.txt
    Skipping installation of /usr/lib/python2.7/site-packages/oslo/__init__.py (namespace package)
    Installing /usr/lib/python2.7/site-packages/oslo.config-1.4.0-py2.7-nspkg.pth
    Installing oslo-config-generator script to /usr/bin
  Running setup.py install for oslo.serialization
    [pbr] Reusing existing SOURCES.txt
    Skipping installation of /usr/lib/python2.7/site-packages/oslo/__init__.py (namespace package)
    Installing /usr/lib/python2.7/site-packages/oslo.serialization-1.0.0-py2.7-nspkg.pth
  Running setup.py install for stevedore
    [pbr] Reusing existing SOURCES.txt
Successfully installed python-novaclient argparse iso8601 oslo.utils PrettyTable requests simplejson six Babel python-keystoneclient pip oslo.i18n pytz netaddr oslo.config oslo.serialization stevedore
Cleaning up...

Administrator@lenovo-9d779749 ~/sgepri-projects/os
$ scp root@192.168.70.162:/root/keystonerc_admin .
root@192.168.70.162's password:
keystonerc_admin                              100%  180     0.2KB/s   00:00

Administrator@lenovo-9d779749 ~/sgepri-projects/os
$ scp root@192.168.70.162:/root/keystonerc_demo .
root@192.168.70.162's password:
keystonerc_demo                               100%  181     0.2KB/s   00:00

Administrator@lenovo-9d779749 ~/sgepri-projects/os
$ scp root@192.168.70.162:/root/liuxm/keystonerc_liuxm .
root@192.168.70.162's password:
keystonerc_liuxm                              100%  173     0.2KB/s   00:00

Administrator@lenovo-9d779749 ~/sgepri-projects/os
$ ls
keystonerc_admin  keystonerc_demo  keystonerc_liuxm

Administrator@lenovo-9d779749 ~/sgepri-projects/os
$ . keystonerc_admin
[Administrator@lenovo-9d779749 os(keystone_admin)]$ nova list
+--------------------------------------+----------------------+--------+------------+-------------+-----------------------------------------+
| ID                                   | Name                 | Status | Task State | Power State | Networks                                |
+--------------------------------------+----------------------+--------+------------+-------------+-----------------------------------------+
| c8b4598a-79cf-40d8-a7b4-971a7e35581c | test                 | ACTIVE | -          | Running     | privateAdmin=10.1.10.16, 192.168.74.131 |
| ad28b0db-5ec1-4437-8726-40e10ead7c8e | wujiasheng-ubuntu-01 | ACTIVE | -          | Running     | privateAdmin=10.1.10.15                 |
+--------------------------------------+----------------------+--------+------------+-------------+-----------------------------------------+
[Administrator@lenovo-9d779749 os(keystone_admin)]$
[Administrator@lenovo-9d779749 os(keystone_admin)]$ keystone user-list
+----------------------------------+------------+---------+----------------------+
|                id                |    name    | enabled |        email         |
+----------------------------------+------------+---------+----------------------+
| 2901a68524fc4dea97018804ff1f98d4 |   admin    |   True  |    root@localhost    |
| f856fa7094cb446281c2218dc6cb55e5 | ceilometer |   True  | ceilometer@localhost |
| a80f918364df49fbb59e934c4a2bd38c |   cinder   |   True  |   cinder@localhost   |
| 05b7f9b0123d4837bcec4bf2bf6dfcd4 |    demo    |   True  |                      |
| f5a76bf523524cb7812e6ffc51e60bcb |   glance   |   True  |   glance@localhost   |
| 6237dd313aea4473b516aa409d5862f2 |   liudan   |   True  |                      |
| 0189c90d965c42eb85d730d81f14bc41 |   liuxm    |   True  |                      |
| 9f15d7a0d0c34ab4b7365ea01807d050 |  neutron   |   True  |  neutron@localhost   |
| 2ee967afb425485382b9b0e6ae41045c |    nova    |   True  |    nova@localhost    |
| 770b9f03cfd84baaabc7f3eb05ce2ce5 |  tiantao   |   True  |                      |
| 34fae764281f47d58a7a12af75911eec |   wangyf   |   True  |                      |
| 8b39d4e2b3f84d6b87c3900087fa6c30 | wujiasheng |   True  |                      |
| 895182ab8edb45d583c7b6c588f2625a |  yangning  |   True  |                      |
+----------------------------------+------------+---------+----------------------+


### python neutron client    http://pypi.douban.com/simple
[root@cfhost01 ~(keystone_admin)]# mkdir .pip
[root@cfhost01 ~(keystone_admin)]# cd .pip
[root@cfhost01 .pip(keystone_admin)]# vim pip.conf
[root@cfhost01 .pip(keystone_admin)]# pip install --upgrade python-neutronclient
Downloading/unpacking python-neutronclient from https://pypi.python.org/packages/source/p/python-neutronclient/python-neutronclient-2.3.9.tar.gz#md5=38e49fa744c482883ecd27863b2b2287
  Downloading python-neutronclient-2.3.9.tar.gz (130kB): 130kB downloaded
  Running setup.py egg_info for package python-neutronclient
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no files found matching '*' under directory 'tests'
Downloading/unpacking pbr>=0.6,!=0.7,<1.0 from https://pypi.python.org/packages/source/p/pbr/pbr-0.10.0.tar.gz#md5=9e02dbfb5e49210c381fd4eea00cf7b7 (from python-neutronclient)
  Downloading pbr-0.10.0.tar.gz (77kB): 77kB downloaded
  Running setup.py egg_info for package pbr
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
Downloading/unpacking argparse from https://pypi.python.org/packages/source/a/argparse/argparse-1.2.2.tar.gz#md5=38589b29d9120b19dfca32f86406a1f5 (from python-neutronclient)
  Downloading argparse-1.2.2.tar.gz (68kB): 68kB downloaded
  Running setup.py egg_info for package argparse
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files matching '*.pyo' found anywhere in distribution
    warning: no previously-included files matching '*.orig' found anywhere in distribution
    warning: no previously-included files matching '*.rej' found anywhere in distribution
    no previously-included directories found matching 'doc/_build'
    no previously-included directories found matching 'env24'
    no previously-included directories found matching 'env25'
    no previously-included directories found matching 'env26'
    no previously-included directories found matching 'env27'
Downloading/unpacking cliff>=1.6.0 (from python-neutronclient)
  Downloading cliff-1.8.0.tar.gz (44kB): 44kB downloaded
  Running setup.py egg_info for package cliff
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no files found matching '*.rst' under directory 'docs'
    warning: no files found matching '*.py' under directory 'docs'
    warning: no files found matching '*.html' under directory 'docs'
    warning: no files found matching '*.css' under directory 'docs'
    warning: no files found matching '*.js' under directory 'docs'
    warning: no files found matching '*.png' under directory 'docs'
    warning: no files found matching '*.txt' under directory 'docs'
    warning: no files found matching '*.py' under directory 'tests'
Requirement already up-to-date: iso8601>=0.1.9 in /usr/lib/python2.6/site-packages (from python-neutronclient)

^COperation cancelled by user
Storing complete log in /root/.pip/pip.log


[root@cfhost01 .pip(keystone_admin)]# pip install -i http://pypi.douban.com/simple --upgrade python-neutronclient
Downloading/unpacking python-neutronclient from http://pypi.douban.com/packages/source/p/python-neutronclient/python-neutronclient-2.3.9.tar.gz#md5=38e49fa744c482883ecd27863b2b2287
  Running setup.py egg_info for package python-neutronclient
    [pbr] Reusing existing SOURCES.txt
Downloading/unpacking pbr>=0.6,!=0.7,<1.0 from http://pypi.douban.com/packages/source/p/pbr/pbr-0.10.0.tar.gz#md5=9e02dbfb5e49210c381fd4eea00cf7b7 (from python-neutronclient)
  Running setup.py egg_info for package pbr
    [pbr] Reusing existing SOURCES.txt
Downloading/unpacking argparse from http://pypi.douban.com/packages/source/a/argparse/argparse-1.2.2.tar.gz#md5=38589b29d9120b19dfca32f86406a1f5 (from python-neutronclient)
  Running setup.py egg_info for package argparse
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files matching '*.pyo' found anywhere in distribution
    warning: no previously-included files matching '*.orig' found anywhere in distribution
    warning: no previously-included files matching '*.rej' found anywhere in distribution
    no previously-included directories found matching 'doc/_build'
    no previously-included directories found matching 'env24'
    no previously-included directories found matching 'env25'
    no previously-included directories found matching 'env26'
    no previously-included directories found matching 'env27'
Downloading/unpacking cliff>=1.6.0 (from python-neutronclient)
  Running setup.py egg_info for package cliff
    [pbr] Reusing existing SOURCES.txt
Requirement already up-to-date: iso8601>=0.1.9 in /usr/lib/python2.6/site-packages (from python-neutronclient)
Downloading/unpacking netaddr>=0.7.12 (from python-neutronclient)
  Downloading netaddr-0.7.12.tar.gz (1.5MB): 1.5MB downloaded
  Running setup.py egg_info for package netaddr
    warning: no previously-included files matching '*.svn*' found anywhere in distribution
    warning: no previously-included files matching '*.git*' found anywhere in distribution
Downloading/unpacking requests>=1.2.1,!=2.4.0 (from python-neutronclient)
  Downloading requests-2.4.3.tar.gz (438kB): 438kB downloaded
  Running setup.py egg_info for package requests
Downloading/unpacking python-keystoneclient>=0.10.0 (from python-neutronclient)
  Downloading python-keystoneclient-0.11.2.tar.gz (325kB): 325kB downloaded
  Running setup.py egg_info for package python-keystoneclient
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no files found matching 'HACKING'
    warning: no files found matching '*' under directory 'tests'
Downloading/unpacking simplejson>=2.2.0 (from python-neutronclient)
  Downloading simplejson-3.6.5.tar.gz (73kB): 73kB downloaded
  Running setup.py egg_info for package simplejson
Downloading/unpacking six>=1.7.0 (from python-neutronclient)
  Downloading six-1.8.0.tar.gz
  Running setup.py egg_info for package six
    no previously-included directories found matching 'documentation/_build'
Downloading/unpacking Babel>=1.3 (from python-neutronclient)
  Downloading Babel-1.3.tar.gz (3.4MB): 3.4MB downloaded
  Running setup.py egg_info for package Babel
    warning: no previously-included files matching '*' found under directory 'docs/_build'
    warning: no previously-included files matching '*.pyc' found under directory 'tests'
    warning: no previously-included files matching '*.pyo' found under directory 'tests'
Downloading/unpacking pip from http://pypi.douban.com/packages/source/p/pip/pip-1.5.6.tar.gz#md5=01026f87978932060cc86c1dc527903e (from pbr>=0.6,!=0.7,<1.0->python-neutronclient)
  Downloading pip-1.5.6.tar.gz (938kB): 938kB downloaded
  Running setup.py egg_info for package pip
    warning: no files found matching 'pip/cacert.pem'
    warning: no files found matching '*.html' under directory 'docs'
    warning: no previously-included files matching '*.rst' found under directory 'docs/_build'
    no previously-included directories found matching 'docs/_build/_sources'
Downloading/unpacking cmd2>=0.6.7 (from cliff>=1.6.0->python-neutronclient)
  Downloading cmd2-0.6.7.tar.gz
  Running setup.py egg_info for package cmd2
Requirement already up-to-date: PrettyTable>=0.7,<0.8 in /usr/lib/python2.6/site-packages (from cliff>=1.6.0->python-neutronclient)
Downloading/unpacking pyparsing>=2.0.1 (from cliff>=1.6.0->python-neutronclient)
  Downloading pyparsing-2.0.3.tar.gz (1.5MB): 1.5MB downloaded
  Running setup.py egg_info for package pyparsing
Downloading/unpacking stevedore>=0.14 from http://pypi.douban.com/packages/source/s/stevedore/stevedore-1.1.0.tar.gz#md5=b7f30055c32410f8f9b6cf1b55bdc68a (from cliff>=1.6.0->python-neutronclient)
  Downloading stevedore-1.1.0.tar.gz (352kB): 352kB downloaded
  Running setup.py egg_info for package stevedore
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no files found matching '*.html' under directory 'docs'
    warning: no files found matching '*.css' under directory 'docs'
    warning: no files found matching '*.js' under directory 'docs'
    warning: no files found matching '*.png' under directory 'docs'
    warning: no files found matching '*.py' under directory 'tests'
Downloading/unpacking oslo.config>=1.4.0 (from python-keystoneclient>=0.10.0->python-neutronclient)
  Downloading oslo.config-1.4.0.tar.gz (66kB): 66kB downloaded
  Running setup.py egg_info for package oslo.config
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
Downloading/unpacking oslo.serialization>=1.0.0 (from python-keystoneclient>=0.10.0->python-neutronclient)
  Downloading oslo.serialization-1.0.0.tar.gz
  Running setup.py egg_info for package oslo.serialization
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
Downloading/unpacking oslo.utils>=1.0.0 (from python-keystoneclient>=0.10.0->python-neutronclient)
  Downloading oslo.utils-1.0.0.tar.gz
  Running setup.py egg_info for package oslo.utils
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
Downloading/unpacking pytz>=0a from http://pypi.douban.com/packages/source/p/pytz/pytz-2014.9.tar.bz2#md5=69d0273f9a1ea53adcc8a114bc43d702 (from Babel>=1.3->python-neutronclient)
  Downloading pytz-2014.9.tar.bz2 (166kB): 166kB downloaded
  Running setup.py egg_info for package pytz
Downloading/unpacking oslo.i18n>=0.2.0 (from oslo.utils>=1.0.0->python-keystoneclient>=0.10.0->python-neutronclient)
  Downloading oslo.i18n-1.0.0.tar.gz
  Running setup.py egg_info for package oslo.i18n
    [pbr] Processing SOURCES.txt
    warning: LocalManifestMaker: standard file '-c' not found
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files found matching '.gitignore'
    warning: no previously-included files found matching '.gitreview'
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
Installing collected packages: python-neutronclient, pbr, argparse, cliff, netaddr, requests, python-keystoneclient, simplejson, six, Babel, pip, cmd2, pyparsing, stevedore, oslo.config, oslo.serialization, oslo.utils, pytz, oslo.i18n
  Found existing installation: python-neutronclient 2.3.4
    Uninstalling python-neutronclient:
      Successfully uninstalled python-neutronclient
  Running setup.py install for python-neutronclient
    [pbr] Reusing existing SOURCES.txt
    Installing neutron script to /usr/bin
  Found existing installation: pbr 0.8.0
    Uninstalling pbr:
      Successfully uninstalled pbr
  Running setup.py install for pbr
    [pbr] Reusing existing SOURCES.txt
  Found existing installation: argparse 1.2.1
    Uninstalling argparse:
      Successfully uninstalled argparse
  Running setup.py install for argparse
    warning: no previously-included files matching '*.pyc' found anywhere in distribution
    warning: no previously-included files matching '*.pyo' found anywhere in distribution
    warning: no previously-included files matching '*.orig' found anywhere in distribution
    warning: no previously-included files matching '*.rej' found anywhere in distribution
    no previously-included directories found matching 'doc/_build'
    no previously-included directories found matching 'env24'
    no previously-included directories found matching 'env25'
    no previously-included directories found matching 'env26'
    no previously-included directories found matching 'env27'
  Found existing installation: cliff 1.4.4
    Uninstalling cliff:
      Successfully uninstalled cliff
  Running setup.py install for cliff
    [pbr] Reusing existing SOURCES.txt
  Found existing installation: netaddr 0.7.5
    Uninstalling netaddr:
      Successfully uninstalled netaddr
  Running setup.py install for netaddr
    changing mode of build/scripts-2.6/netaddr from 644 to 755
    warning: no previously-included files matching '*.svn*' found anywhere in distribution
    warning: no previously-included files matching '*.git*' found anywhere in distribution
    changing mode of /usr/bin/netaddr to 755
  Found existing installation: requests 1.1.0
    Uninstalling requests:
      Successfully uninstalled requests
  Running setup.py install for requests
  Found existing installation: python-keystoneclient 0.9.0
    Uninstalling python-keystoneclient:
      Successfully uninstalled python-keystoneclient
  Running setup.py install for python-keystoneclient
    [pbr] Reusing existing SOURCES.txt
    Installing keystone script to /usr/bin
  Found existing installation: simplejson 2.0.9
    Uninstalling simplejson:
      Successfully uninstalled simplejson
  Running setup.py install for simplejson
    building 'simplejson._speedups' extension
    gcc -pthread -fno-strict-aliasing -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -m64 -mtune=generic -D_GNU_SOURCE -fPIC -fwrapv -DNDEBUG -O2 -g -pipe -Wall -Wp,-D_FORTIFY_SOURCE=2 -fexceptions -fstack-protector --param=ssp-buffer-size=4 -m64 -mtune=generic -D_GNU_SOURCE -fPIC -fwrapv -fPIC -I/usr/include/python2.6 -c simplejson/_speedups.c -o build/temp.linux-x86_64-2.6/simplejson/_speedups.o
    unable to execute gcc: No such file or directory
    ***************************************************************************
    WARNING: The C extension could not be compiled, speedups are not enabled.
    Failure information, if any, is above.
    I'm retrying the build without the C extension now.
    ***************************************************************************
    ***************************************************************************
    WARNING: The C extension could not be compiled, speedups are not enabled.
    Plain-Python installation succeeded.
    ***************************************************************************
  Found existing installation: six 1.6.1
    Uninstalling six:
      Successfully uninstalled six
  Running setup.py install for six
    no previously-included directories found matching 'documentation/_build'
  Found existing installation: Babel 0.9.4
    Uninstalling Babel:
      Successfully uninstalled Babel
  Running setup.py install for Babel
    warning: no previously-included files matching '*' found under directory 'docs/_build'
    warning: no previously-included files matching '*.pyc' found under directory 'tests'
    warning: no previously-included files matching '*.pyo' found under directory 'tests'
    Installing pybabel script to /usr/bin
  Found existing installation: pip 1.3.1
    Uninstalling pip:
      Successfully uninstalled pip
  Running setup.py install for pip
    warning: no files found matching 'pip/cacert.pem'
    warning: no files found matching '*.html' under directory 'docs'
    warning: no previously-included files matching '*.rst' found under directory 'docs/_build'
    no previously-included directories found matching 'docs/_build/_sources'
    Installing pip script to /usr/bin
    Installing pip2.6 script to /usr/bin
    Installing pip2 script to /usr/bin
  Found existing installation: cmd2 0.6.4
    Uninstalling cmd2:
      Successfully uninstalled cmd2
  Running setup.py install for cmd2
  Found existing installation: pyparsing 1.5.6
    Uninstalling pyparsing:
      Successfully uninstalled pyparsing
  Running setup.py install for pyparsing
  Found existing installation: stevedore 0.14
    Uninstalling stevedore:
      Successfully uninstalled stevedore
  Running setup.py install for stevedore
    [pbr] Reusing existing SOURCES.txt
  Found existing installation: oslo.config 1.2.1
    Uninstalling oslo.config:
      Successfully uninstalled oslo.config
  Running setup.py install for oslo.config
    [pbr] Reusing existing SOURCES.txt
    Skipping installation of /usr/lib/python2.6/site-packages/oslo/__init__.py (namespace package)
    Installing /usr/lib/python2.6/site-packages/oslo.config-1.4.0-py2.6-nspkg.pth
    Installing oslo-config-generator script to /usr/bin
  Running setup.py install for oslo.serialization
    [pbr] Reusing existing SOURCES.txt
    Skipping installation of /usr/lib/python2.6/site-packages/oslo/__init__.py (namespace package)
    Installing /usr/lib/python2.6/site-packages/oslo.serialization-1.0.0-py2.6-nspkg.pth
  Running setup.py install for oslo.utils
    [pbr] Reusing existing SOURCES.txt
    Skipping installation of /usr/lib/python2.6/site-packages/oslo/__init__.py (namespace package)
    Installing /usr/lib/python2.6/site-packages/oslo.utils-1.0.0-py2.6-nspkg.pth
  Found existing installation: pytz 2010h
    Uninstalling pytz:
      Successfully uninstalled pytz
  Running setup.py install for pytz
  Running setup.py install for oslo.i18n
    [pbr] Reusing existing SOURCES.txt
    Skipping installation of /usr/lib/python2.6/site-packages/oslo/__init__.py (namespace package)
    Installing /usr/lib/python2.6/site-packages/oslo.i18n-1.0.0-py2.6-nspkg.pth
Successfully installed python-neutronclient pbr argparse cliff netaddr requests python-keystoneclient simplejson six Babel pip cmd2 pyparsing stevedore oslo.config oslo.serialization oslo.utils pytz oslo.i18n
Cleaning up...



  
[Administrator@lenovo-9d779749 com.cygwin(keystone_demo)]$ pip install python-keystoneclient
Requirement already satisfied (use --upgrade to upgrade): python-keystoneclient in /usr/lib/python2.7/site-packages
Requirement already satisfied (use --upgrade to upgrade): pbr>=0.6,!=0.7,<1.0 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): argparse in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): Babel>=1.3 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): iso8601>=0.1.9 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): netaddr>=0.7.12 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): oslo.config>=1.4.0 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): oslo.serialization>=1.0.0 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): oslo.utils>=1.0.0 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): PrettyTable>=0.7,<0.8 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): requests>=2.2.0,!=2.4.0 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): six>=1.7.0 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): stevedore>=1.0.0 in /usr/lib/python2.7/site-packages (from python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): pip in /usr/lib/python2.7/site-packages (from pbr>=0.6,!=0.7,<1.0->python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): pytz>=0a in /usr/lib/python2.7/site-packages (from Babel>=1.3->python-keystoneclient)
Requirement already satisfied (use --upgrade to upgrade): oslo.i18n>=0.2.0 in /usr/lib/python2.7/site-packages (from oslo.utils>=1.0.0->python-keystoneclient)
Cleaning up...
 