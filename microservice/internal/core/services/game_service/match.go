package gameService

const MatchSize = 2

type Match struct {
	// TODO: Refactor this convoluted interaction
	ID                 string
	Players            map[string]Player
	PlayerEliminations chan string
	PlayerCount        int
	NumPlayersAlive    int
	EliminatedPlayers  map[string]bool
	GameService        *GameService
	GameStop           chan bool
}

func (match Match) Start() {
}

func (match *Match) Stop() {
	// TODO: Handle End of match tasks like sending statistics, sending score board data, etc.
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

			for _, player := range match.Players {
				match.GameService.GameAdapter.SendEliminationNotice(player.ID, eliminatedPlayerID)
			}

			if match.NumPlayersAlive == 0 {
				match.GameStop <- true
			}
		}
	}
}
