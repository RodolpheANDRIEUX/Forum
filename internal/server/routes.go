package server

import (
	"forum/internal/controllers"
	"forum/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/home", func(c *gin.Context) { c.HTML(200, "home.html", nil) })
	router.GET("/ws", WebsocketHandler)
}

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) { c.HTML(http.StatusOK, "index.html", nil) })

	router.GET("/signup", func(c *gin.Context) { c.HTML(http.StatusOK, "signup.html", nil) })
	router.POST("/signup", controllers.Signup)

	router.GET("/first_connection", func(c *gin.Context) { c.HTML(http.StatusOK, "newUser.html", nil) })
	router.POST("/first_connection", controllers.SendUsername)

	router.GET("/login", func(c *gin.Context) { c.HTML(http.StatusOK, "login.html", nil) })
	router.POST("/login", controllers.Login)

	// Google OAuth routes
	router.GET("/auth/google", controllers.HandleGoogleAuth)
	router.GET("/auth/google/callback", controllers.HandleGoogleCallback)

	// GitHub OAuth routes
	router.GET("/auth/github", controllers.HandleGitHubLogin)
	router.GET("/auth/github/callback", controllers.HandleGithubCallback)

	router.GET("/logout", controllers.Logout)

	router.GET("/user", middleware.RequireAuth, controllers.User)
}
