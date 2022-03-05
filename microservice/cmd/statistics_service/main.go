package main

import (
	"log"
	postgresRepository "microservice/internal/core/driven_adapters/repository/postgres"
	repository "microservice/internal/core/driven_adapters/repository/postgres"
	statisticsService "microservice/internal/core/services/statistics_service"
	drivingAdapters "microservice/internal/driving_adapters/rest"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "TBR - ", log.Ltime|log.Lshortfile)

	db := repository.MakeDefaultPostgresDB(logger)
	statisticsRepository := postgresRepository.PostgresDatabaseStatisticsRepository{Logger: logger, PostgresDatabase: *db}
	statisticsService := statisticsService.StatisticsService{Logger: logger, StatisticsRepo: statisticsRepository}
	statisticsServiceAdapter := drivingAdapters.StatisticsServiceRestAdapter{Logger: logger, Service: statisticsService}
	statisticsServiceAdapter.Run()
}
