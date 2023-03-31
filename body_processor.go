package rem

type IBodyProcessor interface {
	SerializeResponse(response IResponse) ([]byte, error)
	ParseRequest(request IRequest, out any) IResponse
	// TODO: Add ParseRequest for types implementing validate interface
}

func SerializeResponse(response IResponse) ([]byte, error) {
	return config.BodyProcessor.SerializeResponse(response)
}

func ParseRequest(request IRequest, out any) IResponse {
	return config.BodyProcessor.ParseRequest(request, out)
}
