package rem

import (
	"io"
	"net/http"
	"net/url"
	"path"
)

type IRequest interface {
	GetURL() string
	GetMethod() string
	GetHeaders() http.Header
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
		Headers: req.Header,
		Cookies: req.Cookies(),
		URLParameters: urlParameters,
		QueryParameters: req.URL.Query(),
		Body: req.Body,
	}
}

// -----------------------------------------
// Request Factory Helpers
// -----------------------------------------

// * Borrowed from the net/http package.
// Returns the canonical path for p, eliminating . and .. elements.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		np += "/"
	}

	return np
}