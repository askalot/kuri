package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/askalot/kuri/config"
)

func init() {
	config.LoadEnvironmentVariables()
	config.SetupRoutes()
}

func startServer() {
	port := ":" + os.Getenv("PORT")

	fmt.Println(fmt.Sprintf("\nListening on http://localhost%s.", port))
	fmt.Println("Use Ctrl-C stop.")

	server := http.Server{
		Addr:    port,
		Handler: config.Router,
	}

	server.ListenAndServe()
}

func main() {
	startServer()
}
