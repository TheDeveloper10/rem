package rem

import "net/http"

type Router struct {
	routes []Route
}

// Create a new route
func (r *Router) NewRoute() *Route {
	return &Route{
		handlers: []Handler{},
		methods:  []string{},
	}
}

// Handle HTTP requests
func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}