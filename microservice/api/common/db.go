package common

import (
	"errors"
	"fmt"
	"log"
	"time"

	"microservice/api/model"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5432
	username = "postgres"
	dbname   = "prod"
)

var connectionString = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, username, dbname)

func GetDBConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", connectionString)

	defer db.Close()

	err = db.Ping()

	if err != nil {
		return nil, errors.New("Failed to open db: " + err.Error())
	}

	return db, nil

}
func GetUserFromID(userID int) (model.User, error) {
	user := model.User{}

	db, err := GetDBConnection()
	if err != nil {
		return model.User{}, nil
	}

	err = db.Get(&user, "SELECT * FROM users WHERE ID = $1", userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserFromName(username string) (model.User, error) {
	user := model.User{}

	db, err := GetDBConnection()
	if err != nil {
		return user, err
	}

	err = db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateSession(userID int) (int, error) {

	user, err := GetUserFromID(userID)
	session := model.Session{ID: -1, UserID: user.ID, CreationTime: time.Now()}

	db, err := GetDBConnection()
	if err != nil {
		return 0, err
	}

	err = db.QueryRow("INSERT INTO sessions(user_ID, creation_time) VALUES($1, $2) RETURNING ID", session.UserID, session.CreationTime).Scan(&session.ID)
	if err != nil {
		log.Printf("Error: %v", err)

		return 0, err
	}

	return session.ID, nil
}

func GetSession(userID int) (model.Session, error) {
	session := model.Session{}

	db, err := GetDBConnection()
	if err != nil {
		return session, err
	}

	err = db.Get(&session, "SELECT * FROM sessions WHERE user_ID = $1", userID)
	if err != nil {
		return model.Session{}, errors.New("Failed to retrieve session")
	}

	return session, nil
}

func DeleteSession(sessionID int) error {
	db, err := GetDBConnection()
	if err != nil {
		return err
	}

	res, err := db.Exec("DELETE FROM sessions WHERE ID = $1", sessionID)
	if err != nil {
		return err
	}

	numDeletedSessions, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if numDeletedSessions == 0 {
		return errors.New("Session does not exist")
	}

	return nil
}

func Register(username, password, salt string) (int, error) {
	var userID int

	db, err := GetDBConnection()
	if err != nil {
		return 0, err
	}
	err = db.QueryRow("INSERT INTO users(username, pw_hash, salt) VALUES($1, $2, $3) RETURNING ID", username, string(password), string(salt)).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
