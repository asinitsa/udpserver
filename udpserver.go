package main

import (
    "fmt"
    "net"
    "bufio"
    "strings"
    "time"
    "encoding/json"
    "strconv"
)

type Response struct {
    Timestamp string `json:timestamp`
    Message string `json:message`
}

func JsonMarshal(tstp string, msg string) string {

    g, err := json.Marshal(Response{
        Timestamp: tstp,
        Message: msg,
    })
    if err != nil {
        fmt.Println("ERROR: ", err)
    }

    jsonstr := string(g)

    return jsonstr
}

func TimeParseToUnix(timestr string) string {

    t, err := time.Parse("[02/01/2006 15:04]", timestr)
    if err != nil {
        fmt.Println("ERROR: ", err)
    }

    unixtimestamp := strconv.FormatInt(t.Unix(), 10)

    return unixtimestamp
}

func ScanPacket( buf []byte, n int) []string {

    scanner := bufio.NewScanner(strings.NewReader(string(buf[0:n])))
    if err := scanner.Err(); err != nil {
        fmt.Println("ERROR: ", err)
    }

    scanner.Scan()
    if err := scanner.Err(); err != nil {
        fmt.Println("ERROR: ", err)
    }

    fields := strings.SplitAfter(scanner.Text(), "]")
    if err := scanner.Err(); err != nil {
        fmt.Println("ERROR: ", err)
    }

    return fields
}

func main() {

    buf := make([]byte, 65536) /* size of UDP single UDP datagram */

    /* Lets prepare a address at any address at port 10001*/
    ServerAddr,err := net.ResolveUDPAddr("udp",":1234")
    if err != nil {
        fmt.Println("ERROR: ", err)
    }

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    if err != nil {
        fmt.Println("ERROR: ", err)
    }
    defer ServerConn.Close()

    for {
        n,_,err := ServerConn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("ERROR: ", err)
        }

        fields := ScanPacket(buf, n)
        timestamp := TimeParseToUnix(fields[0])
        fmt.Println(JsonMarshal(timestamp,fields[1]))
    }
}