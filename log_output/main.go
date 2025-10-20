package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"time"
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
	randomStr := randomString()
	log.Println(randomStr)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println(randomStr)
	}
}
