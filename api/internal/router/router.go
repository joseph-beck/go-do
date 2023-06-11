package router

import (
	"go-do/internal/database"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
	Store  *database.TodoStore
}

func MakeRouter(t string) *Router {
	return &Router{
		Engine: gin.Default(),
		Store:  database.MakeTodoStore(t),
	}
}

func (r *Router) Run() {
	r.Engine.Run()
}
