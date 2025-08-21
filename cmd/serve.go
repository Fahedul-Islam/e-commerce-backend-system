package cmd

import (
	"fmt"
	"net/http"

	"github.com/Fahedul-Islam/e-commerce/middleware"
)

func Serve(mux *http.ServeMux) {
	middlewareManager := middleware.NewMiddlewareManager()
	middlewareManager.Use(middleware.Test,middleware.Logger,middleware.CorsWithPreflight)
	initRoutes(mux, middlewareManager)

	wrappedMux := middlewareManager.WrappedMux(mux)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", wrappedMux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}