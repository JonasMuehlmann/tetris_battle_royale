package repository

import (
	"log"
	"os"
	"strconv"

	redis "github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

type RedisStore struct {
	Logger *log.Logger
	Client *redis.Client
}

func MakeDefaultRedisStore(logger *log.Logger) RedisStore {
	err := godotenv.Load(".redis_credentials.env")
	if err != nil {
		log.Fatal(err)
	}

	address := os.Getenv("TBR_REDIS_ADDR")
	dbRaw := os.Getenv("TBR_REDIS_DB")

	db, err := strconv.ParseInt(dbRaw, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	return RedisStore{
		Logger: logger,
		Client: redis.NewClient(&redis.Options{Addr: address, Password: "", DB: int(db)}),
	}
}
