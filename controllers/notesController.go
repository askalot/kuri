package controllers

import (
	"fmt"
	"net/http"
)

func NotesIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
