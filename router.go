package rem

import (
	"net/http"
)

type Router struct {
	routes []IRoute
}

// Create a new route
func (r *Router) NewBasicRoute(url string) IRoute {
	route := BasicRoute{
		url:      url,
		handlers: []Handler{},
		methods:  []string{},
	}
	r.routes = append(r.routes, &route)

	return &route
}

// Handle HTTP requests
func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	//response := WrapHTTPResponseWriter(res)
	//request := NewBasicRequest(req)
	//
	//req.URL
}
