package drivingAdapters

import (
	"log"
	common "microservice/internal"
	drivingPorts "microservice/internal/core/driving_ports"
	"net/http"

	"github.com/gorilla/mux"
)

type MatchmakingServiceRestAdapter struct {
	Service drivingPorts.MatchmakingServicePort
	Logger  *log.Logger
}

func (adapter MatchmakingServiceRestAdapter) JoinHandler(w http.ResponseWriter, r *http.Request) {
	body, err := common.UnmarshalRequestBody(r)
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Could not unmarshal request body")
	}

	err = adapter.Service.Join(body["userId"].(int))
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Could not join matchmaking")
	}
}

func (adapter MatchmakingServiceRestAdapter) LeaveHandler(w http.ResponseWriter, r *http.Request) {
	body, err := common.UnmarshalRequestBody(r)
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Could not unmarshal request body")
	}

	err = adapter.Service.Leave(body["userId"].(int))
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Could not join matchmaking")
	}
}

func (adapter MatchmakingServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/join", adapter.JoinHandler).Methods("POST")
	mux.HandleFunc("/leave", adapter.LeaveHandler).Methods("POST")

	adapter.Logger.Println("Starting server on Port 8080")
	log.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
