package controllers

import (
	"forum/Log"
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPostInDB(content string, c *gin.Context) (error, int) {
	user, err := utils.GetUSer(c)

	if err != nil {
		Logout(c)
		return err, http.StatusUnauthorized
	}

	newPost := models.Post{UserID: user.UserID, Message: content}

	result := initializer.DB.Create(&newPost)

	if result.Error != nil {
		//@todo: voir quelle code erreur mettre
		Log.Err.Printf("Error while saving the post %v", result)
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
