package drivingAdapters

import (
	"log"
	drivingPorts "microservice/internal/core/driving_ports"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

type GameServiceRestAdapter struct {
	Service drivingPorts.GameServicePort
	Logger  *log.Logger
}

func (adapter GameServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	// mux.HandleFunc("/join", adapter.JoinHandler).Methods("POST")
	// mux.HandleFunc("/leave", adapter.LeaveHandler).Methods("POST")

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
