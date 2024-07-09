package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv checks if a .env file exists and loads the environment variables from it
func LoadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Println("Error loading .env file, continuing without it")
		}
	}
}

// LoadApiKey retrieves the AEMET_API_KEY from the environment variables
func LoadApiKey() string {
	LoadEnv() // Load the .env file if it exists
	apiKey := os.Getenv("AEMET_API_KEY")
	if apiKey == "" {
		log.Fatal("AEMET_API_KEY environment variable not set")
	}
	return apiKey
}
