package main

import (
	"errors"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/d27sa/www.yuzuka.tk/app/chatroom"
	"github.com/d27sa/www.yuzuka.tk/model"
)

// head represents the head part of a html5 file
type head struct {
	Title string
	CSS   []string
}

type registerInfo struct {
	Username string
	Email    string
	Password string
	Vericode string
}

type loginInfo struct {
	Username string
	Password string
}

func getRootPath(root string) (string, error) {
	wd, _ := os.Getwd()
	i := strings.Index(wd, root)
	if i == -1 {
		return "", errors.New("failed to locate the root directory")
	}
	return wd[:i+len(root)], nil
}

func writeHead(t *template.Template, w http.ResponseWriter, title string, css ...string) error {
	return t.ExecuteTemplate(w, "head", head{Title: title, CSS: css})
}

func writeScript(t *template.Template, w http.ResponseWriter, js ...string) error {
	return t.ExecuteTemplate(w, "script", js)
}

func parseTemplates(filenames ...string) *template.Template {
	for i := range filenames {
		filenames[i] = filepath.Join(rootPath, "templates", filenames[i])
	}
	return template.Must(template.ParseFiles(filenames...))
}

func initApps() {
	apps = append(apps, model.NewApp(1, "Chatroom", "A simple chatroom.", "chatroom"))
	apps = append(apps, model.NewApp(2, "Translator", "A translator which supports translation between English, Japanese and Chinese.", "translator"))
	appChatroom = chatroom.New()
}
