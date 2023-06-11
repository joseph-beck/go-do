package todohandler

import (
	"go-do/internal/database"
	"go-do/internal/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodoGet(s database.TodoStoreReader) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := make([]todo.TaskModel, 0)
		s.ReadAll(&t)
		c.JSON(http.StatusOK, t)
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
