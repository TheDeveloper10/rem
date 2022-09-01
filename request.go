package rem

import (
	"io"
	"net/http"
	"net/url"
)

type IRequest interface {
	GetURL() string
	GetMethod() string
	GetHeaders() Headers
	GetCookies() []*http.Cookie
	GetURLParameters() map[string]string
	setURLParameters(*map[string]string)
	GetQueryParameters() url.Values
	GetBody() io.ReadCloser
}

// -----------------------------------------
// Request Factories
// -----------------------------------------

func NewBasicRequest(req *http.Request) IRequest {
	urlParametersMaster := req.Context().Value(0)
	var urlParameters map[string]string = nil
	if urlParametersMaster != nil {
		urlParameters = urlParametersMaster.(map[string]string)
	}
	
	return &BasicRequest{
		URL: cleanPath(req.URL.EscapedPath()),
		Method: req.Method,
		Headers: Headers(req.Header),
		Cookies: req.Cookies(),
		URLParameters: urlParameters,
		QueryParameters: req.URL.Query(),
		Body: req.Body,
	}
}
