package matchmakingService

import (
	"context"
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	gameServiceProto "microservice/internal/core/protofiles/game_service"

	"google.golang.org/grpc"
)

const MatchSize = 3

type MatchmakingService struct {
	UserRepo              drivenPorts.UserPort
	Logger                *log.Logger
	Queue                 map[int]bool
	GameServiceGrpcClient gameServiceProto.GameServiceClient
}

func MakeMatchmakingService(userRepo drivenPorts.UserPort, logger *log.Logger) (MatchmakingService, error) {
	grpcConn, err := grpc.Dial("game-service:8081", grpc.WithInsecure())

	if err != nil {
		return MatchmakingService{}, err
	}

	gameServiceGrpcClient := gameServiceProto.NewGameServiceClient(grpcConn)

	matchmakingService := MatchmakingService{
		UserRepo:              userRepo,
		Logger:                logger,
		Queue:                 make(map[int]bool),
		GameServiceGrpcClient: gameServiceGrpcClient,
	}

	return matchmakingService, nil
}

func (service MatchmakingService) Join(userID int, matchStartCallback func(int) error) error {
	service.Queue[userID] = true

	service.Logger.Printf("Player %v joined the queue", userID)

	if len(service.Queue) == MatchSize {
		err := service.startGame()
		if err != nil {
			return err
		}
	}

	return nil
}

func (service MatchmakingService) Leave(userID int) error {
	delete(service.Queue, userID)

	service.Logger.Printf("Player %v left the queue", userID)

	return nil
}

func (service MatchmakingService) startGame() error {
	userIDList := make([]int64, 0, len(service.Queue))

	for k, v := range service.Queue {
		if v {
			userIDList = append(userIDList, int64(k))
		}
	}

	_, err := service.GameServiceGrpcClient.StartGame(context.Background(), &gameServiceProto.UserIDList{Id: userIDList})

	// TODO: Notify clients

	if err != nil {
		return err
	}

	service.Logger.Println("Started a game")

	return nil
}
