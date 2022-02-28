package gameService

import (
	"context"
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	gameServiceProto "microservice/internal/core/protofiles/game_service"
	types "microservice/internal/core/types"
	"net"

	"google.golang.org/grpc"
)

type GameService struct {
	UserRepo   drivenPorts.UserPort
	Logger     *log.Logger
	Matches    map[int]types.Match
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

	return &gameServiceProto.MatchID{Id: int64(matchID)}, nil
}

func MakeGameService(userRepo drivenPorts.UserPort, logger *log.Logger) GameService {
	grpcServer := grpc.NewServer()

	gameService := GameService{
		UserRepo:   userRepo,
		Logger:     logger,
		Matches:    make(map[int]types.Match),
		GrpcServer: grpcServer,
	}

	gameServiceProto.RegisterGameServiceServer(grpcServer, &GameServiceServer{GameService: gameService})

	return gameService
}

func (service GameService) StartGame(userIDList []int64) (int, error) {
	return 0, nil
}

func (service GameService) StartGrpcServer(listener net.Listener) error {
	return service.GrpcServer.Serve(listener)
}
