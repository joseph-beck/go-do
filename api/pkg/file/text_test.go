package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	file := "test.txt"

	err := CreateFile(file)
	assert.NoError(t, err)

	e, err := CheckFile(file)
	assert.NoError(t, err)
	assert.True(t, e)

	err = DeleteFile(file)
	assert.NoError(t, err)

	_, err = CheckFile(file)
	assert.NoError(t, err)
}

func TestWriteFile(t *testing.T) {
	file := "test.txt"

	err := WriteFile(file, "hello file")
	assert.NoError(t, err)

	e, err := CheckFile(file)
	assert.NoError(t, err)
	assert.True(t, e)

	err = DeleteFile(file)
	assert.NoError(t, err)

	_, err = CheckFile(file)
	assert.NoError(t, err)
}

func TestDeleteFile(t *testing.T) {
	file := "test.txt"

	err := CreateFile(file)
	assert.NoError(t, err)

	err = DeleteFile(file)
	assert.NoError(t, err)

	e, err := CheckFile(file)
	assert.NoError(t, err)
	assert.False(t, e)

	err = DeleteFile("test_file.txt")
	assert.Error(t, err)
}

func TestReadFile(t *testing.T) {
	file := "test.txt"

	err := CreateFile(file)
	assert.NoError(t, err)

	err = WriteFile(file, "hello file")
	assert.NoError(t, err)

	c, err := ReadFile(file)
	assert.NoError(t, err)
	assert.Equal(t, "hello file", c[0])

	err = DeleteFile(file)
	assert.NoError(t, err)

	e, err := CheckFile(file)
	assert.NoError(t, err)
	assert.False(t, e)

	_, err = ReadFile("test_file.txt")
	assert.Error(t, err)
}

func TestCheckFile(t *testing.T) {
	g := "test.txt"
	b := "bad.txt"

	err := CreateFile(g)
	assert.NoError(t, err)

	e, err := CheckFile(g)
	assert.NoError(t, err)
	assert.True(t, e)

	err = DeleteFile(g)
	assert.NoError(t, err)

	e, err = CheckFile(g)
	assert.NoError(t, err)
	assert.False(t, e)

	e, err = CheckFile(b)
	assert.NoError(t, err)
	assert.False(t, e)
}

func TestFileContains(t *testing.T) {
	file := "test.txt"

	err := CreateFile(file)
	assert.NoError(t, err)

	err = WriteFile(file, "hello file")
	assert.NoError(t, err)

	c, err := FileContains(file, "hello file")
	assert.NoError(t, err)
	assert.True(t, c)

	c, err = FileContains(file, "world")
	assert.NoError(t, err)
	assert.False(t, c)

	err = DeleteFile(file)
	assert.NoError(t, err)

	e, err := CheckFile(file)
	assert.NoError(t, err)
	assert.False(t, e)

	_, err = FileContains("test_file.txt", "text")
	assert.Error(t, err)
}
