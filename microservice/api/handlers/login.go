package handlers

import (
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Not needed, but better be explicit!
	w.Header().Set("Content-Type", "text/plain; charset")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte("Hello World from server!"))

	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}

func Login() bool {
	return false
}
