package database

import (
	"go-do/internal/todo"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var s Store

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load env")
	}

	s = *MakeStore()

	err = s.Ping()
	if err != nil {
		log.Fatalln("db failed to connect")
	}
}

func TestStoreRead(t *testing.T) {
	q := todo.TaskModel{Id: 1}
	s.Read(&q)

	assert.NotNil(t, q)
	log.Fatalln(q.Name)
}

func TestStoreAdd(t *testing.T) {
	q := todo.TaskModel{
		Name: "Test Task",
	}
	s.Add(q)
}

func TestStoreUpdate(t *testing.T) {

}

func TestStoreDelete(t *testing.T) {

}

func TestStoreCheck(t *testing.T) {
	e := s.Check(todo.MakeTask())
	assert.True(t, e)
}
