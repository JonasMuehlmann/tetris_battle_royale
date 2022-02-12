package main

import (
	"log"
	"microservice/api/handlers"
	"net/http"
)

func main() {
	// Forwards requests to it's registered handlers
	// by matching the endpoint (e.g. "/") to the handler
	// This is the gateway in the microservice diagram
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler)

	log.Println("Starting server on Port 8008")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
