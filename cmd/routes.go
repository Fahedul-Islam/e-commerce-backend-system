package cmd

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/config"
	"github.com/Fahedul-Islam/e-commerce/database"
	"github.com/Fahedul-Islam/e-commerce/rest/handlers/products"
	"github.com/Fahedul-Islam/e-commerce/rest/handlers/users"
	"github.com/Fahedul-Islam/e-commerce/rest/middleware"
)

var user = "user"
var admin = "admin"

func initRoutes(mux *http.ServeMux, middlewareManager *middleware.MiddlewareManager) {
	cfg, _ := config.Load()
	connStr := cfg.GetDBConStr()
	db, err := database.DbConnect(connStr)
	if err != nil {
		panic(err)
	}
	database.Migrate(cfg.GetDBURL())
	database.InitRedis()

	productHandler := products.NewProductHandler(database.NewProductRepository(db))
	userHandler := users.NewUserHandler(database.NewAuthHandler(db, cfg.JWT.Secret))

	mux.Handle("GET /products", middlewareManager.With(middleware.AuthMiddleware(user))(http.HandlerFunc(productHandler.GetAllProducts)))
	mux.Handle("GET /products/{id}", middlewareManager.With(middleware.AuthMiddleware(user))(http.HandlerFunc(productHandler.GetProductByID)))

	// Product management routes. Only admin can access
	mux.Handle("POST /products/create", middlewareManager.With(middleware.AuthMiddleware(admin))(http.HandlerFunc(productHandler.CreateProduct)))
	mux.Handle("DELETE /products/delete/{id}", middlewareManager.With(middleware.AuthMiddleware(admin))(http.HandlerFunc(productHandler.DeleteProduct)))
	mux.Handle("PUT /products/update/{id}", middlewareManager.With(middleware.AuthMiddleware(admin))(http.HandlerFunc(productHandler.UpdateProduct)))

	mux.Handle("GET /users", middlewareManager.With()(http.HandlerFunc(userHandler.GetUsers)))
	mux.Handle("POST /register", middlewareManager.With()(http.HandlerFunc(userHandler.Register)))
	mux.Handle("POST /login", middlewareManager.With()(http.HandlerFunc(userHandler.Login)))
	mux.Handle("POST /refresh-token", middlewareManager.With()(http.HandlerFunc(userHandler.RefreshHandler)))

}
