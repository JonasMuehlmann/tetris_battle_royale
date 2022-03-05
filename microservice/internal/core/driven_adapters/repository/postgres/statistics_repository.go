package repository

import (
	"log"
	types "microservice/internal/core/types"
)

type PostgresDatabaseStatisticsRepository struct {
	PostgresDatabase
	Logger *log.Logger
}

func (repo PostgresDatabaseStatisticsRepository) GetPlayerProfile(userID string) (types.PlayerProfile, error) {
	var playerProfile types.PlayerProfile
	db, err := repo.GetConnection()
	if err != nil {
		return types.PlayerProfile{}, err
	}

	err = db.Get(&playerProfile, "SELECT * FROM player_profiles WHERE user_id = $1", userID)
	if err != nil {
		return types.PlayerProfile{}, err
	}

	return playerProfile, nil
}

func (repo PostgresDatabaseStatisticsRepository) GetPlayerStatistics(userID string) (types.PlayerStatistics, error) {
	var playerStatistics types.PlayerStatistics
	db, err := repo.GetConnection()
	if err != nil {
		return types.PlayerStatistics{}, err
	}

	err = db.Get(&playerStatistics, "SELECT player_statistics.* FROM player_statistics LEFT JOIN player_profiles ON player_profiles.player_statistics_id = player_statistics.id WHERE user_id = $1", userID)
	if err != nil {
		return types.PlayerStatistics{}, err
	}

	return playerStatistics, nil
}

func (repo PostgresDatabaseStatisticsRepository) GetMatchRecords(userID string) ([]types.MatchRecord, error) {
	var matchRecords []types.MatchRecord
	db, err := repo.GetConnection()
	if err != nil {
		return []types.MatchRecord{}, err
	}

	err = db.Select(&matchRecords, "SELECT * FROM  match_records WHERE user_id = $1", userID)
	if err != nil {
		return []types.MatchRecord{}, err
	}

	return matchRecords, nil
}
func (repo PostgresDatabaseStatisticsRepository) GetMatchRecord(matchID string) (types.MatchRecord, error) {
	var matchRecord types.MatchRecord
	db, err := repo.GetConnection()
	if err != nil {
		return types.MatchRecord{}, err
	}

	err = db.Get(&matchRecord, "SELECT * FROM  match_records WHERE id = $1", matchID)
	if err != nil {
		return types.MatchRecord{}, err
	}

	return matchRecord, nil
}
