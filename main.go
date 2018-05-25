package main

import(
    "log"
    "log/syslog"
    "time"
    "math/rand"
    "fmt"
    "strings"
    "strconv"
    "os"
    "net"
)

type InterfaceStruct struct {
    NAME       string
    IPADDR       string
}

func random(min, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return rand.Intn(max - min) + min
}

func ranip() string {
    //create random first octect between 1 and 255
    //and create random second, third and fourth octets between 0 and 255
    return fmt.Sprintf("%d.%d.%d.%d", random(1, 223), random(0, 255), random(0, 255), random(0, 255))
}

func ranport() string {
    // create port between 1 and 65535
    myport := random(1, 65535)
    return fmt.Sprintf("%d", myport)
}

func getinterfaceinfo() (string, string) {
    var ifname string
    var ifip string
    list, err := net.Interfaces()
    if err != nil {
        panic(err)
    }   

    for i, iface := range list {
            if strings.Contains(iface.Name, "eth") {
                ifname = iface.Name
            addrs, err := iface.Addrs()
            if err != nil {
                panic(err)
            }
                for j, addr := range addrs {
                    ifipraw := strings.Split(fmt.Sprintf("%s", addr), "/")
                    ifip = ifipraw[0]             
                    _ = i
                    _ = j
                }
            }
    }
return ifip, ifname
}

func main() {

    sdis, _ := os.LookupEnv("StartDelayInSeconds")
    sdil, _ := strconv.Atoi(sdis)
    fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())
    fmt.Println("Sleeping...")
    time.Sleep(time.Duration(sdil) * time.Second)
    fmt.Printf("Current Unix Time: %v\n", time.Now().Unix())


    myhostname, err := os.LookupEnv("HOSTNAME")
    if len(myhostname) == 0 {
        log.Fatal(err)
    }

    logwriter, e := syslog.New(syslog.LOG_NOTICE, os.Getenv("MyProgramName"))
    if e == nil {
        log.SetOutput(logwriter)
    }

    interfaceip, interfacename := getinterfaceinfo()

    // Make a channel and then rate limit that channel
    lps, _ := strconv.Atoi(os.Getenv("LogsPerSecond"))
    total, _ := strconv.Atoi(os.Getenv("TotalLogs"))
    interval := time.Duration(1000000000 / lps)
    pumplogs := make(chan int, total)
    for i := 1; i <= total; i++ {
        pumplogs <- i
    }
    close(pumplogs)
    limiter := time.Tick(time.Nanosecond * interval) //interval between transactions

    for mylogs := range pumplogs {
        <-limiter
        log.Print(myhostname," kernel: [0000000.000000] iptables: IN=", interfacename," OUT= MAC=00:00:99:ed:32:00:d0:00:e5:6c:7e:00:00:00 SRC=", ranip()," DST=", interfaceip," LEN=40 TOS=0x08 PREC=0x40 TTL=236 ID=5570 DF PROTO=TCP SPT=", ranport()," DPT=", ranport()," WINDOW=14600 RES=0x00 SYN URGP=0")
        _ = mylogs
    }
}