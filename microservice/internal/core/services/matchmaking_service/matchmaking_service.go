package matchmakingService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
)

type MatchmakingService struct {
	UserRepo drivenPorts.UserPort
	Logger   *log.Logger
}

func (service MatchmakingService) Join() (int, error) {
	return 0, nil
}

func (service MatchmakingService) Leave() error {
	return nil
}
