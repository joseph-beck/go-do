package router

import (
	"fmt"

	routey "github.com/joseph-beck/routey/pkg/router"
)

type Router struct {
	App *routey.App
}

func New() Router {
	return Router{
		App: routey.New(),
	}
}

func (r *Router) Register(rs []routey.Route) {
	for _, re := range rs {
		r.App.Route(re)
	}
}

func (r *Router) NoRoute(f routey.HandlerFunc) {
	fmt.Println("todo")
}

func (r *Router) Run() {
	r.App.Run()
}

func (r *Router) Shutdown() {
	r.App.Shutdown()
}
