package store

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	err := godotenv.Load()
	connectionString := os.Getenv("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}
