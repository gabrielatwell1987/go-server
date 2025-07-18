package handlers

import (
	"log"
	"net/http"
	"os"
)

// GetConfig serves the config data from JSON file
func GetConfig(w http.ResponseWriter, r *http.Request) {
    log.Printf("API request: GET /api/config")

    // Read the JSON file
    data, err := os.ReadFile("static/data/config.json")
    if err != nil {
        log.Printf("Error reading config.json: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }

    // Set proper headers for JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*") // Enable CORS

    // Write JSON response
    w.Write(data)
}