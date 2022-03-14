package drivingAdapters

import (
	"log"
	common "microservice/internal"
	gameService "microservice/internal/core/services/game_service"
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

func (adapter GameServiceWebsocketAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/ws", adapter.UpgradeHandler)

	adapter.Logger.Println("Starting server on Port 8080")

	adapter.Logger.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
