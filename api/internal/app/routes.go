package app

import (
	"go-do/internal/pinghandler"
	"go-do/internal/router"
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
		},
	}
}
