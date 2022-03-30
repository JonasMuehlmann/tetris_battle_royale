package gameService

import (
	"microservice/internal/core/types"
	"time"

	"github.com/google/uuid"
)

const MatchSize = 2

// REFACTOR: Refactor this convoluted interaction betweenn the match, playfield and game service
type Match struct {
	ID                 string
	Players            map[string]Player
	PlayerEliminations chan string
	PlayerCount        int
	NumPlayersAlive    int
	EliminatedPlayers  map[string]bool
	AlivePlayers       map[string]bool
	GameService        *GameService
	GameStop           chan bool
}

func (match Match) Start() {
}

func (match *Match) Stop() {
	endOfMatchData, err := match.generateEndOfMatchData()
	if err != nil {
		match.GameService.Logger.Printf("Error: %v\n", err)
	}

	for _, player := range match.Players {

		err := match.GameService.GameAdapter.SendEndOfMatchData(player.ID, endOfMatchData)
		if err != nil {
			match.GameService.Logger.Printf("Error: %v\n", err)
		}

		// FIX: Fill in default initialized statistics
		matchRecord := types.MatchRecord{
			ID:           uuid.NewString(),
			UserID:       player.ID,
			Win:          false,
			WinKind:      0,
			Score:        player.Score,
			Start:        time.Time{},
			Length:       0,
			RatingChange: 0,
		}

		err = match.GameService.StatisticsIPCClient.AddMatchRecord(matchRecord)
		if err != nil {
			match.GameService.Logger.Printf("Error: %v\n", err)
		}
	}

	// HACK: This feels wrong, but it might work for now
	delete(match.GameService.Matches, match.ID)

}

func (match *Match) HandlePlayerEliminations() {
	var eliminatedPlayerID string

	for {
		select {
		case <-match.GameStop:
			match.Stop()

			return
		case eliminatedPlayerID = <-match.PlayerEliminations:
			match.NumPlayersAlive--
			match.EliminatedPlayers[eliminatedPlayerID] = true
			match.AlivePlayers[eliminatedPlayerID] = false

			for _, player := range match.Players {
				match.GameService.GameAdapter.SendEliminationNotice(player.ID, eliminatedPlayerID)
			}

			if match.NumPlayersAlive == 0 {
				match.GameStop <- true
			}
		}
	}
}

func (match *Match) generateEndOfMatchData() (types.EndOfMatchData, error) {
	endOfMatchData := types.EndOfMatchData{}

	// HACK: IDK any other way to retrieve the sole element of a map
	for playerID := range match.AlivePlayers {
		endOfMatchData.Scorboard.Winner = playerID

		break
	}

	for _, player := range match.Players {
		playerPerformance := types.PlayerPerformance{player.ID, player.Score}
		endOfMatchData.Scorboard.PlayerPerformances = append(endOfMatchData.Scorboard.PlayerPerformances, playerPerformance)
	}

	return endOfMatchData, nil
}
