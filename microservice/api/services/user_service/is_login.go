package userService

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
)

func isLoggedIn(db *sqlx.DB, w http.ResponseWriter, r *http.Request, username string) {
	user, err := getUserFromName(db, username)

	if err != nil {
		log.Printf("Error: %v", err)

		_, err = w.Write([]byte("Failed to retrieve login status"))

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}

	session, err := getSession(db, user.ID)

	// TODO: Check if session expired

	if err != nil {
		log.Printf("Error: %v", err)

		_, err = w.Write([]byte("User not logged in"))

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	} else {
		log.Printf("Error: %v", err)

		_, err = w.Write([]byte("User logged in with id " + strconv.FormatInt(int64(session.ID), 10)))

		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}

	}
}

func IsLoginHandler(w http.ResponseWriter, r *http.Request) {
}
