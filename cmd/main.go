package main

import (
	"fmt"
	"net/http"
)

func startServer() {
	PORT := ":3000"

	fmt.Println(fmt.Sprintf("\nListening on http://localhost%s.", PORT))
	fmt.Println("Use Ctrl-C stop.")

	http.ListenAndServe(PORT, nil)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	startServer()
}
