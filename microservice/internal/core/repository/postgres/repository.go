package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	Host     string
	Port     int
	Username string
	DBName   string
	Logger   *log.Logger
}

func MakePostgresDB(host string, port int, username string, dbName string, logger *log.Logger) *PostgresDatabase {
	return &PostgresDatabase{
		Host:     host,
		Port:     port,
		Username: username,
		DBName:   dbName,
		Logger:   logger}
}

func MakeDefaultPostgresDB(logger *log.Logger) *PostgresDatabase {
	return &PostgresDatabase{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "prod",
		Logger:   logger,
	}
}

func (dbImpl *PostgresDatabase) MakeConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		dbImpl.Host,
		dbImpl.Port,
		dbImpl.Username,
		dbImpl.DBName)
}

func (dbImpl *PostgresDatabase) GetConnection() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dbImpl.MakeConnectionString())
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to open db: %v", err)
		dbImpl.Logger.Println(errorMessage)

		return nil, errors.New(errorMessage)
	}

	err = db.Ping()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to open db: %v", err)
		dbImpl.Logger.Println(errorMessage)

		return nil, errors.New(errorMessage)
	}

	dbImpl.Logger.Println("Successfully opened db connection")

	return db, nil
}
