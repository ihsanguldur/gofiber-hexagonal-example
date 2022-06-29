package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Config(key string) string {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("Can not load .env file.")
	}

	return os.Getenv(key)
}
