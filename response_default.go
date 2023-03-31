package rem

func newResponseDefault() IResponse {
	return &responseDefault{
		body:    nil,
		headers: nil,
	}
}

type responseDefault struct {
	status  int
	body    any
	headers map[string]string
}

func (rd *responseDefault) Status(status int) IResponse {
	rd.status = status
	return rd
}

func (rd *responseDefault) Body(body any) IResponse {
	rd.body = body
	return rd
}

func (rd *responseDefault) Header(key, value string) IResponse {
	if rd.headers == nil {
		rd.headers = map[string]string{}
	}

	rd.headers[key] = value
	return rd
}

func (rd *responseDefault) GetWrittenStatus() int {
	return rd.status
}

func (rd *responseDefault) GetWrittenBody() any {
	return rd.body
}

func (rd *responseDefault) GetWrittenHeaders() map[string]string {
	return rd.headers
}
