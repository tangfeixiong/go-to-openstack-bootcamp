在Windows XP下设置Python OpenStack开发环境
==========================================
.. contents::

关于Windows 7/8环境
-------------------
由于Windows 7/8操作系统中，对用户账户有了严格的管理（如"以管理员身份运行"，类似Linux下的sudo命令）。
因此，以下操作可能并不符合Windows7/8环境。可以baidu或bing搜索解决方案。

安装Python27
------------
* 下载
    https://www.python.org/downloads/release/python-278/
	* `Windows x86 MSI Installer (2.7.8) <https://www.python.org/ftp/python/2.7.8/python-2.7.8.msi>`_
	* `Windows X86-64 MSI Installer (2.7.8) <https://www.python.org/ftp/python/2.7.8/python-2.7.8.amd64.msi>`_
    * `在线文档 <https://docs.python.org/2.7/>`_

* 安装
在Windows XP下安装了2个Python程序，在Windows程序菜单里，可以运行窗口版的Python和本地文档。 (windows batch)::

	F:\Python27>dir py*.exe
	 驱动器 F 中的卷没有标签。
	 卷的序列号是 804A-40B6

	 F:\Python27 的目录

	2013-05-15  22:43            26,624 python.exe
	2013-05-15  22:43            27,136 pythonw.exe
				   2 个文件         53,760 字节
				   0 个目录  1,984,139,264 可用字节	
				   
	F:\Python27>python
	Python 2.7.5 (default, May 15 2013, 22:43:36) [MSC v.1500 32 bit (Intel)] on win
	32
	Type "help", "copyright", "credits" or "license" for more information.
	>>> import sys
	>>> sys.path
	['', 'C:\\WINDOWS\\system32\\python27.zip', 'F:\\Python27\\DLLs', 'F:\\Python27\
	\lib', 'F:\\Python27\\lib\\plat-win', 'F:\\Python27\\lib\\lib-tk', 'F:\\Python27
	', 'F:\\Python27\\lib\\site-packages']
	>>> quit()

安装pip
-------
pip用于在线安装Pypi仓库里的开发工具（API，SDK等）

* 参考
    https://pip.pypa.io/en/latest/installing.html
	
    https://pypi.python.org/pypi/pip/
	
    http://stackoverflow.com/questions/4750806/how-to-install-pip-on-windows

* 下载
    * `get_pip.py <https://bootstrap.pypa.io/get-pip.py>`_
    * `pip-1.5.6.win-amd64-py2.7.exe, pip-1.5.6.win32-py2.7.exe <http://www.lfd.uci.edu/~gohlke/pythonlibs/#pip>`_

* 安装
使用get-pip.py安装 <batch>::

    F:\Python27>python G:\_python_tool\io.pypa.pip\get-pip.py
    Requirement already up-to-date: pip in f:\python27\lib\site-packages
    Cleaning up...
	
运行中会出现如下错误：UnicodeDecodeError: ‘ascii’ codec can’t decode byte 0xe5 in position 108: ordinal not in range(128).
解决方法参考：
    http://blog.csdn.net/mindmb/article/details/7898528
	
	http://blog.sina.com.cn/s/blog_6c39196501013s5b.html

在出现错误的文件设置utf-8编码

或运行pip-1.5.6.win-*32/64bit*-py2.7.exe，安装pip.exe程序(pip命令行工具），安装后 <batch>::

	F:\Python27\Scripts>dir
	 驱动器 F 中的卷没有标签。
	 卷的序列号是 804A-40B6

	 F:\Python27\Scripts 的目录

	2014-12-10  13:16    <DIR>          .
	2014-12-10  13:16    <DIR>          ..
	2014-12-10  11:04            91,496 easy_install-2.7.exe
	2014-12-10  11:04            91,496 easy_install.exe
	2014-12-10  09:57         1,340,903 get-pip.py
	2014-12-10  13:16            91,485 keystone.exe
	2014-12-10  13:16             1,210 netaddr
	2014-12-10  13:16            91,486 oslo-config-generator.exe
	2014-12-10  10:56               286 pip-script.py
	2014-12-10  10:56               364 pip-script.pyc
	2014-12-10  10:56               364 pip-script.pyo
	2014-12-10  10:56            65,536 pip.exe
	2014-12-10  10:56               560 pip.exe.manifest
	2014-12-10  10:56               288 pip2-script.py
	2014-12-10  10:56               366 pip2-script.pyc
	2014-12-10  10:56               366 pip2-script.pyo
	2014-12-10  10:56               292 pip2.7-script.py
	2014-12-10  10:56               370 pip2.7-script.pyc
	2014-12-10  10:56               370 pip2.7-script.pyo
	2014-12-10  10:56            65,536 pip2.7.exe
	2014-12-10  10:56               563 pip2.7.exe.manifest
	2014-12-10  10:56            65,536 pip2.exe
	2014-12-10  10:56               561 pip2.exe.manifest
	2014-12-10  13:16               306 pybabel-script.py
	2014-12-10  13:16            65,536 pybabel.exe
	2014-12-10  13:16               642 pybabel.exe.manifest
				  24 个文件      1,975,918 字节
				   2 个目录  1,984,135,168 可用字节

运行pip时的错误信息 (batch)::

    F:\Python27\Scripts>pip
    Traceback (most recent call last):
      File "C:\Python27\Scripts\pip-script.py", line 5, in <module>
        from pkg_resources import load_entry_point
    ImportError: No module named pkg_resources

原因待研究	

设置OpenStack开发环境
---------------------
以keystone client为例

* 安装python-keystoneclient API
安装失败，原因是pypi在美国，下载超时
<batch>::

	F:\Python27\Scripts>pip install python-keystoneclient
	Downloading/unpacking python-keystoneclient
	Downloading/unpacking iso8601>=0.1.9 (from python-keystoneclient)
	  Downloading iso8601-0.1.10.tar.gz
	  Running setup.py (path:c:\docume~1\admini~1\locals~1\temp\pip_build_Administra
	tor\iso8601\setup.py) egg_info for package iso8601

	Downloading/unpacking requests>=2.2.0,!=2.4.0 (from python-keystoneclient)
	Cleaning up...
	Exception:
	Traceback (most recent call last):
	  File "F:\Python27\lib\site-packages\pip\basecommand.py", line 122, in main
		status = self.run(options, args)
	  File "F:\Python27\lib\site-packages\pip\commands\install.py", line 278, in run

		requirement_set.prepare_files(finder, force_root_egg_info=self.bundle, bundl
	e=self.bundle)
	  File "F:\Python27\lib\site-packages\pip\req.py", line 1197, in prepare_files
		do_download,
	  File "F:\Python27\lib\site-packages\pip\req.py", line 1375, in unpack_url
		self.session,
	  File "F:\Python27\lib\site-packages\pip\download.py", line 572, in unpack_http
	_url
		download_hash = _download_url(resp, link, temp_location)
	  File "F:\Python27\lib\site-packages\pip\download.py", line 433, in _download_u
	rl
		for chunk in resp_read(4096):
	  File "F:\Python27\lib\site-packages\pip\download.py", line 421, in resp_read
		chunk_size, decode_content=False):
	  File "F:\Python27\lib\site-packages\pip\_vendor\requests\packages\urllib3\resp
	onse.py", line 240, in stream
		data = self.read(amt=amt, decode_content=decode_content)
	  File "F:\Python27\lib\site-packages\pip\_vendor\requests\packages\urllib3\resp
	onse.py", line 187, in read
		data = self._fp.read(amt)
	  File "F:\Python27\lib\httplib.py", line 567, in read
		s = self.fp.read(amt)
	  File "F:\Python27\lib\socket.py", line 380, in read
		data = self._sock.recv(left)
	  File "F:\Python27\lib\ssl.py", line 241, in recv
		return self.read(buflen)
	  File "F:\Python27\lib\ssl.py", line 160, in read
		return self._sslobj.read(len)
	SSLError: The read operation timed out

	Storing debug log for failure in C:\Documents and Settings\Administrator\pip\pip
	.log

改用douban的镜像，安装失败（网速的原因），再次执行后成功。如下

下载到argparse API时网络超时失败 <batch>::

	F:\Python27\Scripts>pip install -i http://pypi.douban.com/simple python-keystone
	client
	Downloading/unpacking python-keystoneclient
	  http://pypi.douban.com/simple/python-keystoneclient/ uses an insecure transpor
	t scheme (http). Consider using https if pypi.douban.com has it available
	Downloading/unpacking iso8601>=0.1.9 (from python-keystoneclient)
	  http://pypi.douban.com/simple/iso8601/ uses an insecure transport scheme (http
	). Consider using https if pypi.douban.com has it available
	  Downloading iso8601-0.1.10.tar.gz
	  Running setup.py (path:c:\docume~1\admini~1\locals~1\temp\pip_build_Administra
	tor\iso8601\setup.py) egg_info for package iso8601

	Downloading/unpacking requests>=2.2.0,!=2.4.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/requests/ uses an insecure transport scheme (htt
	p). Consider using https if pypi.douban.com has it available
	Downloading/unpacking Babel>=1.3 (from python-keystoneclient)
	  http://pypi.douban.com/simple/Babel/ uses an insecure transport scheme (http).
	 Consider using https if pypi.douban.com has it available
	  Running setup.py (path:c:\docume~1\admini~1\locals~1\temp\pip_build_Administra
	tor\Babel\setup.py) egg_info for package Babel

		warning: no previously-included files matching '*' found under directory 'do
	cs\_build'
		warning: no previously-included files matching '*.pyc' found under directory
	 'tests'
		warning: no previously-included files matching '*.pyo' found under directory
	 'tests'
	Downloading/unpacking netaddr>=0.7.12 (from python-keystoneclient)
	  http://pypi.douban.com/simple/netaddr/ uses an insecure transport scheme (http
	). Consider using https if pypi.douban.com has it available
	  Running setup.py (path:c:\docume~1\admini~1\locals~1\temp\pip_build_Administra
	tor\netaddr\setup.py) egg_info for package netaddr

		warning: no previously-included files matching '*.svn*' found anywhere in di
	stribution
		warning: no previously-included files matching '*.git*' found anywhere in di
	stribution
	Downloading/unpacking six>=1.7.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/six/ uses an insecure transport scheme (http). C
	onsider using https if pypi.douban.com has it available
	  Downloading six-1.8.0-py2.py3-none-any.whl
	Downloading/unpacking oslo.config>=1.4.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/oslo.config/ uses an insecure transport scheme (
	http). Consider using https if pypi.douban.com has it available
	  Downloading oslo.config-1.5.0-py2.py3-none-any.whl
	Downloading/unpacking pbr>=0.6,!=0.7,<1.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/pbr/ uses an insecure transport scheme (http). C
	onsider using https if pypi.douban.com has it available
	Downloading/unpacking oslo.utils>=1.0.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/oslo.utils/ uses an insecure transport scheme (h
	ttp). Consider using https if pypi.douban.com has it available
	  Downloading oslo.utils-1.1.0-py2.py3-none-any.whl
	Downloading/unpacking oslo.serialization>=1.0.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/oslo.serialization/ uses an insecure transport s
	cheme (http). Consider using https if pypi.douban.com has it available
	  Downloading oslo.serialization-1.1.0-py2.py3-none-any.whl
	Downloading/unpacking stevedore>=1.0.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/stevedore/ uses an insecure transport scheme (ht
	tp). Consider using https if pypi.douban.com has it available
	  Downloading stevedore-1.1.0-py2.py3-none-any.whl
	Downloading/unpacking PrettyTable>=0.7,<0.8 (from python-keystoneclient)
	  http://pypi.douban.com/simple/PrettyTable/ uses an insecure transport scheme (
	http). Consider using https if pypi.douban.com has it available
	  Downloading prettytable-0.7.2.zip
	  Running setup.py (path:c:\docume~1\admini~1\locals~1\temp\pip_build_Administra
	tor\PrettyTable\setup.py) egg_info for package PrettyTable

	Downloading/unpacking argparse (from python-keystoneclient)
	  http://pypi.douban.com/simple/argparse/ uses an insecure transport scheme (htt
	p). Consider using https if pypi.douban.com has it available
	  Downloading argparse-1.2.2-py2.py3-none-any.whl
	Downloading/unpacking pytz>=0a (from Babel>=1.3->python-keystoneclient)
	  http://pypi.douban.com/simple/pytz/ uses an insecure transport scheme (http).
	Consider using https if pypi.douban.com has it available
	Cleaning up...
	Exception:
	Traceback (most recent call last):
	  File "F:\Python27\lib\site-packages\pip\basecommand.py", line 122, in main
		status = self.run(options, args)
	  File "F:\Python27\lib\site-packages\pip\commands\install.py", line 278, in run

		requirement_set.prepare_files(finder, force_root_egg_info=self.bundle, bundl
	e=self.bundle)
	  File "F:\Python27\lib\site-packages\pip\req.py", line 1197, in prepare_files
		do_download,
	  File "F:\Python27\lib\site-packages\pip\req.py", line 1375, in unpack_url
		self.session,
	  File "F:\Python27\lib\site-packages\pip\download.py", line 572, in unpack_http
	_url
		download_hash = _download_url(resp, link, temp_location)
	  File "F:\Python27\lib\site-packages\pip\download.py", line 433, in _download_u
	rl
		for chunk in resp_read(4096):
	  File "F:\Python27\lib\site-packages\pip\download.py", line 421, in resp_read
		chunk_size, decode_content=False):
	  File "F:\Python27\lib\site-packages\pip\_vendor\requests\packages\urllib3\resp
	onse.py", line 240, in stream
		data = self.read(amt=amt, decode_content=decode_content)
	  File "F:\Python27\lib\site-packages\pip\_vendor\requests\packages\urllib3\resp
	onse.py", line 187, in read
		data = self._fp.read(amt)
	  File "F:\Python27\lib\httplib.py", line 567, in read
		s = self.fp.read(amt)
	  File "F:\Python27\lib\socket.py", line 380, in read
		data = self._sock.recv(left)
	timeout: timed out

	Storing debug log for failure in C:\Documents and Settings\Administrator\pip\pip
	.log

成功 <batch>::
	
	F:\Python27\Scripts>pip install -i http://pypi.douban.com/simple python-keystone
	client
	Downloading/unpacking python-keystoneclient
	  http://pypi.douban.com/simple/python-keystoneclient/ uses an insecure transpor
	t scheme (http). Consider using https if pypi.douban.com has it available
	Downloading/unpacking iso8601>=0.1.9 (from python-keystoneclient)
	  http://pypi.douban.com/simple/iso8601/ uses an insecure transport scheme (http
	). Consider using https if pypi.douban.com has it available
	  Downloading iso8601-0.1.10.tar.gz
	  Running setup.py (path:c:\docume~1\admini~1\locals~1\temp\pip_build_Administra
	tor\iso8601\setup.py) egg_info for package iso8601

	Downloading/unpacking requests>=2.2.0,!=2.4.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/requests/ uses an insecure transport scheme (htt
	p). Consider using https if pypi.douban.com has it available
	Downloading/unpacking Babel>=1.3 (from python-keystoneclient)
	  http://pypi.douban.com/simple/Babel/ uses an insecure transport scheme (http).
	 Consider using https if pypi.douban.com has it available
	  Running setup.py (path:c:\docume~1\admini~1\locals~1\temp\pip_build_Administra
	tor\Babel\setup.py) egg_info for package Babel

		warning: no previously-included files matching '*' found under directory 'do
	cs\_build'
		warning: no previously-included files matching '*.pyc' found under directory
	 'tests'
		warning: no previously-included files matching '*.pyo' found under directory
	 'tests'
	Downloading/unpacking netaddr>=0.7.12 (from python-keystoneclient)
	  http://pypi.douban.com/simple/netaddr/ uses an insecure transport scheme (http
	). Consider using https if pypi.douban.com has it available
	  Running setup.py (path:c:\docume~1\admini~1\locals~1\temp\pip_build_Administra
	tor\netaddr\setup.py) egg_info for package netaddr

		warning: no previously-included files matching '*.svn*' found anywhere in di
	stribution
		warning: no previously-included files matching '*.git*' found anywhere in di
	stribution
	Downloading/unpacking six>=1.7.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/six/ uses an insecure transport scheme (http). C
	onsider using https if pypi.douban.com has it available
	  Downloading six-1.8.0-py2.py3-none-any.whl
	Downloading/unpacking oslo.config>=1.4.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/oslo.config/ uses an insecure transport scheme (
	http). Consider using https if pypi.douban.com has it available
	  Downloading oslo.config-1.5.0-py2.py3-none-any.whl
	Downloading/unpacking pbr>=0.6,!=0.7,<1.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/pbr/ uses an insecure transport scheme (http). C
	onsider using https if pypi.douban.com has it available
	Downloading/unpacking oslo.utils>=1.0.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/oslo.utils/ uses an insecure transport scheme (h
	ttp). Consider using https if pypi.douban.com has it available
	  Downloading oslo.utils-1.1.0-py2.py3-none-any.whl
	Downloading/unpacking oslo.serialization>=1.0.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/oslo.serialization/ uses an insecure transport s
	cheme (http). Consider using https if pypi.douban.com has it available
	  Downloading oslo.serialization-1.1.0-py2.py3-none-any.whl
	Downloading/unpacking stevedore>=1.0.0 (from python-keystoneclient)
	  http://pypi.douban.com/simple/stevedore/ uses an insecure transport scheme (ht
	tp). Consider using https if pypi.douban.com has it available
	  Downloading stevedore-1.1.0-py2.py3-none-any.whl
	Downloading/unpacking PrettyTable>=0.7,<0.8 (from python-keystoneclient)
	  http://pypi.douban.com/simple/PrettyTable/ uses an insecure transport scheme (
	http). Consider using https if pypi.douban.com has it available
	  Downloading prettytable-0.7.2.zip
	  Running setup.py (path:c:\docume~1\admini~1\locals~1\temp\pip_build_Administra
	tor\PrettyTable\setup.py) egg_info for package PrettyTable

	Downloading/unpacking argparse (from python-keystoneclient)
	  http://pypi.douban.com/simple/argparse/ uses an insecure transport scheme (htt
	p). Consider using https if pypi.douban.com has it available
	  Downloading argparse-1.2.2-py2.py3-none-any.whl
	Downloading/unpacking pytz>=0a (from Babel>=1.3->python-keystoneclient)
	  http://pypi.douban.com/simple/pytz/ uses an insecure transport scheme (http).
	Consider using https if pypi.douban.com has it available
	Requirement already satisfied (use --upgrade to upgrade): pip in f:\python27\lib
	\site-packages (from pbr>=0.6,!=0.7,<1.0->python-keystoneclient)
	Downloading/unpacking oslo.i18n>=1.0.0 (from oslo.utils>=1.0.0->python-keystonec
	lient)
	  http://pypi.douban.com/simple/oslo.i18n/ uses an insecure transport scheme (ht
	tp). Consider using https if pypi.douban.com has it available
	  Downloading oslo.i18n-1.1.0-py2.py3-none-any.whl
	Downloading/unpacking netifaces>=0.10.4 (from oslo.utils>=1.0.0->python-keystone
	client)
	  http://pypi.douban.com/simple/netifaces/ uses an insecure transport scheme (ht
	tp). Consider using https if pypi.douban.com has it available
	  Downloading netifaces-0.10.4-cp27-none-win32.whl
	Installing collected packages: python-keystoneclient, iso8601, requests, Babel,
	netaddr, six, oslo.config, pbr, oslo.utils, oslo.serialization, stevedore, Prett
	yTable, argparse, pytz, oslo.i18n, netifaces
	  Running setup.py install for iso8601

	  Running setup.py install for Babel

		warning: no previously-included files matching '*' found under directory 'do
	cs\_build'
		warning: no previously-included files matching '*.pyc' found under directory
	 'tests'
		warning: no previously-included files matching '*.pyo' found under directory
	 'tests'
		Installing pybabel-script.py script to F:\Python27\Scripts
		Installing pybabel.exe script to F:\Python27\Scripts
		Installing pybabel.exe.manifest script to F:\Python27\Scripts
	  Running setup.py install for netaddr

		warning: no previously-included files matching '*.svn*' found anywhere in di
	stribution
		warning: no previously-included files matching '*.git*' found anywhere in di
	stribution
	  Running setup.py install for PrettyTable

	Successfully installed python-keystoneclient iso8601 requests Babel netaddr six
	oslo.config pbr oslo.utils oslo.serialization stevedore PrettyTable argparse pyt
	z oslo.i18n netifaces
	Cleaning up...

* 使用OpenStack开发环境测试python-keystoneclient
<batch>::

	F:\Python27\Scripts>keystone --version
	0.11.2

	F:\Python27\Scripts>keystone --os-auth-url "http://192.168.1.99:5000/v2.0" --os-
	tenant-name admin --os-username admin --os-password changeme1122 user-list
	+----------------------------------+----------+---------+----------------------+

	|                id                |   name   | enabled |        email         |

	+----------------------------------+----------+---------+----------------------+

	| 9c7e015587264e23b0f16d7857f199b1 |  admin   |   True  |                      |

	| 5eb56929a8964ee3ac2291837a096cf9 | alt_demo |   True  | alt_demo@example.com |

	| ed568f376a8f4293993129ac11041cda |  cinder  |   True  |                      |

	| f26cc8c7d7444636b194c453077646b3 |   demo   |   True  |   demo@example.com   |

	| 56ef0cbb54e24368b3e061cfebc68643 |  glance  |   True  |                      |

	| a46ba4008278439db4bba37eca3fd2ed |   heat   |   True  |                      |

	| 25dea5f2010643f99ca7fb7facaecefb | neutron  |   True  |                      |

	| 26ccbc63c7444850a8be1543db954949 |   nova   |   True  |                      |

	+----------------------------------+----------+---------+----------------------+










