package drivenPorts

import "microservice/internal/core/types"

type GamePort interface {
	ConnectPlayer(userID string, connection interface{}) error
	SendMatchStartNotice(userID string, matchID string, opponents []types.Opponent) error

	SendUpdatedBlockState(userID string, newState types.BlockState) error
	SendBlockLockinNotice(userID string) error
	SendRowClearNotice(userID string, rowNum int) error
	SendBlockSpawnNotice(userID string, dequeuedBlock types.BlockType, enqueuedBlock types.BlockType) error
	SendScoreGain(userID string, score int) error
	SendEventNotice(userID string, event string) error
	SendStartBlockPreview(userID string, newPreview []types.Block) error
	SendEliminationNotice(userID string, eliminatedPlayerID string) error
	SendEndOfMatchData(userID string, endOfMatchData types.EndOfMatchData) error
}
