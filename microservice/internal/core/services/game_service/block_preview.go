package gameService

import (
	"container/list"
	"microservice/internal/core/types"
)

const InitialBlockPreviewSize = 5

type BlockPreview struct {
	blockQueue  *list.List
	queueLength int
}

func MakeBlockPreview() BlockPreview {
	var preview = BlockPreview{list.New(), InitialBlockPreviewSize}

	for i := 0; i < InitialBlockPreviewSize; i++ {
		preview.blockQueue.PushBack(types.GenerateRandomBlock())
	}

	return preview

}

func (preview *BlockPreview) RetrieveBlock() types.Block {
	preview.blockQueue.PushFront(types.GenerateRandomBlock())
	return preview.blockQueue.Remove(preview.blockQueue.Back()).(types.Block)

}
