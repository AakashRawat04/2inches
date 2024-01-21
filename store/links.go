package store

import (
	"database/sql"

	model "github.com/aakashrawat04/2inches/models"
)

func (s *PostgresStore) GetLinks() ([]*model.Link, error) {
	rows, _ := s.db.Query("SELECT * FROM links")

	links, err := scanRowsIntoLinks(rows)

	if err != nil {
		return nil, err
	}

	return links, nil
}

func (s *PostgresStore) GetLinkById(userId string) ([]*model.Link, error) {
	rows, _ := s.db.Query("SELECT * FROM links WHERE id = $1", userId)

	links, err := scanRowsIntoLinks(rows)

	if err != nil {
		return nil, err
	}

	return links, nil
}

func scanIntoLink(rows *sql.Rows) (*model.Link, error) {
	link := new(model.Link)
	err := rows.Scan(&link.ID, &link.CreatedAt, &link.LongUrl)

	return link, err
}

func scanRowsIntoLinks(rows *sql.Rows) ([]*model.Link, error) {
	links := []*model.Link{}

	for rows.Next() {
		link, err := scanIntoLink(rows)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}
