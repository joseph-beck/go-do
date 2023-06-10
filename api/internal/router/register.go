package router

import "github.com/gin-gonic/gin"

func RegisterPing(r *gin.RouterGroup) {
	for _, route := range ping.RouteInfo {
		switch route.Method {
		case Get:
			r.GET(route.Path + route.Handler, route.HandlerFunc)
		case Post:	
			//
		case Patch:
			//
		case Delete:
			//
		}
	} 
}
