package userService

import (
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func logout(db *sqlx.DB, sessionId int) (Session, error) {
	session := Session{}

	err := db.Get(&session, "DELETE FROM sessions WHERE id = $1", sessionId)

	if err != nil {
		return Session{}, errors.New("Failed to end session")
	}

	return session, nil
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
}
