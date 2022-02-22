package userService

import (
	"log"
	"microservice/api/common"
	"microservice/api/model"
)

func createSession(userID int) (model.Session, error) {

	session, err := common.GetSession(userID)
	if err != nil {
		return model.Session{}, nil
	}

	log.Println("Created new session")

	return session, nil
}

func getSession(userID int) (model.Session, error) {
	return common.GetSession(userID)
}
