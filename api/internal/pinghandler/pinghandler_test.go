package pinghandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingGet(t *testing.T) {
	p := PingGet(nil)

	assert.NotNil(t, p)
}
