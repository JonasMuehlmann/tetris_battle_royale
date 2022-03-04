package drivenPorts

import (
	types "microservice/internal/core/types"
)

type StatisticsPort interface {
	GetPlayerProfile(userID string) (types.PlayerProfile, error)
	// GetPlayerPlaytime(userID string) (int, error)
	// GetPlayerRating(userID string) (int, error)
	// GetPlayerProfileLastUpdateTime(userID string) (time.Time, error)

	// SetPlayerPlaytime(userID string) error
	// SetPlayerRating(rating int) error

	// GetPlayerStatistics(userID string) (types.PlayerStatistics, error)
	// GetPlayerScore(userID string) (int, error)
	// GetPlayerScorePerMinute(userID string) (float32, error)
	// GetPlayerWinrate(userID string) (float32, error)
	// GetPlayerNumLosses(userID string) (int, error)
	// GetPlayerNumWinsAsTop10(userID string) (int, error)
	// GetPlayerNumWinsAsTop5(userID string) (int, error)
	// GetPlayerNumWinsAsTop3(userID string) (int, error)
	// GetPlayerNumWinsAsTop1(userID string) (int, error)

	// SetPlayerScore(userID string) error
	// SetPlayerScorePerMinute(userID string) error
	// SetPlayerWinrate(userID string) error
	// SetPlayerNumLosses(userID string) error
	// SetPlayerNumWinsAsTop10(userID string) error
	// SetPlayerNumWinsAsTop5(userID string) error
	// SetPlayerNumWinsAsTop3(userID string) error
	// SetPlayerNumWinsAsTop1(userID string) error

	// GetPlayerMatchRecords(userID string) ([]types.MatchRecord, error)
	// GetMatchRecord(matchID int) (types.MatchRecord, error)

	// AddMatchRecord(userID string, record types.MatchRecord) error
}
