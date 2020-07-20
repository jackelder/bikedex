package domain

import "fmt"

type Model struct {
	Id              int
	Name            string
	Version         string
	Msrp            float32
	ManufacturerUrl string
	Brand           Brand
	Type            string
	Subtype         string
}

func (m *Model) ToJsonString() string {
	return fmt.Sprintf(`{ "id": %d "name": %q "version": %q "msrp": %0.2f "manufacturerUrl": %q "brand": %s "type": %q "subtype": %q }`,
		m.Id, m.Name, m.Version, m.Msrp, m.ManufacturerUrl, m.Brand.ToJsonString(), m.Type, m.Subtype)
}
