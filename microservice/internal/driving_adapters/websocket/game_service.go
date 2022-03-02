package drivingAdapters

import (
	"log"
	common "microservice/internal"
	drivingPorts "microservice/internal/core/driving_ports"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type GameServiceWebsocketAdapter struct {
	Service drivingPorts.GameServicePort
	Logger  *log.Logger
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (adapter GameServiceWebsocketAdapter) UpgradeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := common.UnmarshalRequestBody(r)
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, common.MakeJsonError("Could not establish websocket connection"))
	}

	userID, err := strconv.ParseInt(body["userID"].(string), 10, 32)
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, common.MakeJsonError("Could not establish websocket connection"))
	}

	incomingConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, common.MakeJsonError("Could not establish websocket connection"))
	}

	// TODO: Check if user exists
	err = adapter.Service.ConnectPlayer(int(userID), incomingConn)
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, common.MakeJsonError("Could not establish websocket connection"))
	}
}

func (adapter GameServiceWebsocketAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/ws", adapter.UpgradeHandler).Methods("POST")

	adapter.Logger.Println("Starting server on Port 8080")

	adapter.Logger.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
