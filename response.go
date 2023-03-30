package rem

type response struct {
	status int
	body   any
}

func (r *response) Status(status int) Response {
	r.status = status
	return r
}

func (r *response) Body(body any) Response {
	r.body = body
	return r
}
