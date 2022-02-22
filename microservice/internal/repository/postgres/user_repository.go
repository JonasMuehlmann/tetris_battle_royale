package repository

import (
	"microservice/internal/domain"
)

type PostgresDatabaseUserRepository struct {
	PostgresDatabase
}

func (repo PostgresDatabaseUserRepository) GetUserFromID(userID int) (domain.User, error) {
	user := domain.User{}

	db, err := repo.GetConnection()
	if err != nil {
		return domain.User{}, nil
	}

	err = db.Get(&user, "SELECT * FROM users WHERE ID = $1", userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo PostgresDatabaseUserRepository) GetUserFromName(username string) (domain.User, error) {
	user := domain.User{}

	db, err := repo.GetConnection()
	if err != nil {
		return user, err
	}

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
	err = db.QueryRow("INSERT INTO users(username, pw_hash, salt) VALUES($1, $2, $3) RETURNING ID", username, string(password), string(salt)).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
