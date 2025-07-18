package api

import (
	"go-server/api/handlers"

	"github.com/gorilla/mux"
)

// SetupAPIRoutes configures all API routes
func SetupAPIRoutes(r *mux.Router) {
    // Create API subrouter
    apiRouter := r.PathPrefix("/api").Subrouter()

    // API endpoints
    apiRouter.HandleFunc("/products", handlers.GetProducts).Methods("GET")
    apiRouter.HandleFunc("/config", handlers.GetConfig).Methods("GET")
    apiRouter.HandleFunc("/users", handlers.GetUsers).Methods("GET")
}