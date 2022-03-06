package repository

import (
	"context"
	"encoding/json"
	"log"
	"microservice/internal/core/types"
	"time"

	"github.com/google/uuid"
)

type RedisSessionRepo struct {
	RedisStore
	Logger *log.Logger
}

func (repo RedisSessionRepo) CreateSession(userID string) (string, error) {
	session := types.Session{UserID: userID, CreationTime: time.Now()}

	uuid.EnableRandPool()
	session.ID = uuid.NewString()

	marshalledSession, err := json.Marshal(session)
	if err != nil {
		return "", err
	}

	err = repo.Client.Set(context.Background(), userID, marshalledSession, 0).Err()
	if err != nil {
		return "", err
	}

	return session.ID, nil
}

func (repo RedisSessionRepo) GetSession(userID string) (types.Session, error) {
	var session types.Session

	sessionMarshalled, err := repo.Client.Get(context.Background(), userID).Bytes()
	if err != nil {
		return types.Session{}, err
	}

	err = json.Unmarshal([]byte(sessionMarshalled), &session)
	if err != nil {
		return types.Session{}, err
	}

	return session, nil
}

func (repo RedisSessionRepo) DeleteSession(sessionID string) error {
	return repo.Client.Del(context.Background(), sessionID).Err()
}
