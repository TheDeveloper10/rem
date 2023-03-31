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

func (rw *requestWrapper) BodyBytes() ([]byte, error) {
	return io.ReadAll(rw.req.Body)
}

func (rw *requestWrapper) GetCookie(name string) (*http.Cookie, error) {
	return rw.req.Cookie(name)
}

func (rw *requestWrapper) Cookies() []*http.Cookie {
	return rw.req.Cookies()
}

// func (rw *requestWrapper) GetHeader() {
// 	return rw.req.Header.Get
// }

func (rw *requestWrapper) Headers() http.Header {
	return rw.req.Header
}

func (rw *requestWrapper) RemoteAddress() string {
	return rw.req.RemoteAddr
}

func (rw *requestWrapper) Method() string {
	return rw.req.Method
}

func (rw *requestWrapper) URL() string {
	return cleanPath(rw.req.URL.EscapedPath())
}
