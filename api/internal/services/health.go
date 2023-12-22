package services

import (
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
)

type HealthService struct {
}

func NewHealthService() HealthService {
	return HealthService{}
}

func (s *HealthService) Add() []routey.Route {
	return []routey.Route{
		{
			Path:          "/api/v1",
			Params:        "",
			Method:        routey.Get,
			HandlerFunc:   s.List(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1",
			Params:        "/:id",
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

func (s *HealthService) List() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *HealthService) Get() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *HealthService) Post() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *HealthService) Put() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *HealthService) Patch() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *HealthService) Delete() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *HealthService) Head() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *HealthService) Options() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}
