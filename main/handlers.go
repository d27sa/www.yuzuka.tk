package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/d27sa/www.yuzuka.tk/model"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handleIndex handles requests with root path
func handleIndex(w http.ResponseWriter, r *http.Request) {
	templates := parseTemplates("layout.html", "home.html")
	writeHead(templates, w, "Home", "layout")
	templates.ExecuteTemplate(w, "layout", fmt.Sprintf("Hello, world!\nThe path is %s.", r.URL.Path))
	writeScript(templates, w)
}

// handleRegister handles requests of register
func handleRegister(w http.ResponseWriter, r *http.Request) {
	templates := parseTemplates("layout.html", "register.html")
	writeHead(templates, w, "Register", "layout", "register")
	if r.Method == "POST" {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		user := model.User{Username: username, Email: email, Password: password}
		templates.ExecuteTemplate(w, "layout", &user)
		return
	}
	templates.ExecuteTemplate(w, "layout", nil)
	writeScript(templates, w)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	templates := parseTemplates("layout.html", "about.html")
	writeHead(templates, w, "About", "layout", "about")
	templates.ExecuteTemplate(w, "layout", nil)
	writeScript(templates, w)
}

func handleApp(w http.ResponseWriter, r *http.Request) {
	apps := make([]model.App, 6)
	for i := 0; i < 6; i++ {
		apps[i] = model.NewApp(i+1, "Hello", "wow <em>wow</em> wow what's this?!", "#")
	}
	templates := parseTemplates("layout.html", "app.html")
	writeHead(templates, w, "APP", "layout", "app")
	templates.ExecuteTemplate(w, "layout", apps)
	writeScript(templates, w, "app")
}

func handleAppChatroom(w http.ResponseWriter, r *http.Request) {
	templates := parseTemplates("layout.html", "app/chatroom.html")
	writeHead(templates, w, "Chatroom", "layout")
	templates.ExecuteTemplate(w, "layout", nil)
	writeScript(templates, w, "chatroom")
}

func handleAppChatroomWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	msgType, p, err := conn.ReadMessage()
	if err != nil {
		log.Println(err)
	}
	conn.WriteMessage(msgType, p)
}
