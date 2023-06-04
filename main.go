package main

import (
	"fmt"
	"forum/internal"
	"forum/internal/database"
)

func main() {
	fmt.Println("http://localhost:8080/")
	database.InitDB()
	internal.Initserver()
}
