package app

import (
	"go-do/internal/pinghandler"
	"go-do/internal/router"
	"go-do/internal/todohandler"
)

func makeRoutes() *router.Routes {
	return &router.Routes{
		RouteInfo: []router.Route{
			{
				Name:        "GetPing",
				Method:      router.Get,
				Path:        "/",
				Handler:     "ping",
				HandlerFunc: pinghandler.PingGet(r.Store),
			},
			{
				Name:        "GetTodo",
				Method:      router.Get,
				Path:        "/",
				Handler:     "todo",
				HandlerFunc: todohandler.TodoGet(r.Store),
			},
		},
	}
}
