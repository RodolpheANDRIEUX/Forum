package controllers

import (
	"forum/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Admin(c *gin.Context) {
	user, err := utils.GetUSer(c)

	if err != nil {
		c.Redirect(http.StatusUnauthorized, "/")
	}

	users, err := utils.GetAllUsers()

	if err != nil {
		c.HTML(http.StatusInternalServerError, "admin.html", gin.H{"error": err})
		return
	}

	c.HTML(http.StatusOK, "admin.html",
		gin.H{
			"admin":    user,
			"allUsers": users,
		})
}
