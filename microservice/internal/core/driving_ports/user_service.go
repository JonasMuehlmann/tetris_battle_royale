package drivingPorts

type UserServicePort interface {
	IsLoggedIn(username string) (string, error)
	Login(username string, password string) (string, error)
	Logout(sessionID string) error
	Register(username string, password string) (string, error)
}
