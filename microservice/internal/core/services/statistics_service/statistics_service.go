package statisticsService

import (
	"log"
	ipc "microservice/internal/core/driven_adapters/ipc/grpc"
	"microservice/internal/core/driven_ports/repository"
	"microservice/internal/core/types"
)

type StatisticsService struct {
	UserRepository       repository.UserRepositoryPort
	StatisticsRepository repository.StatisticsRepositoryPort
	IPCServer            ipc.StatisticsServiceIPCServerAdapter
	Logger               *log.Logger
}

// TODO: Validate if user exists
func (service StatisticsService) GetPlayerProfile(userID string) (types.PlayerProfile, error) {
	return service.StatisticsRepository.GetPlayerProfile(userID)
}

func (service StatisticsService) GetPlayerStatistics(userID string) (types.PlayerStatistics, error) {
	return service.StatisticsRepository.GetPlayerStatistics(userID)
}

func (service StatisticsService) GetMatchRecords(userID string) ([]types.MatchRecord, error) {
	return service.StatisticsRepository.GetMatchRecords(userID)
}

func (service StatisticsService) GetMatchRecord(matchID string) (types.MatchRecord, error) {
	return service.StatisticsRepository.GetMatchRecord(matchID)
}

// func (service StatisticsService) AddMatchRecord(matchID string, record types.MatchRecord) error {
// 	return service.StatisticsRepository.AddMatchRecord(matchID, record)
// }

func (service StatisticsService) UpdatePlayerProfile(newProfile types.PlayerProfile) error {
	return service.StatisticsRepository.UpdatePlayerProfile(newProfile)
}

func (service StatisticsService) UpdatePlayerStatistics(newStatistics types.PlayerStatistics) error {
	return service.StatisticsRepository.UpdatePlayerStatistics(newStatistics)
}

func (service StatisticsService) AddMatchRecord(record types.MatchRecord) error {
	return service.StatisticsRepository.AddMatchRecord(record)
}
