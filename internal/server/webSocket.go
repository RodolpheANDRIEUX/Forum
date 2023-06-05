package server

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// pour passer de la connexion http a ws
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// garde une trace de toute les connexions actives
var Connections = make(map[*websocket.Conn]bool)

var Mutex = &sync.Mutex{}

func WebsocketHandler(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade:", err)
		return
	}

	//ajoute connexion a la map
	Mutex.Lock()
	Connections[conn] = true
	Mutex.Unlock()

	defer func() {
		// si il y a une erreur on supprime la connexion de la map
		Mutex.Lock()
		delete(Connections, conn)
		Mutex.Unlock()

		if err := conn.Close(); err != nil {
			fmt.Println("Failed to close connection:", err)
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed to read message:", err)
			break
		}

		// affiche le message sur toutes les connexions active
		Mutex.Lock()
		for conn := range Connections {
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Println("Failed to write message:", err)
				delete(Connections, conn)
			}
		}
		Mutex.Unlock()
	}
}
