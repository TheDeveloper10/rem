package rem

import (
	"encoding/json"
	"net/http"
)

// A wrapper of http.ResponseWriter that implements IResponse
type HTTPResponseWriterWrapper struct {
	rw http.ResponseWriter
	statusCode int
	body       []byte
}

func (w *HTTPResponseWriterWrapper) Status(statusCode int) IResponse {
	w.statusCode = statusCode
	return w
}

func (w *HTTPResponseWriterWrapper) Header(key string, value string) IResponse {
	w.rw.Header().Set(key, value)
	return w
}

func (w *HTTPResponseWriterWrapper) Bytes(bytes []byte) IResponse {
	w.body = bytes
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

	return w.Header("Content-Type", "application/json").Bytes(bytes)
}

func (w *HTTPResponseWriterWrapper) flush() {
	if w.body != nil {
		_, err := w.rw.Write(w.body)
		if err != nil {
			panic(err)
		}
	}
	w.rw.WriteHeader(w.statusCode)
}