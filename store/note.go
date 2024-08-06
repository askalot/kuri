package store

import (
	"errors"

	"github.com/askalot/kuri/models"
)

var noteNotFoundError = errors.New("Note not found.")

type NoteStore struct {
	notes []models.Note
}

func (store *NoteStore) AllNotes() []models.Note {
	return store.notes
}

func NewNoteStore(notes []models.Note) *NoteStore {
	return &NoteStore{
		notes: notes,
	}
}

func (store *NoteStore) CreateNote(note models.Note) {
	store.notes = append(store.notes, note)
}

func (store *NoteStore) FindNote(id string) (models.Note, error) {
	for _, note := range store.notes {
		if note.ID == id {
			return note, nil
		}
	}

	return models.Note{}, noteNotFoundError
}

func (store *NoteStore) Update(id string, newNote models.Note) error {
	for index, note := range store.notes {
		if note.ID == id {
			store.notes[index].Title = newNote.Title
			store.notes[index].Description = newNote.Description

			return nil
		}
	}

	return noteNotFoundError
}

func (store *NoteStore) Destroy(id string) error {
	for index, note := range store.notes {
		if note.ID == id {
			store.notes = append(store.notes[:index], store.notes[index+1:]...)

			return nil
		}
	}

	return noteNotFoundError
}
