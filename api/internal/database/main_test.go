package database

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

var bs BaseStore
var ts TodoStore

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load env")
	}

	bs = *MakeBaseStore()
	ts = *MakeTodoStore()
}
