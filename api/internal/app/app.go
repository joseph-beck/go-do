package app

import (
	p "go-do/internal/pinghandler"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/ping", p.PingGet())

	r.Run()
}

func Shutdown() {

}
