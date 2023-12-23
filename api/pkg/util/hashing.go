package util

import (
	"crypto/sha256"
	"fmt"
)

// Hashes a given string with the SHA256 Hash
func Sha256Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	p := fmt.Sprintf("%x", h.Sum(nil))
	return p
}
