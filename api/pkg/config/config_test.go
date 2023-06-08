package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	c := MakeConfig("")

	assert.NotNil(t, c)
}
