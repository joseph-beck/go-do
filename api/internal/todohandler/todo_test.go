package todohandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoGet(t *testing.T) {
	p := TodoGet(nil, nil)

	assert.NotNil(t, p)
}

func TestTodoPost(t *testing.T) {
	p := TodoPost(nil)

	assert.NotNil(t, p)
}

func TestTodoPatch(t *testing.T) {
	p := TodoPatch(nil)

	assert.NotNil(t, p)
}

func TestTodoDelete(t *testing.T) {
	p := TodoDelete(nil)

	assert.NotNil(t, p)
}
