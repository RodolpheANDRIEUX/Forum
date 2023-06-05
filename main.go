package main

import (
	"forum/controllers"
	"forum/initianlizers"
	"forum/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initianlizers.LoadEnvVariables()
	initianlizers.ConnectToDb()
	initianlizers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.GET("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	err := r.Run()
	if err != nil {
		return
	}
}
