package drivingAdapters

import (
	"log"
	common "microservice/internal"
	drivingPorts "microservice/internal/core/driving_ports"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MatchmakingServiceRestAdapter struct {
	Service drivingPorts.MatchmakingServicePort
	Logger  *log.Logger
}

func (adapter MatchmakingServiceRestAdapter) JoinHandler(w http.ResponseWriter, r *http.Request) {
	body, err := common.UnmarshalRequestBody(r)
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, "Could not unmarshal request body")
	}

	// TODO: Validate if user exists

	userId, err := strconv.ParseInt(body["userId"].(string), 10, 32)
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, "Could not unmarshal request body")
	}

	err = adapter.Service.Join(int(userId))
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, "Could not join matchmaking")
	}
}

func (adapter MatchmakingServiceRestAdapter) LeaveHandler(w http.ResponseWriter, r *http.Request) {
	body, err := common.UnmarshalRequestBody(r)
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, "Could not unmarshal request body")
	}

	// TODO: Validate if user exists

	userId, err := strconv.ParseInt(body["userId"].(string), 10, 32)
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, "Could not unmarshal request body")
	}

	err = adapter.Service.Leave(int(userId))
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, "Could not join matchmaking")
	}
}

func (adapter MatchmakingServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/join", adapter.JoinHandler).Methods("POST")
	mux.HandleFunc("/leave", adapter.LeaveHandler).Methods("POST")

	adapter.Logger.Println("Starting server on Port 8080")
	adapter.Logger.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
