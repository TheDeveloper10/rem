package rem

import "net/http"

func Response(status int, body any) IResponse {
	return newResponseDefault().Status(status).Body(body)
}

func Success(body any) IResponse {
	return Response(http.StatusOK, body)
}

func BadRequest(body any) IResponse {
	return Response(http.StatusBadRequest, body)
}

func ForbiddenAccess(body any) IResponse {
	return Response(http.StatusForbidden, body)
}

func UnsupportedMediaType(body any) IResponse {
	return Response(http.StatusUnsupportedMediaType, body)
}
