package drivingPorts

type MatchmakingServicePort interface {
	Join(userID int) error
	Leave(userID int) error
}
