package app

import (
	"go-do/internal/handlers"
	"go-do/internal/router"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// The makeRoutes method returns a pointer to a Routes struct,
// which stores a slice of Route.
func makeRoutes() router.Routes {
	return router.Routes{
		RouteInfo: []router.Route{
			// Ping Handlers
			{
				Name:        "Get Ping",
				Method:      router.Get,
				Path:        "/ping",
				Handler:     "",
				HandlerFunc: handlers.PingGet(s),
			},
			// Task Handlers
			{
				Name:        "List Task",
				Method:      router.Get,
				Path:        "/task",
				Handler:     "/:list",
				HandlerFunc: handlers.TaskList(s),
			},
			{
				Name:        "Get Task",
				Method:      router.Get,
				Path:        "/task",
				Handler:     "/:list/:task",
				HandlerFunc: handlers.TaskGet(s),
			},
			{
				Name:        "Post Task",
				Method:      router.Post,
				Path:        "/task",
				Handler:     "/:list",
				HandlerFunc: handlers.TaskPost(s),
			},
			{
				Name:        "Put Task",
				Method:      router.Put,
				Path:        "/task",
				Handler:     "/:list",
				HandlerFunc: handlers.TaskPut(s),
			},
			{
				Name:        "Patch Task",
				Method:      router.Patch,
				Path:        "/task",
				Handler:     "/:list",
				HandlerFunc: handlers.TaskPatch(s),
			},
			{
				Name:        "Delete Task",
				Method:      router.Delete,
				Path:        "/task",
				Handler:     "/:list/:task",
				HandlerFunc: handlers.TaskDelete(s),
			},
			// User Handlers
		},
	}
}

func reverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, _ := url.Parse("http://localhost:3000")
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL = c.Request.URL
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
