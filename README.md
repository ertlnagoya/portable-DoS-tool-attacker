Open-source Software-based Portable DoS Test Tool for IoT Devices -Attacker-    
====
The tool comprises a monitor and an attacker.     
The portable DoS tool attacker performs a DoS attack on the target device and transmits the traffic information to the monitor.   
The portable DoS tool monitor is [here](https://github.com/ertlnagoya/portable-DoS-tool-monitor).    
    
<img width="350" alt="2017-09-25 10 22 39" src="https://user-images.githubusercontent.com/26764885/30840032-02bdea00-a2b0-11e7-82ee-2e580704a730.png">    

### attak_bandth   
Attack bandwidth acquisition program.    
An attacker reports the acquired packet amount to the monitor.    
Auto start as service by systemctl.    
### bin   
bot & cnc executable program.    
"/etc/rc.local" to start automatically.    
### bot   
[Mirai BotNet sorcecode](https://github.com/jgamblin/Mirai-Source-Code)    
Mirai's attack bot program.    　　　
### build.sh   
Build script.    
### camera   
Report the camera image acquired with the web camera to the monitor.    
### cnc   
Mirai's command & control server.    
### ping   
[Python ping package](https://pypi.python.org/pypi/ping)    
Send the packet from the attacker to the attack target and report its response time to the monitor.　

## Description
### Open-source Software-based Portable DoS Test Tool for IoT Device    
We constructed a portable denial of service (DoS) test tool based on the malware *Mirai* and conducted a DoS test on several IoT devices. 
The tool can visualize the load and adjust the volume of the attack packet. 
These functionalities enable visual checks of the state of a DoS attack. 
We can also change the attack method and performance in different target devices. 
By applying the tool, we can check the vulnerability of an IoT device and understand the state of the actual DoS attack.     

## Requirement
- odroid-c2    
    - ubuntu 16.04 LTS Xenial Xerus
         - python 2.7
         - gcc    
         - golang    
         - [go-shellwords](https://github.com/mattn/go-shellwords)    
         - electric-fence   
         - mysql-server    
         - mysql-client   
- WiFi Module (odroid)
- ODROID USB-CAM 720P
## Usage
### Attacker program Start    
```
./cnc
taskset -c 1,2,3 ./mirai
./report_packet_stats [moritor IP]
python webcam_client_odroid.py [moritor IP]
python ping.py [target IP] [moritor IP]
```
Or use service by systemctl and "/etc/rc.local".    
## Install    
### Attacker program Install
```
git clone git@github.com:ertlnagoya/portable-DoS-tool-attacker.git
```  
### mysql
```
mysql -u root -p 
db.sql
```
## Licence
[MIT](https://github.com/ertlnagoya/portable-DoS-tool-monitor/blob/master/LICENSE)
## Author
[NGR](https://github.com/KeigoNagara)    
[K-atc](https://github.com/K-atc)    

## Disclaimer
This repository is for academic purposes, the use of this software is your responsibility.

## Warning
The file for this repo is being identified by some AV programs as malware.  Please take caution. 


