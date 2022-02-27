package drivingAdapters

import (
	"log"
	drivingPorts "microservice/internal/core/driving_ports"
	"net/http"

	"github.com/gorilla/mux"
)

type MatchmakingServiceRestAdapter struct {
	Service drivingPorts.MatchmakingServicePort
	Logger  *log.Logger
}

func (adapter MatchmakingServiceRestAdapter) JoinHandler(w http.ResponseWriter, r *http.Request) {
}

func (adapter MatchmakingServiceRestAdapter) LeaveHandler(w http.ResponseWriter, r *http.Request) {
}

func (adapter MatchmakingServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/join", adapter.JoinHandler).Methods("POST")
	mux.HandleFunc("/leave", adapter.LeaveHandler).Methods("POST")

	adapter.Logger.Println("Starting server on Port 8080")
	log.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
