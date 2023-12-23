package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256Hash(t *testing.T) {
	s := Sha256Hash("hello")
	assert.Len(t, s, 64)
}
