package drivenPorts

import "microservice/internal/core/types"

type GamePort interface {
	ConnectPlayer(userID string, connection interface{}) error
	SendMatchStartNotice(userID string, matchID string) error

	SendUpdatedBlockState(userID string, newState types.BlockState) error
	SendBlockLockinNotice(userID string) error
	SendRowClearNotice(userID string, rowNum int) error
	SendBlockSpawnNotice(userID string, newBlock types.BlockType) error
	SendScoreGain(userID string, score int) error
	SendEventNotice(userID string, event string) error
}
