package dbrepository

import (
	"MAD_Assignment/domain"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// type Reader interface {
// 	GetAll() ([]*domain.Restaurant, error)
// 	GetByID() (*domain.Restaurant, error)
// }

// type Writer interface {
// 	Create(*domain.Restaurant) (string, error)
// 	Update(*domain.Restaurant) error
// 	Archive(*domain.Restaurant) error
// }

func (s *Service) GetAll() ([]*domain.Restaurant, error) {
	return s.repo.GetAll()
}

func (s *Service) GetByID(ID domain.ID) (*domain.Restaurant, error) {
	return s.repo.GetByID(ID)
}
func (s *Service) FindByName(name string) ([]*domain.Restaurant, error) {
	return s.repo.FindByName(name)
}
func (s *Service) FindByTypeOfFood(foodType string) ([]*domain.Restaurant, error) {
	return s.repo.FindByTypeOfFood(foodType)
}
func (s *Service) FindByTypeOfPostCode(postCode string) ([]*domain.Restaurant, error) {
	return s.repo.FindByTypeOfPostCode(postCode)
}
func (s *Service) Search(query string) ([]*domain.Restaurant, error) {
	return s.repo.Search(query)
}
func (s *Service) Store(u *domain.Restaurant) (domain.ID, error) {

	// u.ID = utils.NewUUID()
	// u.CreatedOn = utils.GetUTCTimeNow()
	// return s.repo.Create(u)
	return s.repo.Store(u)

}

func (s *Service) Update(inp *domain.Restaurant) error {

	//inp.UpdatedOn = utils.GetUTCTImeNow()
	return s.repo.Update(inp)
}
func (s *Service) Delete(Id domain.ID) error {

	//inp.UpdatedOn = utils.GetUTCTImeNow()
	return s.repo.Delete(Id)
}
