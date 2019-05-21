package dbrepository

import (
	"MAD_Assignment/domain"
)

type Factory struct {
}

func (f *Factory) NewUser(name string, address string, addresline2 string, url string, outcode string, postcode string, rating float32, type_of_food string) *domain.Restaurant {
	return &domain.Restaurant{
		Name:         name,
		Address:      address,
		AddressLine2: addresline2,
		URL:          url,
		Outcode:      outcode,
		Postcode:     postcode,
		Rating:       rating,
		Type_Of_Food: type_of_food,
	}

}
