package types

type Player struct {
	ID           string
	Score        int
	Playfield    *Playfield
	BlockPreview *BlockPreview
}
