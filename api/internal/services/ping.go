package services

import (
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
)

type PingService struct {
}

func NewPingService() PingService {
	return PingService{}
}

func (s *PingService) Add() []routey.Route {
	return []routey.Route{
		{
			Path:          "/api/v1/ping",
			Params:        "",
			Method:        routey.Get,
			HandlerFunc:   s.List(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/ping",
			Params:        "",
			Method:        routey.Get,
			HandlerFunc:   s.Get(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/ping",
			Params:        "",
			Method:        routey.Post,
			HandlerFunc:   s.Post(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/ping",
			Params:        "",
			Method:        routey.Put,
			HandlerFunc:   s.Put(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/ping",
			Params:        "",
			Method:        routey.Patch,
			HandlerFunc:   s.Patch(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/ping",
			Params:        "",
			Method:        routey.Delete,
			HandlerFunc:   s.Delete(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/ping",
			Params:        "",
			Method:        routey.Head,
			HandlerFunc:   s.Head(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/ping",
			Params:        "",
			Method:        routey.Options,
			HandlerFunc:   s.Options(),
			DecoratorFunc: nil,
		},
	}
}

func (s *PingService) List() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *PingService) Get() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *PingService) Post() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *PingService) Put() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *PingService) Patch() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *PingService) Delete() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *PingService) Head() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *PingService) Options() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}
