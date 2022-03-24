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

func (adapter *MatchmakingServiceRestAdapter) JoinHandler(w http.ResponseWriter, r *http.Request) {
	body, err := common.UnmarshalRequestBody(r)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, common.MakeJsonError("Could not unmarshal request body"))
	}

	// TODO: Validate if user exists

	userId, ok := body["userId"].(string)
	if !ok {
		adapter.Logger.Printf("Error: Could not unmarshal request body")
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, common.MakeJsonError("Could not unmarshal request body"))
	}

	err = adapter.Service.Join(userId)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError("Could not join matchmaking"))
	}
}

func (adapter *MatchmakingServiceRestAdapter) LeaveHandler(w http.ResponseWriter, r *http.Request) {
	body, err := common.UnmarshalRequestBody(r)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, common.MakeJsonError("Could not unmarshal request body"))
	}

	// TODO: Validate if user exists

	userID, ok := body["userId"].(string)
	if !ok {
		adapter.Logger.Printf("Error: Could not unmarshal request body")
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, common.MakeJsonError("Could not unmarshal request body"))
	}

	err = adapter.Service.Leave(userID)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		common.TryWriteResponse(w, common.MakeJsonError("Could not join matchmaking"))
	}
}

func (adapter *MatchmakingServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/join", adapter.JoinHandler).Methods("POST")
	mux.HandleFunc("/leave", adapter.LeaveHandler).Methods("POST")

	adapter.Logger.Println("Starting server on Port 8080")
	adapter.Logger.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
