package routes

import (
	"forum/controllers"
	"forum/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.LoadHTMLGlob("views/*")
	router.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })

	router.GET("/signup", func(c *gin.Context) { c.HTML(http.StatusOK, "signup.html", nil) })
	router.POST("/signup", controllers.Signup)

	router.GET("/login", func(c *gin.Context) { c.HTML(http.StatusOK, "login.html", nil) })
	router.POST("/login", controllers.Login)

	router.GET("/logout", controllers.Logout)

	router.GET("/user", middleware.RequireAuth, controllers.User)
}
