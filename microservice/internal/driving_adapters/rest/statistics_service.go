package drivingAdapters

import (
	"log"
	drivingPorts "microservice/internal/core/driving_ports"
	"net/http"

	"github.com/gorilla/mux"
)

type StatisticsServiceRestAdapter struct {
	Service drivingPorts.StatisticsServicePort
	Logger  *log.Logger
}

func (adapter StatisticsServiceRestAdapter) GetHandler(w http.ResponseWriter, r *http.Request) {
}

func (adapter StatisticsServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	// mux.HandleFunc("/login", adapter.LoginHandler).Methods("POST")

	adapter.Logger.Println("Starting server on Port 8080")
	log.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
