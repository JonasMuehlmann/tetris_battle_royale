package main

import (
	"log"
	userService "microservice/api/services/user_service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Forwards requests to it's registered handlers
	// by matching the endpoint (e.g. "/") to the handler
	// This is the gateway in the microservice diagram
	mux := mux.NewRouter()

	// TODO: The routers can be simplified with gorilla/mux
	mux.Handle("/", http.FileServer(http.Dir("../client/build/")))

	mux.HandleFunc("/login", userService.LoginHandler).Methods("POST")
	mux.HandleFunc("/register", userService.RegisterHandler).Methods("POST")
	mux.HandleFunc("/isLogin{userId:[0-9]+}", userService.IsLoginHandler).Methods("GET")
	mux.HandleFunc("/logout{userId:[0-9]+}", userService.LogoutHandler).Methods("DELETE")

	log.Println("Starting server on Port 8080")
	log.Fatalf("server failed to start: %v", http.ListenAndServe(":8080", mux))
}
