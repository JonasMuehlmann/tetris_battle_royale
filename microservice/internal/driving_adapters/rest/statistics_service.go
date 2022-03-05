package drivingAdapters

import (
	"encoding/json"
	"log"
	common "microservice/internal"
	drivingPorts "microservice/internal/core/driving_ports"
	"net/http"

	"github.com/gorilla/mux"
)

type StatisticsServiceRestAdapter struct {
	Service drivingPorts.StatisticsServicePort
	Logger  *log.Logger
}

func (adapter StatisticsServiceRestAdapter) GetPlayerProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	playerProfile, err := adapter.Service.GetPlayerProfile(vars["userID"])
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	marhshalleDplayerProfile, err := json.Marshal(playerProfile)
	if err != nil {
		adapter.Logger.Printf("Error: %v", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	common.TryWriteResponse(w, `{"playerProfile": "`+string(marhshalleDplayerProfile)+`"}`)
}

func (adapter StatisticsServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/playerProfile/{userID:[a-zA-Z0-9]+}", adapter.GetPlayerProfileHandler).Methods("GET")

	adapter.Logger.Println("Starting server on Port 8080")
	log.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
