package types

type TetrominoState struct {
	TetrominoPosition Position          `json:"tetromino_position"`
	RotationChange    RotationDirection `json:"rotation_change"`
}
