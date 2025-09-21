package cmd

import (
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/config"
	"github.com/Fahedul-Islam/e-commerce/database/connections"
	"github.com/Fahedul-Islam/e-commerce/database/repository"
	"github.com/Fahedul-Islam/e-commerce/repo"
	"github.com/Fahedul-Islam/e-commerce/rest"
	"github.com/Fahedul-Islam/e-commerce/rest/handlers/products"
	userservices "github.com/Fahedul-Islam/e-commerce/rest/handlers/user-services"
	usrHandler "github.com/Fahedul-Islam/e-commerce/rest/handlers/users"
	userDomain "github.com/Fahedul-Islam/e-commerce/user"
)

func Serve(mux *http.ServeMux) {
	cfg, _ := config.Load()
	connStr := cfg.GetDBConStr()
	db, err := connections.DbConnect(connStr)
	if err != nil {
		panic(err)
	}
	connections.Migrate(cfg.GetDBURL())
	connections.InitRedis()

	userRepo := repo.NewAuthHandler(db)
	userSrvc := userDomain.NewService(userRepo)
	userHandler := usrHandler.NewUserHandler(cfg, userSrvc)

	productHandler := products.NewProductHandler(repository.NewProductRepository(db))
	orderHandler := userservices.NewOrderHandler(repository.NewOrderRepository(db))
	
	server := rest.NewServer(productHandler, userHandler, orderHandler)
	server.Start(cfg)

}
