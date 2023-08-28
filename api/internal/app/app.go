package app

import (
	"fmt"
	"go-do/internal/database"
	"go-do/internal/router"
	"go-do/pkg/util"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

// Stores the router
var r *router.Router

var s *database.Store

// Runs the app
//
//   - Creates the router.
//   - Registers routes.
//   - Runs the engine.
//   - Waits for interrupt before shutting down app.
func Run() {
	log.Println("Starting app")

	err := godotenv.Load()
	util.ErrOut(err)

	go shutdown()

	s = database.New(database.SQLiteDb())

	r = router.New()
	r.RegisterRoutes(makeRoutes())
	r.NoRoute(reverseProxy())
	r.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down app")
}

// Shuts down the app.
//
//   - Closes Store if not nil.
func shutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("")
	log.Println("Shutting down app")

	log.Println("Closing db conn")
	err := s.Close()
	util.ErrLog(err)

	os.Exit(0)
}
