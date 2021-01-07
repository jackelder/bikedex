package data

import (
	"database/sql"

	"github.com/jackelder/bikedex/domain"
)

func ReadModel(id int, db *sql.DB) (*domain.Model, error) {
	m := domain.Model{}

	query := `SELECT M.id, M.name, M.version, M.msrp, M.manufacturerUrl, 
	B.id, B.name, B.yearEst, B.country, B.websiteUrl,T.name, S.name
	FROM Model M 
	INNER JOIN Brand B ON (M.brandId = B.id)
	INNER JOIN "Type" T ON (M.typeId = T.id)
	INNER JOIN Subtype S ON (M.subtypeId = S.id)
	WHERE M.id=$1;`

	err := db.QueryRow(query, id).Scan(
		&m.Id, &m.Name, &m.Version, &m.Msrp, &m.ManufacturerUrl,
		&m.Brand.Id, &m.Brand.Name, &m.Brand.YearEst, &m.Brand.Country, &m.Brand.WebsiteUrl,
		&m.Type, &m.Subtype)

	if err != nil {
		return nil, err
	}

	return &m, nil
}
