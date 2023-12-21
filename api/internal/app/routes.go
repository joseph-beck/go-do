package app

import (
	"go-do/internal/handlers"
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
)

var routes = []routey.Route{
	// Ping Handlers
	{
		Method:      routey.Get,
		Path:        "/ping",
		Params:      "",
		HandlerFunc: handlers.PingGet(&s),
	},
	// Task Handlers
	{
		Method:      routey.Get,
		Path:        "/task",
		Params:      "/:list",
		HandlerFunc: handlers.TaskList(&s),
	},
	{
		Method:      routey.Get,
		Path:        "/task",
		Params:      "/:list/:task",
		HandlerFunc: handlers.TaskGet(&s),
	},
	{
		Method:      routey.Post,
		Path:        "/task",
		Params:      "/:list",
		HandlerFunc: handlers.TaskPost(&s),
	},
	{
		Method:      routey.Put,
		Path:        "/task",
		Params:      "/:list",
		HandlerFunc: handlers.TaskPut(&s),
	},
	{
		Method:      routey.Patch,
		Path:        "/task",
		Params:      "/:list",
		HandlerFunc: handlers.TaskPatch(&s),
	},
	{
		Method:      routey.Delete,
		Path:        "/task",
		Params:      "/:list/:task",
		HandlerFunc: handlers.TaskDelete(&s),
	},
	// User Handlers
	{
		Method:      routey.Get,
		Path:        "/user",
		Params:      "",
		HandlerFunc: handlers.UserList(&s),
	},
	{
		Method:      routey.Get,
		Path:        "/user",
		Params:      "/:id",
		HandlerFunc: handlers.UserGet(&s),
	},
	{
		Method:      routey.Post,
		Path:        "/user",
		Params:      "/:id",
		HandlerFunc: handlers.UserPost(&s),
	},
	{
		Method:      routey.Put,
		Path:        "/user",
		Params:      "/:list",
		HandlerFunc: handlers.UserPut(&s),
	},
	{
		Method:      routey.Patch,
		Path:        "/user",
		Params:      "",
		HandlerFunc: handlers.UserPatch(&s),
	},
	{
		Method:      routey.Delete,
		Path:        "/user",
		Params:      "/:id",
		HandlerFunc: handlers.UserDelete(&s),
	},
}

func reverseProxy() routey.HandlerFunc {
	return func(c *routey.Context) {
		c.Redirect(http.StatusAccepted, "http://localhost:3000")
	}
}
