package layer

import (
	"github.com/TheDeveloper10/rem"
)

func Query(req rem.IRequest, res rem.IResponse, outs ...IRequestQuery) bool {
	for _, out := range outs {
		err := out.Process(req.GetQueryParameters())

		if err != nil {
			res.Status(ErrStatusInvalidQuery).JSON(&userFriendlyError{Error: err.Error()})
			return false
		}
	}

	return true
}