package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	templates, err := template.ParseFiles("../templates/home.html")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	path := r.URL.Path[1:]
	templates.ExecuteTemplate(w, "home.html", fmt.Sprintf("Hello, world!\nThe path is %s.", path))
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	server := http.Server{
		Addr:    "0.0.0.0:80",
		Handler: mux,
	}
	server.ListenAndServe()
}
