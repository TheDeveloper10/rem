package rem

type Handler func(ctx *Context) IResponse
type ErrorHandlerEmpty func() IResponse
type ErrorHandler func(err error) IResponse
