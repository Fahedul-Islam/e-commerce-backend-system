package user

import (
	"github.com/Fahedul-Islam/e-commerce/domain"
	userHandler "github.com/Fahedul-Islam/e-commerce/rest/handlers/users"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	GetAllUsers() ([]domain.User, error)
	CreateUser(*domain.User) error
	AuthenticateUser(*domain.UserLogin) (*domain.User, error)
	UserValidate(*domain.UserLogin) error
}