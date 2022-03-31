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

func (adapter *WebsocketGameAdapter) ConnectPlayer(userID string, connection interface{}) error {
	conn, ok := connection.(websocket.Conn)
	if !ok {
		return fmt.Errorf("Invalid type %T for argument, expected %T", connection, websocket.Conn{})
	}

	adapter.PlayerConnections[userID] = conn

	return nil
}

func (adapter *WebsocketGameAdapter) SendMatchStartNotice(userID string, matchID string, opponents []types.Opponent) error {
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

func (adapter *WebsocketGameAdapter) SendStartTetrominoPreview(userID string, newPreview []types.Tetromino) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("player with the id %v is not connected", userID)
	}

	out, jsonErr := json.Marshal(newPreview)
	if jsonErr != nil {
		adapter.Logger.Printf("Error: %v\n", jsonErr)

		return jsonErr
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(out))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}

func (adapter *WebsocketGameAdapter) SendUpdatedTetrominoState(userID string, newState types.TetrominoState) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	out, jsonErr := json.Marshal(newState)
	if jsonErr != nil {
		adapter.Logger.Printf("Error: %v\n", jsonErr)

		return jsonErr
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(out))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}

func (adapter *WebsocketGameAdapter) SendTetrominoLockinNotice(userID string) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(`{"LockIn": "true"}`))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}

func (adapter *WebsocketGameAdapter) SendRowClearNotice(userID string, rowNum int) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"rowNum": "%v"}`, rowNum)))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}

func (adapter *WebsocketGameAdapter) SendTetrominoSpawnNotice(userID string, newTetromino types.TetrominoName, enqueuedTetromino types.TetrominoName) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"newTetromino": "%v", "enqueuedTetromino": "%v"}`, newTetromino, enqueuedTetromino)))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}

func (adapter *WebsocketGameAdapter) SendScoreGain(userID string, score int) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"score": "%v"}`, score)))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}

func (adapter *WebsocketGameAdapter) SendEventNotice(userID string, event string) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"event": "%v"}`, event)))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}

func (adapter *WebsocketGameAdapter) SendEliminationNotice(userID string, eliminatedPlayerID string) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("Player with the id %v is not connected", userID)
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(`{"eliminated_player": "%v"}`, userID)))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}
func (adapter *WebsocketGameAdapter) SendEndOfMatchData(userID string, endOfMatchData types.EndOfMatchData) error {
	userConn, ok := adapter.PlayerConnections[userID]
	if !ok {
		return fmt.Errorf("player with the id %v is not connected", userID)
	}

	out, jsonErr := json.Marshal(endOfMatchData)
	if jsonErr != nil {
		adapter.Logger.Printf("Error: %v\n", jsonErr)

		return jsonErr
	}

	err := userConn.WriteMessage(websocket.TextMessage, []byte(out))
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)

		return err
	}

	return nil
}
