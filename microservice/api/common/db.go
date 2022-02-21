package common

import (
	"errors"
	"fmt"

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
func GetUserFromId(userId int) (model.User, error) {
	user := model.User{}

	db, err := GetDBConnection()
	if err != nil {
		return model.User{}, nil
	}

	err = db.Get(&user, "SELECT * FROM users WHERE id = $1", userId)
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

func CreateSession(userId int) error {

}

func GetSession(userId int) (model.Session, error) {
}
