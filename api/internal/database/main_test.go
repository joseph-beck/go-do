package database

import (
	"go-do/pkg/util"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

var bs *BaseStore
var ts *TodoStore

// Loads .env and if not errors out.
//
// Also initializes both stores.
func TestMain(m *testing.M) {
	err := godotenv.Load()
	util.ErrOut(err)

	//bs = MakeBaseStore()
	ts = MakeTodoStore()
}

// Tests the connection to the BaseStore via the Ping method.
//
// Ping returns an error, if it is not nil then the Ping has failed.
// This is checked with the util module.
func TestBaseStore(t *testing.T) {
	err := bs.Ping()
	if err != nil {
		log.Fatalln("db failed to connect")
	}
}

// Tests the connection to the TodoStore via the Ping method.
//
// Ping returns an error, if it is not nil then the Ping has failed.
// This is checked with the util module.
func TestTodoStore(t *testing.T) {
	err := ts.Ping()
	util.ErrMsg(err, "failed to ping db")
}
