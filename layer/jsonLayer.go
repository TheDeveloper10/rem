package layer

import (
	"encoding/json"
	"github.com/TheDeveloper10/rem"
)

func JSON(req rem.IRequest, res rem.IResponse, out IRequestBody) bool {
	if req.GetHeaders().Get("Content-Type") != "application/json" {
		res.Status(ErrStatusContentTypeNotJson)
		return false
	}

	if req.GetBody() == nil {
		res.Status(ErrStatusNoBody).JSON(&userFriendlyError{Error: ErrMsgNoBody})
		return false
	}

	err := json.NewDecoder(req.GetBody()).Decode(&out)
	if err != nil {
		res.Status(ErrStatusInvalidJson).JSON(&userFriendlyError{Error: ErrMsgInvalidJson})
		return false
	}

	err = out.Validate()
	if err != nil {
		res.Status(ErrStatusInvalidBody).JSON(&userFriendlyError{Error: err.Error()})
		return false
	}

	return true
}
