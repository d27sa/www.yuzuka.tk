package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/d27sa/www.yuzuka.tk/model"

	"github.com/d27sa/www.yuzuka.tk/app/chatroom"
)

// RootPath represent the full path of the website root
var RootPath string

// apps stores all available apps
var apps []*model.App

// appChatroom represents a running chatroom app
var appChatroom *chatroom.Chatroom

const (
	// ROOT represents the directory name of the website root
	ROOT = "www.yuzuka.tk"

	// AllowRegister means whether the website allow to register
	AllowRegister = true
)

// init does some initial work
func init() {
	var err error
	RootPath, err = getRootPath(ROOT)
	if err != nil {
		log.Fatalln(err)
	}
	apps = append(apps, model.NewApp(1, "Chatroom", "A simple chatroom.", "chatroom"))
	apps = append(apps, model.NewApp(2, "Translator", "A translator which supports translation between English, Japanese and Chinese.", "translator"))
	appChatroom = chatroom.New()
}

// registerHandlers binds handler functions to specified path
func registerHandlers() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/about/", handleAbout)
	http.HandleFunc("/app/", handleApp)
	http.HandleFunc("/app/chatroom/", handleAppChatroom)
	http.HandleFunc("/app/chatroom/ws", handleAppChatroomWs)
	http.HandleFunc("/app/translator/", handleAppTranslator)
	http.Handle("/static/css/", http.StripPrefix("/static/css", http.FileServer(http.Dir(filepath.Join(RootPath, "static/css")))))
	http.Handle("/static/js/", http.StripPrefix("/static/js", http.FileServer(http.Dir(filepath.Join(RootPath, "static/js")))))
}

// startHTTPServer starts service using http scheme
func startHTTPServer(ch chan<- bool) {
	server := http.Server{
		Addr: ":80",
	}
	log.Println("HTTP server started (listening on port 80).")
	log.Println("HTTP server stopped with error:", server.ListenAndServe())
	ch <- true
}

// startHTTPSServer starts service using https scheme
func startHTTPSServer(ch chan<- bool) {
	server := http.Server{
		Addr: ":443",
	}
	log.Println("HTTPS server started (listening on port 443).")
	log.Println("HTTPS server stopped with error:", server.ListenAndServeTLS(filepath.Join(RootPath, "static/certificate/fullchain.cer"), filepath.Join(RootPath, "static/certificate/www.yuzuka.tk.key")))
	ch <- true
}

// main is the entrance of the program
func main() {
	registerHandlers()

	appChatroom.Run() // run the chatroom app

	ch := make(chan bool) // a channel used to get errors
	defer close(ch)
	go startHTTPServer(ch)
	go startHTTPSServer(ch)
	<-ch
	<-ch
	log.Fatal("Servers stopped with errors.")
}
