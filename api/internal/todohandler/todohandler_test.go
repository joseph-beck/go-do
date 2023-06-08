package todohandler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoGet(t *testing.T) {
	p := TodoGet()

	assert.NotNil(t, p)
}

func TestTodoPost(t *testing.T) {
	p := TodoPost()

	assert.NotNil(t, p)
}

func TestTodoPatch(t *testing.T) {
	p := TodoPatch()

	assert.NotNil(t, p)
}

func TestTodoDelete(t *testing.T) {
	p := TodoDelete()

	assert.NotNil(t, p)
}
