package products

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/rest/middleware"
)
var admin = "admin"
var user = "user"

func (h *ProductHandler) RegisterRoutes(mux *http.ServeMux, manager *middleware.MiddlewareManager) {
	mux.Handle("GET /products", manager.With(middleware.AuthMiddleware(user))(http.HandlerFunc(h.GetAllProducts)))
	mux.Handle("GET /products/{id}", manager.With(middleware.AuthMiddleware(user))(http.HandlerFunc(h.GetProductByID)))

	// Product management routes. Only admin can access
	mux.Handle("POST /products/create", manager.With(middleware.AuthMiddleware(admin))(http.HandlerFunc(h.CreateProduct)))
	mux.Handle("DELETE /products/delete/{id}", manager.With(middleware.AuthMiddleware(admin))(http.HandlerFunc(h.DeleteProduct)))
	mux.Handle("PUT /products/update/{id}", manager.With(middleware.AuthMiddleware(admin))(http.HandlerFunc(h.UpdateProduct)))
}
