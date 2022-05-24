package services

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Encoder(pass string) string {
	str := sha256.Sum256([]byte(pass))
	return fmt.Sprintf("%x", str)
}
