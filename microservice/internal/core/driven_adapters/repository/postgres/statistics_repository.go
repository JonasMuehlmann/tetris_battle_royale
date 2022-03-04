package repository

import (
	"log"
	types "microservice/internal/core/types"
)

type PostgresDatabaseStatisticsRepository struct {
	PostgresDatabase
	Logger *log.Logger
}

func (repo PostgresDatabaseStatisticsRepository) GetPlayerProfile(userID int) (types.PlayerProfile, error) {
	return types.PlayerProfile{}, nil
}
