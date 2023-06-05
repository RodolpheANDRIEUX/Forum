package main

import (
	"forum/internal/database"
	"forum/internal/initianlizers"
	"forum/internal/server"
)

func init() {
	initianlizers.LoadEnvVariables()
	database.ConnectToDb()
	database.SyncDatabase()
}

func main() {
	server.Serve()
}
