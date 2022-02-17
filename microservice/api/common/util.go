package common

import (
	"log"
	"net/http"
)

func TryWriteResponse(w http.ResponseWriter, response string) {

	_, err := w.Write([]byte(response))

	if err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}
