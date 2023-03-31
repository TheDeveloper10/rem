package rem

import "io"

type IRequest interface {
	Body() io.Reader
	Headers() // TODO
	Cookies() // TODO
	IP()      // TODO
}

type IResponse interface {
	Status(status int) IResponse
	Body(body any) IResponse

	GetWrittenStatus() int
	GetWrittenBody() any
}

type WrittenData struct {
	Status int
	Body   any
}

type IBodyProcessor interface {
	Serialize(body any) ([]byte, error)
	Parse(body io.Reader, out any) error
}

type Handler func(ctx *Context) IResponse

type IRoute interface {
	Match(url string) bool
	GetHandlers() []Handler

	Get(handlers ...Handler) IRoute
	Post(handlers ...Handler) IRoute
	Patch(handlers ...Handler) IRoute
	Put(handlers ...Handler) IRoute
	Delete(handlers ...Handler) IRoute
	MultiMethod(methods []string, handlers ...Handler) IRoute
}
