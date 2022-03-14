package drivenPorts

type GameServiceIPCClientPort interface {
	StartGame(userIDList []string) error
	Start(args interface{}) error
}
