package drivenPorts

import (
	types "microservice/internal/core/types"
)

type SessionPort interface {
	CreateSession(userID string) (string, error)
	GetSession(userID string) (types.Session, error)
	DeleteSession(sessionID string) error
}
