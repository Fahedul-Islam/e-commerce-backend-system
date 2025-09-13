package products

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Fahedul-Islam/e-commerce/database/repository"
	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	if productID == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct repository.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid product data", http.StatusBadRequest)
		return
	}
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	if err := h.Repo.Update(pId, &updatedProduct); err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{"message": "Product updated successfully"}, http.StatusOK)
}
