package app

import (
	"go-do/internal/database"
	"go-do/internal/router"
	"go-do/pkg/util"
	"log"

	"github.com/joho/godotenv"
)

// Stores the router
var r router.Router

// Stores the store
var s database.Store

// Runs the app
//   - Creates the router.
//   - Registers routes.
//   - Runs the engine.
//   - Waits for interrupt before shutting down app.
func Run() {
	log.Println("Starting app")

	err := godotenv.Load()
	util.ErrOut(err)

	s = database.New(database.SQLiteDb())

	r = router.New()
	r.Register(routes)
	r.NoRoute(reverseProxy())
	go shutdown()
	r.Run()
}

func shutdown() {
	r.Shutdown()
	err := s.Close()
	if err != nil {
		panic(err)
	}
}
