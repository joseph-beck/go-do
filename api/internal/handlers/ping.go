package handlers

import (
	"go-do/internal/database"
	"net/http"

	routey "github.com/joseph-beck/routey/pkg/router"
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
func PingGet(s *database.Store) routey.HandlerFunc {
	return func(c *routey.Context) {
		err := s.Ping()
		r := "pong"
		if err != nil {
			r = "boom"
		}

		c.JSON(http.StatusOK, routey.M{
			"ping": r,
		})
	}
}
