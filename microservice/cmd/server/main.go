package main

import (
	"log"
	handlers "microservice/api/services/user_service"
	"net/http"
)

func main() {
	// Forwards requests to it's registered handlers
	// by matching the endpoint (e.g. "/") to the handler
	// This is the gateway in the microservice diagram
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/isLogin", handlers.IsLoginHandler)
	mux.HandleFunc("/logut", handlers.LogoutHandler)

	log.Println("Starting server on Port 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
