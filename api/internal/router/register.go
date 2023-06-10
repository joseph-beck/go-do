package router

func (r *Router) RegisterRoutes(rs Routes) {
	for _, route := range rs.RouteInfo {
		switch route.Method {
		case Get:
			r.Engine.GET(route.Path+route.Handler, route.HandlerFunc)
		case Post:
			r.Engine.POST(route.Path+route.Handler, route.HandlerFunc)
		case Patch:
			r.Engine.PATCH(route.Path+route.Handler, route.HandlerFunc)
		case Delete:
			r.Engine.DELETE(route.Path+route.Handler, route.HandlerFunc)
		}
	}
}
