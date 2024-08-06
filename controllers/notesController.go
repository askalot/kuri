package controllers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/askalot/kuri/models"
	"github.com/askalot/kuri/store"
)

const (
	headerPartial       = "views/partials/siteHeader.html"
	applicationLayout   = "views/layouts/application.html"
	notesViewsDirectory = "views/notes/"
)

var noteStore = store.NewNoteStore([]models.Note{
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
})

func NotesIndex(w http.ResponseWriter, r *http.Request) {
	successParam := r.URL.Query().Get("success")
	success, err := strconv.ParseBool(successParam)
	if err != nil {
		success = false
	}

	tmpl := template.Must(template.ParseFiles(applicationLayout, headerPartial, notesViewsDirectory+"index.html"))

	tmpl.Execute(w, struct {
		Notes   []models.Note
		Success bool
	}{
		Notes:   noteStore.AllNotes(),
		Success: success,
	})
}

func NotesNew(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(applicationLayout, headerPartial, notesViewsDirectory+"new.html"))

	tmpl.Execute(w, nil)
}

func NotesCreate(w http.ResponseWriter, r *http.Request) {
	note := models.Note{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	noteStore.CreateNote(note)

	http.Redirect(w, r, "/?success=true", http.StatusSeeOther)
}
