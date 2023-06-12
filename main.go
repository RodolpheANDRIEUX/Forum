package main

import (
	"forum/internal/initializer"
	"forum/internal/server"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDb()
	initializer.SyncDatabase()
	initializer.InitGoogleOAuth()
	initializer.InitGithubOAuth()
}

func main() {
	server.Serve()
}
