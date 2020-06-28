package domain

import "fmt"

type Brand struct {
	Id         int
	Name       string
	YearEst    int
	Country    string
	WebsiteUrl string
}

func (b Brand) ToJsonString() string {
	return fmt.Sprintf("{ \"id\": %d \"name\": %q \"yearEst\": %d, \"country\": %q, \"websiteUrl\": %q }",
		b.Id, b.Name, b.YearEst, b.Country, b.WebsiteUrl)
}
