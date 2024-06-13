package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI(ENV string) string {
	if ENV == "local" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	return os.Getenv("MONGOURI")
}
