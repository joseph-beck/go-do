package app

import (
	"go-do/internal/database"
	"go-do/internal/router"
	"go-do/internal/services"
	"go-do/pkg/util"
	"log"

	"github.com/joho/godotenv"
	routey "github.com/joseph-beck/routey/pkg/router"
)

// Stores the router
var r router.Router

// Stores the store
var s database.Store

// Store the instances of the Services
var (
	healthService = services.NewHealthService(&s)
	pingService   = services.NewPingService(&s)
	taskService   = services.NewTaskService(&s)
	userService   = services.NewUserService(&s)
)

// Store the Service instances in this array
var service = []routey.Service{
	&healthService,
	&pingService,
	&taskService,
	&userService,
}

// Runs the app
//   - Creates the router.
//   - Registers routes & services.
//   - Runs the App.
//   - Waits for interrupt before shutting down app.
func Run() {
	log.Println("Starting app")

	err := godotenv.Load()
	util.ErrOut(err)

	s = database.New(database.SQLiteDb())
	err = s.AutoMigrate()
	util.ErrOut(err)

	r = router.New()
	r.Register(routes)
	r.Services(service)
	r.NoRoute(reverseProxy())
	go shutdown()
	r.Run()
}

// Exit the whole program
func shutdown() {
	r.Shutdown()
	err := s.Close()
	util.ErrOut(err)
}
