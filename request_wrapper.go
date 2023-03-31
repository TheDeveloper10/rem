package rem

import (
	"io"
	"net/http"
)

func wrapHTTPRequest(req *http.Request) IRequest {
	return &requestWrapper{req: req}
}

type requestWrapper struct {
	req *http.Request
}

func (rw *requestWrapper) Body() io.Reader {
	return rw.req.Body
}

// Cookies implements IRequest
func (*requestWrapper) Cookies() {
	panic("unimplemented")
}

// Headers implements IRequest
func (*requestWrapper) Headers() {
	panic("unimplemented")
}

// IP implements IRequest
func (*requestWrapper) IP() {
	panic("unimplemented")
}

func (rw *requestWrapper) Method() string {
	return rw.req.Method
}

func (rw *requestWrapper) URL() string {
	return cleanPath(rw.req.URL.EscapedPath())
}
