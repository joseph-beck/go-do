package app

import (
	"go-do/internal/router"
	"go-do/pkg/util"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

// Stores the router
var r *router.Router

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

	r = router.MakeRouter()
	r.RegisterRoutes(*makeRoutes())
	r.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down app")
	shutdown()
}

// Shuts down the app.
//
//   - Closes Store if not nil.
func shutdown() {
	if r.Store != nil {
		r.Store.Close()
	}
}
