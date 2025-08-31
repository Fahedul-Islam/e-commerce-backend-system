package products

import (
	"net/http"
	"strconv"

	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	pID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.Repo.GetByID(pID)
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	util.SendData(w, product, http.StatusOK)
}