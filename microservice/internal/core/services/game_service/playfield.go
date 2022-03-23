package gameService

import (
	"math"
	types "microservice/internal/core/types"
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

	BlockPreview types.BlockPreview

	Acceleration float64
}

func MakePlayField() Playfield {
	var newField = Playfield{
		width:                        PlayFieldWidth,
		height:                       PlayFieldHeight,
		padding:                      types.BlockWidth,
		curRegularGravityTickLength:  InitialGravityTickLength,
		curSoftDropGravityTickLength: time.Duration(float64(InitialGravityTickLength) * SoftdropTickLengthMultiplier),
		gravityTicker:                *time.NewTicker(InitialGravityTickLength),
		gravityStop:                  make(chan bool, 1),
		GameStop:                     make(chan bool, 1),
		Acceleration:                 1,
	}

	for row := 0; row < newField.height; row++ {
		for col := 0; col < newField.width; col++ {
			if row < newField.padding || row >= newField.height-newField.padding || col < newField.padding || col >= newField.width-newField.padding {
				newField.field[row][col] = true
			}
		}
	}

	return newField
}

func (playfield *Playfield) StartGame() {
	playfield.SpawnNewBlock(types.GenerateRandomBlock())
	playfield.EnableGravity()
	go playfield.setAcceleration()
}

func (playfield *Playfield) StopGame() {
	playfield.GameStop <- true
	playfield.DisableGravity()
}

func (playfield *Playfield) findClearableRows() []int {
	clearableRows := make([]int, 0, 5)
	isRowFull := true

	for row := playfield.padding; row < playfield.height-playfield.padding; row++ {
		for col := playfield.padding; col < playfield.width-playfield.padding; col++ {
			isRowFull = isRowFull && playfield.field[row][col]
		}

		if isRowFull {
			clearableRows = append(clearableRows, row)
		}

		isRowFull = true
	}

	return clearableRows
}

func (playfield *Playfield) tryClearRows() {
	clearableRows := playfield.findClearableRows()

	for clearableRow := range clearableRows {
		// Iterate field upwards from the first (lowest) clearable row
		// -1 to avoid pulling padding row when doing row+1 further down
		for row := clearableRow; row < playfield.height-playfield.padding-1; row++ {
			// Overwrite current row with the one above
			playfield.field[row] = playfield.field[row+1]
		}
	}
}

func (playfield *Playfield) LockInBlock() {
	var curFieldPosition *bool

	for row := 0; row < types.BlockWidth; row++ {
		for col := 0; col < types.BlockWidth; col++ {
			curFieldPosition = &playfield.field[playfield.curBlockPosition.y-row][playfield.curBlockPosition.x+col]
			*curFieldPosition = *curFieldPosition || playfield.curBlock[row][col]
		}
	}

	playfield.tryClearRows()

	playfield.SpawnNewBlock(types.GenerateRandomBlock())
}

func (playfield *Playfield) MoveBlockLeft() {
	var newPosition = playfield.curBlockPosition
	newPosition.x -= 1

	if playfield.CanBlockOccupyPosition(newPosition) {
		playfield.curBlockPosition = newPosition
	}

	playfield.UpdateGhostBlockPosition()
}

func (playfield *Playfield) MoveBlockRight() {
	var newPosition = playfield.curBlockPosition
	newPosition.x += 1

	if playfield.CanBlockOccupyPosition(newPosition) {
		playfield.curBlockPosition = newPosition
	}

	playfield.UpdateGhostBlockPosition()
}

func (playfield *Playfield) MoveBlockDown() {
	var newPosition = playfield.curBlockPosition
	newPosition.y -= math.Round(1 * playfield.Acceleration)

	if playfield.CanBlockOccupyPosition(newPosition) {
		playfield.curBlockPosition = newPosition
	} else {
		playfield.LockInBlock()
	}

	playfield.UpdateGhostBlockPosition()
}

func (playfield *Playfield) HardDropBlock() {

	var newPosition = playfield.curBlockPosition

	for playfield.CanBlockOccupyPosition(newPosition) {
		newPosition.y -= math.Round(1 * playfield.Acceleration)
	}

	newPosition.y += math.Round(1 * playfield.Acceleration)
	playfield.curBlockPosition = newPosition
	playfield.LockInBlock()
}
func (playfield *Playfield) UpdateGhostBlockPosition() {

	var newPosition = playfield.curBlockPosition

	for playfield.CanBlockOccupyPosition(newPosition) {
		newPosition.y -= 1
	}

	newPosition.y += 1
	//playfield.curGhostBlockPosition = newPosition
}

func (playfield *Playfield) ToggleSoftDrop() {
	if playfield.isBlockSoftDropping {
		playfield.gravityTicker = *time.NewTicker(playfield.curRegularGravityTickLength)
	} else {
		playfield.gravityTicker = *time.NewTicker(playfield.curSoftDropGravityTickLength)
	}

	playfield.isBlockSoftDropping = !playfield.isBlockSoftDropping
}

func (playfield *Playfield) RotateBlockCounterClockwise() {
	transposed := types.Block{}

	// TODO: This should be a separate function
	// Transpose
	for row := 0; row < types.BlockWidth; row += 1 {
		for col := 0; col < types.BlockWidth; col += 1 {
			transposed[col][row] = playfield.curBlock[row][col]
		}
	}

	// TODO: This should be a separate function
	// Reverse rows
	for i, j := 0, len(transposed)-1; i < j; i, j = i+1, j-1 {
		transposed[i], transposed[j] = transposed[j], transposed[i]
	}

	playfield.curBlock = transposed

	playfield.UpdateGhostBlockPosition()
}

func (playfield *Playfield) RotateBlockClockwise() {
	// Reverse rows
	for i, j := 0, len(playfield.curBlock)-1; i < j; i, j = i+1, j-1 {
		playfield.curBlock[i], playfield.curBlock[j] = playfield.curBlock[j], playfield.curBlock[i]
	}

	transposed := playfield.curBlock

	// Transpose
	for row := 0; row < types.BlockWidth; row += 1 {
		for col := 0; col < types.BlockWidth; col += 1 {
			transposed[col][row] = playfield.curBlock[row][col]
		}
	}

	playfield.curBlock = transposed

	playfield.UpdateGhostBlockPosition()
}

// Easily implemented as a kernel (Block) on matrix (padded playfield) test
func (playfield *Playfield) CanBlockOccupyPosition(newPosition types.Position) bool {

	for row := 0; row < types.BlockWidth; row += 1 {
		for col := 0; col < types.BlockWidth; col += 1 {
			// if playfield.curBlock[row][col] && playfield.field[playfield.curBlockPosition.y-row-1][playfield.curBlockPosition.x+col] {
			if playfield.curBlock[row][col] && playfield.field[newPosition.y-row][newPosition.x+col] {
				return false
			}
		}
	}

	return true
}

func (playfield *Playfield) EnableGravity() {
	go func() {
		for {
			select {
			case <-playfield.gravityTicker.C:
				playfield.MoveBlockDown()

			case <-playfield.gravityStop:
				return
			}
		}
	}()
}

func (playfield *Playfield) DisableGravity() {
	playfield.gravityStop <- true
}

func (playfield *Playfield) SpawnNewBlock(newBlock types.Block) {
	playfield.curBlock = newBlock

	playfield.curBlockPosition = types.Position{
		x: (playfield.width)/2 - types.BlockWidth/2,
		y: playfield.height - playfield.padding - 1,
	}

	if !playfield.CanBlockOccupyPosition(playfield.curBlockPosition) {
		playfield.StopGame()
	}
}

func (playfield *Playfield) setAcceleration() {

	//TODO: should probably be configurable
	ticker := time.NewTicker(30 * time.Second)
	for _ = range ticker.C {
		playfield.Acceleration += 0.1
	}
}
