package main

import (
	"fmt"
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

	fmt.Println("Writing...")
	_, err = conn.Write([]byte("Hello!\n"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Written.")
	resp := make([]byte, 1024)
	n, err := conn.Read(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp[:n]))
	_, err = conn.Write([]byte("Hello!\n"))
	if err != nil {
		log.Fatal(err)
	}

}
