package repository

import (
	types "microservice/internal/core/types"
)

type SessionRepositoryPort interface {
	CreateSession(userID string) (string, error)
	GetSession(userID string) (types.Session, error)
	DeleteSession(sessionID string) error
}
