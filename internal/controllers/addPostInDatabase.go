package controllers

import (
	"forum/Log"
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddPostInDB(post *models.Post, c *gin.Context) (error, int) {
	user, err := utils.GetUSer(c)

	if err != nil {
		Logout(c)
		return err, http.StatusUnauthorized
	}

	post.User = user
	post.UserID = user.UserID

	result := initializer.DB.Create(&post)

	if result.Error != nil {
		Log.Err.Printf("Error while saving the post %v", result)
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
