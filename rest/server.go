package rest

import (
	"fmt"
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/config"
	"github.com/Fahedul-Islam/e-commerce/rest/handlers/products"
	userservices "github.com/Fahedul-Islam/e-commerce/rest/handlers/user-services"
	"github.com/Fahedul-Islam/e-commerce/rest/handlers/users"
	"github.com/Fahedul-Islam/e-commerce/rest/middleware"
)

type Server struct {
	productsHandler   *products.ProductHandler
	usersHandler      *users.UserHandler
	orderHandler      *userservices.OrderHandler
}
func NewServer(productsHandler *products.ProductHandler, usersHandler *users.UserHandler, orderHandler *userservices.OrderHandler) *Server {
	return &Server{
		productsHandler: productsHandler,
		usersHandler:    usersHandler,
		orderHandler:    orderHandler,
	}
}

func (server *Server) Start(cfg *config.Config) {
	mux := http.NewServeMux()
	middlewareManager := middleware.NewMiddlewareManager()
	middlewareManager.Use(middleware.Logger, middleware.CorsWithPreflight)
	server.productsHandler.RegisterRoutes(mux, middlewareManager)
	server.usersHandler.RegisterRoutes(mux, middlewareManager)
	server.orderHandler.RegisterRoutes(mux, middlewareManager)

	wrappedMux := middlewareManager.WrappedMux(mux)
	fmt.Println("Server is running on port", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, wrappedMux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
