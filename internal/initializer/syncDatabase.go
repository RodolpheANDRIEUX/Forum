package initializer

import (
	"forum/Log"
	"forum/internal/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		Log.Err.Panic("Failed to sync the DB")
	}
}
