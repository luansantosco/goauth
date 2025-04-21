package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ .env file not found, attempting to use environment variables")
	}
}

func GetEnv(key, fallback string) string {
	value := os.Getenv((key))
	if value == "" {
		return fallback
	}
	return value
}
