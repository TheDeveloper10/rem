package rem

import "net/http"

func Success(body any) IResponse {
	return newResponseDefault().Status(http.StatusOK).Body(body)
}

func BadRequest(body any) IResponse {
	return newResponseDefault().Status(http.StatusBadRequest).Body(body)
}

func ForbiddenAccess(body any) IResponse {
	return newResponseDefault().Status(http.StatusForbidden).Body(body)
}

func UnsupportedMediaType(body any) IResponse {
	return newResponseDefault().Status(http.StatusUnsupportedMediaType).Body(body)
}
