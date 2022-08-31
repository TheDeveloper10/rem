package rem

type IRoute interface {
	Match(method string, url string) int
	Handle(response IResponse, request IRequest)
}