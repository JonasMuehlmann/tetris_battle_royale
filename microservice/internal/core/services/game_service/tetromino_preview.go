package gameService

import (
	"container/list"
	"microservice/internal/core/types"
)

const InitialTetrominoPreviewSize = 5

type TetrominoPreview struct {
	tetrominoQueue *list.List
	queueLength    int
}

func MakeTetrominoPreview() TetrominoPreview {
	var preview = TetrominoPreview{list.New(), InitialTetrominoPreviewSize}

	for i := 0; i < InitialTetrominoPreviewSize; i++ {
		preview.tetrominoQueue.PushBack(types.GenerateRandomTetromino())
	}

	return preview

}

func (preview *TetrominoPreview) RetrieveTetromino() types.Tetromino {
	preview.tetrominoQueue.PushFront(types.GenerateRandomTetromino())
	return preview.tetrominoQueue.Remove(preview.tetrominoQueue.Back()).(types.Tetromino)

}
