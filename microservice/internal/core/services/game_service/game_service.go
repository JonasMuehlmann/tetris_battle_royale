package gameService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	repoPorts "microservice/internal/core/driven_ports/repository"
	types "microservice/internal/core/types"

	ipcPorts "microservice/internal/core/driven_ports/ipc"

	"github.com/google/uuid"
)

type GameService struct {
	UserRepo repoPorts.UserRepositoryPort
	// This port/adapter might need refactoring
	Logger      *log.Logger
	Matches     map[string]types.Match
	IPCServer   ipcPorts.GameServiceIPCServerPort
	GameAdapter drivenPorts.GamePort
}

func MakeGameService(userRepo repoPorts.UserRepositoryPort, ipcServerAdapter ipcPorts.GameServiceIPCServerPort, gameAdapter drivenPorts.GamePort, logger *log.Logger) GameService {
	return GameService{
		UserRepo:    userRepo,
		Logger:      logger,
		Matches:     make(map[string]types.Match),
		IPCServer:   ipcServerAdapter,
		GameAdapter: gameAdapter,
	}
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

		err := service.GameAdapter.SendMatchStartNotice(userID, matchID)
		if err != nil {
			service.Logger.Printf("Could not notify client %v of game start", userID)
			service.Logger.Printf("Error: %v\n", err)

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

// NOTE: This function has nothing to do with the matchmaking
func (service GameService) ConnectPlayer(userID string, connection interface{}) error {
	return service.GameAdapter.ConnectPlayer(userID, connection)
}

func (service GameService) MoveBlock(userID string, direction types.MoveDirection) error {
	// TODO: Implement
	return nil
}

func (service GameService) RotateBlock(userID string, direction types.RotationDirection) error {
	// TODO: Implement
	return nil
}

func (service GameService) HardDropBlock(userID string) error {
	// TODO: Implement
	return nil
}

func (service GameService) ToggleSoftDrop(userID string) error {
	// TODO: Implement
	return nil
}
