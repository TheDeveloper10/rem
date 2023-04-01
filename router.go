package rem

import "net/http"

// Create a new router
func NewRouter() *Router {
	return &Router{
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
func (r *Router) NewRoute(URL string) IRoute {
	return r.AddRoute(newRouteDefault(URL))
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	targetURL := cleanPath(req.URL.EscapedPath())

	for _, route := range r.routes {
		if !route.Match(targetURL) {
			continue
		}

		handlers := route.Handlers(req.Method)
		if handlers == nil {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// TODO: improve this part
		paramsURL, matched := route.ExtractURLParameters(targetURL)
		if !matched {
			break
		}

		request := wrapHTTPRequest(req, paramsURL)
		ctx := newContext(request, handlers)
		response := ctx.Next()

		// write status
		rw.WriteHeader(response.GetWrittenStatus())

		// set headers (if any exist)
		for header, value := range response.GetWrittenHeaders() {
			rw.Header().Set(header, value)
		}

		// if the response doesn't have a body
		if rBody := response.GetWrittenBody(); rBody == nil {
			return
		}

		// serialize the body
		data, err := config.BodyProcessor.SerializeResponse(response)
		if err != nil {
			panic(err)
		}

		// write it to ResponseWriter
		_, err = rw.Write(data)
		if err != nil {
			panic(err)
		}

		return
	}

	rw.WriteHeader(http.StatusNotFound)
}

func (r *Router) addHandlersToRoute(URL string, method string, handlers []Handler) *Router {
	var routeToUse IRoute
	cleanedURL := cleanPath(URL)
	for _, route := range r.routes {
		if route.URL() == cleanedURL {
			routeToUse = route
			break
		}
	}
	if routeToUse == nil {
		routeToUse = r.AddRoute(newRouteDefault(URL))
	}

	routeToUse.SetHandlers(method, handlers)
	return r
}

func (r *Router) Get(URL string, handlers ...Handler) *Router {
	return r.addHandlersToRoute(URL, http.MethodGet, handlers)
}
func (r *Router) Post(URL string, handlers ...Handler) *Router {
	return r.addHandlersToRoute(URL, http.MethodPost, handlers)
}
func (r *Router) Patch(URL string, handlers ...Handler) *Router {
	return r.addHandlersToRoute(URL, http.MethodPatch, handlers)
}
func (r *Router) Put(URL string, handlers ...Handler) *Router {
	return r.addHandlersToRoute(URL, http.MethodPut, handlers)
}
func (r *Router) Delete(URL string, handlers ...Handler) *Router {
	return r.addHandlersToRoute(URL, http.MethodDelete, handlers)
}
func (r *Router) Connect(URL string, handlers ...Handler) *Router {
	return r.addHandlersToRoute(URL, http.MethodConnect, handlers)
}
func (r *Router) Head(URL string, handlers ...Handler) *Router {
	return r.addHandlersToRoute(URL, http.MethodHead, handlers)
}
func (r *Router) Options(URL string, handlers ...Handler) *Router {
	return r.addHandlersToRoute(URL, http.MethodOptions, handlers)
}
func (r *Router) Trace(URL string, handlers ...Handler) *Router {
	return r.addHandlersToRoute(URL, http.MethodTrace, handlers)
}
