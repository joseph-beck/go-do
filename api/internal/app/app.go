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

	t := &todo.TaskModel{
		Id: 4,
	}
	r.Store.Read(t)
	log.Println(t.Str())

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
