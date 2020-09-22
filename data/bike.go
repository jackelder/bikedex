package data
//hi
import (
	"database/sql"

	"github.com/jackelder18/bikedex/domain"
)

func ReadBike(id int, db *sql.DB) (*domain.Bike, error) {
	b := domain.Bike{}

	query := `SELECT Bike.id, Bike.imageUrl, 
	M.id, M.name, M.version, M.msrp, M.manufacturerUrl,
	B.id, B.name, B.yearEst, B.country, B.websiteUrl,
	T.name, S.name, C.id, C.name, C.hex
	FROM Bike
	INNER JOIN Model M ON (Bike.modelId = M.id)
	INNER JOIN Brand B ON (M.brandId = B.id)
	INNER JOIN "Type" T ON (M.typeId = T.id)
	INNER JOIN Subtype S ON (M.subtypeId = S.id)
	INNER JOIN Color C ON (Bike.colorId = c.id)
	WHERE Bike.id=$1;`

	err := db.QueryRow(query, id).Scan(
		&b.Id, &b.ImageUrl,
		&b.Model.Id, &b.Model.Name, &b.Model.Version, &b.Model.Msrp, &b.Model.ManufacturerUrl,
		&b.Model.Brand.Id, &b.Model.Brand.Name, &b.Model.Brand.YearEst, &b.Model.Brand.Country, &b.Model.Brand.WebsiteUrl,
		&b.Model.Type, &b.Model.Subtype,
		&b.Color.Id, &b.Color.Name, &b.Color.Hex)

	if err != nil {
		return nil, err
	}

	return &b, nil
}
