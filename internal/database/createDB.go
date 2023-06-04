package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() {

	//Connect to database
	dsn := "host=mel.db.elephantsql.com user=mavowpjc password=eKB5WjbmSuXHKljByVRZeNQ6yemAI21u dbname=mavowpjc sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Create User table
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}

	userThomas := User{
		Username: "Thomas",
		Email:    "mail@mail.fr",
		Password: "password",
	}

	db.Create(&userThomas)

}
