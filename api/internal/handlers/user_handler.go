package handlers

import (
	"go-do/internal/database"
	"go-do/internal/models"
	"go-do/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// /user/login...
func UserLogon(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /user/signup...
func UserSignup(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /user
func UserList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, err := s.ListUser()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// /user/:id
func UserGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		r, err := s.GetUser(uint(id))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
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
