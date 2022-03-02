package drivingPorts

type GameServicePort interface {
	ConnectPlayer(userID int, connection interface{}) error
}
