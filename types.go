package rem

import (
	"io"
	"net/http"
)

type IRequest interface {
	Body() io.Reader
	BodyBytes() ([]byte, error)

	// TODO: add TLS
	// TODO: add Forms
	// TODO: add url variables
	// TODO: add query variables
	Headers() http.Header
	Cookies() []*http.Cookie
	RemoteAddress() string
	URL() string
	Method() string
}

type IResponse interface {
	Status(status int) IResponse
	Body(body any) IResponse
	Header(key, value string) IResponse

	GetWrittenStatus() int
	GetWrittenBody() any
	GetWrittenHeaders() map[string]string
}

type Handler func(ctx *Context) IResponse
type ErrorHandlerEmpty func() IResponse
type ErrorHandler func(err error) IResponse

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
