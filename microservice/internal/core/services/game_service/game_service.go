package gameService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	types "microservice/internal/core/types"
)

type GameService struct {
	UserRepo drivenPorts.UserPort
	GamePort drivenPorts.GamePort
	Logger   *log.Logger
	Matches  map[int]types.Match
}

func MakeGameService(userRepo drivenPorts.UserPort, gameAdapter drivenPorts.GamePort, logger *log.Logger) GameService {

	return GameService{
		UserRepo: userRepo,
		GamePort: gameAdapter,
		Logger:   logger,
		Matches:  make(map[int]types.Match),
	}
}

func (service GameService) StartGame(userIDList []int64) (int, error) {
	return 0, nil
}

func (service GameService) ConnectPlayer(userID int, connection interface{}) error {
	return service.GamePort.ConnectPlayer(userID, connection)
}

func (service GameService) SendMatchStartNotice(userID int, matchID int) error {
	return service.SendMatchStartNotice(userID, matchID)
}
