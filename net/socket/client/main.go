package main

import (
	"fmt"
	"io/ioutil"
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

	_, err = conn.Write([]byte("Hello!\n"))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(resp))
	// _, err = conn.Write([]byte("Hello!\n"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
