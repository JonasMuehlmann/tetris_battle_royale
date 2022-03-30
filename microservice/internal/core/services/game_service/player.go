package gameService

type Player struct {
	ID        string
	Score     int
	Playfield Playfield
	Match     *Match
}
