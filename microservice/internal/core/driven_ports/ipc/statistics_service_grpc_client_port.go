package drivenPorts

import "microservice/internal/core/types"

type StatisticsServiceIPCClientPort interface {
	AddMatchRecord(record types.MatchRecord) error
	Start(args interface{}) error
}
