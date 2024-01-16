package database

import (
	"database/sql"
	"os"

	"github.com/aakashrawat04/2inches/types"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

type Database interface {
	GetLinks() ([]*types.Link, error)
}

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

func (s *PostgresStore) GetLinks() ([]*types.Link, error) {
	rows, err := s.db.Query("SELECT * FROM links")
	if err != nil {
		return nil, err
	}

	links := []*types.Link{}

	for rows.Next() {
		link, err := scanIntoLink(rows)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	println(links)

	return links, nil
}

func scanIntoLink(rows *sql.Rows) (*types.Link, error) {
	link := new(types.Link)
	err := rows.Scan(&link.ID, &link.LongUrl, &link.CreatedAt)

	return link, err
}
