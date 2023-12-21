package handlers

import (
	"go-do/internal/database"
	"go-do/internal/models"
	"go-do/pkg/util"
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
)

// /user/login...
func UserLogon(s *database.Store) routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

// /user/signup...
func UserSignup(s *database.Store) routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

// /user
func UserList(s *database.Store) routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		r, err := s.ListUser()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// /user/:id
func UserGet(s *database.Store) routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		i, err := c.ParamInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		r, err := s.GetUser(uint(i))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// /user
func UserPost(s *database.Store) routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		m := models.User{}
		c.ShouldBindJSON(&m)
		m.Password = util.Sha256Hash(m.Password)

		if m.ID != 0 {
			e := s.CheckUser(m.ID)
			if e {
				c.Status(http.StatusBadRequest)
				return
			}
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
func UserPut(s *database.Store) routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		m := models.User{}
		c.ShouldBindJSON(&m)
		m.Password = util.Sha256Hash(m.Password)

		e := s.CheckUser(m.ID)
		if !e || m.ID == 0 {
			err := s.AddUser(m)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}
		} else {
			err := s.UpdateUser(m)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}
		}

		c.Status(http.StatusNoContent)
	}
}

// /user
func UserPatch(s *database.Store) routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		m := models.User{}
		c.ShouldBindJSON(&m)
		m.Password = util.Sha256Hash(m.Password)

		e := s.CheckUser(m.ID)
		if !e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.UpdateUser(m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// user/:id
func UserDelete(s *database.Store) routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		i, err := c.ParamInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.CheckUser(uint(i))
		if !e {
			c.Status(http.StatusNoContent)
			return
		}

		err = s.DeleteUser(uint(i))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}
