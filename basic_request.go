package rem

type BasicRequest struct {
	URL string
	Method string
	Headers map[string]string
	Cookies map[string]string
	Body string
}

func (br BasicRequest) GetURL() string {
	return br.URL
}

func (br BasicRequest) GetMethod() string {
	return br.Method
}

func (br BasicRequest) GetHeaders() map[string]string {
	return br.Headers
}

func (br BasicRequest) GetCookies() map[string]string {
	return br.Cookies
}

func (br BasicRequest) GetBody() string {
	return br.Body
}
