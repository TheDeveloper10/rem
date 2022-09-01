package rem

import (
	"io"
	"net/http"
)

type IRequest interface {
	GetURL() string
	GetMethod() string
	GetHeaders() KeyValues
	GetCookies() []*http.Cookie
	GetURLParameters() KeyValue
	setURLParameters(*KeyValue)
	GetQueryParameters() KeyValues
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
		URL:             cleanPath(req.URL.EscapedPath()),
		Method:          req.Method,
		Headers:         KeyValues(req.Header),
		Cookies:         req.Cookies(),
		URLParameters:   urlParameters,
		QueryParameters: KeyValues(req.URL.Query()),
		Body:            req.Body,
	}
}
