# OpenStack Pike Installation

## Table of content

控制节点
* [Keystone认证服务](#identity)

## Controller

### Identity

[Keystone](https://docs.openstack.org/keystone/pike/install/keystone-install-rdo.html)

Database
```
[vagrant@localhost ~]$ mysql -u root -e "CREATE DATABASE keystone;GRANT ALL PRIVILEGES ON keystone.* TO 'keystone'@'localhost' IDENTIFIED BY 'KEYSTONE_DBPASS';GRANT ALL PRIVILEGES ON keystone.* TO 'keystone'@'%' IDENTIFIED BY 'KEYSTONE_DBPASS';"
```

```
[vagrant@localhost ~]$ mysql -u root -e "SHOW DATABASES;"
+--------------------+
| Database           |
+--------------------+
| information_schema |
| keystone           |
| mysql              |
| performance_schema |
| test               |
+--------------------+
```

Install
```
[vagrant@localhost ~]$ sudo yum install -y openstack-keystone httpd mod-wsgi
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
No package mod-wsgi available.
Resolving Dependencies
--> Running transaction check
---> Package httpd.x86_64 0:2.4.6-67.el7.centos.6 will be installed
--> Processing Dependency: httpd-tools = 2.4.6-67.el7.centos.6 for package: httpd-2.4.6-67.el7.centos.6.x86_64
--> Processing Dependency: /etc/mime.types for package: httpd-2.4.6-67.el7.centos.6.x86_64
--> Processing Dependency: libaprutil-1.so.0()(64bit) for package: httpd-2.4.6-67.el7.centos.6.x86_64
--> Processing Dependency: libapr-1.so.0()(64bit) for package: httpd-2.4.6-67.el7.centos.6.x86_64
---> Package openstack-keystone.noarch 1:12.0.0-1.el7 will be installed
--> Processing Dependency: python-keystone = 1:12.0.0-1.el7 for package: 1:openstack-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-keystoneclient >= 1:3.8.0 for package: 1:openstack-keystone-12.0.0-1.el7.noarch
--> Running transaction check
---> Package apr.x86_64 0:1.4.8-3.el7 will be installed
---> Package apr-util.x86_64 0:1.5.2-6.el7 will be installed
---> Package httpd-tools.x86_64 0:2.4.6-67.el7.centos.6 will be installed
---> Package mailcap.noarch 0:2.1.41-2.el7 will be installed
---> Package python-keystone.noarch 1:12.0.0-1.el7 will be installed
--> Processing Dependency: python-webob >= 1.7.1 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-stevedore >= 1.20.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-sqlalchemy >= 1.0.10 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-six >= 1.9.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-routes >= 2.3.1 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-pysaml2 >= 2.4.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-pycadf >= 2.1.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-paste-deploy >= 1.5.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-passlib >= 1.7.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-osprofiler >= 1.4.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-utils >= 3.20.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-serialization >= 2.4.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-policy >= 1.23.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-middleware >= 3.27.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-messaging >= 5.24.2 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-log >= 3.22.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-i18n >= 3.4.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-db >= 4.24.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-context >= 2.14.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-config >= 2:4.0.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-concurrency >= 3.8.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oslo-cache >= 1.5.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-oauthlib >= 0.6 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-migrate >= 0.11.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-keystonemiddleware >= 4.12.0 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-dogpile-cache >= 0.6.2 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-babel >= 2.3.4 for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-scrypt for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-pbr for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-paste for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-msgpack for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-memcached for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-ldappool for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-ldap for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-jsonschema for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-cryptography for package: 1:python-keystone-12.0.0-1.el7.noarch
--> Processing Dependency: python-bcrypt for package: 1:python-keystone-12.0.0-1.el7.noarch
---> Package python2-keystoneclient.noarch 1:3.13.0-1.el7 will be installed
--> Processing Dependency: python-requests >= 2.10.0 for package: 1:python2-keystoneclient-3.13.0-1.el7.noarch
--> Processing Dependency: python-positional >= 1.1.1 for package: 1:python2-keystoneclient-3.13.0-1.el7.noarch
--> Processing Dependency: python-keystoneauth1 >= 3.1.0 for package: 1:python2-keystoneclient-3.13.0-1.el7.noarch
--> Processing Dependency: python-keyring >= 5.5.1 for package: 1:python2-keystoneclient-3.13.0-1.el7.noarch
--> Processing Dependency: python-debtcollector >= 1.2.0 for package: 1:python2-keystoneclient-3.13.0-1.el7.noarch
--> Running transaction check
---> Package python-dogpile-cache.noarch 0:0.6.2-1.el7 will be installed
--> Processing Dependency: python-dogpile-core >= 0.4.1 for package: python-dogpile-cache-0.6.2-1.el7.noarch
---> Package python-jsonschema.noarch 0:2.3.0-1.el7 will be installed
---> Package python-keyring.noarch 0:5.7.1-1.el7 will be installed
---> Package python-ldap.x86_64 0:2.4.15-2.el7 will be installed
---> Package python-ldappool.noarch 0:1.0-4.el7 will be installed
---> Package python-memcached.noarch 0:1.58-1.el7 will be installed
---> Package python-migrate.noarch 0:0.11.0-1.el7 will be installed
--> Processing Dependency: python-tempita >= 0.4 for package: python-migrate-0.11.0-1.el7.noarch
--> Processing Dependency: python-sqlparse for package: python-migrate-0.11.0-1.el7.noarch
--> Processing Dependency: python-setuptools for package: python-migrate-0.11.0-1.el7.noarch
---> Package python-msgpack.x86_64 0:0.4.6-3.el7 will be installed
---> Package python-paste.noarch 0:1.7.5.1-9.20111221hg1498.el7 will be installed
--> Processing Dependency: pyOpenSSL for package: python-paste-1.7.5.1-9.20111221hg1498.el7.noarch
---> Package python-paste-deploy.noarch 0:1.5.2-6.el7 will be installed
---> Package python-routes.noarch 0:2.4.1-1.el7 will be installed
--> Processing Dependency: python-repoze-lru for package: python-routes-2.4.1-1.el7.noarch
---> Package python-webob.noarch 0:1.7.2-1.el7 will be installed
---> Package python2-babel.noarch 0:2.3.4-1.el7 will be installed
--> Processing Dependency: pytz for package: python2-babel-2.3.4-1.el7.noarch
---> Package python2-bcrypt.x86_64 0:3.1.2-3.el7 will be installed
--> Processing Dependency: python-cffi for package: python2-bcrypt-3.1.2-3.el7.x86_64
---> Package python2-cryptography.x86_64 0:1.7.2-1.el7_4.1 will be installed
--> Processing Dependency: python-pyasn1 >= 0.1.8 for package: python2-cryptography-1.7.2-1.el7_4.1.x86_64
--> Processing Dependency: python-idna >= 2.0 for package: python2-cryptography-1.7.2-1.el7_4.1.x86_64
--> Processing Dependency: python-ipaddress for package: python2-cryptography-1.7.2-1.el7_4.1.x86_64
--> Processing Dependency: python-enum34 for package: python2-cryptography-1.7.2-1.el7_4.1.x86_64
---> Package python2-debtcollector.noarch 0:1.17.1-1.el7 will be installed
--> Processing Dependency: python-wrapt for package: python2-debtcollector-1.17.1-1.el7.noarch
--> Processing Dependency: python-funcsigs for package: python2-debtcollector-1.17.1-1.el7.noarch
---> Package python2-keystoneauth1.noarch 0:3.1.0-1.el7 will be installed
--> Processing Dependency: python-iso8601 >= 0.1.11 for package: python2-keystoneauth1-3.1.0-1.el7.noarch
---> Package python2-keystonemiddleware.noarch 0:4.17.0-1.el7 will be installed
---> Package python2-oauthlib.noarch 0:2.0.1-8.el7 will be installed
--> Processing Dependency: python-jwcrypto for package: python2-oauthlib-2.0.1-8.el7.noarch
---> Package python2-oslo-cache.noarch 0:1.25.0-1.el7 will be installed
--> Processing Dependency: python-oslo-cache-lang = 1.25.0-1.el7 for package: python2-oslo-cache-1.25.0-1.el7.noarch
---> Package python2-oslo-concurrency.noarch 0:3.21.1-1.el7 will be installed
--> Processing Dependency: python-oslo-concurrency-lang = 3.21.1-1.el7 for package: python2-oslo-concurrency-3.21.1-1.el7.noarch
--> Processing Dependency: python-fixtures for package: python2-oslo-concurrency-3.21.1-1.el7.noarch
--> Processing Dependency: python-fasteners for package: python2-oslo-concurrency-3.21.1-1.el7.noarch
---> Package python2-oslo-config.noarch 2:4.11.1-1.el7 will be installed
--> Processing Dependency: python-rfc3986 >= 0.3.1 for package: 2:python2-oslo-config-4.11.1-1.el7.noarch
--> Processing Dependency: python-netaddr >= 0.7.13 for package: 2:python2-oslo-config-4.11.1-1.el7.noarch
--> Processing Dependency: PyYAML >= 3.10 for package: 2:python2-oslo-config-4.11.1-1.el7.noarch
---> Package python2-oslo-context.noarch 0:2.17.0-1.el7 will be installed
---> Package python2-oslo-db.noarch 0:4.25.0-1.el7 will be installed
--> Processing Dependency: python-oslo-db-lang = 4.25.0-1.el7 for package: python2-oslo-db-4.25.0-1.el7.noarch
--> Processing Dependency: python-alembic >= 0.8.7 for package: python2-oslo-db-4.25.0-1.el7.noarch
--> Processing Dependency: MySQL-python for package: python2-oslo-db-4.25.0-1.el7.noarch
---> Package python2-oslo-i18n.noarch 0:3.17.0-1.el7 will be installed
--> Processing Dependency: python-oslo-i18n-lang = 3.17.0-1.el7 for package: python2-oslo-i18n-3.17.0-1.el7.noarch
---> Package python2-oslo-log.noarch 0:3.30.0-1.el7 will be installed
--> Processing Dependency: python-oslo-log-lang = 3.30.0-1.el7 for package: python2-oslo-log-3.30.0-1.el7.noarch
--> Processing Dependency: python-monotonic for package: python2-oslo-log-3.30.0-1.el7.noarch
--> Processing Dependency: python-inotify for package: python2-oslo-log-3.30.0-1.el7.noarch
--> Processing Dependency: python-dateutil for package: python2-oslo-log-3.30.0-1.el7.noarch
---> Package python2-oslo-messaging.noarch 0:5.30.1-1.el7 will be installed
--> Processing Dependency: python-pika >= 0.10.0 for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-oslo-service >= 1.10.0 for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-kombu >= 1:4.0.0 for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-futurist >= 0.11.0 for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-futures >= 3.0 for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-amqp >= 2.1.0 for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-tenacity for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-pyngus for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-pika_pool for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-eventlet for package: python2-oslo-messaging-5.30.1-1.el7.noarch
--> Processing Dependency: python-cachetools for package: python2-oslo-messaging-5.30.1-1.el7.noarch
---> Package python2-oslo-middleware.noarch 0:3.30.1-1.el7 will be installed
--> Processing Dependency: python-oslo-middleware-lang = 3.30.1-1.el7 for package: python2-oslo-middleware-3.30.1-1.el7.noarch
--> Processing Dependency: python-statsd for package: python2-oslo-middleware-3.30.1-1.el7.noarch
--> Processing Dependency: python-jinja2 for package: python2-oslo-middleware-3.30.1-1.el7.noarch
---> Package python2-oslo-policy.noarch 0:1.25.1-1.el7 will be installed
--> Processing Dependency: python-oslo-policy-lang = 1.25.1-1.el7 for package: python2-oslo-policy-1.25.1-1.el7.noarch
---> Package python2-oslo-serialization.noarch 0:2.20.0-1.el7 will be installed
---> Package python2-oslo-utils.noarch 0:3.28.0-1.el7 will be installed
--> Processing Dependency: python-oslo-utils-lang = 3.28.0-1.el7 for package: python2-oslo-utils-3.28.0-1.el7.noarch
--> Processing Dependency: python-netifaces >= 0.10.4 for package: python2-oslo-utils-3.28.0-1.el7.noarch
--> Processing Dependency: pyparsing for package: python2-oslo-utils-3.28.0-1.el7.noarch
---> Package python2-osprofiler.noarch 0:1.11.0-1.el7 will be installed
---> Package python2-passlib.noarch 0:1.7.0-4.el7 will be installed
---> Package python2-pbr.noarch 0:3.1.1-1.el7 will be installed
---> Package python2-positional.noarch 0:1.1.1-2.el7 will be installed
---> Package python2-pycadf.noarch 0:2.6.0-1.el7 will be installed
--> Processing Dependency: python-pycadf-common = 2.6.0-1.el7 for package: python2-pycadf-2.6.0-1.el7.noarch
---> Package python2-pysaml2.noarch 0:3.0.2-2.el7 will be installed
--> Processing Dependency: pycrypto >= 2.5 for package: python2-pysaml2-3.0.2-2.el7.noarch
--> Processing Dependency: python-zope-interface for package: python2-pysaml2-3.0.2-2.el7.noarch
--> Processing Dependency: python-repoze-who for package: python2-pysaml2-3.0.2-2.el7.noarch
---> Package python2-requests.noarch 0:2.11.1-1.el7 will be installed
--> Processing Dependency: python2-urllib3 = 1.16 for package: python2-requests-2.11.1-1.el7.noarch
---> Package python2-scrypt.x86_64 0:0.8.0-2.el7 will be installed
---> Package python2-six.noarch 0:1.10.0-9.el7 will be installed
---> Package python2-sqlalchemy.x86_64 0:1.1.11-1.el7 will be installed
---> Package python2-stevedore.noarch 0:1.25.0-1.el7 will be installed
--> Running transaction check
---> Package MySQL-python.x86_64 0:1.2.5-1.el7 will be installed
---> Package PyYAML.x86_64 0:3.10-11.el7 will be installed
--> Processing Dependency: libyaml-0.so.2()(64bit) for package: PyYAML-3.10-11.el7.x86_64
---> Package pyparsing.noarch 0:2.1.10-3.el7 will be installed
--> Processing Dependency: python-pyparsing = 2.1.10-3.el7 for package: pyparsing-2.1.10-3.el7.noarch
---> Package python-alembic.noarch 0:0.8.10-1.el7 will be installed
--> Processing Dependency: python-mako for package: python-alembic-0.8.10-1.el7.noarch
--> Processing Dependency: python-editor for package: python-alembic-0.8.10-1.el7.noarch
---> Package python-cachetools.noarch 0:1.1.6-2.el7 will be installed
---> Package python-dateutil.noarch 1:2.4.2-1.el7 will be installed
---> Package python-dogpile-core.noarch 0:0.4.1-2.el7 will be installed
---> Package python-enum34.noarch 0:1.0.4-1.el7 will be installed
---> Package python-fixtures.noarch 0:3.0.0-2.el7 will be installed
--> Processing Dependency: python-testtools >= 0.9.22 for package: python-fixtures-3.0.0-2.el7.noarch
---> Package python-futures.noarch 0:3.0.3-1.el7 will be installed
---> Package python-inotify.noarch 0:0.9.4-4.el7 will be installed
---> Package python-ipaddress.noarch 0:1.0.16-3.el7 will be installed
---> Package python-jwcrypto.noarch 0:0.2.1-1.el7 will be installed
---> Package python-monotonic.noarch 0:0.6-1.el7 will be installed
---> Package python-netaddr.noarch 0:0.7.18-1.el7 will be installed
---> Package python-netifaces.x86_64 0:0.10.4-3.el7 will be installed
---> Package python-oslo-cache-lang.noarch 0:1.25.0-1.el7 will be installed
---> Package python-oslo-concurrency-lang.noarch 0:3.21.1-1.el7 will be installed
---> Package python-oslo-db-lang.noarch 0:4.25.0-1.el7 will be installed
---> Package python-oslo-i18n-lang.noarch 0:3.17.0-1.el7 will be installed
---> Package python-oslo-log-lang.noarch 0:3.30.0-1.el7 will be installed
---> Package python-oslo-middleware-lang.noarch 0:3.30.1-1.el7 will be installed
---> Package python-oslo-policy-lang.noarch 0:1.25.1-1.el7 will be installed
---> Package python-oslo-utils-lang.noarch 0:3.28.0-1.el7 will be installed
---> Package python-pycadf-common.noarch 0:2.6.0-1.el7 will be installed
---> Package python-pyngus.noarch 0:2.0.3-3.el7 will be installed
--> Processing Dependency: qpid-proton-c >= 0.13.0 for package: python-pyngus-2.0.3-3.el7.noarch
--> Processing Dependency: python-qpid-proton >= 0.13.0 for package: python-pyngus-2.0.3-3.el7.noarch
---> Package python-repoze-lru.noarch 0:0.4-3.el7 will be installed
---> Package python-repoze-who.noarch 0:2.1-1.el7 will be installed
---> Package python-sqlparse.noarch 0:0.1.18-5.el7 will be installed
---> Package python-tempita.noarch 0:0.5.1-8.el7 will be installed
---> Package python-wrapt.x86_64 0:1.10.8-2.el7 will be installed
---> Package python-zope-interface.x86_64 0:4.0.5-4.el7 will be installed
---> Package python2-amqp.noarch 0:2.1.4-1.el7 will be installed
--> Processing Dependency: python2-vine for package: python2-amqp-2.1.4-1.el7.noarch
---> Package python2-cffi.x86_64 0:1.5.2-1.el7 will be installed
--> Processing Dependency: python-pycparser for package: python2-cffi-1.5.2-1.el7.x86_64
---> Package python2-crypto.x86_64 0:2.6.1-15.el7 will be installed
--> Processing Dependency: libtomcrypt.so.0()(64bit) for package: python2-crypto-2.6.1-15.el7.x86_64
---> Package python2-eventlet.noarch 0:0.20.1-2.el7 will be installed
--> Processing Dependency: python-greenlet for package: python2-eventlet-0.20.1-2.el7.noarch
---> Package python2-fasteners.noarch 0:0.14.1-6.el7 will be installed
---> Package python2-funcsigs.noarch 0:1.0.2-1.el7 will be installed
---> Package python2-futurist.noarch 0:1.3.0-1.el7 will be installed
--> Processing Dependency: python-contextlib2 >= 0.4.0 for package: python2-futurist-1.3.0-1.el7.noarch
--> Processing Dependency: python-prettytable for package: python2-futurist-1.3.0-1.el7.noarch
---> Package python2-idna.noarch 0:2.5-1.el7 will be installed
---> Package python2-iso8601.noarch 0:0.1.11-1.el7 will be installed
---> Package python2-jinja2.noarch 0:2.8.1-1.el7 will be installed
--> Processing Dependency: python-markupsafe for package: python2-jinja2-2.8.1-1.el7.noarch
---> Package python2-kombu.noarch 1:4.0.2-5.el7 will be installed
--> Processing Dependency: python-anyjson >= 0.3.3 for package: 1:python2-kombu-4.0.2-5.el7.noarch
---> Package python2-oslo-service.noarch 0:1.25.0-1.el7 will be installed
---> Package python2-pika.noarch 0:0.10.0-3.el7 will be installed
---> Package python2-pika_pool.noarch 0:0.1.3-3.el7 will be installed
---> Package python2-pyOpenSSL.noarch 0:16.2.0-3.el7 will be installed
---> Package python2-pyasn1.noarch 0:0.1.9-7.el7 will be installed
---> Package python2-rfc3986.noarch 0:0.3.1-1.el7 will be installed
---> Package python2-setuptools.noarch 0:22.0.5-1.el7 will be installed
---> Package python2-statsd.noarch 0:3.2.1-5.el7 will be installed
---> Package python2-tenacity.noarch 0:4.4.0-1.el7 will be installed
---> Package python2-urllib3.noarch 0:1.16-1.el7 will be installed
--> Processing Dependency: python-pysocks for package: python2-urllib3-1.16-1.el7.noarch
--> Processing Dependency: python-backports-ssl_match_hostname for package: python2-urllib3-1.16-1.el7.noarch
---> Package pytz.noarch 0:2016.10-2.el7 will be installed
--> Running transaction check
---> Package libtomcrypt.x86_64 0:1.17-33.20170623gitcd6e602.el7 will be installed
--> Processing Dependency: libtommath.so.1()(64bit) for package: libtomcrypt-1.17-33.20170623gitcd6e602.el7.x86_64
---> Package libyaml.x86_64 0:0.1.4-11.el7_0 will be installed
---> Package python-anyjson.noarch 0:0.3.3-3.el7 will be installed
---> Package python-backports-ssl_match_hostname.noarch 0:3.4.0.2-4.el7 will be installed
--> Processing Dependency: python-backports for package: python-backports-ssl_match_hostname-3.4.0.2-4.el7.noarch
---> Package python-contextlib2.noarch 0:0.4.0-1.el7 will be installed
---> Package python-editor.noarch 0:0.4-4.el7 will be installed
---> Package python-mako.noarch 0:0.8.1-2.el7 will be installed
--> Processing Dependency: python-beaker for package: python-mako-0.8.1-2.el7.noarch
---> Package python-prettytable.noarch 0:0.7.2-3.el7 will be installed
---> Package python-pycparser.noarch 0:2.14-1.el7 will be installed
--> Processing Dependency: python-ply for package: python-pycparser-2.14-1.el7.noarch
---> Package python-qpid-proton.x86_64 0:0.14.0-2.el7 will be installed
---> Package python-testtools.noarch 0:1.8.0-2.el7 will be installed
--> Processing Dependency: python-unittest2 >= 0.8.0 for package: python-testtools-1.8.0-2.el7.noarch
--> Processing Dependency: python-mimeparse for package: python-testtools-1.8.0-2.el7.noarch
--> Processing Dependency: python-extras for package: python-testtools-1.8.0-2.el7.noarch
---> Package python2-greenlet.x86_64 0:0.4.9-1.el7 will be installed
---> Package python2-markupsafe.x86_64 0:0.23-16.el7 will be installed
---> Package python2-pyparsing.noarch 0:2.1.10-3.el7 will be installed
---> Package python2-pysocks.noarch 0:1.5.6-3.el7 will be installed
---> Package python2-vine.noarch 0:1.1.3-2.el7 will be installed
---> Package qpid-proton-c.x86_64 0:0.14.0-2.el7 will be installed
--> Running transaction check
---> Package libtommath.x86_64 0:1.0-8.el7 will be installed
---> Package python-backports.x86_64 0:1.0-8.el7 will be installed
---> Package python-beaker.noarch 0:1.5.4-10.el7 will be installed
---> Package python-extras.noarch 0:0.0.3-2.el7 will be installed
---> Package python-mimeparse.noarch 0:0.1.4-1.el7 will be installed
---> Package python-ply.noarch 0:3.4-11.el7 will be installed
---> Package python-unittest2.noarch 0:1.0.1-1.el7 will be installed
--> Processing Dependency: python-traceback2 for package: python-unittest2-1.0.1-1.el7.noarch
--> Running transaction check
---> Package python-traceback2.noarch 0:1.4.0-2.el7 will be installed
--> Processing Dependency: python-linecache2 for package: python-traceback2-1.4.0-2.el7.noarch
--> Running transaction check
---> Package python-linecache2.noarch 0:1.0.0-1.el7 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                                       Arch             Version                                     Repository                       Size
==================================================================================================================================================
Installing:
 httpd                                         x86_64           2.4.6-67.el7.centos.6                       updates                         2.7 M
 openstack-keystone                            noarch           1:12.0.0-1.el7                              centos-openstack-pike            84 k
Installing for dependencies:
 MySQL-python                                  x86_64           1.2.5-1.el7                                 base                             90 k
 PyYAML                                        x86_64           3.10-11.el7                                 base                            153 k
 apr                                           x86_64           1.4.8-3.el7                                 base                            103 k
 apr-util                                      x86_64           1.5.2-6.el7                                 base                             92 k
 httpd-tools                                   x86_64           2.4.6-67.el7.centos.6                       updates                          88 k
 libtomcrypt                                   x86_64           1.17-33.20170623gitcd6e602.el7              centos-openstack-pike           390 k
 libtommath                                    x86_64           1.0-8.el7                                   centos-openstack-pike            51 k
 libyaml                                       x86_64           0.1.4-11.el7_0                              base                             55 k
 mailcap                                       noarch           2.1.41-2.el7                                base                             31 k
 pyparsing                                     noarch           2.1.10-3.el7                                centos-openstack-pike           6.8 k
 python-alembic                                noarch           0.8.10-1.el7                                centos-openstack-pike           738 k
 python-anyjson                                noarch           0.3.3-3.el7                                 centos-openstack-pike            12 k
 python-backports                              x86_64           1.0-8.el7                                   base                            5.8 k
 python-backports-ssl_match_hostname           noarch           3.4.0.2-4.el7                               base                             12 k
 python-beaker                                 noarch           1.5.4-10.el7                                base                             80 k
 python-cachetools                             noarch           1.1.6-2.el7                                 centos-openstack-pike            26 k
 python-contextlib2                            noarch           0.4.0-1.el7                                 centos-openstack-pike            12 k
 python-dateutil                               noarch           1:2.4.2-1.el7                               centos-openstack-pike            83 k
 python-dogpile-cache                          noarch           0.6.2-1.el7                                 centos-openstack-pike            69 k
 python-dogpile-core                           noarch           0.4.1-2.el7                                 centos-openstack-pike            19 k
 python-editor                                 noarch           0.4-4.el7                                   centos-openstack-pike            12 k
 python-enum34                                 noarch           1.0.4-1.el7                                 base                             52 k
 python-extras                                 noarch           0.0.3-2.el7                                 centos-openstack-pike            13 k
 python-fixtures                               noarch           3.0.0-2.el7                                 centos-openstack-pike            88 k
 python-futures                                noarch           3.0.3-1.el7                                 centos-openstack-pike            26 k
 python-inotify                                noarch           0.9.4-4.el7                                 base                             49 k
 python-ipaddress                              noarch           1.0.16-3.el7                                centos-openstack-pike            34 k
 python-jsonschema                             noarch           2.3.0-1.el7                                 centos-openstack-pike            60 k
 python-jwcrypto                               noarch           0.2.1-1.el7                                 base                             41 k
 python-keyring                                noarch           5.7.1-1.el7                                 centos-openstack-pike           116 k
 python-keystone                               noarch           1:12.0.0-1.el7                              centos-openstack-pike           963 k
 python-ldap                                   x86_64           2.4.15-2.el7                                base                            159 k
 python-ldappool                               noarch           1.0-4.el7                                   centos-openstack-pike            18 k
 python-linecache2                             noarch           1.0.0-1.el7                                 centos-openstack-pike            11 k
 python-mako                                   noarch           0.8.1-2.el7                                 base                            307 k
 python-memcached                              noarch           1.58-1.el7                                  centos-openstack-pike            38 k
 python-migrate                                noarch           0.11.0-1.el7                                centos-openstack-pike           228 k
 python-mimeparse                              noarch           0.1.4-1.el7                                 centos-openstack-pike           8.6 k
 python-monotonic                              noarch           0.6-1.el7                                   centos-openstack-pike           8.9 k
 python-msgpack                                x86_64           0.4.6-3.el7                                 centos-openstack-pike            73 k
 python-netaddr                                noarch           0.7.18-1.el7                                centos-openstack-pike           1.3 M
 python-netifaces                              x86_64           0.10.4-3.el7                                base                             17 k
 python-oslo-cache-lang                        noarch           1.25.0-1.el7                                centos-openstack-pike            12 k
 python-oslo-concurrency-lang                  noarch           3.21.1-1.el7                                centos-openstack-pike           9.4 k
 python-oslo-db-lang                           noarch           4.25.0-1.el7                                centos-openstack-pike           8.8 k
 python-oslo-i18n-lang                         noarch           3.17.0-1.el7                                centos-openstack-pike           9.6 k
 python-oslo-log-lang                          noarch           3.30.0-1.el7                                centos-openstack-pike           8.7 k
 python-oslo-middleware-lang                   noarch           3.30.1-1.el7                                centos-openstack-pike           7.7 k
 python-oslo-policy-lang                       noarch           1.25.1-1.el7                                centos-openstack-pike            13 k
 python-oslo-utils-lang                        noarch           3.28.0-1.el7                                centos-openstack-pike           8.5 k
 python-paste                                  noarch           1.7.5.1-9.20111221hg1498.el7                base                            866 k
 python-paste-deploy                           noarch           1.5.2-6.el7                                 centos-openstack-pike            46 k
 python-ply                                    noarch           3.4-11.el7                                  base                            123 k
 python-prettytable                            noarch           0.7.2-3.el7                                 base                             37 k
 python-pycadf-common                          noarch           2.6.0-1.el7                                 centos-openstack-pike            10 k
 python-pycparser                              noarch           2.14-1.el7                                  base                            104 k
 python-pyngus                                 noarch           2.0.3-3.el7                                 centos-openstack-pike            41 k
 python-qpid-proton                            x86_64           0.14.0-2.el7                                centos-openstack-pike           212 k
 python-repoze-lru                             noarch           0.4-3.el7                                   centos-openstack-pike            12 k
 python-repoze-who                             noarch           2.1-1.el7                                   centos-openstack-pike           118 k
 python-routes                                 noarch           2.4.1-1.el7                                 centos-openstack-pike           191 k
 python-sqlparse                               noarch           0.1.18-5.el7                                centos-openstack-pike            74 k
 python-tempita                                noarch           0.5.1-8.el7                                 centos-openstack-pike            32 k
 python-testtools                              noarch           1.8.0-2.el7                                 centos-openstack-pike           301 k
 python-traceback2                             noarch           1.4.0-2.el7                                 centos-openstack-pike            18 k
 python-unittest2                              noarch           1.0.1-1.el7                                 centos-openstack-pike           171 k
 python-webob                                  noarch           1.7.2-1.el7                                 centos-openstack-pike           207 k
 python-wrapt                                  x86_64           1.10.8-2.el7                                centos-openstack-pike            46 k
 python-zope-interface                         x86_64           4.0.5-4.el7                                 base                            138 k
 python2-amqp                                  noarch           2.1.4-1.el7                                 centos-openstack-pike            79 k
 python2-babel                                 noarch           2.3.4-1.el7                                 centos-openstack-pike           4.8 M
 python2-bcrypt                                x86_64           3.1.2-3.el7                                 centos-openstack-pike            36 k
 python2-cffi                                  x86_64           1.5.2-1.el7                                 centos-openstack-pike           214 k
 python2-crypto                                x86_64           2.6.1-15.el7                                centos-openstack-pike           476 k
 python2-cryptography                          x86_64           1.7.2-1.el7_4.1                             updates                         502 k
 python2-debtcollector                         noarch           1.17.1-1.el7                                centos-openstack-pike            28 k
 python2-eventlet                              noarch           0.20.1-2.el7                                centos-openstack-pike           512 k
 python2-fasteners                             noarch           0.14.1-6.el7                                centos-openstack-pike            38 k
 python2-funcsigs                              noarch           1.0.2-1.el7                                 centos-openstack-pike            24 k
 python2-futurist                              noarch           1.3.0-1.el7                                 centos-openstack-pike            58 k
 python2-greenlet                              x86_64           0.4.9-1.el7                                 centos-openstack-pike            30 k
 python2-idna                                  noarch           2.5-1.el7                                   centos-openstack-pike            94 k
 python2-iso8601                               noarch           0.1.11-1.el7                                centos-openstack-pike            19 k
 python2-jinja2                                noarch           2.8.1-1.el7                                 centos-openstack-pike           463 k
 python2-keystoneauth1                         noarch           3.1.0-1.el7                                 centos-openstack-pike           369 k
 python2-keystoneclient                        noarch           1:3.13.0-1.el7                              centos-openstack-pike           233 k
 python2-keystonemiddleware                    noarch           4.17.0-1.el7                                centos-openstack-pike            96 k
 python2-kombu                                 noarch           1:4.0.2-5.el7                               centos-openstack-pike           341 k
 python2-markupsafe                            x86_64           0.23-16.el7                                 centos-openstack-pike            32 k
 python2-oauthlib                              noarch           2.0.1-8.el7                                 base                            146 k
 python2-oslo-cache                            noarch           1.25.0-1.el7                                centos-openstack-pike            45 k
 python2-oslo-concurrency                      noarch           3.21.1-1.el7                                centos-openstack-pike            35 k
 python2-oslo-config                           noarch           2:4.11.1-1.el7                              centos-openstack-pike           197 k
 python2-oslo-context                          noarch           2.17.0-1.el7                                centos-openstack-pike            20 k
 python2-oslo-db                               noarch           4.25.0-1.el7                                centos-openstack-pike           145 k
 python2-oslo-i18n                             noarch           3.17.0-1.el7                                centos-openstack-pike            52 k
 python2-oslo-log                              noarch           3.30.0-1.el7                                centos-openstack-pike            55 k
 python2-oslo-messaging                        noarch           5.30.1-1.el7                                centos-openstack-pike           356 k
 python2-oslo-middleware                       noarch           3.30.1-1.el7                                centos-openstack-pike            49 k
 python2-oslo-policy                           noarch           1.25.1-1.el7                                centos-openstack-pike            50 k
 python2-oslo-serialization                    noarch           2.20.0-1.el7                                centos-openstack-pike            27 k
 python2-oslo-service                          noarch           1.25.0-1.el7                                centos-openstack-pike            60 k
 python2-oslo-utils                            noarch           3.28.0-1.el7                                centos-openstack-pike            69 k
 python2-osprofiler                            noarch           1.11.0-1.el7                                centos-openstack-pike           114 k
 python2-passlib                               noarch           1.7.0-4.el7                                 centos-openstack-pike           733 k
 python2-pbr                                   noarch           3.1.1-1.el7                                 centos-openstack-pike           263 k
 python2-pika                                  noarch           0.10.0-3.el7                                centos-openstack-pike           195 k
 python2-pika_pool                             noarch           0.1.3-3.el7                                 centos-openstack-pike            12 k
 python2-positional                            noarch           1.1.1-2.el7                                 centos-openstack-pike            16 k
 python2-pyOpenSSL                             noarch           16.2.0-3.el7                                centos-openstack-pike            88 k
 python2-pyasn1                                noarch           0.1.9-7.el7                                 base                            100 k
 python2-pycadf                                noarch           2.6.0-1.el7                                 centos-openstack-pike            46 k
 python2-pyparsing                             noarch           2.1.10-3.el7                                centos-openstack-pike           135 k
 python2-pysaml2                               noarch           3.0.2-2.el7                                 centos-openstack-pike           524 k
 python2-pysocks                               noarch           1.5.6-3.el7                                 centos-openstack-pike            20 k
 python2-requests                              noarch           2.11.1-1.el7                                centos-openstack-pike           105 k
 python2-rfc3986                               noarch           0.3.1-1.el7                                 centos-openstack-pike            28 k
 python2-scrypt                                x86_64           0.8.0-2.el7                                 centos-openstack-pike            26 k
 python2-setuptools                            noarch           22.0.5-1.el7                                centos-openstack-pike           485 k
 python2-six                                   noarch           1.10.0-9.el7                                centos-openstack-pike            31 k
 python2-sqlalchemy                            x86_64           1.1.11-1.el7                                centos-openstack-pike           1.7 M
 python2-statsd                                noarch           3.2.1-5.el7                                 centos-openstack-pike            28 k
 python2-stevedore                             noarch           1.25.0-1.el7                                centos-openstack-pike            56 k
 python2-tenacity                              noarch           4.4.0-1.el7                                 centos-openstack-pike            39 k
 python2-urllib3                               noarch           1.16-1.el7                                  centos-openstack-pike           126 k
 python2-vine                                  noarch           1.1.3-2.el7                                 centos-openstack-pike            26 k
 pytz                                          noarch           2016.10-2.el7                               base                             46 k
 qpid-proton-c                                 x86_64           0.14.0-2.el7                                centos-openstack-pike           129 k

Transaction Summary
==================================================================================================================================================
Install  2 Packages (+128 Dependent packages)

Total download size: 26 M
Installed size: 114 M
Downloading packages:
(1/130): MySQL-python-1.2.5-1.el7.x86_64.rpm                                                                               |  90 kB  00:00:00     
(2/130): PyYAML-3.10-11.el7.x86_64.rpm                                                                                     | 153 kB  00:00:00     
(3/130): apr-1.4.8-3.el7.x86_64.rpm                                                                                        | 103 kB  00:00:00     
(4/130): apr-util-1.5.2-6.el7.x86_64.rpm                                                                                   |  92 kB  00:00:00     
(5/130): libyaml-0.1.4-11.el7_0.x86_64.rpm                                                                                 |  55 kB  00:00:00     
(6/130): httpd-tools-2.4.6-67.el7.centos.6.x86_64.rpm                                                                      |  88 kB  00:00:00     
(7/130): mailcap-2.1.41-2.el7.noarch.rpm                                                                                   |  31 kB  00:00:00     
(8/130): libtommath-1.0-8.el7.x86_64.rpm                                                                                   |  51 kB  00:00:00     
(9/130): httpd-2.4.6-67.el7.centos.6.x86_64.rpm                                                                            | 2.7 MB  00:00:01     
(10/130): openstack-keystone-12.0.0-1.el7.noarch.rpm                                                                       |  84 kB  00:00:00     
(11/130): pyparsing-2.1.10-3.el7.noarch.rpm                                                                                | 6.8 kB  00:00:00     
(12/130): libtomcrypt-1.17-33.20170623gitcd6e602.el7.x86_64.rpm                                                            | 390 kB  00:00:03     
(13/130): python-backports-ssl_match_hostname-3.4.0.2-4.el7.noarch.rpm                                                     |  12 kB  00:00:00     
(14/130): python-backports-1.0-8.el7.x86_64.rpm                                                                            | 5.8 kB  00:00:00     
(15/130): python-beaker-1.5.4-10.el7.noarch.rpm                                                                            |  80 kB  00:00:00     
(16/130): python-anyjson-0.3.3-3.el7.noarch.rpm                                                                            |  12 kB  00:00:00     
(17/130): python-cachetools-1.1.6-2.el7.noarch.rpm                                                                         |  26 kB  00:00:00     
(18/130): python-contextlib2-0.4.0-1.el7.noarch.rpm                                                                        |  12 kB  00:00:00     
(19/130): python-dateutil-2.4.2-1.el7.noarch.rpm                                                                           |  83 kB  00:00:00     
(20/130): python-dogpile-cache-0.6.2-1.el7.noarch.rpm                                                                      |  69 kB  00:00:00     
(21/130): python-dogpile-core-0.4.1-2.el7.noarch.rpm                                                                       |  19 kB  00:00:00     
(22/130): python-enum34-1.0.4-1.el7.noarch.rpm                                                                             |  52 kB  00:00:00     
(23/130): python-editor-0.4-4.el7.noarch.rpm                                                                               |  12 kB  00:00:00     
(24/130): python-extras-0.0.3-2.el7.noarch.rpm                                                                             |  13 kB  00:00:00     
(25/130): python-fixtures-3.0.0-2.el7.noarch.rpm                                                                           |  88 kB  00:00:00     
(26/130): python-inotify-0.9.4-4.el7.noarch.rpm                                                                            |  49 kB  00:00:00     
(27/130): python-futures-3.0.3-1.el7.noarch.rpm                                                                            |  26 kB  00:00:00     
(28/130): python-ipaddress-1.0.16-3.el7.noarch.rpm                                                                         |  34 kB  00:00:00     
(29/130): python-jwcrypto-0.2.1-1.el7.noarch.rpm                                                                           |  41 kB  00:00:00     
(30/130): python-jsonschema-2.3.0-1.el7.noarch.rpm                                                                         |  60 kB  00:00:00     
(31/130): python-alembic-0.8.10-1.el7.noarch.rpm                                                                           | 738 kB  00:00:06     
(32/130): python-ldap-2.4.15-2.el7.x86_64.rpm                                                                              | 159 kB  00:00:00     
(33/130): python-keyring-5.7.1-1.el7.noarch.rpm                                                                            | 116 kB  00:00:00     
(34/130): python-ldappool-1.0-4.el7.noarch.rpm                                                                             |  18 kB  00:00:00     
(35/130): python-linecache2-1.0.0-1.el7.noarch.rpm                                                                         |  11 kB  00:00:00     
(36/130): python-mako-0.8.1-2.el7.noarch.rpm                                                                               | 307 kB  00:00:00     
(37/130): python-memcached-1.58-1.el7.noarch.rpm                                                                           |  38 kB  00:00:00     
(38/130): python-migrate-0.11.0-1.el7.noarch.rpm                                                                           | 228 kB  00:00:01     
(39/130): python-mimeparse-0.1.4-1.el7.noarch.rpm                                                                          | 8.6 kB  00:00:00     
(40/130): python-monotonic-0.6-1.el7.noarch.rpm                                                                            | 8.9 kB  00:00:00     
(41/130): python-msgpack-0.4.6-3.el7.x86_64.rpm                                                                            |  73 kB  00:00:00     
(42/130): python-netifaces-0.10.4-3.el7.x86_64.rpm                                                                         |  17 kB  00:00:00     
(43/130): python-keystone-12.0.0-1.el7.noarch.rpm                                                                          | 963 kB  00:00:06     
(44/130): python-oslo-cache-lang-1.25.0-1.el7.noarch.rpm                                                                   |  12 kB  00:00:00     
(45/130): python-oslo-concurrency-lang-3.21.1-1.el7.noarch.rpm                                                             | 9.4 kB  00:00:00     
(46/130): python-oslo-db-lang-4.25.0-1.el7.noarch.rpm                                                                      | 8.8 kB  00:00:00     
(47/130): python-oslo-i18n-lang-3.17.0-1.el7.noarch.rpm                                                                    | 9.6 kB  00:00:00     
(48/130): python-oslo-log-lang-3.30.0-1.el7.noarch.rpm                                                                     | 8.7 kB  00:00:00     
(49/130): python-oslo-middleware-lang-3.30.1-1.el7.noarch.rpm                                                              | 7.7 kB  00:00:00     
(50/130): python-oslo-policy-lang-1.25.1-1.el7.noarch.rpm                                                                  |  13 kB  00:00:00     
(51/130): python-oslo-utils-lang-3.28.0-1.el7.noarch.rpm                                                                   | 8.5 kB  00:00:00     
(52/130): python-ply-3.4-11.el7.noarch.rpm                                                                                 | 123 kB  00:00:00     
(53/130): python-prettytable-0.7.2-3.el7.noarch.rpm                                                                        |  37 kB  00:00:00     
(54/130): python-paste-1.7.5.1-9.20111221hg1498.el7.noarch.rpm                                                             | 866 kB  00:00:00     
(55/130): python-paste-deploy-1.5.2-6.el7.noarch.rpm                                                                       |  46 kB  00:00:00     
(56/130): python-pycadf-common-2.6.0-1.el7.noarch.rpm                                                                      |  10 kB  00:00:00     
(57/130): python-pycparser-2.14-1.el7.noarch.rpm                                                                           | 104 kB  00:00:00     
(58/130): python-pyngus-2.0.3-3.el7.noarch.rpm                                                                             |  41 kB  00:00:00     
(59/130): python-qpid-proton-0.14.0-2.el7.x86_64.rpm                                                                       | 212 kB  00:00:01     
(60/130): python-repoze-lru-0.4-3.el7.noarch.rpm                                                                           |  12 kB  00:00:00     
(61/130): python-repoze-who-2.1-1.el7.noarch.rpm                                                                           | 118 kB  00:00:00     
(62/130): python-routes-2.4.1-1.el7.noarch.rpm                                                                             | 191 kB  00:00:01     
(63/130): python-sqlparse-0.1.18-5.el7.noarch.rpm                                                                          |  74 kB  00:00:01     
(64/130): python-tempita-0.5.1-8.el7.noarch.rpm                                                                            |  32 kB  00:00:00     
(65/130): python-netaddr-0.7.18-1.el7.noarch.rpm                                                                           | 1.3 MB  00:00:09     
(66/130): python-traceback2-1.4.0-2.el7.noarch.rpm                                                                         |  18 kB  00:00:00     
(67/130): python-unittest2-1.0.1-1.el7.noarch.rpm                                                                          | 171 kB  00:00:01     
(68/130): python-testtools-1.8.0-2.el7.noarch.rpm                                                                          | 301 kB  00:00:02     
(69/130): python-zope-interface-4.0.5-4.el7.x86_64.rpm                                                                     | 138 kB  00:00:00     
(70/130): python-wrapt-1.10.8-2.el7.x86_64.rpm                                                                             |  46 kB  00:00:00     
(71/130): python-webob-1.7.2-1.el7.noarch.rpm                                                                              | 207 kB  00:00:01     
(72/130): python2-amqp-2.1.4-1.el7.noarch.rpm                                                                              |  79 kB  00:00:00     
(73/130): python2-bcrypt-3.1.2-3.el7.x86_64.rpm                                                                            |  36 kB  00:00:00     
(74/130): python2-cffi-1.5.2-1.el7.x86_64.rpm                                                                              | 214 kB  00:00:01     
(75/130): python2-cryptography-1.7.2-1.el7_4.1.x86_64.rpm                                                                  | 502 kB  00:00:00     
(76/130): python2-crypto-2.6.1-15.el7.x86_64.rpm                                                                           | 476 kB  00:00:02     
(77/130): python2-debtcollector-1.17.1-1.el7.noarch.rpm                                                                    |  28 kB  00:00:00     
(78/130): python2-eventlet-0.20.1-2.el7.noarch.rpm                                                                         | 512 kB  00:00:03     
(79/130): python2-fasteners-0.14.1-6.el7.noarch.rpm                                                                        |  38 kB  00:00:00     
(80/130): python2-funcsigs-1.0.2-1.el7.noarch.rpm                                                                          |  24 kB  00:00:00     
(81/130): python2-futurist-1.3.0-1.el7.noarch.rpm                                                                          |  58 kB  00:00:00     
(82/130): python2-greenlet-0.4.9-1.el7.x86_64.rpm                                                                          |  30 kB  00:00:00     
(83/130): python2-idna-2.5-1.el7.noarch.rpm                                                                                |  94 kB  00:00:00     
(84/130): python2-iso8601-0.1.11-1.el7.noarch.rpm                                                                          |  19 kB  00:00:00     
(85/130): python2-jinja2-2.8.1-1.el7.noarch.rpm                                                                            | 463 kB  00:00:04     
(86/130): python2-keystoneauth1-3.1.0-1.el7.noarch.rpm                                                                     | 369 kB  00:00:03     
(87/130): python2-keystoneclient-3.13.0-1.el7.noarch.rpm                                                                   | 233 kB  00:00:04     
(88/130): python2-keystonemiddleware-4.17.0-1.el7.noarch.rpm                                                               |  96 kB  00:00:00     
(89/130): python2-kombu-4.0.2-5.el7.noarch.rpm                                                                             | 341 kB  00:00:02     
(90/130): python2-oauthlib-2.0.1-8.el7.noarch.rpm                                                                          | 146 kB  00:00:00     
(91/130): python2-markupsafe-0.23-16.el7.x86_64.rpm                                                                        |  32 kB  00:00:00     
(92/130): python2-oslo-cache-1.25.0-1.el7.noarch.rpm                                                                       |  45 kB  00:00:00     
(93/130): python2-oslo-concurrency-3.21.1-1.el7.noarch.rpm                                                                 |  35 kB  00:00:00     
(94/130): python2-babel-2.3.4-1.el7.noarch.rpm                                                                             | 4.8 MB  00:00:27     
(95/130): python2-oslo-context-2.17.0-1.el7.noarch.rpm                                                                     |  20 kB  00:00:00     
(96/130): python2-oslo-config-4.11.1-1.el7.noarch.rpm                                                                      | 197 kB  00:00:01     
(97/130): python2-oslo-i18n-3.17.0-1.el7.noarch.rpm                                                                        |  52 kB  00:00:00     
(98/130): python2-oslo-db-4.25.0-1.el7.noarch.rpm                                                                          | 145 kB  00:00:00     
(99/130): python2-oslo-log-3.30.0-1.el7.noarch.rpm                                                                         |  55 kB  00:00:00     
(100/130): python2-oslo-middleware-3.30.1-1.el7.noarch.rpm                                                                 |  49 kB  00:00:00     
(101/130): python2-oslo-policy-1.25.1-1.el7.noarch.rpm                                                                     |  50 kB  00:00:00     
(102/130): python2-oslo-serialization-2.20.0-1.el7.noarch.rpm                                                              |  27 kB  00:00:00     
(103/130): python2-oslo-service-1.25.0-1.el7.noarch.rpm                                                                    |  60 kB  00:00:00     
(104/130): python2-oslo-messaging-5.30.1-1.el7.noarch.rpm                                                                  | 356 kB  00:00:01     
(105/130): python2-oslo-utils-3.28.0-1.el7.noarch.rpm                                                                      |  69 kB  00:00:00     
(106/130): python2-osprofiler-1.11.0-1.el7.noarch.rpm                                                                      | 114 kB  00:00:00     
(107/130): python2-pbr-3.1.1-1.el7.noarch.rpm                                                                              | 263 kB  00:00:01     
(108/130): python2-pika-0.10.0-3.el7.noarch.rpm                                                                            | 195 kB  00:00:01     
(109/130): python2-pika_pool-0.1.3-3.el7.noarch.rpm                                                                        |  12 kB  00:00:00     
(110/130): python2-positional-1.1.1-2.el7.noarch.rpm                                                                       |  16 kB  00:00:00     
(111/130): python2-pyasn1-0.1.9-7.el7.noarch.rpm                                                                           | 100 kB  00:00:00     
(112/130): python2-passlib-1.7.0-4.el7.noarch.rpm                                                                          | 733 kB  00:00:04     
(113/130): python2-pycadf-2.6.0-1.el7.noarch.rpm                                                                           |  46 kB  00:00:00     
(114/130): python2-pyOpenSSL-16.2.0-3.el7.noarch.rpm                                                                       |  88 kB  00:00:01     
(115/130): python2-pyparsing-2.1.10-3.el7.noarch.rpm                                                                       | 135 kB  00:00:00     
(116/130): python2-pysocks-1.5.6-3.el7.noarch.rpm                                                                          |  20 kB  00:00:00     
(117/130): python2-requests-2.11.1-1.el7.noarch.rpm                                                                        | 105 kB  00:00:00     
(118/130): python2-rfc3986-0.3.1-1.el7.noarch.rpm                                                                          |  28 kB  00:00:00     
(119/130): python2-scrypt-0.8.0-2.el7.x86_64.rpm                                                                           |  26 kB  00:00:00     
(120/130): python2-setuptools-22.0.5-1.el7.noarch.rpm                                                                      | 485 kB  00:00:03     
(121/130): python2-pysaml2-3.0.2-2.el7.noarch.rpm                                                                          | 524 kB  00:00:05     
(122/130): python2-six-1.10.0-9.el7.noarch.rpm                                                                             |  31 kB  00:00:00     
(123/130): python2-statsd-3.2.1-5.el7.noarch.rpm                                                                           |  28 kB  00:00:00     
(124/130): python2-stevedore-1.25.0-1.el7.noarch.rpm                                                                       |  56 kB  00:00:00     
(125/130): python2-tenacity-4.4.0-1.el7.noarch.rpm                                                                         |  39 kB  00:00:00     
(126/130): python2-urllib3-1.16-1.el7.noarch.rpm                                                                           | 126 kB  00:00:00     
(127/130): pytz-2016.10-2.el7.noarch.rpm                                                                                   |  46 kB  00:00:00     
(128/130): python2-vine-1.1.3-2.el7.noarch.rpm                                                                             |  26 kB  00:00:00     
(129/130): qpid-proton-c-0.14.0-2.el7.x86_64.rpm                                                                           | 129 kB  00:00:01     
(130/130): python2-sqlalchemy-1.1.11-1.el7.x86_64.rpm                                                                      | 1.7 MB  00:00:09     
--------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                             351 kB/s |  26 MB  00:01:15     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : python2-six-1.10.0-9.el7.noarch                                                                                              1/130 
  Installing : python2-pbr-3.1.1-1.el7.noarch                                                                                               2/130 
  Installing : python2-setuptools-22.0.5-1.el7.noarch                                                                                       3/130 
  Installing : python2-stevedore-1.25.0-1.el7.noarch                                                                                        4/130 
  Installing : python-webob-1.7.2-1.el7.noarch                                                                                              5/130 
  Installing : python-monotonic-0.6-1.el7.noarch                                                                                            6/130 
  Installing : python2-iso8601-0.1.11-1.el7.noarch                                                                                          7/130 
  Installing : python-netaddr-0.7.18-1.el7.noarch                                                                                           8/130 
  Installing : python2-sqlalchemy-1.1.11-1.el7.x86_64                                                                                       9/130 
  Installing : pytz-2016.10-2.el7.noarch                                                                                                   10/130 
  Installing : python2-babel-2.3.4-1.el7.noarch                                                                                            11/130 
  Installing : apr-1.4.8-3.el7.x86_64                                                                                                      12/130 
  Installing : python-futures-3.0.3-1.el7.noarch                                                                                           13/130 
  Installing : apr-util-1.5.2-6.el7.x86_64                                                                                                 14/130 
  Installing : 1:python-dateutil-2.4.2-1.el7.noarch                                                                                        15/130 
  Installing : python-msgpack-0.4.6-3.el7.x86_64                                                                                           16/130 
  Installing : qpid-proton-c-0.14.0-2.el7.x86_64                                                                                           17/130 
  Installing : python2-funcsigs-1.0.2-1.el7.noarch                                                                                         18/130 
  Installing : python-wrapt-1.10.8-2.el7.x86_64                                                                                            19/130 
  Installing : python2-debtcollector-1.17.1-1.el7.noarch                                                                                   20/130 
  Installing : python2-positional-1.1.1-2.el7.noarch                                                                                       21/130 
  Installing : python2-oslo-context-2.17.0-1.el7.noarch                                                                                    22/130 
  Installing : python2-vine-1.1.3-2.el7.noarch                                                                                             23/130 
  Installing : python2-amqp-2.1.4-1.el7.noarch                                                                                             24/130 
  Installing : python2-markupsafe-0.23-16.el7.x86_64                                                                                       25/130 
  Installing : python-memcached-1.58-1.el7.noarch                                                                                          26/130 
  Installing : python2-pika-0.10.0-3.el7.noarch                                                                                            27/130 
  Installing : python-ipaddress-1.0.16-3.el7.noarch                                                                                        28/130 
  Installing : python-zope-interface-4.0.5-4.el7.x86_64                                                                                    29/130 
  Installing : python2-greenlet-0.4.9-1.el7.x86_64                                                                                         30/130 
  Installing : python2-eventlet-0.20.1-2.el7.noarch                                                                                        31/130 
  Installing : python-tempita-0.5.1-8.el7.noarch                                                                                           32/130 
  Installing : python-ldap-2.4.15-2.el7.x86_64                                                                                             33/130 
  Installing : python-enum34-1.0.4-1.el7.noarch                                                                                            34/130 
  Installing : python-ldappool-1.0-4.el7.noarch                                                                                            35/130 
  Installing : python-repoze-who-2.1-1.el7.noarch                                                                                          36/130 
  Installing : python2-pika_pool-0.1.3-3.el7.noarch                                                                                        37/130 
  Installing : python2-jinja2-2.8.1-1.el7.noarch                                                                                           38/130 
  Installing : python-qpid-proton-0.14.0-2.el7.x86_64                                                                                      39/130 
  Installing : python-pyngus-2.0.3-3.el7.noarch                                                                                            40/130 
  Installing : httpd-tools-2.4.6-67.el7.centos.6.x86_64                                                                                    41/130 
  Installing : python2-tenacity-4.4.0-1.el7.noarch                                                                                         42/130 
  Installing : python2-fasteners-0.14.1-6.el7.noarch                                                                                       43/130 
  Installing : python-oslo-utils-lang-3.28.0-1.el7.noarch                                                                                  44/130 
  Installing : python2-pyasn1-0.1.9-7.el7.noarch                                                                                           45/130 
  Installing : python2-idna-2.5-1.el7.noarch                                                                                               46/130 
  Installing : python2-pyparsing-2.1.10-3.el7.noarch                                                                                       47/130 
  Installing : pyparsing-2.1.10-3.el7.noarch                                                                                               48/130 
  Installing : python-oslo-log-lang-3.30.0-1.el7.noarch                                                                                    49/130 
  Installing : python-oslo-i18n-lang-3.17.0-1.el7.noarch                                                                                   50/130 
  Installing : python2-pysocks-1.5.6-3.el7.noarch                                                                                          51/130 
  Installing : python2-statsd-3.2.1-5.el7.noarch                                                                                           52/130 
  Installing : python-oslo-concurrency-lang-3.21.1-1.el7.noarch                                                                            53/130 
  Installing : python-oslo-db-lang-4.25.0-1.el7.noarch                                                                                     54/130 
  Installing : python-oslo-middleware-lang-3.30.1-1.el7.noarch                                                                             55/130 
  Installing : python-inotify-0.9.4-4.el7.noarch                                                                                           56/130 
  Installing : python2-scrypt-0.8.0-2.el7.x86_64                                                                                           57/130 
  Installing : python-netifaces-0.10.4-3.el7.x86_64                                                                                        58/130 
  Installing : python-editor-0.4-4.el7.noarch                                                                                              59/130 
  Installing : python2-rfc3986-0.3.1-1.el7.noarch                                                                                          60/130 
  Installing : python-pycadf-common-2.6.0-1.el7.noarch                                                                                     61/130 
  Installing : python-extras-0.0.3-2.el7.noarch                                                                                            62/130 
  Installing : python-keyring-5.7.1-1.el7.noarch                                                                                           63/130 
  Installing : python-sqlparse-0.1.18-5.el7.noarch                                                                                         64/130 
  Installing : python-migrate-0.11.0-1.el7.noarch                                                                                          65/130 
  Installing : mailcap-2.1.41-2.el7.noarch                                                                                                 66/130 
  Installing : python-oslo-cache-lang-1.25.0-1.el7.noarch                                                                                  67/130 
  Installing : python-prettytable-0.7.2-3.el7.noarch                                                                                       68/130 
  Installing : python-dogpile-core-0.4.1-2.el7.noarch                                                                                      69/130 
  Installing : python-dogpile-cache-0.6.2-1.el7.noarch                                                                                     70/130 
  Installing : python-contextlib2-0.4.0-1.el7.noarch                                                                                       71/130 
  Installing : python2-futurist-1.3.0-1.el7.noarch                                                                                         72/130 
  Installing : python2-passlib-1.7.0-4.el7.noarch                                                                                          73/130 
  Installing : python-ply-3.4-11.el7.noarch                                                                                                74/130 
  Installing : python-pycparser-2.14-1.el7.noarch                                                                                          75/130 
  Installing : python2-cffi-1.5.2-1.el7.x86_64                                                                                             76/130 
  Installing : python2-cryptography-1.7.2-1.el7_4.1.x86_64                                                                                 77/130 
  Installing : python2-pyOpenSSL-16.2.0-3.el7.noarch                                                                                       78/130 
  Installing : python-paste-1.7.5.1-9.20111221hg1498.el7.noarch                                                                            79/130 
  Installing : python-paste-deploy-1.5.2-6.el7.noarch                                                                                      80/130 
  Installing : python-beaker-1.5.4-10.el7.noarch                                                                                           81/130 
  Installing : python-mako-0.8.1-2.el7.noarch                                                                                              82/130 
  Installing : python-alembic-0.8.10-1.el7.noarch                                                                                          83/130 
  Installing : python-jwcrypto-0.2.1-1.el7.noarch                                                                                          84/130 
  Installing : python2-oauthlib-2.0.1-8.el7.noarch                                                                                         85/130 
  Installing : python2-bcrypt-3.1.2-3.el7.x86_64                                                                                           86/130 
  Installing : python-linecache2-1.0.0-1.el7.noarch                                                                                        87/130 
  Installing : python-traceback2-1.4.0-2.el7.noarch                                                                                        88/130 
  Installing : python-unittest2-1.0.1-1.el7.noarch                                                                                         89/130 
  Installing : python-cachetools-1.1.6-2.el7.noarch                                                                                        90/130 
  Installing : python-repoze-lru-0.4-3.el7.noarch                                                                                          91/130 
  Installing : python-routes-2.4.1-1.el7.noarch                                                                                            92/130 
  Installing : libtommath-1.0-8.el7.x86_64                                                                                                 93/130 
  Installing : libtomcrypt-1.17-33.20170623gitcd6e602.el7.x86_64                                                                           94/130 
  Installing : python2-crypto-2.6.1-15.el7.x86_64                                                                                          95/130 
  Installing : python-backports-1.0-8.el7.x86_64                                                                                           96/130 
  Installing : python-backports-ssl_match_hostname-3.4.0.2-4.el7.noarch                                                                    97/130 
  Installing : python2-urllib3-1.16-1.el7.noarch                                                                                           98/130 
  Installing : python2-requests-2.11.1-1.el7.noarch                                                                                        99/130 
  Installing : python2-keystoneauth1-3.1.0-1.el7.noarch                                                                                   100/130 
  Installing : python2-pysaml2-3.0.2-2.el7.noarch                                                                                         101/130 
  Installing : libyaml-0.1.4-11.el7_0.x86_64                                                                                              102/130 
  Installing : PyYAML-3.10-11.el7.x86_64                                                                                                  103/130 
  Installing : python-anyjson-0.3.3-3.el7.noarch                                                                                          104/130 
  Installing : 1:python2-kombu-4.0.2-5.el7.noarch                                                                                         105/130 
  Installing : python-oslo-policy-lang-1.25.1-1.el7.noarch                                                                                106/130 
  Installing : python-mimeparse-0.1.4-1.el7.noarch                                                                                        107/130 
  Installing : python-testtools-1.8.0-2.el7.noarch                                                                                        108/130 
  Installing : python-fixtures-3.0.0-2.el7.noarch                                                                                         109/130 
  Installing : python2-oslo-i18n-3.17.0-1.el7.noarch                                                                                      110/130 
  Installing : python2-oslo-utils-3.28.0-1.el7.noarch                                                                                     111/130 
  Installing : 2:python2-oslo-config-4.11.1-1.el7.noarch                                                                                  112/130 
  Installing : python2-oslo-serialization-2.20.0-1.el7.noarch                                                                             113/130 
  Installing : python2-oslo-log-3.30.0-1.el7.noarch                                                                                       114/130 
  Installing : python2-oslo-concurrency-3.21.1-1.el7.noarch                                                                               115/130 
  Installing : 1:python2-keystoneclient-3.13.0-1.el7.noarch                                                                               116/130 
  Installing : python2-oslo-middleware-3.30.1-1.el7.noarch                                                                                117/130 
  Installing : python2-oslo-service-1.25.0-1.el7.noarch                                                                                   118/130 
  Installing : python2-oslo-messaging-5.30.1-1.el7.noarch                                                                                 119/130 
  Installing : python2-pycadf-2.6.0-1.el7.noarch                                                                                          120/130 
  Installing : python2-keystonemiddleware-4.17.0-1.el7.noarch                                                                             121/130 
  Installing : python2-osprofiler-1.11.0-1.el7.noarch                                                                                     122/130 
  Installing : python2-oslo-cache-1.25.0-1.el7.noarch                                                                                     123/130 
  Installing : python2-oslo-policy-1.25.1-1.el7.noarch                                                                                    124/130 
  Installing : MySQL-python-1.2.5-1.el7.x86_64                                                                                            125/130 
  Installing : python2-oslo-db-4.25.0-1.el7.noarch                                                                                        126/130 
  Installing : python-jsonschema-2.3.0-1.el7.noarch                                                                                       127/130 
  Installing : 1:python-keystone-12.0.0-1.el7.noarch                                                                                      128/130 
  Installing : 1:openstack-keystone-12.0.0-1.el7.noarch                                                                                   129/130 
  Installing : httpd-2.4.6-67.el7.centos.6.x86_64                                                                                         130/130 
  Verifying  : python-jsonschema-2.3.0-1.el7.noarch                                                                                         1/130 
  Verifying  : MySQL-python-1.2.5-1.el7.x86_64                                                                                              2/130 
  Verifying  : python2-oslo-concurrency-3.21.1-1.el7.noarch                                                                                 3/130 
  Verifying  : python2-jinja2-2.8.1-1.el7.noarch                                                                                            4/130 
  Verifying  : python2-cryptography-1.7.2-1.el7_4.1.x86_64                                                                                  5/130 
  Verifying  : python2-stevedore-1.25.0-1.el7.noarch                                                                                        6/130 
  Verifying  : libtomcrypt-1.17-33.20170623gitcd6e602.el7.x86_64                                                                            7/130 
  Verifying  : python-mako-0.8.1-2.el7.noarch                                                                                               8/130 
  Verifying  : httpd-tools-2.4.6-67.el7.centos.6.x86_64                                                                                     9/130 
  Verifying  : python-mimeparse-0.1.4-1.el7.noarch                                                                                         10/130 
  Verifying  : python-paste-deploy-1.5.2-6.el7.noarch                                                                                      11/130 
  Verifying  : python-oslo-policy-lang-1.25.1-1.el7.noarch                                                                                 12/130 
  Verifying  : python-qpid-proton-0.14.0-2.el7.x86_64                                                                                      13/130 
  Verifying  : python-anyjson-0.3.3-3.el7.noarch                                                                                           14/130 
  Verifying  : python2-pbr-3.1.1-1.el7.noarch                                                                                              15/130 
  Verifying  : libyaml-0.1.4-11.el7_0.x86_64                                                                                               16/130 
  Verifying  : apr-util-1.5.2-6.el7.x86_64                                                                                                 17/130 
  Verifying  : python-backports-1.0-8.el7.x86_64                                                                                           18/130 
  Verifying  : libtommath-1.0-8.el7.x86_64                                                                                                 19/130 
  Verifying  : python-routes-2.4.1-1.el7.noarch                                                                                            20/130 
  Verifying  : python-repoze-lru-0.4-3.el7.noarch                                                                                          21/130 
  Verifying  : python-cachetools-1.1.6-2.el7.noarch                                                                                        22/130 
  Verifying  : python-pycparser-2.14-1.el7.noarch                                                                                          23/130 
  Verifying  : python2-oslo-log-3.30.0-1.el7.noarch                                                                                        24/130 
  Verifying  : python-beaker-1.5.4-10.el7.noarch                                                                                           25/130 
  Verifying  : python2-bcrypt-3.1.2-3.el7.x86_64                                                                                           26/130 
  Verifying  : python-linecache2-1.0.0-1.el7.noarch                                                                                        27/130 
  Verifying  : python-jwcrypto-0.2.1-1.el7.noarch                                                                                          28/130 
  Verifying  : python-repoze-who-2.1-1.el7.noarch                                                                                          29/130 
  Verifying  : PyYAML-3.10-11.el7.x86_64                                                                                                   30/130 
  Verifying  : python-enum34-1.0.4-1.el7.noarch                                                                                            31/130 
  Verifying  : 1:python2-kombu-4.0.2-5.el7.noarch                                                                                          32/130 
  Verifying  : python2-oslo-utils-3.28.0-1.el7.noarch                                                                                      33/130 
  Verifying  : python-unittest2-1.0.1-1.el7.noarch                                                                                         34/130 
  Verifying  : python2-positional-1.1.1-2.el7.noarch                                                                                       35/130 
  Verifying  : 1:openstack-keystone-12.0.0-1.el7.noarch                                                                                    36/130 
  Verifying  : 2:python2-oslo-config-4.11.1-1.el7.noarch                                                                                   37/130 
  Verifying  : python-ply-3.4-11.el7.noarch                                                                                                38/130 
  Verifying  : python2-debtcollector-1.17.1-1.el7.noarch                                                                                   39/130 
  Verifying  : python2-oslo-serialization-2.20.0-1.el7.noarch                                                                              40/130 
  Verifying  : pytz-2016.10-2.el7.noarch                                                                                                   41/130 
  Verifying  : python2-passlib-1.7.0-4.el7.noarch                                                                                          42/130 
  Verifying  : python-contextlib2-0.4.0-1.el7.noarch                                                                                       43/130 
  Verifying  : python2-pysaml2-3.0.2-2.el7.noarch                                                                                          44/130 
  Verifying  : python2-sqlalchemy-1.1.11-1.el7.x86_64                                                                                      45/130 
  Verifying  : python2-oslo-middleware-3.30.1-1.el7.noarch                                                                                 46/130 
  Verifying  : python-dogpile-core-0.4.1-2.el7.noarch                                                                                      47/130 
  Verifying  : python2-setuptools-22.0.5-1.el7.noarch                                                                                      48/130 
  Verifying  : python2-pycadf-2.6.0-1.el7.noarch                                                                                           49/130 
  Verifying  : python2-futurist-1.3.0-1.el7.noarch                                                                                         50/130 
  Verifying  : python-ldappool-1.0-4.el7.noarch                                                                                            51/130 
  Verifying  : python-paste-1.7.5.1-9.20111221hg1498.el7.noarch                                                                            52/130 
  Verifying  : python-ldap-2.4.15-2.el7.x86_64                                                                                             53/130 
  Verifying  : python-futures-3.0.3-1.el7.noarch                                                                                           54/130 
  Verifying  : python-monotonic-0.6-1.el7.noarch                                                                                           55/130 
  Verifying  : python-prettytable-0.7.2-3.el7.noarch                                                                                       56/130 
  Verifying  : python-tempita-0.5.1-8.el7.noarch                                                                                           57/130 
  Verifying  : python-oslo-cache-lang-1.25.0-1.el7.noarch                                                                                  58/130 
  Verifying  : python-fixtures-3.0.0-2.el7.noarch                                                                                          59/130 
  Verifying  : python-backports-ssl_match_hostname-3.4.0.2-4.el7.noarch                                                                    60/130 
  Verifying  : python2-oslo-policy-1.25.1-1.el7.noarch                                                                                     61/130 
  Verifying  : mailcap-2.1.41-2.el7.noarch                                                                                                 62/130 
  Verifying  : 1:python2-keystoneclient-3.13.0-1.el7.noarch                                                                                63/130 
  Verifying  : python2-oauthlib-2.0.1-8.el7.noarch                                                                                         64/130 
  Verifying  : python2-greenlet-0.4.9-1.el7.x86_64                                                                                         65/130 
  Verifying  : pyparsing-2.1.10-3.el7.noarch                                                                                               66/130 
  Verifying  : python-sqlparse-0.1.18-5.el7.noarch                                                                                         67/130 
  Verifying  : python2-oslo-service-1.25.0-1.el7.noarch                                                                                    68/130 
  Verifying  : python-keyring-5.7.1-1.el7.noarch                                                                                           69/130 
  Verifying  : python-webob-1.7.2-1.el7.noarch                                                                                             70/130 
  Verifying  : python2-keystoneauth1-3.1.0-1.el7.noarch                                                                                    71/130 
  Verifying  : python2-requests-2.11.1-1.el7.noarch                                                                                        72/130 
  Verifying  : python-zope-interface-4.0.5-4.el7.x86_64                                                                                    73/130 
  Verifying  : python-extras-0.0.3-2.el7.noarch                                                                                            74/130 
  Verifying  : python-pycadf-common-2.6.0-1.el7.noarch                                                                                     75/130 
  Verifying  : python-ipaddress-1.0.16-3.el7.noarch                                                                                        76/130 
  Verifying  : python2-rfc3986-0.3.1-1.el7.noarch                                                                                          77/130 
  Verifying  : python2-six-1.10.0-9.el7.noarch                                                                                             78/130 
  Verifying  : python-editor-0.4-4.el7.noarch                                                                                              79/130 
  Verifying  : python-dogpile-cache-0.6.2-1.el7.noarch                                                                                     80/130 
  Verifying  : python-traceback2-1.4.0-2.el7.noarch                                                                                        81/130 
  Verifying  : python-netifaces-0.10.4-3.el7.x86_64                                                                                        82/130 
  Verifying  : python2-scrypt-0.8.0-2.el7.x86_64                                                                                           83/130 
  Verifying  : python2-pika_pool-0.1.3-3.el7.noarch                                                                                        84/130 
  Verifying  : python-migrate-0.11.0-1.el7.noarch                                                                                          85/130 
  Verifying  : python2-pika-0.10.0-3.el7.noarch                                                                                            86/130 
  Verifying  : python2-oslo-context-2.17.0-1.el7.noarch                                                                                    87/130 
  Verifying  : python2-oslo-cache-1.25.0-1.el7.noarch                                                                                      88/130 
  Verifying  : httpd-2.4.6-67.el7.centos.6.x86_64                                                                                          89/130 
  Verifying  : python-inotify-0.9.4-4.el7.noarch                                                                                           90/130 
  Verifying  : python-oslo-middleware-lang-3.30.1-1.el7.noarch                                                                             91/130 
  Verifying  : python-memcached-1.58-1.el7.noarch                                                                                          92/130 
  Verifying  : python2-oslo-db-4.25.0-1.el7.noarch                                                                                         93/130 
  Verifying  : python2-eventlet-0.20.1-2.el7.noarch                                                                                        94/130 
  Verifying  : python-oslo-db-lang-4.25.0-1.el7.noarch                                                                                     95/130 
  Verifying  : python-oslo-concurrency-lang-3.21.1-1.el7.noarch                                                                            96/130 
  Verifying  : python-testtools-1.8.0-2.el7.noarch                                                                                         97/130 
  Verifying  : python2-tenacity-4.4.0-1.el7.noarch                                                                                         98/130 
  Verifying  : python2-statsd-3.2.1-5.el7.noarch                                                                                           99/130 
  Verifying  : python2-markupsafe-0.23-16.el7.x86_64                                                                                      100/130 
  Verifying  : python2-oslo-i18n-3.17.0-1.el7.noarch                                                                                      101/130 
  Verifying  : 1:python-keystone-12.0.0-1.el7.noarch                                                                                      102/130 
  Verifying  : apr-1.4.8-3.el7.x86_64                                                                                                     103/130 
  Verifying  : python-pyngus-2.0.3-3.el7.noarch                                                                                           104/130 
  Verifying  : python-netaddr-0.7.18-1.el7.noarch                                                                                         105/130 
  Verifying  : python2-amqp-2.1.4-1.el7.noarch                                                                                            106/130 
  Verifying  : python2-pysocks-1.5.6-3.el7.noarch                                                                                         107/130 
  Verifying  : python2-iso8601-0.1.11-1.el7.noarch                                                                                        108/130 
  Verifying  : python2-vine-1.1.3-2.el7.noarch                                                                                            109/130 
  Verifying  : python-oslo-i18n-lang-3.17.0-1.el7.noarch                                                                                  110/130 
  Verifying  : python-wrapt-1.10.8-2.el7.x86_64                                                                                           111/130 
  Verifying  : python2-keystonemiddleware-4.17.0-1.el7.noarch                                                                             112/130 
  Verifying  : python2-funcsigs-1.0.2-1.el7.noarch                                                                                        113/130 
  Verifying  : python-oslo-log-lang-3.30.0-1.el7.noarch                                                                                   114/130 
  Verifying  : python2-osprofiler-1.11.0-1.el7.noarch                                                                                     115/130 
  Verifying  : python2-pyparsing-2.1.10-3.el7.noarch                                                                                      116/130 
  Verifying  : python2-babel-2.3.4-1.el7.noarch                                                                                           117/130 
  Verifying  : 1:python-dateutil-2.4.2-1.el7.noarch                                                                                       118/130 
  Verifying  : python-alembic-0.8.10-1.el7.noarch                                                                                         119/130 
  Verifying  : python2-fasteners-0.14.1-6.el7.noarch                                                                                      120/130 
  Verifying  : python2-urllib3-1.16-1.el7.noarch                                                                                          121/130 
  Verifying  : qpid-proton-c-0.14.0-2.el7.x86_64                                                                                          122/130 
  Verifying  : python2-oslo-messaging-5.30.1-1.el7.noarch                                                                                 123/130 
  Verifying  : python-msgpack-0.4.6-3.el7.x86_64                                                                                          124/130 
  Verifying  : python2-crypto-2.6.1-15.el7.x86_64                                                                                         125/130 
  Verifying  : python2-cffi-1.5.2-1.el7.x86_64                                                                                            126/130 
  Verifying  : python2-idna-2.5-1.el7.noarch                                                                                              127/130 
  Verifying  : python2-pyasn1-0.1.9-7.el7.noarch                                                                                          128/130 
  Verifying  : python2-pyOpenSSL-16.2.0-3.el7.noarch                                                                                      129/130 
  Verifying  : python-oslo-utils-lang-3.28.0-1.el7.noarch                                                                                 130/130 

Installed:
  httpd.x86_64 0:2.4.6-67.el7.centos.6                                  openstack-keystone.noarch 1:12.0.0-1.el7                                 

Dependency Installed:
  MySQL-python.x86_64 0:1.2.5-1.el7                                   PyYAML.x86_64 0:3.10-11.el7                                                
  apr.x86_64 0:1.4.8-3.el7                                            apr-util.x86_64 0:1.5.2-6.el7                                              
  httpd-tools.x86_64 0:2.4.6-67.el7.centos.6                          libtomcrypt.x86_64 0:1.17-33.20170623gitcd6e602.el7                        
  libtommath.x86_64 0:1.0-8.el7                                       libyaml.x86_64 0:0.1.4-11.el7_0                                            
  mailcap.noarch 0:2.1.41-2.el7                                       pyparsing.noarch 0:2.1.10-3.el7                                            
  python-alembic.noarch 0:0.8.10-1.el7                                python-anyjson.noarch 0:0.3.3-3.el7                                        
  python-backports.x86_64 0:1.0-8.el7                                 python-backports-ssl_match_hostname.noarch 0:3.4.0.2-4.el7                 
  python-beaker.noarch 0:1.5.4-10.el7                                 python-cachetools.noarch 0:1.1.6-2.el7                                     
  python-contextlib2.noarch 0:0.4.0-1.el7                             python-dateutil.noarch 1:2.4.2-1.el7                                       
  python-dogpile-cache.noarch 0:0.6.2-1.el7                           python-dogpile-core.noarch 0:0.4.1-2.el7                                   
  python-editor.noarch 0:0.4-4.el7                                    python-enum34.noarch 0:1.0.4-1.el7                                         
  python-extras.noarch 0:0.0.3-2.el7                                  python-fixtures.noarch 0:3.0.0-2.el7                                       
  python-futures.noarch 0:3.0.3-1.el7                                 python-inotify.noarch 0:0.9.4-4.el7                                        
  python-ipaddress.noarch 0:1.0.16-3.el7                              python-jsonschema.noarch 0:2.3.0-1.el7                                     
  python-jwcrypto.noarch 0:0.2.1-1.el7                                python-keyring.noarch 0:5.7.1-1.el7                                        
  python-keystone.noarch 1:12.0.0-1.el7                               python-ldap.x86_64 0:2.4.15-2.el7                                          
  python-ldappool.noarch 0:1.0-4.el7                                  python-linecache2.noarch 0:1.0.0-1.el7                                     
  python-mako.noarch 0:0.8.1-2.el7                                    python-memcached.noarch 0:1.58-1.el7                                       
  python-migrate.noarch 0:0.11.0-1.el7                                python-mimeparse.noarch 0:0.1.4-1.el7                                      
  python-monotonic.noarch 0:0.6-1.el7                                 python-msgpack.x86_64 0:0.4.6-3.el7                                        
  python-netaddr.noarch 0:0.7.18-1.el7                                python-netifaces.x86_64 0:0.10.4-3.el7                                     
  python-oslo-cache-lang.noarch 0:1.25.0-1.el7                        python-oslo-concurrency-lang.noarch 0:3.21.1-1.el7                         
  python-oslo-db-lang.noarch 0:4.25.0-1.el7                           python-oslo-i18n-lang.noarch 0:3.17.0-1.el7                                
  python-oslo-log-lang.noarch 0:3.30.0-1.el7                          python-oslo-middleware-lang.noarch 0:3.30.1-1.el7                          
  python-oslo-policy-lang.noarch 0:1.25.1-1.el7                       python-oslo-utils-lang.noarch 0:3.28.0-1.el7                               
  python-paste.noarch 0:1.7.5.1-9.20111221hg1498.el7                  python-paste-deploy.noarch 0:1.5.2-6.el7                                   
  python-ply.noarch 0:3.4-11.el7                                      python-prettytable.noarch 0:0.7.2-3.el7                                    
  python-pycadf-common.noarch 0:2.6.0-1.el7                           python-pycparser.noarch 0:2.14-1.el7                                       
  python-pyngus.noarch 0:2.0.3-3.el7                                  python-qpid-proton.x86_64 0:0.14.0-2.el7                                   
  python-repoze-lru.noarch 0:0.4-3.el7                                python-repoze-who.noarch 0:2.1-1.el7                                       
  python-routes.noarch 0:2.4.1-1.el7                                  python-sqlparse.noarch 0:0.1.18-5.el7                                      
  python-tempita.noarch 0:0.5.1-8.el7                                 python-testtools.noarch 0:1.8.0-2.el7                                      
  python-traceback2.noarch 0:1.4.0-2.el7                              python-unittest2.noarch 0:1.0.1-1.el7                                      
  python-webob.noarch 0:1.7.2-1.el7                                   python-wrapt.x86_64 0:1.10.8-2.el7                                         
  python-zope-interface.x86_64 0:4.0.5-4.el7                          python2-amqp.noarch 0:2.1.4-1.el7                                          
  python2-babel.noarch 0:2.3.4-1.el7                                  python2-bcrypt.x86_64 0:3.1.2-3.el7                                        
  python2-cffi.x86_64 0:1.5.2-1.el7                                   python2-crypto.x86_64 0:2.6.1-15.el7                                       
  python2-cryptography.x86_64 0:1.7.2-1.el7_4.1                       python2-debtcollector.noarch 0:1.17.1-1.el7                                
  python2-eventlet.noarch 0:0.20.1-2.el7                              python2-fasteners.noarch 0:0.14.1-6.el7                                    
  python2-funcsigs.noarch 0:1.0.2-1.el7                               python2-futurist.noarch 0:1.3.0-1.el7                                      
  python2-greenlet.x86_64 0:0.4.9-1.el7                               python2-idna.noarch 0:2.5-1.el7                                            
  python2-iso8601.noarch 0:0.1.11-1.el7                               python2-jinja2.noarch 0:2.8.1-1.el7                                        
  python2-keystoneauth1.noarch 0:3.1.0-1.el7                          python2-keystoneclient.noarch 1:3.13.0-1.el7                               
  python2-keystonemiddleware.noarch 0:4.17.0-1.el7                    python2-kombu.noarch 1:4.0.2-5.el7                                         
  python2-markupsafe.x86_64 0:0.23-16.el7                             python2-oauthlib.noarch 0:2.0.1-8.el7                                      
  python2-oslo-cache.noarch 0:1.25.0-1.el7                            python2-oslo-concurrency.noarch 0:3.21.1-1.el7                             
  python2-oslo-config.noarch 2:4.11.1-1.el7                           python2-oslo-context.noarch 0:2.17.0-1.el7                                 
  python2-oslo-db.noarch 0:4.25.0-1.el7                               python2-oslo-i18n.noarch 0:3.17.0-1.el7                                    
  python2-oslo-log.noarch 0:3.30.0-1.el7                              python2-oslo-messaging.noarch 0:5.30.1-1.el7                               
  python2-oslo-middleware.noarch 0:3.30.1-1.el7                       python2-oslo-policy.noarch 0:1.25.1-1.el7                                  
  python2-oslo-serialization.noarch 0:2.20.0-1.el7                    python2-oslo-service.noarch 0:1.25.0-1.el7                                 
  python2-oslo-utils.noarch 0:3.28.0-1.el7                            python2-osprofiler.noarch 0:1.11.0-1.el7                                   
  python2-passlib.noarch 0:1.7.0-4.el7                                python2-pbr.noarch 0:3.1.1-1.el7                                           
  python2-pika.noarch 0:0.10.0-3.el7                                  python2-pika_pool.noarch 0:0.1.3-3.el7                                     
  python2-positional.noarch 0:1.1.1-2.el7                             python2-pyOpenSSL.noarch 0:16.2.0-3.el7                                    
  python2-pyasn1.noarch 0:0.1.9-7.el7                                 python2-pycadf.noarch 0:2.6.0-1.el7                                        
  python2-pyparsing.noarch 0:2.1.10-3.el7                             python2-pysaml2.noarch 0:3.0.2-2.el7                                       
  python2-pysocks.noarch 0:1.5.6-3.el7                                python2-requests.noarch 0:2.11.1-1.el7                                     
  python2-rfc3986.noarch 0:0.3.1-1.el7                                python2-scrypt.x86_64 0:0.8.0-2.el7                                        
  python2-setuptools.noarch 0:22.0.5-1.el7                            python2-six.noarch 0:1.10.0-9.el7                                          
  python2-sqlalchemy.x86_64 0:1.1.11-1.el7                            python2-statsd.noarch 0:3.2.1-5.el7                                        
  python2-stevedore.noarch 0:1.25.0-1.el7                             python2-tenacity.noarch 0:4.4.0-1.el7                                      
  python2-urllib3.noarch 0:1.16-1.el7                                 python2-vine.noarch 0:1.1.3-2.el7                                          
  pytz.noarch 0:2016.10-2.el7                                         qpid-proton-c.x86_64 0:0.14.0-2.el7                                        

Complete!
```

```
[vagrant@localhost ~]$ ls /usr/lib/python2.7/site-packages/keystone
assignment  common      endpoint_policy  federation  identity      middleware         notifications.pyo  revoke  v2_crud
auth        conf        exception.py     i18n.py     __init__.py   models             oauth1             server  version
catalog     contrib     exception.pyc    i18n.pyc    __init__.pyc  notifications.py   policy             token
cmd         credential  exception.pyo    i18n.pyo    __init__.pyo  notifications.pyc  resource           trust
```

```
[vagrant@localhost ~]$ ls /usr/share/keystone/
keystone-dist.conf  keystone-schema.json  keystone-schema.yaml  policy.v3cloudsample.json  sample_data.sh  wsgi-keystone.conf
```

```
[vagrant@localhost ~]$ sed '0,/chrony/d' /etc/passwd
sshd:x:74:74:Privilege-separated SSH:/var/empty/sshd:/sbin/nologin
vagrant:x:1000:1000:vagrant:/home/vagrant:/bin/bash
mysql:x:27:27:MySQL Server:/var/lib/mysql:/sbin/nologin
epmd:x:997:995:Erlang Port Mapper Daemon:/tmp:/sbin/nologin
rabbitmq:x:996:994:RabbitMQ messaging server:/var/lib/rabbitmq:/sbin/nologin
memcached:x:995:993:Memcached daemon:/run/memcached:/sbin/nologin
keystone:x:163:163:OpenStack Keystone Daemons:/var/lib/keystone:/sbin/nologin
apache:x:48:48:Apache:/usr/share/httpd:/sbin/nologin
```

Configure
```
[vagrant@localhost ~]$ sudo sed -i 's/^\[DEFAULT\]$/&\ndebug=true\nverbose=true\n/;s%^\[database\]$%&\nconnection=mysql+pymysql://keystone:KEYSTONE_DBPASS@10.64.33.64/keystone\n%;s/^\[token\]$/&\nprovider=fernet\n/' /etc/keystone/keystone.conf
```

```
[vagrant@localhost ~]$ sudo su -s /bin/sh -c "keystone-manage db_sync" keystone
```

```
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "show tables in keystone;"
+------------------------+
| Tables_in_keystone     |
+------------------------+
| access_token           |
| assignment             |
| config_register        |
| consumer               |
| credential             |
| endpoint               |
| endpoint_group         |
| federated_user         |
| federation_protocol    |
| group                  |
| id_mapping             |
| identity_provider      |
| idp_remote_ids         |
| implied_role           |
| local_user             |
| mapping                |
| migrate_version        |
| nonlocal_user          |
| password               |
| policy                 |
| policy_association     |
| project                |
| project_endpoint       |
| project_endpoint_group |
| region                 |
| request_token          |
| revocation_event       |
| role                   |
| sensitive_config       |
| service                |
| service_provider       |
| token                  |
| trust                  |
| trust_role             |
| user                   |
| user_group_membership  |
| user_option            |
| whitelisted_config     |
+------------------------+
```

Using _fernet_
```
[vagrant@localhost ~]$ sudo keystone-manage fernet_setup --keystone-user keystone --keystone-group keystone
```

```
[vagrant@localhost ~]$ sudo keystone-manage credential_setup --keystone-user keystone --keystone-group keystone
```

```
[vagrant@localhost ~]$ sudo keystone-manage bootstrap --bootstrap-password ADMIN_PASS --bootstrap-admin-url http://10.64.33.64:35357/v3/ --bootstrap-internal-url http://10.64.33.64:5000/v3/ --bootstrap-public-url http://10.64.33.64:5000/v3/ --bootstrap-region-id RegionOne
```

```
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.project;"
+----------------------------------+--------------------------+-------+-----------------------------------------------+---------+--------------------------+-----------+-----------+
| id                               | name                     | extra | description                                   | enabled | domain_id                | parent_id | is_domain |
+----------------------------------+--------------------------+-------+-----------------------------------------------+---------+--------------------------+-----------+-----------+
| <<keystone.domain.root>>         | <<keystone.domain.root>> | {}    |                                               |       0 | <<keystone.domain.root>> | NULL      |         1 |
| a0be38aef8c74d4abca3e4e100ee7910 | admin                    | {}    | Bootstrap project for initializing the cloud. |       1 | default                  | default   |         0 |
| default                          | Default                  | {}    | The default domain                            |       1 | <<keystone.domain.root>> | NULL      |         1 |
+----------------------------------+--------------------------+-------+-----------------------------------------------+---------+--------------------------+-----------+-----------+
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.user;"
+----------------------------------+-------+---------+--------------------+---------------------+----------------+-----------+
| id                               | extra | enabled | default_project_id | created_at          | last_active_at | domain_id |
+----------------------------------+-------+---------+--------------------+---------------------+----------------+-----------+
| 44e6ee1df8ae436986d2d50f7b358aa0 | {}    |       1 | NULL               | 2017-10-20 21:52:05 | NULL           | default   |
+----------------------------------+-------+---------+--------------------+---------------------+----------------+-----------+
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.local_user;"
+----+----------------------------------+-----------+-------+-------------------+----------------+
| id | user_id                          | domain_id | name  | failed_auth_count | failed_auth_at |
+----+----------------------------------+-----------+-------+-------------------+----------------+
|  1 | 44e6ee1df8ae436986d2d50f7b358aa0 | default   | admin |                 0 | NULL           |
+----+----------------------------------+-----------+-------+-------------------+----------------+
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.password;"
+----+---------------+----------+------------+--------------+--------------------------------------------------------------+------------------+----------------+---------------------+
| id | local_user_id | password | expires_at | self_service | password_hash                                                | created_at_int   | expires_at_int | created_at          |
+----+---------------+----------+------------+--------------+--------------------------------------------------------------+------------------+----------------+---------------------+
|  1 |             1 | NULL     | NULL       |            0 | $2b$12$vipVz/bPCPjkybYp1dSFgO3KMYkOQzLgFM32Uo6mTfjSeXt5X75rq | 1508536324819918 |           NULL | 2017-10-20 21:52:04 |
+----+---------------+----------+------------+--------------+--------------------------------------------------------------+------------------+----------------+---------------------+
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.role;"
+----------------------------------+----------+-------+-----------+
| id                               | name     | extra | domain_id |
+----------------------------------+----------+-------+-----------+
| 9ce0832112464e818686139e46cb1d78 | admin    | {}    | <<null>>  |
| 9fe2ff9ee4384b1894a90878d3e92bab | _member_ | {}    | <<null>>  |
+----------------------------------+----------+-------+-----------+
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.assignment;"
+-------------+----------------------------------+----------------------------------+----------------------------------+-----------+
| type        | actor_id                         | target_id                        | role_id                          | inherited |
+-------------+----------------------------------+----------------------------------+----------------------------------+-----------+
| UserProject | 44e6ee1df8ae436986d2d50f7b358aa0 | a0be38aef8c74d4abca3e4e100ee7910 | 9ce0832112464e818686139e46cb1d78 |         0 |
+-------------+----------------------------------+----------------------------------+----------------------------------+-----------+
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.region;"
+-----------+-------------+------------------+-------+
| id        | description | parent_region_id | extra |
+-----------+-------------+------------------+-------+
| RegionOne |             | NULL             | {}    |
+-----------+-------------+------------------+-------+
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.service;"
+----------------------------------+----------+---------+----------------------+
| id                               | type     | enabled | extra                |
+----------------------------------+----------+---------+----------------------+
| 896a7add738d40fb931cd704001d99e9 | identity |       1 | {"name": "keystone"} |
+----------------------------------+----------+---------+----------------------+
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.endpoint;"
+----------------------------------+--------------------+-----------+----------------------------------+------------------------------+-------+---------+-----------+
| id                               | legacy_endpoint_id | interface | service_id                       | url                          | extra | enabled | region_id |
+----------------------------------+--------------------+-----------+----------------------------------+------------------------------+-------+---------+-----------+
| 1c05dda3e9144422bd31c9c91ace6011 | NULL               | internal  | 896a7add738d40fb931cd704001d99e9 | http://10.64.33.64:5000/v3/  | {}    |       1 | RegionOne |
| 638c40e1f33f42d78bc4cead7e6e2438 | NULL               | admin     | 896a7add738d40fb931cd704001d99e9 | http://10.64.33.64:35357/v3/ | {}    |       1 | RegionOne |
| a11e9e70b7f74b43856d349e6ff8e9da | NULL               | public    | 896a7add738d40fb931cd704001d99e9 | http://10.64.33.64:5000/v3/  | {}    |       1 | RegionOne |
+----------------------------------+--------------------+-----------+----------------------------------+------------------------------+-------+---------+-----------+
[vagrant@localhost ~]$ mysql -u keystone --password=KEYSTONE_DBPASS -e "select * from keystone.migrate_version;"
+-----------------------+--------------------------------------------------------------------------+---------+
| repository_id         | repository_path                                                          | version |
+-----------------------+--------------------------------------------------------------------------+---------+
| keystone              | /usr/lib/python2.7/site-packages/keystone/common/sql/migrate_repo        |     109 |
| keystone_contract     | /usr/lib/python2.7/site-packages/keystone/common/sql/contract_repo       |      24 |
| keystone_data_migrate | /usr/lib/python2.7/site-packages/keystone/common/sql/data_migration_repo |      24 |
| keystone_expand       | /usr/lib/python2.7/site-packages/keystone/common/sql/expand_repo         |      24 |
+-----------------------+--------------------------------------------------------------------------+---------+
```

Apache HTTP server
```
[vagrant@localhost ~]$ sudo sed -i '1 i\ServerName 10.64.33.64\n' /etc/httpd/conf/httpd.conf
```

```
[vagrant@localhost ~]$ sudo ln -s /usr/share/keystone/wsgi-keystone.conf /etc/httpd/conf.d/
```

Trouble shooting
```
[vagrant@localhost ~]$ sudo systemctl start httpd.service
Job for httpd.service failed because the control process exited with error code. See "systemctl status httpd.service" and "journalctl -xe" for details.
```

```
[vagrant@localhost ~]$ sudo systemctl -l status httpd.service
● httpd.service - The Apache HTTP Server
   Loaded: loaded (/usr/lib/systemd/system/httpd.service; disabled; vendor preset: disabled)
   Active: failed (Result: exit-code) since Fri 2017-10-20 22:15:20 UTC; 8min ago
     Docs: man:httpd(8)
           man:apachectl(8)
  Process: 26805 ExecStop=/bin/kill -WINCH ${MAINPID} (code=exited, status=1/FAILURE)
  Process: 26804 ExecStart=/usr/sbin/httpd $OPTIONS -DFOREGROUND (code=exited, status=1/FAILURE)
 Main PID: 26804 (code=exited, status=1/FAILURE)

Oct 20 22:15:20 localhost.localdomain systemd[1]: Starting The Apache HTTP Server...
Oct 20 22:15:20 localhost.localdomain httpd[26804]: AH00526: Syntax error on line 5 of /etc/httpd/conf.d/wsgi-keystone.conf:
Oct 20 22:15:20 localhost.localdomain httpd[26804]: Invalid command 'WSGIDaemonProcess', perhaps misspelled or defined by a module not included in the server configuration
Oct 20 22:15:20 localhost.localdomain systemd[1]: httpd.service: main process exited, code=exited, status=1/FAILURE
Oct 20 22:15:20 localhost.localdomain kill[26805]: kill: cannot find process ""
Oct 20 22:15:20 localhost.localdomain systemd[1]: httpd.service: control process exited, code=exited status=1
Oct 20 22:15:20 localhost.localdomain systemd[1]: Failed to start The Apache HTTP Server.
Oct 20 22:15:20 localhost.localdomain systemd[1]: Unit httpd.service entered failed state.
Oct 20 22:15:20 localhost.localdomain systemd[1]: httpd.service failed.
```

```
[vagrant@localhost ~]$ sudo yum list | grep mod_wsgi
mod_wsgi.x86_64                    3.4-12.el7_0            base                 
```

```
[vagrant@localhost ~]$ sudo yum install -y mod_wsgi
Loaded plugins: fastestmirror
Loading mirror speeds from cached hostfile
 * base: mirrors.aliyun.com
 * extras: mirrors.aliyun.com
 * updates: mirrors.sohu.com
Resolving Dependencies
--> Running transaction check
---> Package mod_wsgi.x86_64 0:3.4-12.el7_0 will be installed
--> Finished Dependency Resolution

Dependencies Resolved

==================================================================================================================================================
 Package                            Arch                             Version                                 Repository                      Size
==================================================================================================================================================
Installing:
 mod_wsgi                           x86_64                           3.4-12.el7_0                            base                            76 k

Transaction Summary
==================================================================================================================================================
Install  1 Package

Total download size: 76 k
Installed size: 197 k
Downloading packages:
mod_wsgi-3.4-12.el7_0.x86_64.rpm                                                                                           |  76 kB  00:00:00     
Running transaction check
Running transaction test
Transaction test succeeded
Running transaction
  Installing : mod_wsgi-3.4-12.el7_0.x86_64                                                                                                   1/1 
  Verifying  : mod_wsgi-3.4-12.el7_0.x86_64                                                                                                   1/1 

Installed:
  mod_wsgi.x86_64 0:3.4-12.el7_0                                                                                                                  

Complete!
```

```
[vagrant@localhost ~]$ sudo httpd -M | grep wsgi
 wsgi_module (shared)
```

```
[vagrant@localhost ~]$ sudo systemctl -l status httpd.service
● httpd.service - The Apache HTTP Server
   Loaded: loaded (/usr/lib/systemd/system/httpd.service; disabled; vendor preset: disabled)
   Active: failed (Result: exit-code) since Fri 2017-10-20 22:28:09 UTC; 2min 55s ago
     Docs: man:httpd(8)
           man:apachectl(8)
  Process: 27042 ExecStop=/bin/kill -WINCH ${MAINPID} (code=exited, status=1/FAILURE)
  Process: 27041 ExecStart=/usr/sbin/httpd $OPTIONS -DFOREGROUND (code=exited, status=1/FAILURE)
 Main PID: 27041 (code=exited, status=1/FAILURE)

Oct 20 22:28:09 localhost.localdomain httpd[27041]: (13)Permission denied: AH00072: make_sock: could not bind to address [::]:5000
Oct 20 22:28:09 localhost.localdomain httpd[27041]: (13)Permission denied: AH00072: make_sock: could not bind to address 0.0.0.0:5000
Oct 20 22:28:09 localhost.localdomain httpd[27041]: no listening sockets available, shutting down
Oct 20 22:28:09 localhost.localdomain httpd[27041]: AH00015: Unable to open logs
Oct 20 22:28:09 localhost.localdomain systemd[1]: httpd.service: main process exited, code=exited, status=1/FAILURE
Oct 20 22:28:09 localhost.localdomain kill[27042]: kill: cannot find process ""
Oct 20 22:28:09 localhost.localdomain systemd[1]: httpd.service: control process exited, code=exited status=1
Oct 20 22:28:09 localhost.localdomain systemd[1]: Failed to start The Apache HTTP Server.
Oct 20 22:28:09 localhost.localdomain systemd[1]: Unit httpd.service entered failed state.
Oct 20 22:28:09 localhost.localdomain systemd[1]: httpd.service failed.
```

```
[vagrant@localhost ~]$ sudo getenforce 
Enforcing
```

```
[vagrant@localhost ~]$ sudo setenforce Permissive
```

```
[vagrant@localhost ~]$ sudo systemctl start httpd.service
```

```
[vagrant@localhost ~]$ sudo systemctl -l status httpd.service
● httpd.service - The Apache HTTP Server
   Loaded: loaded (/usr/lib/systemd/system/httpd.service; disabled; vendor preset: disabled)
   Active: active (running) since Fri 2017-10-20 22:34:39 UTC; 7s ago
     Docs: man:httpd(8)
           man:apachectl(8)
  Process: 27042 ExecStop=/bin/kill -WINCH ${MAINPID} (code=exited, status=1/FAILURE)
 Main PID: 27161 (httpd)
   Status: "Processing requests..."
   CGroup: /system.slice/httpd.service
           ├─27161 /usr/sbin/httpd -DFOREGROUND
           ├─27162 (wsgi:keystone- -DFOREGROUND
           ├─27163 (wsgi:keystone- -DFOREGROUND
           ├─27164 (wsgi:keystone- -DFOREGROUND
           ├─27165 (wsgi:keystone- -DFOREGROUND
           ├─27166 (wsgi:keystone- -DFOREGROUND
           ├─27167 (wsgi:keystone- -DFOREGROUND
           ├─27168 (wsgi:keystone- -DFOREGROUND
           ├─27169 (wsgi:keystone- -DFOREGROUND
           ├─27170 (wsgi:keystone- -DFOREGROUND
           ├─27171 (wsgi:keystone- -DFOREGROUND
           ├─27172 /usr/sbin/httpd -DFOREGROUND
           ├─27173 /usr/sbin/httpd -DFOREGROUND
           ├─27174 /usr/sbin/httpd -DFOREGROUND
           ├─27175 /usr/sbin/httpd -DFOREGROUND
           └─27176 /usr/sbin/httpd -DFOREGROUND

Oct 20 22:34:39 localhost.localdomain systemd[1]: Starting The Apache HTTP Server...
Oct 20 22:34:39 localhost.localdomain systemd[1]: Started The Apache HTTP Server.
```

```
[vagrant@localhost ~]$ sudo systemctl enable httpd.service
Created symlink from /etc/systemd/system/multi-user.target.wants/httpd.service to /usr/lib/systemd/system/httpd.service.
```

client
```
[vagrant@localhost ~]$ cat << EOF > openstack-admin.sh
> export OS_USERNAME=admin
> export OS_PASSWORD=ADMIN_PASS
> export OS_PROJECT_NAME=admin
> export OS_USER_DOMAIN_NAME=Default
> export OS_PROJECT_DOMAIN_NAME=Default
> export OS_AUTH_URL=http://10.64.33.64:35357/v3
> export OS_IDENTITY_API_VERSION=3
> EOF
```

```
[vagrant@localhost ~]$ . openstack-admin.sh 
```

```
[vagrant@localhost ~]$ echo $OS_USERNAME
admin
```

```
[vagrant@localhost ~]$ openstack project list
+----------------------------------+-------+
| ID                               | Name  |
+----------------------------------+-------+
| a0be38aef8c74d4abca3e4e100ee7910 | admin |
+----------------------------------+-------+
[vagrant@localhost ~]$ openstack user list
+----------------------------------+-------+
| ID                               | Name  |
+----------------------------------+-------+
| 44e6ee1df8ae436986d2d50f7b358aa0 | admin |
+----------------------------------+-------+
[vagrant@localhost ~]$ openstack role list
+----------------------------------+----------+
| ID                               | Name     |
+----------------------------------+----------+
| 9ce0832112464e818686139e46cb1d78 | admin    |
| 9fe2ff9ee4384b1894a90878d3e92bab | _member_ |
+----------------------------------+----------+
[vagrant@localhost ~]$ openstack token issue
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Field      | Value                                                                                                                                                                                   |
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| expires    | 2017-10-20T23:52:00+0000                                                                                                                                                                |
| id         | gAAAAABZ6n4QQ_CyuypadbPv0uOeTKwgjvBDSUmagF0m6pQE0dIMV3UIbVLCfPUkB3E3sjR0n0W1IVDGerCZqLpbtG-auNDQu0ZmAG1h4KHffvk4xpjhntmVqeB74Uu59N_s3wT8RwZiO35whnzX-rccBDHp65ZCBt_6Ecfu7c0T3EYNLCYqwLc |
| project_id | a0be38aef8c74d4abca3e4e100ee7910                                                                                                                                                        |
| user_id    | 44e6ee1df8ae436986d2d50f7b358aa0                                                                                                                                                        |
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
[vagrant@localhost ~]$ openstack service list
+----------------------------------+----------+----------+
| ID                               | Name     | Type     |
+----------------------------------+----------+----------+
| 896a7add738d40fb931cd704001d99e9 | keystone | identity |
+----------------------------------+----------+----------+
[vagrant@localhost ~]$ openstack endpoint list
+----------------------------------+-----------+--------------+--------------+---------+-----------+------------------------------+
| ID                               | Region    | Service Name | Service Type | Enabled | Interface | URL                          |
+----------------------------------+-----------+--------------+--------------+---------+-----------+------------------------------+
| 1c05dda3e9144422bd31c9c91ace6011 | RegionOne | keystone     | identity     | True    | internal  | http://10.64.33.64:5000/v3/  |
| 638c40e1f33f42d78bc4cead7e6e2438 | RegionOne | keystone     | identity     | True    | admin     | http://10.64.33.64:35357/v3/ |
| a11e9e70b7f74b43856d349e6ff8e9da | RegionOne | keystone     | identity     | True    | public    | http://10.64.33.64:5000/v3/  |
+----------------------------------+-----------+--------------+--------------+---------+-----------+------------------------------+
```

### Project service

create 
```
[vagrant@localhost ~]$ openstack project create --domain default --description "Service Project" service
+-------------+----------------------------------+
| Field       | Value                            |
+-------------+----------------------------------+
| description | Service Project                  |
| domain_id   | default                          |
| enabled     | True                             |
| id          | 250900c3faf44b3197befe24d0b8f681 |
| is_domain   | False                            |
| name        | service                          |
| parent_id   | default                          |
+-------------+----------------------------------+
```

### User

demo
```
[vagrant@localhost ~]$ openstack project create --domain default --description "Demo Project" demo
+-------------+----------------------------------+
| Field       | Value                            |
+-------------+----------------------------------+
| description | Demo Project                     |
| domain_id   | default                          |
| enabled     | True                             |
| id          | b7ce36a2bb124b9095b22ba8557e0fc4 |
| is_domain   | False                            |
| name        | demo                             |
| parent_id   | default                          |
+-------------+----------------------------------+
[vagrant@localhost ~]$ openstack user create --domain default --password DEMO_PASS demo
+---------------------+----------------------------------+
| Field               | Value                            |
+---------------------+----------------------------------+
| domain_id           | default                          |
| enabled             | True                             |
| id                  | e9524a68c74a4c6badcf36cb325b708a |
| name                | demo                             |
| options             | {}                               |
| password_expires_at | None                             |
+---------------------+----------------------------------+
[vagrant@localhost ~]$ openstack role create user
+-----------+----------------------------------+
| Field     | Value                            |
+-----------+----------------------------------+
| domain_id | None                             |
| id        | 5ea27488aec94433baddca083ce893be |
| name      | user                             |
+-----------+----------------------------------+
[vagrant@localhost ~]$ openstack role add --project demo --user demo user
```

```
[vagrant@localhost ~]$ . openstack-demo.sh 
```

```
[vagrant@localhost ~]$ openstack user show demo
+---------------------+----------------------------------+
| Field               | Value                            |
+---------------------+----------------------------------+
| domain_id           | default                          |
| enabled             | True                             |
| id                  | e9524a68c74a4c6badcf36cb325b708a |
| name                | demo                             |
| options             | {}                               |
| password_expires_at | None                             |
+---------------------+----------------------------------+
[vagrant@localhost ~]$ openstack project show demo
+-------------+----------------------------------+
| Field       | Value                            |
+-------------+----------------------------------+
| description | Demo Project                     |
| domain_id   | default                          |
| enabled     | True                             |
| id          | b7ce36a2bb124b9095b22ba8557e0fc4 |
| is_domain   | False                            |
| name        | demo                             |
| parent_id   | default                          |
+-------------+----------------------------------+
[vagrant@localhost ~]$ openstack token issue
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| Field      | Value                                                                                                                                                                                   |
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| expires    | 2017-10-21T00:00:28+0000                                                                                                                                                                |
| id         | gAAAAABZ6oAMt3x50KwjQSwLhO2zM6HWuZ5KjvEhbDjVHJEmgci8Ldjzsb3IEInn7qe38ADuFP5MMaDBOgQgr2niu72qRJu9-Bl9cgtx8iJ8tCqcyiChrQChpRG8sx4MCb7KDY4D9r04li9Otr05RDgvoWXhXQIjWFSb0sy8y-pxlHKdYQIvEHE |
| project_id | b7ce36a2bb124b9095b22ba8557e0fc4                                                                                                                                                        |
| user_id    | e9524a68c74a4c6badcf36cb325b708a                                                                                                                                                        |
+------------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
```

```
[vagrant@localhost ~]$ sudo ls /var/log/httpd
access_log  error_log  keystone_access.log  keystone.log
```

```
[vagrant@localhost ~]$ sudo ls -l /var/log/httpd
total 20
-rw-r--r--. 1 root root     0 Oct 20 22:34 access_log
-rw-r--r--. 1 root root   782 Oct 20 22:34 error_log
-rw-r--r--. 1 root root 13482 Oct 20 23:00 keystone_access.log
-rw-r--r--. 1 root root     0 Oct 20 22:34 keystone.log
[vagrant@localhost ~]$ sudo tail /var/log/httpd/error_log
[Fri Oct 20 22:34:39.284530 2017] [core:notice] [pid 27161] SELinux policy enabled; httpd running as context system_u:system_r:httpd_t:s0
[Fri Oct 20 22:34:39.285354 2017] [suexec:notice] [pid 27161] AH01232: suEXEC mechanism enabled (wrapper: /usr/sbin/suexec)
[Fri Oct 20 22:34:39.293713 2017] [auth_digest:notice] [pid 27161] AH01757: generating secret for digest authentication ...
[Fri Oct 20 22:34:39.294646 2017] [lbmethod_heartbeat:notice] [pid 27161] AH02282: No slotmem from mod_heartmonitor
[Fri Oct 20 22:34:39.298641 2017] [mpm_prefork:notice] [pid 27161] AH00163: Apache/2.4.6 (CentOS) mod_wsgi/3.4 Python/2.7.5 configured -- resuming normal operations
[Fri Oct 20 22:34:39.298663 2017] [core:notice] [pid 27161] AH00094: Command line: '/usr/sbin/httpd -D FOREGROUND'
```

```
[vagrant@localhost ~]$ sudo ls /var/log/keystone/
keystone.log
```

```
[vagrant@localhost ~]$ sudo tail 10 /var/log/keystone/keystone.log
tail: cannot open ‘10’ for reading: No such file or directory
==> /var/log/keystone/keystone.log <==
2017-10-20 23:00:06.942 27169 INFO keystone.common.wsgi [req-6a15be5f-0f5e-4a55-9cbf-7e3c9bb97f61 e9524a68c74a4c6badcf36cb325b708a b7ce36a2bb124b9095b22ba8557e0fc4 - default default] GET http://10.64.33.64:35357/v3/projects/b7ce36a2bb124b9095b22ba8557e0fc4
2017-10-20 23:00:06.943 27169 DEBUG keystone.common.authorization [req-6a15be5f-0f5e-4a55-9cbf-7e3c9bb97f61 e9524a68c74a4c6badcf36cb325b708a b7ce36a2bb124b9095b22ba8557e0fc4 - default default] RBAC: Authorizing identity:get_project(project_id=b7ce36a2bb124b9095b22ba8557e0fc4) _build_policy_check_credentials /usr/lib/python2.7/site-packages/keystone/common/authorization.py:137
2017-10-20 23:00:06.945 27169 DEBUG keystone.policy.backends.rules [req-6a15be5f-0f5e-4a55-9cbf-7e3c9bb97f61 e9524a68c74a4c6badcf36cb325b708a b7ce36a2bb124b9095b22ba8557e0fc4 - default default] enforce identity:get_project: {'is_delegated_auth': False, 'access_token_id': None, 'user_id': u'e9524a68c74a4c6badcf36cb325b708a', 'roles': [u'user'], 'user_domain_id': u'default', 'consumer_id': None, 'trustee_id': None, 'is_domain': False, 'is_admin_project': True, 'trustor_id': None, 'token': <KeystoneToken (audit_id=agRg6402Q-qi7S9iPn514A, audit_chain_id=agRg6402Q-qi7S9iPn514A) at 0x7f9d140579f0>, 'project_id': u'b7ce36a2bb124b9095b22ba8557e0fc4', 'trust_id': None, 'project_domain_id': u'default'} enforce /usr/lib/python2.7/site-packages/keystone/policy/backends/rules.py:33
2017-10-20 23:00:06.947 27169 DEBUG keystone.common.authorization [req-6a15be5f-0f5e-4a55-9cbf-7e3c9bb97f61 e9524a68c74a4c6badcf36cb325b708a b7ce36a2bb124b9095b22ba8557e0fc4 - default default] RBAC: Authorization granted check_policy /usr/lib/python2.7/site-packages/keystone/common/authorization.py:240
2017-10-20 23:00:28.644 27166 DEBUG keystone.middleware.auth [req-7a47326f-8390-4a7f-839a-052f09ee1400 - - - - -] There is either no auth token in the request or the certificate issuer is not trusted. No auth context will be set. fill_context /usr/lib/python2.7/site-packages/keystone/middleware/auth.py:203
2017-10-20 23:00:28.644 27166 INFO keystone.common.wsgi [req-7a47326f-8390-4a7f-839a-052f09ee1400 - - - - -] GET http://10.64.33.64:5000/v3/
2017-10-20 23:00:28.651 27163 DEBUG keystone.middleware.auth [req-49791b1b-fd98-41e5-89fb-30745b997a1a - - - - -] There is either no auth token in the request or the certificate issuer is not trusted. No auth context will be set. fill_context /usr/lib/python2.7/site-packages/keystone/middleware/auth.py:203
2017-10-20 23:00:28.652 27163 INFO keystone.common.wsgi [req-49791b1b-fd98-41e5-89fb-30745b997a1a - - - - -] POST http://10.64.33.64:5000/v3/auth/tokens
2017-10-20 23:00:28.959 27163 DEBUG keystone.auth.core [req-49791b1b-fd98-41e5-89fb-30745b997a1a - - - - -] MFA Rules not processed for user `e9524a68c74a4c6badcf36cb325b708a`. Rule list: `[]` (Enabled: `True`). check_auth_methods_against_rules /usr/lib/python2.7/site-packages/keystone/auth/core.py:388
2017-10-20 23:00:28.993 27163 DEBUG keystone.common.fernet_utils [req-49791b1b-fd98-41e5-89fb-30745b997a1a - - - - -] Loaded 2 Fernet keys from /etc/keystone/fernet-keys/, but `[fernet_tokens] max_active_keys = 3`; perhaps there have not been enough key rotations to reach `max_active_keys` yet? load_keys /usr/lib/python2.7/site-packages/keystone/common/fernet_utils.py:306
```