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
