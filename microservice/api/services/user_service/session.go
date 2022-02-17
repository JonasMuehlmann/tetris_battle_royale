package userService

import (
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type Session struct {
	ID           int       `db:"id"`
	UserID       int       `db:"user_id"`
	CreationTime time.Time `db:"creation_time"`
}

func createSession(db *sqlx.DB, userid int) (Session, error) {
	user, err := getUserFromId(db, userid)
	if err != nil {
		log.Printf("Failed to open db: %v", err)

		return Session{}, errors.New("Failed to create session")
	}

	session := Session{-1, user.ID, time.Now()}

	err = db.QueryRow("INSERT INTO sessions(user_id, creation_time) VALUES($1, $2) RETURNING id", session.UserID, session.CreationTime).Scan(&session.ID)
	if err != nil {
		log.Printf("Error: %v", err)

		return session, err
	}

	log.Println("Created new session")
	return session, nil
}

func getSession(db *sqlx.DB, userId int) (Session, error) {
	session := Session{}

	err := db.Get(&session, "SELECT * FROM sessions WHERE user_id = $1", userId)

	if err != nil {
		return Session{}, errors.New("Failed to retrieve session")
	}

	return session, nil
}
