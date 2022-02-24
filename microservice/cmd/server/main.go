package main

import (
	"log"
	repository2 "microservice/internal/core/repository/postgres"
	userService2 "microservice/internal/core/services/user_service"
	drivingAdapters "microservice/internal/driving_adapters/rest"
	"os"
)

func main() {
	// Forwards requests to it's registered handlers
	// by matching the endpoint (e.g. "/") to the handler
	// This is the gateway in the microservice diagram

	logger := log.New(os.Stdout, "TBR - ", log.Ltime|log.Lshortfile)

	// TODO: Set correct response codes
	db := repository2.MakeDefaultPostgresDB(logger)
	userRepository := repository2.PostgresDatabaseUserRepository{Logger: logger, PostgresDatabase: *db}
	sessionRepository := repository2.PostgresDatabaseSessionRepository{Logger: logger, PostgresDatabase: *db}
	userService := userService2.UserService{Logger: logger, UserRepo: userRepository, SessionRepo: sessionRepository}
	userServiceAdapter := drivingAdapters.UserServiceRestAdapter{Logger: logger, Service: userService}
	userServiceAdapter.Run()
}
