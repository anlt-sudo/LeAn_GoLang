package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil{
		log.Println("Error: While get .env file!")
	}
}

func GetEnv(key, fallback string) string{
	value := os.Getenv(key)

	if value == ""{
		return fallback
	}

	return value
}