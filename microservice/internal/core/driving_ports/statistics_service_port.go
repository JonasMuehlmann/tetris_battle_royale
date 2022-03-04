package drivingPorts

import (
	"microservice/internal/core/types"
	"time"
)

type StatisticsServicePort interface {
	GetPlayerProfile(userID int) (types.PlayerProfile, error)
	GetPlayerPlaytime(userID int) (int, error)
	GetPlayerRating(userID int) (int, error)
	GetPlayerProfileLastUpdateTime(userID int) (time.Time, error)

	GetPlayerStatistics(userID int) (types.PlayerStatistics, error)
	GetPlayerScore(userID int) (int, error)
	GetPlayerScorePerMinute(userID int) (float32, error)
	GetPlayerWinrate(userID int) (float32, error)
	GetPlayerNumLosses(userID int) (int, error)
	GetPlayerNumWinsAsTop10(userID int) (int, error)
	GetPlayerNumWinsAsTop5(userID int) (int, error)
	GetPlayerNumWinsAsTop3(userID int) (int, error)
	GetPlayerNumWinsAsTop1(userID int) (int, error)

	GetPlayerMatchRecords(userID int) ([]types.MatchRecord, error)
	GetMatchRecord(matchID int) (types.MatchRecord, error)
}
