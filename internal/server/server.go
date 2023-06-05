package server

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Initserver() {
	router := gin.Default()

	// Charger les templates HTML
	router.LoadHTMLGlob("web/*.html")

	// sert pour  les fichiers JavaScript
	router.Static("/web", "./web")

	// lier au fichier de log
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	InitRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
