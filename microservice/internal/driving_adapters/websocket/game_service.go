package drivingAdapters

import (
	"log"
	common "microservice/internal"
	drivingPorts "microservice/internal/core/driving_ports"
	"net"
	"net/http"

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
	incomingConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, common.MakeJsonError("Could not establish websocket connection"))
	}

	// TODO: Read userID from request
	userID := 0

	// TODO: Handler error
	err = adapter.Service.ConnectPlayer(userID, incomingConn)
}
func (adapter GameServiceWebsocketAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/ws", adapter.UpgradeHandler)

	adapter.Logger.Println("Starting server on Port 8080")

	go func() {
		adapter.Logger.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
	}()

	adapter.Logger.Println("Starting grpc server on Port 8081")

	grpcListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		adapter.Logger.Fatalf("Could not start listener for grcp server: %v", err)
	}

	adapter.Logger.Fatalf("Error: GRPC Server failed to start: %v", adapter.Service.StartGrpcServer(grpcListener))
}
