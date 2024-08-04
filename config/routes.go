package config

import (
	"github.com/go-chi/chi/v5"

	"github.com/askalot/kuri/controllers"
)

var Router = chi.NewRouter()

const (
	NotesIndex = "/"
)

func SetupRoutes() {
	Router.Get(NotesIndex, controllers.NotesIndex)
}
