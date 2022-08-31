package rem

import (
	"io"
	"net/http"
	"net/url"
)

type BasicRequest struct {
	URL string
	Method string
	Headers http.Header
	Cookies []*http.Cookie
	URLParameters map[string]string
	QueryParameters url.Values
	Body io.ReadCloser
}

func (br *BasicRequest) GetURL() string { return br.URL }
func (br *BasicRequest) GetMethod() string { return br.Method }
func (br *BasicRequest) GetHeaders() http.Header { return br.Headers }
func (br *BasicRequest) GetCookies() []*http.Cookie { return br.Cookies }
func (br *BasicRequest) GetURLParameters() map[string]string { return br.URLParameters }
func (br *BasicRequest) setURLParameters(params *map[string]string) { br.URLParameters = *params }
func (br *BasicRequest) GetQueryParameters() url.Values { return br.QueryParameters }
func (br *BasicRequest) GetBody() io.ReadCloser { return br.Body }
