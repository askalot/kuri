package controllers

import (
	"fmt"
	"net/http"
	"strconv"
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

	successParam := r.URL.Query().Get("success")
	success, err := strconv.ParseBool(successParam)
	if err != nil {
		success = false
	}

	tmpl := template.Must(template.ParseFiles(applicationLayout, notesViewsDirectory+"index.html"))

	tmpl.Execute(w, struct {
		Notes   []models.Note
		Success bool
	}{
		Notes:   notes,
		Success: success,
	})
}

func NotesNew(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(applicationLayout, notesViewsDirectory+"new.html"))

	tmpl.Execute(w, nil)
}

func NotesCreate(w http.ResponseWriter, r *http.Request) {
	note := models.Note{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	fmt.Println(note)

	http.Redirect(w, r, "/?success=true", http.StatusSeeOther)
}
