package main

import (
	"log"
	"net/http"
	"os"

	"tesodev-product-api/middleware"
	"tesodev-product-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(r)

	// Apply middleware
	r.Use(middleware.Logger)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
