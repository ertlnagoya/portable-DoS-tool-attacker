Open-source Software-based Portable DoS Test Tool for IoT Devices -Attacker-    
====
The tool comprises a monitor and an attacker.     
The portable DoS tool attacker performs a DoS attack on the target device and transmits the traffic information to the monitor.   
The portable DoS tool monitor is [here](https://github.com/ertlnagoya/portable-DoS-tool-monitor).    
    
<img width="350" alt="2017-09-25 10 22 39" src="https://user-images.githubusercontent.com/26764885/30840032-02bdea00-a2b0-11e7-82ee-2e580704a730.png">    

### attak_bandth   
攻撃帯域取得プログラム   
アタッカーホストにて取得したパケット量をモニターに報告   
systemctlによりserviceとして自動起動   
### bin   
bot, cncの実行プログラム      
"/etc/rc.local" により自動起動 
### bot   
Miraiの攻撃ボットによるプログラム   
### camera   
ウェブカメラ似て取得したカメラ画像をモニターに報告   
systemctlによりserviceとして自動起動 
### cnc   
Miraiのコマンド＆コントロールサーバによるプログラム   
### ping   
アタッカーから攻撃対象にパケットを送りその応答時間をモニターに報告 
### build.sh   
ビルド用スクロプト   
"bash build.sh debug telnet"

## Description
### Open-source Software-based Portable DoS Test Tool for IoT Device    
We constructed a portable denial of service (DoS) test tool based on the malware *Mirai* and conducted a DoS test on several IoT devices. 
The tool can visualize the load and adjust the volume of the attack packet. 
These functionalities enable visual checks of the state of a DoS attack. 
We can also change the attack method and performance in different target devices. 
By applying the tool, we can check the vulnerability of an IoT device and understand the state of the actual DoS attack.     

## Requirement
- ubuntu 16.04 LTS Xenial Xerus
     - python 2.7
     -    
## Usage
### Attacker program Start    
```
./cnc
./mirai.dbg
./report_packet_stats
python webcam_client_odroid.py 
python ping.py
```
## Install    
### Attacker program Install
```
git clone git@github.com:ertlnagoya/portable-DoS-tool-attacker.git
```
## Licence
[MIT](https://github.com/ertlnagoya/portable-DoS-tool-monitor/blob/master/LICENSE)
## Author
[Nagara](https://github.com/KeigoNagara)    
[K-atc](https://github.com/K-atc)    

  
