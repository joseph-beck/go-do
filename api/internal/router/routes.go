package router

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Name        string
	Method      Method
	Path        string
	Handler     string
	HandlerFunc gin.HandlerFunc
}

type Routes struct {
	RouteInfo []Route
}
