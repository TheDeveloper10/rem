package rem

import (
	"encoding/json"
)

func DefaultJSONBodyProcessor() IBodyProcessor {
	return &JSONBodyProcessor{
		UnsupportedContentType: func() IResponse { return UnsupportedMediaType(nil) },
		NoBody:                 func() IResponse { return BadRequest("No Body") },
		InvalidBody:            func(err error) IResponse { return BadRequest("Invalid Body") },
	}
}

type JSONBodyProcessor struct {
	UnsupportedContentType ErrorHandlerEmpty
	NoBody                 ErrorHandlerEmpty
	InvalidBody            ErrorHandler
}

func (bp *JSONBodyProcessor) SerializeResponse(response IResponse) ([]byte, error) {
	response.Header("Content-Type", "application/json")
	return json.Marshal(response.GetWrittenBody())
}

func (bp *JSONBodyProcessor) ParseRequest(req IRequest, out any) IResponse {
	ct := req.Headers().Get("Content-Type")
	if ct != "application/json" {
		return bp.UnsupportedContentType()
	}

	body := req.Body()
	if body == nil {
		return bp.NoBody()
	}

	err := json.NewDecoder(body).Decode(out)
	if err != nil {
		return bp.InvalidBody(err)
	}

	return nil
}
