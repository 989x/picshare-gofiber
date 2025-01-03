package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv checks if the .env file exists and loads it if present
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env file not found or failed to load: %v", err)
	} else {
		log.Println(".env file loaded successfully")
	}
}
