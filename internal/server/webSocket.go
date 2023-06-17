package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"forum/internal/controllers"
	"forum/internal/models"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Body struct {
	Message string `json:"message"`
	File    string `json:"file"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Connections = make(map[*websocket.Conn]bool)
var Mutex = &sync.Mutex{}

func WebsocketHandler(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade:", err)
		return
	}

	Mutex.Lock()
	Connections[conn] = true
	Mutex.Unlock()

	defer func() {
		Mutex.Lock()
		delete(Connections, conn)
		Mutex.Unlock()
		if err := conn.Close(); err != nil {
			fmt.Println("Failed to close connection:", err)
			return
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message:", err)
			break
		}

		post := models.Post{}

		// Try to parse the incoming message as JSON
		var incoming Body
		err = json.Unmarshal(msg, &incoming)

		// If the message could be parsed as JSON and contains a file, decode the file
		if err == nil {
			post.Message = incoming.Message // Set the message from the parsed JSON

			if incoming.File != "" {
				fileData, err := base64.StdEncoding.DecodeString(incoming.File)
				if err != nil {
					fmt.Println("Failed to decode file data:", err)
					break
				}
				post.Picture = fileData // Save the file data in the Post object
			}
		} else {
			// If the message could not be parsed as JSON, treat it as a plain message
			post.Message = string(msg)
		}

		// Save the post in the database
		err, code := controllers.AddPostInDB(&post, c)

		if err != nil {
			fmt.Printf("Error: %s, Code: %d\n", err, code)
			break
		}

		broadcastMessage, err := json.Marshal(post)
		if err != nil {
			fmt.Println("Failed to marshal post:", err)
			break
		}

		Mutex.Lock()
		for conn := range Connections {
			err = conn.WriteMessage(websocket.TextMessage, broadcastMessage)
			if err != nil {
				fmt.Println("Failed to write message:", err)
				delete(Connections, conn)
				Mutex.Unlock()
				return
			}
		}
		Mutex.Unlock()
	}
}
