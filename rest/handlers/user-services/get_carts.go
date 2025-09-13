package userservices

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *OrderHandler) GetCarts(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("user_id").(string)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	carts, err := h.Repo.GetCart(userId)
	if err != nil {
		http.Error(w, "Failed to retrieve carts", http.StatusInternalServerError)
		return
	}
	if len(carts) == 0 {
		http.Error(w, "No carts available", http.StatusNotFound)
		return
	}
	util.SendData(w, carts, http.StatusOK)
}
