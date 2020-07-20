package data

import (
	"database/sql"

	"github.com/jackelder18/bikedex/domain"
)

func ReadColor(id int, db *sql.DB) (*domain.Color, error) {
	c := domain.Color{}

	err := db.QueryRow(`SELECT * FROM Color WHERE id = $1`, id).Scan(&c.Id, &c.Name, &c.Hex)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
