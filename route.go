package rem

type IRoute interface {
	Match(url string) bool
	handle(response IResponse, request IRequest)

	extractURLParameters(url string) *KeyValue

	Get(handlers ...Handler) IRoute
	Post(handlers ...Handler) IRoute
	Patch(handlers ...Handler) IRoute
	Put(handlers ...Handler) IRoute
	Delete(handlers ...Handler) IRoute
	MultiMethod(methods []string, handlers ...Handler) IRoute
}
