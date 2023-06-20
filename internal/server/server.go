package server

import (
	"forum/Log"
	"forum/internal/initializer"
	"github.com/gin-gonic/gin"
	"log"
)

func Serve() {
	router := gin.Default()

	// parse assets and templates
	router.Static("/css", "./web/css")
	router.Static("/script", "./web/script")
	router.Static("/img", "./web/img")

	router.LoadHTMLGlob("web/*.html")
	router.MaxMultipartMemory = 20 << 20 // 20 MiB

	gin.DefaultWriter = initializer.LogFile
	defer func() {
		if err := initializer.LogFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// routes
	InitRoutes(router)
	Routes(router)

	// run
	log.Println("running the server...")
	err := router.Run()
	if err != nil {
		Log.Err.Println(err)
		panic(err)
	}
}
