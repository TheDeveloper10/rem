package rem

func newResponseDefault() IResponse {
	return &responseDefault{}
}

type responseDefault struct {
	status int
	body   any
}

func (rd *responseDefault) Status(status int) IResponse {
	rd.status = status
	return rd
}

func (rd *responseDefault) Body(body any) IResponse {
	rd.body = body
	return rd
}

func (rd *responseDefault) GetWrittenStatus() int {
	return rd.status
}

func (rd *responseDefault) GetWrittenBody() any {
	return rd.body
}
