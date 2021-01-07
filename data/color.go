package data

import (
	"database/sql"

	"github.com/jackelder/bikedex/domain"
)

func ReadColor(id int, db *sql.DB) (*domain.Color, error) {
	c := domain.Color{}

	query := `SELECT C.id, C.name, C.hex FROM Color C WHERE id = $1`

	err := db.QueryRow(query, id).Scan(&c.Id, &c.Name, &c.Hex)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
