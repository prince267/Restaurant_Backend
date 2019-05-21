package dbrepository

import "MAD_Assignment/domain"

//Reader read from db
type Reader interface {
	GetByID(id domain.ID) (*domain.Restaurant, error)
	GetAll() ([]*domain.Restaurant, error)
	//Regex Substring Match on the name field
	FindByName(name string) ([]*domain.Restaurant, error)
}

//Writer  write to db
type Writer interface {
	//Create Or update
	Update(inp *domain.Restaurant) error
	//Create(u *domain.Restaurant) (string, error)
	Store(b *domain.Restaurant) (domain.ID, error)
	Delete(id domain.ID) error
}

//Filter Find objects by additional filters
type Filter interface {
	FindByTypeOfFood(foodType string) ([]*domain.Restaurant, error)
	FindByTypeOfPostCode(postCode string) ([]*domain.Restaurant, error)
	//(foodType string) int
	//Search --> across all string fields regex match with case insensitive
	//substring match accross all string fields
	Search(query string) ([]*domain.Restaurant, error)
}

//Repository db interface
type Repository interface {
	Reader
	Writer
	Filter
}
