package controllers

import (
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
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

	// Open the file
	fileData, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	defer fileData.Close()

	// Read the file data into a byte slice
	fileBytes, err := ioutil.ReadAll(fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
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

	// Assign the file data to the user's profile image field
	user.ProfileImg = fileBytes

	// save the user
	result := initializer.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err,
		})
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
