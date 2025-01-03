package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file into environment variables
func LoadEnv() error {
	return godotenv.Load()
}

// MustGetEnv retrieves the value of the specified environment variable
// and logs a fatal error if the variable is not set
func MustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
