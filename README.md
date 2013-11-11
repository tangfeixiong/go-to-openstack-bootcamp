OpenStack-Grizzly版在Ubuntu-Precise即12.04.3-LTS的安装实验-
==============================

参考[mseknibilel](https://github.com/mseknibilel/OpenStack-Grizzly-Install-Guide/tree/OVS_SingleNode)的安装指南，向作者致谢

准备
----

1. 切换到超级用户root
    sudo su或sudo -i
su是linux命令，Switch User的意思，但当前用户没有执行该命令的权限
sudo命令即Super User Do，使当前用户可以执行root命令

2. 添加Grizzly的apt-get在线安装仓库
apt-get的在线安装仓库配置文件在/etc/apt-get目录下
source.list配置文件是Ubuntu系统的在线更新或升级配置文件
使用国内镜像mirrors.163.com
而第三方的在线更新配置文件可以放置在/etc/apt/source.list.d/目录下
这里将OpenStack的配置文件取名为grizzly.list
可以用vi或vim来创建和编辑
cd /etc/apt/source.list.d/
vi grizzly.list
按insert键
输入deb http://ubuntu-cloud.archive.canonical.com/ubuntu precise-updates/grizzly main
按escape键
按wq键
mseknibilel的文章中采用了重定向stdout到grizzly.list文件的写法：
echo deb http://ubuntu-cloud.archive.canonical.com/ubuntu precise-updates/grizzly main >> /etc/apt/sources.list.d/grizzly.list

3. 添加在Ubuntu Precise上安装OpenStack必要的keyring和python等工具
    apt-get install ubuntu-cloud-keyring python-software-properties software-properties-common python-keyring

4. 更新Ubuntu
    apt-get update
    apt-get upgrade
    apt-get dist-upgrade
