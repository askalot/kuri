package config

import (
	"fmt"
	"net/http"
)

const (
	NotesIndex = "/"
)

func SetupRoutes() {
	http.HandleFunc(NotesIndex, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})
}
