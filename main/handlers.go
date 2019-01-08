package main

import (
	"fmt"
	"net/http"

	"github.com/d27sa/www.yuzuka.tk/model"
)

// indexHandler handles requests with root path
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates := parseTemplates("layout.html", "home.html")
	path := r.URL.Path[1:]
	writeHead(templates, w, "Home", "layout")
	templates.ExecuteTemplate(w, "layout", fmt.Sprintf("Hello, world!\nThe path is %s.", path))
}

// registerHandler handles requests of register
func registerHandler(w http.ResponseWriter, r *http.Request) {
	templates := parseTemplates("layout.html", "register.html")
	writeHead(templates, w, "Register", "layout")
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
