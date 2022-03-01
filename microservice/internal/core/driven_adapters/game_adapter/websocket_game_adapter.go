package drivenAdapters

import (
	"log"

	"github.com/gorilla/websocket"
)

type WebsocketGameAdapter struct {
	Logger            *log.Logger
	PlayerConnections map[int]websocket.Conn
}

func (adapter WebsocketGameAdapter) ConnectPlayer(userID int, connection websocket.Conn) error {
	adapter.PlayerConnections[userID] = connection
}
