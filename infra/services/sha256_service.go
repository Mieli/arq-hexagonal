package services

import (
	"crypto/sha256"
	"fmt"
)

func Sha256Encoder(value string) string {
	str := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", str)
}
