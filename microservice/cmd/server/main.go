package main

import (
	"log"
	drivingAdapters "microservice/internal/driving_adapters/rest"
	repository "microservice/internal/repository/postgres"
	userService "microservice/internal/services/user_service"
	"os"
)

func main() {
	// Forwards requests to it's registered handlers
	// by matching the endpoint (e.g. "/") to the handler
	// This is the gateway in the microservice diagram

	logger := log.New(os.Stdout, "TBR", log.Ltime|log.Lshortfile)

	db := repository.MakeDefaultPostgresDB(logger)
	userRepository := repository.PostgresDatabaseUserRepository{Logger: logger, PostgresDatabase: *db}
	sessionRepository := repository.PostgresDatabaseSessionRepository{Logger: logger, PostgresDatabase: *db}
	userService := userService.UserService{Logger: logger, UserRepo: userRepository, SessionRepo: sessionRepository}
	userServiceAdapter := drivingAdapters.UserServiceRestAdapter{Logger: logger, Service: userService}
	userServiceAdapter.Run()
}
