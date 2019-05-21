package domain

//Restaurant is a domain object
type Restaurant struct {
	DBID         ID      `json:"id" bson:"_id"`
	Name         string  `json:"name" bson:"name"`
	Address      string  `json:"address" bson:"address"`
	AddressLine2 string  `json:"addressline2" bson:"addressline2"`
	URL          string  `json:"url" bson:"url"`
	Outcode      string  `json:"outcode" bson:"outcode"`
	Postcode     string  `json:"postcode" bson:"postcode"`
	Rating       float32 `json:"rating" bson:"rating"`
	Type_Of_Food string  `json:"type_Of_food" bson:"type_Of_food"`
}
