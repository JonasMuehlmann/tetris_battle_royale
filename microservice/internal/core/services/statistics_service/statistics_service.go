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

func (service StatisticsService) GetMatchRecords(userID string) ([]types.MatchRecord, error) {
	return service.StatisticsRepo.GetMatchRecords(userID)
}

func (service StatisticsService) GetMatchRecord(matchID string) (types.MatchRecord, error) {
	return service.StatisticsRepo.GetMatchRecord(matchID)
}

// func (service StatisticsService) AddMatchRecord(matchID string, record types.MatchRecord) error {
// 	return service.StatisticsRepo.AddMatchRecord(matchID, record)
// }

func (service StatisticsService) UpdatePlayerProfile(newProfile types.PlayerProfile) error {
	return service.StatisticsRepo.UpdatePlayerProfile(newProfile)
}

func (service StatisticsService) UpdatePlayerStatistics(newStatistics types.PlayerStatistics) error {
	return service.StatisticsRepo.UpdatePlayerStatistics(newStatistics)
}
