package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVairables() {

	// "../.env" - doesnt work with tests
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
