package types

const MatchSize = 10

type Match struct {
	ID      string
	Players [MatchSize]Player
}
