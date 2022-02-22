package drivenPorts

import (
	"microservice/internal/domain"
)

type UserPort interface {
	GetUserFromID(userID int) (domain.User, error)
	GetUserFromName(username string) (domain.User, error)
	Register(username, password, salt string) (int, error)
}
