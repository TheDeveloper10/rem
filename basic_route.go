package rem

type BasicRoute struct {
	url      string
	handlers []Handler
	methods  []string
}

// This method sets all handlers on this route (that may include middlewares and endpoint functions)
func (br *BasicRoute) SetHandlers(handlers ...Handler) IRoute {
	br.handlers = handlers
	return br
}

// This method sets all allowed methods on this route
func (br *BasicRoute) SetMethods(methods ...string) IRoute {
	br.methods = methods
	return br
}

// This method compares incoming path to the allowed path of the route
func (br *BasicRoute) Match(method string, url string) int {
	if url != br.url {
		return routeMismatch
	}

	methodMatch := false
	for _, m := range br.methods {
		if m == method {
			methodMatch = true
			break
		}
	}
	if methodMatch {
		return routePerfectMatch
	} else {
		return routeMethodMismatch
	}
}

// This method manages the incoming request
func (br *BasicRoute) handle(response IResponse, request IRequest) {
	for _, handler := range br.handlers {
		status := handler(response, request)
		if !status {
			return
		}
	}
}