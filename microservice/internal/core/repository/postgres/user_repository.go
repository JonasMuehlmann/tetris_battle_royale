package repository

import (
	"log"
	types "microservice/internal/core/types"
)

type PostgresDatabaseUserRepository struct {
	PostgresDatabase
	Logger *log.Logger
}

func (repo PostgresDatabaseUserRepository) GetUserFromID(userID int) (types.User, error) {
	user := types.User{}

	db, err := repo.GetConnection()
	if err != nil {
		return types.User{}, nil
	}

	defer db.Close()

	err = db.Get(&user, "SELECT * FROM users WHERE ID = $1", userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo PostgresDatabaseUserRepository) GetUserFromName(username string) (types.User, error) {
	user := types.User{}

	db, err := repo.GetConnection()
	if err != nil {
		return user, err
	}

	defer db.Close()

	err = db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo PostgresDatabaseUserRepository) Register(username, password, salt string) (int, error) {
	var userID int

	db, err := repo.GetConnection()
	if err != nil {
		return 0, err
	}

	defer db.Close()

	err = db.QueryRow("INSERT INTO users(username, pw_hash, salt) VALUES($1, $2, $3) RETURNING ID", username, string(password), string(salt)).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
