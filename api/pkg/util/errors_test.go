package util

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrOut(t *testing.T) {
	err := errors.New("error panic")
	assert.Panics(t, func() { ErrOut(err) })
}

func TestErrMsg(t *testing.T) {
	err := errors.New("error message")
	assert.NotPanics(t, func() { ErrMsg(err, "message") })
}

func TestErrLog(t *testing.T) {
	err := errors.New("error log")
	assert.NotPanics(t, func() { ErrLog(err) })
}
