package main

import (
	"log"
	userService "microservice/api/services/user_service"
	"net/http"
)

func main() {
	// Forwards requests to it's registered handlers
	// by matching the endpoint (e.g. "/") to the handler
	// This is the gateway in the microservice diagram
	mux := http.NewServeMux()
	userServiceMux := userService.MakeUserServiceMux()

	mux.Handle("/user/", http.StripPrefix("/user", userServiceMux))

	log.Println("Starting server on Port 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
