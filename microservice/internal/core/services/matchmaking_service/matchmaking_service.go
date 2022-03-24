package matchmakingService

import (
	"context"
	"log"
	repoPorts "microservice/internal/core/driven_ports/repository"
	gameServiceProto "microservice/internal/core/protofiles/game_service"

	"google.golang.org/grpc"
)

const MatchSize = 2

type MatchmakingService struct {
	UserRepository        repoPorts.UserRepositoryPort
	Logger                *log.Logger
	Queue                 map[string]bool
	GameServiceGrpcClient gameServiceProto.GameServiceClient
}

func MakeMatchmakingService(userRepo repoPorts.UserRepositoryPort, logger *log.Logger) (MatchmakingService, error) {
	grpcConn, err := grpc.Dial("game-service:8081", grpc.WithInsecure())
	if err != nil {
		return MatchmakingService{}, err
	}

	gameServiceGrpcClient := gameServiceProto.NewGameServiceClient(grpcConn)

	matchmakingService := MatchmakingService{
		UserRepository:        userRepo,
		Logger:                logger,
		Queue:                 make(map[string]bool),
		GameServiceGrpcClient: gameServiceGrpcClient,
	}

	return matchmakingService, nil
}

func (service *MatchmakingService) Join(userID string) error {
	service.Queue[userID] = true

	service.Logger.Printf("Player %v joined the queue", userID)

	if len(service.Queue) == MatchSize {
		err := service.startGame()
		if err != nil {
			service.Logger.Printf("Error: %v\n", err)

			return err
		}
	}

	return nil
}

func (service *MatchmakingService) Leave(userID string) error {
	delete(service.Queue, userID)

	service.Logger.Printf("Player %v left the queue", userID)

	return nil
}

func (service *MatchmakingService) startGame() error {
	userIDList := make([]string, 0, len(service.Queue))

	for k, v := range service.Queue {
		if v {
			userIDList = append(userIDList, k)
		}
	}

	_, err := service.GameServiceGrpcClient.StartGame(context.Background(), &gameServiceProto.UserIDList{Id: userIDList})
	if err != nil {
		service.Logger.Printf("Error: %v\n", err)

		return err
	}

	service.Logger.Println("Started a game")

	return nil
}
