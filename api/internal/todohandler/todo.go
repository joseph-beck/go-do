package todohandler

import (
	"go-do/internal/database"
	"go-do/internal/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodoGet(s database.StoreReader[todo.TaskModel], v database.StoreChecker[todo.TaskModel]) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Query("table")
		if q == "" {
			q = "tasks"
		}

		l := make([]todo.TaskModel, 0)
		if v.Check(todo.TaskModel{}, q) {
			s.Read(&l, q)
		}
		c.JSON(http.StatusOK, l)
	}
}

func TodoPost(s database.StoreAdder[todo.TaskModel]) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func TodoPatch(s database.StoreUpdater[todo.TaskModel]) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func TodoDelete(s database.StoreDeleter[todo.TaskModel]) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
