package cmd

import (
	"fmt"
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/config"
	"github.com/Fahedul-Islam/e-commerce/rest/middleware"
)

func Serve(mux *http.ServeMux) {
	cfg, err := config.Load()
	if err != nil {
		panic("Failed to load config: " + err.Error())
	}

	middlewareManager := middleware.NewMiddlewareManager()
	middlewareManager.Use(middleware.Logger, middleware.CorsWithPreflight)
	initRoutes(mux, middlewareManager)

	wrappedMux := middlewareManager.WrappedMux(mux)
	fmt.Println("Server is running on port", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, wrappedMux); err != nil {
		fmt.Println("Error starting server:", err)
	}

}
