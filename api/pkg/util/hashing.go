package util

import (
	"crypto/sha256"
	"fmt"
)

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	p := fmt.Sprintf("%x", h.Sum(nil))
	return p
}
