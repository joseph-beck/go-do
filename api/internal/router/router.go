package router

import (
	"fmt"

	routey "github.com/joseph-beck/routey/pkg/router"
)

// Stores the Routey App
type Router struct {
	App *routey.App
}

// Creates a new Router
func New() Router {
	return Router{
		App: routey.New(),
	}
}

// Adds a route to the Router
func (r *Router) Route(re routey.Route) {
	r.App.Route(re)
}

// Function for NoRoute
func (r *Router) NoRoute(f routey.HandlerFunc) {
	fmt.Println("todo")
}

// Register an array of routey.Route
func (r *Router) Register(rs []routey.Route) {
	for _, re := range rs {
		r.App.Route(re)
	}
}

// Register a Service with Routey
func (r *Router) Service(s routey.Service) {
	r.App.Service(s)
}

// Register an array of Services
func (r *Router) Services(s []routey.Service) {
	for _, se := range s {
		r.App.Service(se)
	}
}

// Run the App
func (r *Router) Run() {
	r.App.Run()
}

// Shutdown the App
func (r *Router) Shutdown() {
	r.App.Shutdown()
}
