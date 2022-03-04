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
	GetWinrate(userID int) (float32, error)
	GetNumLosses(userID int) (int, error)
	GetNumWinsAsTop10(userID int) (int, error)
	GetNumWinsAsTop5(userID int) (int, error)
	GetNumWinsAsTop3(userID int) (int, error)
	GetNumWinsAsTop1(userID int) (int, error)

	GetPlayerMatchRecords(userID int) ([]types.MatchRecord, error)
	GetMatchRecord(matchID int) (types.MatchRecord, error)
}
