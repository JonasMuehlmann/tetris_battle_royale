package types

import (
	"container/list"
	"math/rand"
)

const BlockWidth = 4

type Block [BlockWidth][BlockWidth]bool

var BlockShapeL = Block{
	{false, false, false, false},
	{true, false, false, false},
	{true, true, true, true},
	{false, false, false, false},
}
var BlockShapeJ = Block{
	{false, false, false, false},
	{false, false, false, true},
	{true, true, true, true},
	{false, false, false, false},
}
var BlockShapeS = Block{
	{false, true, true, false},
	{true, true, false, false},
	{false, false, false, false},
	{false, false, false, false},
}
var BlockShapeZ = Block{
	{false, false, false, false},
	{true, true, false, false},
	{false, true, true, false},
	{false, false, false, false},
}
var BlockShapeStick = Block{
	{false, false, false, false},
	{true, true, true, true},
	{false, false, false, false},
	{false, false, false, false},
}
var BlockShapeSquare = Block{
	{false, false, false, false},
	{false, true, true, false},
	{false, true, true, false},
	{false, false, false, false},
}
var BlockShapeTriangle = Block{
	{false, false, false, false},
	{false, true, false, false},
	{true, true, true, false},
	{false, false, false, false},
}

var BlockTypes = []Block{
	BlockShapeL,
	BlockShapeJ,
	BlockShapeS,
	BlockShapeZ,
	BlockShapeStick,
	BlockShapeSquare,
	BlockShapeTriangle,
}

func GenerateRandomBlock() Block {
	return BlockTypes[rand.Intn(len(BlockTypes))]
}

const InitialBlockPreviewSize = 5

type BlockPreview struct {
	blockQueue  *list.List
	queueLength int
}

func MakeBlockPreview() BlockPreview {
	var preview = BlockPreview{list.New(), InitialBlockPreviewSize}

	for i := 0; i < InitialBlockPreviewSize; i++ {
		preview.blockQueue.PushBack(GenerateRandomBlock())
	}

	return preview

}

func (preview *BlockPreview) RetrieveBlock() Block {
	preview.blockQueue.PushFront(GenerateRandomBlock())
	return preview.blockQueue.Remove(preview.blockQueue.Back()).(Block)

}

type BlockState struct {
	BlockType           BlockType `json:"block_type"`
	BlockPosition       Position  `json:"block_position"`
	ShadowBlockPosition Position  `json:"shadow_block_position"`
}

type BlockType string

const (
	BlockTypeL        = "l"
	BlockTypeJ        = "j"
	BlockTypeS        = "S"
	BlockTypeZ        = "Z"
	BlockTypeStick    = "stick"
	BlockTypeSquare   = "square"
	BlockTypeTriangle = "triangle"
)

const MatchSize = 2

type Match struct {
	ID      string
	Players [MatchSize]Player
}

func (match Match) Start() {
}

type MoveDirection string

const (
	MoveLeft  = "left"
	MoveRight = "right"
	MoveDown  = "down"
)

type Player struct {
	ID           string
	Score        int
	Playfield    *Playfield
	BlockPreview *BlockPreview
}

type Playfield struct {
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type RotationDirection string

const (
	RotateLeft  = "left"
	RotateRight = "right"
)
