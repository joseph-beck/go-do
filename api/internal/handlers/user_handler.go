package handlers

import (
	"go-do/internal/database"
	"go-do/internal/models"
	"go-do/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserLogon(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func UserSignup(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /user
func UserList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /user/:id
func UserGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /user
func UserPost(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		m := models.User{}
		c.ShouldBindJSON(&m)
		m.Password = util.Sha256Hash(m.Password)

		e := s.CheckUser(m.ID)
		if e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.AddUser(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// /user
func UserPut(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /user
func UserPatch(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// user/:id
func UserDelete(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
