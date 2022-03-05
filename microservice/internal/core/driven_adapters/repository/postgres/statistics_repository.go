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

	err := repo.DBConn.Get(&playerProfile, "SELECT * FROM player_profiles WHERE user_id = $1", userID)
	if err != nil {
		return types.PlayerProfile{}, err
	}

	return playerProfile, nil
}

func (repo PostgresDatabaseStatisticsRepository) GetPlayerStatistics(userID string) (types.PlayerStatistics, error) {
	var playerStatistics types.PlayerStatistics

	err := repo.DBConn.Get(&playerStatistics, "SELECT player_statistics.* FROM player_statistics LEFT JOIN player_profiles ON player_profiles.player_statistics_id = player_statistics.id WHERE user_id = $1", userID)
	if err != nil {
		return types.PlayerStatistics{}, err
	}

	return playerStatistics, nil
}

func (repo PostgresDatabaseStatisticsRepository) GetMatchRecords(userID string) ([]types.MatchRecord, error) {
	var matchRecords []types.MatchRecord

	err := repo.DBConn.Select(&matchRecords, "SELECT * FROM  match_records WHERE user_id = $1", userID)
	if err != nil {
		return []types.MatchRecord{}, err
	}

	return matchRecords, nil
}

func (repo PostgresDatabaseStatisticsRepository) GetMatchRecord(matchID string) (types.MatchRecord, error) {
	var matchRecord types.MatchRecord

	err := repo.DBConn.Get(&matchRecord, "SELECT * FROM  match_records WHERE id = $1", matchID)
	if err != nil {
		return types.MatchRecord{}, err
	}

	return matchRecord, nil
}

func (repo PostgresDatabaseStatisticsRepository) UpdatePlayerProfile(newProfile types.PlayerProfile) error {
	statement := `UPDATE
    player_profiles
SET
    id = :id,
    user_id = :user_id,
    playtime = :playtime,
    player_rating_id = :player_rating_id,
    player_statistics_id = :player_statistics_id,
    last_update = :last_update
WHERE
    id = :id`

	_, err := repo.DBConn.NamedExec(statement, &newProfile)
	if err != nil {
		return err
	}

	return nil
}
