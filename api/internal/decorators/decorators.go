package decorators

import (
	"fmt"

	routey "github.com/joseph-beck/routey/pkg/router"
)

func Admin() routey.DecoratorFunc {
	return func(f routey.HandlerFunc) routey.HandlerFunc {
		return func(c *routey.Context) {
			fmt.Println("Admin route!")
		}
	}
}
