package repository

import (
	"database/sql"
	"fmt"

	"github.com/Fahedul-Islam/e-commerce/database/connections"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (h *OrderRepository) AddToCart(productId uint, quantity int, userId string) error {
	key := "cart_user_" + userId

	var existingCart map[string]string
	var err error
	existingCart, err = h.GetCart(userId)
	if err != nil || existingCart == nil {
		existingCart = make(map[string]string)
	}
	existingCart[fmt.Sprintf("%d", productId)] = fmt.Sprintf("%d", quantity)
	// Save updated cart back to Redis
	return connections.SaveCartToRedis(key, existingCart)

}

func (h *OrderRepository) GetCart(userId string) (map[string]string, error) {
	key := "cart_user_" + userId
	var cartData map[string]string
	err := connections.GetCartFromRedis(key, &cartData)
	return cartData, err
}

func (r *OrderRepository)AvailableStock(pId int) (int, error){
	var p Product
	query := `SELECT * FROM products WHERE id = $1`
	err := r.DB.QueryRow(query, pId).Scan(&p.ID, &p.Name, &p.Price, &p.ImageUrl, &p.IsAvailable, &p.StockQuantity)
	if err != nil {
		return 0,err
	}
	return p.StockQuantity, err
}