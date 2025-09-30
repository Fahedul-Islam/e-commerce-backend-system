package product

import (
	"github.com/Fahedul-Islam/e-commerce/domain"
	productHandler "github.com/Fahedul-Islam/e-commerce/rest/handlers/products"
)

type Service interface {
	productHandler.Service
}

type ProductRepo interface {
	Create(product *domain.Product) error
	Delete(id string) error
	GetByID(id int) (*domain.Product, error)
	GetAll() ([]*domain.Product, error)
	Update(id int, product *domain.Product) error
}
