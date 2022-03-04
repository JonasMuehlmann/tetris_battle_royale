package drivenPorts

import (
	types "microservice/internal/core/types"
	"time"
)

type StatisticsPort interface {
	GetPlayerProfile(userID int) (types.PlayerProfile, error)
	GetPlayerPlaytime(userID int) (int, error)
	GetPlayerRating(userID int) (int, error)
	GetPlayerProfileLastUpdateTime(userID int) (time.Time, error)

	SetPlayerPlaytime(userID int) error
	SetPlayerRating(rating int) error

	GetPlayerStatistics(userID int) (types.PlayerStatistics, error)
	GetPlayerScore(userID int) (int, error)
	GetPlayerScorePerMinute(userID int) (float32, error)
	GetPlayerWinrate(userID int) (float32, error)
	GetPlayerNumLosses(userID int) (int, error)
	GetPlayerNumWinsAsTop10(userID int) (int, error)
	GetPlayerNumWinsAsTop5(userID int) (int, error)
	GetPlayerNumWinsAsTop3(userID int) (int, error)
	GetPlayerNumWinsAsTop1(userID int) (int, error)

	SetPlayerScore(userID int) error
	SetPlayerScorePerMinute(userID int) error
	SetPlayerWinrate(userID int) error
	SetPlayerNumLosses(userID int) error
	SetPlayerNumWinsAsTop10(userID int) error
	SetPlayerNumWinsAsTop5(userID int) error
	SetPlayerNumWinsAsTop3(userID int) error
	SetPlayerNumWinsAsTop1(userID int) error

	GetPlayerMatchRecords(userID int) ([]types.MatchRecord, error)
	GetMatchRecord(matchID int) (types.MatchRecord, error)

	AddMatchRecord(userID int, record types.MatchRecord) error
}
