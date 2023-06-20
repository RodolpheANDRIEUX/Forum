package controllers

import (
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Admin(c *gin.Context) {
	user, err := utils.GetUSer(c)

	if err != nil {
		c.Redirect(http.StatusUnauthorized, "/")
	}

	users, err := utils.GetAllUsersExceptAdmins()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin.html", gin.H{"error": err})
		return
	}

	var reportedPosts []models.Post
	err = initializer.DB.Where("report > ? AND deleted = ?", 0, false).Preload("User").Find(&reportedPosts).Error
	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin.html", gin.H{"error": err})
		return
	}

	c.HTML(http.StatusOK, "admin.html",
		gin.H{
			"admin":         user,
			"allUsers":      users,
			"reportedPosts": reportedPosts,
		})
}

func UpdateUser(c *gin.Context) {
	type Body struct {
		UserID   uint   `json:"userID"`
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	var body Body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var user models.User
	initializer.DB.First(&user, body.UserID)
	user.Username = body.Username
	user.Role = body.Role

	result := initializer.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func DeletePost(c *gin.Context) {
	type Body struct {
		PostID uint   `json:"postID"`
		Admin  string `json:"admin"`
	}
	var body Body
	err := c.ShouldBind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	post, err := utils.GetPost(body.PostID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// delete the post
	err = utils.UpdatePost(body.PostID, post.Message, true)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// TODO: send a notification au user

	// respond
	c.JSON(http.StatusOK, gin.H{"error": "Post deleted"})
}

func BanUser(c *gin.Context) {
	type Body struct {
		UserID uint   `json:"userID"`
		Admin  string `json:"admin"`
	}
	var body Body
	err := c.ShouldBind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user, err := utils.GetUser(body.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// TODO: send a notification au user
	// TODO: delete the jwt

	// ban the user
	err = utils.UpdateUser(body.UserID, utils.CreateUniqueUsername(user.Email), "banned")

	// respond
	c.JSON(http.StatusOK, gin.H{"error": "User banned"})
}
