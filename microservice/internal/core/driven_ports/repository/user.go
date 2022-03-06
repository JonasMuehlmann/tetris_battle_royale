package repository

import (
	types "microservice/internal/core/types"
)

type UserRepositoryPort interface {
	GetUserFromID(userID string) (types.User, error)
	GetUserFromName(username string) (types.User, error)
	Register(username, password, salt string) (string, error)
}
