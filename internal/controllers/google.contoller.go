package controllers

import (
	"context"
	"encoding/json"
	"forum/internal/initializer"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func HandleGoogleAuth(c *gin.Context) {
	state := c.Query("state")
	url := initializer.GoogleOauthConfig.AuthCodeURL(initializer.GoogleState)
	queryUrl := utils.AddQuery(url, "state", state)
	c.Redirect(http.StatusTemporaryRedirect, queryUrl)
}

func HandleGoogleCallback(c *gin.Context) {
	code := c.Query("code")
	token, err := initializer.GoogleOauthConfig.Exchange(context.Background(), code)

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to exchange token")
		return
	}

	client := initializer.GoogleOauthConfig.Client(context.Background(), token)
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
	var body Body

	err = json.Unmarshal(data, &body)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "signup.html", gin.H{"error": "Failed to parse user info"})
		return
	}

	phase := c.Query("state")

	// if login
	if phase == "login" {
		message, errorCode := Authorize(c, body)

		if errorCode != http.StatusOK {
			c.HTML(errorCode, "login.html", gin.H{"error": message})
			return
		}
		// Respond
		c.Redirect(http.StatusFound, "/user")
	} else {
		// if register
		dbErr, errorCode := SignupAndStore(c, body)

		if errorCode != http.StatusOK {
			c.HTML(errorCode, "signup.html", gin.H{"error": dbErr})
			return
		}
		// redirect to the configuration of the account
		c.Redirect(http.StatusFound, "/user")
		return
	}
}
