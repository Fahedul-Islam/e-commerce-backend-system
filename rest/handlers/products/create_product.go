package products

import (
	"encoding/json"
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/database/repository"
	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct repository.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Create(&newProduct); err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	util.SendData(w, newProduct, http.StatusCreated)
}
