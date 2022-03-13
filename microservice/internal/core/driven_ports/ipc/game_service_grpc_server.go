package drivenPorts

type GameServiceIPCServerPort interface {
	Start(args interface{}) error
}
