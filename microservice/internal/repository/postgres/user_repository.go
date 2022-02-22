package repository

import (
	"microservice/api/model"
)

type PostgresDatabaseUserRepository struct {
	PostgresDatabase
}

func (repo *PostgresDatabaseUserRepository) GetUserFromID(userID int) (model.User, error) {
	user := model.User{}

	db, err := repo.GetDBConnection()
	if err != nil {
		return model.User{}, nil
	}

	err = db.Get(&user, "SELECT * FROM users WHERE ID = $1", userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *PostgresDatabaseUserRepository) GetUserFromName(username string) (model.User, error) {
	user := model.User{}

	db, err := repo.GetDBConnection()
	if err != nil {
		return user, err
	}

	err = db.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *PostgresDatabaseUserRepository) Register(username, password, salt string) (int, error) {
	var userID int

	db, err := repo.GetDBConnection()
	if err != nil {
		return 0, err
	}
	err = db.QueryRow("INSERT INTO users(username, pw_hash, salt) VALUES($1, $2, $3) RETURNING ID", username, string(password), string(salt)).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
