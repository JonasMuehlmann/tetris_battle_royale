package drivingAdapters

import (
	"bytes"
	"errors"
	"log"
	common "microservice/internal"
	drivingPorts "microservice/internal/core/driving_ports"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MatchmakingServiceRestAdapter struct {
	Service drivingPorts.MatchmakingServicePort
	Logger  *log.Logger
}

// TODO: This callback mechanic is not really the way to go.
// Instead, we should open a websocket connection where the server can send a message
func (adapter MatchmakingServiceRestAdapter) buildMatchStartCallBack(w http.ResponseWriter, clientAddress string) func(int) error {
	clientHost, clientPort, err := net.SplitHostPort(clientAddress)
	clientPort = "8082"

	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Could not register callback for match start")
	}
	return func(matchID int) error {

		matchIDStr := strconv.FormatInt(int64(matchID), 10)
		responseBuffer := bytes.NewBuffer([]byte("{matchID: " + matchIDStr + "}"))

		callbackResponse, err := http.Post(clientHost+":"+clientPort, "application/json", responseBuffer)

		if err != nil {
			return err
		}
		if callbackResponse.StatusCode != http.StatusOK {
			return errors.New("Failed to send callback to client")
		}

		return nil
	}
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

	err = adapter.Service.Join(int(userId), adapter.buildMatchStartCallBack(w, r.RemoteAddr))
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

	userID, err := strconv.ParseInt(body["userId"].(string), 10, 32)
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		common.TryWriteResponse(w, "Could not unmarshal request body")
	}

	err = adapter.Service.Leave(int(userID))
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
