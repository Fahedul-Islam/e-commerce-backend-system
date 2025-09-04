package util

import (
	"crypto/rand"
	"fmt"
	"log"
)

func GenerateOTP() (string, error) {
	b := make([]byte, 3)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Failed to generate OTP: %v", err)
		return "", err
	}
	return fmt.Sprintf("%06d", int(b[0])<<16|int(b[1])<<8|int(b[2])%1000000), nil
}