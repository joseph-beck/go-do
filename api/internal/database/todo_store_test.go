package database

import (
	"go-do/internal/todo"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoStore(t *testing.T) {
	err := ts.Ping()
	if err != nil {
		log.Fatalln("db failed to connect")
	}
}

func TestTodoStoreRead(t *testing.T) {
	q := todo.TaskModel{Id: 1}
	ts.Read(&q, "tasks")

	assert.NotNil(t, q)
	log.Fatalln(q.Name)
}

func TestTodoStoreAdd(t *testing.T) {
	q := todo.TaskModel{
		Name: "Test Task",
	}
	ts.Add(q, "tasks")
}

func TestTodoStoreUpdate(t *testing.T) {

}

func TestTodoStoreDelete(t *testing.T) {

}

func TestTodoStoreCheck(t *testing.T) {
	tm := todo.TaskModel{
		Id: 1,
	}

	e := ts.Check(tm, "tasks")
	assert.True(t, e)
}