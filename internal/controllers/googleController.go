package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// Google OAuth
func HandleAuth(c *gin.Context) {
	url := initializer.OauthConfig.AuthCodeURL(initializer.State)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func HandleAuthCallback(c *gin.Context) {
	code := c.Query("code")

	token, err := initializer.OauthConfig.Exchange(context.Background(), code)

	fmt.Println("error=", err)

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to exchange token")
		return
	}

	client := initializer.OauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get user info")
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to read response body"})
		return
	}

	// Parse user information
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err = json.Unmarshal(data, &body)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to parse user info"})
		return
	}

	// Hash the password
	hash, err := utils.HasPassword(body.Password)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to hash password"})
		return
	}

	// Create the user
	user := models.User{Role: "member", Email: body.Email, Password: string(hash)}
	result := initializer.DB.Create(&user)
	if result.Error != nil {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{"error": "This user already exist"})
		return
	}

	// Respond
	c.HTML(http.StatusOK, "user.html", gin.H{"username": body.Email})
}
