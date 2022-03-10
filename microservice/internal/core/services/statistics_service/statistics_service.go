package statisticsService

import (
	"log"
	ipc "microservice/internal/core/driven_adapters/ipc/grpc"
	"microservice/internal/core/driven_ports/repository"
	"microservice/internal/core/types"
	"time"
)

type StatisticsService struct {
	UserRepository       repository.UserRepositoryPort
	StatisticsRepository repository.StatisticsRepositoryPort
	// TODO: If this has a port as a type, we don't have an import cycle and can refer to this struct in the adapter
	IPCServer ipc.StatisticsServiceIPCServerAdapter
	Logger    *log.Logger
}

// TODO: Validate if user exists
func (service StatisticsService) GetPlayerProfile(userID string) (types.PlayerProfile, error) {
	return service.StatisticsRepository.GetPlayerProfile(userID)
}

func (service StatisticsService) GetPlayerStatistics(userID string) (types.PlayerStatistics, error) {
	return service.StatisticsRepository.GetPlayerStatistics(userID)
}

func (service StatisticsService) GetPlayerRating(userID string) (types.PlayerRating, error) {
	return service.StatisticsRepository.GetPlayerRating(userID)
}

func (service StatisticsService) GetMatchRecords(userID string) ([]types.MatchRecord, error) {
	return service.StatisticsRepository.GetMatchRecords(userID)
}

func (service StatisticsService) GetMatchRecord(matchID string) (types.MatchRecord, error) {
	return service.StatisticsRepository.GetMatchRecord(matchID)
}

func (service StatisticsService) UpdatePlayerProfile(newProfile types.PlayerProfile) error {
	return service.StatisticsRepository.UpdatePlayerProfile(newProfile)
}

func (service StatisticsService) UpdatePlayerStatistics(newStatistics types.PlayerStatistics) error {
	return service.StatisticsRepository.UpdatePlayerStatistics(newStatistics)
}

func (service StatisticsService) UpdatePlayerRating(newRating types.PlayerRating) error {
	return service.StatisticsRepository.UpdatePlayerRating(newRating)
}

func (service StatisticsService) AddMatchRecord(record types.MatchRecord) error {
	playerProfile, err := service.GetPlayerProfile(record.UserID)
	if err != nil {
		return err
	}

	playerStatistics, err := service.GetPlayerStatistics(record.UserID)
	if err != nil {
		return err
	}

	playerRating, err := service.GetPlayerRating(record.UserID)
	if err != nil {
		return err
	}

	playerProfile.Playtime += record.Length
	playerProfile.LastUpdate = record.Start.Add(time.Minute * time.Duration(record.Length))

	playerStatistics.Score += record.Score
	playerStatistics.ScorePerMinute = float32(playerStatistics.Score) / float32(playerProfile.Playtime)

	if record.Win {
		playerStatistics.Wins += 1
	} else {
		playerStatistics.Losses += 1
	}

	playerStatistics.Winrate = float32(playerStatistics.Wins) / float32(playerStatistics.Wins+playerStatistics.Losses)

	playerRating.MMR += record.RatingChange

	switch record.WinKind {
	case types.WinTop10:
		playerStatistics.WinsAsTop10 += 1
	case types.WinTop5:
		playerStatistics.WinsAsTop5 += 1
	case types.WinTop3:
		playerStatistics.WinsAsTop3 += 1
	case types.WinTop1:
		playerStatistics.WinsAsTop1 += 1
	default:
		log.Println("Unhandled WinKinnd")
	}

	err = service.UpdatePlayerProfile(playerProfile)
	if err != nil {
		return err
	}

	err = service.UpdatePlayerStatistics(playerStatistics)
	if err != nil {
		return err
	}

	err = service.UpdatePlayerRating(playerRating)
	if err != nil {
		return err
	}

	return service.StatisticsRepository.AddMatchRecord(record)
}
