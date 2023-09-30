package controllers

import (
	"forum/Log"
	"forum/internal/initializer"
	"forum/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func getPost(offset int, limit int) ([]models.Post, error) {
	var posts []models.Post

	err := initializer.DB.Order("created_at desc").Offset(offset).Limit(limit).Preload("User").Find(&posts).Error
	if err != nil {
		Log.Err.Println("Failed to retrieve posts:", err)
		return nil, err
	}

	return posts, nil
}

func DisplayPost(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		Log.Err.Println("Error while getting page: ", err)
		page = 1
	}

	limit := 10

	offset := (page - 1) * limit

	posts, err := getPost(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusFound, gin.H{
		"posts": posts,
	})
	return
}

func getPostByID(postID int) (*models.Post, error) {
	var post models.Post

	err := initializer.DB.Preload("User").Where("post_id = ?", postID).First(&post).Error
	if err != nil {
		Log.Err.Println("Failed to retrieve post:", err)
		return nil, err
	}

	return &post, nil
}

func ShowPostPage(c *gin.Context) {
	postID, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid post ID",
		})
		return
	}

	post, err := getPostByID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("Successfully fetched post:", post)

	c.HTML(http.StatusOK, "post.html", gin.H{
		"post": post,
	})
}
