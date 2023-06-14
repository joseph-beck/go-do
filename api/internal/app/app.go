package app

import (
	"go-do/internal/router"
	"go-do/internal/todo"
	"go-do/pkg/util"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

var r *router.Router

func Run() {
	log.Println("Starting app")

	err := godotenv.Load()
	util.ErrOut(err)

	r = router.MakeRouter()

	log.Println("ADDING")
	a := todo.TaskModel{}
	r.Store.Add(a, "tasks")
	log.Println("Done")

	r.RegisterRoutes(*makeRoutes())
	r.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down app")
	shutdown()
}

func shutdown() {
	if r != nil {
		r.Store.Close()
	}
}
