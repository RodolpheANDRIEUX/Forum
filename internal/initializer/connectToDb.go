package initializer

import (
	"forum/Log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	dsn := os.Getenv("LOGIN_DB")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		Log.Err.Panic("Failed to connect to DB")
	}

	log.Println("connected to ", DB)
}
