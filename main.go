package main

// File server for serving static files and handling file requests
// http://localhost:8080/static/images/image6.png
// http://localhost:8080/static/videos/fed.mp4

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

type Config struct {
    Port      string
    StaticDir string
}

func loadConfig() *Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    staticDir := os.Getenv("STATIC_DIR")
    if staticDir == "" {
        staticDir = "./static"
    }

    return &Config{
        Port:      port,
        StaticDir: staticDir,
    }
}

func serveFile(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    fileType := vars["type"]
    fileName := vars["name"]
    
    // Basic security check to prevent directory traversal
    if strings.Contains(fileName, "..") {
        log.Printf("Invalid file path detected: %s", fileName)
        http.Error(w, "Invalid file path", http.StatusBadRequest)
        return
    }
    
    // Construct file path
    filePath := filepath.Join("static", fileType, fileName)
    log.Printf("Looking for file at: %s", filePath)
    
    // Check if file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        log.Printf("File not found: %s", filePath)
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }

    log.Printf("Serving file: %s", filePath)
    // Serve the file
    http.ServeFile(w, r, filePath)
}

// API endpoint to serve products data
func getProducts(w http.ResponseWriter, r *http.Request) {
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

// API endpoint to serve config data
func getConfig(w http.ResponseWriter, r *http.Request) {
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

// API endpoint to serve users data
func getUsers(w http.ResponseWriter, r *http.Request) {
    log.Printf("API request: GET /api/users")
    
    // Read the JSON file
    data, err := os.ReadFile("static/data/users.json")
    if err != nil {
        log.Printf("Error reading users.json: %v", err)
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    
    // Set proper headers for JSON response
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*") // Enable CORS
    
    // Write JSON response
    w.Write(data)
}

func main() {
    cfg := loadConfig()

    r := mux.NewRouter()
    
    // File serving endpoints
    r.HandleFunc("/files/{type}/{name}", serveFile).Methods("GET")
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    
    // API endpoints
    r.HandleFunc("/api/products", getProducts).Methods("GET")
    r.HandleFunc("/api/config", getConfig).Methods("GET")
    r.HandleFunc("/api/users", getUsers).Methods("GET")

    log.Printf("Starting server on port %s...", cfg.Port)
    log.Printf("Static files will be served from ./static/")
    log.Printf("File endpoints:")
    log.Printf("  - http://localhost:%s/files/{type}/{filename}", cfg.Port)
    log.Printf("  - http://localhost:%s/static/{type}/{filename}", cfg.Port)
    log.Printf("API endpoints:")
    log.Printf("  - http://localhost:%s/api/products", cfg.Port)
    log.Printf("  - http://localhost:%s/api/config", cfg.Port)
    log.Printf("  - http://localhost:%s/api/users", cfg.Port)

    // Check if static directory exists
    if _, err := os.Stat("./static"); os.IsNotExist(err) {
        log.Printf("WARNING: ./static directory does not exist!")
    } else {
        log.Printf("Static directory exists")
    }
    
    if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}