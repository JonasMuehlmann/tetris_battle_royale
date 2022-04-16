package gameService

import (
	types "microservice/internal/core/types"
	"time"
)

const PlayFieldWidth = 2*types.TetrominoWidth + 10
const PlayFieldHeight = 2*types.TetrominoWidth + 20

const InitialGravityTickLength = 500 * time.Millisecond

const SoftdropTickLengthMultiplier float64 = 0.2

type Playfield struct {
	width    int
	height   int
	padding  int
	field    [PlayFieldHeight][PlayFieldWidth]bool
	curTetromino types.Tetromino
	// This is the top left corner of the tetromino
	curTetrominoPosition             types.Position
	curRegularGravityTickLength  time.Duration
	curSoftDropGravityTickLength time.Duration
	gravityTicker                time.Ticker
	gravityStop                  chan bool
	GameStop                     chan bool
	isTetrominoSoftDropping          bool
	softDropScoreMultiplication  int
	hardDropScoreMultiplication  int

	TetrominoPreview TetrominoPreview
	// TODO: Having this much logic in the Playfield and referencing back to
	// the player looks like a code smell
	Player *Player
}

func MakePlayField() Playfield {
	var newField = Playfield{
		width:                        PlayFieldWidth,
		height:                       PlayFieldHeight,
		padding:                      types.TetrominoWidth,
		curRegularGravityTickLength:  InitialGravityTickLength,
		curSoftDropGravityTickLength: time.Duration(float64(InitialGravityTickLength) * SoftdropTickLengthMultiplier),
		gravityTicker:                *time.NewTicker(InitialGravityTickLength),
		gravityStop:                  make(chan bool, 1),
		GameStop:                     make(chan bool, 1),
		softDropScoreMultiplication:  1,
		hardDropScoreMultiplication:  1,
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
	playfield.SpawnNewTetromino(types.GenerateRandomTetromino())
	playfield.EnableGravity()
}

func (playfield *Playfield) StopGame() {
	playfield.GameStop <- true
	playfield.Player.Match.PlayerEliminations <- playfield.Player.ID
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

	//add the score for cleared rows to player
	if len(clearableRows) > 0 {
		scoreToAdd := types.ScoreForRows[len(clearableRows)-1] * playfield.softDropScoreMultiplication * playfield.hardDropScoreMultiplication
		playfield.Player.Score += scoreToAdd

		err := playfield.Player.Match.GameService.GameAdapter.SendScoreGain(playfield.Player.ID, scoreToAdd)

		if err != nil {
			playfield.Player.Match.GameService.Logger.Printf("Error: %v\n", err)
			return
		}
	}
}

func (playfield *Playfield) LockInTetromino() {
	var curFieldPosition *bool

	for row := 0; row < types.TetrominoWidth; row++ {
		for col := 0; col < types.TetrominoWidth; col++ {
			curFieldPosition = &playfield.field[playfield.curTetrominoPosition.Y-row][playfield.curTetrominoPosition.X+col]
			*curFieldPosition = *curFieldPosition || playfield.curTetromino[row][col]
		}
	}

	playfield.tryClearRows()

	playfield.SpawnNewTetromino(types.GenerateRandomTetromino())
	//reset hardDrop score
	playfield.hardDropScoreMultiplication = 1
	if playfield.isTetrominoSoftDropping {
		playfield.softDropScoreMultiplication = 1
	}

	playfield.SpawnNewTetromino(types.GenerateRandomTetromino())
}

func (playfield *Playfield) MoveTetrominoLeft() {
	var newPosition = playfield.curTetrominoPosition
	newPosition.X -= 1

	if playfield.CanTetrominoOccupyPosition(newPosition) {
		playfield.curTetrominoPosition = newPosition
	}

	playfield.UpdateGhostTetrominoPosition()
}

func (playfield *Playfield) MoveTetrominoRight() {
	var newPosition = playfield.curTetrominoPosition
	newPosition.X += 1

	if playfield.CanTetrominoOccupyPosition(newPosition) {
		playfield.curTetrominoPosition = newPosition
	}

	playfield.UpdateGhostTetrominoPosition()
}

func (playfield *Playfield) MoveTetrominoDown() {
	var newPosition = playfield.curTetrominoPosition
	newPosition.Y -= 1

	if playfield.CanTetrominoOccupyPosition(newPosition) {
		playfield.curTetrominoPosition = newPosition
		if playfield.isTetrominoSoftDropping {
			playfield.softDropScoreMultiplication++
		}
	} else {
		playfield.LockInTetromino()
	}

	playfield.UpdateGhostTetrominoPosition()
}

func (playfield *Playfield) HardDropTetromino() {

	var newPosition = playfield.curTetrominoPosition

	for playfield.CanTetrominoOccupyPosition(newPosition) {
		newPosition.Y -= 1

		//add hard drop score for each skipped row
		if playfield.hardDropScoreMultiplication == 1 {
			playfield.hardDropScoreMultiplication++
		} else {
			playfield.hardDropScoreMultiplication += 2
		}
	}

	newPosition.Y += 1
	playfield.curTetrominoPosition = newPosition
	playfield.LockInTetromino()
}
func (playfield *Playfield) UpdateGhostTetrominoPosition() {

	var newPosition = playfield.curTetrominoPosition

	for playfield.CanTetrominoOccupyPosition(newPosition) {
		newPosition.Y -= 1
	}

	newPosition.Y += 1
	//playfield.curGhostTetrominoPosition = newPosition
}

func (playfield *Playfield) ToggleSoftDrop() {
	if playfield.isTetrominoSoftDropping {
		playfield.gravityTicker = *time.NewTicker(playfield.curRegularGravityTickLength)
	} else {
		playfield.gravityTicker = *time.NewTicker(playfield.curSoftDropGravityTickLength)
	}

	playfield.isTetrominoSoftDropping = !playfield.isTetrominoSoftDropping
	playfield.softDropScoreMultiplication = 1
}

func (playfield *Playfield) RotateTetrominoCounterClockwise() {
	transposed := types.Tetromino{}

	// TODO: This should be a separate function
	// Transpose
	for row := 0; row < types.TetrominoWidth; row += 1 {
		for col := 0; col < types.TetrominoWidth; col += 1 {
			transposed[col][row] = playfield.curTetromino[row][col]
		}
	}

	// TODO: This should be a separate function
	// Reverse rows
	for i, j := 0, len(transposed)-1; i < j; i, j = i+1, j-1 {
		transposed[i], transposed[j] = transposed[j], transposed[i]
	}

	playfield.curTetromino = transposed

	playfield.UpdateGhostTetrominoPosition()
}

func (playfield *Playfield) RotateTetrominoClockwise() {
	// Reverse rows
	for i, j := 0, len(playfield.curTetromino)-1; i < j; i, j = i+1, j-1 {
		playfield.curTetromino[i], playfield.curTetromino[j] = playfield.curTetromino[j], playfield.curTetromino[i]
	}

	transposed := playfield.curTetromino

	// Transpose
	for row := 0; row < types.TetrominoWidth; row += 1 {
		for col := 0; col < types.TetrominoWidth; col += 1 {
			transposed[col][row] = playfield.curTetromino[row][col]
		}
	}

	playfield.curTetromino = transposed

	playfield.UpdateGhostTetrominoPosition()
}

// Easily implemented as a kernel (Tetromino) on matrix (padded playfield) test
func (playfield *Playfield) CanTetrominoOccupyPosition(newPosition types.Position) bool {

	for row := 0; row < types.TetrominoWidth; row += 1 {
		for col := 0; col < types.TetrominoWidth; col += 1 {
			// if playfield.curTetromino[row][col] && playfield.field[playfield.curTetrominoPosition.y-row-1][playfield.curTetrominoPosition.x+col] {
			if playfield.curTetromino[row][col] && playfield.field[newPosition.Y-row][newPosition.X+col] {
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
				playfield.MoveTetrominoDown()

			case <-playfield.gravityStop:
				return
			}
		}
	}()
}

func (playfield *Playfield) DisableGravity() {
	playfield.gravityStop <- true
}

func (playfield *Playfield) SpawnNewTetromino(newTetromino types.Tetromino) {
	playfield.curTetromino = newTetromino

	playfield.curTetrominoPosition = types.Position{
		X: (playfield.width)/2 - types.TetrominoWidth/2,
		Y: playfield.height - playfield.padding - 1,
	}

	if !playfield.CanTetrominoOccupyPosition(playfield.curTetrominoPosition) {
		playfield.StopGame()
	}
}
