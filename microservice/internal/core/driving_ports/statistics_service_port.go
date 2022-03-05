package drivingPorts

import (
	"microservice/internal/core/types"
)

type StatisticsServicePort interface {
	GetPlayerProfile(userID string) (types.PlayerProfile, error)
	// GetPlayerPlaytime(userID string) (int, error)
	// GetPlayerRating(userID string) (int, error)
	// GetPlayerProfileLastUpdateTime(userID string) (time.Time, error)

	GetPlayerStatistics(userID string) (types.PlayerStatistics, error)
	// GetPlayerScore(userID string) (int, error)
	// GetPlayerScorePerMinute(userID string) (float32, error)
	// GetPlayerWinrate(userID string) (float32, error)
	// GetPlayerNumLosses(userID string) (int, error)
	// GetPlayerNumWinsAsTop10(userID string) (int, error)
	// GetPlayerNumWinsAsTop5(userID string) (int, error)
	// GetPlayerNumWinsAsTop3(userID string) (int, error)
	// GetPlayerNumWinsAsTop1(userID string) (int, error)

	GetMatchRecords(userID string) ([]types.MatchRecord, error)
	GetMatchRecord(matchID string) (types.MatchRecord, error)
}
