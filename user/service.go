package user

import "github.com/Fahedul-Islam/e-commerce/domain"

type service struct {
	repo UserRepo
}

func NewService(r UserRepo) Service {
	return &service{repo: r}
}

func (s *service) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAllUsers()
}

func (s *service) CreateUser(user *domain.User) error {
	return s.repo.CreateUser(user)
}

func (s *service) AuthenticateUser(login *domain.UserLogin) (*domain.User, error) {
	return s.repo.AuthenticateUser(login)
}

func (s *service) UserValidate(user *domain.UserLogin) error {
	return s.repo.UserValidate(user)
}