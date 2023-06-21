package controllers

import (
	"forum/internal/initializer"
	"forum/internal/models"
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendNotification(userID uint, message string) error {
	notif := models.Notifications{UserID: userID, Message: message}
	result := initializer.DB.Create(&notif)
	return result.Error
}

func IgnoreNotification(c *gin.Context) {
	type Body struct {
		NotificationID uint `json:"notificationID"`
	}
	var body Body
	err := c.ShouldBind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	notif, err := utils.GetNotification(body.NotificationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// delete the notif
	err = initializer.DB.Unscoped().Delete(&notif).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{})
}
