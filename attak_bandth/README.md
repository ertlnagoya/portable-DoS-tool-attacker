systemdで攻撃ホストのパケット送信量を報告するやつ
====

必要物
----
* g++
* report\_packet\_stats.c
* recieve\_packet\_stats.py
    * 時刻補正機能付き
* report-packet-stats.service
    * systemdのサービスとしてreport\_packet\_statsを起動する
* correct\_time\_server.py
    * 時刻補正サーバー（モニターホストに設置を想定）
/home/katc/lab/MiraiTower/report_packet_stats 

準備
----
```bash
g++ -o report_packet_stats report_packet_stats.c
cp report-packet-stats.service ~/.config/systemd/user/
systemctl --user daemon-reload
### startの代わりにrestartするとプログラムを再読込してくれる
systemctl --user start  report-packet-stats.service
```

* report-packet-stats.service に書かれているIPアドレスはモニター用のホストアドレスを書くこと
* report-packet-stats.service に書かれているプログラムのパスは適切なものに書き換えること
* サービスは自動起動するように設定したほうが楽

### Advanced
* モニターでcorrect\_time\_server.pyをsystemdで自動起動するようにする（既存の`.service`ファイルをいじって所定の場所に設置するだけ）


使い方
----
ローカルホストで検証するとき

```bash
### in Terminal I
./recieve_packet_stats.py
### in Terminal II (Optional)
systemctl --user start  report-packet-stats.service
### in Terminal III (Optional)
journalctl --user -f
### in Terminal IV (Optional)
./correct_time_server.py
```

データフォーマット
----
```
タイムスタンプ（マイクロ秒単位`%.6f`） パケット数\n
… 以降繰り返し ...
```

例：
```
% ./recieve_packet_stats.py
1490366281.597008 147740
1490366282.097463 147742
1490366282.598024 147745
1490366283.098695 147745
1490366283.599037 147745
1490366284.099294 147751
1490366284.599828 147757
1490366285.100270 147757
1490366285.600719 147764
```

TODO
----
* [ ] モニターホストを自動で識別できる手立てを用意する（IPアドレス設定の自動化）