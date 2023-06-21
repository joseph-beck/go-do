package pinghandler

import (
	"go-do/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Creates a gin HandlerFunc.
//
// Takes a parameter of a database StorePinger.
//
// Pings the database, if no error is returned the response is ping : pong,
// otherwise it is ping : boom.
//
// Example usage:
//   - /ping : returns as stated above.
func PingGet(p database.StorePinger) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := p.Ping()
		r := "pong"
		if err != nil {
			r = "boom"
		}

		c.JSON(http.StatusOK, gin.H{
			"ping": r,
		})
	}
}
