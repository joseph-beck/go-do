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
