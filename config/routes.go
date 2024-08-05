package config

import (
	"net/http"

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
	Router.Use(middleware.RequestID)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)

	staticAssetsDirectory := http.Dir("assets")
	Router.Handle(StaticAssetsPath+"*", http.StripPrefix(StaticAssetsPath, http.FileServer(staticAssetsDirectory)))

	Router.Get(NotesIndex, controllers.NotesIndex)
	Router.Get(NotesNew, controllers.NotesNew)
}
