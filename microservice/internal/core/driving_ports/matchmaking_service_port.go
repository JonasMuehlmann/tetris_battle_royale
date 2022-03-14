package drivingPorts

type MatchmakingServicePort interface {
	Join(userID string) error
	Leave(userID string) error
}
