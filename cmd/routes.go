package cmd

import (
	"log"
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/config"
	"github.com/Fahedul-Islam/e-commerce/database"
	"github.com/Fahedul-Islam/e-commerce/handlers"
	"github.com/Fahedul-Islam/e-commerce/middleware"
)

func initRoutes(mux *http.ServeMux, middlewareManager *middleware.MiddlewareManager) {
	cfg, _ := config.Load()
	connStr := cfg.GetDBConStr()
	db, err := database.DbConnect(connStr)
	if err != nil {
		panic(err)
	}
	productHandler := handlers.NewProductHandler(database.NewProductRepository(db))
	if err := productHandler.CreateTable(); err != nil {
		log.Fatalf("Error creating product table: %v", err)
	}

	userHandler := handlers.NewUserHandler(database.NewAuthHandler(db, cfg.JWT.Secret))
	if err := userHandler.CreateTable(); err != nil {
		log.Fatalf("Error creating user table: %v", err)
	}

	mux.Handle("GET /products", middlewareManager.With()(http.HandlerFunc(productHandler.GetProducts)))
	mux.Handle("GET /products/{id}", middlewareManager.With()(http.HandlerFunc(productHandler.GetProductByID)))
	mux.Handle("POST /products/create", middlewareManager.With()(http.HandlerFunc(productHandler.CreateProduct)))
	mux.Handle("GET /users", middlewareManager.With()(http.HandlerFunc(userHandler.GetUsers)))
	mux.Handle("POST /users/create", middlewareManager.With()(http.HandlerFunc(userHandler.CreateUser)))
}
