#!/bin/bash

#export PYTHONPATH=/home/odroid/.local/lib/python2.7/site-packages/elasticsearch:$PYTHONPATH
#export PYTHONPATH="/usr/.local/lib/python2.7/site-packages"
su odroid && sudo taskset -c 0 python /home/odroid/ping.py 172.24.8.100 192.168.11.3
