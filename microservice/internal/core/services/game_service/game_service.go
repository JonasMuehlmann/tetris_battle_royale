package gameService

import (
	"context"
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	repoPorts "microservice/internal/core/driven_ports/repository"
	gameServiceProto "microservice/internal/core/protofiles/game_service"
	types "microservice/internal/core/types"
	"net"

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

func (service GameServiceServer) StartGame(context context.Context, userIDList *gameServiceProto.UserIDList) (*gameServiceProto.MatchID, error) {
	matchID, err := service.GameService.StartGame(userIDList.GetId())

	if err != nil {
		return nil, err
	}

	return &gameServiceProto.MatchID{Id: matchID}, nil
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

func (service GameService) StartGame(userIDList []string) (string, error) {
	// TODO: Generate UUID and add players to match map
	return "", nil
}

func (service GameService) StartGrpcServer(listener net.Listener) error {
	return service.GrpcServer.Serve(listener)
}

func (service GameService) ConnectPlayer(userID string, connection interface{}) error {
	return service.GamePort.ConnectPlayer(userID, connection)
}

func (service GameService) SendMatchStartNotice(userID string, matchID string) error {
	return service.SendMatchStartNotice(userID, matchID)
}
