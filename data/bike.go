package data

import (
	"database/sql"

	"github.com/jackelder/bikedex/domain"
)

const HYDRATE_BIKE_QUERY = `SELECT Bike.id, Bike.imageUrl, 
M.id, M.name, M.version, M.msrp, M.manufacturerUrl,
B.id, B.name, B.yearEst, B.country, B.websiteUrl,
T.name, S.name, C.id, C.name, C.hex
FROM Bike
INNER JOIN Model M ON (Bike.modelId = M.id)
INNER JOIN Brand B ON (M.brandId = B.id)
INNER JOIN "Type" T ON (M.typeId = T.id)
INNER JOIN Subtype S ON (M.subtypeId = S.id)
INNER JOIN Color C ON (Bike.colorId = c.id)`

func ReadBike(id int, db *sql.DB) (*domain.Bike, error) {
	b := domain.Bike{}

	query := HYDRATE_BIKE_QUERY + `WHERE Bike.id=$1`

	row := db.QueryRow(query, id)

	err := scanBike(row, b)

	if err != nil {
		return nil, err
	}

	return &b, nil
}

func ReadAllBikes(db *sql.DB) ([]*domain.Bike, error) {
	rows, err := db.Query(HYDRATE_BIKE_QUERY)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bikes := make([]*domain.Bike, 0)

	for rows.Next() {
		b := domain.Bike{}

		err := scanBikes(rows, &b)

		if err != nil {
			return nil, err
		}

		bikes = append(bikes, &b)
	}

	if rerr := rows.Close(); rerr != nil {
		return nil, rerr
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bikes, nil
}

func scanBike(row *sql.Row, b domain.Bike) error {
	return row.Scan(
		&b.Id, &b.ImageUrl,
		&b.Model.Id, &b.Model.Name, &b.Model.Version, &b.Model.Msrp, &b.Model.ManufacturerUrl,
		&b.Model.Brand.Id, &b.Model.Brand.Name, &b.Model.Brand.YearEst, &b.Model.Brand.Country, &b.Model.Brand.WebsiteUrl,
		&b.Model.Type, &b.Model.Subtype,
		&b.Color.Id, &b.Color.Name, &b.Color.Hex)
}

func scanBikes(rows *sql.Rows, b *domain.Bike) error {
	return rows.Scan(
		&b.Id, &b.ImageUrl,
		&b.Model.Id, &b.Model.Name, &b.Model.Version, &b.Model.Msrp, &b.Model.ManufacturerUrl,
		&b.Model.Brand.Id, &b.Model.Brand.Name, &b.Model.Brand.YearEst, &b.Model.Brand.Country, &b.Model.Brand.WebsiteUrl,
		&b.Model.Type, &b.Model.Subtype,
		&b.Color.Id, &b.Color.Name, &b.Color.Hex)
}
