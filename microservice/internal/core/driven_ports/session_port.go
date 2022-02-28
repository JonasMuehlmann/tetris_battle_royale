package drivenPorts

import (
	types "microservice/internal/core/types"
)

type SessionPort interface {
	CreateSession(userID int) (int, error)
	GetSession(userID int) (types.Session, error)
	DeleteSession(sessionID int) error
}
