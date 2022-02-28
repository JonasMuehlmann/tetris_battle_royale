package drivenPorts

import (
	types "microservice/internal/core/types"
)

type UserPort interface {
	GetUserFromID(userID int) (types.User, error)
	GetUserFromName(username string) (types.User, error)
	Register(username, password, salt string) (int, error)
}
