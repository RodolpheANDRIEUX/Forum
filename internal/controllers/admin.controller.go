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

	//users, err := utils.GetAllUsersExcept(user.UserID)
	users, err := utils.GetAllUsers()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin.html", gin.H{"error": err})
		return
	}

	var reportedPosts []models.Post
	err = initializer.DB.Where("report > ?", 0).Preload("User").Find(&reportedPosts).Error
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
