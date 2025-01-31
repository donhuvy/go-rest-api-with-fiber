package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config func to get env value from key.
func Config(key string) string {
	// Load .env file.
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}
