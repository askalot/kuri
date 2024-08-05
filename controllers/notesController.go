package controllers

import (
	"net/http"
	"text/template"

	"github.com/askalot/kuri/models"
)

const (
	applicationLayout   = "views/layouts/application.html"
	notesViewsDirectory = "views/notes/"
)

func NotesIndex(w http.ResponseWriter, r *http.Request) {
	notes := []models.Note{
		{
			Title:       "First note",
			Description: "Note Uno Details",
		},
		{
			Title:       "Second note",
			Description: "Note Dos Details",
		},
		{
			Title:       "Third note",
			Description: "Note Tres Details",
		},
	}

	tmpl := template.Must(template.ParseFiles(applicationLayout, notesViewsDirectory+"index.html"))

	tmpl.Execute(w, struct {
		Notes []models.Note
	}{
		Notes: notes,
	})
}

func NotesNew(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(applicationLayout, notesViewsDirectory+"new.html"))

	tmpl.Execute(w, nil)
}
