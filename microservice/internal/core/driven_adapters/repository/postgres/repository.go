package repository

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
	Host     string
	Port     int
	Username string
	Password string
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
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(configDir, "tbr", ".postgres_credentials.env"))
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("TBR_PG_HOST")
	portRaw := os.Getenv("TBR_PG_PORT")

	var port int64

	if portRaw == "" {
		port = -1
	} else {
		port, err = strconv.ParseInt(portRaw, 10, 32)
	}

	username := os.Getenv("TBR_PG_USER")
	dbName := os.Getenv("TBR_PG_DB")
	password := os.Getenv("TBR_PG_PASSWORD")

	return &PostgresDatabase{
		Host:     host,
		Port:     int(port),
		Username: username,
		Password: password,
		DBName:   dbName,
		Logger:   logger,
	}
}
func MakeDefaultPostgresTestDB(logger *log.Logger) *PostgresDatabase {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(configDir, "tbr", ".postgres_credentials_test.env"))
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("TBR_PG_HOST")
	portRaw := os.Getenv("TBR_PG_PORT")

	var port int64

	if portRaw == "" {
		port = -1
	} else {
		port, err = strconv.ParseInt(portRaw, 10, 32)
	}

	username := os.Getenv("TBR_PG_USER")
	dbName := os.Getenv("TBR_PG_DB")
	password := os.Getenv("TBR_PG_PASSWORD")

	return &PostgresDatabase{
		Host:     host,
		Port:     int(port),
		Username: username,
		Password: password,
		DBName:   dbName,
		Logger:   logger,
	}
}

func (dbImpl *PostgresDatabase) MakeConnectionString() string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbImpl.Host,
		dbImpl.Username,
		dbImpl.Password,
		dbImpl.DBName)

	if dbImpl.Port != -1 {
		connectionString += strconv.FormatInt(int64(dbImpl.Port), 10)
	}

	return connectionString
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
