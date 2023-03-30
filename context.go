package rem

type Context struct {
	request Request
	data    map[string]any

	handlers       []Handler
	currentHandler uint8
}

func (ctx *Context) Request() Request {
	return ctx.request
}

func (ctx *Context) Data() map[string]any {
	return ctx.data
}

func (ctx *Context) Next() Handler {
	ctx.currentHandler++

	if ctx.currentHandler >= uint8(len(ctx.handlers)) {
		return settings.DefaultHandler
	} else {
		return ctx.handlers[ctx.currentHandler]
	}
}
