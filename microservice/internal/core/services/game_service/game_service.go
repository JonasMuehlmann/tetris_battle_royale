package matchmakingService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	types "microservice/internal/core/types"
)

type GameService struct {
	UserRepo drivenPorts.UserPort
	Logger   *log.Logger
	Matches  map[int]types.Match
}
