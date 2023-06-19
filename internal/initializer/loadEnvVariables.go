package initializer

import (
	"forum/Log"
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		Log.Err.Printf("Error loading .env file")
	}
	log.Println(".env file loaded properly")
}
