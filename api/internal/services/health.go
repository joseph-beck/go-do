package services

import (
	"go-do/internal/database"

	routey "github.com/joseph-beck/routey/pkg/router"
	"github.com/joseph-beck/routey/pkg/status"
)

type HealthService struct {
	db *database.Store
}

func NewHealthService(db *database.Store) HealthService {
	return HealthService{
		db: db,
	}
}

func (s *HealthService) Add() []routey.Route {
	return []routey.Route{
		{
			Path:          "/api/v1",
			Params:        "",
			Method:        routey.Get,
			HandlerFunc:   s.Get(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1",
			Params:        "",
			Method:        routey.Post,
			HandlerFunc:   s.Post(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1",
			Params:        "",
			Method:        routey.Put,
			HandlerFunc:   s.Put(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1",
			Params:        "",
			Method:        routey.Patch,
			HandlerFunc:   s.Patch(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1",
			Params:        "",
			Method:        routey.Delete,
			HandlerFunc:   s.Delete(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1",
			Params:        "",
			Method:        routey.Head,
			HandlerFunc:   s.Head(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1",
			Params:        "",
			Method:        routey.Options,
			HandlerFunc:   s.Options(),
			DecoratorFunc: nil,
		},
	}
}

func (s *HealthService) Get() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.OK)
	}
}

func (s *HealthService) Post() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *HealthService) Put() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *HealthService) Patch() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *HealthService) Delete() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *HealthService) Head() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *HealthService) Options() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}
