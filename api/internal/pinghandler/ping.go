package pinghandler

import (
	"go-do/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingGet(p database.StorePinger) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := p.Ping()
		r := "pong"
		if err != nil {
			r = "boom"
		}

		c.JSON(http.StatusOK, map[string]string{
			"ping": r,
		})
	}
}
