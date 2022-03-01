package drivenAdapters

import (
	"fmt"
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

func (adapter WebsocketGameAdapter) SendMatchStartNotice(userID int, matchID int) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"matchID": %v}`, matchID)))
	if err != nil {
		return err
	}

	return nil
}
