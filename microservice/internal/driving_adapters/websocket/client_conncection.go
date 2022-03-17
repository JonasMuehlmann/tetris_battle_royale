package drivingAdapters

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientConnection struct {
	userID string
	conn   *websocket.Conn
}

func (c *ClientConnection) ReadPump(IncomingMesssages chan []byte) {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		IncomingMesssages <- message
	}
}
