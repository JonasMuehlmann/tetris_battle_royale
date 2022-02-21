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
	mux.HandleFunc("/login", userService.LoginHandler)
	mux.HandleFunc("/isLogin", userService.IsLoginHandler)
	mux.HandleFunc("/logout", userService.LogoutHandler)

	log.Println("Starting server on Port 8080")
	log.Fatalf("server failed to start: %v", http.ListenAndServe(":8080", mux))
}
