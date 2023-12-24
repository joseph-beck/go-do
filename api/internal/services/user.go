package services

import (
	"fmt"
	"go-do/internal/database"
	"go-do/internal/models"
	"go-do/pkg/util"
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
)

type UserService struct {
	db    *database.Store
	table string
}

func NewUserService(db *database.Store) UserService {
	return UserService{
		db:    db,
		table: "users",
	}
}

func (s *UserService) Add() []routey.Route {
	return []routey.Route{
		// sign in and sign up
		{
			Path:          "/api/v1/signin",
			Params:        "",
			Method:        routey.Get,
			HandlerFunc:   s.SignIn(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/signup",
			Params:        "",
			Method:        routey.Post,
			HandlerFunc:   s.SignUp(),
			DecoratorFunc: nil,
		},
		// generic end points
		{
			Path:          "/api/v1/users",
			Params:        "",
			Method:        routey.Get,
			HandlerFunc:   s.List(),
			DecoratorFunc: s.Authorization(),
		},
		{
			Path:          "/api/v1/users",
			Params:        "/:id",
			Method:        routey.Get,
			HandlerFunc:   s.Get(),
			DecoratorFunc: s.Authorization(),
		},
		{
			Path:          "/api/v1/users",
			Params:        "",
			Method:        routey.Post,
			HandlerFunc:   s.Post(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/users",
			Params:        "",
			Method:        routey.Put,
			HandlerFunc:   s.Put(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/users",
			Params:        "",
			Method:        routey.Patch,
			HandlerFunc:   s.Patch(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/users",
			Params:        "/:id",
			Method:        routey.Delete,
			HandlerFunc:   s.Delete(),
			DecoratorFunc: s.Authorization(),
		},
		{
			Path:          "/api/v1/users",
			Params:        "",
			Method:        routey.Head,
			HandlerFunc:   s.Head(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/users",
			Params:        "",
			Method:        routey.Options,
			HandlerFunc:   s.Options(),
			DecoratorFunc: nil,
		},
	}
}

func (s *UserService) SignIn() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Render(http.StatusOK, "sign in")
	}
}

func (s *UserService) SignUp() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Render(http.StatusOK, "sign up")
	}
}

// GET /api/v1/users?limit=0&offset=0
func (s *UserService) List() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		l, _ := c.QueryInt("limit")
		o, _ := c.QueryInt("offset")

		var m []models.User
		err := s.db.Read(&m, s.table)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		if l == 0 && o == 0 {
			c.JSON(http.StatusOK, m)
			return
		}

		e := o + l
		if e > len(m) {
			e = len(m)
		}

		r := m[o:e]
		c.JSON(http.StatusOK, r)
	}
}

// GET /api/v1/users/:id
func (s *UserService) Get() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		i, err := c.ParamInt("id")
		if err != nil {
			c.Status(http.StatusBadGateway)
			return
		}

		m := models.User{
			Model: models.Model{ID: uint(i)},
		}
		err = s.db.Get(&m, s.table)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, m)
	}
}

// POST /api/v1/users
func (s *UserService) Post() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		m := models.User{}
		c.ShouldBindJSON(&m)
		m.Password = util.Sha256Hash(m.Password)

		if m.ID != 0 {
			if s.db.Check(&m, s.table) {
				fmt.Println("yo")
				c.Status(http.StatusBadRequest)
				return
			}
		}

		err := s.db.Add(&m, s.table)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// PUT /api/v1/users
func (s *UserService) Put() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		m := models.User{}
		c.ShouldBindJSON(&m)
		m.Password = util.Sha256Hash(m.Password)

		if m.ID != 0 {
			if s.db.Check(&models.User{Model: models.Model{ID: uint(m.ID)}}, s.table) {
				err := s.db.Update(&m, s.table)
				if err != nil {
					fmt.Println(err, "1")
					c.Status(http.StatusBadRequest)
					return
				}
				c.Status(http.StatusNoContent)
				return
			}
		}

		err := s.db.Add(&m, s.table)
		if err != nil {
			fmt.Println(err, "2")
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// PATCH /api/v1/users
func (s *UserService) Patch() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		m := models.User{}
		c.ShouldBindJSON(&m)
		m.Password = util.Sha256Hash(m.Password)

		if !s.db.Check(&models.User{Model: models.Model{ID: uint(m.ID)}}, s.table) {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.db.Update(&m, s.table)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// DELETE /api/v1/users
func (s *UserService) Delete() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		i, err := c.ParamInt("id")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		m := models.User{Model: models.Model{ID: uint(i)}}
		if !s.db.Check(&m, s.table) {
			c.Status(http.StatusBadRequest)
			return
		}

		err = s.db.Delete(&m, s.table)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func (s *UserService) Head() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *UserService) Options() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *UserService) Authorization() routey.DecoratorFunc {
	return func(f routey.HandlerFunc) routey.HandlerFunc {
		return func(c *routey.Context) {
			a := c.GetHeader("Authorization")
			if !s.db.Check(&models.Admin{Token: a}, "admins") {
				c.Status(http.StatusForbidden)
				return
			}

			f(c)
		}
	}
}
