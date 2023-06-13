package controllers

import (
	"fmt"
	"forum/Log"
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPostInDB(content string, c *gin.Context) error {
	user, err := utils.GetUSer(c)

	fmt.Println(user)

	if err != nil {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{"error": err})
		Logout(c)
		return err
	}

	newPost := models.Post{UserID: user.UserID, Message: content}

	result := initializer.DB.Create(&newPost)

	if result.Error != nil {
		//@todo voir quelle code erreur mettre
		Log.Err.Printf("Error while saving the post %v", result)
	}

	return nil
}
