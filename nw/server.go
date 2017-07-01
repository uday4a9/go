package main

import (
        "encoding/gob"
//        "io"
        "net"
        "os"
        "fmt"
    )

func server(port string) {
    ln, err := net.Listen("tcp", port) 
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Listener started as", ln)
    for {
        conn, err := ln.Accept()
        if err!=nil{
            fmt.Println(err)
            continue
        }
        fmt.Println(conn)
//        handler(conn)
    }
}

func handler(conn net.Conn) {
    var msg string
    err := gob.NewDecoder(conn).Decode(&msg)
    if err!=nil {
        fmt.Println(err)
    } else {
        fmt.Println("Received", msg)
    }
    conn.Close()
}

func main() {
//    fmt.Println("Args are:", os.Args)
//    fmt.Println(os.Args[1])
    server(os.Args[1])
}
