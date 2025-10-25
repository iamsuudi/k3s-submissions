package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func randomString() string {
	b := make([]byte, 12)

	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("failed to generate random bytes: %w", err)
	}

	s := base64.RawURLEncoding.EncodeToString(b)

	if len(s) > 12 {
		s = s[:12]
	}

	return s
}

func main() {
	// Get port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Skip favicon requests
		if strings.Contains(r.URL.Path, "favicon.ico") {
			return
		}

		// Get current time
		currentTime := time.Now()

		// Format the timestamp
		formattedTime := currentTime.Format("20060-01-02 15:04:05")

		stringNow := randomString()
		fmt.Println("--------------------")
		fmt.Printf("[%s] Responding with %s\n", formattedTime, stringNow)
		
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s: %s\n", formattedTime, stringNow)
	})

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
