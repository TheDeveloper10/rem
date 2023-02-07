package layer

import (
	"github.com/TheDeveloper10/rem"
)

func URL(req rem.IRequest, res rem.IResponse, outs ...IRequestURL) bool {
	for _, out := range outs {
		err := out.Process(req.GetURLParameters())

		if err != nil {
			res.Status(ErrStatusInvalidURL).JSON(&userFriendlyError{Error: err.Error()})
			return false
		}
	}

	return true
}