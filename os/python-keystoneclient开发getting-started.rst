OpenStack python-keystoneclient开发 getting started
===================================================
.. contents::

Pre-requisite
-------------
1. Linux (以下为cygwin模拟环境)
2. Python (包括开发环境和工具，如pip，virtualenv)
3. Git

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


开发第一步
----------

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

* keystoneclient包和shell模块
*bash*::

    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ ls .tox/py27/lib/python2.7/site-packages/
    Babel-1.3-py2.7.egg-info           oslo.config-1.5.0-py2.7.egg-info
    Crypto                             oslo.i18n-1.1.0-py2.7-nspkg.pth
    Jinja2-2.7.3-py2.7.egg-info        oslo.i18n-1.1.0-py2.7.egg-info
    MarkupSafe-0.23-py2.7.egg-info     oslo.serialization-1.1.0-py2.7-nspkg.pth
    Pygments-2.0.1-py2.7.egg-info      oslo.serialization-1.1.0-py2.7.egg-info
    Sphinx-1.2.3-py2.7.egg-info        oslo.utils-1.1.0-py2.7-nspkg.pth
    WebOb-1.4-py2.7.egg-info           oslo.utils-1.1.0-py2.7.egg-info
    _markerlib                         oslosphinx
    argparse-1.2.2-py2.7.egg-info      oslosphinx-2.3.0-py2.7.egg-info
    argparse.py                        oslotest
    argparse.pyc                       oslotest-1.3.0-py2.7.egg-info
    babel                              pbr
    coverage                           pbr-0.10.0-py2.7.egg-info
    coverage-3.7.1-py2.7.egg-info      pep8-1.5.6-py2.7.egg-info
    discover-0.4.0-py2.7.egg-info      pep8.py
    discover.py                        pep8.pyc
    discover.pyc                       pip
    docutils                           pip-1.5.6-py2.7.egg-info
    docutils-0.12-py2.7.egg-info       pkg_resources.py
    easy-install.pth                   pkg_resources.pyc
    easy_install.py                    prettytable-0.7.2-py2.7.egg-info
    easy_install.pyc                   prettytable.py
    extras                             prettytable.pyc
    extras-0.0.3-py2.7.egg-info        pycrypto-2.6.1-py2.7.egg-info
    fixtures                           pyflakes
    fixtures-1.0.0-py2.7.egg-info      pyflakes-0.8.1-py2.7.egg-info
    flake8                             pygments
    flake8-2.1.0-py2.7.egg-info        python-keystoneclient.egg-link
    hacking                            python_mimeparse-0.1.4-py2.7.egg-info
    hacking-0.9.5-py2.7.egg-info       python_subunit-1.0.0-py2.7.egg-info
    iso8601                            pytz
    iso8601-0.1.10-py2.7.egg-info      pytz-2014.10-py2.7.egg-info
    jinja2                             requests
    keyring                            requests-2.5.0-py2.7.egg-info
    keyring-4.0-py2.7.egg-info         requests_mock
    lxml                               requests_mock-0.5.1-py2.7.egg-info
    lxml-3.4.1-py2.7.egg-info          setuptools
    markupsafe                         setuptools-0.9.8-py2.7.egg-info
    mccabe-0.2.1-py2.7.egg-info        six-1.8.0-py2.7.egg-info
    mccabe.py                          six.py
    mccabe.pyc                         six.pyc
    mimeparse.py                       sphinx
    mimeparse.pyc                      stevedore
    mock-1.0.1-py2.7.egg-info          stevedore-1.1.0-py2.7.egg-info
    mock.py                            subunit
    mock.pyc                           testrepository
    mox3                               testrepository-0.0.20-py2.7.egg-info
    mox3-0.7.0-py2.7.egg-info          testresources
    netaddr                            testresources-0.2.7-py2.7.egg-info
    netaddr-0.7.12-py2.7.egg-info      testscenarios
    netifaces-0.10.4-py2.7.egg-info    testscenarios-0.4-py2.7.egg-info
    netifaces.dll                      testtools
    oauthlib                           testtools-1.5.0-py2.7.egg-info
    oauthlib-0.7.2-py2.7.egg-info      unittest2
    oslo                               unittest2-0.8.0-py2.7.egg-info
    oslo.config-1.5.0-py2.7-nspkg.pth  webob
    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ cat .tox/py27/lib/python2.7/site-packages/python-keystoneclient.egg-link
    /home/Administrator/github.com/openstack/python-keystoneclient
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    (py27)
    $ ls keystoneclient/
    __init__.py    auth            discover.py     i18n.pyc             shell.pyc
    __init__.pyc   base.py         discover.pyc    locale               tests
    _discover.py   base.pyc        exceptions.py   middleware           utils.py
    _discover.pyc  baseclient.py   exceptions.pyc  openstack            utils.pyc
    access.py      baseclient.pyc  fixture         service_catalog.py   v2_0
    access.pyc     client.py       generic         service_catalog.pyc  v3
    adapter.py     client.pyc      httpclient.py   session.py
    adapter.pyc    common          httpclient.pyc  session.pyc
    apiclient      contrib         i18n.py         shell.py

* shell.py的main函数和OpenStackIdentityShell类，类方法main()
*bash*::

    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ cat keystoneclient/shell.py
    ...

    class OpenStackIdentityShell(object):

        def __init__(self, parser_class=argparse.ArgumentParser):
            self.parser_class = parser_class

        ...

        def main(self, argv):
            
            ...

            # Handle top-level --help/-h before attempting to parse
            # a command off the command line
            if not argv or options.help:
                self.do_help(options)
                return 0

            # Parse args again and call whatever callback was selected
            args = subcommand_parser.parse_args(argv)

            # Short-circuit and deal with help command right away.
            if args.func == self.do_help:
                self.do_help(args)
                return 0
            elif args.func == self.do_bash_completion:
                self.do_bash_completion(args)
                return 0

            if args.debug:
                logging_level = logging.DEBUG
                iso_logger = logging.getLogger('iso8601')
                iso_logger.setLevel('WARN')
            else:
                logging_level = logging.WARNING

            logging.basicConfig(level=logging_level)

            ...

            if utils.isunauthenticated(args.func):
                self.cs = shell_generic.CLIENT_CLASS(endpoint=args.os_auth_url,
                                                     cacert=args.os_cacert,
                                                     key=args.os_key,
                                                     cert=args.os_cert,
                                                     insecure=args.insecure,
                                                     timeout=args.timeout)
            else:
                self.auth_check(args)
                token = None
                if args.os_token and args.os_endpoint:
                    token = args.os_token
                api_version = options.os_identity_api_version
                self.cs = self.get_api_class(api_version)(
                    username=args.os_username,
                    tenant_name=args.os_tenant_name,
                    tenant_id=args.os_tenant_id,
                    token=token,
                    endpoint=args.os_endpoint,
                    password=args.os_password,
                    auth_url=args.os_auth_url,
                    region_name=args.os_region_name,
                    cacert=args.os_cacert,
                    key=args.os_key,
                    cert=args.os_cert,
                    insecure=args.insecure,
                    debug=args.debug,
                    use_keyring=args.os_cache,
                    force_new_token=args.force_new_token,
                    stale_duration=args.stale_duration,
                    timeout=args.timeout)

            try:
                args.func(self.cs, args)
            except exc.Unauthorized:
                raise exc.CommandError("Invalid OpenStack Identity credentials.")
            except exc.AuthorizationFailure:
                raise exc.CommandError("Unable to authorize user")

        ......
        
    def main():
        try:
            OpenStackIdentityShell().main(sys.argv[1:])

        except Exception as e:
            print(encodeutils.safe_encode(six.text_type(e)), file=sys.stderr)
            sys.exit(1)

    ...

* keystoneclient包的初始元数据
*bash*::

    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ cat keystoneclient/__init__.py
    ...

    """The python bindings for the OpenStack Identity (Keystone) project.

    See :py:class:`keystoneclient.v3.client.Client` for the Identity V3 client.

    See :py:class:`keystoneclient.v2_0.client.Client` for the Identity V2.0 client.

    """

    import pbr.version

    from keystoneclient import access
    from keystoneclient import client
    from keystoneclient import exceptions
    from keystoneclient import generic
    from keystoneclient import httpclient
    from keystoneclient import service_catalog
    from keystoneclient import v2_0
    from keystoneclient import v3


    __version__ = pbr.version.VersionInfo('python-keystoneclient').version_string()

    __all__ = [
        # Modules
        'generic',
        'v2_0',
        'v3',

        # Packages
        'access',
        'client',
        'exceptions',
        'httpclient',
        'service_catalog',
    ]

    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ python
    Python 2.7.8 (default, Jul 28 2014, 01:34:03)
    [GCC 4.8.3] on cygwin
    Type "help", "copyright", "credits" or "license" for more information.
    >>> import keystoneclient
    >>> dir(keystoneclient)
    ['__all__', '__builtins__', '__doc__', '__file__', '__name__', '__package__', '__path__', '__version__', '_discover', 'access', 'adapter', 'auth', 'base', 'baseclient', 'client', 'discover', 'exceptions', 'generic', 'httpclient', 'i18n', 'openstack', 'pbr', 'service_catalog', 'session', 'utils', 'v2_0', 'v3']
    >>> keystoneclient.__version__
    '0.11.2.43'
    >>> keystoneclient.__all__
    ['generic', 'v2_0', 'v3', 'access', 'client', 'exceptions', 'httpclient', 'service_catalog']

    >>> import keystoneclient.shell as python_keystoneclient_shell
    >>> dir(keystoneclient.shell)
    ['OpenStackHelpFormatter', 'OpenStackIdentityShell', '__builtins__', '__doc__', '__file__', '__name__', '__package__', 'access', 'argparse', 'encodeutils', 'env', 'exc', 'getpass', 'keystoneclient', 'logging', 'main', 'os', 'print_function', 'session', 'shell_bootstrap', 'shell_generic', 'shell_v2_0', 'six', 'sys', 'utils']

* 编程调用shell.py的main函数
*bash*::

    >>> python_keystoneclient_shell.main()
    usage: keystone [--version] [--debug] [--os-username <auth-user-name>]
                    [--os-password <auth-password>]
                    [--os-tenant-name <auth-tenant-name>]
                    [--os-tenant-id <tenant-id>] [--os-auth-url <auth-url>]
                    [--os-region-name <region-name>]
                    [--os-identity-api-version <identity-api-version>]
                    [--os-token <service-token>]
                    [--os-endpoint <service-endpoint>] [--os-cache]
                    [--force-new-token] [--stale-duration <seconds>] [--insecure]
                    [--os-cacert <ca-certificate>] [--os-cert <certificate>]
                    [--os-key <key>] [--timeout <seconds>]
                    <subcommand> ...

    Pending deprecation: Command-line interface to the OpenStack Identity API.
    This CLI is pending deprecation in favor of python-openstackclient. For a
    Python library, continue using python-keystoneclient.

    Positional arguments:
      <subcommand>
        catalog                 List service catalog, possibly filtered by
                                service.
        ec2-credentials-create  Create EC2-compatible credentials for user per
                                tenant.
        ec2-credentials-delete  Delete EC2-compatible credentials.
        ec2-credentials-get     Display EC2-compatible credentials.
        ec2-credentials-list    List EC2-compatible credentials for a user.
        endpoint-create         Create a new endpoint associated with a service.
        endpoint-delete         Delete a service endpoint.
        endpoint-get            Find endpoint filtered by a specific attribute or
                                service type.
        endpoint-list           List configured service endpoints.
        password-update         Update own password.
        role-create             Create new role.
        role-delete             Delete role.
        role-get                Display role details.
        role-list               List all roles.
        service-create          Add service to Service Catalog.
        service-delete          Delete service from Service Catalog.
        service-get             Display service from Service Catalog.
        service-list            List all services in Service Catalog.
        tenant-create           Create new tenant.
        tenant-delete           Delete tenant.
        tenant-get              Display tenant details.
        tenant-list             List all tenants.
        tenant-update           Update tenant name, description, enabled status.
        token-get               Display the current user token.
        user-create             Create new user.
        user-delete             Delete user.
        user-get                Display user details.
        user-list               List users.
        user-password-update    Update user password.
        user-role-add           Add role to user.
        user-role-list          List roles granted to a user.
        user-role-remove        Remove role from user.
        user-update             Update user's name, email, and enabled status.
        discover                Discover Keystone servers, supported API versions
                                and extensions.
        bootstrap               Grants a new role to a new user on a new tenant,
                                after creating each.
        bash-completion         Prints all of the commands and options to stdout.
        help                    Display help about this program or one of its
                                subcommands.

    Optional arguments:
      --version                 Shows the client version and exits.
      --debug                   Prints debugging output onto the console, this
                                includes the curl request and response calls.
                                Helpful for debugging and understanding the API
                                calls.
      --os-username <auth-user-name>
                                Name used for authentication with the OpenStack
                                Identity service. Defaults to env[OS_USERNAME].
      --os-password <auth-password>
                                Password used for authentication with the
                                OpenStack Identity service. Defaults to
                                env[OS_PASSWORD].
      --os-tenant-name <auth-tenant-name>
                                Tenant to request authorization on. Defaults to
                                env[OS_TENANT_NAME].
      --os-tenant-id <tenant-id>
                                Tenant to request authorization on. Defaults to
                                env[OS_TENANT_ID].
      --os-auth-url <auth-url>  Specify the Identity endpoint to use for
                                authentication. Defaults to env[OS_AUTH_URL].
      --os-region-name <region-name>
                                Specify the region to use. Defaults to
                                env[OS_REGION_NAME].
      --os-identity-api-version <identity-api-version>
                                Specify Identity API version to use. Defaults to
                                env[OS_IDENTITY_API_VERSION] or 2.0.
      --os-token <service-token>
                                Specify an existing token to use instead of
                                retrieving one via authentication (e.g. with
                                username & password). Defaults to
                                env[OS_SERVICE_TOKEN].
      --os-endpoint <service-endpoint>
                                Specify an endpoint to use instead of retrieving
                                one from the service catalog (via authentication).
                                Defaults to env[OS_SERVICE_ENDPOINT].
      --os-cache                Use the auth token cache. Defaults to
                                env[OS_CACHE].
      --force-new-token         If the keyring is available and in use, token will
                                always be stored and fetched from the keyring
                                until the token has expired. Use this option to
                                request a new token and replace the existing one
                                in the keyring.
      --stale-duration <seconds>
                                Stale duration (in seconds) used to determine
                                whether a token has expired when retrieving it
                                from keyring. This is useful in mitigating process
                                or network delays. Default is 30 seconds.
      --insecure                Explicitly allow client to perform "insecure" TLS
                                (https) requests. The server's certificate will
                                not be verified against any certificate
                                authorities. This option should be used with
                                caution.
      --os-cacert <ca-certificate>
                                Specify a CA bundle file to use in verifying a TLS
                                (https) server certificate. Defaults to
                                env[OS_CACERT].
      --os-cert <certificate>   Defaults to env[OS_CERT].
      --os-key <key>            Defaults to env[OS_KEY].
      --timeout <seconds>       Set request timeout (in seconds).

    See "keystone help COMMAND" for help on a specific command.

* 编程调用shell.py的OpenStackIdentityShell类
*bash*::

    >>> dir(python_keystoneclient_shell.OpenStackIdentityShell)
    ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', '_add_bash_completion_subparser', '_find_actions', 'auth_check', 'do_bash_completion', 'do_help', 'get_api_class', 'get_base_parser', 'get_subcommand_parser', 'main']
    >>> myObj = python_keystoneclient_shell.OpenStackIdentityShell()
    >>> myObj.main(['--version'])
    0.11.2.43
    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $

使用pdb调试shell.py
^^^^^^^^^^^^^^^^^^^
*bash* and *pdb*::

    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ python
    Python 2.7.8 (default, Jul 28 2014, 01:34:03)
    [GCC 4.8.3] on cygwin
    Type "help", "copyright", "credits" or "license" for more information.
    >>> import pdb, keystoneclient.shell
    >>> pdb.run('keystoneclient.shell.main()')
    > /usr/lib/python2.7/pdb.py(1238)run()
    -> Pdb().run(statement, globals, locals)

    (Pdb) help step
    s(tep)
    Execute the current line, stop at the first possible occasion
    (either in a function that is called or in the current function).

    (Pdb) help return
    r(eturn)
    Continue execution until the current function returns.
    (Pdb) r
    --Return--
    > /usr/lib/python2.7/pdb.py(104)__init__()->None
    -> self.commands_bnum = None # The breakpoint number for which we are
    (Pdb) r
    > <string>(1)<module>()
    (Pdb) s
    --Call--
    > /home/Administrator/github.com/openstack/python-keystoneclient/keystoneclient/shell.py(461)main()
    -> def main():

    (Pdb) help next
    n(ext)
    Continue execution until the next line in the current function
    is reached or it returns.
    (Pdb) n
    > /home/Administrator/github.com/openstack/python-keystoneclient/keystoneclient/shell.py(462)main()
    -> try:

    (Pdb) l
    457             heading = '%s%s' % (heading[0].upper(), heading[1:])
    458             super(OpenStackHelpFormatter, self).start_section(heading)
    459
    460
    461     def main():
    462  ->     try:
    463             OpenStackIdentityShell().main(sys.argv[1:])
    464
    465         except Exception as e:
    466             print(encodeutils.safe_encode(six.text_type(e)), file=sys.stderr)
    467             sys.exit(1)

    (Pdb) l 315, 396
    315
    316         def main(self, argv):
    317             # Parse args once to find version
    318             parser = self.get_base_parser()
    319             (options, args) = parser.parse_known_args(argv)
    320
    321             # build available subcommands based on version
    322             api_version = options.os_identity_api_version
    323             subcommand_parser = self.get_subcommand_parser(api_version)
    324             self.parser = subcommand_parser
    325
    326             # Handle top-level --help/-h before attempting to parse
    327             # a command off the command line
    328             if not argv or options.help:
    329                 self.do_help(options)
    330                 return 0
    331
    332             # Parse args again and call whatever callback was selected
    333             args = subcommand_parser.parse_args(argv)
    334
    335             # Short-circuit and deal with help command right away.
    336             if args.func == self.do_help:
    337                 self.do_help(args)
    338                 return 0
    339             elif args.func == self.do_bash_completion:
    340                 self.do_bash_completion(args)
    341                 return 0
    342
    343             if args.debug:
    344                 logging_level = logging.DEBUG
    345                 iso_logger = logging.getLogger('iso8601')
    346                 iso_logger.setLevel('WARN')
    347             else:
    348                 logging_level = logging.WARNING
    349
    350             logging.basicConfig(level=logging_level)
    351
    352             # TODO(heckj): supporting backwards compatibility with environment
    353             # variables. To be removed after DEVSTACK is updated, ideally in
    354             # the Grizzly release cycle.
    355             args.os_token = args.os_token or env('SERVICE_TOKEN')
    356             args.os_endpoint = args.os_endpoint or env('SERVICE_ENDPOINT')
    357
    358             if utils.isunauthenticated(args.func):
    359                 self.cs = shell_generic.CLIENT_CLASS(endpoint=args.os_auth_url,
    360                                                      cacert=args.os_cacert,
    361                                                      key=args.os_key,
    362                                                      cert=args.os_cert,
    363                                                      insecure=args.insecure,
    364                                                      timeout=args.timeout)
    365             else:
    366                 self.auth_check(args)
    367                 token = None
    368                 if args.os_token and args.os_endpoint:
    369                     token = args.os_token
    370                 api_version = options.os_identity_api_version
    371                 self.cs = self.get_api_class(api_version)(
    372                     username=args.os_username,
    373                     tenant_name=args.os_tenant_name,
    374                     tenant_id=args.os_tenant_id,
    375                     token=token,
    376                     endpoint=args.os_endpoint,
    377                     password=args.os_password,
    378                     auth_url=args.os_auth_url,
    379                     region_name=args.os_region_name,
    380                     cacert=args.os_cacert,
    381                     key=args.os_key,
    382                     cert=args.os_cert,
    383                     insecure=args.insecure,
    384                     debug=args.debug,
    385                     use_keyring=args.os_cache,
    386                     force_new_token=args.force_new_token,
    387                     stale_duration=args.stale_duration,
    388                     timeout=args.timeout)
    389
    390             try:
    391                 args.func(self.cs, args)
    392             except exc.Unauthorized:
    393                 raise exc.CommandError("Invalid OpenStack Identity credentials.")
    394             except exc.AuthorizationFailure:
    395                 raise exc.CommandError("Unable to authorize user")
    396

    (Pdb) help break
    b(reak) ([file:]lineno | function) [, condition]
    With a line number argument, set a break there in the current
    file.  With a function name, set a break at first executable line
    of that function.  Without argument, list all breaks.  If a second
    argument is present, it is a string specifying an expression
    which must evaluate to true before the breakpoint is honored.

    The line number may be prefixed with a filename and a colon,
    to specify a breakpoint in another file (probably one that
    hasn't been loaded yet).  The file is searched for on sys.path;
    the .py suffix may be omitted.
    (Pdb) b 328
    Breakpoint 1 at /home/Administrator/github.com/openstack/python-keystoneclient/keystoneclient/shell.py:328

    (Pdb) help continue
    c(ont(inue))
    Continue execution, only stop when a breakpoint is encountered.

    (Pdb) options
    Namespace(debug=False, force_new_token=False, help=False, insecure=False, os_auth_url='', os_cacert=None, os_cache=False, os_cert=None, os_endpoint='', os_identity_api_version='', os_key=None, os_password='', os_region_name='', os_tenant_id='', os_tenant_name='', os_token='', os_username='', stale_duration=30, timeout=600)
    (Pdb) args
    self = <keystoneclient.shell.OpenStackIdentityShell object at 0x7f8ffdec>
    argv = []
    ...
    (Pdb) l
    431                     self.subcommands[args.command].print_help()
    432                 else:
    433                     raise exc.CommandError("'%s' is not a valid subcommand" %
    434                                            args.command)
    435             else:
    436  ->             self.parser.print_help()
    437
    438
    439     # I'm picky about my shell help.
    440     class OpenStackHelpFormatter(argparse.HelpFormatter):
    441         INDENT_BEFORE_ARGUMENTS = 6
    (Pdb) n
    > /home/Administrator/github.com/openstack/python-keystoneclient/keystoneclient/shell.py(330)main()
    -> return 0
    (Pdb) n
    --Return--
    > /home/Administrator/github.com/openstack/python-keystoneclient/keystoneclient/shell.py(330)main()->0
    -> return 0
    (Pdb) n
    --Return--
    > /home/Administrator/github.com/openstack/python-keystoneclient/keystoneclient/shell.py(463)main()->None
    -> OpenStackIdentityShell().main(sys.argv[1:])
    (Pdb) n
    --Return--
    > <string>(1)<module>()->None
    (Pdb) n

    >>> quit()
    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $

**关于pdb，详细参考**

    https://docs.python.org/2.7/library/pdb.html
