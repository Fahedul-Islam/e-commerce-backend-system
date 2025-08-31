package products

import (
	"github.com/Fahedul-Islam/e-commerce/database"
)

type ProductHandler struct {
	Repo *database.ProductRepository
}

func NewProductHandler(repo *database.ProductRepository) *ProductHandler {
	return &ProductHandler{Repo: repo}
}

func (h *ProductHandler) CreateTable() error {
	return h.Repo.InitTable()
}
