package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("AEMET_API_KEY")
	env := os.Getenv("ENV")

	fmt.Printf("AEMET_API_KEY: %s\n", apiKey)
	fmt.Printf("ENV: %s\n", env)
}
