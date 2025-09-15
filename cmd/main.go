package main

import (
	"fmt"
	"log"

	"github.com/askalot/kuri/config"
	"github.com/askalot/kuri/internal/htmlutils"
)

func init() {
	config.LoadEnvironmentVariables()
	config.SetupRoutes()
}

func main() {
	// config.StartServer()
	title, err := htmlutils.GetHTMLTagTextFromURL("title", "https://nimal.info")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(title)
}
