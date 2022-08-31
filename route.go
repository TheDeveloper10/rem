package rem

type IRoute interface {
	Match(url string) bool
	handle(response IResponse, request IRequest)

	GetRoute(handlers ...Handler) IRoute
	PostRoute(handlers ...Handler) IRoute
	PatchRoute(handlers ...Handler) IRoute
	PutRoute(handlers ...Handler) IRoute
	DeleteRoute(handlers ...Handler) IRoute
}
