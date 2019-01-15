package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handleTCPConn(conn *net.TCPConn) error {
	defer conn.Close()
	data := make([]byte, 4096)
	for {
		n, err := conn.Read(data)
		if err != nil && err != io.EOF {
			return err
		}
		fmt.Println("Received message: ", string(data[:n]))
		_, err = conn.Write(data[:n])
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listening on port 8080 ...")
	var conn *net.TCPConn
	for {
		conn, err = listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("New Conn Accepted: ", conn.RemoteAddr())
		go handleTCPConn(conn)
	}
}
