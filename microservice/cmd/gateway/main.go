package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	mux := http.NewServeMux()

	services := []string{"user", "statistics"}

	for _, service := range services {
		// NOTE: The names are determined by the kubernetes services
		serviceURL, err := url.Parse("http://" + service + "-service:8080")

		if err != nil {
			log.Fatalf("Could not set up backend service routing: %v", err)
		}

		// Forwards the requests made to the specified endpoint to the respective backend service
		mux.Handle("/"+service, httputil.NewSingleHostReverseProxy(serviceURL))

		log.Printf("Registered backend service at endpoint /%v\n", service)
	}

	log.Println("Starting http server on Port 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}