package userservices

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/rest/middleware"
)

var user = "user"

func (h *OrderHandler) RegisterRoutes(mux *http.ServeMux, manager *middleware.MiddlewareManager) {
	// Order and Cart routes. Only user can access
	mux.Handle("POST /cart/add", manager.With(middleware.AuthMiddleware(user))(http.HandlerFunc(h.CartAdd)))
	mux.Handle("GET /cart", manager.With(middleware.AuthMiddleware(user))(http.HandlerFunc(h.GetCarts)))
}
