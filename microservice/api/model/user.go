package model

type User struct {
	ID       int    `db:"ID"`
	Username string `db:"username"`
	PwHash   string `db:"pw_hash"`
	Salt     string `db:"salt"`
}
