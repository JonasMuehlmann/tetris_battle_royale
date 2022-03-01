package main

import (
	"log"
	drivenAdapters "microservice/internal/core/driven_adapters/game_adapter"
	repository "microservice/internal/core/driven_adapters/repository/postgres"
	gameService "microservice/internal/core/services/game_service"
	drivingAdapters "microservice/internal/driving_adapters/websocket"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "TBR - ", log.Ltime|log.Lshortfile)

	db := repository.MakeDefaultPostgresDB(logger)
	userRepo := repository.PostgresDatabaseUserRepository{Logger: logger, PostgresDatabase: *db}
	gameAdapter := drivenAdapters.MakeWebsocketGameAdapter(logger)
	game_service := gameService.MakeGameService(userRepo, gameAdapter, logger)
	userServiceAdapter := drivingAdapters.GameServiceWebsocketAdapter{Logger: logger, Service: game_service}
	userServiceAdapter.Run()
}
