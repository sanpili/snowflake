package main

import (
        "fmt"
        "net"
        id "github.com/sanpili/algorithm/id"
)

var idgen id.IdGen = &id.Snowflake{ProcessId: 1}

func StartServer(port string) {
        service := ":" + port
        tcpAddr, err := net.ResolveTCPAddr("tcp", service)
        if nil != err {
                fmt.Println(err.Error())
                return
        }
        link, err := net.ListenTCP("tcp", tcpAddr)
        if nil != err {
                fmt.Println(err.Error())
                return
        }
        for {
                fmt.Println("Listening....")
                conn, err := link.Accept()
                if nil != err {
                        fmt.Println(err.Error)
                        continue
                }
                go Handler(conn)
        }
}

func Handler(conn net.Conn) {
        buf := make([]byte, 1024)
        for {
                _, err := conn.Read(buf)
                if err != nil {
                        conn.Close()
                        fmt.Println(err.Error())
                        break
                }
                id := idgen.Next()
                msg := fmt.Sprintf("%d", id)
                conn.Write([]byte(msg))
        }
}

func main() {
        StartServer("9762")
}
