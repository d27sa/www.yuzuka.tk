package model

import "html/template"

// App represent a preview of an application
type App struct {
	ID    int
	Name  string
	Intro template.HTML
	URL   string
}

// NewApp returns a new instance of App
func NewApp(id int, name, intro, url string) *App {
	return &App{
		ID:    id,
		Name:  name,
		Intro: template.HTML(intro),
		URL:   url,
	}
}
