package drivingAdapters

import (
	"log"
	drivingPorts "microservice/internal/core/driving_ports"
)

type MatchmakingServiceRestAdapter struct {
	Service drivingPorts.MatchmakingServicePort
	Logger  *log.Logger
}

func (adapter MatchmakingServiceRestAdapter) Run() {
	// mux := mux.NewRouter()

	// mux.Handle("/", http.FileServer(http.Dir("../client/build/")))

	// // NOTE: The api gateay should contain a prefix user/, which is stripped before forwarding
	// mux.HandleFunc("/login", adapter.LoginHandler).Methods("POST")
	// mux.HandleFunc("/register", adapter.RegisterHandler).Methods("POST")
	// mux.HandleFunc("/isLogin/{username:[a-zA-Z0-9]+}", adapter.IsLoginHandler).Methods("GET")
	// mux.HandleFunc("/logout", adapter.LogoutHandler).Methods("DELETE")

	// adapter.Logger.Println("Starting server on Port 8080")
	// log.Fatalf("Error: Server failed to start: %v", http.ListenAndServe(":8080", mux))
}
