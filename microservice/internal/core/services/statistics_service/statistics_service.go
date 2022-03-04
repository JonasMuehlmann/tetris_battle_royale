package statisticsService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
)

type StatisticsService struct {
	UserRepo drivenPorts.UserPort
	Logger   *log.Logger
}
