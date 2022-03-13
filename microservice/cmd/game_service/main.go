package main

import (
	"log"
	drivenAdapters "microservice/internal/core/driven_adapters/game_adapter"
	ipc "microservice/internal/core/driven_adapters/ipc/grpc"
	repository "microservice/internal/core/driven_adapters/repository/postgres"
	gameService "microservice/internal/core/services/game_service"
	"microservice/internal/core/types"
	drivingAdapters "microservice/internal/driving_adapters/websocket"
	"net"
	"os"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)

	logger := log.New(os.Stdout, "TBR - ", log.Ltime|log.Lshortfile)

	db := repository.MakeDefaultPostgresDB(logger)
	userRepo := repository.PostgresDatabaseUserRepository{Logger: logger, PostgresDatabase: *db}
	gameAdapter := drivenAdapters.MakeWebsocketGameAdapter(logger)

	gameService := gameService.MakeGameService(userRepo, nil, gameAdapter, logger)

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Fatalf("Error: %v", err)
	}

	ipcServer := ipc.GameServiceIPCServerAdapter{
		GameServiceServer: ipc.GameServiceServer{
			Logger: logger,
		},
		Logger: logger,
	}

	gameService.IPCServer = ipcServer
	gameServiceAdapter := drivingAdapters.GameServiceWebsocketAdapter{Logger: logger, Service: gameService}

	grpcServerArgs := types.DrivenAdapterGRPCArgs{
		Service:  &gameService,
		Listener: listener,
	}

	go func() {
		err := gameService.IPCServer.Start(grpcServerArgs)
		if err != nil {
			logger.Fatalf("Error: %v", err)
		}
	}()

	gameServiceAdapter.Run()
}
