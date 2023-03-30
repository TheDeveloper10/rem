package rem

import "net/http"

func Success(body any) Response {
	res := &response{}
	return res.Status(http.StatusOK).Body(body)
}

func ForbiddenAccess(body any) Response {
	res := &response{}
	return res.Status(http.StatusForbidden).Body(body)
}
