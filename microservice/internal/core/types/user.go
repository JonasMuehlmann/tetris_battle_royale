package types

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	PwHash   string `db:"pw_hash"`
	Salt     string `db:"salt"`
}
