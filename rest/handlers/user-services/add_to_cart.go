package userservices

import (
	"encoding/json"
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/database/repository"
	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *OrderHandler) CartAdd(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	var item repository.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid item data", http.StatusBadRequest)
		return
	}
	stock, err := h.Repo.AvailableStock(int(item.ProductID))
	if err != nil {
		http.Error(w, "Product not found ", http.StatusBadRequest)
	}
	if stock< item.Quantity {
		http.Error(w, "Not enough stock!", http.StatusNotFound)
	}
	if err := h.Repo.AddToCart(uint(item.ProductID), item.Quantity, userId); err != nil {
		http.Error(w, "Failed to add item to cart", http.StatusInternalServerError)
		return
	}
	util.SendData(w, map[string]string{"message": "Item added to cart successfully"}, http.StatusOK)

}
