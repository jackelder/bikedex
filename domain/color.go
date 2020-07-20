package domain

import "fmt"

type Color struct {
	Id   int
	Name string
	Hex  string
}

func (c *Color) ToJsonString() string {
	return fmt.Sprintf(`{ "id": %d "name": %q "hex": %q }`,
		c.Id, c.Name, c.Hex)
}
