
__Refer__

+ https://medium.com/@remy.hosseinkhan/installing-ubuntu-19-10-on-lenovo-thinkpad-x1-extreme-gen-2-and-p1-gen-2-ba4c9c7c7ed2
+ https://www.linuxbabe.com/ubuntu/install-nvidia-driver-ubuntu-18-04
+ https://linuxconfig.org/how-to-disable-nouveau-nvidia-driver-on-ubuntu-18-04-bionic-beaver-linux

__Ops__

create a bootable USB Stick for Ubuntu

- https://www.pendrivelinux.com/universal-usb-installer-easy-as-1-2-3/
- https://mirrors.tuna.tsinghua.edu.cn/ubuntu-releases/18.04/

Setup Windows 10. Disable Windows 10 fast startup

Update UEFI/BIOS to 1.26 (initially my UEFI/BIOS was up to v1.17)

Disable Secure Boot

Install Ubuntu 

Reboot the computer. In the GRUB menu press E to edit the command before booting then append ‘nomodeset’ at the end of the ‘linux…’ line 

Thus, install last NVIDIA proprietary driver

Blacklist Nouveau driver to avoid conflicts

- https://linuxconfig.org/how-to-disable-nouveau-nvidia-driver-on-ubuntu-18-04-bionic-beaver-linux
