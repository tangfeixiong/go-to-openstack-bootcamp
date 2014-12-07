OpenStack Keystone开始
======================
main函数
--------
Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ vim bin/keystone-all
#!/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/.tox/py27/bin/python2.7
# PBR Generated from u'console_scripts'

import sys

from keystoneclient.shell import main


if __name__ == "__main__":
    sys.exit(main())


	
Developer Documentation
-----------------------
http://docs.openstack.org/developer/keystone/

Repositories
------------
* Github.com
git clone -b stable/juno https://github.com/openstack/keystone.git
* Launchpad.net
https://launchpad.net/keystone/+download?direction=backwards&start=50
* OpenStack.org
http://git.openstack.org/cgit/openstack/keystone/

Developing With Keystone
------------------------
* Pre-requisite
    install Python 2.6/2.7 
	install Phthon Utilities:
	    1. pip
		2. virtualenv
* Test
Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ pip install tox
Requirement already satisfied (use --upgrade to upgrade): tox in /usr/lib/python2.7/site-packages
Requirement already satisfied (use --upgrade to upgrade): virtualenv>=1.9.1 in /usr/lib/python2.7/site-packages/virtualenv-1.10.1-py2.7.egg (from tox)
Requirement already satisfied (use --upgrade to upgrade): py>=1.4.15 in /usr/lib/python2.7/site-packages (from tox)
Cleaning up...

Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ tox
py26 create: /home/Administrator/python-workspace/openstack-devel/keystone-2014.2/.tox/py26
ERROR: InterpreterNotFound: python2.6
py27 create: /home/Administrator/python-workspace/openstack-devel/keystone-2014.2/.tox/py27
py27 installdeps: -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt
...
ERROR: could not install deps [-r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt]
py33 create: /home/Administrator/python-workspace/openstack-devel/keystone-2014.2/.tox/py33
ERROR: InterpreterNotFound: python3.3
py34 create: /home/Administrator/python-workspace/openstack-devel/keystone-2014.2/.tox/py34
ERROR: InterpreterNotFound: python3.4
pep8 create: /home/Administrator/python-workspace/openstack-devel/keystone-2014.2/.tox/pep8
pep8 installdeps: -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt
...
ERROR: could not install deps [-r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt]
docs create: /home/Administrator/python-workspace/openstack-devel/keystone-2014.2/.tox/docs
docs installdeps: -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt
...
ERROR: could not install deps [-r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt]
sample_config create: /home/Administrator/python-workspace/openstack-devel/keystone-2014.2/.tox/sample_config
ERROR: could not install deps [-r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt]
___________________________________ summary ____________________________________
ERROR:   py26: InterpreterNotFound: python2.6
ERROR:   py27: could not install deps [-r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt]
ERROR:   py33: InterpreterNotFound: python3.3
ERROR:   py34: InterpreterNotFound: python3.4
ERROR:   pep8: could not install deps [-r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt]
ERROR:   docs: could not install deps [-r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt]
ERROR:   sample_config: could not install deps [-r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/requirements.txt, -r/home/Administrator/python-workspace/openstack-devel/keystone-2014.2/test-requirements.txt]

Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ cat tox.ini
[tox]
minversion = 1.6
skipsdist = True
envlist = py26,py27,py33,py34,pep8,docs,sample_config

[testenv]
usedevelop = True
install_command = pip install -U {opts} {packages}
setenv = VIRTUAL_ENV={envdir}
         # FIXME(dolph): overriding the hash seed with a constant is a
         # workaround for bug 1348818
         PYTHONHASHSEED=0
deps = -r{toxinidir}/requirements.txt
       -r{toxinidir}/test-requirements.txt
commands = python setup.py testr --slowest --testr-args='{posargs}'
whitelist_externals = bash

[testenv:py33]
deps = -r{toxinidir}/requirements-py3.txt
       -r{toxinidir}/test-requirements-py3.txt
       nose
commands =
  nosetests --with-coverage --cover-package=keystone \
      --exclude test_ldap \
      keystone/tests/test_auth_plugin.py \
      keystone/tests/test_backend.py \
      keystone/tests/test_backend_rules.py \
      keystone/tests/test_cache_backend_mongo.py \
      keystone/tests/test_contrib_stats_core.py \
      keystone/tests/test_driver_hints.py \
      keystone/tests/test_hacking_checks.py \
      keystone/tests/test_injection.py \
      keystone/tests/test_matchers.py \
      keystone/tests/test_policy.py \
      keystone/tests/test_singular_plural.py \
      keystone/tests/test_sizelimit.py \
      keystone/tests/test_sql_migrate_extensions.py \
      keystone/tests/test_token_bind.py \
      keystone/tests/test_url_middleware.py \
      keystone/tests/test_utils.py \
      keystone/tests/test_validation.py \
      keystone/tests/test_v3_controller.py \
      keystone/tests/test_wsgi.py \
      keystone/tests/unit

[testenv:py34]
deps = -r{toxinidir}/requirements-py3.txt
       -r{toxinidir}/test-requirements-py3.txt
       nose
commands =
  nosetests --with-coverage --cover-package=keystone \
      --exclude test_ldap \
      keystone/tests/test_auth_plugin.py \
      keystone/tests/test_backend.py \
      keystone/tests/test_backend_rules.py \
      keystone/tests/test_cache_backend_mongo.py \
      keystone/tests/test_contrib_stats_core.py \
      keystone/tests/test_driver_hints.py \
      keystone/tests/test_hacking_checks.py \
      keystone/tests/test_injection.py \
      keystone/tests/test_matchers.py \
      keystone/tests/test_policy.py \
      keystone/tests/test_singular_plural.py \
      keystone/tests/test_sizelimit.py \
      keystone/tests/test_sql_migrate_extensions.py \
      keystone/tests/test_token_bind.py \
      keystone/tests/test_url_middleware.py \
      keystone/tests/test_utils.py \
      keystone/tests/test_validation.py \
      keystone/tests/test_v3_controller.py \
      keystone/tests/test_wsgi.py \
      keystone/tests/unit

[testenv:pep8]
commands =
  flake8 {posargs}
  # Run bash8 during pep8 runs to ensure violations are caught by
  # the check and gate queues
  bashate examples/pki/gen_pki.sh
  # Check that .po and .pot files are valid.
  # NOTE(jaegerandi): We search for files ending with '.po' or '.pot'.
  # The regex '.*\.pot?' does not work on OS X and we assume there are no
  # files with more than one "t" that have to be ignored.
  bash -c "find keystone -type f -regex '.*\.pot*' -print0| \
           xargs -0 -n 1 msgfmt --check-format -o /dev/null"

[tox:jenkins]
downloadcache = ~/cache/pip

[testenv:cover]
commands = python setup.py testr --coverage --testr-args='{posargs}'

[testenv:venv]
commands = {posargs}

[testenv:debug]
commands = oslo_debug_helper.sh {posargs}

[flake8]
filename= *.py,keystone-*
show-source = true

# H104  File contains nothing but comments
# H405  multi line docstring summary not separated with an empty line
# H803  Commit message should not end with a period (do not remove per list discussion)
# H904  Wrap long lines in parentheses instead of a backslash
ignore = H104,H405,H803,H904

builtins = _
exclude=.venv,.git,.tox,build,dist,doc,*openstack/common*,*lib/python*,*egg,tools,vendor,.update-venv,*.ini,*.po,*.pot

[testenv:docs]
commands=
    python setup.py build_sphinx

[testenv:sample_config]
commands = {toxinidir}/tools/config/generate_sample.sh

[hacking]
import_exceptions =
  keystone.i18n
local-check-factory = keystone.hacking.checks.factory

Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ ls .tox
log  pep8  py27  sample_config

Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ ls .tox/py27/
bin  etc  include  lib  log

Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ ls .tox/py27/bin/
activate          easy_install      migrate-repository           pip-2.7
activate.csh      easy_install-2.7  netaddr                      pybabel
activate.fish     jsonschema        oslo-config-generator        python
activate_this.py  keystone          oslo-messaging-zmq-receiver  python2
bashate           migrate           pip                          python2.7




tox
===
https://pypi.python.org/pypi/tox/1.8.1
https://testrun.org/tox/latest/

Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ tox --version
1.6.1 imported from /usr/lib/python2.7/site-packages/tox/__init__.pyc

Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ pip install --upgrade tox
Cannot fetch index base URL http://pypi.douban.com/simple/
http://pypi.douban.com/simple/tox/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
Downloading/unpacking tox from http://pypi.douban.com/packages/source/t/tox/tox-1.8.1.tar.gz#md5=c4423cc6512932b37e5b0d1faa87bef2
  Downloading tox-1.8.1.tar.gz (90kB): 90kB downloaded
  Running setup.py (path:/tmp/pip_build_Administrator/tox/setup.py) egg_info for package tox

Downloading/unpacking virtualenv>=1.11.2 (from tox)
  http://pypi.douban.com/simple/virtualenv/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
  Downloading virtualenv-1.11.6-py2.py3-none-any.whl (1.6MB): 1.6MB downloaded
http://pypi.douban.com/simple/py/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
Downloading/unpacking py>=1.4.17 from http://pypi.douban.com/packages/source/p/py/py-1.4.26.tar.gz#md5=30c3fd92a53f1a5ed6f3591c1fe75c0e (from tox)
  Downloading py-1.4.26.tar.gz (190kB): 190kB downloaded
  Running setup.py (path:/tmp/pip_build_Administrator/py/setup.py) egg_info for package py

Installing collected packages: tox, virtualenv, py
  Found existing installation: tox 1.6.1
    Uninstalling tox:
      Successfully uninstalled tox
  Running setup.py install for tox

    Installing tox script to /usr/bin
    Installing tox-quickstart script to /usr/bin
  Found existing installation: virtualenv 1.10.1
    Uninstalling virtualenv:
      Successfully uninstalled virtualenv
  Found existing installation: py 1.4.18
    Uninstalling py:
      Successfully uninstalled py
  Running setup.py install for py

Successfully installed tox virtualenv py
Cleaning up...

Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ tox -e py27
...

Pypi Mirror
===========
Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
$ cat ~/.pip/pip.conf
[global]
#index-url=https://pypi.python.org/simple
index-url=http://pypi.douban.com/simple

http://www.pypi-mirrors.org/
https://pypi.python.org/mirrors