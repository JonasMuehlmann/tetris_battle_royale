package drivenPorts

import "microservice/internal/core/types"

type GamePort interface {
	ConnectPlayer(userID string, connection interface{}) error
	SendMatchStartNotice(userID string, matchID string, opponents []types.Opponent) error

	SendUpdatedTetrominoState(userID string, newState types.TetrominoState) error
	SendTetrominoLockinNotice(userID string) error
	SendRowClearNotice(userID string, rowNum int) error
	SendTetrominoSpawnNotice(userID string, dequeuedTetromino types.TetrominoName, enqueuedTetromino types.TetrominoName) error
	SendScoreGain(userID string, score int) error
	SendEventNotice(userID string, event string) error
	SendStartTetrominoPreview(userID string, newPreview []types.Tetromino) error
	SendEliminationNotice(userID string, eliminatedPlayerID string) error
	SendEndOfMatchData(userID string, endOfMatchData types.EndOfMatchData) error
}
