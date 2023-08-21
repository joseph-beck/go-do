package handlers

import (
	"go-do/internal/database"
	"go-do/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TaskList is the HandlerFunc for getting all Tasks from a Table.
//
// If the queried Table queried does not exist then a http.StatusBadRequest is given.
//
// Example usage:
//   - /task/:list : returns the results of the given List.
func TaskList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		l := c.Param("list")
		m, err := s.ListTask(l)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		r := make([]models.TaskPost, len(l))
		for i, task := range m {
			r[i] = task.ToTaskPost()
		}

		c.JSON(http.StatusOK, r)
	}
}

// TaskGet is the HandlerFunc for getting a Task from a Table.
//
// If the queried Table queried does not exist then a http.StatusBadRequest is given.
//
// Example usage:
//   - /task/:list/:task : returns a specific Task from a given List.
func TaskGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		t := c.Param("list")
		i := c.Param("task")

		id, err := strconv.Atoi(i)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.CheckTask(t, id)
		if !e {
			c.Status(http.StatusNoContent)
			return
		}

		m, err := s.GetTask(t, id)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}
		r := m.ToTaskPost()

		c.JSON(http.StatusOK, r)
	}
}

// TaskPost is the HandlerFunc for posting a Task to the database.
//
// The Task is posted to the given table.
//
// Example usage:
//
//   - /task/:list : this is the html query for this function
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
func TaskPost(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		l := c.Param("list")
		b := models.TaskPost{}
		c.ShouldBindJSON(&b)
		m := b.ToTaskModel()

		e := s.CheckTask(l, m.Id)
		if e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.AddTask(l, m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// TaskPost is the HandlerFunc for patching a Task already in the database.
//
// The table is given and the body is used to patch.
//
// Example usage:
//
//   - /task?table=tasks : this is the html query for this function
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
func TaskPut(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
	}
}

// TaskPost is the HandlerFunc for patching a Task already in the database.
//
// The table is given and the body is used to patch.
//
// Example usage:
//
//   - /task?table=tasks : this is the html query for this function
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
func TaskPatch(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		l := c.Param("list")
		b := models.TaskPost{}
		c.ShouldBindJSON(&b)
		m := b.ToTaskModel()

		e := s.CheckTask(l, m.Id)
		if !e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.UpdateTask(l, m)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// TaskDelete is the HandlerFunc for deleting a Task.
//
// If neither the user or the table exist then a http.StatusBadRequest is given.
//
// Example usage:
//   - /task?table='table'&id='id' : deletes the Task Model at the given table and id.
func TaskDelete(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		l := c.Param("list")
		i := c.Param("task")

		id, err := strconv.Atoi(i)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		e := s.CheckTask(l, id)
		if !e {
			c.Status(http.StatusNoContent)
			return
		}

		err = s.DeleteTask(l, id)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}
