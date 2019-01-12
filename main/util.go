package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// head represents the head part of a html5 file
type head struct {
	Title string
	CSS   []string
}

func writeHead(t *template.Template, w http.ResponseWriter, title string, css ...string) {
	t.ExecuteTemplate(w, "head", head{Title: title, CSS: css})
}

func parseTemplates(filenames ...string) *template.Template {
	for i := range filenames {
		filenames[i] = filepath.Join(RootPath, "templates", filenames[i])
	}
	return template.Must(template.ParseFiles(filenames...))
}
