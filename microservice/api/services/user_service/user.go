package userService

import "github.com/jmoiron/sqlx"

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	PwHash   string `db:"pw_hash"`
	Salt     string `db:"salt"`
}

func getUserFromId(db *sqlx.DB, userId int) (User, error) {
	user := User{}

	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func getUserFromName(db *sqlx.DB, username string) (User, error) {

	user := User{}

	err := db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return user, err
	}

	return user, nil
}
