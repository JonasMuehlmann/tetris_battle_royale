package matchmakingService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
)

type GameService struct {
	UserRepo drivenPorts.UserPort
	Logger   *log.Logger
	Matches  map[int]bool
}
