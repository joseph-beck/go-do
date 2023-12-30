package services

import (
	"errors"
	"go-do/internal/database"
	"go-do/internal/models"
	"go-do/pkg/util"
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
)

type TaskService struct {
	db *database.Store
}

func NewTaskService(db *database.Store) TaskService {
	return TaskService{
		db: db,
	}
}

func (s *TaskService) Add() []routey.Route {
	return []routey.Route{
		{
			Path:          "/api/v1/tasks",
			Params:        "/:list",
			Method:        routey.Get,
			HandlerFunc:   s.List(),
			DecoratorFunc: nil,
		},
		{
			Path:          "/api/v1/tasks",
			Params:        "/:list/:task",
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
			Params:        "/:list",
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
	type Response struct {
		Tasks     []models.Task `json:"tasks"`
		TaskCount int           `json:"task_count"`
	}

	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		t, err := c.Param("list")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		err = initTaskTable(s.db, t)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		var m []models.Task
		err = s.db.Read(&m, t)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		r := Response{
			Tasks:     m,
			TaskCount: len(m),
		}

		c.JSON(http.StatusOK, r)
	}
}

func (s *TaskService) Get() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Header("Content-Type", "application/json")

		t, err := c.Param("list")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		i, err := c.ParamInt("task")
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		err = initTaskTable(s.db, t)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.db.Contains(&models.Task{Model: models.Model{ID: uint(i)}}, t)
		if !e {
			c.Status(http.StatusNotFound)
			return
		}

		r := models.Task{Model: models.Model{ID: uint(i)}}
		err = s.db.Get(&r, t)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

func (s *TaskService) Post() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusOK)
	}
}

func (s *TaskService) Put() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusOK)
	}
}

func (s *TaskService) Patch() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusOK)
	}
}

func (s *TaskService) Delete() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Status(http.StatusOK)
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

func (s *TaskService) Authorization() routey.DecoratorFunc {
	return func(f routey.HandlerFunc) routey.HandlerFunc {
		return func(c *routey.Context) {
			a := c.GetHeader("Authorization")
			if !s.db.Contains(&models.Admin{Token: a}, "admins") {
				c.Status(http.StatusForbidden)
				return
			}

			f(c)
		}
	}
}

// Create task table if it does not exist
func initTaskTable(s *database.Store, n string) error {
	var m models.Task
	err := s.CreateTable(n, &m)
	if errors.Is(err, util.ErrTableAlreadyExists) {
		return nil
	}

	return err
}
