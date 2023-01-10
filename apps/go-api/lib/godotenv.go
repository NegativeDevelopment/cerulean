package lib

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializeDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file")
	}
}
