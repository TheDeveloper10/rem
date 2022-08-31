package rem

import "net/http"

type BasicRoute struct {
	url      string
	endpoints map[string][]Handler
}

// This method compares incoming path to the allowed path of the route
func (br *BasicRoute) Match(url string) bool {
	return url == br.url
}

// This method manages the incoming request
func (br *BasicRoute) handle(response IResponse, request IRequest) {
	handlers, ok := br.endpoints[request.GetMethod()]
	if !ok {
		response.Status(http.StatusMethodNotAllowed)
		return
	}

	for _, handler := range handlers {
		status := handler(response, request)
		if !status {
			return
		}
	}
}

func (br *BasicRoute) setHandlers(method string, handlers []Handler) IRoute {
	br.endpoints[method] = handlers
	return br
}

func (br *BasicRoute) GetRoute(handlers ...Handler) IRoute {
	return br.setHandlers(http.MethodGet, handlers)
}

func (br *BasicRoute) PostRoute(handlers ...Handler) IRoute {
	return br.setHandlers(http.MethodPost, handlers)
}

func (br *BasicRoute) PatchRoute(handlers ...Handler) IRoute {
	return br.setHandlers(http.MethodPatch, handlers)
}

func (br *BasicRoute) PutRoute(handlers ...Handler) IRoute {
	return br.setHandlers(http.MethodPut, handlers)
}

func (br *BasicRoute) DeleteRoute(handlers ...Handler) IRoute {
	return br.setHandlers(http.MethodDelete, handlers)
}