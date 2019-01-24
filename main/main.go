package main

// codes with "temp" in the comments should be removed later

import (
	"log"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"

	"github.com/d27sa/www.yuzuka.tk/model"

	"github.com/d27sa/www.yuzuka.tk/app/chatroom"
)

var (
	// rootPath represent the full path of the website root
	rootPath string
	// apps stores all available apps
	apps []*model.App
	// appChatroom represents a running chatroom app
	appChatroom *chatroom.Chatroom
	// verificationCode maps email address to verification code used for registeration
	verificationCode map[string]string
	// serverMailers is the mailers used to send mail to users
	serverMailers *mailers
)

const (
	// root represents the directory name of the website root
	root = "www.yuzuka.tk"
)

// init does some initial work
func init() {
	rand.Seed(time.Now().UnixNano())           // set the random seed
	verificationCode = make(map[string]string) // store verification code for each mail address
	initMailers()                              // initialize the mailers used to send verification mails
	// get the root path of the website
	var err error
	rootPath, err = getRootPath(root)
	if err != nil {
		log.Fatalln(err)
	}
	initApps()     // initialize apps
	initPostList() // temp - this is only for test, remember to remove it
}

// registerHandlers binds handler functions to specified path
func registerHandlers() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/register/ajax/register", handleRegisterAjaxRegister)
	http.HandleFunc("/register/ajax/vericode", handleRegisterAjaxVericode)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/login/ajax", handleLoginAjax)
	http.HandleFunc("/blog/", handleBlog)
	http.HandleFunc("/about/", handleAbout)
	http.HandleFunc("/app/", handleApp)
	http.HandleFunc("/app/chatroom/", handleAppChatroom)
	http.HandleFunc("/app/chatroom/ws", handleAppChatroomWs)
	http.HandleFunc("/app/translator/", handleAppTranslator)
	http.HandleFunc("/app/translator/ajax", handleAppTranslatorAjax)
	http.Handle("/static/css/", http.StripPrefix("/static/css", http.FileServer(http.Dir(filepath.Join(rootPath, "static/css")))))
	http.Handle("/static/js/", http.StripPrefix("/static/js", http.FileServer(http.Dir(filepath.Join(rootPath, "static/js")))))
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
	log.Println("HTTPS server stopped with error:", server.ListenAndServeTLS(filepath.Join(rootPath, "static/certificate/fullchain.cer"), filepath.Join(rootPath, "static/certificate/www.yuzuka.tk.key")))
	ch <- true
}

// main is the entrance of the program
func main() {
	registerHandlers()
	appChatroom.Run() // run the chatroom app
	// start the server
	ch := make(chan bool) // a channel used to get errors
	defer close(ch)
	go startHTTPServer(ch)
	go startHTTPSServer(ch)
	<-ch
	<-ch
	log.Fatal("Servers stopped with errors.")
}
