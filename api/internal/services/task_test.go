package services

import (
	"go-do/internal/database"
	"net/http"
	"net/http/httptest"
	"testing"

	routey "github.com/joseph-beck/routey/pkg/router"
	"github.com/stretchr/testify/assert"
)

func TestTaskList(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewTaskService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("GET", "/api/v1/tasks", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTaskGet(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewTaskService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("GET", "/api/v1/tasks/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTaskPost(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewTaskService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("POST", "/api/v1/tasks", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTaskPut(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewTaskService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("PUT", "/api/v1/tasks", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTaskPatch(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewTaskService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("PATCH", "/api/v1/tasks", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTaskDelete(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewTaskService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("DELETE", "/api/v1/tasks/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestTaskHead(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewTaskService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("HEAD", "/api/v1/tasks", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestTaskOptions(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewTaskService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("OPTIONS", "/api/v1/tasks", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
