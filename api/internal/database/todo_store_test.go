package database

import (
	"go-do/internal/todo"
	"go-do/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests the connection to the TodoStore via the Ping method.
//
// Ping returns an error, if it is not nil then the Ping has failed.
// This is checked with the util module.
func TestTodoStore(t *testing.T) {
	err := ts.Ping()
	util.ErrMsg(err, "failed to ping db")
}

// Tests the Scan method of the TodoStore
//
// Creates a new TaskModel, of Id 1, which exists in the DB already.
// Using Scan, the value from the database is received and mapped to the given TaskModel
func TestTodoStoreScan(t *testing.T) {
	table := "tasks"
	q := todo.TaskModel{Id: 1}
	ts.Scan(&q, table)

	assert.NotNil(t, q)
}

// Tests the Read method of the TodoStore
//
// Creates an empty slice of TaskModel's and this is passed with the table to the
// Read method, which then reads all pieices of data from the table into the slice.
func TestTodoStoreRead(t *testing.T) {
	table := "tasks"
	var tasks []todo.TaskModel
	ts.Read(&tasks, table)

	assert.NotEmpty(t, tasks)
}

// Tests the Add method from the TodoStore.
//
// It creates a new TaskModel, with only a name.
// This is then added, checked if it exists and then deleted.
func TestTodoStoreAdd(t *testing.T) {
	table := "tasks"
	q := todo.TaskModel{
		Name: "Another Task",
	}

	ts.Add(q, table)
	e := ts.Check(q, table)
	assert.True(t, e)

	ts.Delete(q, table)
}

// Tests the Update method from the TodoStore.
//
// Checks against the first added value and the updated value,
// if they are not equal the update has been successful.
func TestTodoStoreUpdate(t *testing.T) {
	table := "tasks"
	q := todo.TaskModel{
		Id:   9999,
		Name: "Testing Update",
	}

	ts.Add(q, table)
	e := ts.Check(q, table)
	assert.True(t, e)

	c := q
	c.Name = "Update Testing"
	ts.Update(c, table)

	a := q
	ts.Scan(&a, table)

	assert.NotEqual(t, q, a)
}

func TestTodoStoreDelete(t *testing.T) {

}

// Tests the Check method from the TodoStore.
//
// Check returns a boolean value based on whether the queried Id exists.
func TestTodoStoreCheck(t *testing.T) {
	tm := todo.TaskModel{
		Id: 1,
	}

	e := ts.Check(tm, "tasks")
	assert.True(t, e)
}
