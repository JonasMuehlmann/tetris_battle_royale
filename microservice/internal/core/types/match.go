package types

const MatchSize = 10

type Match struct {
	ID      int
	Players [MatchSize]Player
}
