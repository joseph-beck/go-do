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
				Path:        "/ping",
				Handler:     "",
				HandlerFunc: pinghandler.PingGet(r.Store),
			},
			{
				Name:        "GetTodo",
				Method:      router.Get,
				Path:        "/todo",
				Handler:     "",
				HandlerFunc: todohandler.TodoGet(r.Store),
			},
		},
	}
}
