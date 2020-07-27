package domain

import "fmt"

type Bike struct {
	Id       int
	ImageUrl string
	Model    Model
	Color    Color
}

func (b *Bike) ToJsonString() string {
	return fmt.Sprintf(`{ "id": %d "imageUrl": %q "model": %s "color": %s }`,
		b.Id, b.ImageUrl, b.Model.ToJsonString(), b.Color.ToJsonString())
}
