package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	LoadEnvVairables()
	ConnectToDB()
}

func LoadEnvVairables() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
