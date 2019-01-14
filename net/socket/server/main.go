package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		data := make([]byte, 1024)
		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(data[:n]))
			_, err = conn.Write(data[:n])
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}
