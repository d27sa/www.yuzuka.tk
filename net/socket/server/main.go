package main

import (
	"fmt"
	"io/ioutil"
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
		data, err := ioutil.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(data))
		_, err = conn.Write([]byte("Hello~\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}
