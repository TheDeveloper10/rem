package rem

type Route struct {
	handlers []Handler
	methods  []string
}

// AddHandle can be used for adding middlewares or just the endpoint handling function
func (r *Route) AddHandle(handler Handler) *Route {
	r.handlers = append(r.handlers, handler)
	return r
}
