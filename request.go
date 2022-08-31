package rem

import "net/http"

type IRequest interface {
	GetURL() string
	GetMethod() string
	GetHeaders() map[string]string
	GetCookies() map[string]string
	GetBody() string
}

// -----------------------------------------
// Request Factories
// -----------------------------------------

func NewBasicRequest(req *http.Request) IRequest {
	return &BasicRequest{
		
	}
}