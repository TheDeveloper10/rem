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

// Create a new route
func (r *Router) NewBasicRoute(url string) IRoute {
	route := BasicRoute{
		url:       url,
		endpoints: map[string][]Handler{},
	}
	r.routes = append(r.routes, &route)

	return &route
}

// Handle HTTP requests
func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	response := WrapHTTPResponseWriter(res)
	request := NewBasicRequest(req)

	for _, route := range r.routes {
		status := route.Match(request.GetURL())
		if status {
			route.handle(response, request)
			return
		}
	}

	response.Status(http.StatusNotFound)
}
