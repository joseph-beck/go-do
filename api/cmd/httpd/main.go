package main

import(
	"go-do/internal/app"
)

func main() {
	app.Run()
	defer app.Shutdown()
}