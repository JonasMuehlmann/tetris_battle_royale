package matchmakingService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
)

const MatchSize = 10

type MatchmakingService struct {
	UserRepo drivenPorts.UserPort
	Logger   *log.Logger
	Queue    map[int]bool
}

func (service MatchmakingService) Join(userID int) error {
	service.Queue[userID] = true

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

	return nil
}

func (service MatchmakingService) startGame() error {
	// TODO: Hand off players to game service over grpc
	return nil
}
