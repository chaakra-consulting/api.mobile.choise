package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

// GenerateRandomSecret generates a cryptographically secure random secret of a specific length
func GenerateRandomSecret(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	// Encode to Base64 to make it a usable string for environment variables/config files
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func main() {
	// Generate a 256-bit (32-byte) secret
	secret, err := GenerateRandomSecret(32)
	if err != nil {
		log.Fatalf("Error generating secret: %v", err)
	}
	fmt.Printf("Generated Secure JWT Secret: %s\n", secret)
}
