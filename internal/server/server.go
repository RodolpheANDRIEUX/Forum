package server

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Serve() {
	router := gin.Default()

	// parse assets and templates
	router.Static("/web", "./web")
	router.LoadHTMLGlob("web/*.html")

	// create log file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// routes
	InitRoutes(router)
	Routes(router)

	// run
	err := router.Run()
	if err != nil {
		panic(err)
	}
}
