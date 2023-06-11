package database

import (
	"go-do/internal/todo"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	err := rs.Ping()
	if err != nil {
		log.Fatalln("db failed to connect")
	}
}

func TestStoreRead(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load env")
	}

	s := *MakeStore()

	q := todo.TaskModel{Id: 1}
	s.Scan(&q)

	assert.NotNil(t, q)
	log.Fatalln(q.Name)
}

func TestStoreAdd(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load env")
	}

	s := *MakeStore()

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
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load env")
	}

	s := *MakeStore()

	e := s.Check(todo.MakeTask())
	assert.True(t, e)
}
