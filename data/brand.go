package data

import (
	"database/sql"

	"github.com/jackelder18/bikedex/domain"
)

func ReadBrand(id int, db *sql.DB) (*domain.Brand, error) {
	b := domain.Brand{}

	err := db.QueryRow(`SELECT * FROM Brand WHERE id = $1`, id).Scan(&b.Id, &b.Name, &b.YearEst, &b.Country, &b.WebsiteUrl)

	if err != nil {
		return nil, err
	}

	return &b, nil
}
