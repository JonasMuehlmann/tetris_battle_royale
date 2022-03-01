package drivenPorts

type GamePort interface {
	ConnectPlayer(userID int, connection interface{}) error
	SendMatchStartNotice(userID int, matchID int) error
}
