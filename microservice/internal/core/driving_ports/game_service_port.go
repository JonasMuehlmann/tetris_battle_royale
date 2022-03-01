package drivingPorts

import "net"

type GameServicePort interface {
	StartGrpcServer(net.Listener) error
	ConnectPlayer(userID int, connection interface{}) error
}
