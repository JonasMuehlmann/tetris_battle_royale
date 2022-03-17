package drivingAdapters

import (
	"encoding/json"
	"log"
	common "microservice/internal"
	gameService "microservice/internal/core/services/game_service"
	"microservice/internal/core/types"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type GameServiceWebsocketAdapter struct {
	Service           gameService.GameService
	Logger            *log.Logger
	clientConnections []ClientConnection
	IncomingMesssages chan []byte
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (adapter GameServiceWebsocketAdapter) UpgradeHandler(w http.ResponseWriter, r *http.Request) {
	incomingConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError("Could not establish websocket connection"))
	}

	userID := struct {
		UserID string `json:"userID"`
	}{}

	err = incomingConn.ReadJSON(&userID)
	if err != nil {
		log.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError("Could not read from websocket connection"))
	}

	// TODO: Check if user exists
	err = adapter.Service.ConnectPlayer(userID.UserID, *incomingConn)
	if err != nil {
		log.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError("Could not establish websocket connection"))
	}

	connection := ClientConnection{
		userID: userID.UserID,
		conn:   incomingConn,
	}

	adapter.clientConnections = append(adapter.clientConnections, connection)

	go connection.ReadPump(adapter.IncomingMesssages)
}

func (adapter GameServiceWebsocketAdapter) HandleMoveBlock(message map[string]string) error {

	var userID string = message["userID"]
	var matchID string = message["matchID"]
	var direction types.MoveDirection = types.MoveDirection(message["direction"])

	return adapter.Service.MoveBlock(userID, matchID, direction)
}

func (adapter GameServiceWebsocketAdapter) HandleRotateBlock(message map[string]string) error {
	var userID string = message["userID"]
	var matchID string = message["matchID"]
	var direction types.RotationDirection = types.RotationDirection(message["direction"])

	return adapter.Service.RotateBlock(userID, matchID, direction)
}

func (adapter GameServiceWebsocketAdapter) HandleHardDropBlock(message map[string]string) error {
	var userID string = message["userID"]
	var matchID string = message["matchID"]

	return adapter.Service.HardDropBlock(userID, matchID)
}

func (adapter GameServiceWebsocketAdapter) HandleToggleSoftDrop(message map[string]string) error {
	var userID string = message["userID"]
	var matchID string = message["matchID"]

	return adapter.Service.ToggleSoftDrop(userID, matchID)
}

func (adapter GameServiceWebsocketAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/ws", adapter.UpgradeHandler)

	adapter.Logger.Println("Starting server on Port 8080")

	go func() {
		adapter.Logger.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
	}()

	for {
		raw := <-adapter.IncomingMesssages
		var message map[string]string
		json.Unmarshal(raw, &message)
		switch message["type"] {
		case "MoveBlock":
			adapter.HandleMoveBlock(message)
		case "RotateBlock":
			adapter.HandleRotateBlock(message)
		case "HardDrop":
			adapter.HandleHardDropBlock(message)
		case "SoftDrop":
			adapter.HandleToggleSoftDrop(message)
		}

	}
}
