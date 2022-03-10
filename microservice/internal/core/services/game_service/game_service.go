package gameService

import (
	"context"
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	repoPorts "microservice/internal/core/driven_ports/repository"
	gameServiceProto "microservice/internal/core/protofiles/game_service"
	types "microservice/internal/core/types"
	"net"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type GameService struct {
	UserRepo repoPorts.UserRepositoryPort
	// This port/adapter might need refactoring
	GamePort   drivenPorts.GamePort
	Logger     *log.Logger
	Matches    map[string]types.Match
	GrpcServer *grpc.Server
}

type GameServiceServer struct {
	gameServiceProto.UnimplementedGameServiceServer
	GameService GameService
}

// TODO: This should be a driven adapter
func (service GameServiceServer) StartGame(context context.Context, userIDList *gameServiceProto.UserIDList) (*gameServiceProto.EmptyMessage, error) {
	matchID, err := service.GameService.StartGame(userIDList.GetId())

	if err != nil {
		return nil, err
	}

	return &gameServiceProto.EmptyMessage{}, nil
}

func MakeGameService(userRepo repoPorts.UserRepositoryPort, gameAdapter drivenPorts.GamePort, logger *log.Logger) GameService {
	grpcServer := grpc.NewServer()

	gameService := GameService{
		UserRepo:   userRepo,
		GamePort:   gameAdapter,
		Logger:     logger,
		Matches:    make(map[string]types.Match),
		GrpcServer: grpcServer,
	}

	gameServiceProto.RegisterGameServiceServer(grpcServer, &GameServiceServer{GameService: gameService})

	return gameService
}

// TODO: This should not return the match id
func (service GameService) StartGame(userIDList []string) error {
	matchID := uuid.NewString()

	players := [types.MatchSize]types.Player{}
	for i, userID := range userIDList {
		players[i] = types.Player{
			ID:        userID,
			Score:     0,
			Playfield: &types.Playfield{},
		}

		err := service.GamePort.SendMatchStartNotice(userID, matchID)
		if err != nil {
			service.Logger.Printf("Could not notify client %v of game start", userID)
			service.Logger.Println("Error: %v", err)
			return err
		}
	}

	service.Matches[matchID] = types.Match{
		ID:      matchID,
		Players: players,
	}

	return nil
}

func (service GameService) StartGrpcServer(listener net.Listener) error {
	return service.GrpcServer.Serve(listener)
}

func (service GameService) ConnectPlayer(userID string, connection interface{}) error {
	return service.GamePort.ConnectPlayer(userID, connection)
}
