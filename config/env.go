package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvironmentVariables() {
	const environmentConfigFile = ".env"

	err := godotenv.Load(environmentConfigFile)
	if err != nil {
		log.Fatalf("Error loading %s file.\n%s", environmentConfigFile, err)
	}
}
