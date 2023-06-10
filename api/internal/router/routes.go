package router

import (
	"go-do/internal/pinghandler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      Method
	Path        string
	Handler     string
	HandlerFunc gin.HandlerFunc
}

type Routes struct {
	RouteInfo []Route
}

var ping = Routes{
	[]Route{
		Route{
			Method:      Get,
			Path:        "/",
			Handler:     "ping",
			HandlerFunc: pinghandler.PingGet(),
		},
	},
}
