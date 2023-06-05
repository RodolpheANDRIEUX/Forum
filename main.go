package main

import (
	"forum/initianlizers"
	"forum/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	initianlizers.LoadEnvVariables()
	initianlizers.ConnectToDb()
	initianlizers.SyncDatabase()
}

func main() {
	router := gin.Default()

	routes.Routes(router)

	err := router.Run()
	if err != nil {
		return
	}
}
