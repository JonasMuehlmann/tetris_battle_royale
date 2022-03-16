package drivenAdapters

import (
	"encoding/json"
	"fmt"
	"log"
	"microservice/internal/core/types"

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

func (adapter WebsocketGameAdapter) SendMatchStartNotice(userID string, matchID string, opponents []types.Opponent) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	data := map[string]interface{}{
		"matchID":   matchID,
		"opponents": opponents,
	}

	dataJson, err := json.Marshal(data)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	err = userConn.WriteMessage(websocket.TextMessage, dataJson)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}

func (adapter WebsocketGameAdapter) SendUpdatedBlockState(userID string, newState types.BlockState) error {
	// TODO: Implement
	return nil
}

func (adapter WebsocketGameAdapter) SendBlockLockinNotice(userID string) error {
	// TODO: Implement
	return nil
}

func (adapter WebsocketGameAdapter) SendRowClearNotice(userID string, rowNum int) error {
	// TODO: Implement
	return nil
}

func (adapter WebsocketGameAdapter) SendBlockSpawnNotice(userID string, dequeuedBlock types.BlockType, enqueuedBlock types.BlockType) error {
	// TODO: Implement
	return nil
}

func (adapter WebsocketGameAdapter) SendScoreGain(userID string, score int) error {
	// TODO: Implement
	return nil
}

func (adapter WebsocketGameAdapter) SendEventNotice(userID string, event string) error {
	// TODO: Implement
	return nil
}
