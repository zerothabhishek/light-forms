package main

import (
	"fmt"
	"net/http"

	"github.com/light-forms/backend/handlers"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.Route("/", func(route chi.Router) {
		route.Get("/health", handlers.HealthCheck)
		route.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Root World!"))
		})
	})

	server := &http.Server{
		Addr:    ":3003",
		Handler: router,
	}

	if sErr := server.ListenAndServe(); sErr != nil && sErr != http.ErrServerClosed {
		fmt.Printf("Server is not listening! %v", sErr)
	}
}
