package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/d27sa/www.yuzuka.tk/app/chatroom"
)

// RootPath represent the full path of the website root
var RootPath string

var appChatroom *chatroom.Chatroom

const (
	// ROOT represents the directory name of the website root
	ROOT = "www.yuzuka.tk"

	// AllowRegister means whether the website allow to register
	AllowRegister = true
)

func init() {
	var err error
	RootPath, err = getRootPath(ROOT)
	if err != nil {
		log.Fatalln(err)
	}
	appChatroom = chatroom.New()
}

func registerHandlers() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/about/", handleAbout)
	http.HandleFunc("/app/", handleApp)
	http.HandleFunc("/app/chatroom/", handleAppChatroom)
	http.HandleFunc("/app/chatroom/ws", handleAppChatroomWs)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(filepath.Join(RootPath, "static")))))
}

func startHTTPServer(ch chan<- bool) {
	server := http.Server{
		Addr: ":80",
	}
	log.Println("HTTP server started (listening on port 80).")
	log.Println("HTTP server stopped with error:", server.ListenAndServe())
	ch <- true
}

func startHTTPSServer(ch chan<- bool) {
	server := http.Server{
		Addr: ":443",
	}
	log.Println("HTTPS server started (listening on port 443).")
	log.Println("HTTPS server stopped with error:", server.ListenAndServeTLS(filepath.Join(RootPath, "static/certificate/fullchain.cer"), filepath.Join(RootPath, "static/certificate/www.yuzuka.tk.key")))
	ch <- true
}

func main() {
	appChatroom.Run()
	registerHandlers()
	ch := make(chan bool)
	defer close(ch)
	go startHTTPServer(ch)
	go startHTTPSServer(ch)
	<-ch
	<-ch
	log.Fatal("Servers stopped with errors.")
}
