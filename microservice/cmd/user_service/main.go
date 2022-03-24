package main

import (
	common "microservice/internal"
	postgresRepository "microservice/internal/core/driven_adapters/repository/postgres"
	userService "microservice/internal/core/services/user_service"
	drivingAdapters "microservice/internal/driving_adapters/rest"
)

func main() {
	// Forwards requests to it's registered handlers
	// by matching the endpoint (e.g. "/") to the handler
	// This is the gateway in the microservice diagram

	logger := common.NewDefaultLogger()

	db := postgresRepository.MakeDefaultPostgresDB(logger)
	userRepository := postgresRepository.PostgresDatabaseUserRepository{Logger: logger, PostgresDatabase: *db}
	sessionRepository := postgresRepository.PostgresDatabaseSessionRepository{Logger: logger, PostgresDatabase: *db}
	// sessionRepository := redisRepository.RedisSessionRepo{Logger: logger, RedisStore: redisRepository.MakeDefaultRedisStore(logger)}
	userService := userService.UserService{Logger: logger, UserRepository: &userRepository, SessionRepository: &sessionRepository}
	userServiceAdapter := drivingAdapters.UserServiceRestAdapter{Logger: logger, Service: &userService}
	userServiceAdapter.Run()
}
