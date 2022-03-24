package repository

import (
	"log"
	types "microservice/internal/core/types"
)

type PostgresDatabaseUserRepository struct {
	PostgresDatabase
	Logger *log.Logger
}

func (repo *PostgresDatabaseUserRepository) GetUserFromID(userID string) (types.User, error) {
	user := types.User{}

	err := repo.DBConn.Get(&user, "SELECT * FROM users WHERE ID = $1", userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *PostgresDatabaseUserRepository) GetUserFromName(username string) (types.User, error) {
	user := types.User{}

	err := repo.DBConn.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *PostgresDatabaseUserRepository) Register(username, password, salt string) (string, error) {
	var userID string

	err := repo.DBConn.QueryRow("INSERT INTO users(id, username, pw_hash, salt) VALUES(uuid_generate_v4(), $1, $2, $3) RETURNING ID", username, string(password), string(salt)).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}
