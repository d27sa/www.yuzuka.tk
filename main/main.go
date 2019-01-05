package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	fmt.Fprintf(w, "Hello, world!\nThe path is %s.", path)
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	http.ListenAndServe(":80", mux)
}
