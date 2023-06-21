package app

import (
	"go-do/internal/pinghandler"
	"go-do/internal/router"
	"go-do/internal/todohandler"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// The makeRoutes method returns a pointer to a Routes struct,
// which stores a slice of Route.
func makeRoutes() *router.Routes {
	return &router.Routes{
		RouteInfo: []router.Route{
			// Ping Handlers
			{
				Name:        "GetPing",
				Method:      router.Get,
				Path:        "/ping",
				Handler:     "",
				HandlerFunc: pinghandler.PingGet(r.Store),
			},
			// Todo Handlers
			{
				Name:        "GetTodo",
				Method:      router.Get,
				Path:        "/todo",
				Handler:     "",
				HandlerFunc: todohandler.TodoGet(r.Store),
			},
			{
				Name:        "PostTodo",
				Method:      router.Post,
				Path:        "/todo",
				Handler:     "",
				HandlerFunc: todohandler.TodoPost(r.Store),
			},
			{
				Name:        "PatchTodo",
				Method:      router.Patch,
				Path:        "/todo",
				Handler:     "",
				HandlerFunc: todohandler.TodoPatch(r.Store),
			},
			{
				Name:        "DeleteTodo",
				Method:      router.Delete,
				Path:        "/todo",
				Handler:     "",
				HandlerFunc: todohandler.TodoDelete(r.Store),
			},
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
