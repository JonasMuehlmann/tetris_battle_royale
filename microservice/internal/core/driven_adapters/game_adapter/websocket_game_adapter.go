package drivenAdapters

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type WebsocketGameAdapter struct {
	Logger            *log.Logger
	PlayerConnections map[string]websocket.Conn
}

func MakeWebsocketGameAdapter(logger *log.Logger) WebsocketGameAdapter {
	return WebsocketGameAdapter{
		Logger:            logger,
		PlayerConnections: make(map[string]websocket.Conn),
	}
}

func (adapter WebsocketGameAdapter) ConnectPlayer(userID string, connection interface{}) error {
	conn, ok := connection.(websocket.Conn)
	if !ok {
		return fmt.Errorf("Invalid type %T for argument, expected %T", connection, websocket.Conn{})
	}

	adapter.PlayerConnections[userID] = conn

	return nil
}

func (adapter WebsocketGameAdapter) SendMatchStartNotice(userID string, matchID string) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"matchID": "%v"}`, matchID)))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}
