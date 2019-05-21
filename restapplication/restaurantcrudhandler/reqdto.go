package restaurantcrudhandler

type RestaurantCreateReqDTO struct {
	Name         string  `json:"name" `
	Address      string  `json:"address" `
	AddressLine2 string  `json:"addressline2" `
	URL          string  `json:"url"`
	Outcode      string  `json:"outcode" `
	Postcode     string  `json:"postcode"`
	Rating       float32 `json:"rating"`
	Type_Of_Food string  `json:"type_of_food"`
}

type RestaurantUpdateDTO struct {
	ID           string  `json:"id" `
	Name         string  `json:"name" `
	Address      string  `json:"address" `
	AddressLine2 string  `json:"addressline2" `
	URL          string  `json:"url"`
	Outcode      string  `json:"outcode" `
	Postcode     string  `json:"postcode"`
	Rating       float32 `json:"rating"`
	Type_Of_Food string  `json:"type_of_food"`
}
