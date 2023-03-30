package rem

import "io"

type Request interface {
	Body() io.Reader
	Headers() // TODO
	Cookies() // TODO
	IP()      // TODO
}

type Response interface {
	Status(status int) Response
	Body(body any) Response
}

type BodyProcessor interface {
	Serialize(body any) ([]byte, error)
	Parse(body io.Reader, out any) error
}

type Handler func(ctx *Context) Response
