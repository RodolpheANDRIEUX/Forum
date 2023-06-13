package controllers

import (
	"fmt"
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func SendUsername(c *gin.Context) {
	user, err := utils.GetUSer(c)

	if err != nil {
		c.HTML(http.StatusUnauthorized, "index.html", gin.H{"error": err})
		Logout(c)
		return
	}

	newUsername := c.PostForm("username")
	user.Username = newUsername

	result := initializer.DB.Save(&user)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "This username already exist"})
		return
	}
	utils.CreateJWT(c, &user)
	c.Next()
	//c.Redirect(http.StatusFound, "/user")
}

func UploadProfileImg(c *gin.Context) {
	// get the file
	file, err := c.FormFile("profile-img")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	// get the user
	user, err := utils.GetUSer(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	// save the profile image path to db
	path := filepath.Join("/uploads", file.Filename)
	user.ProfileImg = path

	result := initializer.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	// save the file locally
	err = c.SaveUploadedFile(file, filepath.Join("./assets/uploads", file.Filename))

	if err != nil {
		fmt.Println("err= ", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Render the page
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Profile updated successfully",
	})
}

func SendProfileData(c *gin.Context) {
	var user models.User

	user, err := utils.GetUSer(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "failed to load profile",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
