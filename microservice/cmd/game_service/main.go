package main

import (
	"log"
	gameService "microservice/internal/core/services/game_service"
	drivingAdapters "microservice/internal/driving_adapters/rest"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "TBR - ", log.Ltime|log.Lshortfile)

	game_service := gameService.GameService{Logger: logger}
	userServiceAdapter := drivingAdapters.GameServiceRestAdapter{Logger: logger, Service: game_service}
	userServiceAdapter.Run()
}
