package users

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/rest/middleware"
)

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux, manager *middleware.MiddlewareManager) {
	mux.Handle("GET /users", manager.With()(http.HandlerFunc(h.GetUsers)))
	mux.Handle("POST /register", manager.With()(http.HandlerFunc(h.Register)))
	mux.Handle("POST /register/verify-otp", manager.With()(http.HandlerFunc(h.VerifyOTP)))
	mux.Handle("POST /login", manager.With()(http.HandlerFunc(h.Login)))
	mux.Handle("POST /refresh-token", manager.With()(http.HandlerFunc(h.RefreshHandler)))
}
