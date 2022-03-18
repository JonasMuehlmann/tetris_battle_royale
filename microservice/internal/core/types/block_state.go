package types

type BlockState struct {
	BlockPosition  Position          `json:"block_position"`
	RotationChange RotationDirection `json:"rotation_change"`
}
