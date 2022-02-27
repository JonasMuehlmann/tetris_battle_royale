package drivingAdapters

import (
	"log"
	drivingPorts "microservice/internal/core/driving_ports"
)

type GameServiceRestAdapter struct {
	Service drivingPorts.GameServicePort
	Logger  *log.Logger
}

func (adapter GameServiceRestAdapter) Run() {
	// mux := mux.NewRouter()

	// mux.HandleFunc("/join", adapter.JoinHandler).Methods("POST")
	// mux.HandleFunc("/leave", adapter.LeaveHandler).Methods("POST")

	// adapter.Logger.Println("Starting server on Port 8080")
	// log.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
