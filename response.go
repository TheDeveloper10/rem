package rem

import (
	"net/http"
)

type IResponse interface {
	Status(statusCode int) IResponse
	Header(key string, value string) IResponse
	Bytes(data []byte) IResponse
	Text(text string) IResponse
	JSON(data interface{}) IResponse
}

// -----------------------------------------
// Response Factories
// -----------------------------------------

func WrapHTTPResponseWriter(rw http.ResponseWriter) *HTTPResponseWriterWrapper {
	return &HTTPResponseWriterWrapper{
		rw: 		rw,
		statusCode: http.StatusInternalServerError,
		body: 	 	nil,
	}
}

func NewMockResponse() *MockResponse {
	return &MockResponse{
		StatusCode: 0,
		Headers: map[string][]string{},
		Body: []byte{},
	}
}