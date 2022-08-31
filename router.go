package rem

import (
	"net/http"
)

func CreateDefaultRouter() *Router {
	return &Router {
		routes: []IRoute{},
		IncorrectRoutes: map[int]Handler{},
	}
}

type Router struct {
	routes []IRoute

	IncorrectRoutes map[int]Handler
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

func (r *Router) GetRoute(url string, handlers ...Handler) IRoute {
	return r.NewBasicRoute(url).SetMethods(http.MethodGet).SetHandlers(handlers...)
}

func (r *Router) PostRoute(url string, handlers ...Handler) IRoute {
	return r.NewBasicRoute(url).SetMethods(http.MethodPost).SetHandlers(handlers...)
}

func (r *Router) PatchRoute(url string, handlers ...Handler) IRoute {
	return r.NewBasicRoute(url).SetMethods(http.MethodPatch).SetHandlers(handlers...)
}

func (r *Router) PutRoute(url string, handlers ...Handler) IRoute {
	return r.NewBasicRoute(url).SetMethods(http.MethodPut).SetHandlers(handlers...)
}

func (r *Router) DeleteRoute(url string, handlers ...Handler) IRoute {
	return r.NewBasicRoute(url).SetMethods(http.MethodDelete).SetHandlers(handlers...)
}


// Handle HTTP requests
func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	response := WrapHTTPResponseWriter(res)
	request := NewBasicRequest(req)

	finalStatusCode := http.StatusNotFound
	for _, route := range r.routes {
		status := route.Match(request.GetMethod(), request.GetURL())
		if status == routePerfectMatch {
			route.handle(response, request)
			return
		} else if status == routeMethodMismatch {
			finalStatusCode = http.StatusMethodNotAllowed
		}
	}

	handler, ok := r.IncorrectRoutes[finalStatusCode]
	if !ok {
		response.Status(finalStatusCode)
	} else {
		handler(response, request)
	}
}
