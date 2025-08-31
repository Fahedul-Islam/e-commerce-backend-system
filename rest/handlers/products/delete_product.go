package products

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/util"
)

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	if productID == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	if err := h.Repo.Delete(productID); err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	util.SendData(w, map[string]string{"message": "Product deleted successfully"}, http.StatusOK)
}
