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

func (repo *PostgresDatabaseSessionRepository) CreateSession(userID string) (string, error) {
	session := types.Session{ID: "", UserID: userID, CreationTime: time.Now()}

	err := repo.DBConn.QueryRow("INSERT INTO sessions(id, user_ID, creation_time) VALUES(uuid_generate_v4(), $1, $2) RETURNING ID", session.UserID, session.CreationTime).Scan(&session.ID)
	if err != nil {
		repo.Logger.Printf("Error: %v\n", err)

		return "", err
	}

	return session.ID, nil
}

func (repo *PostgresDatabaseSessionRepository) GetSession(userID string) (types.Session, error) {
	session := types.Session{}

	err := repo.DBConn.Get(&session, "SELECT * FROM sessions WHERE user_ID = $1", userID)
	if err != nil {
		return types.Session{}, errors.New("Failed to retrieve session")
	}

	return session, nil
}

func (repo *PostgresDatabaseSessionRepository) DeleteSession(sessionID string) error {
	res, err := repo.DBConn.Exec("DELETE FROM sessions WHERE ID = $1", sessionID)
	if err != nil {
		repo.Logger.Printf("Error: %v\n", err)

		return err
	}

	numDeletedSessions, err := res.RowsAffected()
	if err != nil {
		repo.Logger.Printf("Error: %v\n", err)

		return err
	}

	if numDeletedSessions == 0 {
		return errors.New("Session does not exist")
	}

	return nil
}
