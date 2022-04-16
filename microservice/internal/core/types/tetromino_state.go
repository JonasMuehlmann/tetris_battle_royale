package types

type TetrominoState struct {
	TetrominoPosition Position          `json:"tetrominoPosition"`
	RotationChange    RotationDirection `json:"rotationChange"`
}
