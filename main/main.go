package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/d27sa/www.yuzuka.tk/model"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")
	templates := template.Must(template.ParseFiles("../templates/layout.html", "../templates/home.html"))
	path := r.URL.Path[1:]
	templates.ExecuteTemplate(w, "layout", fmt.Sprintf("Hello, world!\nThe path is %s.", path))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseFiles("../templates/layout.html", "../templates/register.html"))
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

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/register", registerHandler)
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("../static/css"))))
	server := http.Server{
		Addr: ":80",
	}
	server.ListenAndServe()
}
