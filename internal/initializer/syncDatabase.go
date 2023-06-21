package initializer

import (
	"fmt"
	"forum/Log"
	"forum/internal/models"
	"log"
)

func SyncDatabase() {
	log.Println("Syncing User model with database")
	fmt.Println("Syncing User model with database")
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		Log.Err.Println("Failed to sync the User table:", err)
	}

	DB.Raw("")

	log.Println("Syncing Notifications model with database")
	fmt.Println("Syncing Notifications model with database")
	err = DB.AutoMigrate(&models.Notifications{})
	if err != nil {
		Log.Err.Println("Failed to sync the Notifications table:", err)
	}

	log.Println("Syncing Post model with database")
	fmt.Println("Syncing Post model with database")
	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		Log.Err.Println("Failed to sync the Post table:", err)
	}

	log.Println("Syncing Reply model with database")
	fmt.Println("Syncing Reply model with database")
	err = DB.AutoMigrate(&models.Reply{})
	if err != nil {
		Log.Err.Println("Failed to sync the Reply table:", err)
	}

	log.Println("Syncing PostLike model with database")
	fmt.Println("Syncing PostLike model with database")
	err = DB.AutoMigrate(&models.PostLike{})
	if err != nil {
		Log.Err.Println("Failed to sync the PostLike table:", err)
	}

	log.Println("Syncing ReplyLike model with database")
	fmt.Println("Syncing ReplyLike model with database")
	err = DB.AutoMigrate(&models.ReplyLike{})
	if err != nil {
		Log.Err.Println("Failed to sync the ReplyLike table:", err)
	}

	log.Println("Database sync complete")
	fmt.Println("Database sync complete")
}
