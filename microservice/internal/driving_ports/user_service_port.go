package drivingPorts

type UserServicePorts interface {
	Register(username string, password string)
	Login(username string, password string)
	Logout(sessionID int)
	IsLoggedIn(username string)
}
