package gameService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	repoPorts "microservice/internal/core/driven_ports/repository"
	types "microservice/internal/core/types"

	"github.com/google/uuid"
)

type GameService struct {
	UserRepo repoPorts.UserRepositoryPort
	// This port/adapter might need refactoring
	Logger  *log.Logger
	Matches map[string]types.Match
	// TODO: This type's name is not consistent with the satistics service' ipc
	IPCServer drivenPorts.GamePort
}

func MakeGameService(userRepo repoPorts.UserRepositoryPort, gameAdapter drivenPorts.GamePort, logger *log.Logger) GameService {
	return GameService{
		UserRepo:  userRepo,
		Logger:    logger,
		Matches:   make(map[string]types.Match),
		IPCServer: gameAdapter,
	}
	// TODO: This belongs in the main file
	// gameServiceProto.RegisterGameServiceServer(grpcServer, &GameServiceServer{GameService: gameService})
}

func (service GameService) StartGame(userIDList []string) error {
	matchID := uuid.NewString()

	players := [types.MatchSize]types.Player{}
	for i, userID := range userIDList {
		// TODO: This should probably be refactored into a separate function and will include more complex setup logic
		players[i] = types.Player{
			ID:        userID,
			Score:     0,
			Playfield: &types.Playfield{},
		}

		err := service.IPCServer.SendMatchStartNotice(userID, matchID)
		if err != nil {
			service.Logger.Printf("Could not notify client %v of game start", userID)
			service.Logger.Printf("Error: %v", err)

			return err
		}
	}

	service.Matches[matchID] = types.Match{
		ID:      matchID,
		Players: players,
	}

	go service.Matches[matchID].Start()

	return nil
}

// TODO: This belongs in the main file
// func (service GameService) StartGrpcServer(listener net.Listener) error {
// 	return service.GrpcServer.Serve(listener)
// }

// NOTE: This function has nothing to do with the matchmaking
func (service GameService) ConnectPlayer(userID string, connection interface{}) error {
	return service.IPCServer.ConnectPlayer(userID, connection)
}
