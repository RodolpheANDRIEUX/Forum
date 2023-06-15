package controllers

import (
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
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

	if newUsername != "" {
		user.Username = newUsername

		result := initializer.DB.Save(&user)
		if result.Error != nil {
			c.Set("message", "Username is already taken")
			c.Set("status", 409)
			c.Next()
		}
		utils.CreateJWT(c, &user)
		c.Set("message", "Username updated successfully")
	}
	c.Set("status", 200)
	c.Next()
	//c.Redirect(http.StatusFound, "/user")
}

func UploadProfileImg(c *gin.Context) {
	// get the last message (if exist)
	message := c.GetString("message")
	status := c.GetInt("status")

	// get the file
	file, err := c.FormFile("profile-img")

	if file == nil {
		c.JSON(status, gin.H{
			"message": message,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	// convert the file to blob image
	fileBytes, err := fileToBlob(file)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// save profile picture
	code, err := saveProfilePicture(c, fileBytes)

	if err != nil {
		c.JSON(code, gin.H{
			"success": false,
			"message": message + err.Error(),
		})
		return
	}

	// Render the page
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message + " Profile updated successfully",
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

func saveProfilePicture(c *gin.Context, file []byte) (int, error) {
	user, err := utils.GetUSer(c)
	if err != nil {
		return http.StatusUnauthorized, err
	}

	// Assign the file data to the user's profile image field
	user.ProfileImg = file

	// save the user
	result := initializer.DB.Save(&user)
	if result.Error != nil {
		return http.StatusInternalServerError, err
	}
	return 200, nil
}

func fileToBlob(file *multipart.FileHeader) ([]byte, error) {
	// Open the file
	fileData, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fileData.Close()

	// Read the file data into a byte slice
	fileBytes, err := ioutil.ReadAll(fileData)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}
