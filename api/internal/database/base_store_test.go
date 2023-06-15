package database

import (
	"go-do/internal/todo"
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	err := bs.Ping()
	if err != nil {
		log.Fatalln("db failed to connect")
	}
}

func TestStoreScan(t *testing.T) {
	// not really sure if this is working as intended
	m := todo.TaskModel{Id: 1}
	q := reflect.ValueOf(m).Interface()

	bs.Scan(&q, "tasks")
	m = reflect.ValueOf(q).Interface().(todo.TaskModel)

	assert.NotNil(t, m)
	assert.NotNil(t, m.Name)
	assert.NotNil(t, m.Deadline)
	assert.NotNil(t, m.Complete)
}

func TestScoreRead(t *testing.T) {

}

func TestStoreAdd(t *testing.T) {

}

func TestStoreUpdate(t *testing.T) {

}

func TestStoreDelete(t *testing.T) {

}

func TestStoreCheck(t *testing.T) {
	e := bs.Check(todo.MakeTask(), "tasks")
	assert.True(t, e)
}
