package todohandler

import (
	"go-do/internal/database"
	"go-do/internal/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodoGet(s database.TodoStoreReader, v database.TodoStoreChecker) gin.HandlerFunc {
	return func(c *gin.Context) {
		q := c.Query("table")
		if q == "" {
			q = "tasks"
		}

		var t todo.TaskModel
		l := make([]todo.TaskModel, 0)
		if v.Check(t, q) {
			s.ReadAll(&l, q)
		}
		c.JSON(http.StatusOK, l)
	}
}

func TodoPost(s *database.Store) gin.HandlerFunc {
	return nil
}

func TodoPatch(s *database.Store) gin.HandlerFunc {
	return nil
}

func TodoDelete(s *database.Store) gin.HandlerFunc {
	return nil
}
