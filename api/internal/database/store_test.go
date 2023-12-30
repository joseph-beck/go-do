package database

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var s Store

func beforeEach() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	s = New(SQLiteDb())
	return nil
}

func afterEach() error {
	if s.HasTable("test_table") {
		err := s.DropTable("test_table")
		if err != nil {
			return err
		}
	}

	err := s.Close()
	return err
}

func TestNew(t *testing.T) {
	err := godotenv.Load()
	assert.NoError(t, err)

	s := New(SQLiteDb())
	assert.NotNil(t, s)
}

func TestCreateTable(t *testing.T) {
	err := beforeEach()
	assert.NoError(t, err)

	type Table struct {
		ID   uint
		Name string
	}

	err = s.CreateTable("test_table", &Table{})
	assert.NoError(t, err)

	e := s.HasTable("test_table")
	assert.True(t, e)

	err = s.CreateTable("test_table", &Table{})
	assert.Error(t, err)

	err = afterEach()
	assert.NoError(t, err)
}

func TestDeleteTable(t *testing.T) {
	err := beforeEach()
	assert.NoError(t, err)

	type Table struct {
		ID   uint
		Name string
	}

	err = s.CreateTable("test_table", &Table{})
	assert.NoError(t, err)

	err = s.DropTable("test_table")
	assert.NoError(t, err)

	e := s.HasTable("test_table")
	assert.False(t, e)

	err = s.DropTable("test_table")
	assert.Error(t, err)

	err = afterEach()
	assert.NoError(t, err)
}

func TestCheckTable(t *testing.T) {
	err := beforeEach()
	assert.NoError(t, err)

	type Table struct {
		ID   uint
		Name string
	}

	err = s.CreateTable("test_table", &Table{})
	assert.NoError(t, err)

	e := s.HasTable("test_table")
	assert.True(t, e)

	e = s.HasTable("another_test_table")
	assert.False(t, e)

	err = afterEach()
	assert.NoError(t, err)
}
