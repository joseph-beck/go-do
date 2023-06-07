package router

import "github.com/gin-gonic/gin"

type Routes []gin.RouteInfo

var routes = Routes{
	gin.RouteInfo{
		Method:      "GET",
		Path:        "/",
		Handler:     "Ping",
		HandlerFunc: nil,
	},
}
