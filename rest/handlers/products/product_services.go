package products

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Fahedul-Islam/e-commerce/database"
	"github.com/Fahedul-Islam/e-commerce/util"
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


func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
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

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
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
