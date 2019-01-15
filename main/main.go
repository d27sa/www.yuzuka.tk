package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

// RootPath represent the full path of the website root
var RootPath string

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
		log.Fatal(err)
	}
}

func registerHandlers() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/about/", handleAbout)
	http.HandleFunc("/app/", handleApp)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir(filepath.Join(RootPath, "static")))))
}

func main() {
	registerHandlers()
	server := http.Server{
		Addr: ":443",
	}

	fmt.Println("Listening at port 443 ...")
	// log.Fatal(server.ListenAndServe())
	log.Fatal(server.ListenAndServeTLS(filepath.Join(RootPath, "static/certificate/fullchain.cer"), filepath.Join(RootPath, "static/certificate/www.yuzuka.tk.key")))
}
