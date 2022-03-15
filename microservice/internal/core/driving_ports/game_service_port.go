package drivingPorts

import (
	"microservice/internal/core/types"
	"net"
)

type GameServicePort interface {
	StartGrpcServer(net.Listener) error
	// TODO: Session IDs would be better for security reasons, but using userIDs is a bit simpler
	ConnectPlayer(userID string, connection interface{}) error

	MoveBlock(userID string, matchID string, direction types.MoveDirection) error
	RotateBlock(userID string, matchID string, direction types.RotationDirection) error
	HardDropBlock(userID string, matchID string) error
	ToggleSoftDrop(userID string, matchID string) error
}
