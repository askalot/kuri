package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/askalot/kuri/config"
)

func init() {
	config.LoadEnvironmentVariables()
	config.SetupStaticAssets()
}

func startServer() {
	port := ":" + os.Getenv("PORT")

	fmt.Println(fmt.Sprintf("\nListening on http://localhost%s.", port))
	fmt.Println("Use Ctrl-C stop.")

	http.ListenAndServe(port, nil)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	startServer()
}
