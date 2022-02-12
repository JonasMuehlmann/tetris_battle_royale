package internal

import (
	"container/list"
)

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
