package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToJsonString(t *testing.T) {
	testBrand := Brand{
		Id:         1,
		Name:       "Fuji",
		YearEst:    1899,
		Country:    "Japan",
		WebsiteUrl: "www.fuji.com",
	}

	testModel := Model{
		Id:              3,
		Name:            "Cross",
		Version:         "1.3",
		Msrp:            1200.00,
		ManufacturerUrl: "www.fuji.com/cross",
		Brand:           testBrand,
		Type:            "Road",
		Subtype:         "Cyclocross",
	}

	t.Run("print JSON string", func(t *testing.T) {
		want := `{ "id": 3 "name": "Cross" "version": "1.3" "msrp": 1200.00 "manufacturerUrl": "www.fuji.com/cross"`
		assert.Contains(t, testModel.ToJsonString(), want)
		assert.Contains(t, testModel.ToJsonString(), testBrand.ToJsonString())
	})
}
