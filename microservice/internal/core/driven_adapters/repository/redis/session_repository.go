package repository

import (
	"context"
	"encoding/json"
	"log"
	"microservice/internal/core/types"
	"strconv"
	"time"
)

type RedisSessionRepo struct {
	RedisStore
	Logger *log.Logger
}

func (repo RedisSessionRepo) CreateSession(userID int) (int, error) {
	session := types.Session{UserID: userID, CreationTime: time.Now()}

	result := repo.Client.Incr(context.Background(), "sessionCount")

	err := result.Err()
	if err != nil {
		return 0, err
	}

	session.ID = int(result.Val())

	marshalledSession, err := json.Marshal(session)
	if err != nil {
		return 0, err
	}

	err = repo.Client.Set(context.Background(), strconv.FormatInt(int64(userID), 10), marshalledSession, 0).Err()
	if err != nil {
		return 0, err
	}

	return session.ID, nil
}

func (repo RedisSessionRepo) GetSession(userID int) (types.Session, error) {
	var session types.Session

	sessionMarshalled, err := repo.Client.Get(context.Background(), strconv.FormatInt(int64(userID), 10)).Bytes()
	if err != nil {
		return types.Session{}, err
	}

	err = json.Unmarshal([]byte(sessionMarshalled), &session)
	if err != nil {
		return types.Session{}, err
	}

	return session, nil
}

func (repo RedisSessionRepo) DeleteSession(sessionID int) error {
	return repo.Client.Del(context.Background(), strconv.FormatInt(int64(sessionID), 10)).Err()
}
