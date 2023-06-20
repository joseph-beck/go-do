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
		if err != nil {
			i = 0
		}

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

// TodoPost is the HandlerFunc for posting a Task to the database.
//
// The Task is posted to the given table.
//
// Example usage:
//
//   - /todo?table=tasks : this is the html query for this function
//     however the body of the request must contain the Task submitted.
//
//   - body :
//
//     {
//
//     "name" : "Example Task",
//
//     "description" : "This is an Example Task",
//
//     "complete" : false,
//
//     "deadline" : "00/00/0000-00:00"
//
//     }
//
//     an id can also be given if the sender wishes to specify this.
func TodoPost(
	s database.Storer[todo.TaskModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.Query("table")

		if !s.Check(todo.TaskModel{Id: 0}, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		b := todo.TaskPost{}
		c.Bind(&b)
		q := b.ToTaskModel()
		s.Add(q, t)

		c.Status(http.StatusNoContent)
	}
}

// TodoPost is the HandlerFunc for patching a Task already in the database.
//
// The table is given and the body is used to patch.
//
// Example usage:
//
//   - /todo?table=tasks : this is the html query for this function
//     however the body of the request must contain the Task submitted.
//
//   - body :
//
//     {
//
//     "id" : 70,
//
//     "name" : "Example Task",
//
//     "description" : "This is an Example Task",
//
//     "complete" : false,
//
//     "deadline" : "00/00/0000-00:00"
//
//     }
//
//     an id must be given here otherwise a StatusBadRequest will be given.
func TodoPatch(
	s database.Storer[todo.TaskModel],
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.Query("table")

		if !s.Check(todo.TaskModel{Id: 0}, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		b := todo.TaskPost{}
		c.Bind(&b)
		q := b.ToTaskModel()

		if !s.Check(q, t) {
			c.Status(http.StatusBadRequest)
			return
		}

		s.Update(q, t)

		c.Status(http.StatusNoContent)
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
