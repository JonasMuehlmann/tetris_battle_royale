package gameService

const MatchSize = 2

type Match struct {
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

func (match *Match) HandlePlayerEliminations() {
	var playerID string

	for {
		select {
		case <-match.GameStop:
			// TODO: Handle End of match tasks
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
