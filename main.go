package main

import (
	"forum/internal/initializer"
	"forum/internal/server"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDb()
	initializer.SyncDatabase()
}

func main() {
	server.Serve()
}
