package statisticsService

import (
	"log"
	drivenPorts "microservice/internal/core/driven_ports"
	"microservice/internal/core/types"
)

type StatisticsService struct {
	UserRepo       drivenPorts.UserPort
	StatisticsRepo drivenPorts.StatisticsPort
	Logger         *log.Logger
}

// TODO: Validate if user exists
func (service StatisticsService) GetPlayerProfile(userID string) (types.PlayerProfile, error) {
	return service.StatisticsRepo.GetPlayerProfile(userID)
}

func (service StatisticsService) GetPlayerStatistics(userID string) (types.PlayerStatistics, error) {
	return service.StatisticsRepo.GetPlayerStatistics(userID)
}
