package data

import (
	"database/sql"

	"github.com/jackelder18/bikedex/domain"
)

func ReadBrand(id int, db *sql.DB) (*domain.Brand, error) {
	b := domain.Brand{}

	query := `SELECT B.id, B.name, B.yearEst, B.country, B.websiteUrl FROM Brand B WHERE id = $1`

	err := db.QueryRow(query, id).Scan(
		&b.Id, &b.Name, &b.YearEst, &b.Country, &b.WebsiteUrl)

	if err != nil {
		return nil, err
	}

	return &b, nil
}
