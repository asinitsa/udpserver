package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
    "strings"
)

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
            fmt.Println("Error: ",err)
        }

        scanner := bufio.NewScanner(strings.NewReader(string(buf[0:n])))
        if err := scanner.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "error:", err)
            os.Exit(1)
        }

        scanner.Scan()

        ucl := strings.SplitAfter(scanner.Text(), "]")

        fmt.Println(ucl, addr)
    }
}