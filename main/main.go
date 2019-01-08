package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/register", registerHandler)
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("../static/css"))))
	server := http.Server{
		Addr: ":80",
	}
	server.ListenAndServe()
}
