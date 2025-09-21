package users

import "github.com/Fahedul-Islam/e-commerce/domain"

type Service interface {
	CreateUser(*domain.User) error
	GetAllUsers() ([]domain.User, error)
	AuthenticateUser(*domain.UserLogin) (*domain.User, error)
	UserValidate(*domain.UserLogin) error
}