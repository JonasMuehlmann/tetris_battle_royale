package drivenPorts

import (
	"microservice/internal/domain"
)

type SessionPort interface {
	CreateSession(userID int) (int, error)
	GetSession(userID int) (domain.Session, error)
	DeleteSession(sessionID int) error
}
