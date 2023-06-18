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
// If the queried Table queried does not exist then a http.StatusBadRequest is given.
//
// Example usage:
//   - /todo : returns the results of the tasks table.
//   - /todo?table='table' : returns the result of the given table.
//   - /todo?table='table'&id='id' : returns the result of an id from a given table.
func TodoGet(
	s database.Storer[todo.TaskModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		l := make([]todo.TaskModel, 0)

		t := c.Query("table")
		if t == "" {
			t = "tasks"
		}

		if !s.Check(todo.TaskModel{Id: 0}, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		q := c.Query("id")
		i, err := strconv.Atoi(q)
		util.ErrLog(err)

		if i != 0 {
			u := todo.TaskModel{Id: i}
			s.Scan(&u, t)
			l = append(l, u)
		} else {
			s.Read(&l, t)
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

// TodoDelete is the HandlerFunc for deleting a Task.
//
// If neither the user or the table exist then a http.StatusBadRequest is given.
//
// Example usage:
//   - /todo?table='table'&id='id' : deletes the Task Model at the given table and id.
func TodoDelete(
	s database.Storer[todo.TaskModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.Query("table")

		if !s.Check(todo.TaskModel{Id: 0}, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		q := c.Query("id")
		i, err := strconv.Atoi(q)
		util.ErrLog(err)

		u := todo.TaskModel{Id: i}

		if i == 0 || s.Check(u, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		s.Delete(u, t)

		c.Status(http.StatusNoContent)
	}
}
