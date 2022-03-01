package repository

import (
	"errors"
	"log"
	types "microservice/internal/core/types"
	"time"
)

type PostgresDatabaseSessionRepository struct {
	PostgresDatabase
	Logger *log.Logger
}

func (repo PostgresDatabaseSessionRepository) CreateSession(userID int) (int, error) {

	session := types.Session{ID: -1, UserID: userID, CreationTime: time.Now()}

	db, err := repo.GetConnection()

	if err != nil {
		return 0, err
	}

	defer db.Close()

	err = db.QueryRow("INSERT INTO sessions(user_ID, creation_time) VALUES($1, $2) RETURNING ID", session.UserID, session.CreationTime).Scan(&session.ID)
	if err != nil {
		repo.Logger.Printf("Error: %v", err)

		return 0, err
	}

	return session.ID, nil
}

func (repo PostgresDatabaseSessionRepository) GetSession(userID int) (types.Session, error) {
	session := types.Session{}

	db, err := repo.GetConnection()

	if err != nil {
		return session, err
	}

	defer db.Close()

	err = db.Get(&session, "SELECT * FROM sessions WHERE user_ID = $1", userID)
	if err != nil {
		return types.Session{}, errors.New("Failed to retrieve session")
	}

	return session, nil
}

func (repo PostgresDatabaseSessionRepository) DeleteSession(sessionID int) error {
	db, err := repo.GetConnection()

	if err != nil {
		return err
	}

	defer db.Close()

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
