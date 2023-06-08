package server

import (
	"forum/internal/initializer"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func Serve() {
	router := gin.Default()

	// parse assets and templates
	router.Static("/web", "./web")
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
	errorLog := log.New(io.MultiWriter(logFile, os.Stderr), "[ERROR] ", log.LstdFlags)

	// routes
	InitRoutes(router)
	Routes(router)

	// run
	log.Println("running the server...")
	err := router.Run()
	if err != nil {
		errorLog.Println(err)
		panic(err)
	}
}
