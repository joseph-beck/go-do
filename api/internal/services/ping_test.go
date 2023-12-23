package services

import (
	"go-do/internal/database"
	"net/http"
	"net/http/httptest"
	"testing"

	routey "github.com/joseph-beck/routey/pkg/router"
	"github.com/stretchr/testify/assert"
)

func TestPingGet(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewPingService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("GET", "/api/v1/ping", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPingPost(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewPingService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("POST", "/api/v1/ping", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPingPut(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewPingService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("PUT", "/api/v1/ping", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPingPatch(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewPingService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("PATCH", "/api/v1/ping", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPingDelete(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewPingService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("DELETE", "/api/v1/ping", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPingHead(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewPingService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("HEAD", "/api/v1/ping", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestPingOptions(t *testing.T) {
	db := database.New(database.MockDb())
	s := NewPingService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("OPTIONS", "/api/v1/ping", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
