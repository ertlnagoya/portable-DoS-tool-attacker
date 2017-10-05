Open-source Software-based Portable DoS Test Tool for IoT Devices -Attacker-    
====
This DoS tool consists of the _monitor_ and _the attacker_. 

The attacker performs a Denial-of-Service (DoS) attack on the target device, 
and transmits its traffic information to the monitor. 

The monitor tool is available [here](https://github.com/ertlnagoya/portable-DoS-tool-monitor).    
    
<img width="350" alt="2017-09-25 10 22 39" src="https://user-images.githubusercontent.com/26764885/30840032-02bdea00-a2b0-11e7-82ee-2e580704a730.png">    

## Description
### What is "Open-source Software-based Portable DoS Test Tool for IoT Device?"
We implemented a portable denial of service (DoS) test tool based on the *Mirai* malware and conducted a DoS test on several IoT devices. 
The tool visualizes the load, and adjusts the volume of the attack packet manually. 
These functionalities enable visual checks of the state of a DoS attack. 
We can also change the attack method and its performance in different target devices. 
By applying the tool, we can check the vulnerability of an IoT device and understand the state of the actual DoS attack.

## Directories
### attak_bandth   
Attack bandwidth acquisition program.
The attacker reports the amount of sent packet to the monitor.    

This program can be started automatically using a systemctl service `report-packet-stats.service`. 

### bin   
There's two executable `./bot` and `./cnc`.

Using `/etc/rc.local` may help you to start these program.

### bot   
Mirai's attack bot program.
Originates from [Mirai BotNet sorcecode](https://github.com/jgamblin/Mirai-Source-Code).
This is source code of `bin/mirai`
 　　　
### build.sh   
A build script for mirai/cnc.    

### camera   
Reports the camera image acquired with the web camera to the monitor.  

### cnc   
Mirai's command & control server. 
This is source code of `bin/cnc`

### ping   
This python script referred [Python ping package](https://pypi.python.org/pypi/ping).    

Pings from the attacker to the attack target, and reports its response time to the monitor.　

## Hardware & Software Requirements
- odroid-c2    
    - ubuntu 16.04 LTS Xenial Xerus
         - python 2.7
         - gcc    
         - golang    
         - [go-shellwords](https://github.com/mattn/go-shellwords)   
         - [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)    
         - electric-fence   
         - mysql-server    
         - mysql-client   
- WiFi Dongle
- ODROID USB-CAM 720P

## Usage
### To Start Attacker program
```bash
### Terminal 1
./cnc
### Terminal 2
taskset -c 1,2,3 ./mirai
### Terminal 3
./report_packet_stats [moritor IP]
### Terminal 4
python webcam_client_odroid.py [moritor IP]
### Terminal 5
python ping.py [target IP] [moritor IP]
```

Or use a systemctl service and `/etc/rc.local`.   

## Installation
### odroid-c2 setup
Get image file [here](https://odroid.in/ubuntu_16.04lts/ubuntu64-16.04.2lts-mate-odroid-c2-20170301.img.xz).    
Write image in MacOS.
```
unxz ubuntu64-16.04.2lts-mate-odroid-c2-20170301.img.xz
df
diskutil unmount /dev/disk2s1
sudo dd bs=1m if=ubuntu64-16.04.2lts-mate-odroid-c2-20170301.img of=/dev/rdisk2
```
Edit `boot.ini` for mysql-server.
```
setenv mesontimer "0"  # setenv mesontimer "1"
```
### Ubuntu setup  
```bash
sudo apt-get update    
sudo apt-get upgrade    
sudo apt-get install golang    
wget https://bootstrap.pypa.io/get-pip.py    
sudo python get-pip.py     
sudo apt-get install mysql-server mysql-client    
```
### Github setup
```bash
sudo apt-get install git
git config --global user.name "<ユーザ名>"
git config --global user.email "<メールアドレス>"
```
### Golang library setup
```bash
git clone https://github.com/mattn/go-shellwords.git
cd /usr/lib/go-1.6/src/     
sudo mkdir github.com    
sudo mkdir github.com/mattn    
cd github.com/mattn/    
sudo mv ~/go-shellwords/ ./    

git clone https://github.com/go-sql-driver/mysql.git
cd /usr/lib/go-1.6/src/
sudo mkdir github.com/go-sql-driver    
cd github.com/go-sql-driver/    
sudo mv ~/mysql/ ./     
```

### Attacker program Install
```bash
git clone git@github.com:ertlnagoya/portable-DoS-tool-attacker.git
cd portable-DoS-tool-attacker
bash build.sh
```

### mysql
```bash
mysql -u root -p 
```
Exect `db.sql`.    
And, add user.
```mysql
INSERT INTO users VALUES (NULL, 'root', 'password', 0, 0, 0, 0, -1, 1, 30, '');
```

## Licence
[MIT](https://github.com/ertlnagoya/portable-DoS-tool-monitor/blob/master/LICENSE)

## Author
* [NGR](https://github.com/KeigoNagara)    
* [K-atc](https://github.com/K-atc)    

## Disclaimer
This repository is for academic purposes, the use of this software is your responsibility.

## Warning
The file for this repository is being identified by some AV programs as malware. Please take caution. 
