package initializer

import (
	"forum/internal/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to sync the DB")
	}
}
