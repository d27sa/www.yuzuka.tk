package main

import (
	"log"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "www.yuzuka.tk:8080")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	conn.Write([]byte("Hello, world!\n"))
}
