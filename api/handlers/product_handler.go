package handlers

import (
	"log"
	"net/http"
	"os"
)

// GetProducts serves the products data from JSON file
func GetProducts(w http.ResponseWriter, r *http.Request) {
    log.Printf("API request: GET /api/products")

    // Read the JSON file
    data, err := os.ReadFile("static/data/products.json")
    if err != nil {
        log.Printf("Error reading products.json: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Set proper headers for JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*") // Enable CORS

    // Write JSON response
    w.Write(data)
}