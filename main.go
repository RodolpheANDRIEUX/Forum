package main

import (
	"fmt"
	"forum/internal/database"
	"forum/internal/server"
)

func main() {
	fmt.Println("http://localhost:8080/")
	database.InitDB()
	server.Initserver()
}
