package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	templates, _ := template.ParseFiles("templates/home.html")
	path := r.URL.Path[1:]
	templates.ExecuteTemplate(w, "home", fmt.Sprintf("Hello, world!\nThe path is %s.", path))
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	http.ListenAndServe(":80", mux)
}
