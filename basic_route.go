package rem

import "net/http"

type BasicRoute struct {
	url      string
	handlers []Handler
	methods  []string
}

// AddHandle can be used for adding middlewares or just the endpoint handling function
func (r *BasicRoute) AddHandle(handler Handler) *BasicRoute {
	r.handlers = append(r.handlers, handler)
	return r
}

func (r *BasicRoute) Match(method string, url string) int {
	methodMatch := false
	for _, m := range r.methods {
		if m == method {
			methodMatch = true
			break
		}
	}
	if !methodMatch {
		return http.StatusMethodNotAllowed
	}

	return http.StatusOK
}

func (r *BasicRoute) Handle(response IResponse, request IRequest) {

}