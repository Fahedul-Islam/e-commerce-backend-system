package middleware

import (
	"log"
	"net/http"
)

func Test(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is a test middleware that does nothing
		// You can add your test logic here if needed
		log.Printf("Test middleware executed for %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
		
	})
}