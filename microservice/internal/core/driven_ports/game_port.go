package drivenPorts

type GamePort interface {
	ConnectPlayer(userID string, connection interface{}) error
	SendMatchStartNotice(userID string, matchID string) error
}
