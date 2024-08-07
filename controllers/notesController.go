package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"

	"github.com/askalot/kuri/internal/fileutils"
	"github.com/askalot/kuri/models"
	"github.com/askalot/kuri/store"
)

const (
	headerPartial       = "views/partials/siteHeader.html"
	applicationLayout   = "views/layouts/application.html"
	notesViewsDirectory = "views/notes/"
)

var noteStore = store.NewNoteStore([]models.Note{})

func init() {
	dataDirectory := "./data"
	allowedExtensions := []string{".md"}

	fileNames, err := fileutils.ListFileNames(dataDirectory)
	if err != nil {
		fmt.Println(err)
	}

	for _, fileName := range fileNames {
		if fileutils.HasAllowedExtension(fileName, allowedExtensions) {
			content, err := fileutils.GetFileContentString(filepath.Join(dataDirectory, fileName))
			if err != nil {
				fmt.Println(err)
			}

			noteStore.CreateNote(models.Note{
				Title:       fileName,
				Description: content,
			})
		}
	}
}

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
