package initializer

import (
	"forum/Log"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		Log.Err.Panic("Error loading .env file")
	}
}
