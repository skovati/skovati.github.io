## Linux Terminology - Distros, Package Managers, and More
*Author: Skovati, Date: 2020-12-26*

### Distros

So, say you wanted to install Microsoft Windows on a new computer. You'd go to the official [Windows download](https://www.microsoft.com/en-us/software-download/windows10ISO) page, download that *one* specific ISO, burn it to a USB or DVD, and insert it into your computer. The installation process for Linux is quite similar, except instead of there being one official download, there are *thousands* of different distributions of Linux, know as distros. The only thing each of these distros have in common is that they all use some form of the Linux kernel. Some have different package managers (which we'll get to shortly), different versions of the kernel, different core utilities (like the GNU system I talked about last chapter), different release schedules, or different initialization systems. A common behavior with these distros, since all the code is free and open source, is "forking", where a new distro uses an existing distro as a base. An example of this is the relationship between Debian and Ubuntu, two very popular distros. Since Ubuntu is based on Debian, anything written for Debian will also work out of the box on Ubuntu. This relationship means that you can have interoperability between distros, and it limits fragmentation within the ecosystem, but still offers thousands of choices for the user on which flavor of Linux they want to install. 

Here I will give a small summary of the most popular distros and their derivatives. As previously mentioned, there are thousands of distros, but a dozen or so hold the vast majority of the user base.

- [Debian](https://www.debian.org/): The grand-daddy of distros. Debian 0.01, the first release, came out all the way back in 1993, *before* the first 1.0 Linux kernel release. Because of this long history, Debian is likely the most forked distro of all time. It uses the apt package manager, and has stable releases about every 2 years which use a Long Term Support (LTS) version of the Linux kernel, meaning it's perfect for server applications. This is my personal server distro choice, and is what this blog is hosted on.

    - [Ubuntu](https://ubuntu.com/): The most popular Debian based distro, as well as the most popular Linux distro *period*. Since it is based on Debian, it also uses the apt package manager, but releases a new image every 6 months, and is more lenient with updating new packages, meaning it's a more "fresh" version of Debian. For example, Debian, since it's focused on stability and security, might only have version 80 of Firefox in the repositories with security updates only, but Ubuntu would have the latest Firefox version 84.2, with all new features, meaning it is more sutible for desktop users. Every two years, Ubuntu releases a LTS version, just like Debian. The current LTS is Ubuntu 20.04 LTS, which runs on my Raspberry Pi. Ubuntu is perfect for newcomers as they make it a design goal to be easy to download, install, and use.

    - [Linux Mint](https://linuxmint.com/): Based on Ubuntu, Linux Mint aims to be even *more* user-friendly than Ubuntu, by removing some of Ubuntu's questionable decisions and offering many different desktop environment flavors. This is the distro I recommend for beginners.

    - [Raspberry Pi OS](https://www.raspberrypi.org/software/): Debian based distro meant specifically for Raspberry Pis.

    - [Tails OS](https://tails.boum.org/): Debian based, security and privacy focused distro. Mainly used live on USB sticks to anonymously access the Internet through the [Tor](https://www.torproject.org/) Network. Perfect for journalists, those with government-restricted Internet access, or those who wish to access the Dark Web. Not recommended for daily use or for beginners.

- [Red Hat Enterprise Linux](https://www.redhat.com/en/technologies/linux-platforms/enterprise-linux): Red Hat, an IBM subsidiary, makes an enterprise focused Linux distro known as RHEL. Uses the dnf package manager. Perfect for servers and mission critical software deployments, RHEL is open-source, but you can pay an optional fee for on-demand enterprise-grade tech support.

    - [Fedora](https://getfedora.org/): Technically, RHEL is based on Fedora, but they share most of the same source code. When Red Hat wants to add a new feature, it first enters the Fedora distro as a sort of testing ground for the feature. Because of this, Fedora carries very up to date software and is perfect for desktop use. If a feature is deemed ready for the super stable and security focused RHEL, it gets added to a new version of RHEL a few months down the line. This is my recommended desktop distro for people who are familiar with Linux, but don't want to mess around with configuring some of the more advanced distros below.

    - [CentOS](https://www.centos.org/): Before December 2020, CentOS used to be an exact clone of RHEL, with the branding and paywalls removed, making it perfect for enthusiast home servers. Unfortunately, Red Hat bought out the CentOS project, and has recently turned it into more of a testing ground like Fedora, making it useless for many who wanted a stable server OS without the RHEL license subscription. Currently, one of my home servers runs CentOS 8, but I will be moving it to a different platform because of this news.

    - [Rocky Linux](https://rockylinux.org/): Essentially a clone of the CentOS project since Red Hat shut that one down. Currently has no stable release but I will be transitioning to this platform once they release a 1.0.

- [Arch Linux](https://archlinux.org/): Arch is a rolling release distro, which means as soon as a new version of software is available, like Firefox 85 or Bash 5.1, it will be added to the Arch repositories. Uses the pacman package manager. Arch is considered "bleeding-edge" because of this; you always have the latest version of software on Arch. This is in contrast with the distros above, as those software packages are only updated based on the release schedule of the distro, meaning Ubuntu won't get the latest version of Bash for a few months, and Debian won't get it for years. Arch also doesn't have an installer, meaning the user has to partition disks, make filesystems, bootstrap the package manager, and install all of their desired software manually, making it a much more advanced distro. This is what I currently use on desktop, and I recommend it to any advanced Linux user who craves the ability to customize any and every part of their system.

    - [Manjaro](https://manjaro.org/): Essentially a pre-configured version of Arch, so you don't have to worry about manually installing and getting things configured. Perfect for those who want bleeding-edge software like developers or enthusiasts without the hassle of installing Arch.


*Other more obscure distros worth mentioning*

- [Gentoo](https://www.gentoo.org/): Sourced based distro, meaning every piece of software you install must be manually compiled. A pain in the ass to install and configure, but offers near infinite customizability because of this.

- [Alpine](https://www.alpinelinux.org/): Linux as small and lightweight as possible, only takes up ~100MB when installed, perfect for embedded systems or old hardware.

- [OpenSUSE](https://www.opensuse.org/)

- [Slackware](http://www.slackware.com/)

[Here](https://upload.wikimedia.org/wikipedia/commons/1/1b/Linux_Distribution_Timeline.svg) is a ridiculous look at all the different distros of Linux.

### Package Managers

I've been talking about these package managers like apt, dnf, and pacman, but what actually *are* they? I think the best way to explain it is to compare the software installation process in Linux vs Windows. So, you are likely familiar with how to install something on Windows: you open up Microsoft Edge, suffer through a Bing search to find the Google Chrome installation page, download the .exe file, and run it as an administrator, which installs Chrome on your machine for you. It is *way easier* on Linux. Essentially, each distro has what are called repositories, which hold the files for tens of thousands of different packages of software made specifically for that distro. So, to install chrome on something like Ubuntu, you would just type `sudo apt install google-chrome` in the terminal, apt (the package manager on Debian/Ubuntu) would run its magic, and a few seconds later Chrome would be fully installed on your system, no downloading random .exes from the internet and running them blindly. For what it's worth, you can still manually install things if it happens to not be in your distro's repos.

This style of package management has many pros:

- No downloading random installers from the internet and running them with admin privileges, packages are instead scrutinized and validated by package managers for the distro.

- No wasting time searching the internet and clicking around in a downloads folder for an installer.

- All installed packages can be checked for updates simultaneously with a command like `sudo apt update`, rather than relying on the software itself to have a built in updater, or just running out of date software since you're too lazy to download a new exe.

- Package managers automatically link and install shared dependencies to optimize disk space.

- Because Linux is modular by design and each program does one thing and one things only, system updates are administered through package managers. In contrast, Windows has to release a whole new version and force you to restart just to add a feature to File Manager or Edge. No more "Updating Windows..." screens while you're trying to look something up quick on your laptop. Even the kernel is updated through package managers.

### Desktop Environments and Window Managers

The most obvious form of customization in Linux is how the desktop looks and feels. This is the responsibility of a desktop environment, which manages windows, the dock, status bar, icons, settings menus, login pages, and more. Some popular options include [Gnome](https://www.gnome.org/) and [KDE Plasma](https://kde.org/plasma-desktop/). When you download a distro's ISO, they will likely give you multiple options for "flavors", which include different DEs. So you can install Ubuntu with a Gnome DE, or install Ubuntu with KDE Plasma. Both are still Ubuntu, but will look and feel vastly different because the main visual software is by a different organization. I recommend Gnome as a starting place, but it is *very* easy to install and switch to different ones. A "Window Manager" is essentially *just* the part of a DE that controls the windows. Personally I just use a WM rather than a full DE, because I don't need a fancy dock or a settings menu. I use [bspwm](https://github.com/baskerville/bspwm) and [polybar](https://github.com/polybar/polybar) for a lightweight desktop config.

<img src="https://www.gnome.org/wp-content/uploads/2020/12/wgo-splash-338.png" alt="drawing" width="100%"/>

*Gnome Desktop*

<img src="https://raw.githubusercontent.com/skovati/dotfiles/master/scrot.png" alt="drawing" width="100%"/>

*My bspwm and polybar config*

### Configuration Files aka "Dotfiles"

Another fantastic feature of Linux distros is that all of the configuration is held in so-called "configuration files". Often these files reside in the users home directory (we'll talk about filesystem organization in a later chapter) and are prefixed with a "." like ".vimrc", hence the nickname "dotfiles". You can see examples of some dotfiles in my [GitHub repo](https://github.com/skovati/dotfiles). These hold all of the config for my window manager, text editor, pdf and image viewer, terminal, and more. Because this config is stored in text files, I can back-up and version control them with git, and if I want to reinstall my OS or switch distros, all I have to do it clone that repo and 95% of my configuration is done already—including automatically installing important packages. Try doing that in windows where you have to manually install every piece of software and change each setting by hand with a new install. This makes Linux very portable and replaceable; my laptop's SSD could get wiped and I wouldn't break a sweat, because I didn't really lose anything. You will spend a lot of time customizing these files, but it's time well spent, as they can stay with you during your whole career.

[Chapter 4 - Installing Linux](/linux/intro/ch4)
