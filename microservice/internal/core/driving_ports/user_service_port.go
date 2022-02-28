package drivingPorts

type UserServicePort interface {
	IsLoggedIn(username string) (int, error)
	Login(username string, password string) (int, error)
	Logout(sessionID int) error
	Register(username string, password string) (int, error)
}
