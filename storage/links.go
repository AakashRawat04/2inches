package storage

import (
	"github.com/jmoiron/sqlx"
)

type LinkStorage struct {
	db *sqlx.DB
}

func NewLinkStorage(db *sqlx.DB) *LinkStorage {
	return &LinkStorage{
		db: db,
	}
}

func (l *LinkStorage) CreateNewLink() error {
	return nil
}

func (l *LinkStorage) GetLinks() error {
	return nil
}
