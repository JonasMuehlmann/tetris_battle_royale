package internal

import "time"

const PlayFieldWidth = 2*BlockWidth + 10
const PlayFieldHeight = 2*BlockWidth + 20

// const InitialGravityTickLength = 2 * time.Second
const InitialGravityTickLength = 500 * time.Millisecond

const SoftdropTickLengthMultiplier float64 = 0.2

type Playfield struct {
	width    int
	height   int
	padding  int
	field    [PlayFieldHeight][PlayFieldWidth]bool
	curBlock Block
	// This is the top left corner of the block
	curBlockPosition             Position
	curGhostBlockPosition        Position
	curRegularGravityTickLength  time.Duration
	curSoftDropGravityTickLength time.Duration
	gravityTicker                time.Ticker
	gravityStop                  chan bool
	GameStop                     chan bool
	isBlockSoftDropping          bool
	score                        int
}

func MakePlayField() Playfield {
	var newField = Playfield{
		width:                        PlayFieldWidth,
		height:                       PlayFieldHeight,
		padding:                      BlockWidth,
		curRegularGravityTickLength:  InitialGravityTickLength,
		curSoftDropGravityTickLength: time.Duration(float64(InitialGravityTickLength) * SoftdropTickLengthMultiplier),
		gravityTicker:                *time.NewTicker(InitialGravityTickLength),
		gravityStop:                  make(chan bool, 1),
		GameStop:                     make(chan bool, 1),
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
	playfield.SpawnNewBlock(GenerateRandomBlock())
	playfield.EnableGravity()
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

	for row := 0; row < BlockWidth; row++ {
		for col := 0; col < BlockWidth; col++ {
			curFieldPosition = &playfield.field[playfield.curBlockPosition.y-row][playfield.curBlockPosition.x+col]
			*curFieldPosition = *curFieldPosition || playfield.curBlock[row][col]
		}
	}

	playfield.tryClearRows()

	playfield.SpawnNewBlock(GenerateRandomBlock())
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
	newPosition.y -= 1

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
		newPosition.y -= 1
	}

	newPosition.y += 1
	playfield.curBlockPosition = newPosition
	playfield.LockInBlock()
}
func (playfield *Playfield) UpdateGhostBlockPosition() {

	var newPosition = playfield.curBlockPosition

	for playfield.CanBlockOccupyPosition(newPosition) {
		newPosition.y -= 1
	}

	newPosition.y += 1
	playfield.curGhostBlockPosition = newPosition
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
	transposed := Block{}

	// TODO: This should be a separate function
	// Transpose
	for row := 0; row < BlockWidth; row += 1 {
		for col := 0; col < BlockWidth; col += 1 {
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
	for row := 0; row < BlockWidth; row += 1 {
		for col := 0; col < BlockWidth; col += 1 {
			transposed[col][row] = playfield.curBlock[row][col]
		}
	}

	playfield.curBlock = transposed

	playfield.UpdateGhostBlockPosition()
}

// Easily implemented as a kernel (Block) on matrix (padded playfield) test
func (playfield *Playfield) CanBlockOccupyPosition(newPosition Position) bool {

	for row := 0; row < BlockWidth; row += 1 {
		for col := 0; col < BlockWidth; col += 1 {
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

func (playfield *Playfield) SpawnNewBlock(newBlock Block) {
	playfield.curBlock = newBlock

	playfield.curBlockPosition = Position{
		x: (playfield.width)/2 - BlockWidth/2,
		y: playfield.height - playfield.padding - 1,
	}

	if !playfield.CanBlockOccupyPosition(playfield.curBlockPosition) {
		playfield.StopGame()
	}
}

func isPointInRect(px int, py int, rectTopLeft Position, w int, h int) bool {
	// w and h are total lengths, but adding them would move by this many from the first point, hence the -1 and +1
	return px >= rectTopLeft.x && px <= rectTopLeft.x+w-1 && py <= rectTopLeft.y && py >= rectTopLeft.y-h+1
}

func (playfield *Playfield) PrintPlayField() {
	// printHorizontalLine(playfield.width*2 + 4)
	print("\n")

	for row := playfield.height - playfield.padding - 1; row >= playfield.padding; row-- {
		// print("┆")
		for col := playfield.padding; col < playfield.width-playfield.padding; col++ {
			var cell = playfield.field[row][col]

			if isPointInRect(col, row, playfield.curBlockPosition, BlockWidth, BlockWidth) && (playfield.curBlock[playfield.curBlockPosition.y-row][col-playfield.curBlockPosition.x]) {
				print("\033[38;2;255;180;0m██\033[0m")

			} else if isPointInRect(col, row, playfield.curGhostBlockPosition, BlockWidth, BlockWidth) && (playfield.curBlock[playfield.curGhostBlockPosition.y-row][col-playfield.curGhostBlockPosition.x]) {
				print("\033[38;2;100;100;100m██\033[0m")
			} else if cell {
				print("\033[38;2;255;255;255m██\033[0m")
			} else {
				print("  ")
			}

		}
		// print(" ")
		print(row - playfield.padding)
		// print(" ")
		// print(row)
		print("\n")

	}
	print(" ")
	for col := playfield.padding; col < playfield.width-playfield.padding; col++ {
		print(col - playfield.padding)
		print(" ")
	}
	// print(" Logical field size\n")
	// print(" ")
	// for col := playfield.padding; col < playfield.width-playfield.padding; col++ {
	// 	print(col)
	// 	if col >= 10 {
	// 		print(" ")
	// 	} else {
	// 		print("  ")
	// 	}
	// }
	// print("    Internal field size\n")
	print("\n")
}
