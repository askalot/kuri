package controllers

import (
	"net/http"
	"text/template"
)

const (
	notesViewsDirectory = "views/notes/"
)

func NotesIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(notesViewsDirectory + "index.html"))

	tmpl.Execute(w, nil)
}

func NotesNew(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(notesViewsDirectory + "new.html"))

	tmpl.Execute(w, nil)
}
