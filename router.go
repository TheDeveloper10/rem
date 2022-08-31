package rem

import (
	"net/http"
)

func CreateDefaultRouter() *Router {
	return &Router {
		routes: []IRoute{},
	}
}

type Router struct {
	routes []IRoute
}

// Add a new route to the router
func (r *Router) AddRoute(route IRoute) IRoute {
	r.routes = append(r.routes, route)
	return route
}

// Create and add a new basic route
func (r *Router) NewBasicRoute(url string) IRoute {
	return r.AddRoute(NewBasicRoute(url))
}

// Handle HTTP requests
func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	response := WrapHTTPResponseWriter(res)
	request := NewBasicRequest(req)

	targetURL := request.GetURL()
	for _, route := range r.routes {
		status := route.Match(targetURL)
		if status {
			request.setURLParameters(route.extractURLParameters(targetURL))
			route.handle(response, request)
			return
		}
	}

	response.Status(http.StatusNotFound)
}
