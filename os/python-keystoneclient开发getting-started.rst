OpenStack python-keystoneclient开发 getting started
===================================================
.. contents::

.. _prerequisite:
Pre-requisite
-------------
1. Linux (以下为cygwin模拟环境)
2. Python (包括开发环境和工具，如pip，virtualenv)
3. Git

.. _communityresources:
Community Resources
-------------------
1. Repositories
    * `github.com`_ - https://github.com/openstack/python-keystoneclient
    * `openstack.org`_ - http://git.openstack.org/cgit/
2. Bug Report
    * `launchpad.net`_ - https://launchpad.net/python-keystoneclient
3. Developer Documentation
    * keystone docs - http://docs.openstack.org/developer/keystone/
    * keystone wiki - https://wiki.openstack.org/wiki/Keystone
        
.. _github.com: https://github.com/
.. _openstack.org: https://www.openstack.org/
.. _launchpad.net: https://www.launchpad.net/

示例 git clone 
^^^^^^^^^^^^^^^^
*bash*::

    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel
    $ git clone https://github.com/openstack/python-keystoneclient.git

.. _staging:
安装Staging环境
---------------
* 使用tox将keystone client安装到默认的virtuaenv，后面有tox的较详细的安装和设置章节
*bash*::

    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ tox -e py27
    py27 create: /home/Administrator/github.com/openstack/python-keystoneclient/.tox/py27
    py27 installdeps: -r/home/Administrator/github.com/openstack/python-keystoneclient/requirements.txt, -r/home/Administrator/github.com/openstack/python-keystoneclient/test-requirements.txt
    py27 develop-inst: /home/Administrator/github.com/openstack/python-keystoneclient
    py27 runtests: commands[0] | python setup.py testr --testr-args=
    running testr
    running=${PYTHON:-python} -m subunit.run discover -t ./ ./keystoneclient/tests --list
    running=${PYTHON:-python} -m subunit.run discover -t ./ ./keystoneclient/tests  --load-list /tmp/tmpoS1qqw
    running=${PYTHON:-python} -m subunit.run discover -t ./ ./keystoneclient/tests  --load-list /tmp/tmpafUSWb
    No handlers could be found for logger "keystoneclient.middleware.auth_token"
    No handlers could be found for logger "keystoneclient.middleware.auth_token"
    ======================================================================
    FAIL: keystoneclient.tests.v3.test_access.AccessInfoTest.test_will_expire_soon
    tags: worker-1
    ----------------------------------------------------------------------
    pythonlogging:'': {{{
    Parsed 2014-12-07T14:18:40.781250 into {'tz_sign': None, 'second_fraction': '781250', 'hour': '14', 'daydash': '07', 'tz_hour': None, 'month': None, 'timezone': None, 'second': '40', 'tz_minute': None, 'year': '2014', 'separator': 'T', 'monthdash': '12', 'day': None, 'minute': '18'} with default timezone <iso8601.iso8601.Utc object at 0x7fb084cc>
    Got '2014' for 'year' with default None    
    Got '12' for 'monthdash' with default 1
    Got 12 for 'month' with default 12
    Got '07' for 'daydash' with default 1
    Got 7 for 'day' with default 7
    Got '14' for 'hour' with default None
    Got '18' for 'minute' with default None
    Got '40' for 'second' with default None
    Parsed 2014-12-07T14:18:40.781250 into {'tz_sign': None, 'second_fraction': '781250', 'hour': '14', 'daydash': '07', 'tz_hour': None, 'month': None, 'timezone': None, 'second': '40', 'tz_minute': None, 'year': '2014', 'separator': 'T', 'monthdash': '12', 'day': None, 'minute': '18'} with default timezone <iso8601.iso8601.Utc object at 0x7fb084cc>
    Got '2014' for 'year' with default None    
    Got '12' for 'monthdash' with default 1
    Got 12 for 'month' with default 12
    Got '07' for 'daydash' with default 1
    Got 7 for 'day' with default 7
    Got '14' for 'hour' with default None
    Got '18' for 'minute' with default None
    Got '40' for 'second' with default None
    }}}
    Traceback (most recent call last):
      File "keystoneclient/tests/v3/test_access.py", line 79, in test_will_expire_soon
        self.assertTrue(auth_ref.will_expire_soon(stale_duration=300))
      File "/home/Administrator/github.com/openstack/python-keystoneclient/.tox/py27/lib/python2.7/site-packages/unittest2/case.py", line 678, in assertTrue
        raise self.failureException(msg)
    AssertionError: False is not true
    Ran 976 tests in 16.297s
    FAILED (id=0, failures=1, skips=3)
    error: testr failed (1)
    ERROR: InvocationError: '/home/Administrator/github.com/openstack/python-keystoneclient/.tox/py27/bin/python setup.py testr --testr-args='
    ___________________________________ summary ____________________________________
    ERROR:   py27: commands failed

* 切换到tox创建的virtualenv，操作keystone client命令
*bash*::

    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ . .tox/py27/bin/activate
    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ keystone --version
    0.11.2.43
    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ which keystone
    /home/Administrator/github.com/openstack/python-keystoneclient/.tox/py27/bin/keystone


.. _gettingstarted:
开发第一步
----------
.. _appentry:
Keystone Client代码入口
^^^^^^^^^^^^^^^^^^^^^^^
* main函数
*bash*::

    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ ls .tox/py27/bin/
    activate               pip2.7                 sphinx-build
    activate.csh           pybabel                sphinx-quickstart
    activate.fish          pyflakes               subunit-1to2
    activate_this.py       pygmentize             subunit-2to1
    coverage               python                 subunit-filter
    coverage-2.7           python2                subunit-ls
    coverage2              python2.7              subunit-notify
    discover               rst2html.py            subunit-output
    easy_install           rst2latex.py           subunit-stats
    easy_install-2.7       rst2man.py             subunit-tags
    flake8                 rst2odt.py             subunit2csv
    keyring                rst2odt_prepstyles.py  subunit2gtk
    keystone               rst2pseudoxml.py       subunit2junitxml
    netaddr                rst2s5.py              subunit2pyunit
    oslo-config-generator  rst2xetex.py           tap2subunit
    oslo_debug_helper      rst2xml.py             testr
    pep8                   rstpep2html.py         unit2
    pip                    sphinx-apidoc
    pip2                   sphinx-autogen
    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ cat .tox/py27/bin/keystone
    #!/home/Administrator/github.com/openstack/python-keystoneclient/.tox/py27/bin/python2.7
    # PBR Generated from u'console_scripts'

    import sys

    from keystoneclient.shell import main

    if __name__ == "__main__":
        sys.exit(main())
