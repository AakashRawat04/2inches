package storage

import (
	"github.com/aakashrawat04/2inches/lib"
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

func (l *LinkStorage) CreateLink(link *models.CreateLinkRequest) (*string, error) {
	if link.ShortCode == "" {
		id, err := lib.GenerateShortCode()
		if err != nil {
			return nil, err
		}
		link.ShortCode = id
	}
	_, err := l.db.Exec("INSERT INTO links (user_id, long_url, expires_at, short_code) VALUES ($1, $2, $3, $4)", link.UserID, link.LongURL, link.ExpiresAt, link.ShortCode)
	if err != nil {
		return nil, err
	}
	return &link.ShortCode, nil
}

func (l *LinkStorage) GetLinks() (*[]*models.Link, error) {
	links := []*models.Link{}
	err := l.db.Select(&links, "SELECT * FROM links")
	if err != nil {
		return nil, err
	}
	return &links, nil
}

func (l *LinkStorage) GetAllLinksByUserId(id int) (*[]*models.Link, error) {
	links := []*models.Link{}
	err := l.db.Select(&links, "SELECT * FROM links WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	return &links, nil
}
