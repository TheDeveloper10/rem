package rem

func newContext(request IRequest, handlers []Handler) *Context {
	return &Context{
		request:        request,
		data:           nil,
		handlers:       handlers,
		handlerCount:   uint8(len(handlers)),
		currentHandler: 0,
	}
}

type Context struct {
	request IRequest
	data    map[string]any

	handlers       []Handler
	handlerCount   uint8
	currentHandler uint8
}

func (ctx *Context) Request() IRequest {
	return ctx.request
}

func (ctx *Context) SetData(key string, value any) {
	if ctx.data == nil {
		ctx.data = map[string]any{key: value}
	} else {
		ctx.data[key] = value
	}
}

func (ctx *Context) GetData(key string) any {
	if ctx.data == nil {
		return nil
	}
	return ctx.data[key]
}

func (ctx *Context) Next() Handler {
	var h Handler
	if ctx.currentHandler >= ctx.handlerCount {
		h = config.DefaultHandler
	} else {
		h = ctx.handlers[ctx.currentHandler]
	}

	ctx.currentHandler++
	return h
}

func (ctx *Context) Body(out any) IResponse {
	return ParseRequest(ctx.request, out)
}
