package main

// File server for serving static files and handling file requests
// http://localhost:8080/static/images/image6.png
// http://localhost:8080/static/videos/fed.mp4
// http://localhost:8080/static/web/index.html

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"go-server/api"

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

func main() {
    cfg := loadConfig()

    r := mux.NewRouter()
    
    // File serving endpoints
    r.HandleFunc("/files/{type}/{name}", serveFile).Methods("GET")
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
    
    // Setup API routes
    api.SetupAPIRoutes(r)

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