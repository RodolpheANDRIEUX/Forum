package initianlizers

import "forum/models"

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to sync the DB")
	}
}
