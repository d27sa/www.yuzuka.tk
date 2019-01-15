package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "www.yuzuka.tk:8080")
	if err != nil {
		log.Fatal(err)
	}
	var conn *net.TCPConn
	resp := make([]byte, 4096)
	for {
		conn, err = net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}
		_, err = conn.Write([]byte("Hello! This is a message.\n"))
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}
		n, err := conn.Read(resp)
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}
		fmt.Println("The server responds: ", string(resp[:n]))
		time.Sleep(time.Second * 2)
	}
}
