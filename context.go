package rem

func newContext(request IRequest, handlers []Handler) *Context {
	return &Context{
		request:        request,
		data:           nil,
		handlers:       handlers,
		currentHandler: 0,
	}
}

type Context struct {
	request IRequest
	data    map[string]any

	handlers       []Handler
	currentHandler uint8
}

func (ctx *Context) Request() IRequest {
	return ctx.request
}

func (ctx *Context) SetData(key string, value any) {
	if ctx.data == nil {
		ctx.data = map[string]any{}
	}
	ctx.data[key] = value
}

func (ctx *Context) GetData(key string) any {
	if ctx.data == nil {
		return nil
	}
	return ctx.data[key]
}

func (ctx *Context) Next() IResponse {
	ctx.currentHandler++

	if ctx.currentHandler >= uint8(len(ctx.handlers)) {
		return config.DefaultHandler(ctx)
	} else {
		return ctx.handlers[ctx.currentHandler](ctx)
	}
}

func (ctx *Context) Body(out any) IResponse {
	return ParseRequest(ctx.request, out)
}
