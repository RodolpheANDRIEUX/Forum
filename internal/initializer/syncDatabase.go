package initializer

import (
	"forum/Log"
	"forum/internal/models"
	"log"
)

func SyncDatabase() {
	log.Println("Syncing User model with database")
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		Log.Err.Println("Failed to sync the User table:", err)
	}

	DB.Raw("")

	log.Println("Syncing Post model with database")
	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		Log.Err.Println("Failed to sync the Post table:", err)
	}

	log.Println("Syncing Reply model with database")
	err = DB.AutoMigrate(&models.Reply{})
	if err != nil {
		Log.Err.Println("Failed to sync the Reply table:", err)
	}

	log.Println("Syncing PostLike model with database")
	err = DB.AutoMigrate(&models.PostLike{})
	if err != nil {
		Log.Err.Println("Failed to sync the PostLike table:", err)
	}

	log.Println("Syncing ReplyLike model with database")
	err = DB.AutoMigrate(&models.ReplyLike{})
	if err != nil {
		Log.Err.Println("Failed to sync the ReplyLike table:", err)
	}

	log.Println("Database sync complete")
}
