package gameService

import (
	"microservice/internal/core/types"
	"time"
)

const PlayFieldWidth = 2*types.BlockWidth + 10
const PlayFieldHeight = 2*types.BlockWidth + 20

const InitialGravityTickLength = 500 * time.Millisecond

const SoftdropTickLengthMultiplier float64 = 0.2

type Playfield struct {
	width    int
	height   int
	padding  int
	field    [PlayFieldHeight][PlayFieldWidth]bool
	curBlock types.Block
	// This is the top left corner of the block
	curBlockPosition             types.Position
	curRegularGravityTickLength  time.Duration
	curSoftDropGravityTickLength time.Duration
	gravityTicker                time.Ticker
	gravityStop                  chan bool
	GameStop                     chan bool
	isBlockSoftDropping          bool

	BlockPreview BlockPreview
}
