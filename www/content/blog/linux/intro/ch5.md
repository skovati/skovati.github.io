---
title: "Intro to Linux - Chapter 5"
date: 2021-05-25T18:00:51-05:00
draft: true
---

## Chapter 5 - Server Configuration
*Author: Skovati, Date: 2020-12-29*

This chapter is optional if you're strictly using Linux as a desktop OS, but goes over some useful ideas like user management and security that are still applicable.

### SSH Keys

First of all, what *is* SSH? It stands for Secure Shell, and is a protocol design to gain remote shell access to Linux servers. All traffic is encrypted by default, and the popular OpenSSH implementation of the protocol is one of the most secure pieces of software that exists. You can authenticate through SSH in a few ways, one of which being just your username and password on the server, like `ssh username@server-ip`, which will prompt you for your user's password on that server, and then give you access if it was correct. This is relatively insecure, since lots of cloud based servers will accept SSH requests from any IP, meaning anyone can try and guess your password to gain access to your server. To mitigate this, I recommend the use of SSH public/private key pairs for authentication. If you aren't familiar with public key cryptography—you should be, here's a very brief primer. 

Using advanced number theory, an algorithm generates two keys—one public and one private—that are cryptographically linked, such that using your private key to login to a server that has your public key on a whitelist with authenticate you and essentially guarantee you are who you say you are—that is, if you're responsible with your private key and keep it *private*. Guessing someone else's private key takes somewhere on the order of 2^256 guesses, which is significantly more secure than using an 8 digit password to authenticate. I recommend [reading](https://www.cloudflare.com/learning/ssl/how-does-public-key-encryption-work/) and [watching](https://www.youtube.com/watch?v=GSIDS_lvRv4) some more resources to get a good understanding. Essentially all you need to know to use SSH keys, is that a really long, really secure password is generated for you and stored in a file, and you'll use that "password" (private key) to login to your server. The server, which will have your public key in an "authorized keys" list, will then check for a cryptographic match between SSH requests with different private keys to see if it matches a key in this authorized list. If it does, you are given access.

Because of this vastly increased security, myself, and cloud VM providers like Google Cloud, *only* recommend the use of SSH keys for authentication, and encourage disabling password authentication entirely to increase security. We'll learn how to do that after we generate some SSH keys.

On both Linux and Windows, to generate SSH key pairs, run the command `ssh-keygen`, which will launch an interactive generation program. Selecting the defaults here is alright until you learn a little more, but I'd recommend using a 4096 bit RSA key for increased security, as shorter keys like 1024 bit have shown weaknesses. Once generated, your public key will live at either /home/*username*/.ssh/id_rsa.pub or C:\\Users\\*username*\\.ssh\\id_rsa.pub on Linux and Windows respectively. Your private key will be the file there without the .pub extension. If you're following the GCP guide from Chapter 4, go back and read the next part to add it to your VMs whitelist. If you're setting up a Linux server manually, edit your users ~/.ssh/authorized_keys file and paste your public key there. 

---

### User Management and the Root User

User management on Linux is fantastically simple. Each user has a home directory located at /home/*username*/ (the first / refers to the root of the filesystem, similar to C:/ in Windows), where all of the user-specific configuration files, media, downloads, cached Spotify music, and much more. Users can belong to groups too; each user is automatically a member of a group with the same name as the username. Files and directories can be "owned" by one user, and one group at a time, and each file and directory has certain permissions that say which user or group can read, edit, or execute the given file or directory. We'll learn about filesystems and permissions more in a later chapter. The most important group is the sudo group (also known as the wheel group), which allows any user in the sudo group to use the sudo command.

For those who aren't familiar, the root user on Linux is like the Administrator. The root user runs all the low level services, installs system-wide packages, makes important configuration changes, and can access any information on the system. Because of this extreme power, it isn't recommended to login and use the system as the root user for long periods of time, as you can easily break things since you have no limitations. But, often a regular user will need root access in order to say, install a new package like Firefox, or add a new VPN connection. To remedy this, we use the `sudo` command, which when prefixed to another command, will run that other command as the root user. This is the reason you must run `sudo apt install firefox` rather than just `apt install firefox`, since you are installing a system-wide package that can affect other users. Only certain users, which could be considered administrators, will have the ability to use this `sudo` command. These administrator users are put in the sudo group, so that they can run commands as root. 

This whole user system will make more sense as you gain experience using Linux, but make sure your user is in the sudo group and is able to run commands as root. On GCP, I believe your user is automatically in the sudo group, but doesn't have a password. If you want to add a password, run the command `sudo su` (su is a command that lets you switch to be a different user, when run with no arguments it switches you to the root user) and then run `passwd your-username`, which will prompt you for a new password. Then  type `exit` to stop being the root user. Test your sudo access by running something like `sudo apt update`. If your root access checks out, we can move onto the next steps of configuration.

---

### Hardening your server with SSH configurations and a Firewall

Time to edit our first configuration file! Once you're sure your SSH key auth is working, it's time to disable password-based logins and remote root login in order to secure our server. The configuration file for the SSH daemon (a daemon is like a background service) is located at /etc/ssh/sshd_config. Open this with `sudo nano /etc/ssh/sshd_config`. Scroll down using the arrow keys and find the commented option called PermitRootLogin, uncomment this, and set it to "no". Next find the option PubkeyAuthentication and set it to yes. Finally find the option, PasswordAuthentication and set it to no. Then, exit and save the file by pressing Ctrl-X, Y, and then enter. Finally, we must restart the SSH daemon with the command `sudo systemctl restart sshd` (systemctl is the command used to interact with system processes and daemons). Congrats! You just edited your first system configuration file and applied it. I know that was kind of a lot, but taking these few steps just made your system significantly more secure.

If you're running a server with a public IP address, I recommend running a firewall on it. Normally, a networks router or dedicated firewall will block unauthorized packets, but having a per-server firewall is recommended. If you're running a GCP VM, the firewall can be found under VPC Network and firewall in the nav bar to the left. I recommend ufw as a firewall—it's super simple to setup. Run `sudo apt install ufw` to install the firewall, and then run `sudo ufw allow ssh` to allow our ssh requests through the firewall. Then, run `sudo ufw default deny` so that the firewall denies any traffic we don't explicitly allow. Finally, run `sudo ufw enable` to start the firewall. 

A lot of that might've gone over your head, but that's fine. We secured your system and made it essentially impenetrable while you learn the ropes. Keep in mind that you'll have to change these things if you want to, say, give a new user ssh or sudo access, or let certain traffic through the firewall if you want to run a web server.

Congratulations though! If you've been following Chapters 4-5, you've set up a remote Linux virtual machine or bare metal install, given yourself encrypted access, locked down the traffic with a firewall, and now you have a perfect environment set up to learn the ins and outs of the command line. The next few chapters will likely be about specific topics, like the bash programming language or how to set up a web server, so feel free to skip around to the stuff that interests you!

[Chapter 6 - Bash && "Extensibilty"](/linux/intro/ch6)

