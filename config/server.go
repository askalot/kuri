package config

import (
	"fmt"
	"net/http"
	"os"
)

func StartServer() {
	port := ":" + os.Getenv("PORT")

	fmt.Println(fmt.Sprintf("\nListening on http://localhost%s.", port))
	fmt.Println("Use Ctrl-C stop.")

	server := http.Server{
		Addr:    port,
		Handler: Router,
	}

	server.ListenAndServe()
}
