package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostgresMaker(t *testing.T) {
	assert.NotPanics(t, func() { _ = PostgresDb() })
}

func TestSQLiteMaker(t *testing.T) {
	assert.NotPanics(t, func() { _ = SQLiteDb() })
}

func TestMockMaker(t *testing.T) {
	assert.NotPanics(t, func() { _ = MockDb() })
}
