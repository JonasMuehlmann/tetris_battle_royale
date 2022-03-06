package types

import "net"

type DrivenAdapterGRPCArgs struct {
	Service  interface{}
	Listener net.Listener
}
