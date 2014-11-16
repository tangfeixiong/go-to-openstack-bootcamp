Manually Setup Python 2.7.8 Developer Environment
=================================================
Preparation
-----------
A CentOS 6.5 VM is used as lab. VM can be launched with any Hypervisor (such as VMware, VirtualBox, virsh).

.. image:: /os/image/launch-centos-inst.png

Using easy_install with setuptools from ez_setup.py
----------------------------------------------------------
the default python installation in CentOS instance.

.. image:: /os/image/instance-default-python.png 

1. install setuptools
    * `setuptools in PYPI <https://pypi.python.org/pypi/setuptools>`_
    * download ``ez_setup.py`` with wget command
        .. image:: /os/image/ez_setup.py-dl.png
    * run ``ez_setup.py`` to download and install setuptools
        .. image:: /os/image/ez_setup.setuptools.png
    * verify ``easy_install`` in setuptools
        .. image:: /os/image/easy_install.verify.png
2. install pip
    * install `pip` with easy_install
        .. image:: /os/image/easy_install-pip.png
    * pip search packages
        .. image:: /os/image/pip-search-example.png

Install python 2.7.8 from source
--------------------------------
1. setup python 2.7.8
    * go to download site
        .. image:: /os/image/python-dl.png
    * python source
        .. image:: /os/image/python-2.7.8.src.link.png
    * download
        .. image:: /os/image/python-2.7.8.src-wget.png
    * extract
        .. image:: /os/image/python-2.7.8.tar-zxf.png
    * before install, do setup CentOS core devel (akka gcc) env
        .. image:: /os/image/yum-core-devel-install.png
    * update CentOS
        .. image:: /os/image/yum-update.png
    * show help for configure tool
        .. image:: /os/image/do-configure-help.png
    * specify location for python installation
        .. image:: /os/image/do-configure-prefix.png
    * generate makefile by configure tool
        .. image:: /os/image/do-configure-makefile.png
    * make and install
        .. image:: /os/image/make-and-install.png
    * execute python
        .. image:: /os/image/execute-python-2.7.8.png

2. install pip in alternative way
    https://pip.pypa.io/en/latest/installing.html

    * download ``get-pip.py``
        .. image:: /os/image/get-pip-wget.png
    * investigate issues
        #) missing zlib
            .. image:: /os/image/get-pip-missing-zlib.png
        #) make and make install again
            .. image:: /os/image/get-pip-zip-re-make.png
            note: the screenshot shows incorrect operation. the correct is following

            .. code:: bash
            [root@host-172-16-32-173 ~]# sed -i 's/#zlib/zlib/g' Python-2.7.8/Modules/Setup

            [root@host-172-16-32-173 ~]# cat Python-2.7.8/Modules/Setup | grep zlib

            # Andrew Kuchling's zlib module.

            # This require zlib 1.1.3 (or later).

            # See http://www.gzip.org/zlib/

            zlib zlibmodule.c -I$(prefix)/include -L$(exec_prefix)/lib -lz

        #) missing openssl
            .. image:: /os/image/get-pip-miss-httpshandler.png

            install openssl-devel

            .. image:: /os/image/get-pip-yum-openssl-devel.png
        #) run get-pip.py correctly
            ``make`` and ``make install`` again

            .. image:: /os/image/get-pip.png

3. about ``virtualenv``
        .. code:: bash
        [root@host-172-16-32-173 ~]# pip install virtualenv

        Downloading/unpacking virtualenv

          Downloading virtualenv-1.11.6-py2.py3-none-any.whl (1.6MB): 1.6MB downloaded

        Installing collected packages: virtualenv

        Successfully installed virtualenv

        Cleaning up...
    
    * ``virtualenv`` 
        .. image:: /os/image/virtualenv-activate.png
    * install ``python-novaclient`` in the virtual environment
        .. image:: /os/image/virtualenv-install-novaclient.png

        .. image:: /os/image/virtualenv-install-novaclient1.png

4. fully installation of developer environment
    .. code:: bash
(stagingenv)[root@host-172-16-32-173 ~]# yum install python-devel

