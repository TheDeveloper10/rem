package rem

type BasicRoute struct {
	url      string
	handlers []Handler
	methods  []string
}

// AddHandle can be used for adding middlewares or just the endpoint handling function
func (br *BasicRoute) AddHandle(handler Handler) *BasicRoute {
	br.handlers = append(br.handlers, handler)
	return br
}

func (br *BasicRoute) Match(method string, url string) int {
	methodMatch := false
	for _, m := range br.methods {
		if m == method {
			methodMatch = true
			break
		}
	}
	if !methodMatch {
		return routeMethodMismatch
	}

	if url == br.url {
		return routePerfectMatch
	} else {
		return routeMismatch
	}
}

func (br *BasicRoute) Handle(response IResponse, request IRequest) {
	for _, handler := range br.handlers {
		status := handler(response, request)
		if !status {
			return
		}
	}
}