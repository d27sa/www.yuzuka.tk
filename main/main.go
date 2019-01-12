package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// RootPath represent the full path of the website root
var RootPath string

// ROOT represents the directory name of the website root
const ROOT = "www.yuzuka.tk"

func init() {
	wd, _ := os.Getwd()
	RootPath = wd[:(strings.Index(wd, ROOT) + len(ROOT))]
	fmt.Println(RootPath)
}

func registerHandlers() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/about/", handleAbout)
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir(filepath.Join(RootPath, "static/css")))))
}

func main() {
	registerHandlers()
	server := http.Server{
		Addr: ":80",
	}
	fmt.Println("Listening at port 80 ...")
	log.Fatal(server.ListenAndServe())
}
