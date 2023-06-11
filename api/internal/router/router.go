package router

import (
	"go-do/internal/database"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	Store  *database.TodoStore
}

func MakeRouter() *Router {
	return &Router{
		Engine: gin.Default(),
		Store:  database.MakeTodoStore(),
	}
}

func (r *Router) Run() {
	r.Engine.Run()
}

func (r *Router) RegisterRoutes(rl Routes) {
	for _, route := range rl.RouteInfo {
		switch route.Method {
		case Get:
			r.Engine.GET(route.Path+route.Handler, route.HandlerFunc)
		case Post:
			r.Engine.POST(route.Path+route.Handler, route.HandlerFunc)
		case Patch:
			r.Engine.PATCH(route.Path+route.Handler, route.HandlerFunc)
		case Delete:
			r.Engine.DELETE(route.Path+route.Handler, route.HandlerFunc)
		}
	}
}
