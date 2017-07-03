# portable-DoS-tool-attacker   
## attak_bandth   
攻撃帯域取得プログラム   
アタッカーホストにて取得したパケット量をモニターに報告   
systemctlによりserviceとして自動起動   
## bin   
bot, cncの実行プログラム   
   
/etc/rc．local により自動起動 
## bot   
Miraiの攻撃ボットによるプログラム   
## camera   
ウェブカメラ似て取得したカメラ画像をモニターに報告   
systemctlによりserviceとして自動起動 
## cnc   
Miraiのコマンド＆コントロールサーバによるプログラム   
## ping   
アタッカーから攻撃対象にパケットを送りその応答時間をモニターに報告 
### ToDo   
* [ ] systemctlによりserviceとして自動起動 
## build.sh   
ビルド用スクロプト   
"bash build.sh debug telnet"
