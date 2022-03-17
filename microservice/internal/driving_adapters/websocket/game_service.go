package drivingAdapters

import (
	"log"
	common "microservice/internal"
	gameService "microservice/internal/core/services/game_service"
	"microservice/internal/core/types"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type GameServiceWebsocketAdapter struct {
	Service gameService.GameService
	Logger  *log.Logger
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
}

func (adapter GameServiceWebsocketAdapter) HandleMoveBlock(message map[string]string) error {

	var userID string = message["userID"]
	var matchID string = message["matchID"]
	var direction types.MoveDirection = message["direction"]

	return adapter.Service.MoveBlock(userID, matchID, direction)
}

func (adapter GameServiceWebsocketAdapter) HandleRotateBlock(message map[string]string) error {
	// TODO: Implement
	var userID string = message["userID"]
	var matchID string = message["matchID"]
	var direction types.RotationDirection = message["direction"]

	return adapter.Service.RotateBlock(userID, matchID, direction)
}

func (adapter GameServiceWebsocketAdapter) HandleHardDropBlock(message map[string]string) error {
	var userID string = message["userID"]
	var matchID string = message["matchID"]

	return adapter.Service.HardDropBlock(userID, matchID)
}

func (adapter GameServiceWebsocketAdapter) HandleToggleSoftDrop(message map[string]string) error {
	// TODO: Implement
	var userID string
	var matchID string

	return adapter.Service.ToggleSoftDrop(userID, matchID)
}

func (adapter GameServiceWebsocketAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/ws", adapter.UpgradeHandler)

	adapter.Logger.Println("Starting server on Port 8080")

	go func() {
		adapter.Logger.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
	}()

	// TODO: Implement read loop and dispatch to the below functions
	for {
	}
}
