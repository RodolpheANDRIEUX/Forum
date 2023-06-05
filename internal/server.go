package internal

import (
	"forum/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Initserver() {
	router := gin.Default()

	// Charger les templates HTML
	router.LoadHTMLGlob("templates/*")

	// Servir les fichiers JavaScript
	router.Static("/script", "./static/script")

	router.GET("/", func(c *gin.Context) {
		// Renvoie la page HTML index.html
		c.HTML(200, "home.html", gin.H{})

	})

	router.GET("/ws", controllers.WebsocketHandler)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
