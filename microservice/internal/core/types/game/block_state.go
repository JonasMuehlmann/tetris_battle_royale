package types

type BlockState struct {
	BlockType           BlockType `json:"block_type"`
	BlockPosition       Position  `json:"block_position"`
	ShadowBlockPosition Position  `json:"shadow_block_position"`
}
