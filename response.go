package rem

import (
	"net/http"
)

type IResponse interface {
	Status(statusCode int) IResponse
	Bytes(data []byte) IResponse
	Text(text string) IResponse
	JSON(data interface{}) IResponse
}

// -----------------------------------------
// Response Factories
// -----------------------------------------

func WrapHTTPResponseWriter(rw http.ResponseWriter) IResponse {
	return &HTTPResponseWriterWrapper{ rw }
}
