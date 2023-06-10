package router

import "github.com/gin-gonic/gin"

type Router struct {
	Engine *gin.Engine
}

func MakeRouter() *Router {
	return &Router{
		Engine: gin.Default(),
	}
}

func (r *Router) Run() {
	r.Engine.Run()
}