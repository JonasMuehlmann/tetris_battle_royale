package gameService

import "microservice/internal/core/types"

const MatchSize = 2

type Match struct {
	// TODO: Refactor this convoluted interaction
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

		match.GameService.GameAdapter.SendEndOfMatchData(player.ID, endOfMatchData)
	}

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

	// NOTE: This should only do one iteration, IDK any other way to retrieve the sole element of a map
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
