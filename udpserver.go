package main

import (
    "fmt"
    "net"
    "os"
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



/* A Simple function to verify error */
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func main() {
    /* Lets prepare a address at any address at port 10001*/
    ServerAddr,err := net.ResolveUDPAddr("udp",":1234")
    CheckError(err)

    /* Now listen at selected port */
    ServerConn, err := net.ListenUDP("udp", ServerAddr)
    CheckError(err)
    defer ServerConn.Close()

    buf := make([]byte, 1024)

    for {
        n,addr,err := ServerConn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("ERROR: ", err)
        }

        scanner := bufio.NewScanner(strings.NewReader(string(buf[0:n])))
        if err != nil {
            fmt.Println("ERROR: ", err)
        }

        scanner.Scan()
        if err != nil {
            fmt.Println("ERROR: ", err)
        }

        ucl := strings.SplitAfter(scanner.Text(), "]")

        t, err := time.Parse("[02/01/2006 15:04]", ucl[0])
        if err != nil {
            fmt.Println("ERROR: ", err)
        }



        g, _ := json.Marshal(Response{
            Timestamp: strconv.FormatInt(t.Unix(), 10),
            Message: ucl[1],
        })
        if err != nil {
            fmt.Println("ERROR: ", err)
        }

        fmt.Println(string(g), addr)
    }
}