package server

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})

	router.GET("/ws", WebsocketHandler)

}
