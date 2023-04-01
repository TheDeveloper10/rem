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
	QueryVars() MapStringsReadOnly
	URLVars() MapStringReadOnly
	URL() string
	Method() string

	OriginalRequest() *http.Request
}

type IResponse interface {
	Status(status int) IResponse
	Body(body any) IResponse
	Header(key, value string) IResponse

	GetWrittenStatus() int
	GetWrittenBody() any
	GetWrittenHeaders() map[string]string
}

type IRoute interface {
	Match(targetURL string) bool
	ExtractURLParameters(targetURL string) (map[string]string, bool)
	Handlers(method string) []Handler
	SetHandlers(method string, handlers []Handler) IRoute
	URL() string

	Get(handlers ...Handler) IRoute
	Head(handlers ...Handler) IRoute
	Post(handlers ...Handler) IRoute
	Put(handlers ...Handler) IRoute
	Delete(handlers ...Handler) IRoute
	Connect(handlers ...Handler) IRoute
	Options(handlers ...Handler) IRoute
	Trace(handlers ...Handler) IRoute
	Patch(handlers ...Handler) IRoute
	MultiMethod(methods []string, handlers ...Handler) IRoute
}
