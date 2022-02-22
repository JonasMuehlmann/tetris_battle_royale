package main

import (
	repository "microservice/internal/repository/postgres"
	userService "microservice/internal/services/user_service"
)

func main() {
	// Forwards requests to it's registered handlers
	// by matching the endpoint (e.g. "/") to the handler
	// This is the gateway in the microservice diagram

	// TODO: Add loggers to all structs
	db := repository.MakeDefaultPostgresDB()
	userRepository := repository.PostgresDatabaseUserRepository{*db}
	sessionRepository := repository.PostgresDatabaseSessionRepository{*db}
	userService := userService.UserService{userRepository, sessionRepository}
	userServiceAdapter := drivingAdapters.UserServiceRestAdapter{userService}
	userServiceAdapter.Run()
}
