package controllers

import (
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

// on appel cette fonction a chaque fois que la route /ws est solicité
func WebsocketHandler(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	//ajoute connexion a la map
	Mutex.Lock()
	Connections[conn] = true
	Mutex.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// si il y a une erreur on supprime la connexion de la map
			Mutex.Lock()
			delete(Connections, conn)
			Mutex.Unlock()
			return
		}

		// affiche le message sur toutes les connexions active
		Mutex.Lock()
		for conn := range Connections {
			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				delete(Connections, conn)
			}
		}
		Mutex.Unlock()
	}
}
