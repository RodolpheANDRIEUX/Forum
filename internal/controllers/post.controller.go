package controllers

import (
	"encoding/base64"
	"forum/Log"
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func UniquePost(c *gin.Context) {
	postID, err := strconv.Atoi(c.Query("post_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	post, err := utils.GetPost(uint(postID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"post": post, "replies": post.Replies})
}

func Reply(c *gin.Context) {
	type Body struct {
		Message string `json:"message"`
		File    string `json:"file"`
		PostID  uint   `json:"postID"`
	}
	var body Body
	err := c.ShouldBind(&body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	// convert file to blob
	pic, err := base64.StdEncoding.DecodeString(body.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reply := models.Reply{
		Message: body.Message,
		Picture: pic,
		PostID:  body.PostID,
	}

	// get the user
	reply.User, err = utils.ParseUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save the reply to the db
	result := initializer.DB.Create(&reply)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// send a notification to the creator of the post

	// respond
	c.JSON(http.StatusOK, gin.H{
		"reply": reply,
	})
}

func RepostPost(c *gin.Context) {
	type Body struct {
		PostID string `json:"postID"`
	}

	var body Body
	err := c.ShouldBind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	postID, err := strconv.Atoi(body.PostID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// get the post
	post, err := utils.GetPost(uint(postID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// increment the report counter
	err = utils.UpdatePost(post.PostID, post.Message, post.Deleted, post.Report+1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// send a notification to the creator of the post
	err = sendNotification(post.UserID, "Your post "+strconv.Itoa(int(post.PostID))+" has been reported.")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{})
}
