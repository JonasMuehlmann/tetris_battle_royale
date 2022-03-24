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

func (adapter *StatisticsServiceRestAdapter) GetPlayerProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	playerProfile, err := adapter.Service.GetPlayerProfile(vars["userID"])
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	marhshalledPlayerProfile, err := json.Marshal(playerProfile)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	common.TryWriteResponse(w, `{"playerProfile": "`+string(marhshalledPlayerProfile)+`"}`)
}

func (adapter *StatisticsServiceRestAdapter) GetPlayerStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	playerStatistics, err := adapter.Service.GetPlayerStatistics(vars["userID"])
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	marhshalledPlayerStatistics, err := json.Marshal(playerStatistics)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	common.TryWriteResponse(w, `{"playerStatistics": "`+string(marhshalledPlayerStatistics)+`"}`)
}

func (adapter *StatisticsServiceRestAdapter) GetPlayerMatchRecordsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	matchRecords, err := adapter.Service.GetMatchRecords(vars["userID"])
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	marhshalledMatchRecords, err := json.Marshal(matchRecords)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	common.TryWriteResponse(w, `{"matchRecords": "`+string(marhshalledMatchRecords)+`"}`)
}

func (adapter *StatisticsServiceRestAdapter) GetPlayerMatchRecordHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	matchRecord, err := adapter.Service.GetMatchRecord(vars["matchID"])
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	marhshalledMatchRecords, err := json.Marshal(matchRecord)
	if err != nil {
		adapter.Logger.Printf("Error: %v\n", err)
		common.TryWriteResponse(w, common.MakeJsonError(err.Error()))

		return
	}

	common.TryWriteResponse(w, `{"matchRecord": "`+string(marhshalledMatchRecords)+`"}`)
}

func (adapter *StatisticsServiceRestAdapter) Run() {
	mux := mux.NewRouter()

	mux.HandleFunc("/playerProfile/{userID:[a-zA-Z0-9]+}", adapter.GetPlayerProfileHandler).Methods("GET")
	mux.HandleFunc("/playerStatistics/{userID:[a-zA-Z0-9]+}", adapter.GetPlayerStatisticsHandler).Methods("GET")
	mux.HandleFunc("/matchRecords/{userID:[a-zA-Z0-9]+}", adapter.GetPlayerMatchRecordsHandler).Methods("GET")
	mux.HandleFunc("/matchRecord/{matchID:[a-zA-Z0-9]+}", adapter.GetPlayerMatchRecordHandler).Methods("GET")

	adapter.Logger.Println("Starting server on Port 8080")
	log.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
