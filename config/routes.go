package config

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/askalot/kuri/controllers"
)

var Router = chi.NewRouter()

const (
	NotesIndex = "/"
	NotesNew   = "/new"
)

func SetupRoutes() {
	Router.Use(middleware.Logger)

	Router.Get(NotesIndex, controllers.NotesIndex)
	Router.Get(NotesNew, controllers.NotesNew)
}
