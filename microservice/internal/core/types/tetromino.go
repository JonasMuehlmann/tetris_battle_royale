package types

import "math/rand"

const TetrominoWidth = 4

type Tetromino [TetrominoWidth][TetrominoWidth]bool


var Tetrominos = map[TetrominoName]Tetromino{
	TetrominoL : 
	{
		{false, false, false, false},
		{true, false, false, false},
		{true, true, true, true},
		{false, false, false, false},
	},
	TetrominoJ :
	{
		{false, false, false, false},
		{false, false, false, true},
		{true, true, true, true},
		{false, false, false, false},
	},
	TetrominoS :
	{
		{false, true, true, false},
		{true, true, false, false},
		{false, false, false, false},
		{false, false, false, false},
	},
	TetrominoZ :
	{
		{false, false, false, false},
		{true, true, false, false},
		{false, true, true, false},
		{false, false, false, false},
	},
	TetrominoI :
	{
		{false, false, false, false},
		{true, true, true, true},
		{false, false, false, false},
		{false, false, false, false},
	},
	TetrominoO :
	{
		{false, false, false, false},
		{false, true, true, false},
		{false, true, true, false},
		{false, false, false, false},
	},
	TetrominoT :
	{
		{false, false, false, false},
		{false, true, false, false},
		{true, true, true, false},
		{false, false, false, false},
	},
}

var tetrominoNames = []TetrominoName{TetrominoL, TetrominoJ, TetrominoS, TetrominoZ, TetrominoI, TetrominoO, TetrominoT}

func GenerateRandomTetromino() Tetromino {
	// HACK: This is not so clean, but it works and probably won't ever change.
	return  Tetrominos[tetrominoNames[rand.Intn(len(Tetrominos))]]
}
