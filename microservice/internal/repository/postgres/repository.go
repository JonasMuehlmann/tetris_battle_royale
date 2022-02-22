package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	Host     string
	Port     int
	Username string
	DBName   string
}

func MakePostgresDB(host string, port int, username string, dbName string) *PostgresDatabase {
	return &PostgresDatabase{
		Host:     host,
		Port:     port,
		Username: username,
		DBName:   dbName}
}

func MakeDefaultPostgresDB() *PostgresDatabase {
	return &PostgresDatabase{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "prod"}
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
		return nil, errors.New("Failed to open db: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("Failed to open db: " + err.Error())
	}

	return db, nil
}
