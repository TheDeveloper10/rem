package rem

type IRoute interface {
	Match(method string, url string) int
	Handle(response IResponse, request IRequest)
}

const (
	routePerfectMatch = 0
	routeMethodMismatch = 1
	routeMismatch = 2
)