package types

type User struct {
	ID       string `db:"id"`
	Username string `db:"username"`
	PwHash   string `db:"pw_hash"`
	Salt     string `db:"salt"`
}
