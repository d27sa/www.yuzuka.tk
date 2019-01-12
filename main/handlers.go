package main

import (
	"fmt"
	"net/http"

	"github.com/d27sa/www.yuzuka.tk/model"
)

// handleIndex handles requests with root path
func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in handleIndex")
	templates := parseTemplates("layout.html", "home.html")
	writeHead(templates, w, "Home", "layout")
	templates.ExecuteTemplate(w, "layout", fmt.Sprintf("Hello, world!\nThe path is %s.", r.URL.Path))
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
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	templates := parseTemplates("layout.html", "about.html")
	writeHead(templates, w, "About", "layout", "about")
	templates.ExecuteTemplate(w, "layout", nil)
}
