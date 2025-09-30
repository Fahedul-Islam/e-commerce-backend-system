package products

import (
	"github.com/Fahedul-Islam/e-commerce/config"
	"github.com/Fahedul-Islam/e-commerce/domain"
)

type Service interface {
	Create(product *domain.Product) error
	Delete(id string) error
	GetByID(id int) (*domain.Product, error)
	GetAll() ([]*domain.Product, error)
	Update(id int, product *domain.Product) error
}

type ProductHandler struct {
	cfg     *config.Config
	service Service
}

func NewProductHandler(cfg *config.Config, service Service) *ProductHandler {
	return &ProductHandler{
		cfg:     cfg,
		service: service,
	}
}
