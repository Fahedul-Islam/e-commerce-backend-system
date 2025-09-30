package product

import "github.com/Fahedul-Islam/e-commerce/domain"

type service struct {
	repo ProductRepo
}

func NewService(r ProductRepo) Service {
	return &service{repo: r}
}

func (s *service) Create(product *domain.Product) error {
	return s.repo.Create(product)
}
func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
func (s *service) GetByID(id int) (*domain.Product, error) {
	return s.repo.GetByID(id)
}
func (s *service) GetAll() ([]*domain.Product, error) {
	return s.repo.GetAll()
}
func (s *service) Update(id int, product *domain.Product) error {
	return s.repo.Update(id, product)
}