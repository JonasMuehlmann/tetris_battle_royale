package gameService

import "time"

const MatchSize = 2

type Match struct {
	ID           string
	Players      [MatchSize]Player
	Acceleration float32
}

func (match Match) Start() {
	go match.setAcceleration()
}

func (match Match) setAcceleration() {
	ticker := time.NewTicker(30 * time.Second)
	for _ = range ticker.C {
		match.Acceleration += 0.1
	}
}
