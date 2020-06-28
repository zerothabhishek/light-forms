package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Hello World!")
	router := chi.NewRouter()

	router.Route("/", func(route chi.Router) {
		route.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if sErr := server.ListenAndServe(); sErr != nil && sErr != http.ErrServerClosed {
		fmt.Printf("Server is not listening! %v", sErr)
	}
}
