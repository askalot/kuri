package controllers

import (
	"net/http"
	"text/template"
)

const (
	applicationLayout   = "views/layouts/application.html"
	notesViewsDirectory = "views/notes/"
)

func NotesIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(applicationLayout, notesViewsDirectory+"index.html"))

	tmpl.Execute(w, nil)
}

func NotesNew(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(applicationLayout, notesViewsDirectory+"new.html"))

	tmpl.Execute(w, nil)
}
