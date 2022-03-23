package gameService

const MatchSize = 2

type Match struct {
	// TODO: Refactor this convoluted interaction
	ID                 string
	Players            [MatchSize]Player
	PlayerEliminations chan string
	PlayerCount        int
	NumPlayersAlive    int
	EliminatedPlayers  map[string]*Player
	GameService        *GameService
	GameStop           chan bool
}

func (match Match) Start() {
}

func (match *Match) Stop() {
	// TODO: Handle End of match tasks like sending statistics, sending score board data, etc.
}

func (match *Match) HandlePlayerEliminations() {
	var playerID string

	for {
		select {
		case <-match.GameStop:
			match.Stop()
			return
		case playerID = <-match.PlayerEliminations:
			match.NumPlayersAlive--
			match.EliminatedPlayers[playerID] = &match.Players[playerID]

			if match.NumPlayersAlive == 0 {
				match.GameStop <- true
			}
		}
	}
}
