package drivenAdapters

import (
	"log"

	"github.com/gorilla/websocket"
)

type WebsocketGameAdapter struct {
	Logger            *log.Logger
	PlayerConnections map[int]websocket.Conn
}

func MakeWebsocketGameAdapter(logger *log.Logger) WebsocketGameAdapter {
	return WebsocketGameAdapter{
		Logger:            logger,
		PlayerConnections: make(map[int]websocket.Conn),
	}
}

func (adapter WebsocketGameAdapter) ConnectPlayer(userID int, connection interface{}) error {
	adapter.PlayerConnections[userID] = connection.(websocket.Conn)

	return nil
}

func (adapter WebsocketGameAdapter) SendMatchStartNotice(userID int) error {
	return nil
}
