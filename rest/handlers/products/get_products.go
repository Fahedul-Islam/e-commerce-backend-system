package products

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}
	if len(products) == 0 {
		http.Error(w, "No products available", http.StatusNotFound)
		return
	}
	util.SendData(w, products, http.StatusOK)
}
