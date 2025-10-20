package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Get port from environment variable, default to 8080 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a basic handler that does nothing
	handler := http.HandlerFunc(handler)

	// Start the server
	log.Printf("Server started in port %s", port)

	// Start the HTTP Server
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Empty handler - does nothing
}
