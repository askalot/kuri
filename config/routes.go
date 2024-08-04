package config

import (
	"net/http"

	"github.com/askalot/kuri/controllers"
)

const (
	NotesIndex = "/"
)

func SetupRoutes() {
	http.HandleFunc(NotesIndex, controllers.NotesIndex)
}
