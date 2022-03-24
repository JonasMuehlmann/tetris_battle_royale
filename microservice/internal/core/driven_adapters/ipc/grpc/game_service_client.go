package ipc

import (
	"context"
	"log"
	gameServiceProto "microservice/internal/core/protofiles/game_service"

	"google.golang.org/grpc"
)

type GameServiceIPCClientAdapter struct {
	GrpcClient gameServiceProto.GameServiceClient
	Logger     *log.Logger
}

func (adapter *GameServiceIPCClientAdapter) StartGame(userIDList []string) error {
	message := &gameServiceProto.UserIDList{Id: userIDList}

	_, err := adapter.GrpcClient.StartGame(context.Background(), message)

	return err
}

func (adapter *GameServiceIPCClientAdapter) Start(args interface{}) error {
	serverAddr := "game-service:8081"

	grpcConn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}

	adapter.Logger.Printf("Connected to GRPC server at %v", serverAddr)

	adapter.GrpcClient = gameServiceProto.NewGameServiceClient(grpcConn)

	return nil
}
