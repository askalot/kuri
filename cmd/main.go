package main

import (
	"github.com/askalot/kuri/config"
)

func init() {
	config.LoadEnvironmentVariables()
	config.SetupRoutes()
}

func main() {
	config.StartServer()
}
