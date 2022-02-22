package repository

type UserRepository interface {
	RegisterUser(username, password, salt string) error
}
