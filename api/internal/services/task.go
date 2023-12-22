package services

import (
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
)

type TaskService struct {
}

func NewTaskService() TaskService {
	return TaskService{}
}

func (s *TaskService) Add() []routey.Route {
	return []routey.Route{
		{
			Path:          "/api/v1/tasks",
			Params:        "",
			Method:        routey.Get,
			HandlerFunc:   s.List(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/tasks",
			Params:        "/:id",
			Method:        routey.Get,
			HandlerFunc:   s.Get(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/tasks",
			Params:        "",
			Method:        routey.Post,
			HandlerFunc:   s.Post(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/tasks",
			Params:        "",
			Method:        routey.Put,
			HandlerFunc:   s.Put(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/tasks",
			Params:        "",
			Method:        routey.Patch,
			HandlerFunc:   s.Patch(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/tasks",
			Params:        "/:id",
			Method:        routey.Delete,
			HandlerFunc:   s.Delete(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/tasks",
			Params:        "",
			Method:        routey.Head,
			HandlerFunc:   s.Head(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/tasks",
			Params:        "",
			Method:        routey.Options,
			HandlerFunc:   s.Options(),
			DecoratorFunc: nil,
		},
	}
}

func (s *TaskService) List() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *TaskService) Get() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *TaskService) Post() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *TaskService) Put() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *TaskService) Patch() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *TaskService) Delete() routey.HandlerFunc {
	return func(c *routey.Context) {

	}
}

func (s *TaskService) Head() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}

func (s *TaskService) Options() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusNotFound)
	}
}
