# Go File Server

This project is a simple file server implemented in Go (Golang) that serves static files such as images and videos. It is designed to handle incoming HTTP requests and serve the requested files from the `static` directory.

## Project Structure

```
go-file-server
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── file_handler.go  # Handles serving files
│   ├── middleware
│   │   └── logging.go       # Middleware for logging requests
│   └── config
│       └── config.go        # Configuration settings
├── static
│   ├── images               # Directory for image files
│   └── videos               # Directory for video files
├── go.mod                   # Module definition file
├── go.sum                   # Checksums for module dependencies
└── README.md                # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd go-file-server
   ```

2. **Install dependencies:**
   Ensure you have Go installed on your machine. Run the following command to install the necessary dependencies:
   ```bash
   go mod tidy
   ```

3. **Run the server:**
   You can start the server by running:
   ```bash
   go run cmd/server/main.go
   ```

4. **Access the server:**
   Open your web browser and navigate to `http://localhost:8080` to access the file server. You can request files from the `static/images` and `static/videos` directories.

## Usage

- To serve an image, navigate to `http://localhost:8080/images/<image-file-name>`.
- To serve a video, navigate to `http://localhost:8080/videos/<video-file-name>`.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.