package initializer

import (
	"fmt"
	"forum/Log"
	"forum/internal/models"
)

func SyncDatabase() {
	fmt.Println("1")
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Printf("Failed to sync the DB: %v\n", err)

	}

	DB.Raw("")
	fmt.Println("2")

	err = DB.AutoMigrate(&models.Post{})
	if err != nil {
		fmt.Printf("Failed to sync the DB: %v\n", err)
		Log.Err.Printf("Failed to sync the Post table: %v", err)
	}
	fmt.Println("3")

	err = DB.AutoMigrate(&models.Reply{})
	if err != nil {
		fmt.Printf("Failed to sync the DB: %v\n", err)
		Log.Err.Printf("Failed to sync the Reply table: %v", err)
	}
	fmt.Println("4")

	err = DB.AutoMigrate(&models.PostLike{})
	if err != nil {
		fmt.Printf("Failed to sync the DB: %v\n", err)
		Log.Err.Printf("Failed to sync the Like table: %v", err)
	}
	fmt.Println("5")

	err = DB.AutoMigrate(&models.ReplyLike{})
	if err != nil {
		fmt.Printf("Failed to sync the DB: %v\n", err)
		Log.Err.Printf("Failed to sync the Like table: %v", err)
	}
	fmt.Println("6")

}
