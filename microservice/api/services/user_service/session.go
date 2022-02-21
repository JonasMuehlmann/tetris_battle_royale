package userService

import (
	"errors"
	"log"
	"microservice/api/common"
	"microservice/api/model"
	"time"
)

func createSession(userid int) (model.Session, error) {
	user, err := common.GetUserFromId(userid)
	if err != nil {
		log.Printf("Failed to open db: %v", err)

		return model.Session{}, errors.New("Failed to create session")
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

func getSession(userId int) (model.Session, error) {
	session := model.Session{}

	err := db.Get(&session, "SELECT * FROM sessions WHERE user_id = $1", userId)

	if err != nil {
		return model.Session{}, errors.New("Failed to retrieve session")
	}

	return session, nil
}
