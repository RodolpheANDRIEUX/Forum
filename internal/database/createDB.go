package database

import (
	"forum/internal/models"
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
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}

	//créer un User pour tester
	userThomas := models.User{
		Username: "Thomas",
		Email:    "mail@mail.fr",
		Password: "password",
	}

	//mets les données dans la db en pointant ce qu'on ajoute
	db.Create(&userThomas)

}
