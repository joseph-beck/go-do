package services

import (
	"go-do/internal/database"
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
	"github.com/joseph-beck/routey/pkg/status"
)

type PingService struct {
	db *database.Store
}

func NewPingService(db *database.Store) PingService {
	return PingService{
		db: db,
	}
}

func (s *PingService) Add() []routey.Route {
	return []routey.Route{
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

func (s *PingService) Get() routey.HandlerFunc {
	return func(c *routey.Context) {
		err := s.db.Ping()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Status(status.OK)
	}
}

func (s *PingService) Post() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *PingService) Put() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *PingService) Patch() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *PingService) Delete() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *PingService) Head() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}

func (s *PingService) Options() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(status.NotFound)
	}
}
