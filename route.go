package rem

type IRoute interface {
	SetHandlers(handlers ...Handler) IRoute
	SetMethods(methods ...string) IRoute
	Match(method string, url string) int
	handle(response IResponse, request IRequest)
}

const (
	routePerfectMatch = 0
	routeMethodMismatch = 1
	routeMismatch = 2
)