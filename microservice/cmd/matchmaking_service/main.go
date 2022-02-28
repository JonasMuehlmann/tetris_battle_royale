package main

import (
	"log"
	repository "microservice/internal/core/repository/postgres"
	matchmakingService "microservice/internal/core/services/matchmaking_service"
	drivingAdapters "microservice/internal/driving_adapters/rest"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "TBR - ", log.Ltime|log.Lshortfile)

	db := repository.MakeDefaultPostgresDB(logger)
	userRepository := repository.PostgresDatabaseUserRepository{Logger: logger, PostgresDatabase: *db}
	matchmakingService, err := matchmakingService.MakeMatchmakingService(userRepository, logger)

	if err != nil {
		logger.Fatal(err)
	}

	userServiceAdapter := drivingAdapters.MatchmakingServiceRestAdapter{Logger: logger, Service: matchmakingService}
	userServiceAdapter.Run()
}
