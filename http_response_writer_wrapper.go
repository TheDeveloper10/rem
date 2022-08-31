package rem

import (
	"encoding/json"
	"net/http"
)

// A wrapper of http.ResponseWriter that implements IResponse
type HTTPResponseWriterWrapper struct {
	rw http.ResponseWriter
}

func (w *HTTPResponseWriterWrapper) Status(statusCode int) IResponse {
	w.rw.WriteHeader(statusCode)
	return w
}

func (w *HTTPResponseWriterWrapper) Bytes(bytes []byte) IResponse {
	_, err := w.rw.Write(bytes)
	if err != nil {
		panic(err)
	}
	return w
}

func (w *HTTPResponseWriterWrapper) Text(text string) IResponse {
	return w.Bytes([]byte(text))
}

func (w *HTTPResponseWriterWrapper) JSON(data interface{}) IResponse {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return w.Bytes(bytes)
}