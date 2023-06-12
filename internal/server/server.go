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
	router.LoadHTMLGlob("web/*.html")

	// init log files
	logFile := initializer.InitLogs()
	gin.DefaultWriter = logFile
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// set default log
	log.SetOutput(gin.DefaultWriter)

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
