package services

import (
	"go-do/internal/database"
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
)

type UserService struct {
	db *database.Store
}

func NewUserService(db *database.Store) UserService {
	return UserService{
		db: db,
	}
}

func (s *UserService) Add() []routey.Route {
	return []routey.Route{
		// sign in and sign up
		{
			Path:          "/api/v1/users/signin",
			Params:        "",
			Method:        routey.Get,
			HandlerFunc:   s.SignIn(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/users/signup",
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
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/users",
			Params:        "/:id",
			Method:        routey.Get,
			HandlerFunc:   s.Get(),
			DecoratorFunc: nil,
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
			DecoratorFunc: nil,
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

func (s *UserService) List() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Render(http.StatusOK, "users")
	}
}

func (s *UserService) Get() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Render(http.StatusOK, "user")
	}
}

func (s *UserService) Post() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusOK)
	}
}

func (s *UserService) Put() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusOK)
	}
}

func (s *UserService) Patch() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusOK)
	}
}

func (s *UserService) Delete() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusOK)
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
