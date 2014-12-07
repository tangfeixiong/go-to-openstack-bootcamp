OpenStack python-keystoneclient开发 - getting started
=====================================================
Pre-requisite
-------------
1. Linux (以下为cygwin模拟环境)
2. Python
3. Git

Community Resources
-------------------
1. Repositories
    * `github.com`_ - https://github.com/openstack/python-keystoneclient
    * openstack.org - http://git.openstack.org/cgit/
2. Bug Report
    * launchpad.net - https://launchpad.net/python-keystoneclient
3. Developer Documentation
    * http://docs.openstack.org/developer/keystone/
    * wiki - https://wiki.openstack.org/wiki/Keystone
    
    .. _github.com: https://github.com/

示例 - git clone 
^^^^^^^^^^^^^^^^
*bash*  
.. code:: 
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel$
    git clone https://github.com/openstack/python-keystoneclient.git


Staging
-------
#. 使用tox将keystone client安装到默认的virtuaenv，后面有tox的较详细的安装和设置章节
*bash*  .. code::
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

#. 切换到tox创建的virtualenv，操作keystone client命令
*bash*
    .. code:: bash
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
#. main函数
*bash*
    .. code:: bash
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

#. keystoneclient包和shell模块
*bash*
    .. code:: bash
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

#. shell.py的main函数和OpenStackIdentityShell类，类方法main()
*bash*
    .. code:: bash
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

#. keystoneclient包的初始元数据
*bash*
    .. code:: bash
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

#. 编程调用shell.py的main函数
*bash*
    .. code:: bash
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

#. 编程调用shell.py的OpenStackIdentityShell类
*bash*
    .. code:: bash
    >>> dir(python_keystoneclient_shell.OpenStackIdentityShell)
    ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', '_add_bash_completion_subparser', '_find_actions', 'auth_check', 'do_bash_completion', 'do_help', 'get_api_class', 'get_base_parser', 'get_subcommand_parser', 'main']
    >>> myObj = python_keystoneclient_shell.OpenStackIdentityShell()
    >>> myObj.main(['--version'])
    0.11.2.43
    (py27)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $

调试shell.py
^^^^^^^^^^^^
*bash* and *pdb*
    .. code::
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

关于tox
-------
# 参考
  https://pypi.python.org/pypi/tox/1.8.1
  
  https://testrun.org/tox/latest/
# 安装（升级）
*bash*
    .. code:: bash
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

#. tox.ini
*bash*
    .. code::
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
    
tox按照tox.ini中的envlist依次测试   
*bash*
    .. code::
    dministrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
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

因此，可以选择目标环境执行tox
*bash*
    .. code:: bash
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
    $ tox -e py27
    ...
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
    $ ls .tox/py27/
    bin  etc  include  lib  log

Pypi Mirror
-----------
以下在virtualenv中安装python-keystoneclient时使用的是pypi服务器是镜像
*bash*
    .. code:: bash
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/keystone-2014.2
    $ cat ~/.pip/pip.conf
    [global]
    #index-url=https://pypi.python.org/simple
    index-url=http://pypi.douban.com/simple
详细参考
    http://www.pypi-mirrors.org/

    https://pypi.python.org/mirrors
    
关于virtualenv
--------------
详细参考请到pypi网站搜索
# 创建virtualenv
*bash*
    .. code:: bash
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel
    $ virtualenv stagingenv/
    Overwriting stagingenv/lib/python2.7/site.py with new content
    New python executable in stagingenv/bin/python2.7
    Not overwriting existing python script stagingenv/bin/python (you must use stagingenv/bin/python2.7)
    Installing setuptools, pip...done.
    Overwriting stagingenv/bin/activate.fish with new content
使用virtualenv
*bash*
    .. code:: bash
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel
    $ source stagingenv/bin/activate
    (stagingenv)
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel
    $ echo $VIRTUAL_ENV
    /home/Administrator/python-workspace/openstack-devel/stagingenv
    (stagingenv)
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel
    $ ls stagingenv/
    bin  include  lib
    (stagingenv)
在virtualenv中安装tox
*bash*
    .. code:：bash
    (stagingenv)
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/
    $ pip install tox
    Requirement already satisfied (use --upgrade to upgrade): tox in /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages
    Requirement already satisfied (use --upgrade to upgrade): virtualenv>=1.9.1 in /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages (from tox)
    Requirement already satisfied (use --upgrade to upgrade): py>=1.4.15 in /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages (from tox)
    Cleaning up...
在virtualenv中安装python-keystoneclient
*bash*
    .. code:: bash
    (stagingenv)
    Administrator@lenovo-9d779749 ~/python-workspace/openstack-devel/
    $ cd ~/github.com/openstack/python-keystoneclient
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ python setup.py install
    running install
    Requirement already satisfied (use --upgrade to upgrade): pbr>=0.6,!=0.7,<1.0 in /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages
    Requirement already satisfied (use --upgrade to upgrade): argparse in /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages
    Downloading/unpacking Babel>=1.3
      http://pypi.douban.com/simple/Babel/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading Babel-1.3.tar.gz (3.4MB): 3.4MB downloaded
      Running setup.py (path:/home/Administrator/python-workspace/openstack-devel/stagingenv/build/Babel/setup.py) egg_info for package Babel

        warning: no previously-included files matching '*' found under directory 'docs/_build'
        warning: no previously-included files matching '*.pyc' found under directory 'tests'
        warning: no previously-included files matching '*.pyo' found under directory 'tests'
    Downloading/unpacking iso8601>=0.1.9
      http://pypi.douban.com/simple/iso8601/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading iso8601-0.1.10.tar.gz
      Running setup.py (path:/home/Administrator/python-workspace/openstack-devel/stagingenv/build/iso8601/setup.py) egg_info for package iso8601

    Downloading/unpacking netaddr>=0.7.12
      http://pypi.douban.com/simple/netaddr/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading netaddr-0.7.12.tar.gz (1.5MB): 1.5MB downloaded
      Running setup.py (path:/home/Administrator/python-workspace/openstack-devel/stagingenv/build/netaddr/setup.py) egg_info for package netaddr

        warning: no previously-included files matching '*.svn*' found anywhere in distribution
        warning: no previously-included files matching '*.git*' found anywhere in distribution
    Downloading/unpacking oslo.config>=1.4.0
      http://pypi.douban.com/simple/oslo.config/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading oslo.config-1.5.0-py2.py3-none-any.whl
    Downloading/unpacking oslo.i18n>=1.0.0
      http://pypi.douban.com/simple/oslo.i18n/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading oslo.i18n-1.1.0-py2.py3-none-any.whl
    Downloading/unpacking oslo.serialization>=1.0.0
      http://pypi.douban.com/simple/oslo.serialization/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading oslo.serialization-1.1.0-py2.py3-none-any.whl
    Downloading/unpacking oslo.utils>=1.0.0
      http://pypi.douban.com/simple/oslo.utils/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading oslo.utils-1.1.0-py2.py3-none-any.whl
    Downloading/unpacking PrettyTable>=0.7,<0.8
      http://pypi.douban.com/simple/PrettyTable/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading prettytable-0.7.2.zip
      Running setup.py (path:/home/Administrator/python-workspace/openstack-devel/stagingenv/build/PrettyTable/setup.py) egg_info for package PrettyTable

    Downloading/unpacking requests>=2.2.0,!=2.4.0
      http://pypi.douban.com/simple/requests/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading requests-2.5.0-py2.py3-none-any.whl (464kB): 464kB downloaded
    Downloading/unpacking six>=1.7.0
      http://pypi.douban.com/simple/six/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading six-1.8.0-py2.py3-none-any.whl
    Downloading/unpacking stevedore>=1.1.0
      http://pypi.douban.com/simple/stevedore/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading stevedore-1.1.0-py2.py3-none-any.whl
    Requirement already satisfied (use --upgrade to upgrade): pip in /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages (from pbr>=0.6,!=0.7,<1.0)
    Downloading/unpacking pytz>=0a (from Babel>=1.3)
      http://pypi.douban.com/simple/pytz/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading pytz-2014.10-py2.py3-none-any.whl (477kB): 477kB downloaded
    Downloading/unpacking netifaces>=0.10.4 (from oslo.utils>=1.0.0)
      http://pypi.douban.com/simple/netifaces/ uses an insecure transport scheme (http). Consider using https if pypi.douban.com has it available
      Downloading netifaces-0.10.4.tar.gz
      Running setup.py (path:/home/Administrator/python-workspace/openstack-devel/stagingenv/build/netifaces/setup.py) egg_info for package netifaces

    Installing collected packages: Babel, iso8601, netaddr, oslo.config, oslo.i18n, oslo.serialization, oslo.utils, PrettyTable, requests, six, stevedore, pytz, netifaces
      Running setup.py install for Babel

        warning: no previously-included files matching '*' found under directory 'docs/_build'
        warning: no previously-included files matching '*.pyc' found under directory 'tests'
        warning: no previously-included files matching '*.pyo' found under directory 'tests'
        Installing pybabel script to /home/Administrator/python-workspace/openstack-devel/stagingenv/bin
      Running setup.py install for iso8601

      Running setup.py install for netaddr
        changing mode of build/scripts-2.7/netaddr from 644 to 755

        warning: no previously-included files matching '*.svn*' found anywhere in distribution
        warning: no previously-included files matching '*.git*' found anywhere in distribution
        changing mode of /home/Administrator/python-workspace/openstack-devel/stagingenv/bin/netaddr to 755
      Found existing installation: oslo.config 1.2.1
        Uninstalling oslo.config:
          Successfully uninstalled oslo.config
      Running setup.py install for PrettyTable

      Found existing installation: six 1.4.1
        Uninstalling six:
          Successfully uninstalled six
      Running setup.py install for netifaces
        checking for getifaddrs...found.
        checking for getnameinfo...found.
        checking for optional header files...none found.
        checking whether struct sockaddr has a length field...no.
        checking which sockaddr_xxx structs are defined...in in6 un.
        checking for routing socket support...no.
        checking for sysctl(CTL_NET...) support...no.
        checking for netlink support...no.
        building 'netifaces' extension
        gcc -fno-strict-aliasing -ggdb -O2 -pipe -Wimplicit-function-declaration -fdebug-prefix-map=/usr/src/ports/python/python-2.7.8-1.i686/build=/usr/src/debug/python-2.7.8-1 -fdebug-prefix-map=/usr/src/ports/python/python-2.7.8-1.i686/src/Python-2.7.8=/usr/src/debug/python-2.7.8-1 -DNDEBUG -g -fwrapv -O3 -Wall -Wstrict-prototypes -DNETIFACES_VERSION=0.10.4 -DHAVE_GETIFADDRS=1 -DHAVE_GETNAMEINFO=1 -DHAVE_SOCKADDR_IN=1 -DHAVE_SOCKADDR_IN6=1 -DHAVE_SOCKADDR_UN=1 -I/usr/include/python2.7 -c netifaces.c -o build/temp.cygwin-1.7.32-i686-2.7/netifaces.o
        netifaces.c: In function ‘gateways’:
        netifaces.c:1213:22: warning: unused variable ‘defaults’ [-Wunused-variable]
           PyObject *result, *defaults;
                              ^
        netifaces.c:2262:3: warning: ‘result’ is used uninitialized in this function [-Wuninitialized]
           return result;
           ^
        gcc -shared -Wl,--enable-auto-image-base -L. build/temp.cygwin-1.7.32-i686-2.7/netifaces.o -L/home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/config -L/usr/lib -lpython2.7 -o build/lib.cygwin-1.7.32-i686-2.7/netifaces.dll

    Successfully installed Babel iso8601 netaddr oslo.config oslo.i18n oslo.serialization oslo.utils PrettyTable requests six stevedore pytz netifaces
    Cleaning up...
    running build
    running build_py
    creating build
    creating build/lib
    creating build/lib/keystoneclient
    creating build/lib/keystoneclient/auth
    copying keystoneclient/auth/base.py -> build/lib/keystoneclient/auth
    copying keystoneclient/auth/cli.py -> build/lib/keystoneclient/auth
    copying keystoneclient/auth/conf.py -> build/lib/keystoneclient/auth
    copying keystoneclient/auth/token_endpoint.py -> build/lib/keystoneclient/auth
    copying keystoneclient/auth/__init__.py -> build/lib/keystoneclient/auth
    creating build/lib/keystoneclient/auth/identity
    creating build/lib/keystoneclient/auth/identity/generic
    copying keystoneclient/auth/identity/generic/base.py -> build/lib/keystoneclient/auth/identity/generic
    copying keystoneclient/auth/identity/generic/password.py -> build/lib/keystoneclient/auth/identity/generic
    copying keystoneclient/auth/identity/generic/token.py -> build/lib/keystoneclient/auth/identity/generic
    copying keystoneclient/auth/identity/generic/__init__.py -> build/lib/keystoneclient/auth/identity/generic
    copying keystoneclient/auth/identity/base.py -> build/lib/keystoneclient/auth/identity
    copying keystoneclient/auth/identity/v2.py -> build/lib/keystoneclient/auth/identity
    copying keystoneclient/auth/identity/v3.py -> build/lib/keystoneclient/auth/identity
    copying keystoneclient/auth/identity/__init__.py -> build/lib/keystoneclient/auth/identity
    creating build/lib/keystoneclient/tests
    creating build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/client_fixtures.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_access.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_auth.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_client.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_discovery.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_ec2.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_endpoints.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_extensions.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_roles.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_services.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_service_catalog.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_shell.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_tenants.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_tokens.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/test_users.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/utils.py -> build/lib/keystoneclient/tests/v2_0
    copying keystoneclient/tests/v2_0/__init__.py -> build/lib/keystoneclient/tests/v2_0
    creating build/lib/keystoneclient/openstack
    creating build/lib/keystoneclient/openstack/common
    copying keystoneclient/openstack/common/memorycache.py -> build/lib/keystoneclient/openstack/common
    copying keystoneclient/openstack/common/uuidutils.py -> build/lib/keystoneclient/openstack/common
    copying keystoneclient/openstack/common/_i18n.py -> build/lib/keystoneclient/openstack/common
    copying keystoneclient/openstack/common/__init__.py -> build/lib/keystoneclient/openstack/common
    creating build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/client_fixtures.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/saml2_fixtures.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_access.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_auth.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_auth_saml2.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_client.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_credentials.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_discover.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_domains.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_endpoints.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_endpoint_filter.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_endpoint_policy.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_federation.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_groups.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_oauth1.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_policies.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_projects.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_regions.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_roles.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_role_assignments.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_services.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_service_catalog.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_tokens.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_trusts.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/test_users.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/utils.py -> build/lib/keystoneclient/tests/v3
    copying keystoneclient/tests/v3/__init__.py -> build/lib/keystoneclient/tests/v3
    creating build/lib/keystoneclient/generic
    copying keystoneclient/generic/client.py -> build/lib/keystoneclient/generic
    copying keystoneclient/generic/shell.py -> build/lib/keystoneclient/generic
    copying keystoneclient/generic/__init__.py -> build/lib/keystoneclient/generic
    creating build/lib/keystoneclient/openstack/common/apiclient
    copying keystoneclient/openstack/common/apiclient/auth.py -> build/lib/keystoneclient/openstack/common/apiclient
    copying keystoneclient/openstack/common/apiclient/base.py -> build/lib/keystoneclient/openstack/common/apiclient
    copying keystoneclient/openstack/common/apiclient/client.py -> build/lib/keystoneclient/openstack/common/apiclient
    copying keystoneclient/openstack/common/apiclient/exceptions.py -> build/lib/keystoneclient/openstack/common/apiclient
    copying keystoneclient/openstack/common/apiclient/fake_client.py -> build/lib/keystoneclient/openstack/common/apiclient
    copying keystoneclient/openstack/common/apiclient/utils.py -> build/lib/keystoneclient/openstack/common/apiclient
    copying keystoneclient/openstack/common/apiclient/__init__.py -> build/lib/keystoneclient/openstack/common/apiclient
    creating build/lib/keystoneclient/contrib
    copying keystoneclient/contrib/__init__.py -> build/lib/keystoneclient/contrib
    creating build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/client.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/ec2.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/endpoints.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/extensions.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/roles.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/services.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/shell.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/tenants.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/tokens.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/users.py -> build/lib/keystoneclient/v2_0
    copying keystoneclient/v2_0/__init__.py -> build/lib/keystoneclient/v2_0
    creating build/lib/keystoneclient/common
    copying keystoneclient/common/cms.py -> build/lib/keystoneclient/common
    copying keystoneclient/common/__init__.py -> build/lib/keystoneclient/common
    creating build/lib/keystoneclient/fixture
    copying keystoneclient/fixture/discovery.py -> build/lib/keystoneclient/fixture
    copying keystoneclient/fixture/exception.py -> build/lib/keystoneclient/fixture
    copying keystoneclient/fixture/v2.py -> build/lib/keystoneclient/fixture
    copying keystoneclient/fixture/v3.py -> build/lib/keystoneclient/fixture
    copying keystoneclient/fixture/__init__.py -> build/lib/keystoneclient/fixture
    creating build/lib/keystoneclient/v3
    creating build/lib/keystoneclient/v3/contrib
    creating build/lib/keystoneclient/v3/contrib/federation
    copying keystoneclient/v3/contrib/federation/base.py -> build/lib/keystoneclient/v3/contrib/federation
    copying keystoneclient/v3/contrib/federation/core.py -> build/lib/keystoneclient/v3/contrib/federation
    copying keystoneclient/v3/contrib/federation/domains.py -> build/lib/keystoneclient/v3/contrib/federation
    copying keystoneclient/v3/contrib/federation/identity_providers.py -> build/lib/keystoneclient/v3/contrib/federation
    copying keystoneclient/v3/contrib/federation/mappings.py -> build/lib/keystoneclient/v3/contrib/federation
    copying keystoneclient/v3/contrib/federation/projects.py -> build/lib/keystoneclient/v3/contrib/federation
    copying keystoneclient/v3/contrib/federation/protocols.py -> build/lib/keystoneclient/v3/contrib/federation
    copying keystoneclient/v3/contrib/federation/__init__.py -> build/lib/keystoneclient/v3/contrib/federation
    creating build/lib/keystoneclient/middleware
    copying keystoneclient/middleware/auth_token.py -> build/lib/keystoneclient/middleware
    copying keystoneclient/middleware/memcache_crypt.py -> build/lib/keystoneclient/middleware
    copying keystoneclient/middleware/s3_token.py -> build/lib/keystoneclient/middleware
    copying keystoneclient/middleware/__init__.py -> build/lib/keystoneclient/middleware
    creating build/lib/keystoneclient/contrib/bootstrap
    copying keystoneclient/contrib/bootstrap/shell.py -> build/lib/keystoneclient/contrib/bootstrap
    copying keystoneclient/contrib/bootstrap/__init__.py -> build/lib/keystoneclient/contrib/bootstrap
    creating build/lib/keystoneclient/contrib/auth
    creating build/lib/keystoneclient/contrib/auth/v3
    copying keystoneclient/contrib/auth/v3/saml2.py -> build/lib/keystoneclient/contrib/auth/v3
    copying keystoneclient/contrib/auth/v3/__init__.py -> build/lib/keystoneclient/contrib/auth/v3
    copying keystoneclient/v3/client.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/credentials.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/domains.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/endpoints.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/groups.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/policies.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/projects.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/regions.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/roles.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/role_assignments.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/services.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/tokens.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/users.py -> build/lib/keystoneclient/v3
    copying keystoneclient/v3/__init__.py -> build/lib/keystoneclient/v3
    creating build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/test_cli.py -> build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/test_conf.py -> build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/test_identity_common.py -> build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/test_identity_v2.py -> build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/test_identity_v3.py -> build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/test_password.py -> build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/test_token.py -> build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/test_token_endpoint.py -> build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/utils.py -> build/lib/keystoneclient/tests/auth
    copying keystoneclient/tests/auth/__init__.py -> build/lib/keystoneclient/tests/auth
    creating build/lib/keystoneclient/apiclient
    copying keystoneclient/apiclient/exceptions.py -> build/lib/keystoneclient/apiclient
    copying keystoneclient/apiclient/__init__.py -> build/lib/keystoneclient/apiclient
    creating build/lib/keystoneclient/v3/contrib/oauth1
    copying keystoneclient/v3/contrib/oauth1/access_tokens.py -> build/lib/keystoneclient/v3/contrib/oauth1
    copying keystoneclient/v3/contrib/oauth1/auth.py -> build/lib/keystoneclient/v3/contrib/oauth1
    copying keystoneclient/v3/contrib/oauth1/consumers.py -> build/lib/keystoneclient/v3/contrib/oauth1
    copying keystoneclient/v3/contrib/oauth1/core.py -> build/lib/keystoneclient/v3/contrib/oauth1
    copying keystoneclient/v3/contrib/oauth1/request_tokens.py -> build/lib/keystoneclient/v3/contrib/oauth1
    copying keystoneclient/v3/contrib/oauth1/utils.py -> build/lib/keystoneclient/v3/contrib/oauth1
    copying keystoneclient/v3/contrib/oauth1/__init__.py -> build/lib/keystoneclient/v3/contrib/oauth1
    creating build/lib/keystoneclient/tests/generic
    copying keystoneclient/tests/generic/test_client.py -> build/lib/keystoneclient/tests/generic
    copying keystoneclient/tests/generic/test_shell.py -> build/lib/keystoneclient/tests/generic
    copying keystoneclient/tests/generic/__init__.py -> build/lib/keystoneclient/tests/generic
    creating build/lib/keystoneclient/contrib/ec2
    copying keystoneclient/contrib/ec2/utils.py -> build/lib/keystoneclient/contrib/ec2
    copying keystoneclient/contrib/ec2/__init__.py -> build/lib/keystoneclient/contrib/ec2
    copying keystoneclient/tests/client_fixtures.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_auth_token_middleware.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_base.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_cms.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_discovery.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_ec2utils.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_fixtures.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_http.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_https.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_keyring.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_memcache_crypt.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_s3_token_middleware.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_session.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_shell.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/test_utils.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/utils.py -> build/lib/keystoneclient/tests
    copying keystoneclient/tests/__init__.py -> build/lib/keystoneclient/tests
    creating build/lib/keystoneclient/contrib/revoke
    copying keystoneclient/contrib/revoke/model.py -> build/lib/keystoneclient/contrib/revoke
    copying keystoneclient/contrib/revoke/__init__.py -> build/lib/keystoneclient/contrib/revoke
    creating build/lib/keystoneclient/tests/apiclient
    copying keystoneclient/tests/apiclient/test_exceptions.py -> build/lib/keystoneclient/tests/apiclient
    copying keystoneclient/tests/apiclient/__init__.py -> build/lib/keystoneclient/tests/apiclient
    copying keystoneclient/contrib/auth/__init__.py -> build/lib/keystoneclient/contrib/auth
    copying keystoneclient/v3/contrib/endpoint_filter.py -> build/lib/keystoneclient/v3/contrib
    copying keystoneclient/v3/contrib/endpoint_policy.py -> build/lib/keystoneclient/v3/contrib
    copying keystoneclient/v3/contrib/trusts.py -> build/lib/keystoneclient/v3/contrib
    copying keystoneclient/v3/contrib/__init__.py -> build/lib/keystoneclient/v3/contrib
    copying keystoneclient/openstack/__init__.py -> build/lib/keystoneclient/openstack
    copying keystoneclient/access.py -> build/lib/keystoneclient
    copying keystoneclient/adapter.py -> build/lib/keystoneclient
    copying keystoneclient/base.py -> build/lib/keystoneclient
    copying keystoneclient/baseclient.py -> build/lib/keystoneclient
    copying keystoneclient/client.py -> build/lib/keystoneclient
    copying keystoneclient/discover.py -> build/lib/keystoneclient
    copying keystoneclient/exceptions.py -> build/lib/keystoneclient
    copying keystoneclient/httpclient.py -> build/lib/keystoneclient
    copying keystoneclient/i18n.py -> build/lib/keystoneclient
    copying keystoneclient/service_catalog.py -> build/lib/keystoneclient
    copying keystoneclient/session.py -> build/lib/keystoneclient
    copying keystoneclient/shell.py -> build/lib/keystoneclient
    copying keystoneclient/utils.py -> build/lib/keystoneclient
    copying keystoneclient/_discover.py -> build/lib/keystoneclient
    copying keystoneclient/__init__.py -> build/lib/keystoneclient
    running egg_info
    writing requirements to python_keystoneclient.egg-info/requires.txt
    writing python_keystoneclient.egg-info/PKG-INFO
    writing top-level names to python_keystoneclient.egg-info/top_level.txt
    writing dependency_links to python_keystoneclient.egg-info/dependency_links.txt
    writing entry points to python_keystoneclient.egg-info/entry_points.txt
    [pbr] Reusing existing SOURCES.txt
    creating build/lib/keystoneclient/tests/v3/examples
    creating build/lib/keystoneclient/tests/v3/examples/xml
    copying keystoneclient/tests/v3/examples/xml/ADFS_RequestSecurityTokenResponse.xml -> build/lib/keystoneclient/tests/v3/examples/xml
    copying keystoneclient/tests/v3/examples/xml/ADFS_fault.xml -> build/lib/keystoneclient/tests/v3/examples/xml
    creating build/lib/keystoneclient/locale
    copying keystoneclient/locale/keystoneclient.pot -> build/lib/keystoneclient/locale
    running install_lib
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    copying build/lib/keystoneclient/access.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    copying build/lib/keystoneclient/adapter.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/apiclient
    copying build/lib/keystoneclient/apiclient/exceptions.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/apiclient
    copying build/lib/keystoneclient/apiclient/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/apiclient
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth
    copying build/lib/keystoneclient/auth/base.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth
    copying build/lib/keystoneclient/auth/cli.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth
    copying build/lib/keystoneclient/auth/conf.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity
    copying build/lib/keystoneclient/auth/identity/base.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/generic
    copying build/lib/keystoneclient/auth/identity/generic/base.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/generic
    copying build/lib/keystoneclient/auth/identity/generic/password.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/generic
    copying build/lib/keystoneclient/auth/identity/generic/token.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/generic
    copying build/lib/keystoneclient/auth/identity/generic/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/generic
    copying build/lib/keystoneclient/auth/identity/v2.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity
    copying build/lib/keystoneclient/auth/identity/v3.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity
    copying build/lib/keystoneclient/auth/identity/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity
    copying build/lib/keystoneclient/auth/token_endpoint.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth
    copying build/lib/keystoneclient/auth/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth
    copying build/lib/keystoneclient/base.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    copying build/lib/keystoneclient/baseclient.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    copying build/lib/keystoneclient/client.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/common
    copying build/lib/keystoneclient/common/cms.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/common
    copying build/lib/keystoneclient/common/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/common
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/auth
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/auth/v3
    copying build/lib/keystoneclient/contrib/auth/v3/saml2.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/auth/v3
    copying build/lib/keystoneclient/contrib/auth/v3/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/auth/v3
    copying build/lib/keystoneclient/contrib/auth/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/auth
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/bootstrap
    copying build/lib/keystoneclient/contrib/bootstrap/shell.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/bootstrap
    copying build/lib/keystoneclient/contrib/bootstrap/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/bootstrap
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/ec2
    copying build/lib/keystoneclient/contrib/ec2/utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/ec2
    copying build/lib/keystoneclient/contrib/ec2/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/ec2
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/revoke
    copying build/lib/keystoneclient/contrib/revoke/model.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/revoke
    copying build/lib/keystoneclient/contrib/revoke/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/revoke
    copying build/lib/keystoneclient/contrib/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib
    copying build/lib/keystoneclient/discover.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    copying build/lib/keystoneclient/exceptions.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture
    copying build/lib/keystoneclient/fixture/discovery.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture
    copying build/lib/keystoneclient/fixture/exception.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture
    copying build/lib/keystoneclient/fixture/v2.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture
    copying build/lib/keystoneclient/fixture/v3.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture
    copying build/lib/keystoneclient/fixture/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/generic
    copying build/lib/keystoneclient/generic/client.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/generic
    copying build/lib/keystoneclient/generic/shell.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/generic
    copying build/lib/keystoneclient/generic/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/generic
    copying build/lib/keystoneclient/httpclient.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    copying build/lib/keystoneclient/i18n.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/locale
    copying build/lib/keystoneclient/locale/keystoneclient.pot -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/locale
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/middleware
    copying build/lib/keystoneclient/middleware/auth_token.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/middleware
    copying build/lib/keystoneclient/middleware/memcache_crypt.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/middleware
    copying build/lib/keystoneclient/middleware/s3_token.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/middleware
    copying build/lib/keystoneclient/middleware/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/middleware
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient
    copying build/lib/keystoneclient/openstack/common/apiclient/auth.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient
    copying build/lib/keystoneclient/openstack/common/apiclient/base.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient
    copying build/lib/keystoneclient/openstack/common/apiclient/client.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient
    copying build/lib/keystoneclient/openstack/common/apiclient/exceptions.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient
    copying build/lib/keystoneclient/openstack/common/apiclient/fake_client.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient
    copying build/lib/keystoneclient/openstack/common/apiclient/utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient
    copying build/lib/keystoneclient/openstack/common/apiclient/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient
    copying build/lib/keystoneclient/openstack/common/memorycache.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common
    copying build/lib/keystoneclient/openstack/common/uuidutils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common
    copying build/lib/keystoneclient/openstack/common/_i18n.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common
    copying build/lib/keystoneclient/openstack/common/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common
    copying build/lib/keystoneclient/openstack/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack
    copying build/lib/keystoneclient/service_catalog.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    copying build/lib/keystoneclient/session.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    copying build/lib/keystoneclient/shell.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/apiclient
    copying build/lib/keystoneclient/tests/apiclient/test_exceptions.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/apiclient
    copying build/lib/keystoneclient/tests/apiclient/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/apiclient
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/test_cli.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/test_conf.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/test_identity_common.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/test_identity_v2.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/test_identity_v3.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/test_password.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/test_token.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/test_token_endpoint.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/auth/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth
    copying build/lib/keystoneclient/tests/client_fixtures.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/generic
    copying build/lib/keystoneclient/tests/generic/test_client.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/generic
    copying build/lib/keystoneclient/tests/generic/test_shell.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/generic
    copying build/lib/keystoneclient/tests/generic/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/generic
    copying build/lib/keystoneclient/tests/test_auth_token_middleware.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_base.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_cms.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_discovery.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_ec2utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_fixtures.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_http.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_https.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_keyring.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_memcache_crypt.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_s3_token_middleware.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_session.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_shell.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/test_utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/tests/utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/client_fixtures.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_access.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_auth.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_client.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_discovery.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_ec2.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_endpoints.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_extensions.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_roles.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_services.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_service_catalog.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_shell.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_tenants.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_tokens.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/test_users.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    copying build/lib/keystoneclient/tests/v2_0/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/client_fixtures.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/examples
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/examples/xml
    copying build/lib/keystoneclient/tests/v3/examples/xml/ADFS_fault.xml -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/examples/xml
    copying build/lib/keystoneclient/tests/v3/examples/xml/ADFS_RequestSecurityTokenResponse.xml -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/examples/xml
    copying build/lib/keystoneclient/tests/v3/saml2_fixtures.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_access.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_auth.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_auth_saml2.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_client.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_credentials.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_discover.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_domains.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_endpoints.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_endpoint_filter.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_endpoint_policy.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_federation.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_groups.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_oauth1.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_policies.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_projects.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_regions.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_roles.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_role_assignments.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_services.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_service_catalog.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_tokens.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_trusts.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/test_users.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/v3/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3
    copying build/lib/keystoneclient/tests/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests
    copying build/lib/keystoneclient/utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/client.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/ec2.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/endpoints.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/extensions.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/roles.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/services.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/shell.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/tenants.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/tokens.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/users.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    copying build/lib/keystoneclient/v2_0/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/client.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib
    copying build/lib/keystoneclient/v3/contrib/endpoint_filter.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib
    copying build/lib/keystoneclient/v3/contrib/endpoint_policy.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation
    copying build/lib/keystoneclient/v3/contrib/federation/base.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation
    copying build/lib/keystoneclient/v3/contrib/federation/core.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation
    copying build/lib/keystoneclient/v3/contrib/federation/domains.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation
    copying build/lib/keystoneclient/v3/contrib/federation/identity_providers.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation
    copying build/lib/keystoneclient/v3/contrib/federation/mappings.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation
    copying build/lib/keystoneclient/v3/contrib/federation/projects.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation
    copying build/lib/keystoneclient/v3/contrib/federation/protocols.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation
    copying build/lib/keystoneclient/v3/contrib/federation/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation
    creating /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1
    copying build/lib/keystoneclient/v3/contrib/oauth1/access_tokens.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1
    copying build/lib/keystoneclient/v3/contrib/oauth1/auth.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1
    copying build/lib/keystoneclient/v3/contrib/oauth1/consumers.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1
    copying build/lib/keystoneclient/v3/contrib/oauth1/core.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1
    copying build/lib/keystoneclient/v3/contrib/oauth1/request_tokens.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1
    copying build/lib/keystoneclient/v3/contrib/oauth1/utils.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1
    copying build/lib/keystoneclient/v3/contrib/oauth1/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1
    copying build/lib/keystoneclient/v3/contrib/trusts.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib
    copying build/lib/keystoneclient/v3/contrib/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib
    copying build/lib/keystoneclient/v3/credentials.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/domains.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/endpoints.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/groups.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/policies.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/projects.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/regions.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/roles.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/role_assignments.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/services.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/tokens.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/users.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/v3/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3
    copying build/lib/keystoneclient/_discover.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    copying build/lib/keystoneclient/__init__.py -> /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/access.py to access.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/adapter.py to adapter.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/apiclient/exceptions.py to exceptions.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/apiclient/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/base.py to base.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/cli.py to cli.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/conf.py to conf.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/base.py to base.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/generic/base.py to base.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/generic/password.py to password.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/generic/token.py to token.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/generic/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/v2.py to v2.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/v3.py to v3.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/identity/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/token_endpoint.py to token_endpoint.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/auth/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/base.py to base.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/baseclient.py to baseclient.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/client.py to client.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/common/cms.py to cms.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/common/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/auth/v3/saml2.py to saml2.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/auth/v3/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/auth/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/bootstrap/shell.py to shell.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/bootstrap/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/ec2/utils.py to utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/ec2/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/revoke/model.py to model.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/revoke/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/contrib/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/discover.py to discover.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/exceptions.py to exceptions.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture/discovery.py to discovery.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture/exception.py to exception.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture/v2.py to v2.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture/v3.py to v3.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/fixture/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/generic/client.py to client.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/generic/shell.py to shell.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/generic/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/httpclient.py to httpclient.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/i18n.py to i18n.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/middleware/auth_token.py to auth_token.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/middleware/memcache_crypt.py to memcache_crypt.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/middleware/s3_token.py to s3_token.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/middleware/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient/auth.py to auth.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient/base.py to base.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient/client.py to client.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient/exceptions.py to exceptions.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient/fake_client.py to fake_client.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient/utils.py to utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/apiclient/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/memorycache.py to memorycache.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/uuidutils.py to uuidutils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/_i18n.py to _i18n.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/common/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/openstack/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/service_catalog.py to service_catalog.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/session.py to session.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/shell.py to shell.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/apiclient/test_exceptions.py to test_exceptions.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/apiclient/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/test_cli.py to test_cli.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/test_conf.py to test_conf.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/test_identity_common.py to test_identity_common.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/test_identity_v2.py to test_identity_v2.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/test_identity_v3.py to test_identity_v3.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/test_password.py to test_password.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/test_token.py to test_token.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/test_token_endpoint.py to test_token_endpoint.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/utils.py to utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/auth/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/client_fixtures.py to client_fixtures.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/generic/test_client.py to test_client.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/generic/test_shell.py to test_shell.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/generic/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_auth_token_middleware.py to test_auth_token_middleware.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_base.py to test_base.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_cms.py to test_cms.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_discovery.py to test_discovery.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_ec2utils.py to test_ec2utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_fixtures.py to test_fixtures.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_http.py to test_http.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_https.py to test_https.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_keyring.py to test_keyring.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_memcache_crypt.py to test_memcache_crypt.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_s3_token_middleware.py to test_s3_token_middleware.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_session.py to test_session.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_shell.py to test_shell.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/test_utils.py to test_utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/utils.py to utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/client_fixtures.py to client_fixtures.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_access.py to test_access.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_auth.py to test_auth.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_client.py to test_client.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_discovery.py to test_discovery.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_ec2.py to test_ec2.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_endpoints.py to test_endpoints.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_extensions.py to test_extensions.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_roles.py to test_roles.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_services.py to test_services.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_service_catalog.py to test_service_catalog.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_shell.py to test_shell.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_tenants.py to test_tenants.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_tokens.py to test_tokens.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/test_users.py to test_users.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/utils.py to utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v2_0/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/client_fixtures.py to client_fixtures.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/saml2_fixtures.py to saml2_fixtures.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_access.py to test_access.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_auth.py to test_auth.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_auth_saml2.py to test_auth_saml2.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_client.py to test_client.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_credentials.py to test_credentials.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_discover.py to test_discover.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_domains.py to test_domains.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_endpoints.py to test_endpoints.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_endpoint_filter.py to test_endpoint_filter.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_endpoint_policy.py to test_endpoint_policy.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_federation.py to test_federation.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_groups.py to test_groups.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_oauth1.py to test_oauth1.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_policies.py to test_policies.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_projects.py to test_projects.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_regions.py to test_regions.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_roles.py to test_roles.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_role_assignments.py to test_role_assignments.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_services.py to test_services.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_service_catalog.py to test_service_catalog.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_tokens.py to test_tokens.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_trusts.py to test_trusts.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/test_users.py to test_users.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/utils.py to utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/v3/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/tests/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/utils.py to utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/client.py to client.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/ec2.py to ec2.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/endpoints.py to endpoints.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/extensions.py to extensions.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/roles.py to roles.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/services.py to services.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/shell.py to shell.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/tenants.py to tenants.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/tokens.py to tokens.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/users.py to users.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v2_0/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/client.py to client.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/endpoint_filter.py to endpoint_filter.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/endpoint_policy.py to endpoint_policy.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation/base.py to base.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation/core.py to core.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation/domains.py to domains.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation/identity_providers.py to identity_providers.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation/mappings.py to mappings.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation/projects.py to projects.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation/protocols.py to protocols.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/federation/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1/access_tokens.py to access_tokens.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1/auth.py to auth.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1/consumers.py to consumers.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1/core.py to core.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1/request_tokens.py to request_tokens.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1/utils.py to utils.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/oauth1/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/trusts.py to trusts.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/contrib/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/credentials.py to credentials.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/domains.py to domains.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/endpoints.py to endpoints.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/groups.py to groups.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/policies.py to policies.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/projects.py to projects.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/regions.py to regions.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/roles.py to roles.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/role_assignments.py to role_assignments.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/services.py to services.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/tokens.py to tokens.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/users.py to users.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/v3/__init__.py to __init__.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/_discover.py to _discover.pyc
    byte-compiling /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/keystoneclient/__init__.py to __init__.pyc
    running install_egg_info
    Copying python_keystoneclient.egg-info to /home/Administrator/python-workspace/openstack-devel/stagingenv/lib/python2.7/site-packages/python_keystoneclient-0.11.2.43.gf8f81bb-py2.7.egg-info
    running install_scripts
    Installing keystone script to /home/Administrator/python-workspace/openstack-devel/stagingenv/bin
    (stagingenv)
    Administrator@lenovo-9d779749 ~/github.com/openstack/python-keystoneclient
    $ which keystone
    /home/Administrator/python-workspace/openstack-devel/stagingenv/bin/keystone




