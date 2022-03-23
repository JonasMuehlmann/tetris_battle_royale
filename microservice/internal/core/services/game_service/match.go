package gameService

const MatchSize = 2

type Match struct {
	ID      string
	Players [MatchSize]Player
}

func (match Match) Start() {

}
