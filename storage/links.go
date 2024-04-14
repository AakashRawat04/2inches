package storage

import (
	"github.com/aakashrawat04/2inches/models"
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

func (l *LinkStorage) GetLinks() (*[]*models.Link, error) {
	links := []*models.Link{}
	err := l.db.Select(&links, "SELECT * FROM links")
	if err != nil {
		return nil, err
	}
	return &links, nil
}
