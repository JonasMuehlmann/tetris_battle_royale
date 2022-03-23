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
