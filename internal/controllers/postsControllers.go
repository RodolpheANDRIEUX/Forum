package controllers

import (
	"forum/Log"
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// provisoire
func IncrementLikes(c *gin.Context) {
	// Convert the post ID to an integer.
	postID, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		Log.Err.Println("Post ID should be an integer")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Post ID should be an integer",
		})
		return
	}

	// Retrieve the post from the database.
	var post models.Post
	if err := initializer.DB.Preload("User").Where("post_id = ?", postID).First(&post).Error; err != nil {
		Log.Err.Println("Error retrieving post from the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error retrieving post from the database",
		})
		return
	}

	post.Like++

	if err := initializer.DB.Save(&post).Error; err != nil {
		Log.Err.Println("Error saving updated post to the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error saving updated post to the database",
		})
		return
	}

	// send a notifications
	user, err := utils.ParseUser(c)
	if err != nil {
		Log.Err.Println("Error parsing the user JWT", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error parsing the user JWT",
		})
		return
	}

	message := user.Username + " liked your post #" + strconv.Itoa(int(post.PostID))
	err = sendNotification(post.UserID, message)
	if err != nil {
		Log.Err.Println("Error sending a notification", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error sending a notification",
		})
		return
	}

	// Success
	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"newLikes": post.Like,
	})
}
