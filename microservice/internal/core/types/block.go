package types

import "math/rand"

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
