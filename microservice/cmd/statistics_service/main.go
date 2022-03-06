package main

import (
	common "microservice/internal"
	ipc "microservice/internal/core/driven_adapters/ipc/grpc"
	postgresRepository "microservice/internal/core/driven_adapters/repository/postgres"
	repository "microservice/internal/core/driven_adapters/repository/postgres"
	statisticsService "microservice/internal/core/services/statistics_service"
	"microservice/internal/core/types"
	drivingAdapters "microservice/internal/driving_adapters/rest"
	"net"
)

func main() {
	logger := common.NewDefaultLogger()

	db := repository.MakeDefaultPostgresDB(logger)
	statisticsRepository := postgresRepository.PostgresDatabaseStatisticsRepository{Logger: logger, PostgresDatabase: *db}

	statisticsService := statisticsService.StatisticsService{Logger: logger, StatisticsRepository: statisticsRepository}

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	ipcServer := ipc.StatisticsServiceIPCServerAdapter{
		StatisticsServiceServer: ipc.StatisticsServiceServer{
			Logger: logger,
		},
		Logger: logger,
	}

	statisticsService.IPCServer = ipcServer
	statisticsServiceAdapter := drivingAdapters.StatisticsServiceRestAdapter{Logger: logger, Service: statisticsService}

	grpcServerArgs := types.DrivenAdapterGRPCArgs{
		Service:  &statisticsService,
		Listener: listener,
	}

	go func() {
		err := statisticsService.IPCServer.Start(grpcServerArgs)
		if err != nil {
			logger.Fatalf("Error: %v", err)
		}
	}()

	statisticsServiceAdapter.Run()
}
