package controllers

import (
	"forum/internal/initializer"
	"forum/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getPost() ([]models.PostWeb, error) {
	var posts []models.Post
	result := initializer.DB.Order("created_at desc").Limit(15).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	var postsWeb []models.PostWeb
	for _, post := range posts {
		var user models.User
		err := initializer.DB.Model(&models.User{}).Where("user_id = ?", post.UserID).First(&user).Error
		if err != nil {
			return nil, err
		}
		var likeNumber int
		postWeb := models.PostWeb{
			PostID:         post.PostID,
			UserID:         post.UserID,
			Username:       user.Username,
			ProfilePicture: user.ProfileImg,
			Message:        post.Message,
			Picture:        post.Picture,
			Topic:          post.Topic,
			Like:           likeNumber,
		}
		postsWeb = append(postsWeb, postWeb)
	}
	return postsWeb, nil
}
func DisplayPost(c *gin.Context) {
	postsWeb, err := getPost()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusFound, gin.H{
		"posts": postsWeb,
	})
	return
}
