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

func (repo PostgresDatabaseStatisticsRepository) UpdatePlayerStatistics(newStatistics types.PlayerStatistics) error {
	statement := `UPDATE
    player_statistics
SET
    id = :id,
    score = :score,
    score_per_minute = :score_per_minute,
    wins = :wins,
    losses = :losses,
    winrate = :winrate,
    wins_as_top_10 = :wins_as_top_10,
    wins_as_top_5 = :wins_as_top_5,
    wins_as_top_3 = :wins_as_top_3,
    wins_as_top_1 = :wins_as_top_1
WHERE
    id = :id`

	_, err := repo.DBConn.NamedExec(statement, &newStatistics)
	if err != nil {
		return err
	}

	return nil
}

func (repo PostgresDatabaseStatisticsRepository) AddMatchRecord(record types.MatchRecord) error {

	statement := "INSERT INTO match_records VALUES(:id, :user_id, :win, :score, :length, :start, :rating_change)"

	_, err := repo.DBConn.NamedExec(statement, &record)
	if err != nil {
		return err
	}

	return nil
}
