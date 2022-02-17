package userService

import (
	"errors"
	"log"
	"microservice/api/common"
	"net/http"
	"strconv"

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
	sessionIdParam, okSessionId := r.URL.Query()["sessionId"]

	if !okSessionId {
		common.TryWriteResponse(w, "Missing sessionId")
		return
	}

	sessionId, err := strconv.Atoi(sessionIdParam[0])

	db, err := sqlx.Open("postgres", connectionString)

	if err != nil {
		common.TryWriteResponse(w, "Invalid sessionId")
		return
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Printf("Failed to open db: %v", err)

		return
	}

	logout(db, sessionId)
}
