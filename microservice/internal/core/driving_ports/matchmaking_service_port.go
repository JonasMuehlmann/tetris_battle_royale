package drivingPorts

type MatchmakingServicePort interface {
	Join(userID int) (int, error)
	Leave(userID int) error
}
