package gameService

const MatchSize = 2

type Match struct {
	ID      string
	Players map[string]Player
}

func (match Match) Start() {
}
