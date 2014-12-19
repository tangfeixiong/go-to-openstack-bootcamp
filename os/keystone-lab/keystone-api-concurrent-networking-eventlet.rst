Keystone API concurrent networking - eventlet
=============================================
(python)::

	#代码eventlet-client-example.py
	#eventlet是高并发，无阻塞的事件驱动tcp io api，为openstack api所使用，这是eventlet网站上示例程序的简单修改，成为登录openstack keystone的简单客户端

	import eventlet
	from eventlet.green import urllib2
	import re, json


	urls = ["http://127.0.0.1:5000/v2.0/tokens", "http://python.org/images/python-logo.gif", "http://www.baidu.com"]

	#data = {"auth":{"identity":{"methods":["password"],"password":{"user":{"name":"admin","domain":{"id":"default"},"password":"changeme1122"}}}}}
	data = {"auth": {"tenantName": "admin", "passwordCredentials": {"username": "admin", "password": "changeme1122"}}}
	headers1 = {'User-Agent': 'Mozilla/5.0 (Windows NT 5.1; rv:35.0) Gecko/20100101 Firefox/35.0', 'Content-Type': 'application/json'}
	headers2 = {'User-Agent': 'python-keystoneclient', 'Content-Type': 'application/json'}

	def fetch(url):
		print("opening", url)
		body = None
		if (re.search(r'\:5000/v2\.0', url) != None) :
			d = json.dumps(data)
			req = urllib2.Request(url, d, headers1)
			f = urllib2.urlopen(req)
			body = f.read()
			print  '\n%(response)s\n' % {"response": body}
		else : 
			body = urllib2.urlopen(url).read()
		print("done with", url)
		return url, body

	pool = eventlet.GreenPool(200)
	for url, body in pool.imap(fetch, urls):
		print("got body from", url, "of length", len(body))

(bash)::
	
	#关掉我的笔记本的代理	
	Administrator@lenovo-9d779749 ~/github.com/tangfeixiong/learning-openstack-and-cloudfoundry/os/keystone-lab
	$ unset http_proxy https_proxy

	#运行出错，因为没有eventlet
	Administrator@lenovo-9d779749 ~/github.com/tangfeixiong/learning-openstack-and-cloudfoundry/os/keystone-lab
	$ python eventlet-client-example.py
	Traceback (most recent call last):
	  File "eventlet-client-example.py", line 1, in <module>
		import eventlet
	ImportError: No module named eventlet

	#用pip在线安装eventlet
	Administrator@lenovo-9d779749 ~/github.com/tangfeixiong/learning-openstack-and-cloudfoundry/os/keystone-lab
	$ pip install eventlet
	Downloading/unpacking eventlet
	  http://pypi.douban.com/simple/eventlet/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
	  Downloading eventlet-0.15.2-py2.py3-none-any.whl (150kB): 150kB downloaded
	Downloading/unpacking greenlet>=0.3 (from eventlet)
	  http://pypi.douban.com/simple/greenlet/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
	  Downloading greenlet-0.4.5.zip (77kB): 77kB downloaded
	  Running setup.py (path:/tmp/pip_build_Administrator/greenlet/setup.py) egg_info for package greenlet

	Installing collected packages: eventlet, greenlet
	  Running setup.py install for greenlet
		building 'greenlet' extension
		gcc -fno-strict-aliasing -ggdb -O2 -pipe -Wimplicit-function-declaration -fdebug-prefix-map=/usr/src/ports/python/python-2.7.8-1.i686/build=/usr/src/debug/python-2.7.8-1 -fdebug-prefix-map=/usr/src/ports/python/python-2.7.8-1.i686/src/Python-2.7.8=/usr/src/debug/python-2.7.8-1 -DNDEBUG -g -fwrapv -O3 -Wall -Wstrict-prototypes -I/usr/include/python2.7 -c greenlet.c -o build/temp.cygwin-1.7.32-i686-2.7/greenlet.o
		gcc -shared -Wl,--enable-auto-image-base -L. build/temp.cygwin-1.7.32-i686-2.7/greenlet.o -L/usr/lib/python2.7/config -L/usr/lib -lpython2.7 -o build/lib.cygwin-1.7.32-i686-2.7/greenlet.dll

	Successfully installed eventlet greenlet
	Cleaning up...

	#运行，同时请求3个url，演示高并发和无阻塞的特性
	Administrator@lenovo-9d779749 ~/github.com/tangfeixiong/learning-openstack-and-cloudfoundry/os/keystone-lab
	$ python eventlet-client-example.py
	('opening', 'http://127.0.0.1:5000/v2.0/tokens')
	('opening', 'http://python.org/images/python-logo.gif')
	('opening', 'http://www.baidu.com')
	('done with', 'http://www.baidu.com')

	{"access": {"token": {"issued_at": "2014-12-18T14:44:39.593750", "expires": "2014-12-18T15:44:39Z", "id": "05ac66a15e61442ba5c20fd331fa7737", "tenant": {"description": "Admin Tenant - cygwin", "enabled": true, "id": "996b6c042b8f43bcb07ace4122a2d5a4", "name": "admin"}, "audit_ids": ["qqHoWDfSRsGUMHtJUlb4tQ"]}, "serviceCatalog": [{"endpoints": [{"adminURL": "http://controller-cygwin.openstack-tangfx.local:35357/v2.0", "region": "regionOne", "internalURL": "http://controller-cygwin.openstack-tangfx.local:5000/v3", "id": "014a0d11e4164e76b6507489b1f24c22", "publicURL": "http://controller-cygwin.openstack-tangfx.local:5000/v3"}], "endpoints_links": [], "type": "identity", "name": "keystone"}], "user": {"username": "admin", "roles_links": [], "id": "fd857b0ed3f14fe5b6811e1116d17ade", "roles": [{"name": "_member_"}, {"name": "admin"}], "name": "admin"}, "metadata": {"is_admin": 0, "roles": ["72f9cffcfd6447afb1ac9eb2dec28b75", "91d7efa486d34a27b0f057d35c1f9534"]}}}

	('done with', 'http://127.0.0.1:5000/v2.0/tokens')
	('got body from', 'http://127.0.0.1:5000/v2.0/tokens', 'of length', 972)
	('done with', 'http://python.org/images/python-logo.gif')
	('got body from', 'http://python.org/images/python-logo.gif', 'of length', 2549)
	('got body from', 'http://www.baidu.com', 'of length', 85761)

	#在keystone服务端输出了log
	(stagingenv)
	Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/stagingenv
	$ mysqld_safe
	141218 20:58:12 mysqld_safe Logging to '/var/lib/mysql//lenovo-9d779749.err'.
	141218 20:58:15 mysqld_safe Starting mysqld daemon with databases from /var/lib/mysql/

	[1]+  Stopped                 mysqld_safe

	Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/stagingenv
	$ bin/venv-keystone-all
	2014-12-18 21:00:02.390 6096 WARNING keystone.openstack.common.versionutils [-] Deprecated: keystone.middleware.core.XmlBodyMiddleware is deprecated as of Icehouse in favor of support for "application/json" only and may be removed in Kilo.
	2014-12-18 21:00:03.187 6096 WARNING keystone.openstack.common.versionutils [-] Deprecated: keystone.contrib.revoke.backends.kvs is deprecated as of Juno in favor of keystone.contrib.revoke.backends.sql and may be removed in Kilo.
	2014-12-18 21:00:03.218 6096 INFO keystone.common.kvs.core [-] Using default dogpile sha1_mangle_key as KVS region os-revoke-driver key_mangler
	2014-12-18 21:00:03.421 6096 INFO keystone.common.environment.eventlet_server [-] Starting bin/venv-keystone-all on 0.0.0.0:35357
	2014-12-18 21:00:03.421 6096 INFO keystone.openstack.common.service [-] Starting 2 workers
	2014-12-18 21:00:03.640 6096 INFO keystone.openstack.common.service [-] Started child 4916
	2014-12-18 21:00:03.718 4916 INFO eventlet.wsgi.server [-] (4916) wsgi starting up on http://0.0.0.0:35357/
	2014-12-18 21:00:03.750 6096 INFO keystone.openstack.common.service [-] Started child 5476
	2014-12-18 21:00:03.765 6096 INFO keystone.common.environment.eventlet_server [-] Starting bin/venv-keystone-all on 0.0.0.0:5000
	2014-12-18 21:00:03.765 6096 INFO keystone.openstack.common.service [-] Starting 2 workers
	2014-12-18 21:00:03.828 5476 INFO eventlet.wsgi.server [-] (5476) wsgi starting up on http://0.0.0.0:35357/
	2014-12-18 21:00:03.875 6096 INFO keystone.openstack.common.service [-] Started child 2592
	2014-12-18 21:00:03.937 2592 INFO eventlet.wsgi.server [-] (2592) wsgi starting up on http://0.0.0.0:5000/
	2014-12-18 21:00:03.984 6096 INFO keystone.openstack.common.service [-] Started child 3368
	2014-12-18 21:00:04.046 3368 INFO eventlet.wsgi.server [-] (3368) wsgi starting up on http://0.0.0.0:5000/
	...显示收到了request
	2014-12-18 22:44:39.593 3368 INFO eventlet.wsgi.server [-] 127.0.0.1 - - [18/Dec/2014 22:44:39] "POST /v2.0/tokens HTTP/1.1" 200 1120 0.250000

	#与openstack的python-keystoneclient的执行对比
	Administrator@lenovo-9d779749 ~/github.com/tangfeixiong/learning-openstack-and-cloudfoundry/os/keystone-lab
	$ keystone --debug --os-auth-url=http://127.0.0.1:5000/v2.0 --os-username=admin --os-password=changeme1122 --os-tenant-name=admin user-list
	DEBUG:keystoneclient.auth.identity.v2:Making authentication request to http://127.0.0.1:5000/v2.0/tokens
	INFO:requests.packages.urllib3.connectionpool:Starting new HTTP connection (1): 127.0.0.1
	DEBUG:requests.packages.urllib3.connectionpool:"POST /v2.0/tokens HTTP/1.1" 200 972
	DEBUG:keystoneclient.session:REQ: curl -i -X GET http://controller-cygwin.openstack-tangfx.local:35357/v2.0/users -H "User-Agent: python-keystoneclient" -H "X-Auth-Token: {SHA1}eba1f1aeffe97076ec188a0c5df4d378a9ed4490"
	INFO:requests.packages.urllib3.connectionpool:Starting new HTTP connection (1): controller-cygwin.openstack-tangfx.local
	DEBUG:requests.packages.urllib3.connectionpool:"GET /v2.0/users HTTP/1.1" 200 276
	DEBUG:keystoneclient.session:RESP: [200] date: Thu, 18 Dec 2014 15:01:37 GMT vary: X-Auth-Token content-length: 276 content-type: application/json connection: keep-alive
	RESP BODY: {"users": [{"username": "admin", "name": "admin", "enabled": true, "email": "admin-cygwin@localhost", "id": "fd857b0ed3f14fe5b6811e1116d17ade"}, {"username": "demo", "name": "demo", "enabled": true, "email": "demo-cygwin@localhost", "id": "7a0e132d88684aeaa1fc7e8b74fbfd13"}]}

	+----------------------------------+-------+---------+------------------------+
	|                id                |  name | enabled |         email          |
	+----------------------------------+-------+---------+------------------------+
	| fd857b0ed3f14fe5b6811e1116d17ade | admin |   True  | admin-cygwin@localhost |
	| 7a0e132d88684aeaa1fc7e8b74fbfd13 |  demo |   True  | demo-cygwin@localhost  |
	+----------------------------------+-------+---------+------------------------+

#api网址 http://eventlet.net/ http://eventlet.net/doc/index.html