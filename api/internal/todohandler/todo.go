package todohandler

import (
	"go-do/internal/database"
	"go-do/internal/todo"
	"go-do/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TodoGet is the HandlerFunc for getting a Task or a table of Tasks.
//
// Example usage:
//   - /todo : returns the results of the tasks table.
//   - /todo?table='table' : returns the result of the given table.
//   - /todo?table='table'&id='id' : returns the result of an id from a given table.
func TodoGet(
	g database.StoreGetter[todo.TaskModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := make([]todo.TaskModel, 0)

		table := c.Query("table")
		if table == "" {
			table = "tasks"
		}

		if g.Check(todo.TaskModel{}, table) {
			q := c.Query("id")
			id, err := strconv.Atoi(q)
			util.ErrLog(err)

			if id != 0 {
				u := todo.TaskModel{Id: id}
				g.Scan(&u, table)
				l = append(l, u)
			} else {
				g.Read(&l, table)
			}
		}

		c.JSON(http.StatusOK, l)
	}
}

func TodoPost(
	s database.StoreAdder[todo.TaskModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func TodoPatch(
	s database.StoreUpdater[todo.TaskModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func TodoDelete(
	s database.StoreDeleter[todo.TaskModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
