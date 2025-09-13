package userservices

import (
	"github.com/Fahedul-Islam/e-commerce/database/repository"
)

type OrderHandler struct {
	Repo *repository.OrderRepository
}

func NewOrderHandler(repo *repository.OrderRepository) *OrderHandler {
	return &OrderHandler{Repo: repo}
}
