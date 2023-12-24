package services

import (
	"go-do/internal/database"
	"net/http"
	"net/http/httptest"
	"testing"

	routey "github.com/joseph-beck/routey/pkg/router"
	"github.com/stretchr/testify/assert"
)

func TestUserSignIn(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("GET", "/api/v1/signin", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserSignUp(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("POST", "/api/v1/signup", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserList(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("GET", "/api/v1/users", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestUserGet(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("GET", "/api/v1/users/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestUserPost(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("POST", "/api/v1/users", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserPut(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("PUT", "/api/v1/users", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserPatch(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("PATCH", "/api/v1/users", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUserDelete(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestUserHead(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("HEAD", "/api/v1/users", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestUserOptions(t *testing.T) {
	db := database.New(database.SQLiteDb())
	s := NewUserService(&db)
	app := routey.New()
	app.Service(&s)

	req, err := http.NewRequest("OPTIONS", "/api/v1/users", nil)
	assert.NoError(t, err)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
