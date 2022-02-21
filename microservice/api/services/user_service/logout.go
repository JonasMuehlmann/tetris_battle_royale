package userService

import (
	"errors"
	"log"
	"microservice/api/common"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func logout(db *sqlx.DB, w http.ResponseWriter, sessionId int) (Session, error) {
	session := Session{}

	res, err := db.Exec("DELETE FROM sessions WHERE id = $1", sessionId)

	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "Failed to end session")
		return Session{}, errors.New("Failed to end session")
	}

	numDeletedSessions, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error: %v", err)
		common.TryWriteResponse(w, "The given session does not exist")
		return Session{}, errors.New("The given session does not exist")
	}

	if numDeletedSessions == 0 {
		common.TryWriteResponse(w, "The given session does not exist")
		return Session{}, errors.New("The given session does not exist")
	}

	common.TryWriteResponse(w, "Successfully ended session")
	return session, nil
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

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

	userId, _ := strconv.ParseInt(vars["userId"], 10, 32)
	logout(db, w, int(userId))
}
