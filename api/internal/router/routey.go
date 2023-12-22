package router

import (
	"fmt"
	"go-do/internal/services"

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

func (r *Router) Route(re routey.Route) {
	r.App.Route(re)
}

func (r *Router) NoRoute(f routey.HandlerFunc) {
	fmt.Println("todo")
}

func (r *Router) Register(rs []routey.Route) {
	for _, re := range rs {
		r.App.Route(re)
	}
}

func (r *Router) Service(s services.Service) {
	rs := s.Add()
	r.Register(rs)
}

func (r *Router) Services(s []services.Service) {
	for _, se := range s {
		rs := se.Add()
		r.Register(rs)
	}
}

func (r *Router) Run() {
	r.App.Run()
}

func (r *Router) Shutdown() {
	r.App.Shutdown()
}
